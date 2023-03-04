package lookup

import (
	"context"
	"net"

	"github.com/huin/goupnp"
)

type ServiceTable map[string]ServiceTableProvider

type ServiceTableProvider func(goupnp.MultiServiceClient) ServiceImplementations

type ServiceImplementation interface {
	LocalAddr() net.IP
	GetServiceClient() *goupnp.ServiceClient
}

type ServiceImplementations []ServiceImplementation

func (s ServiceTable) Targets() (out []string) {
	for urn := range s {
		out = append(out, urn)
	}
	return
}

func (s ServiceTable) Lookup(client goupnp.Discovery) (ServiceImplementations, error) {

	searchTargets := s.Targets()

	services, _, err := client.NewServiceClientsCtx(context.Background(), searchTargets...)
	if err != nil {
		return nil, err
	}

	return s.Implement(services), nil
}

func (s ServiceTable) Implement(c goupnp.MultiServiceClient) (out ServiceImplementations) {
	for st, provider := range s {
		svcs := c.Filter(goupnp.ServiceType(st))
		out = append(out, provider(svcs)...)
	}
	return out
}

func (s ServiceTable) ExpandAnyAddr(services goupnp.MultiServiceClient) (impls ServiceImplementations) {
	isAnyAddrSvcs := services.Filter(goupnp.ServiceLocalAddr("0.0.0.0"))
	isNotAnyAddrSvcs := services.Filter(goupnp.NotServiceLocalAddr("0.0.0.0"))

	for _, svc := range isAnyAddrSvcs {
		bySvcHost := goupnp.ServiceHost(svc.Location.Host)
		bySvcType := goupnp.ServiceType(svc.Service.ServiceType)
		for _, related := range isNotAnyAddrSvcs.Filter(bySvcHost, bySvcType) {
			for _, v := range s.Implement(isAnyAddrSvcs.Filter(bySvcHost, bySvcType)) {
				impls = append(impls, AnyAddrServiceImplementation{
					ServiceImplementation: v,
					Related:               related,
				})
			}
		}
	}

	impls = append(impls, s.Implement(isNotAnyAddrSvcs)...)

	return
}

type AnyAddrServiceImplementation struct {
	ServiceImplementation
	Related goupnp.ServiceClient
}

func (v AnyAddrServiceImplementation) LocalAddr() net.IP {
	return v.Related.LocalAddr()
}
