package main

// DCP contains extra metadata to use when generating DCP source files.
type DCPMetadata struct {
	Name         string // What to name the Go DCP package.
	OfficialName string // Official name for the DCP.
	Src          dcpProvider
}

var dcpMetadata = []DCPMetadata{
	{
		Name:         "internetgateway1",
		OfficialName: "Internet Gateway Device v1",
		Src: multiProvider{
			// upnpdotorg{
			// 	DocURL:     "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v1-Device.pdf",
			// 	XMLSpecURL: "http://upnp.org/specs/gw/UPnP-gw-IGD-TestFiles-20010921.zip",
			// 	Hacks:      []DCPHackFn{totalBytesHack},
			// },
			openconnectivitydotorg{
				SpecsURL:       allSpecsURL,
				DocPath:        "*/*/UPnP-gw-*v1*.pdf",
				XMLSpecZipPath: "*/*/UPnP-gw-IGD-TestFiles-*.zip",
				XMLServicePath: []string{"*/service/*1.xml", "*/service/WANIPConnection2.xml"},
				XMLDevicePath:  []string{"*/device/*1.xml", "*/device/WANConnectionDevice2.xml", "*/device/WANDevice2.xml"},
				Hacks:          []DCPHackFn{fixMissingURN, totalBytesHack},
			},
		},
	},
	{
		Name:         "internetgateway2",
		OfficialName: "Internet Gateway Device v2",
		Src: multiProvider{
			// upnpdotorg{
			// 	DocURL:     "http://upnp.org/specs/gw/UPnP-gw-InternetGatewayDevice-v2-Device.pdf",
			// 	XMLSpecURL: "http://upnp.org/specs/gw/UPnP-gw-IGD-Testfiles-20110224.zip",
			// 	Hacks: []DCPHackFn{fixMissingURN, totalBytesHack},
			// },
			openconnectivitydotorg{
				SpecsURL:       allSpecsURL,
				DocPath:        "*/*/UPnP-gw-*v2*.pdf",
				XMLSpecZipPath: "*/*/UPnP-gw-IGD-TestFiles-*.zip",
				XMLServicePath: []string{"*/service/*1.xml", "*/service/*2.xml"},
				XMLDevicePath:  []string{"*/device/*1.xml", "*/device/*2.xml"},
				Hacks:          []DCPHackFn{fixMissingURN, totalBytesHack},
			},
		},
	},
	{
		Name:         "av1",
		OfficialName: "MediaServer v1 and MediaRenderer v1",
		Src: multiProvider{
			upnpdotorg{
				DocURL:     "http://upnp.org/specs/av/av1/",
				XMLSpecURL: "http://upnp.org/specs/av/UPnP-av-TestFiles-20070927.zip",
			},
		},
	},
	{
		Name:         "av3",
		OfficialName: "MediaServer v3 and MediaRenderer v4",
		Src: multiProvider{
			openconnectivitydotorg{
				SpecsURL:       allSpecsURL,
				DocPath:        "*/*/UPnP-gw-*v{3,4}*.pdf",
				XMLSpecZipPath: "*/*/UPnP-av-TestFiles-*.zip",
				XMLServicePath: []string{"*/*/service/*3.xml", "*/*/service/*4.xml"},
				XMLDevicePath:  []string{"*/*/device/*3.xml", "*/*/device/*4.xml"},
			},
		},
	},
}

func totalBytesHack(dcp *DCP) error {
	for _, service := range dcp.Services {
		if service.URN == "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1" {
			variables := service.SCPD.StateVariables
			for key, variable := range variables {
				varName := variable.Name
				if varName == "TotalBytesSent" || varName == "TotalBytesReceived" {
					// Fix size of total bytes which is by default ui4 or maximum 4 GiB.
					variable.DataType.Name = "ui8"
					variables[key] = variable
				}
			}

			break
		}
	}

	return nil
}

func fixMissingURN(dcp *DCP) error {
	missingURN := "urn:schemas-upnp-org:service:WANIPv6FirewallControl:1"
	if _, ok := dcp.ServiceTypes[missingURN]; ok {
		return nil
	}
	urnParts, err := extractURNParts(missingURN, serviceURNPrefix)
	if err != nil {
		return err
	}
	dcp.ServiceTypes[missingURN] = urnParts
	return nil
}

type DCPHackFn func(*DCP) error
