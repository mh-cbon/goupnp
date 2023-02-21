package portmapping

import (
	"context"
	"log"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
	"github.com/huin/goupnp/lookup"
)

var IGW = lookup.ServiceTable{
	internetgateway1.URN_WANIPConnection_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
		for _, v := range internetgateway1.NewWANIPConnection1MultiClients(services) {
			out = append(out, v)
		}
		return
	},
	internetgateway2.URN_WANIPConnection_2: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
		for _, v := range internetgateway2.NewWANIPConnection2MultiClients(services) {
			out = append(out, v)
		}
		return
	},
}

func Lookup(client goupnp.Discovery) ([]PortMapperClient, error) {
	var out []PortMapperClient

	searchTargets := IGW.Targets()

	services, _, err := client.NewServiceClientsCtx(context.Background(), searchTargets...)
	if err != nil {
		return out, err
	}

	impls := IGW.ExpandAnyAddr(services)

	for _, impl := range impls {
		switch x := impl.(type) {
		case lookup.AnyAddrServiceImplementation:
			out = append(out, PortMapperClient{
				PortMapper:            x.ServiceImplementation.(PortMapper),
				ServiceImplementation: impl,
			})
		case PortMapper:
			out = append(out, PortMapperClient{
				PortMapper:            x,
				ServiceImplementation: impl,
			})
		default:
			log.Printf("not a port mapper %T", x)
		}
	}

	return out, nil
}

type PortMapper interface {
	AddPortMappingCtx(
		ctx context.Context,
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
		NewInternalPort uint16,
		NewInternalClient string,
		NewEnabled bool,
		NewPortMappingDescription string,
		NewLeaseDuration uint32,
	) (err error)
	AddPortMapping(
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
		NewInternalPort uint16,
		NewInternalClient string,
		NewEnabled bool,
		NewPortMappingDescription string,
		NewLeaseDuration uint32,
	) (err error)
	DeletePortMappingCtx(
		ctx context.Context,
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
	) (err error)
	DeletePortMapping(
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
	) (err error)
	GetExternalIPAddressCtx(
		ctx context.Context,
	) (NewExternalIPAddress string, err error)
	GetExternalIPAddress() (NewExternalIPAddress string, err error)
}

type PortMapperClient struct {
	PortMapper
	lookup.ServiceImplementation
}
