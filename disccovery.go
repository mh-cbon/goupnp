package goupnp

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/huin/goupnp/httpu"
	"github.com/huin/goupnp/ssdp"
)

type clientProvider func() (httpu.ClientInterface, func(), error)

type Discovery struct {
	Client         clientProvider
	MaxWaitSeconds int
	NumSends       int
}

func (d *Discovery) DiscoverDevicesCtx(ctx context.Context, searchTargets ...string) ([]MaybeRootDevice, error) {
	hc, hcCleanup, err := d.Client()
	if err != nil {
		return nil, err
	}
	defer hcCleanup()
	maxWaitSeconds := d.MaxWaitSeconds
	if maxWaitSeconds < 1 {
		maxWaitSeconds = 2
	}
	numSends := d.NumSends
	if numSends < 1 {
		numSends = 3
	}
	var results []MaybeRootDevice
	for _, searchTarget := range searchTargets {
		responses, err := ssdp.SSDPRawSearchCtx(ctx, hc, string(searchTarget), maxWaitSeconds, numSends)
		if err != nil {
			return nil, err
		}

		res := make([]MaybeRootDevice, len(responses))
		for i, response := range responses {
			maybe := &res[i]
			maybe.SearchTarget = searchTarget
			maybe.USN = response.Header.Get("USN")
			loc, err := response.Location()
			if err != nil {
				maybe.Err = ContextError{"unexpected bad location from search", err}
				continue
			}
			maybe.Location = loc
			if root, err := DeviceByURLCtx(ctx, loc); err != nil {
				maybe.Err = err
			} else {
				maybe.Root = root
			}
			if i := response.Header.Get(httpu.LocalAddressHeader); len(i) > 0 {
				maybe.LocalAddr = net.ParseIP(i)
			}
		}
		results = append(results, res...)
	}

	return results, nil
}

// DiscoverDevices is the legacy version of DiscoverDevicesCtx, but uses
// context.Background() as the context.
func (d *Discovery) DiscoverDevices(searchTargets ...string) ([]MaybeRootDevice, error) {
	return d.DiscoverDevicesCtx(context.Background(), searchTargets...)
}

// NewServiceClientsCtx discovers services, and returns clients for them. err will
// report any error with the discovery process (blocking any device/service
// discovery), errors reports errors on a per-root-device basis.
func (d *Discovery) NewServiceClientsCtx(ctx context.Context, searchTargets ...string) (clients MultiServiceClient, errors []error, err error) {
	var maybeRootDevices []MaybeRootDevice
	if maybeRootDevices, err = d.DiscoverDevicesCtx(ctx, searchTargets...); err != nil {
		return
	}

	clients = make([]ServiceClient, 0, len(maybeRootDevices))
	uniq := map[string]ServiceClient{}
	for _, maybeRootDevice := range maybeRootDevices {
		if maybeRootDevice.Err != nil {
			errors = append(errors, maybeRootDevice.Err)
			continue
		}

		deviceClients, err := newServiceClientsFromRootDevice(
			maybeRootDevice.Root,
			maybeRootDevice.Location,
			maybeRootDevice.SearchTarget,
			maybeRootDevice.LocalAddr,
		)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		for _, deviceClient := range deviceClients {
			id := fmt.Sprintf("%v:%v-%v.%v",
				deviceClient.Service.ServiceId,
				deviceClient.Service.ServiceType,
				deviceClient.RootDevice.Device.UDN,
				deviceClient.LocalAddr(),
			)
			if _, ok := uniq[id]; ok {
				continue
			}
			uniq[id] = deviceClient
			clients = append(clients, deviceClient)
		}
	}

	return
}

// NewServiceClients is the legacy version of NewServiceClientsCtx, but uses
// context.Background() as the context.
func (d *Discovery) NewServiceClients(searchTarget string) (clients MultiServiceClient, errors []error, err error) {
	return d.NewServiceClientsCtx(context.Background(), searchTarget)
}

func AutoDiscover(also ...string) clientProvider {
	return func() (httpu.ClientInterface, func(), error) {
		addrs, err := localIPv4MCastAddrs()
		if err != nil {
			return nil, nil, ctxError(err, "requesting host IPv4 addresses")
		}
		return DiscoverFromAddrs(append(addrs, also...)...)()
	}
}

func DiscoverFromAddrs(addrs ...string) clientProvider {
	return func() (httpu.ClientInterface, func(), error) {

		closers := make([]io.Closer, 0, len(addrs))
		delegates := make([]httpu.ClientInterface, 0, len(addrs))
		for _, addr := range addrs {
			c, err := httpu.NewHTTPUClientAddr(addr)
			if err != nil {
				return nil, nil, ctxErrorf(err,
					"creating HTTPU client for address %s", addr)
			}
			closers = append(closers, c)
			delegates = append(delegates, c)
		}

		closer := func() {
			for _, c := range closers {
				c.Close()
			}
		}

		return httpu.NewMultiClient(delegates), closer, nil
	}
}

var DefaultDiscoveryClient = Discovery{
	Client: AutoDiscover(),
}

type MultiServiceClient []ServiceClient

func (m MultiServiceClient) Filter(f ServiceClientFilter, andAlso ...ServiceClientFilter) MultiServiceClient {
	all := f.And(andAlso...)
	out := make(MultiServiceClient, 0, len(m))
	for _, sc := range m {
		if all(sc) {
			out = append(out, sc)
		}
	}
	return out
}
func (m MultiServiceClient) Has(f ServiceClientFilter, andAlso ...ServiceClientFilter) bool {
	all := f.And(andAlso...)
	for _, sc := range m {
		if all(sc) {
			return true
		}
	}
	return false
}

type ServiceClientFilter func(ServiceClient) bool

func (f ServiceClientFilter) And(filters ...ServiceClientFilter) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		if !f(sc) {
			return false
		}
		for _, filter := range filters {
			if !filter(sc) {
				return false
			}
		}
		return true
	}
}
func (f ServiceClientFilter) Or(filters ...ServiceClientFilter) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		if f(sc) {
			return true
		}
		for _, filter := range filters {
			if filter(sc) {
				return true
			}
		}
		return false
	}
}

func ServiceLocalAddr(addr string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		return sc.LAddr.String() == addr
	}
}

func NotServiceLocalAddr(addr string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		return sc.LAddr.String() != addr
	}
}

func ServiceHost(host string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		return host == sc.Location.Host
	}
}

func ServiceID(id string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		return sc.Service.ServiceId == id
	}
}

func ServiceType(ts ...string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		for _, t := range ts {
			if sc.Service.ServiceType == t {
				return true
			}
		}
		return false
	}
}

func DeviceUDN(udn string) ServiceClientFilter {
	return func(sc ServiceClient) bool {
		return sc.RootDevice.Device.UDN == udn
	}
}
