// Example program to display all devices discovered on the local network.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
	"github.com/huin/goupnp/lookup"
	"github.com/huin/goupnp/portmapping"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	client := goupnp.Discovery{
		Client: goupnp.AutoDiscover("0.0.0.0"),
	}

	svcTable := lookup.ServiceTable{
		internetgateway2.URN_WANIPConnection_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
			for _, v := range internetgateway2.NewWANIPConnection1MultiClients(services) {
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
		internetgateway2.URN_WANIPv6FirewallControl_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
			for _, v := range internetgateway2.NewWANIPv6FirewallControl1MultiClients(services) {
				out = append(out, v)
			}
			return
		},
		internetgateway2.URN_WANCableLinkConfig_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
			for _, v := range internetgateway2.NewWANCableLinkConfig1MultiClients(services) {
				out = append(out, v)
			}
			return
		},
		internetgateway1.URN_WANDSLLinkConfig_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
			for _, v := range internetgateway1.NewWANDSLLinkConfig1MultiClients(services) {
				out = append(out, v)
			}
			return
		},
		internetgateway1.URN_Layer3Forwarding_1: func(services goupnp.MultiServiceClient) (out lookup.ServiceImplementations) {
			for _, v := range internetgateway1.NewLayer3Forwarding1MultiClients(services) {
				out = append(out, v)
			}
			return
		},
	}

	// impls, err := svcTable.Lookup(client)
	// or...

	searchTargets := svcTable.Targets()
	services, _, err := client.NewServiceClientsCtx(context.Background(), searchTargets...)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("found %v services\n", len(services))
	fmt.Println()

	for _, svc := range services {
		fmt.Printf("%T\n", svc)
		printService(svc)
	}

	// do some filterings...
	// services.Filter(goupnp.ServiceHost("192..."))

	// create implementations
	// impls := svcTable.Implement(services)
	impls := svcTable.ExpandAnyAddr(services)

	// It is the same as rolling a manual loop of the desired clients.

	// out = out[:0]
	// for _, v := range internetgateway2.NewWANIPConnection1MultiClients(services) {
	// 	out = append(out, v)
	// }
	// for _, v := range internetgateway2.NewWANIPConnection2MultiClients(services) {
	// 	out = append(out, v)
	// }

	fmt.Println()
	fmt.Printf("found %v implemented clients\n", len(impls))
	fmt.Println()

	for _, impl := range impls {
		switch x := impl.(type) {
		case lookup.AnyAddrServiceImplementation:
			fmt.Printf("%v should port map using %v\n", x.ServiceImplementation.LocalAddr(), impl.LocalAddr())
		}
		_, ok := impl.(portmapping.PortMapper)
		fmt.Printf("%T PortMapper=%v\n", impl, ok)
		printService(*impl.GetServiceClient())
	}

	return nil
}

func printService(svc goupnp.ServiceClient) {
	fmt.Printf("LocalAddr: %v\n", svc.LocalAddr().String())
	fmt.Printf("Location: %v\n", svc.Location)
	fmt.Printf("ServiceId: %v\n", svc.Service.ServiceId)
	fmt.Printf("ServiceType: %v\n", svc.Service.ServiceType)
	fmt.Printf("ControlURL: %v\n", svc.Service.ControlURL)
	device := svc.RootDevice
	fmt.Printf("  Root v%d.%d @ %s\n",
		device.SpecVersion.Major, device.SpecVersion.Minor,
		device.URLBaseStr)
	fmt.Printf("  Type: %s\n", device.Device.DeviceType)
	fmt.Printf("  Friendly name: %s\n", device.Device.FriendlyName)
	fmt.Printf("  UDN: %s\n", device.Device.UDN)
	fmt.Printf("  UPC: %s\n", device.Device.UPC)
	fmt.Printf("  Num devices: %d\n", len(device.Device.Devices))
	fmt.Println()
}
