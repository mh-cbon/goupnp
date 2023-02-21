package goupnp

import (
	"net"
)

// localIPv2MCastAddrs returns the set of IPv4 addresses on multicast-able
// network interfaces.
func localIPv4MCastAddrs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, ctxError(err, "requesting host interfaces")
	}

	// Find the set of addresses to listen on.
	var addrs []string
	for _, iface := range ifaces {
		if iface.Flags&net.FlagMulticast == 0 || iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			// Does not support multicast or is a loopback address.
			continue
		}
		ifaceAddrs, err := iface.Addrs()
		if err != nil {
			return nil, ctxErrorf(err,
				"finding addresses on interface %s", iface.Name)
		}
		for _, netAddr := range ifaceAddrs {
			addr, ok := netAddr.(*net.IPNet)
			if !ok {
				// Not an IPNet address.
				continue
			}
			if addr.IP.To4() == nil {
				// Not IPv4.
				continue
			}
			addrs = append(addrs, addr.IP.String())
		}
	}

	return addrs, nil
}
