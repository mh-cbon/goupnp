// Client for UPnP Device Control Protocol MediaServer v3 and MediaRenderer v4.
//
//
// Typically, use one of the New* functions to create clients for services.
package av3

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
	"context"
	"net/url"
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ()

// Service URNs:
const (
	URN_AVTransport_3        = "urn:schemas-upnp-org:service:AVTransport:3"
	URN_ConnectionManager_3  = "urn:schemas-upnp-org:service:ConnectionManager:3"
	URN_ContentDirectory_4   = "urn:schemas-upnp-org:service:ContentDirectory:4"
	URN_RenderingControl_3   = "urn:schemas-upnp-org:service:RenderingControl:3"
	URN_ScheduledRecording_2 = "urn:schemas-upnp-org:service:ScheduledRecording:2"
)

// AVTransport3 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport3 struct {
	goupnp.ServiceClient
}

// NewAVTransport3Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport3Clients() (clients []*AVTransport3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients(URN_AVTransport_3); err != nil {
		return
	}
	clients = newAVTransport3ClientsFromGenericClients(genericClients)
	return
}

// NewAVTransport3ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport3ClientsByURL(loc *url.URL) ([]*AVTransport3, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_AVTransport_3)
	if err != nil {
		return nil, err
	}
	return newAVTransport3ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_3)
	if err != nil {
		return nil, err
	}
	return newAVTransport3ClientsFromGenericClients(genericClients), nil
}

func newAVTransport3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport3 {
	clients := make([]*AVTransport3, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport3{genericClients[i]}
	}
	return clients
}

func (client *AVTransport3) SetAVTransportURICtx(
	ctx context.Context,
	InstanceID uint32,
	CurrentURI string,
	CurrentURIMetaData string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID         string
		CurrentURI         string
		CurrentURIMetaData string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.CurrentURI, err = soap.MarshalString(CurrentURI); err != nil {
		return
	}
	if request.CurrentURIMetaData, err = soap.MarshalString(CurrentURIMetaData); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAVTransportURI is the legacy version of SetAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetAVTransportURI(InstanceID uint32, CurrentURI string, CurrentURIMetaData string) (err error) {
	return client.SetAVTransportURICtx(context.Background(),
		InstanceID,
		CurrentURI,
		CurrentURIMetaData,
	)
}

func (client *AVTransport3) SetNextAVTransportURICtx(
	ctx context.Context,
	InstanceID uint32,
	NextURI string,
	NextURIMetaData string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID      string
		NextURI         string
		NextURIMetaData string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.NextURI, err = soap.MarshalString(NextURI); err != nil {
		return
	}
	if request.NextURIMetaData, err = soap.MarshalString(NextURIMetaData); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetNextAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetNextAVTransportURI is the legacy version of SetNextAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetNextAVTransportURI(InstanceID uint32, NextURI string, NextURIMetaData string) (err error) {
	return client.SetNextAVTransportURICtx(context.Background(),
		InstanceID,
		NextURI,
		NextURIMetaData,
	)
}

//
// Return values:
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport3) GetMediaInfoCtx(
	ctx context.Context,
	InstanceID uint32,
) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NrTracks           string
		MediaDuration      string
		CurrentURI         string
		CurrentURIMetaData string
		NextURI            string
		NextURIMetaData    string
		PlayMedium         string
		RecordMedium       string
		WriteStatus        string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetMediaInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NrTracks, err = soap.UnmarshalUi4(response.NrTracks); err != nil {
		return
	}
	if MediaDuration, err = soap.UnmarshalString(response.MediaDuration); err != nil {
		return
	}
	if CurrentURI, err = soap.UnmarshalString(response.CurrentURI); err != nil {
		return
	}
	if CurrentURIMetaData, err = soap.UnmarshalString(response.CurrentURIMetaData); err != nil {
		return
	}
	if NextURI, err = soap.UnmarshalString(response.NextURI); err != nil {
		return
	}
	if NextURIMetaData, err = soap.UnmarshalString(response.NextURIMetaData); err != nil {
		return
	}
	if PlayMedium, err = soap.UnmarshalString(response.PlayMedium); err != nil {
		return
	}
	if RecordMedium, err = soap.UnmarshalString(response.RecordMedium); err != nil {
		return
	}
	if WriteStatus, err = soap.UnmarshalString(response.WriteStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetMediaInfo is the legacy version of GetMediaInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetMediaInfo(InstanceID uint32) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	return client.GetMediaInfoCtx(context.Background(),
		InstanceID,
	)
}

//
// Return values:
//
// * CurrentType: allowed values: NO_MEDIA, TRACK_AWARE, TRACK_UNAWARE
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport3) GetMediaInfo_ExtCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentType string, NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentType        string
		NrTracks           string
		MediaDuration      string
		CurrentURI         string
		CurrentURIMetaData string
		NextURI            string
		NextURIMetaData    string
		PlayMedium         string
		RecordMedium       string
		WriteStatus        string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetMediaInfo_Ext", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentType, err = soap.UnmarshalString(response.CurrentType); err != nil {
		return
	}
	if NrTracks, err = soap.UnmarshalUi4(response.NrTracks); err != nil {
		return
	}
	if MediaDuration, err = soap.UnmarshalString(response.MediaDuration); err != nil {
		return
	}
	if CurrentURI, err = soap.UnmarshalString(response.CurrentURI); err != nil {
		return
	}
	if CurrentURIMetaData, err = soap.UnmarshalString(response.CurrentURIMetaData); err != nil {
		return
	}
	if NextURI, err = soap.UnmarshalString(response.NextURI); err != nil {
		return
	}
	if NextURIMetaData, err = soap.UnmarshalString(response.NextURIMetaData); err != nil {
		return
	}
	if PlayMedium, err = soap.UnmarshalString(response.PlayMedium); err != nil {
		return
	}
	if RecordMedium, err = soap.UnmarshalString(response.RecordMedium); err != nil {
		return
	}
	if WriteStatus, err = soap.UnmarshalString(response.WriteStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetMediaInfo_Ext is the legacy version of GetMediaInfo_ExtCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetMediaInfo_Ext(InstanceID uint32) (CurrentType string, NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	return client.GetMediaInfo_ExtCtx(context.Background(),
		InstanceID,
	)
}

//
// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PLAYING
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport3) GetTransportInfoCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentTransportState  string
		CurrentTransportStatus string
		CurrentSpeed           string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetTransportInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentTransportState, err = soap.UnmarshalString(response.CurrentTransportState); err != nil {
		return
	}
	if CurrentTransportStatus, err = soap.UnmarshalString(response.CurrentTransportStatus); err != nil {
		return
	}
	if CurrentSpeed, err = soap.UnmarshalString(response.CurrentSpeed); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransportInfo is the legacy version of GetTransportInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetTransportInfo(InstanceID uint32) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
	return client.GetTransportInfoCtx(context.Background(),
		InstanceID,
	)
}

//
// Return values:
//
// * Track: allowed value range: minimum=0, step=1
func (client *AVTransport3) GetPositionInfoCtx(
	ctx context.Context,
	InstanceID uint32,
) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount uint32, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Track         string
		TrackDuration string
		TrackMetaData string
		TrackURI      string
		RelTime       string
		AbsTime       string
		RelCount      string
		AbsCount      string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetPositionInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Track, err = soap.UnmarshalUi4(response.Track); err != nil {
		return
	}
	if TrackDuration, err = soap.UnmarshalString(response.TrackDuration); err != nil {
		return
	}
	if TrackMetaData, err = soap.UnmarshalString(response.TrackMetaData); err != nil {
		return
	}
	if TrackURI, err = soap.UnmarshalString(response.TrackURI); err != nil {
		return
	}
	if RelTime, err = soap.UnmarshalString(response.RelTime); err != nil {
		return
	}
	if AbsTime, err = soap.UnmarshalString(response.AbsTime); err != nil {
		return
	}
	if RelCount, err = soap.UnmarshalI4(response.RelCount); err != nil {
		return
	}
	if AbsCount, err = soap.UnmarshalUi4(response.AbsCount); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPositionInfo is the legacy version of GetPositionInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetPositionInfo(InstanceID uint32) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount uint32, err error) {
	return client.GetPositionInfoCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) GetDeviceCapabilitiesCtx(
	ctx context.Context,
	InstanceID uint32,
) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PlayMedia       string
		RecMedia        string
		RecQualityModes string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetDeviceCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PlayMedia, err = soap.UnmarshalString(response.PlayMedia); err != nil {
		return
	}
	if RecMedia, err = soap.UnmarshalString(response.RecMedia); err != nil {
		return
	}
	if RecQualityModes, err = soap.UnmarshalString(response.RecQualityModes); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDeviceCapabilities is the legacy version of GetDeviceCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetDeviceCapabilities(InstanceID uint32) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
	return client.GetDeviceCapabilitiesCtx(context.Background(),
		InstanceID,
	)
}

//
// Return values:
//
// * PlayMode: allowed values: NORMAL
func (client *AVTransport3) GetTransportSettingsCtx(
	ctx context.Context,
	InstanceID uint32,
) (PlayMode string, RecQualityMode string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PlayMode       string
		RecQualityMode string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetTransportSettings", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PlayMode, err = soap.UnmarshalString(response.PlayMode); err != nil {
		return
	}
	if RecQualityMode, err = soap.UnmarshalString(response.RecQualityMode); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransportSettings is the legacy version of GetTransportSettingsCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetTransportSettings(InstanceID uint32) (PlayMode string, RecQualityMode string, err error) {
	return client.GetTransportSettingsCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) StopCtx(
	ctx context.Context,
	InstanceID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Stop(InstanceID uint32) (err error) {
	return client.StopCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Speed: allowed values: 1

func (client *AVTransport3) PlayCtx(
	ctx context.Context,
	InstanceID uint32,
	Speed string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Speed      string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Speed, err = soap.MarshalString(Speed); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Play(InstanceID uint32, Speed string) (err error) {
	return client.PlayCtx(context.Background(),
		InstanceID,
		Speed,
	)
}

func (client *AVTransport3) PauseCtx(
	ctx context.Context,
	InstanceID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Pause(InstanceID uint32) (err error) {
	return client.PauseCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) RecordCtx(
	ctx context.Context,
	InstanceID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Record", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Record is the legacy version of RecordCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Record(InstanceID uint32) (err error) {
	return client.RecordCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Unit: allowed values: TRACK_NR

func (client *AVTransport3) SeekCtx(
	ctx context.Context,
	InstanceID uint32,
	Unit string,
	Target string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Unit       string
		Target     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Unit, err = soap.MarshalString(Unit); err != nil {
		return
	}
	if request.Target, err = soap.MarshalString(Target); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Seek", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Seek is the legacy version of SeekCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Seek(InstanceID uint32, Unit string, Target string) (err error) {
	return client.SeekCtx(context.Background(),
		InstanceID,
		Unit,
		Target,
	)
}

func (client *AVTransport3) NextCtx(
	ctx context.Context,
	InstanceID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Next", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Next is the legacy version of NextCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Next(InstanceID uint32) (err error) {
	return client.NextCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) PreviousCtx(
	ctx context.Context,
	InstanceID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "Previous", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Previous is the legacy version of PreviousCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) Previous(InstanceID uint32) (err error) {
	return client.PreviousCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * NewPlayMode: allowed values: NORMAL

func (client *AVTransport3) SetPlayModeCtx(
	ctx context.Context,
	InstanceID uint32,
	NewPlayMode string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID  string
		NewPlayMode string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.NewPlayMode, err = soap.MarshalString(NewPlayMode); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetPlayMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetPlayMode is the legacy version of SetPlayModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetPlayMode(InstanceID uint32, NewPlayMode string) (err error) {
	return client.SetPlayModeCtx(context.Background(),
		InstanceID,
		NewPlayMode,
	)
}

func (client *AVTransport3) SetRecordQualityModeCtx(
	ctx context.Context,
	InstanceID uint32,
	NewRecordQualityMode string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID           string
		NewRecordQualityMode string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.NewRecordQualityMode, err = soap.MarshalString(NewRecordQualityMode); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetRecordQualityMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRecordQualityMode is the legacy version of SetRecordQualityModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetRecordQualityMode(InstanceID uint32, NewRecordQualityMode string) (err error) {
	return client.SetRecordQualityModeCtx(context.Background(),
		InstanceID,
		NewRecordQualityMode,
	)
}

func (client *AVTransport3) GetCurrentTransportActionsCtx(
	ctx context.Context,
	InstanceID uint32,
) (Actions string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Actions string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetCurrentTransportActions", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Actions, err = soap.UnmarshalString(response.Actions); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetCurrentTransportActions is the legacy version of GetCurrentTransportActionsCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetCurrentTransportActions(InstanceID uint32) (Actions string, err error) {
	return client.GetCurrentTransportActionsCtx(context.Background(),
		InstanceID,
	)
}

//
// Return values:
//
// * CurrentDRMState: allowed values: OK
func (client *AVTransport3) GetDRMStateCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentDRMState string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentDRMState string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetDRMState", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentDRMState, err = soap.UnmarshalString(response.CurrentDRMState); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDRMState is the legacy version of GetDRMStateCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetDRMState(InstanceID uint32) (CurrentDRMState string, err error) {
	return client.GetDRMStateCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) GetStateVariablesCtx(
	ctx context.Context,
	InstanceID uint32,
	StateVariableList string,
) (StateVariableValuePairs string, err error) {
	// Request structure.
	request := &struct {
		InstanceID        string
		StateVariableList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.StateVariableList, err = soap.MarshalString(StateVariableList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StateVariableValuePairs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetStateVariables", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StateVariableValuePairs, err = soap.UnmarshalString(response.StateVariableValuePairs); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetStateVariables is the legacy version of GetStateVariablesCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetStateVariables(InstanceID uint32, StateVariableList string) (StateVariableValuePairs string, err error) {
	return client.GetStateVariablesCtx(context.Background(),
		InstanceID,
		StateVariableList,
	)
}

func (client *AVTransport3) SetStateVariablesCtx(
	ctx context.Context,
	InstanceID uint32,
	AVTransportUDN string,
	ServiceType string,
	ServiceId string,
	StateVariableValuePairs string,
) (StateVariableList string, err error) {
	// Request structure.
	request := &struct {
		InstanceID              string
		AVTransportUDN          string
		ServiceType             string
		ServiceId               string
		StateVariableValuePairs string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.AVTransportUDN, err = soap.MarshalString(AVTransportUDN); err != nil {
		return
	}
	if request.ServiceType, err = soap.MarshalString(ServiceType); err != nil {
		return
	}
	if request.ServiceId, err = soap.MarshalString(ServiceId); err != nil {
		return
	}
	if request.StateVariableValuePairs, err = soap.MarshalString(StateVariableValuePairs); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StateVariableList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetStateVariables", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StateVariableList, err = soap.UnmarshalString(response.StateVariableList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SetStateVariables is the legacy version of SetStateVariablesCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetStateVariables(InstanceID uint32, AVTransportUDN string, ServiceType string, ServiceId string, StateVariableValuePairs string) (StateVariableList string, err error) {
	return client.SetStateVariablesCtx(context.Background(),
		InstanceID,
		AVTransportUDN,
		ServiceType,
		ServiceId,
		StateVariableValuePairs,
	)
}

func (client *AVTransport3) AdjustSyncOffsetCtx(
	ctx context.Context,
	InstanceID uint32,
	Adjustment string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Adjustment string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Adjustment, err = soap.MarshalString(Adjustment); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "AdjustSyncOffset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// AdjustSyncOffset is the legacy version of AdjustSyncOffsetCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) AdjustSyncOffset(InstanceID uint32, Adjustment string) (err error) {
	return client.AdjustSyncOffsetCtx(context.Background(),
		InstanceID,
		Adjustment,
	)
}

func (client *AVTransport3) GetSyncOffsetCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentSyncOffset string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentSyncOffset string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetSyncOffset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentSyncOffset, err = soap.UnmarshalString(response.CurrentSyncOffset); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSyncOffset is the legacy version of GetSyncOffsetCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetSyncOffset(InstanceID uint32) (CurrentSyncOffset string, err error) {
	return client.GetSyncOffsetCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport3) SetSyncOffsetCtx(
	ctx context.Context,
	InstanceID uint32,
	NewSyncOffset string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID    string
		NewSyncOffset string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.NewSyncOffset, err = soap.MarshalString(NewSyncOffset); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetSyncOffset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSyncOffset is the legacy version of SetSyncOffsetCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetSyncOffset(InstanceID uint32, NewSyncOffset string) (err error) {
	return client.SetSyncOffsetCtx(context.Background(),
		InstanceID,
		NewSyncOffset,
	)
}

//
// Arguments:
//
// * Speed: allowed values: 1
//
// * ReferencePositionUnits: allowed values: TRACK_NR

func (client *AVTransport3) SyncPlayCtx(
	ctx context.Context,
	InstanceID uint32,
	Speed string,
	ReferencePositionUnits string,
	ReferencePosition string,
	ReferencePresentationTime string,
	ReferenceClockId string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID                string
		Speed                     string
		ReferencePositionUnits    string
		ReferencePosition         string
		ReferencePresentationTime string
		ReferenceClockId          string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Speed, err = soap.MarshalString(Speed); err != nil {
		return
	}
	if request.ReferencePositionUnits, err = soap.MarshalString(ReferencePositionUnits); err != nil {
		return
	}
	if request.ReferencePosition, err = soap.MarshalString(ReferencePosition); err != nil {
		return
	}
	if request.ReferencePresentationTime, err = soap.MarshalString(ReferencePresentationTime); err != nil {
		return
	}
	if request.ReferenceClockId, err = soap.MarshalString(ReferenceClockId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SyncPlay", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SyncPlay is the legacy version of SyncPlayCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SyncPlay(InstanceID uint32, Speed string, ReferencePositionUnits string, ReferencePosition string, ReferencePresentationTime string, ReferenceClockId string) (err error) {
	return client.SyncPlayCtx(context.Background(),
		InstanceID,
		Speed,
		ReferencePositionUnits,
		ReferencePosition,
		ReferencePresentationTime,
		ReferenceClockId,
	)
}

func (client *AVTransport3) SyncPauseCtx(
	ctx context.Context,
	InstanceID uint32,
	PauseTime string,
	ReferenceClockId string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID       string
		PauseTime        string
		ReferenceClockId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.PauseTime, err = soap.MarshalString(PauseTime); err != nil {
		return
	}
	if request.ReferenceClockId, err = soap.MarshalString(ReferenceClockId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SyncPause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SyncPause is the legacy version of SyncPauseCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SyncPause(InstanceID uint32, PauseTime string, ReferenceClockId string) (err error) {
	return client.SyncPauseCtx(context.Background(),
		InstanceID,
		PauseTime,
		ReferenceClockId,
	)
}

func (client *AVTransport3) SyncStopCtx(
	ctx context.Context,
	InstanceID uint32,
	StopTime string,
	ReferenceClockId string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID       string
		StopTime         string
		ReferenceClockId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.StopTime, err = soap.MarshalString(StopTime); err != nil {
		return
	}
	if request.ReferenceClockId, err = soap.MarshalString(ReferenceClockId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SyncStop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SyncStop is the legacy version of SyncStopCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SyncStop(InstanceID uint32, StopTime string, ReferenceClockId string) (err error) {
	return client.SyncStopCtx(context.Background(),
		InstanceID,
		StopTime,
		ReferenceClockId,
	)
}

func (client *AVTransport3) SetStaticPlaylistCtx(
	ctx context.Context,
	InstanceID uint32,
	PlaylistData string,
	PlaylistDataLength uint32,
	PlaylistOffset uint32,
	PlaylistTotalLength uint32,
	PlaylistMIMEType string,
	PlaylistExtendedType string,
	PlaylistStartObj string,
	PlaylistStartGroup string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID           string
		PlaylistData         string
		PlaylistDataLength   string
		PlaylistOffset       string
		PlaylistTotalLength  string
		PlaylistMIMEType     string
		PlaylistExtendedType string
		PlaylistStartObj     string
		PlaylistStartGroup   string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.PlaylistData, err = soap.MarshalString(PlaylistData); err != nil {
		return
	}
	if request.PlaylistDataLength, err = soap.MarshalUi4(PlaylistDataLength); err != nil {
		return
	}
	if request.PlaylistOffset, err = soap.MarshalUi4(PlaylistOffset); err != nil {
		return
	}
	if request.PlaylistTotalLength, err = soap.MarshalUi4(PlaylistTotalLength); err != nil {
		return
	}
	if request.PlaylistMIMEType, err = soap.MarshalString(PlaylistMIMEType); err != nil {
		return
	}
	if request.PlaylistExtendedType, err = soap.MarshalString(PlaylistExtendedType); err != nil {
		return
	}
	if request.PlaylistStartObj, err = soap.MarshalString(PlaylistStartObj); err != nil {
		return
	}
	if request.PlaylistStartGroup, err = soap.MarshalString(PlaylistStartGroup); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetStaticPlaylist", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetStaticPlaylist is the legacy version of SetStaticPlaylistCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetStaticPlaylist(InstanceID uint32, PlaylistData string, PlaylistDataLength uint32, PlaylistOffset uint32, PlaylistTotalLength uint32, PlaylistMIMEType string, PlaylistExtendedType string, PlaylistStartObj string, PlaylistStartGroup string) (err error) {
	return client.SetStaticPlaylistCtx(context.Background(),
		InstanceID,
		PlaylistData,
		PlaylistDataLength,
		PlaylistOffset,
		PlaylistTotalLength,
		PlaylistMIMEType,
		PlaylistExtendedType,
		PlaylistStartObj,
		PlaylistStartGroup,
	)
}

//
// Arguments:
//
// * PlaylistStep: allowed values: Initial, Continue, Stop, Reset

func (client *AVTransport3) SetStreamingPlaylistCtx(
	ctx context.Context,
	InstanceID uint32,
	PlaylistData string,
	PlaylistDataLength uint32,
	PlaylistMIMEType string,
	PlaylistExtendedType string,
	PlaylistStep string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID           string
		PlaylistData         string
		PlaylistDataLength   string
		PlaylistMIMEType     string
		PlaylistExtendedType string
		PlaylistStep         string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.PlaylistData, err = soap.MarshalString(PlaylistData); err != nil {
		return
	}
	if request.PlaylistDataLength, err = soap.MarshalUi4(PlaylistDataLength); err != nil {
		return
	}
	if request.PlaylistMIMEType, err = soap.MarshalString(PlaylistMIMEType); err != nil {
		return
	}
	if request.PlaylistExtendedType, err = soap.MarshalString(PlaylistExtendedType); err != nil {
		return
	}
	if request.PlaylistStep, err = soap.MarshalString(PlaylistStep); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "SetStreamingPlaylist", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetStreamingPlaylist is the legacy version of SetStreamingPlaylistCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) SetStreamingPlaylist(InstanceID uint32, PlaylistData string, PlaylistDataLength uint32, PlaylistMIMEType string, PlaylistExtendedType string, PlaylistStep string) (err error) {
	return client.SetStreamingPlaylistCtx(context.Background(),
		InstanceID,
		PlaylistData,
		PlaylistDataLength,
		PlaylistMIMEType,
		PlaylistExtendedType,
		PlaylistStep,
	)
}

//
// Arguments:
//
// * PlaylistType: allowed values: Static, Streaming

func (client *AVTransport3) GetPlaylistInfoCtx(
	ctx context.Context,
	InstanceID uint32,
	PlaylistType string,
) (PlaylistInfo string, err error) {
	// Request structure.
	request := &struct {
		InstanceID   string
		PlaylistType string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.PlaylistType, err = soap.MarshalString(PlaylistType); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PlaylistInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_3, "GetPlaylistInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PlaylistInfo, err = soap.UnmarshalString(response.PlaylistInfo); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPlaylistInfo is the legacy version of GetPlaylistInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport3) GetPlaylistInfo(InstanceID uint32, PlaylistType string) (PlaylistInfo string, err error) {
	return client.GetPlaylistInfoCtx(context.Background(),
		InstanceID,
		PlaylistType,
	)
}

// ConnectionManager3 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager3 struct {
	goupnp.ServiceClient
}

// NewConnectionManager3Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager3Clients() (clients []*ConnectionManager3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients(URN_ConnectionManager_3); err != nil {
		return
	}
	clients = newConnectionManager3ClientsFromGenericClients(genericClients)
	return
}

// NewConnectionManager3ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager3ClientsByURL(loc *url.URL) ([]*ConnectionManager3, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ConnectionManager_3)
	if err != nil {
		return nil, err
	}
	return newConnectionManager3ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_3)
	if err != nil {
		return nil, err
	}
	return newConnectionManager3ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager3 {
	clients := make([]*ConnectionManager3, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager3{genericClients[i]}
	}
	return clients
}

func (client *ConnectionManager3) GetProtocolInfoCtx(
	ctx context.Context,
) (Source string, Sink string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Source string
		Sink   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "GetProtocolInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Source, err = soap.UnmarshalString(response.Source); err != nil {
		return
	}
	if Sink, err = soap.UnmarshalString(response.Sink); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetProtocolInfo is the legacy version of GetProtocolInfoCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) GetProtocolInfo() (Source string, Sink string, err error) {
	return client.GetProtocolInfoCtx(context.Background())
}

//
// Arguments:
//
// * Direction: allowed values: Input, Output

func (client *ConnectionManager3) PrepareForConnectionCtx(
	ctx context.Context,
	RemoteProtocolInfo string,
	PeerConnectionManager string,
	PeerConnectionID int32,
	Direction string,
) (ConnectionID int32, AVTransportID int32, RcsID int32, err error) {
	// Request structure.
	request := &struct {
		RemoteProtocolInfo    string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RemoteProtocolInfo, err = soap.MarshalString(RemoteProtocolInfo); err != nil {
		return
	}
	if request.PeerConnectionManager, err = soap.MarshalString(PeerConnectionManager); err != nil {
		return
	}
	if request.PeerConnectionID, err = soap.MarshalI4(PeerConnectionID); err != nil {
		return
	}
	if request.Direction, err = soap.MarshalString(Direction); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ConnectionID  string
		AVTransportID string
		RcsID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "PrepareForConnection", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ConnectionID, err = soap.UnmarshalI4(response.ConnectionID); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(response.AVTransportID); err != nil {
		return
	}
	if RcsID, err = soap.UnmarshalI4(response.RcsID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PrepareForConnection is the legacy version of PrepareForConnectionCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) PrepareForConnection(RemoteProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string) (ConnectionID int32, AVTransportID int32, RcsID int32, err error) {
	return client.PrepareForConnectionCtx(context.Background(),
		RemoteProtocolInfo,
		PeerConnectionManager,
		PeerConnectionID,
		Direction,
	)
}

func (client *ConnectionManager3) ConnectionCompleteCtx(
	ctx context.Context,
	ConnectionID int32,
) (err error) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ConnectionID, err = soap.MarshalI4(ConnectionID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "ConnectionComplete", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ConnectionComplete is the legacy version of ConnectionCompleteCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) ConnectionComplete(ConnectionID int32) (err error) {
	return client.ConnectionCompleteCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager3) GetCurrentConnectionIDsCtx(
	ctx context.Context,
) (ConnectionIDs string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ConnectionIDs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "GetCurrentConnectionIDs", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ConnectionIDs, err = soap.UnmarshalString(response.ConnectionIDs); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetCurrentConnectionIDs is the legacy version of GetCurrentConnectionIDsCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) GetCurrentConnectionIDs() (ConnectionIDs string, err error) {
	return client.GetCurrentConnectionIDsCtx(context.Background())
}

//
// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager3) GetCurrentConnectionInfoCtx(
	ctx context.Context,
	ConnectionID int32,
) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
	// Request structure.
	request := &struct {
		ConnectionID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ConnectionID, err = soap.MarshalI4(ConnectionID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RcsID                 string
		AVTransportID         string
		ProtocolInfo          string
		PeerConnectionManager string
		PeerConnectionID      string
		Direction             string
		Status                string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "GetCurrentConnectionInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RcsID, err = soap.UnmarshalI4(response.RcsID); err != nil {
		return
	}
	if AVTransportID, err = soap.UnmarshalI4(response.AVTransportID); err != nil {
		return
	}
	if ProtocolInfo, err = soap.UnmarshalString(response.ProtocolInfo); err != nil {
		return
	}
	if PeerConnectionManager, err = soap.UnmarshalString(response.PeerConnectionManager); err != nil {
		return
	}
	if PeerConnectionID, err = soap.UnmarshalI4(response.PeerConnectionID); err != nil {
		return
	}
	if Direction, err = soap.UnmarshalString(response.Direction); err != nil {
		return
	}
	if Status, err = soap.UnmarshalString(response.Status); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetCurrentConnectionInfo is the legacy version of GetCurrentConnectionInfoCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) GetCurrentConnectionInfo(ConnectionID int32) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
	return client.GetCurrentConnectionInfoCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager3) GetRendererItemInfoCtx(
	ctx context.Context,
	ItemInfoFilter string,
	ItemMetadataList string,
) (ItemRenderingInfoList string, err error) {
	// Request structure.
	request := &struct {
		ItemInfoFilter   string
		ItemMetadataList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ItemInfoFilter, err = soap.MarshalString(ItemInfoFilter); err != nil {
		return
	}
	if request.ItemMetadataList, err = soap.MarshalString(ItemMetadataList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ItemRenderingInfoList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "GetRendererItemInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ItemRenderingInfoList, err = soap.UnmarshalString(response.ItemRenderingInfoList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRendererItemInfo is the legacy version of GetRendererItemInfoCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) GetRendererItemInfo(ItemInfoFilter string, ItemMetadataList string) (ItemRenderingInfoList string, err error) {
	return client.GetRendererItemInfoCtx(context.Background(),
		ItemInfoFilter,
		ItemMetadataList,
	)
}

func (client *ConnectionManager3) GetFeatureListCtx(
	ctx context.Context,
) (FeatureList string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		FeatureList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_3, "GetFeatureList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if FeatureList, err = soap.UnmarshalString(response.FeatureList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetFeatureList is the legacy version of GetFeatureListCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager3) GetFeatureList() (FeatureList string, err error) {
	return client.GetFeatureListCtx(context.Background())
}

// RenderingControl3 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl3 struct {
	goupnp.ServiceClient
}

// NewRenderingControl3Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl3Clients() (clients []*RenderingControl3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients(URN_RenderingControl_3); err != nil {
		return
	}
	clients = newRenderingControl3ClientsFromGenericClients(genericClients)
	return
}

// NewRenderingControl3ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl3ClientsByURL(loc *url.URL) ([]*RenderingControl3, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_RenderingControl_3)
	if err != nil {
		return nil, err
	}
	return newRenderingControl3ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_3)
	if err != nil {
		return nil, err
	}
	return newRenderingControl3ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl3 {
	clients := make([]*RenderingControl3, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl3{genericClients[i]}
	}
	return clients
}

func (client *RenderingControl3) ListPresetsCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentPresetNameList string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentPresetNameList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "ListPresets", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentPresetNameList, err = soap.UnmarshalString(response.CurrentPresetNameList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ListPresets is the legacy version of ListPresetsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) ListPresets(InstanceID uint32) (CurrentPresetNameList string, err error) {
	return client.ListPresetsCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * PresetName: allowed values: FactoryDefaults

func (client *RenderingControl3) SelectPresetCtx(
	ctx context.Context,
	InstanceID uint32,
	PresetName string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		PresetName string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.PresetName, err = soap.MarshalString(PresetName); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SelectPreset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SelectPreset is the legacy version of SelectPresetCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SelectPreset(InstanceID uint32, PresetName string) (err error) {
	return client.SelectPresetCtx(context.Background(),
		InstanceID,
		PresetName,
	)
}

//
// Return values:
//
// * CurrentBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetBrightnessCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentBrightness uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentBrightness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetBrightness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentBrightness, err = soap.UnmarshalUi2(response.CurrentBrightness); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetBrightness is the legacy version of GetBrightnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetBrightness(InstanceID uint32) (CurrentBrightness uint16, err error) {
	return client.GetBrightnessCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredBrightness: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetBrightnessCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredBrightness uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID        string
		DesiredBrightness string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredBrightness, err = soap.MarshalUi2(DesiredBrightness); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetBrightness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBrightness is the legacy version of SetBrightnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetBrightness(InstanceID uint32, DesiredBrightness uint16) (err error) {
	return client.SetBrightnessCtx(context.Background(),
		InstanceID,
		DesiredBrightness,
	)
}

//
// Return values:
//
// * CurrentContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetContrastCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentContrast uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentContrast string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetContrast", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentContrast, err = soap.UnmarshalUi2(response.CurrentContrast); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetContrast is the legacy version of GetContrastCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetContrast(InstanceID uint32) (CurrentContrast uint16, err error) {
	return client.GetContrastCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredContrast: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetContrastCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredContrast uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID      string
		DesiredContrast string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredContrast, err = soap.MarshalUi2(DesiredContrast); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetContrast", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetContrast is the legacy version of SetContrastCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetContrast(InstanceID uint32, DesiredContrast uint16) (err error) {
	return client.SetContrastCtx(context.Background(),
		InstanceID,
		DesiredContrast,
	)
}

//
// Return values:
//
// * CurrentSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetSharpnessCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentSharpness uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentSharpness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetSharpness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentSharpness, err = soap.UnmarshalUi2(response.CurrentSharpness); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSharpness is the legacy version of GetSharpnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetSharpness(InstanceID uint32) (CurrentSharpness uint16, err error) {
	return client.GetSharpnessCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredSharpness: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetSharpnessCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredSharpness uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID       string
		DesiredSharpness string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredSharpness, err = soap.MarshalUi2(DesiredSharpness); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetSharpness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSharpness is the legacy version of SetSharpnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetSharpness(InstanceID uint32, DesiredSharpness uint16) (err error) {
	return client.SetSharpnessCtx(context.Background(),
		InstanceID,
		DesiredSharpness,
	)
}

//
// Return values:
//
// * CurrentRedVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetRedVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentRedVideoGain uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentRedVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetRedVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentRedVideoGain, err = soap.UnmarshalUi2(response.CurrentRedVideoGain); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRedVideoGain is the legacy version of GetRedVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetRedVideoGain(InstanceID uint32) (CurrentRedVideoGain uint16, err error) {
	return client.GetRedVideoGainCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredRedVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetRedVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoGain uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID          string
		DesiredRedVideoGain string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredRedVideoGain, err = soap.MarshalUi2(DesiredRedVideoGain); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetRedVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoGain is the legacy version of SetRedVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetRedVideoGain(InstanceID uint32, DesiredRedVideoGain uint16) (err error) {
	return client.SetRedVideoGainCtx(context.Background(),
		InstanceID,
		DesiredRedVideoGain,
	)
}

//
// Return values:
//
// * CurrentGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetGreenVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentGreenVideoGain uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentGreenVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetGreenVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentGreenVideoGain, err = soap.UnmarshalUi2(response.CurrentGreenVideoGain); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetGreenVideoGain is the legacy version of GetGreenVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetGreenVideoGain(InstanceID uint32) (CurrentGreenVideoGain uint16, err error) {
	return client.GetGreenVideoGainCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetGreenVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoGain uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID            string
		DesiredGreenVideoGain string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredGreenVideoGain, err = soap.MarshalUi2(DesiredGreenVideoGain); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetGreenVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoGain is the legacy version of SetGreenVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetGreenVideoGain(InstanceID uint32, DesiredGreenVideoGain uint16) (err error) {
	return client.SetGreenVideoGainCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoGain,
	)
}

//
// Return values:
//
// * CurrentBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetBlueVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentBlueVideoGain uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentBlueVideoGain string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetBlueVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentBlueVideoGain, err = soap.UnmarshalUi2(response.CurrentBlueVideoGain); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetBlueVideoGain is the legacy version of GetBlueVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetBlueVideoGain(InstanceID uint32) (CurrentBlueVideoGain uint16, err error) {
	return client.GetBlueVideoGainCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetBlueVideoGainCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoGain uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID           string
		DesiredBlueVideoGain string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredBlueVideoGain, err = soap.MarshalUi2(DesiredBlueVideoGain); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetBlueVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoGain is the legacy version of SetBlueVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetBlueVideoGain(InstanceID uint32, DesiredBlueVideoGain uint16) (err error) {
	return client.SetBlueVideoGainCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoGain,
	)
}

//
// Return values:
//
// * CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetRedVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentRedVideoBlackLevel uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentRedVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetRedVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentRedVideoBlackLevel, err = soap.UnmarshalUi2(response.CurrentRedVideoBlackLevel); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRedVideoBlackLevel is the legacy version of GetRedVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetRedVideoBlackLevel(InstanceID uint32) (CurrentRedVideoBlackLevel uint16, err error) {
	return client.GetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetRedVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredRedVideoBlackLevel uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredRedVideoBlackLevel string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredRedVideoBlackLevel, err = soap.MarshalUi2(DesiredRedVideoBlackLevel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetRedVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoBlackLevel is the legacy version of SetRedVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetRedVideoBlackLevel(InstanceID uint32, DesiredRedVideoBlackLevel uint16) (err error) {
	return client.SetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredRedVideoBlackLevel,
	)
}

//
// Return values:
//
// * CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetGreenVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentGreenVideoBlackLevel uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentGreenVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetGreenVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentGreenVideoBlackLevel, err = soap.UnmarshalUi2(response.CurrentGreenVideoBlackLevel); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetGreenVideoBlackLevel is the legacy version of GetGreenVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetGreenVideoBlackLevel(InstanceID uint32) (CurrentGreenVideoBlackLevel uint16, err error) {
	return client.GetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetGreenVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredGreenVideoBlackLevel uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID                  string
		DesiredGreenVideoBlackLevel string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredGreenVideoBlackLevel, err = soap.MarshalUi2(DesiredGreenVideoBlackLevel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetGreenVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoBlackLevel is the legacy version of SetGreenVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetGreenVideoBlackLevel(InstanceID uint32, DesiredGreenVideoBlackLevel uint16) (err error) {
	return client.SetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoBlackLevel,
	)
}

//
// Return values:
//
// * CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetBlueVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentBlueVideoBlackLevel uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentBlueVideoBlackLevel string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetBlueVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentBlueVideoBlackLevel, err = soap.UnmarshalUi2(response.CurrentBlueVideoBlackLevel); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetBlueVideoBlackLevel is the legacy version of GetBlueVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetBlueVideoBlackLevel(InstanceID uint32) (CurrentBlueVideoBlackLevel uint16, err error) {
	return client.GetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetBlueVideoBlackLevelCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredBlueVideoBlackLevel uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID                 string
		DesiredBlueVideoBlackLevel string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredBlueVideoBlackLevel, err = soap.MarshalUi2(DesiredBlueVideoBlackLevel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetBlueVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoBlackLevel is the legacy version of SetBlueVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetBlueVideoBlackLevel(InstanceID uint32, DesiredBlueVideoBlackLevel uint16) (err error) {
	return client.SetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoBlackLevel,
	)
}

//
// Return values:
//
// * CurrentColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetColorTemperatureCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentColorTemperature uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentColorTemperature string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetColorTemperature", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentColorTemperature, err = soap.UnmarshalUi2(response.CurrentColorTemperature); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetColorTemperature is the legacy version of GetColorTemperatureCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetColorTemperature(InstanceID uint32) (CurrentColorTemperature uint16, err error) {
	return client.GetColorTemperatureCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredColorTemperature: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetColorTemperatureCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredColorTemperature uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredColorTemperature string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredColorTemperature, err = soap.MarshalUi2(DesiredColorTemperature); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetColorTemperature", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetColorTemperature is the legacy version of SetColorTemperatureCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetColorTemperature(InstanceID uint32, DesiredColorTemperature uint16) (err error) {
	return client.SetColorTemperatureCtx(context.Background(),
		InstanceID,
		DesiredColorTemperature,
	)
}

//
// Return values:
//
// * CurrentHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl3) GetHorizontalKeystoneCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentHorizontalKeystone int16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentHorizontalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetHorizontalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentHorizontalKeystone, err = soap.UnmarshalI2(response.CurrentHorizontalKeystone); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetHorizontalKeystone is the legacy version of GetHorizontalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetHorizontalKeystone(InstanceID uint32) (CurrentHorizontalKeystone int16, err error) {
	return client.GetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredHorizontalKeystone: allowed value range: step=1

func (client *RenderingControl3) SetHorizontalKeystoneCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredHorizontalKeystone int16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID                string
		DesiredHorizontalKeystone string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredHorizontalKeystone, err = soap.MarshalI2(DesiredHorizontalKeystone); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetHorizontalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetHorizontalKeystone is the legacy version of SetHorizontalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetHorizontalKeystone(InstanceID uint32, DesiredHorizontalKeystone int16) (err error) {
	return client.SetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredHorizontalKeystone,
	)
}

//
// Return values:
//
// * CurrentVerticalKeystone: allowed value range: step=1
func (client *RenderingControl3) GetVerticalKeystoneCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentVerticalKeystone int16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentVerticalKeystone string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetVerticalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentVerticalKeystone, err = soap.UnmarshalI2(response.CurrentVerticalKeystone); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetVerticalKeystone is the legacy version of GetVerticalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetVerticalKeystone(InstanceID uint32) (CurrentVerticalKeystone int16, err error) {
	return client.GetVerticalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * DesiredVerticalKeystone: allowed value range: step=1

func (client *RenderingControl3) SetVerticalKeystoneCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredVerticalKeystone int16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID              string
		DesiredVerticalKeystone string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredVerticalKeystone, err = soap.MarshalI2(DesiredVerticalKeystone); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetVerticalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVerticalKeystone is the legacy version of SetVerticalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetVerticalKeystone(InstanceID uint32, DesiredVerticalKeystone int16) (err error) {
	return client.SetVerticalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredVerticalKeystone,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) GetMuteCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (CurrentMute bool, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentMute string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentMute, err = soap.UnmarshalBoolean(response.CurrentMute); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetMute is the legacy version of GetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetMute(InstanceID uint32, Channel string) (CurrentMute bool, err error) {
	return client.GetMuteCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) SetMuteCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredMute bool,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID  string
		Channel     string
		DesiredMute string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.DesiredMute, err = soap.MarshalBoolean(DesiredMute); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetMute(InstanceID uint32, Channel string, DesiredMute bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredMute,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

//
// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl3) GetVolumeCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (CurrentVolume uint16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentVolume, err = soap.UnmarshalUi2(response.CurrentVolume); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetVolume is the legacy version of GetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetVolume(InstanceID uint32, Channel string) (CurrentVolume uint16, err error) {
	return client.GetVolumeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master
//
// * DesiredVolume: allowed value range: minimum=0, step=1

func (client *RenderingControl3) SetVolumeCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume uint16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalUi2(DesiredVolume); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetVolume(InstanceID uint32, Channel string, DesiredVolume uint16) (err error) {
	return client.SetVolumeCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredVolume,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) GetVolumeDBCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (CurrentVolume int16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentVolume string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetVolumeDB", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentVolume, err = soap.UnmarshalI2(response.CurrentVolume); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetVolumeDB is the legacy version of GetVolumeDBCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetVolumeDB(InstanceID uint32, Channel string) (CurrentVolume int16, err error) {
	return client.GetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) SetVolumeDBCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredVolume int16,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID    string
		Channel       string
		DesiredVolume string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.DesiredVolume, err = soap.MarshalI2(DesiredVolume); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetVolumeDB", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolumeDB is the legacy version of SetVolumeDBCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetVolumeDB(InstanceID uint32, Channel string, DesiredVolume int16) (err error) {
	return client.SetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredVolume,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) GetVolumeDBRangeCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (MinValue int16, MaxValue int16, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		MinValue string
		MaxValue string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetVolumeDBRange", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if MinValue, err = soap.UnmarshalI2(response.MinValue); err != nil {
		return
	}
	if MaxValue, err = soap.UnmarshalI2(response.MaxValue); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetVolumeDBRange is the legacy version of GetVolumeDBRangeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetVolumeDBRange(InstanceID uint32, Channel string) (MinValue int16, MaxValue int16, err error) {
	return client.GetVolumeDBRangeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) GetLoudnessCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
) (CurrentLoudness bool, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
		Channel    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentLoudness string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetLoudness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentLoudness, err = soap.UnmarshalBoolean(response.CurrentLoudness); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetLoudness is the legacy version of GetLoudnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetLoudness(InstanceID uint32, Channel string) (CurrentLoudness bool, err error) {
	return client.GetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl3) SetLoudnessCtx(
	ctx context.Context,
	InstanceID uint32,
	Channel string,
	DesiredLoudness bool,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID      string
		Channel         string
		DesiredLoudness string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.DesiredLoudness, err = soap.MarshalBoolean(DesiredLoudness); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetLoudness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetLoudness is the legacy version of SetLoudnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetLoudness(InstanceID uint32, Channel string, DesiredLoudness bool) (err error) {
	return client.SetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredLoudness,
	)
}

func (client *RenderingControl3) GetStateVariablesCtx(
	ctx context.Context,
	InstanceID uint32,
	StateVariableList string,
) (StateVariableValuePairs string, err error) {
	// Request structure.
	request := &struct {
		InstanceID        string
		StateVariableList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.StateVariableList, err = soap.MarshalString(StateVariableList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StateVariableValuePairs string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetStateVariables", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StateVariableValuePairs, err = soap.UnmarshalString(response.StateVariableValuePairs); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetStateVariables is the legacy version of GetStateVariablesCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetStateVariables(InstanceID uint32, StateVariableList string) (StateVariableValuePairs string, err error) {
	return client.GetStateVariablesCtx(context.Background(),
		InstanceID,
		StateVariableList,
	)
}

func (client *RenderingControl3) SetStateVariablesCtx(
	ctx context.Context,
	InstanceID uint32,
	RenderingControlUDN string,
	ServiceType string,
	ServiceId string,
	StateVariableValuePairs string,
) (StateVariableList string, err error) {
	// Request structure.
	request := &struct {
		InstanceID              string
		RenderingControlUDN     string
		ServiceType             string
		ServiceId               string
		StateVariableValuePairs string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.RenderingControlUDN, err = soap.MarshalString(RenderingControlUDN); err != nil {
		return
	}
	if request.ServiceType, err = soap.MarshalString(ServiceType); err != nil {
		return
	}
	if request.ServiceId, err = soap.MarshalString(ServiceId); err != nil {
		return
	}
	if request.StateVariableValuePairs, err = soap.MarshalString(StateVariableValuePairs); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StateVariableList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetStateVariables", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StateVariableList, err = soap.UnmarshalString(response.StateVariableList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SetStateVariables is the legacy version of SetStateVariablesCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetStateVariables(InstanceID uint32, RenderingControlUDN string, ServiceType string, ServiceId string, StateVariableValuePairs string) (StateVariableList string, err error) {
	return client.SetStateVariablesCtx(context.Background(),
		InstanceID,
		RenderingControlUDN,
		ServiceType,
		ServiceId,
		StateVariableValuePairs,
	)
}

func (client *RenderingControl3) GetAllowedTransformsCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentAllowedTransformSettings string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentAllowedTransformSettings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetAllowedTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentAllowedTransformSettings, err = soap.UnmarshalString(response.CurrentAllowedTransformSettings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllowedTransforms is the legacy version of GetAllowedTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetAllowedTransforms(InstanceID uint32) (CurrentAllowedTransformSettings string, err error) {
	return client.GetAllowedTransformsCtx(context.Background(),
		InstanceID,
	)
}

func (client *RenderingControl3) SetTransformsCtx(
	ctx context.Context,
	InstanceID uint32,
	DesiredTransformValues string,
) (err error) {
	// Request structure.
	request := &struct {
		InstanceID             string
		DesiredTransformValues string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	if request.DesiredTransformValues, err = soap.MarshalString(DesiredTransformValues); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetTransforms is the legacy version of SetTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetTransforms(InstanceID uint32, DesiredTransformValues string) (err error) {
	return client.SetTransformsCtx(context.Background(),
		InstanceID,
		DesiredTransformValues,
	)
}

func (client *RenderingControl3) GetTransformsCtx(
	ctx context.Context,
	InstanceID uint32,
) (CurrentTransformValues string, err error) {
	// Request structure.
	request := &struct {
		InstanceID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentTransformValues string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentTransformValues, err = soap.UnmarshalString(response.CurrentTransformValues); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransforms is the legacy version of GetTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetTransforms(InstanceID uint32) (CurrentTransformValues string, err error) {
	return client.GetTransformsCtx(context.Background(),
		InstanceID,
	)
}

func (client *RenderingControl3) GetAllAvailableTransformsCtx(
	ctx context.Context,
) (AllAllowedTransformSettings string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		AllAllowedTransformSettings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetAllAvailableTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if AllAllowedTransformSettings, err = soap.UnmarshalString(response.AllAllowedTransformSettings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllAvailableTransforms is the legacy version of GetAllAvailableTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetAllAvailableTransforms() (AllAllowedTransformSettings string, err error) {
	return client.GetAllAvailableTransformsCtx(context.Background())
}

func (client *RenderingControl3) GetAllowedDefaultTransformsCtx(
	ctx context.Context,
) (AllowedDefaultTransformSettings string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		AllowedDefaultTransformSettings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetAllowedDefaultTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if AllowedDefaultTransformSettings, err = soap.UnmarshalString(response.AllowedDefaultTransformSettings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllowedDefaultTransforms is the legacy version of GetAllowedDefaultTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetAllowedDefaultTransforms() (AllowedDefaultTransformSettings string, err error) {
	return client.GetAllowedDefaultTransformsCtx(context.Background())
}

func (client *RenderingControl3) GetDefaultTransformsCtx(
	ctx context.Context,
) (CurrentDefaultTransformSettings string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentDefaultTransformSettings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "GetDefaultTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentDefaultTransformSettings, err = soap.UnmarshalString(response.CurrentDefaultTransformSettings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDefaultTransforms is the legacy version of GetDefaultTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) GetDefaultTransforms() (CurrentDefaultTransformSettings string, err error) {
	return client.GetDefaultTransformsCtx(context.Background())
}

func (client *RenderingControl3) SetDefaultTransformsCtx(
	ctx context.Context,
	DesiredDefaultTransformSettings string,
) (err error) {
	// Request structure.
	request := &struct {
		DesiredDefaultTransformSettings string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DesiredDefaultTransformSettings, err = soap.MarshalString(DesiredDefaultTransformSettings); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_3, "SetDefaultTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetDefaultTransforms is the legacy version of SetDefaultTransformsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl3) SetDefaultTransforms(DesiredDefaultTransformSettings string) (err error) {
	return client.SetDefaultTransformsCtx(context.Background(),
		DesiredDefaultTransformSettings,
	)
}

// ContentDirectory4 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:4". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory4 struct {
	goupnp.ServiceClient
}

// NewContentDirectory4Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory4Clients() (clients []*ContentDirectory4, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients(URN_ContentDirectory_4); err != nil {
		return
	}
	clients = newContentDirectory4ClientsFromGenericClients(genericClients)
	return
}

// NewContentDirectory4ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory4ClientsByURL(loc *url.URL) ([]*ContentDirectory4, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, URN_ContentDirectory_4)
	if err != nil {
		return nil, err
	}
	return newContentDirectory4ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory4ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory4ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory4, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_4)
	if err != nil {
		return nil, err
	}
	return newContentDirectory4ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory4ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory4 {
	clients := make([]*ContentDirectory4, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory4{genericClients[i]}
	}
	return clients
}

func (client *ContentDirectory4) GetSearchCapabilitiesCtx(
	ctx context.Context,
) (SearchCaps string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SearchCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetSearchCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SearchCaps, err = soap.UnmarshalString(response.SearchCaps); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSearchCapabilities is the legacy version of GetSearchCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetSearchCapabilities() (SearchCaps string, err error) {
	return client.GetSearchCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory4) GetSortCapabilitiesCtx(
	ctx context.Context,
) (SortCaps string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SortCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetSortCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SortCaps, err = soap.UnmarshalString(response.SortCaps); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSortCapabilities is the legacy version of GetSortCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetSortCapabilities() (SortCaps string, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory4) GetSortExtensionCapabilitiesCtx(
	ctx context.Context,
) (SortExtensionCaps string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SortExtensionCaps string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetSortExtensionCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SortExtensionCaps, err = soap.UnmarshalString(response.SortExtensionCaps); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSortExtensionCapabilities is the legacy version of GetSortExtensionCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetSortExtensionCapabilities() (SortExtensionCaps string, err error) {
	return client.GetSortExtensionCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory4) GetFeatureListCtx(
	ctx context.Context,
) (FeatureList string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		FeatureList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetFeatureList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if FeatureList, err = soap.UnmarshalString(response.FeatureList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetFeatureList is the legacy version of GetFeatureListCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetFeatureList() (FeatureList string, err error) {
	return client.GetFeatureListCtx(context.Background())
}

func (client *ContentDirectory4) GetSystemUpdateIDCtx(
	ctx context.Context,
) (Id uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Id string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetSystemUpdateID", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Id, err = soap.UnmarshalUi4(response.Id); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSystemUpdateID is the legacy version of GetSystemUpdateIDCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetSystemUpdateID() (Id uint32, err error) {
	return client.GetSystemUpdateIDCtx(context.Background())
}

func (client *ContentDirectory4) GetServiceResetTokenCtx(
	ctx context.Context,
) (ResetToken string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ResetToken string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetServiceResetToken", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ResetToken, err = soap.UnmarshalString(response.ResetToken); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetServiceResetToken is the legacy version of GetServiceResetTokenCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetServiceResetToken() (ResetToken string, err error) {
	return client.GetServiceResetTokenCtx(context.Background())
}

//
// Arguments:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren

func (client *ContentDirectory4) BrowseCtx(
	ctx context.Context,
	ObjectID string,
	BrowseFlag string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		ObjectID       string
		BrowseFlag     string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ObjectID, err = soap.MarshalString(ObjectID); err != nil {
		return
	}
	if request.BrowseFlag, err = soap.MarshalString(BrowseFlag); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(StartingIndex); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(RequestedCount); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(SortCriteria); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "Browse", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(response.NumberReturned); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(response.TotalMatches); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Browse is the legacy version of BrowseCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) Browse(ObjectID string, BrowseFlag string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseCtx(context.Background(),
		ObjectID,
		BrowseFlag,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory4) SearchCtx(
	ctx context.Context,
	ContainerID string,
	SearchCriteria string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		ContainerID    string
		SearchCriteria string
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ContainerID, err = soap.MarshalString(ContainerID); err != nil {
		return
	}
	if request.SearchCriteria, err = soap.MarshalString(SearchCriteria); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	if request.StartingIndex, err = soap.MarshalUi4(StartingIndex); err != nil {
		return
	}
	if request.RequestedCount, err = soap.MarshalUi4(RequestedCount); err != nil {
		return
	}
	if request.SortCriteria, err = soap.MarshalString(SortCriteria); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result         string
		NumberReturned string
		TotalMatches   string
		UpdateID       string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "Search", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if NumberReturned, err = soap.UnmarshalUi4(response.NumberReturned); err != nil {
		return
	}
	if TotalMatches, err = soap.UnmarshalUi4(response.TotalMatches); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Search is the legacy version of SearchCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) Search(ContainerID string, SearchCriteria string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.SearchCtx(context.Background(),
		ContainerID,
		SearchCriteria,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory4) CreateObjectCtx(
	ctx context.Context,
	ContainerID string,
	Elements string,
) (ObjectID string, Result string, err error) {
	// Request structure.
	request := &struct {
		ContainerID string
		Elements    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ContainerID, err = soap.MarshalString(ContainerID); err != nil {
		return
	}
	if request.Elements, err = soap.MarshalString(Elements); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ObjectID string
		Result   string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "CreateObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ObjectID, err = soap.UnmarshalString(response.ObjectID); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// CreateObject is the legacy version of CreateObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) CreateObject(ContainerID string, Elements string) (ObjectID string, Result string, err error) {
	return client.CreateObjectCtx(context.Background(),
		ContainerID,
		Elements,
	)
}

func (client *ContentDirectory4) DestroyObjectCtx(
	ctx context.Context,
	ObjectID string,
) (err error) {
	// Request structure.
	request := &struct {
		ObjectID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ObjectID, err = soap.MarshalString(ObjectID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "DestroyObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DestroyObject is the legacy version of DestroyObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) DestroyObject(ObjectID string) (err error) {
	return client.DestroyObjectCtx(context.Background(),
		ObjectID,
	)
}

func (client *ContentDirectory4) UpdateObjectCtx(
	ctx context.Context,
	ObjectID string,
	CurrentTagValue string,
	NewTagValue string,
) (err error) {
	// Request structure.
	request := &struct {
		ObjectID        string
		CurrentTagValue string
		NewTagValue     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ObjectID, err = soap.MarshalString(ObjectID); err != nil {
		return
	}
	if request.CurrentTagValue, err = soap.MarshalString(CurrentTagValue); err != nil {
		return
	}
	if request.NewTagValue, err = soap.MarshalString(NewTagValue); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "UpdateObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// UpdateObject is the legacy version of UpdateObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) UpdateObject(ObjectID string, CurrentTagValue string, NewTagValue string) (err error) {
	return client.UpdateObjectCtx(context.Background(),
		ObjectID,
		CurrentTagValue,
		NewTagValue,
	)
}

func (client *ContentDirectory4) MoveObjectCtx(
	ctx context.Context,
	ObjectID string,
	NewParentID string,
) (NewObjectID string, err error) {
	// Request structure.
	request := &struct {
		ObjectID    string
		NewParentID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ObjectID, err = soap.MarshalString(ObjectID); err != nil {
		return
	}
	if request.NewParentID, err = soap.MarshalString(NewParentID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewObjectID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "MoveObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewObjectID, err = soap.UnmarshalString(response.NewObjectID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// MoveObject is the legacy version of MoveObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) MoveObject(ObjectID string, NewParentID string) (NewObjectID string, err error) {
	return client.MoveObjectCtx(context.Background(),
		ObjectID,
		NewParentID,
	)
}

func (client *ContentDirectory4) ImportResourceCtx(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (TransferID uint32, err error) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}
	// BEGIN Marshal arguments into request.

	if request.SourceURI, err = soap.MarshalURI(SourceURI); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(DestinationURI); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "ImportResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransferID, err = soap.UnmarshalUi4(response.TransferID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ImportResource is the legacy version of ImportResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) ImportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ImportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory4) ExportResourceCtx(
	ctx context.Context,
	SourceURI *url.URL,
	DestinationURI *url.URL,
) (TransferID uint32, err error) {
	// Request structure.
	request := &struct {
		SourceURI      string
		DestinationURI string
	}{}
	// BEGIN Marshal arguments into request.

	if request.SourceURI, err = soap.MarshalURI(SourceURI); err != nil {
		return
	}
	if request.DestinationURI, err = soap.MarshalURI(DestinationURI); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransferID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "ExportResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransferID, err = soap.UnmarshalUi4(response.TransferID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ExportResource is the legacy version of ExportResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) ExportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ExportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory4) DeleteResourceCtx(
	ctx context.Context,
	ResourceURI *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		ResourceURI string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ResourceURI, err = soap.MarshalURI(ResourceURI); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "DeleteResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteResource is the legacy version of DeleteResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) DeleteResource(ResourceURI *url.URL) (err error) {
	return client.DeleteResourceCtx(context.Background(),
		ResourceURI,
	)
}

func (client *ContentDirectory4) StopTransferResourceCtx(
	ctx context.Context,
	TransferID uint32,
) (err error) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransferID, err = soap.MarshalUi4(TransferID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "StopTransferResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// StopTransferResource is the legacy version of StopTransferResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) StopTransferResource(TransferID uint32) (err error) {
	return client.StopTransferResourceCtx(context.Background(),
		TransferID,
	)
}

//
// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory4) GetTransferProgressCtx(
	ctx context.Context,
	TransferID uint32,
) (TransferStatus string, TransferLength string, TransferTotal string, err error) {
	// Request structure.
	request := &struct {
		TransferID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransferID, err = soap.MarshalUi4(TransferID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransferStatus string
		TransferLength string
		TransferTotal  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetTransferProgress", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransferStatus, err = soap.UnmarshalString(response.TransferStatus); err != nil {
		return
	}
	if TransferLength, err = soap.UnmarshalString(response.TransferLength); err != nil {
		return
	}
	if TransferTotal, err = soap.UnmarshalString(response.TransferTotal); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransferProgress is the legacy version of GetTransferProgressCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetTransferProgress(TransferID uint32) (TransferStatus string, TransferLength string, TransferTotal string, err error) {
	return client.GetTransferProgressCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory4) CreateReferenceCtx(
	ctx context.Context,
	ContainerID string,
	ObjectID string,
) (NewID string, err error) {
	// Request structure.
	request := &struct {
		ContainerID string
		ObjectID    string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ContainerID, err = soap.MarshalString(ContainerID); err != nil {
		return
	}
	if request.ObjectID, err = soap.MarshalString(ObjectID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "CreateReference", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewID, err = soap.UnmarshalString(response.NewID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// CreateReference is the legacy version of CreateReferenceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) CreateReference(ContainerID string, ObjectID string) (NewID string, err error) {
	return client.CreateReferenceCtx(context.Background(),
		ContainerID,
		ObjectID,
	)
}

func (client *ContentDirectory4) FreeFormQueryCtx(
	ctx context.Context,
	ContainerID string,
	CDSView uint32,
	QueryRequest string,
) (QueryResult string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		ContainerID  string
		CDSView      string
		QueryRequest string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ContainerID, err = soap.MarshalString(ContainerID); err != nil {
		return
	}
	if request.CDSView, err = soap.MarshalUi4(CDSView); err != nil {
		return
	}
	if request.QueryRequest, err = soap.MarshalString(QueryRequest); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		QueryResult string
		UpdateID    string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "FreeFormQuery", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if QueryResult, err = soap.UnmarshalString(response.QueryResult); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// FreeFormQuery is the legacy version of FreeFormQueryCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) FreeFormQuery(ContainerID string, CDSView uint32, QueryRequest string) (QueryResult string, UpdateID uint32, err error) {
	return client.FreeFormQueryCtx(context.Background(),
		ContainerID,
		CDSView,
		QueryRequest,
	)
}

func (client *ContentDirectory4) GetFreeFormQueryCapabilitiesCtx(
	ctx context.Context,
) (FFQCapabilities string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		FFQCapabilities string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetFreeFormQueryCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if FFQCapabilities, err = soap.UnmarshalString(response.FFQCapabilities); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetFreeFormQueryCapabilities is the legacy version of GetFreeFormQueryCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetFreeFormQueryCapabilities() (FFQCapabilities string, err error) {
	return client.GetFreeFormQueryCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory4) GetPermissionsInfoCtx(
	ctx context.Context,
) (PermissionsInfo string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PermissionsInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetPermissionsInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PermissionsInfo, err = soap.UnmarshalString(response.PermissionsInfo); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPermissionsInfo is the legacy version of GetPermissionsInfoCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetPermissionsInfo() (PermissionsInfo string, err error) {
	return client.GetPermissionsInfoCtx(context.Background())
}

func (client *ContentDirectory4) RequestDeviceModeCtx(
	ctx context.Context,
	CPID string,
	DeviceModeRequest string,
) (DeviceModeID string, DeviceModeStatus string, err error) {
	// Request structure.
	request := &struct {
		CPID              string
		DeviceModeRequest string
	}{}
	// BEGIN Marshal arguments into request.

	if request.CPID, err = soap.MarshalString(CPID); err != nil {
		return
	}
	if request.DeviceModeRequest, err = soap.MarshalString(DeviceModeRequest); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		DeviceModeID     string
		DeviceModeStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "RequestDeviceMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if DeviceModeID, err = soap.UnmarshalString(response.DeviceModeID); err != nil {
		return
	}
	if DeviceModeStatus, err = soap.UnmarshalString(response.DeviceModeStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// RequestDeviceMode is the legacy version of RequestDeviceModeCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) RequestDeviceMode(CPID string, DeviceModeRequest string) (DeviceModeID string, DeviceModeStatus string, err error) {
	return client.RequestDeviceModeCtx(context.Background(),
		CPID,
		DeviceModeRequest,
	)
}

func (client *ContentDirectory4) GetDeviceModeCtx(
	ctx context.Context,
) (DeviceMode string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		DeviceMode string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetDeviceMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if DeviceMode, err = soap.UnmarshalString(response.DeviceMode); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDeviceMode is the legacy version of GetDeviceModeCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetDeviceMode() (DeviceMode string, err error) {
	return client.GetDeviceModeCtx(context.Background())
}

func (client *ContentDirectory4) GetDeviceModeStatusCtx(
	ctx context.Context,
) (DeviceModeStatus string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		DeviceModeStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetDeviceModeStatus", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if DeviceModeStatus, err = soap.UnmarshalString(response.DeviceModeStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDeviceModeStatus is the legacy version of GetDeviceModeStatusCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetDeviceModeStatus() (DeviceModeStatus string, err error) {
	return client.GetDeviceModeStatusCtx(context.Background())
}

func (client *ContentDirectory4) CancelDeviceModeCtx(
	ctx context.Context,
	DeviceModeID string,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceModeID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceModeID, err = soap.MarshalString(DeviceModeID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "CancelDeviceMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// CancelDeviceMode is the legacy version of CancelDeviceModeCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) CancelDeviceMode(DeviceModeID string) (err error) {
	return client.CancelDeviceModeCtx(context.Background(),
		DeviceModeID,
	)
}

func (client *ContentDirectory4) ExtendDeviceModeCtx(
	ctx context.Context,
	DeviceModeID string,
	DeviceModeRequest string,
) (DeviceModeStatus string, err error) {
	// Request structure.
	request := &struct {
		DeviceModeID      string
		DeviceModeRequest string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceModeID, err = soap.MarshalString(DeviceModeID); err != nil {
		return
	}
	if request.DeviceModeRequest, err = soap.MarshalString(DeviceModeRequest); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		DeviceModeStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "ExtendDeviceMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if DeviceModeStatus, err = soap.UnmarshalString(response.DeviceModeStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ExtendDeviceMode is the legacy version of ExtendDeviceModeCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) ExtendDeviceMode(DeviceModeID string, DeviceModeRequest string) (DeviceModeStatus string, err error) {
	return client.ExtendDeviceModeCtx(context.Background(),
		DeviceModeID,
		DeviceModeRequest,
	)
}

func (client *ContentDirectory4) GetAllAvailableTransformsCtx(
	ctx context.Context,
) (AllAvailableTransforms string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		AllAvailableTransforms string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetAllAvailableTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if AllAvailableTransforms, err = soap.UnmarshalString(response.AllAvailableTransforms); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllAvailableTransforms is the legacy version of GetAllAvailableTransformsCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetAllAvailableTransforms() (AllAvailableTransforms string, err error) {
	return client.GetAllAvailableTransformsCtx(context.Background())
}

func (client *ContentDirectory4) GetAllowedTransformsCtx(
	ctx context.Context,
	TransformResourceObjectDesc string,
) (AllowedTransforms string, err error) {
	// Request structure.
	request := &struct {
		TransformResourceObjectDesc string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformResourceObjectDesc, err = soap.MarshalString(TransformResourceObjectDesc); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		AllowedTransforms string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetAllowedTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if AllowedTransforms, err = soap.UnmarshalString(response.AllowedTransforms); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllowedTransforms is the legacy version of GetAllowedTransformsCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetAllowedTransforms(TransformResourceObjectDesc string) (AllowedTransforms string, err error) {
	return client.GetAllowedTransformsCtx(context.Background(),
		TransformResourceObjectDesc,
	)
}

func (client *ContentDirectory4) GetCurrentTransformStatusListCtx(
	ctx context.Context,
) (TransformStatus string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransformStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetCurrentTransformStatusList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransformStatus, err = soap.UnmarshalString(response.TransformStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetCurrentTransformStatusList is the legacy version of GetCurrentTransformStatusListCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetCurrentTransformStatusList() (TransformStatus string, err error) {
	return client.GetCurrentTransformStatusListCtx(context.Background())
}

func (client *ContentDirectory4) StartTransformTaskCtx(
	ctx context.Context,
	TransformResourceDesc string,
	TransformSettings string,
	TransformOverwrite bool,
	TransformRollback bool,
) (TransformTaskID string, err error) {
	// Request structure.
	request := &struct {
		TransformResourceDesc string
		TransformSettings     string
		TransformOverwrite    string
		TransformRollback     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformResourceDesc, err = soap.MarshalString(TransformResourceDesc); err != nil {
		return
	}
	if request.TransformSettings, err = soap.MarshalString(TransformSettings); err != nil {
		return
	}
	if request.TransformOverwrite, err = soap.MarshalBoolean(TransformOverwrite); err != nil {
		return
	}
	if request.TransformRollback, err = soap.MarshalBoolean(TransformRollback); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransformTaskID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "StartTransformTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransformTaskID, err = soap.UnmarshalString(response.TransformTaskID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// StartTransformTask is the legacy version of StartTransformTaskCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) StartTransformTask(TransformResourceDesc string, TransformSettings string, TransformOverwrite bool, TransformRollback bool) (TransformTaskID string, err error) {
	return client.StartTransformTaskCtx(context.Background(),
		TransformResourceDesc,
		TransformSettings,
		TransformOverwrite,
		TransformRollback,
	)
}

func (client *ContentDirectory4) GetTransformsCtx(
	ctx context.Context,
	TransformTaskID string,
) (CurrentTransformSettings string, err error) {
	// Request structure.
	request := &struct {
		TransformTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CurrentTransformSettings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CurrentTransformSettings, err = soap.UnmarshalString(response.CurrentTransformSettings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransforms is the legacy version of GetTransformsCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetTransforms(TransformTaskID string) (CurrentTransformSettings string, err error) {
	return client.GetTransformsCtx(context.Background(),
		TransformTaskID,
	)
}

func (client *ContentDirectory4) GetTransformTaskResultCtx(
	ctx context.Context,
	TransformTaskID string,
	TransformTaskResultFilter string,
) (TransformTaskResult string, err error) {
	// Request structure.
	request := &struct {
		TransformTaskID           string
		TransformTaskResultFilter string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	if request.TransformTaskResultFilter, err = soap.MarshalString(TransformTaskResultFilter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TransformTaskResult string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "GetTransformTaskResult", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TransformTaskResult, err = soap.UnmarshalString(response.TransformTaskResult); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetTransformTaskResult is the legacy version of GetTransformTaskResultCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) GetTransformTaskResult(TransformTaskID string, TransformTaskResultFilter string) (TransformTaskResult string, err error) {
	return client.GetTransformTaskResultCtx(context.Background(),
		TransformTaskID,
		TransformTaskResultFilter,
	)
}

func (client *ContentDirectory4) CancelTransformTaskCtx(
	ctx context.Context,
	TransformTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		TransformTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "CancelTransformTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// CancelTransformTask is the legacy version of CancelTransformTaskCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) CancelTransformTask(TransformTaskID string) (err error) {
	return client.CancelTransformTaskCtx(context.Background(),
		TransformTaskID,
	)
}

func (client *ContentDirectory4) PauseTransformTaskCtx(
	ctx context.Context,
	TransformTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		TransformTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "PauseTransformTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PauseTransformTask is the legacy version of PauseTransformTaskCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) PauseTransformTask(TransformTaskID string) (err error) {
	return client.PauseTransformTaskCtx(context.Background(),
		TransformTaskID,
	)
}

func (client *ContentDirectory4) ResumeTransformTaskCtx(
	ctx context.Context,
	TransformTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		TransformTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "ResumeTransformTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ResumeTransformTask is the legacy version of ResumeTransformTaskCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) ResumeTransformTask(TransformTaskID string) (err error) {
	return client.ResumeTransformTaskCtx(context.Background(),
		TransformTaskID,
	)
}

func (client *ContentDirectory4) RollbackTransformTaskCtx(
	ctx context.Context,
	TransformTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		TransformTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformTaskID, err = soap.MarshalString(TransformTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "RollbackTransformTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// RollbackTransformTask is the legacy version of RollbackTransformTaskCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) RollbackTransformTask(TransformTaskID string) (err error) {
	return client.RollbackTransformTaskCtx(context.Background(),
		TransformTaskID,
	)
}

func (client *ContentDirectory4) EvaluateTransformsCtx(
	ctx context.Context,
	TransformResourceDesc string,
	TransformSettings string,
) (EvaluationResult string, err error) {
	// Request structure.
	request := &struct {
		TransformResourceDesc string
		TransformSettings     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.TransformResourceDesc, err = soap.MarshalString(TransformResourceDesc); err != nil {
		return
	}
	if request.TransformSettings, err = soap.MarshalString(TransformSettings); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		EvaluationResult string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_4, "EvaluateTransforms", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if EvaluationResult, err = soap.UnmarshalString(response.EvaluationResult); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// EvaluateTransforms is the legacy version of EvaluateTransformsCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory4) EvaluateTransforms(TransformResourceDesc string, TransformSettings string) (EvaluationResult string, err error) {
	return client.EvaluateTransformsCtx(context.Background(),
		TransformResourceDesc,
		TransformSettings,
	)
}
