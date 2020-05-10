package httpu

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/huin/goupnp/v2/errkind"
)

// Client is a client for dealing with HTTPU (HTTP over UDP). Its typical
// function is for HTTPMU, and particularly SSDP.
type Client struct {
	connLock sync.Mutex // Protects use of conn.
	conn     net.PacketConn
}

// NewClient creates a new Client, opening up a new UDP socket for the purpose.
func NewClient(ctx context.Context) (*Client, error) {
	lc := &net.ListenConfig{}

	conn, err := lc.ListenPacket(ctx, "udp", ":0")
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

// NewClientAddr creates a new Client which will broadcast packets from the
// specified address, opening up a new UDP socket for the purpose
func NewClientAddr(ctx context.Context, addr string) (*Client, error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return nil, errkind.New(
			errkind.InvalidArgument, "invalid listening address %q", addr,
		)
	}

	lc := &net.ListenConfig{}

	conn, err := lc.ListenPacket(ctx, "udp", ip.String()+":0")
	if err != nil {
		return nil, errkind.Network.Wrap(err)
	}
	return &Client{conn: conn}, nil
}

// Close shuts down the client. The client will no longer be useful following
// this.
func (httpu *Client) Close() error {
	httpu.connLock.Lock()
	defer httpu.connLock.Unlock()
	return httpu.conn.Close()
}

// Do performs a request. An error is only returned for failing to send the
// request. Failures in receipt simply do not add to the resulting responses.
//
// By default it sends 2 requests, and waits 3 seconds for responses to them.
//
// Note that at present only one concurrent request will happen per Client.
func (httpu *Client) Do(
	req *http.Request,
	options ...RequestOption,
) ([]*http.Response, error) {
	ctx := req.Context()

	now := time.Now()
	rs := &requestSettings{
		numSends:  2,
		now:       now,
		waitUntil: now.Add(time.Second * 3),
	}
	rs.applyOptions(options)

	if ctxDeadline, ok := ctx.Deadline(); ok {
		if ctxDeadline.Before(rs.waitUntil) {
			rs.waitUntil = ctxDeadline
		}
	}

	// Create the request. This is a subset of what http.Request.Write does
	// deliberately to avoid creating extra fields which may confuse some
	// devices.
	var requestBuf bytes.Buffer
	method := req.Method
	if method == "" {
		method = "GET"
	}
	fmt.Fprintf(&requestBuf, "%s %s HTTP/1.1\r\n", method, req.URL.RequestURI())
	if err := req.Header.Write(&requestBuf); err != nil {
		return nil, err
	}
	requestBuf.Write([]byte{'\r', '\n'})

	destAddr, err := net.ResolveUDPAddr("udp", req.Host)
	if err != nil {
		return nil, errkind.Network.Wrap(err)
	}

	return httpu.doInternal(destAddr, req, requestBuf.Bytes(), rs)
}

func (httpu *Client) doInternal(
	destAddr net.Addr,
	req *http.Request,
	reqBytes []byte,
	rs *requestSettings,
) ([]*http.Response, error) {
	httpu.connLock.Lock()
	defer httpu.connLock.Unlock()

	if err := httpu.conn.SetDeadline(rs.waitUntil); err != nil {
		return nil, err
	}

	// Send request.
	for i := 0; i < rs.numSends; i++ {
		if n, err := httpu.conn.WriteTo(reqBytes, destAddr); err != nil {
			return nil, errkind.Wrap(
				errkind.Network, err, "sending HTTPU request",
			)
		} else if n < len(reqBytes) {
			return nil, errkind.New(
				errkind.Network,
				"wrote %d bytes rather than full %d in request",
				n, len(reqBytes),
			)
		}
		time.Sleep(5 * time.Millisecond)
	}

	// Await responses until timeout.
	var responses []*http.Response
	responseBytes := make([]byte, 2048)
	for {
		// 2048 bytes should be sufficient for most networks.
		n, _, err := httpu.conn.ReadFrom(responseBytes)
		if err != nil {
			if err, ok := err.(net.Error); ok {
				if err.Timeout() {
					// Timeout reached - return discovered responses.
					return responses, nil
				}
				if err.Temporary() {
					// Sleep in case this is a persistent error to avoid pegging
					// CPU until deadline.
					time.Sleep(10 * time.Millisecond)
					continue
				}
			}
			return nil, err
		}

		// Parse response.
		response, err := http.ReadResponse(
			bufio.NewReader(bytes.NewBuffer(responseBytes[:n])),
			req,
		)
		if err != nil {
			log.Printf("httpu: error while parsing response: %v", err)
			continue
		}

		responses = append(responses, response)
	}
}

type requestSettings struct {
	numSends  int
	now       time.Time
	waitUntil time.Time
}

func (rs *requestSettings) applyOptions(options []RequestOption) {
	for _, o := range options {
		o(rs)
	}
}

// RequestOption is for types that tune HTTPU requests.
type RequestOption func(*requestSettings)

// NumSends controls how many redundant requests to send.
func NumSends(numSends int) RequestOption {
	return func(rs *requestSettings) {
		rs.numSends = numSends
	}
}

// WaitFor controls how long to wait for HTTPU responses.
func WaitFor(d time.Duration) RequestOption {
	return func(rs *requestSettings) {
		rs.waitUntil = rs.now.Add(d)
	}
}