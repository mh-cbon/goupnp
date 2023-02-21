// Example program to display all devices discovered on the local network.
package main

import (
	"fmt"
	"log"

	"github.com/huin/goupnp"
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

	portmappers, err := portmapping.Lookup(client)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("found %v portMapper\n", len(portmappers))
	fmt.Println()

	for _, pm := range portmappers {
		switch x := pm.ServiceImplementation.(type) {
		case lookup.AnyAddrServiceImplementation:
			fmt.Printf("%v should port map using %v\n", x.ServiceImplementation.LocalAddr(), pm.LocalAddr())
		}
		fmt.Printf("%T\n", pm)
		printService(pm.GetServiceClient())
		_ = pm.AddPortMapping
	}

	return nil
}

func printService(svc *goupnp.ServiceClient) {
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
