// Client for UPnP Device Control Protocol MediaServer v1 and MediaRenderer v1.
//
// This DCP is documented in detail at:
// - http://upnp.org/specs/av/av1/
//
// Typically, use one of the New* functions to create clients for services.
package av1

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
	URN_AVTransport_1        = "urn:schemas-upnp-org:service:AVTransport:1"
	URN_AVTransport_2        = "urn:schemas-upnp-org:service:AVTransport:2"
	URN_ConnectionManager_1  = "urn:schemas-upnp-org:service:ConnectionManager:1"
	URN_ConnectionManager_2  = "urn:schemas-upnp-org:service:ConnectionManager:2"
	URN_ContentDirectory_1   = "urn:schemas-upnp-org:service:ContentDirectory:1"
	URN_ContentDirectory_2   = "urn:schemas-upnp-org:service:ContentDirectory:2"
	URN_ContentDirectory_3   = "urn:schemas-upnp-org:service:ContentDirectory:3"
	URN_RenderingControl_1   = "urn:schemas-upnp-org:service:RenderingControl:1"
	URN_RenderingControl_2   = "urn:schemas-upnp-org:service:RenderingControl:2"
	URN_ScheduledRecording_1 = "urn:schemas-upnp-org:service:ScheduledRecording:1"
	URN_ScheduledRecording_2 = "urn:schemas-upnp-org:service:ScheduledRecording:2"
)

// AVTransport1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport1 struct {
	goupnp.ServiceClient
}

// NewAVTransport1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport1ClientsCtx(ctx context.Context) (clients []*AVTransport1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_AVTransport_1); err != nil {
		return
	}
	clients = newAVTransport1ClientsFromGenericClients(genericClients)
	return
}

// NewAVTransport1Clients is the legacy version of NewAVTransport1ClientsCtx, but uses
// context.Background() as the context.
func NewAVTransport1Clients() (clients []*AVTransport1, errors []error, err error) {
	return NewAVTransport1ClientsCtx(context.Background())
}

// NewAVTransport1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*AVTransport1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_AVTransport_1)
	if err != nil {
		return nil, err
	}
	return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport1ClientsByURL is the legacy version of NewAVTransport1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewAVTransport1ClientsByURL(loc *url.URL) ([]*AVTransport1, error) {
	return NewAVTransport1ClientsByURLCtx(context.Background(), loc)
}

// NewAVTransport1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_1)
	if err != nil {
		return nil, err
	}
	return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

func newAVTransport1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport1 {
	clients := make([]*AVTransport1, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport1{genericClients[i]}
	}
	return clients
}

func NewAVTransport1MultiClients(mutiClients goupnp.MultiServiceClient) []*AVTransport1 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_AVTransport_1))
	clients := make([]*AVTransport1, len(services))
	for i := range services {
		clients[i] = &AVTransport1{services[i]}
	}
	return clients
}

func (client *AVTransport1) GetCurrentTransportActionsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetCurrentTransportActions", request, response); err != nil {
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
func (client *AVTransport1) GetCurrentTransportActions(InstanceID uint32) (Actions string, err error) {
	return client.GetCurrentTransportActionsCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport1) GetDeviceCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetDeviceCapabilities", request, response); err != nil {
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
func (client *AVTransport1) GetDeviceCapabilities(InstanceID uint32) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
	return client.GetDeviceCapabilitiesCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport1) GetMediaInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetMediaInfo", request, response); err != nil {
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
func (client *AVTransport1) GetMediaInfo(InstanceID uint32) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	return client.GetMediaInfoCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * Track: allowed value range: minimum=0, step=1
func (client *AVTransport1) GetPositionInfoCtx(
	ctx context.Context,
	InstanceID uint32,
) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetPositionInfo", request, response); err != nil {
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
	if AbsCount, err = soap.UnmarshalI4(response.AbsCount); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPositionInfo is the legacy version of GetPositionInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetPositionInfo(InstanceID uint32) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
	return client.GetPositionInfoCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PLAYING
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport1) GetTransportInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetTransportInfo", request, response); err != nil {
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
func (client *AVTransport1) GetTransportInfo(InstanceID uint32) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
	return client.GetTransportInfoCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * PlayMode: allowed values: NORMAL
func (client *AVTransport1) GetTransportSettingsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetTransportSettings", request, response); err != nil {
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
func (client *AVTransport1) GetTransportSettings(InstanceID uint32) (PlayMode string, RecQualityMode string, err error) {
	return client.GetTransportSettingsCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport1) NextCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Next", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Next is the legacy version of NextCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Next(InstanceID uint32) (err error) {
	return client.NextCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport1) PauseCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Pause(InstanceID uint32) (err error) {
	return client.PauseCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Speed: allowed values: 1

func (client *AVTransport1) PlayCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Play(InstanceID uint32, Speed string) (err error) {
	return client.PlayCtx(context.Background(),
		InstanceID,
		Speed,
	)
}

func (client *AVTransport1) PreviousCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Previous", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Previous is the legacy version of PreviousCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Previous(InstanceID uint32) (err error) {
	return client.PreviousCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport1) RecordCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Record", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Record is the legacy version of RecordCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Record(InstanceID uint32) (err error) {
	return client.RecordCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Unit: allowed values: TRACK_NR

func (client *AVTransport1) SeekCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Seek", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Seek is the legacy version of SeekCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Seek(InstanceID uint32, Unit string, Target string) (err error) {
	return client.SeekCtx(context.Background(),
		InstanceID,
		Unit,
		Target,
	)
}

func (client *AVTransport1) SetAVTransportURICtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAVTransportURI is the legacy version of SetAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetAVTransportURI(InstanceID uint32, CurrentURI string, CurrentURIMetaData string) (err error) {
	return client.SetAVTransportURICtx(context.Background(),
		InstanceID,
		CurrentURI,
		CurrentURIMetaData,
	)
}

func (client *AVTransport1) SetNextAVTransportURICtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetNextAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetNextAVTransportURI is the legacy version of SetNextAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetNextAVTransportURI(InstanceID uint32, NextURI string, NextURIMetaData string) (err error) {
	return client.SetNextAVTransportURICtx(context.Background(),
		InstanceID,
		NextURI,
		NextURIMetaData,
	)
}

//
// Arguments:
//
// * NewPlayMode: allowed values: NORMAL

func (client *AVTransport1) SetPlayModeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetPlayMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetPlayMode is the legacy version of SetPlayModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetPlayMode(InstanceID uint32, NewPlayMode string) (err error) {
	return client.SetPlayModeCtx(context.Background(),
		InstanceID,
		NewPlayMode,
	)
}

func (client *AVTransport1) SetRecordQualityModeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetRecordQualityMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRecordQualityMode is the legacy version of SetRecordQualityModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetRecordQualityMode(InstanceID uint32, NewRecordQualityMode string) (err error) {
	return client.SetRecordQualityModeCtx(context.Background(),
		InstanceID,
		NewRecordQualityMode,
	)
}

func (client *AVTransport1) StopCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Stop(InstanceID uint32) (err error) {
	return client.StopCtx(context.Background(),
		InstanceID,
	)
}

// AVTransport2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport2 struct {
	goupnp.ServiceClient
}

// NewAVTransport2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport2ClientsCtx(ctx context.Context) (clients []*AVTransport2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_AVTransport_2); err != nil {
		return
	}
	clients = newAVTransport2ClientsFromGenericClients(genericClients)
	return
}

// NewAVTransport2Clients is the legacy version of NewAVTransport2ClientsCtx, but uses
// context.Background() as the context.
func NewAVTransport2Clients() (clients []*AVTransport2, errors []error, err error) {
	return NewAVTransport2ClientsCtx(context.Background())
}

// NewAVTransport2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*AVTransport2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_AVTransport_2)
	if err != nil {
		return nil, err
	}
	return newAVTransport2ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport2ClientsByURL is the legacy version of NewAVTransport2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewAVTransport2ClientsByURL(loc *url.URL) ([]*AVTransport2, error) {
	return NewAVTransport2ClientsByURLCtx(context.Background(), loc)
}

// NewAVTransport2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_2)
	if err != nil {
		return nil, err
	}
	return newAVTransport2ClientsFromGenericClients(genericClients), nil
}

func newAVTransport2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport2 {
	clients := make([]*AVTransport2, len(genericClients))
	for i := range genericClients {
		clients[i] = &AVTransport2{genericClients[i]}
	}
	return clients
}

func NewAVTransport2MultiClients(mutiClients goupnp.MultiServiceClient) []*AVTransport2 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_AVTransport_2))
	clients := make([]*AVTransport2, len(services))
	for i := range services {
		clients[i] = &AVTransport2{services[i]}
	}
	return clients
}

func (client *AVTransport2) GetCurrentTransportActionsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetCurrentTransportActions", request, response); err != nil {
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
func (client *AVTransport2) GetCurrentTransportActions(InstanceID uint32) (Actions string, err error) {
	return client.GetCurrentTransportActionsCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentDRMState: allowed values: OK
func (client *AVTransport2) GetDRMStateCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetDRMState", request, response); err != nil {
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
func (client *AVTransport2) GetDRMState(InstanceID uint32) (CurrentDRMState string, err error) {
	return client.GetDRMStateCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport2) GetDeviceCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetDeviceCapabilities", request, response); err != nil {
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
func (client *AVTransport2) GetDeviceCapabilities(InstanceID uint32) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
	return client.GetDeviceCapabilitiesCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport2) GetMediaInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetMediaInfo", request, response); err != nil {
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
func (client *AVTransport2) GetMediaInfo(InstanceID uint32) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	return client.GetMediaInfoCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentType: allowed values: NO_MEDIA, TRACK_AWARE, TRACK_UNAWARE
//
// * NrTracks: allowed value range: minimum=0
func (client *AVTransport2) GetMediaInfo_ExtCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetMediaInfo_Ext", request, response); err != nil {
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
func (client *AVTransport2) GetMediaInfo_Ext(InstanceID uint32) (CurrentType string, NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
	return client.GetMediaInfo_ExtCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * Track: allowed value range: minimum=0, step=1
func (client *AVTransport2) GetPositionInfoCtx(
	ctx context.Context,
	InstanceID uint32,
) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetPositionInfo", request, response); err != nil {
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
	if AbsCount, err = soap.UnmarshalI4(response.AbsCount); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPositionInfo is the legacy version of GetPositionInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) GetPositionInfo(InstanceID uint32) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
	return client.GetPositionInfoCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport2) GetStateVariablesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetStateVariables", request, response); err != nil {
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
func (client *AVTransport2) GetStateVariables(InstanceID uint32, StateVariableList string) (StateVariableValuePairs string, err error) {
	return client.GetStateVariablesCtx(context.Background(),
		InstanceID,
		StateVariableList,
	)
}

// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PLAYING
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport2) GetTransportInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetTransportInfo", request, response); err != nil {
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
func (client *AVTransport2) GetTransportInfo(InstanceID uint32) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
	return client.GetTransportInfoCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * PlayMode: allowed values: NORMAL
func (client *AVTransport2) GetTransportSettingsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "GetTransportSettings", request, response); err != nil {
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
func (client *AVTransport2) GetTransportSettings(InstanceID uint32) (PlayMode string, RecQualityMode string, err error) {
	return client.GetTransportSettingsCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport2) NextCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Next", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Next is the legacy version of NextCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Next(InstanceID uint32) (err error) {
	return client.NextCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport2) PauseCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Pause(InstanceID uint32) (err error) {
	return client.PauseCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Speed: allowed values: 1

func (client *AVTransport2) PlayCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Play(InstanceID uint32, Speed string) (err error) {
	return client.PlayCtx(context.Background(),
		InstanceID,
		Speed,
	)
}

func (client *AVTransport2) PreviousCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Previous", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Previous is the legacy version of PreviousCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Previous(InstanceID uint32) (err error) {
	return client.PreviousCtx(context.Background(),
		InstanceID,
	)
}

func (client *AVTransport2) RecordCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Record", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Record is the legacy version of RecordCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Record(InstanceID uint32) (err error) {
	return client.RecordCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Unit: allowed values: TRACK_NR

func (client *AVTransport2) SeekCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Seek", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Seek is the legacy version of SeekCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Seek(InstanceID uint32, Unit string, Target string) (err error) {
	return client.SeekCtx(context.Background(),
		InstanceID,
		Unit,
		Target,
	)
}

func (client *AVTransport2) SetAVTransportURICtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "SetAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAVTransportURI is the legacy version of SetAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport2) SetAVTransportURI(InstanceID uint32, CurrentURI string, CurrentURIMetaData string) (err error) {
	return client.SetAVTransportURICtx(context.Background(),
		InstanceID,
		CurrentURI,
		CurrentURIMetaData,
	)
}

func (client *AVTransport2) SetNextAVTransportURICtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "SetNextAVTransportURI", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetNextAVTransportURI is the legacy version of SetNextAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport2) SetNextAVTransportURI(InstanceID uint32, NextURI string, NextURIMetaData string) (err error) {
	return client.SetNextAVTransportURICtx(context.Background(),
		InstanceID,
		NextURI,
		NextURIMetaData,
	)
}

//
// Arguments:
//
// * NewPlayMode: allowed values: NORMAL

func (client *AVTransport2) SetPlayModeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "SetPlayMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetPlayMode is the legacy version of SetPlayModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) SetPlayMode(InstanceID uint32, NewPlayMode string) (err error) {
	return client.SetPlayModeCtx(context.Background(),
		InstanceID,
		NewPlayMode,
	)
}

func (client *AVTransport2) SetRecordQualityModeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "SetRecordQualityMode", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRecordQualityMode is the legacy version of SetRecordQualityModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) SetRecordQualityMode(InstanceID uint32, NewRecordQualityMode string) (err error) {
	return client.SetRecordQualityModeCtx(context.Background(),
		InstanceID,
		NewRecordQualityMode,
	)
}

func (client *AVTransport2) SetStateVariablesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "SetStateVariables", request, response); err != nil {
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
func (client *AVTransport2) SetStateVariables(InstanceID uint32, AVTransportUDN string, ServiceType string, ServiceId string, StateVariableValuePairs string) (StateVariableList string, err error) {
	return client.SetStateVariablesCtx(context.Background(),
		InstanceID,
		AVTransportUDN,
		ServiceType,
		ServiceId,
		StateVariableValuePairs,
	)
}

func (client *AVTransport2) StopCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_2, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *AVTransport2) Stop(InstanceID uint32) (err error) {
	return client.StopCtx(context.Background(),
		InstanceID,
	)
}

// ConnectionManager1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager1 struct {
	goupnp.ServiceClient
}

// NewConnectionManager1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager1ClientsCtx(ctx context.Context) (clients []*ConnectionManager1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ConnectionManager_1); err != nil {
		return
	}
	clients = newConnectionManager1ClientsFromGenericClients(genericClients)
	return
}

// NewConnectionManager1Clients is the legacy version of NewConnectionManager1ClientsCtx, but uses
// context.Background() as the context.
func NewConnectionManager1Clients() (clients []*ConnectionManager1, errors []error, err error) {
	return NewConnectionManager1ClientsCtx(context.Background())
}

// NewConnectionManager1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ConnectionManager1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ConnectionManager_1)
	if err != nil {
		return nil, err
	}
	return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager1ClientsByURL is the legacy version of NewConnectionManager1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewConnectionManager1ClientsByURL(loc *url.URL) ([]*ConnectionManager1, error) {
	return NewConnectionManager1ClientsByURLCtx(context.Background(), loc)
}

// NewConnectionManager1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_1)
	if err != nil {
		return nil, err
	}
	return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager1 {
	clients := make([]*ConnectionManager1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager1{genericClients[i]}
	}
	return clients
}

func NewConnectionManager1MultiClients(mutiClients goupnp.MultiServiceClient) []*ConnectionManager1 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ConnectionManager_1))
	clients := make([]*ConnectionManager1, len(services))
	for i := range services {
		clients[i] = &ConnectionManager1{services[i]}
	}
	return clients
}

func (client *ConnectionManager1) ConnectionCompleteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "ConnectionComplete", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ConnectionComplete is the legacy version of ConnectionCompleteCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager1) ConnectionComplete(ConnectionID int32) (err error) {
	return client.ConnectionCompleteCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager1) GetCurrentConnectionIDsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetCurrentConnectionIDs", request, response); err != nil {
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
func (client *ConnectionManager1) GetCurrentConnectionIDs() (ConnectionIDs string, err error) {
	return client.GetCurrentConnectionIDsCtx(context.Background())
}

// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager1) GetCurrentConnectionInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetCurrentConnectionInfo", request, response); err != nil {
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
func (client *ConnectionManager1) GetCurrentConnectionInfo(ConnectionID int32) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
	return client.GetCurrentConnectionInfoCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager1) GetProtocolInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetProtocolInfo", request, response); err != nil {
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
func (client *ConnectionManager1) GetProtocolInfo() (Source string, Sink string, err error) {
	return client.GetProtocolInfoCtx(context.Background())
}

//
// Arguments:
//
// * Direction: allowed values: Input, Output

func (client *ConnectionManager1) PrepareForConnectionCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "PrepareForConnection", request, response); err != nil {
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
func (client *ConnectionManager1) PrepareForConnection(RemoteProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string) (ConnectionID int32, AVTransportID int32, RcsID int32, err error) {
	return client.PrepareForConnectionCtx(context.Background(),
		RemoteProtocolInfo,
		PeerConnectionManager,
		PeerConnectionID,
		Direction,
	)
}

// ConnectionManager2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager2 struct {
	goupnp.ServiceClient
}

// NewConnectionManager2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager2ClientsCtx(ctx context.Context) (clients []*ConnectionManager2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ConnectionManager_2); err != nil {
		return
	}
	clients = newConnectionManager2ClientsFromGenericClients(genericClients)
	return
}

// NewConnectionManager2Clients is the legacy version of NewConnectionManager2ClientsCtx, but uses
// context.Background() as the context.
func NewConnectionManager2Clients() (clients []*ConnectionManager2, errors []error, err error) {
	return NewConnectionManager2ClientsCtx(context.Background())
}

// NewConnectionManager2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ConnectionManager2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ConnectionManager_2)
	if err != nil {
		return nil, err
	}
	return newConnectionManager2ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager2ClientsByURL is the legacy version of NewConnectionManager2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewConnectionManager2ClientsByURL(loc *url.URL) ([]*ConnectionManager2, error) {
	return NewConnectionManager2ClientsByURLCtx(context.Background(), loc)
}

// NewConnectionManager2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_2)
	if err != nil {
		return nil, err
	}
	return newConnectionManager2ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager2 {
	clients := make([]*ConnectionManager2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ConnectionManager2{genericClients[i]}
	}
	return clients
}

func NewConnectionManager2MultiClients(mutiClients goupnp.MultiServiceClient) []*ConnectionManager2 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ConnectionManager_2))
	clients := make([]*ConnectionManager2, len(services))
	for i := range services {
		clients[i] = &ConnectionManager2{services[i]}
	}
	return clients
}

func (client *ConnectionManager2) ConnectionCompleteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_2, "ConnectionComplete", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ConnectionComplete is the legacy version of ConnectionCompleteCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager2) ConnectionComplete(ConnectionID int32) (err error) {
	return client.ConnectionCompleteCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager2) GetCurrentConnectionIDsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_2, "GetCurrentConnectionIDs", request, response); err != nil {
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
func (client *ConnectionManager2) GetCurrentConnectionIDs() (ConnectionIDs string, err error) {
	return client.GetCurrentConnectionIDsCtx(context.Background())
}

// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager2) GetCurrentConnectionInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_2, "GetCurrentConnectionInfo", request, response); err != nil {
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
func (client *ConnectionManager2) GetCurrentConnectionInfo(ConnectionID int32) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
	return client.GetCurrentConnectionInfoCtx(context.Background(),
		ConnectionID,
	)
}

func (client *ConnectionManager2) GetProtocolInfoCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_2, "GetProtocolInfo", request, response); err != nil {
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
func (client *ConnectionManager2) GetProtocolInfo() (Source string, Sink string, err error) {
	return client.GetProtocolInfoCtx(context.Background())
}

//
// Arguments:
//
// * Direction: allowed values: Input, Output

func (client *ConnectionManager2) PrepareForConnectionCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_2, "PrepareForConnection", request, response); err != nil {
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
func (client *ConnectionManager2) PrepareForConnection(RemoteProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string) (ConnectionID int32, AVTransportID int32, RcsID int32, err error) {
	return client.PrepareForConnectionCtx(context.Background(),
		RemoteProtocolInfo,
		PeerConnectionManager,
		PeerConnectionID,
		Direction,
	)
}

// ContentDirectory1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory1 struct {
	goupnp.ServiceClient
}

// NewContentDirectory1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory1ClientsCtx(ctx context.Context) (clients []*ContentDirectory1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ContentDirectory_1); err != nil {
		return
	}
	clients = newContentDirectory1ClientsFromGenericClients(genericClients)
	return
}

// NewContentDirectory1Clients is the legacy version of NewContentDirectory1ClientsCtx, but uses
// context.Background() as the context.
func NewContentDirectory1Clients() (clients []*ContentDirectory1, errors []error, err error) {
	return NewContentDirectory1ClientsCtx(context.Background())
}

// NewContentDirectory1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ContentDirectory1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ContentDirectory_1)
	if err != nil {
		return nil, err
	}
	return newContentDirectory1ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory1ClientsByURL is the legacy version of NewContentDirectory1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewContentDirectory1ClientsByURL(loc *url.URL) ([]*ContentDirectory1, error) {
	return NewContentDirectory1ClientsByURLCtx(context.Background(), loc)
}

// NewContentDirectory1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_1)
	if err != nil {
		return nil, err
	}
	return newContentDirectory1ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory1 {
	clients := make([]*ContentDirectory1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory1{genericClients[i]}
	}
	return clients
}

func NewContentDirectory1MultiClients(mutiClients goupnp.MultiServiceClient) []*ContentDirectory1 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ContentDirectory_1))
	clients := make([]*ContentDirectory1, len(services))
	for i := range services {
		clients[i] = &ContentDirectory1{services[i]}
	}
	return clients
}

//
// Arguments:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren

func (client *ContentDirectory1) BrowseCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "Browse", request, response); err != nil {
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
func (client *ContentDirectory1) Browse(ObjectID string, BrowseFlag string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseCtx(context.Background(),
		ObjectID,
		BrowseFlag,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory1) CreateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "CreateObject", request, response); err != nil {
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
func (client *ContentDirectory1) CreateObject(ContainerID string, Elements string) (ObjectID string, Result string, err error) {
	return client.CreateObjectCtx(context.Background(),
		ContainerID,
		Elements,
	)
}

func (client *ContentDirectory1) CreateReferenceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "CreateReference", request, response); err != nil {
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
func (client *ContentDirectory1) CreateReference(ContainerID string, ObjectID string) (NewID string, err error) {
	return client.CreateReferenceCtx(context.Background(),
		ContainerID,
		ObjectID,
	)
}

func (client *ContentDirectory1) DeleteResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "DeleteResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteResource is the legacy version of DeleteResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory1) DeleteResource(ResourceURI *url.URL) (err error) {
	return client.DeleteResourceCtx(context.Background(),
		ResourceURI,
	)
}

func (client *ContentDirectory1) DestroyObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "DestroyObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DestroyObject is the legacy version of DestroyObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory1) DestroyObject(ObjectID string) (err error) {
	return client.DestroyObjectCtx(context.Background(),
		ObjectID,
	)
}

func (client *ContentDirectory1) ExportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "ExportResource", request, response); err != nil {
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
func (client *ContentDirectory1) ExportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ExportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory1) GetSearchCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "GetSearchCapabilities", request, response); err != nil {
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
func (client *ContentDirectory1) GetSearchCapabilities() (SearchCaps string, err error) {
	return client.GetSearchCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory1) GetSortCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "GetSortCapabilities", request, response); err != nil {
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
func (client *ContentDirectory1) GetSortCapabilities() (SortCaps string, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory1) GetSystemUpdateIDCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "GetSystemUpdateID", request, response); err != nil {
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
func (client *ContentDirectory1) GetSystemUpdateID() (Id uint32, err error) {
	return client.GetSystemUpdateIDCtx(context.Background())
}

// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory1) GetTransferProgressCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "GetTransferProgress", request, response); err != nil {
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
func (client *ContentDirectory1) GetTransferProgress(TransferID uint32) (TransferStatus string, TransferLength string, TransferTotal string, err error) {
	return client.GetTransferProgressCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory1) ImportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "ImportResource", request, response); err != nil {
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
func (client *ContentDirectory1) ImportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ImportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory1) SearchCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "Search", request, response); err != nil {
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
func (client *ContentDirectory1) Search(ContainerID string, SearchCriteria string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.SearchCtx(context.Background(),
		ContainerID,
		SearchCriteria,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory1) StopTransferResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "StopTransferResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// StopTransferResource is the legacy version of StopTransferResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory1) StopTransferResource(TransferID uint32) (err error) {
	return client.StopTransferResourceCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory1) UpdateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_1, "UpdateObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// UpdateObject is the legacy version of UpdateObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory1) UpdateObject(ObjectID string, CurrentTagValue string, NewTagValue string) (err error) {
	return client.UpdateObjectCtx(context.Background(),
		ObjectID,
		CurrentTagValue,
		NewTagValue,
	)
}

// ContentDirectory2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory2 struct {
	goupnp.ServiceClient
}

// NewContentDirectory2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory2ClientsCtx(ctx context.Context) (clients []*ContentDirectory2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ContentDirectory_2); err != nil {
		return
	}
	clients = newContentDirectory2ClientsFromGenericClients(genericClients)
	return
}

// NewContentDirectory2Clients is the legacy version of NewContentDirectory2ClientsCtx, but uses
// context.Background() as the context.
func NewContentDirectory2Clients() (clients []*ContentDirectory2, errors []error, err error) {
	return NewContentDirectory2ClientsCtx(context.Background())
}

// NewContentDirectory2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ContentDirectory2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ContentDirectory_2)
	if err != nil {
		return nil, err
	}
	return newContentDirectory2ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory2ClientsByURL is the legacy version of NewContentDirectory2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewContentDirectory2ClientsByURL(loc *url.URL) ([]*ContentDirectory2, error) {
	return NewContentDirectory2ClientsByURLCtx(context.Background(), loc)
}

// NewContentDirectory2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_2)
	if err != nil {
		return nil, err
	}
	return newContentDirectory2ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory2 {
	clients := make([]*ContentDirectory2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory2{genericClients[i]}
	}
	return clients
}

func NewContentDirectory2MultiClients(mutiClients goupnp.MultiServiceClient) []*ContentDirectory2 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ContentDirectory_2))
	clients := make([]*ContentDirectory2, len(services))
	for i := range services {
		clients[i] = &ContentDirectory2{services[i]}
	}
	return clients
}

//
// Arguments:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren

func (client *ContentDirectory2) BrowseCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "Browse", request, response); err != nil {
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
func (client *ContentDirectory2) Browse(ObjectID string, BrowseFlag string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseCtx(context.Background(),
		ObjectID,
		BrowseFlag,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory2) CreateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "CreateObject", request, response); err != nil {
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
func (client *ContentDirectory2) CreateObject(ContainerID string, Elements string) (ObjectID string, Result string, err error) {
	return client.CreateObjectCtx(context.Background(),
		ContainerID,
		Elements,
	)
}

func (client *ContentDirectory2) CreateReferenceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "CreateReference", request, response); err != nil {
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
func (client *ContentDirectory2) CreateReference(ContainerID string, ObjectID string) (NewID string, err error) {
	return client.CreateReferenceCtx(context.Background(),
		ContainerID,
		ObjectID,
	)
}

func (client *ContentDirectory2) DeleteResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "DeleteResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteResource is the legacy version of DeleteResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory2) DeleteResource(ResourceURI *url.URL) (err error) {
	return client.DeleteResourceCtx(context.Background(),
		ResourceURI,
	)
}

func (client *ContentDirectory2) DestroyObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "DestroyObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DestroyObject is the legacy version of DestroyObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory2) DestroyObject(ObjectID string) (err error) {
	return client.DestroyObjectCtx(context.Background(),
		ObjectID,
	)
}

func (client *ContentDirectory2) ExportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "ExportResource", request, response); err != nil {
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
func (client *ContentDirectory2) ExportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ExportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory2) GetFeatureListCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetFeatureList", request, response); err != nil {
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
func (client *ContentDirectory2) GetFeatureList() (FeatureList string, err error) {
	return client.GetFeatureListCtx(context.Background())
}

func (client *ContentDirectory2) GetSearchCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetSearchCapabilities", request, response); err != nil {
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
func (client *ContentDirectory2) GetSearchCapabilities() (SearchCaps string, err error) {
	return client.GetSearchCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory2) GetSortCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetSortCapabilities", request, response); err != nil {
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
func (client *ContentDirectory2) GetSortCapabilities() (SortCaps string, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory2) GetSortExtensionCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetSortExtensionCapabilities", request, response); err != nil {
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
func (client *ContentDirectory2) GetSortExtensionCapabilities() (SortExtensionCaps string, err error) {
	return client.GetSortExtensionCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory2) GetSystemUpdateIDCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetSystemUpdateID", request, response); err != nil {
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
func (client *ContentDirectory2) GetSystemUpdateID() (Id uint32, err error) {
	return client.GetSystemUpdateIDCtx(context.Background())
}

// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory2) GetTransferProgressCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "GetTransferProgress", request, response); err != nil {
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
func (client *ContentDirectory2) GetTransferProgress(TransferID uint32) (TransferStatus string, TransferLength string, TransferTotal string, err error) {
	return client.GetTransferProgressCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory2) ImportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "ImportResource", request, response); err != nil {
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
func (client *ContentDirectory2) ImportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ImportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory2) MoveObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "MoveObject", request, response); err != nil {
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
func (client *ContentDirectory2) MoveObject(ObjectID string, NewParentID string) (NewObjectID string, err error) {
	return client.MoveObjectCtx(context.Background(),
		ObjectID,
		NewParentID,
	)
}

func (client *ContentDirectory2) SearchCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "Search", request, response); err != nil {
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
func (client *ContentDirectory2) Search(ContainerID string, SearchCriteria string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.SearchCtx(context.Background(),
		ContainerID,
		SearchCriteria,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory2) StopTransferResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "StopTransferResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// StopTransferResource is the legacy version of StopTransferResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory2) StopTransferResource(TransferID uint32) (err error) {
	return client.StopTransferResourceCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory2) UpdateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_2, "UpdateObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// UpdateObject is the legacy version of UpdateObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory2) UpdateObject(ObjectID string, CurrentTagValue string, NewTagValue string) (err error) {
	return client.UpdateObjectCtx(context.Background(),
		ObjectID,
		CurrentTagValue,
		NewTagValue,
	)
}

// ContentDirectory3 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ContentDirectory:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ContentDirectory3 struct {
	goupnp.ServiceClient
}

// NewContentDirectory3ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewContentDirectory3ClientsCtx(ctx context.Context) (clients []*ContentDirectory3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ContentDirectory_3); err != nil {
		return
	}
	clients = newContentDirectory3ClientsFromGenericClients(genericClients)
	return
}

// NewContentDirectory3Clients is the legacy version of NewContentDirectory3ClientsCtx, but uses
// context.Background() as the context.
func NewContentDirectory3Clients() (clients []*ContentDirectory3, errors []error, err error) {
	return NewContentDirectory3ClientsCtx(context.Background())
}

// NewContentDirectory3ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewContentDirectory3ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ContentDirectory3, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ContentDirectory_3)
	if err != nil {
		return nil, err
	}
	return newContentDirectory3ClientsFromGenericClients(genericClients), nil
}

// NewContentDirectory3ClientsByURL is the legacy version of NewContentDirectory3ClientsByURLCtx, but uses
// context.Background() as the context.
func NewContentDirectory3ClientsByURL(loc *url.URL) ([]*ContentDirectory3, error) {
	return NewContentDirectory3ClientsByURLCtx(context.Background(), loc)
}

// NewContentDirectory3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewContentDirectory3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ContentDirectory3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ContentDirectory_3)
	if err != nil {
		return nil, err
	}
	return newContentDirectory3ClientsFromGenericClients(genericClients), nil
}

func newContentDirectory3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ContentDirectory3 {
	clients := make([]*ContentDirectory3, len(genericClients))
	for i := range genericClients {
		clients[i] = &ContentDirectory3{genericClients[i]}
	}
	return clients
}

func NewContentDirectory3MultiClients(mutiClients goupnp.MultiServiceClient) []*ContentDirectory3 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ContentDirectory_3))
	clients := make([]*ContentDirectory3, len(services))
	for i := range services {
		clients[i] = &ContentDirectory3{services[i]}
	}
	return clients
}

//
// Arguments:
//
// * BrowseFlag: allowed values: BrowseMetadata, BrowseDirectChildren

func (client *ContentDirectory3) BrowseCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "Browse", request, response); err != nil {
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
func (client *ContentDirectory3) Browse(ObjectID string, BrowseFlag string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseCtx(context.Background(),
		ObjectID,
		BrowseFlag,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory3) CreateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "CreateObject", request, response); err != nil {
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
func (client *ContentDirectory3) CreateObject(ContainerID string, Elements string) (ObjectID string, Result string, err error) {
	return client.CreateObjectCtx(context.Background(),
		ContainerID,
		Elements,
	)
}

func (client *ContentDirectory3) CreateReferenceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "CreateReference", request, response); err != nil {
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
func (client *ContentDirectory3) CreateReference(ContainerID string, ObjectID string) (NewID string, err error) {
	return client.CreateReferenceCtx(context.Background(),
		ContainerID,
		ObjectID,
	)
}

func (client *ContentDirectory3) DeleteResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "DeleteResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteResource is the legacy version of DeleteResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory3) DeleteResource(ResourceURI *url.URL) (err error) {
	return client.DeleteResourceCtx(context.Background(),
		ResourceURI,
	)
}

func (client *ContentDirectory3) DestroyObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "DestroyObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DestroyObject is the legacy version of DestroyObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory3) DestroyObject(ObjectID string) (err error) {
	return client.DestroyObjectCtx(context.Background(),
		ObjectID,
	)
}

func (client *ContentDirectory3) ExportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "ExportResource", request, response); err != nil {
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
func (client *ContentDirectory3) ExportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ExportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory3) FreeFormQueryCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "FreeFormQuery", request, response); err != nil {
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
func (client *ContentDirectory3) FreeFormQuery(ContainerID string, CDSView uint32, QueryRequest string) (QueryResult string, UpdateID uint32, err error) {
	return client.FreeFormQueryCtx(context.Background(),
		ContainerID,
		CDSView,
		QueryRequest,
	)
}

func (client *ContentDirectory3) GetFeatureListCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetFeatureList", request, response); err != nil {
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
func (client *ContentDirectory3) GetFeatureList() (FeatureList string, err error) {
	return client.GetFeatureListCtx(context.Background())
}

func (client *ContentDirectory3) GetFreeFormQueryCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetFreeFormQueryCapabilities", request, response); err != nil {
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
func (client *ContentDirectory3) GetFreeFormQueryCapabilities() (FFQCapabilities string, err error) {
	return client.GetFreeFormQueryCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory3) GetSearchCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetSearchCapabilities", request, response); err != nil {
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
func (client *ContentDirectory3) GetSearchCapabilities() (SearchCaps string, err error) {
	return client.GetSearchCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory3) GetServiceResetTokenCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetServiceResetToken", request, response); err != nil {
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
func (client *ContentDirectory3) GetServiceResetToken() (ResetToken string, err error) {
	return client.GetServiceResetTokenCtx(context.Background())
}

func (client *ContentDirectory3) GetSortCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetSortCapabilities", request, response); err != nil {
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
func (client *ContentDirectory3) GetSortCapabilities() (SortCaps string, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory3) GetSortExtensionCapabilitiesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetSortExtensionCapabilities", request, response); err != nil {
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
func (client *ContentDirectory3) GetSortExtensionCapabilities() (SortExtensionCaps string, err error) {
	return client.GetSortExtensionCapabilitiesCtx(context.Background())
}

func (client *ContentDirectory3) GetSystemUpdateIDCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetSystemUpdateID", request, response); err != nil {
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
func (client *ContentDirectory3) GetSystemUpdateID() (Id uint32, err error) {
	return client.GetSystemUpdateIDCtx(context.Background())
}

// Return values:
//
// * TransferStatus: allowed values: COMPLETED, ERROR, IN_PROGRESS, STOPPED
func (client *ContentDirectory3) GetTransferProgressCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "GetTransferProgress", request, response); err != nil {
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
func (client *ContentDirectory3) GetTransferProgress(TransferID uint32) (TransferStatus string, TransferLength string, TransferTotal string, err error) {
	return client.GetTransferProgressCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory3) ImportResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "ImportResource", request, response); err != nil {
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
func (client *ContentDirectory3) ImportResource(SourceURI *url.URL, DestinationURI *url.URL) (TransferID uint32, err error) {
	return client.ImportResourceCtx(context.Background(),
		SourceURI,
		DestinationURI,
	)
}

func (client *ContentDirectory3) MoveObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "MoveObject", request, response); err != nil {
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
func (client *ContentDirectory3) MoveObject(ObjectID string, NewParentID string) (NewObjectID string, err error) {
	return client.MoveObjectCtx(context.Background(),
		ObjectID,
		NewParentID,
	)
}

func (client *ContentDirectory3) SearchCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "Search", request, response); err != nil {
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
func (client *ContentDirectory3) Search(ContainerID string, SearchCriteria string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.SearchCtx(context.Background(),
		ContainerID,
		SearchCriteria,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ContentDirectory3) StopTransferResourceCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "StopTransferResource", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// StopTransferResource is the legacy version of StopTransferResourceCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory3) StopTransferResource(TransferID uint32) (err error) {
	return client.StopTransferResourceCtx(context.Background(),
		TransferID,
	)
}

func (client *ContentDirectory3) UpdateObjectCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ContentDirectory_3, "UpdateObject", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// UpdateObject is the legacy version of UpdateObjectCtx, but uses
// context.Background() as the context.
func (client *ContentDirectory3) UpdateObject(ObjectID string, CurrentTagValue string, NewTagValue string) (err error) {
	return client.UpdateObjectCtx(context.Background(),
		ObjectID,
		CurrentTagValue,
		NewTagValue,
	)
}

// RenderingControl1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl1 struct {
	goupnp.ServiceClient
}

// NewRenderingControl1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl1ClientsCtx(ctx context.Context) (clients []*RenderingControl1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_RenderingControl_1); err != nil {
		return
	}
	clients = newRenderingControl1ClientsFromGenericClients(genericClients)
	return
}

// NewRenderingControl1Clients is the legacy version of NewRenderingControl1ClientsCtx, but uses
// context.Background() as the context.
func NewRenderingControl1Clients() (clients []*RenderingControl1, errors []error, err error) {
	return NewRenderingControl1ClientsCtx(context.Background())
}

// NewRenderingControl1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*RenderingControl1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_RenderingControl_1)
	if err != nil {
		return nil, err
	}
	return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl1ClientsByURL is the legacy version of NewRenderingControl1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewRenderingControl1ClientsByURL(loc *url.URL) ([]*RenderingControl1, error) {
	return NewRenderingControl1ClientsByURLCtx(context.Background(), loc)
}

// NewRenderingControl1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_1)
	if err != nil {
		return nil, err
	}
	return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl1 {
	clients := make([]*RenderingControl1, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl1{genericClients[i]}
	}
	return clients
}

func NewRenderingControl1MultiClients(mutiClients goupnp.MultiServiceClient) []*RenderingControl1 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_RenderingControl_1))
	clients := make([]*RenderingControl1, len(services))
	for i := range services {
		clients[i] = &RenderingControl1{services[i]}
	}
	return clients
}

// Return values:
//
// * CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBlueVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetBlueVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl1) GetBlueVideoBlackLevel(InstanceID uint32) (CurrentBlueVideoBlackLevel uint16, err error) {
	return client.GetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBlueVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetBlueVideoGain", request, response); err != nil {
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
func (client *RenderingControl1) GetBlueVideoGain(InstanceID uint32) (CurrentBlueVideoGain uint16, err error) {
	return client.GetBlueVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetBrightnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetBrightness", request, response); err != nil {
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
func (client *RenderingControl1) GetBrightness(InstanceID uint32) (CurrentBrightness uint16, err error) {
	return client.GetBrightnessCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetColorTemperatureCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetColorTemperature", request, response); err != nil {
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
func (client *RenderingControl1) GetColorTemperature(InstanceID uint32) (CurrentColorTemperature uint16, err error) {
	return client.GetColorTemperatureCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetContrastCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetContrast", request, response); err != nil {
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
func (client *RenderingControl1) GetContrast(InstanceID uint32) (CurrentContrast uint16, err error) {
	return client.GetContrastCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetGreenVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetGreenVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl1) GetGreenVideoBlackLevel(InstanceID uint32) (CurrentGreenVideoBlackLevel uint16, err error) {
	return client.GetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetGreenVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetGreenVideoGain", request, response); err != nil {
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
func (client *RenderingControl1) GetGreenVideoGain(InstanceID uint32) (CurrentGreenVideoGain uint16, err error) {
	return client.GetGreenVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl1) GetHorizontalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetHorizontalKeystone", request, response); err != nil {
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
func (client *RenderingControl1) GetHorizontalKeystone(InstanceID uint32) (CurrentHorizontalKeystone int16, err error) {
	return client.GetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) GetLoudnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetLoudness", request, response); err != nil {
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
func (client *RenderingControl1) GetLoudness(InstanceID uint32, Channel string) (CurrentLoudness bool, err error) {
	return client.GetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) GetMuteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetMute", request, response); err != nil {
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
func (client *RenderingControl1) GetMute(InstanceID uint32, Channel string) (CurrentMute bool, err error) {
	return client.GetMuteCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

// Return values:
//
// * CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetRedVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetRedVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl1) GetRedVideoBlackLevel(InstanceID uint32) (CurrentRedVideoBlackLevel uint16, err error) {
	return client.GetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

func (client *RenderingControl1) GetRedVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetRedVideoGain", request, response); err != nil {
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
func (client *RenderingControl1) GetRedVideoGain(InstanceID uint32) (CurrentRedVideoGain uint16, err error) {
	return client.GetRedVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetSharpnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetSharpness", request, response); err != nil {
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
func (client *RenderingControl1) GetSharpness(InstanceID uint32) (CurrentSharpness uint16, err error) {
	return client.GetSharpnessCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentVerticalKeystone: allowed value range: step=1
func (client *RenderingControl1) GetVerticalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetVerticalKeystone", request, response); err != nil {
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
func (client *RenderingControl1) GetVerticalKeystone(InstanceID uint32) (CurrentVerticalKeystone int16, err error) {
	return client.GetVerticalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl1) GetVolumeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetVolume", request, response); err != nil {
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
func (client *RenderingControl1) GetVolume(InstanceID uint32, Channel string) (CurrentVolume uint16, err error) {
	return client.GetVolumeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) GetVolumeDBCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetVolumeDB", request, response); err != nil {
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
func (client *RenderingControl1) GetVolumeDB(InstanceID uint32, Channel string) (CurrentVolume int16, err error) {
	return client.GetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) GetVolumeDBRangeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetVolumeDBRange", request, response); err != nil {
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
func (client *RenderingControl1) GetVolumeDBRange(InstanceID uint32, Channel string) (MinValue int16, MaxValue int16, err error) {
	return client.GetVolumeDBRangeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

func (client *RenderingControl1) ListPresetsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "ListPresets", request, response); err != nil {
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
func (client *RenderingControl1) ListPresets(InstanceID uint32) (CurrentPresetNameList string, err error) {
	return client.ListPresetsCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * PresetName: allowed values: FactoryDefaults

func (client *RenderingControl1) SelectPresetCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SelectPreset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SelectPreset is the legacy version of SelectPresetCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SelectPreset(InstanceID uint32, PresetName string) (err error) {
	return client.SelectPresetCtx(context.Background(),
		InstanceID,
		PresetName,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetBlueVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetBlueVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoBlackLevel is the legacy version of SetBlueVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetBlueVideoBlackLevel(InstanceID uint32, DesiredBlueVideoBlackLevel uint16) (err error) {
	return client.SetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoBlackLevel,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetBlueVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetBlueVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoGain is the legacy version of SetBlueVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetBlueVideoGain(InstanceID uint32, DesiredBlueVideoGain uint16) (err error) {
	return client.SetBlueVideoGainCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoGain,
	)
}

//
// Arguments:
//
// * DesiredBrightness: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetBrightnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetBrightness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBrightness is the legacy version of SetBrightnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetBrightness(InstanceID uint32, DesiredBrightness uint16) (err error) {
	return client.SetBrightnessCtx(context.Background(),
		InstanceID,
		DesiredBrightness,
	)
}

//
// Arguments:
//
// * DesiredColorTemperature: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetColorTemperatureCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetColorTemperature", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetColorTemperature is the legacy version of SetColorTemperatureCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetColorTemperature(InstanceID uint32, DesiredColorTemperature uint16) (err error) {
	return client.SetColorTemperatureCtx(context.Background(),
		InstanceID,
		DesiredColorTemperature,
	)
}

//
// Arguments:
//
// * DesiredContrast: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetContrastCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetContrast", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetContrast is the legacy version of SetContrastCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetContrast(InstanceID uint32, DesiredContrast uint16) (err error) {
	return client.SetContrastCtx(context.Background(),
		InstanceID,
		DesiredContrast,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetGreenVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetGreenVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoBlackLevel is the legacy version of SetGreenVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetGreenVideoBlackLevel(InstanceID uint32, DesiredGreenVideoBlackLevel uint16) (err error) {
	return client.SetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoBlackLevel,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetGreenVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetGreenVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoGain is the legacy version of SetGreenVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetGreenVideoGain(InstanceID uint32, DesiredGreenVideoGain uint16) (err error) {
	return client.SetGreenVideoGainCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoGain,
	)
}

//
// Arguments:
//
// * DesiredHorizontalKeystone: allowed value range: step=1

func (client *RenderingControl1) SetHorizontalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetHorizontalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetHorizontalKeystone is the legacy version of SetHorizontalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetHorizontalKeystone(InstanceID uint32, DesiredHorizontalKeystone int16) (err error) {
	return client.SetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredHorizontalKeystone,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) SetLoudnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetLoudness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetLoudness is the legacy version of SetLoudnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetLoudness(InstanceID uint32, Channel string, DesiredLoudness bool) (err error) {
	return client.SetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredLoudness,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl1) SetMuteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetMute(InstanceID uint32, Channel string, DesiredMute bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredMute,
	)
}

//
// Arguments:
//
// * DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetRedVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetRedVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoBlackLevel is the legacy version of SetRedVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetRedVideoBlackLevel(InstanceID uint32, DesiredRedVideoBlackLevel uint16) (err error) {
	return client.SetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredRedVideoBlackLevel,
	)
}

func (client *RenderingControl1) SetRedVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetRedVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoGain is the legacy version of SetRedVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetRedVideoGain(InstanceID uint32, DesiredRedVideoGain uint16) (err error) {
	return client.SetRedVideoGainCtx(context.Background(),
		InstanceID,
		DesiredRedVideoGain,
	)
}

//
// Arguments:
//
// * DesiredSharpness: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetSharpnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetSharpness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSharpness is the legacy version of SetSharpnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetSharpness(InstanceID uint32, DesiredSharpness uint16) (err error) {
	return client.SetSharpnessCtx(context.Background(),
		InstanceID,
		DesiredSharpness,
	)
}

//
// Arguments:
//
// * DesiredVerticalKeystone: allowed value range: step=1

func (client *RenderingControl1) SetVerticalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetVerticalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVerticalKeystone is the legacy version of SetVerticalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetVerticalKeystone(InstanceID uint32, DesiredVerticalKeystone int16) (err error) {
	return client.SetVerticalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredVerticalKeystone,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master
//
// * DesiredVolume: allowed value range: minimum=0, step=1

func (client *RenderingControl1) SetVolumeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetVolume(InstanceID uint32, Channel string, DesiredVolume uint16) (err error) {
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

func (client *RenderingControl1) SetVolumeDBCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetVolumeDB", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolumeDB is the legacy version of SetVolumeDBCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetVolumeDB(InstanceID uint32, Channel string, DesiredVolume int16) (err error) {
	return client.SetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredVolume,
	)
}

// RenderingControl2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl2 struct {
	goupnp.ServiceClient
}

// NewRenderingControl2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl2ClientsCtx(ctx context.Context) (clients []*RenderingControl2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_RenderingControl_2); err != nil {
		return
	}
	clients = newRenderingControl2ClientsFromGenericClients(genericClients)
	return
}

// NewRenderingControl2Clients is the legacy version of NewRenderingControl2ClientsCtx, but uses
// context.Background() as the context.
func NewRenderingControl2Clients() (clients []*RenderingControl2, errors []error, err error) {
	return NewRenderingControl2ClientsCtx(context.Background())
}

// NewRenderingControl2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*RenderingControl2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_RenderingControl_2)
	if err != nil {
		return nil, err
	}
	return newRenderingControl2ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl2ClientsByURL is the legacy version of NewRenderingControl2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewRenderingControl2ClientsByURL(loc *url.URL) ([]*RenderingControl2, error) {
	return NewRenderingControl2ClientsByURLCtx(context.Background(), loc)
}

// NewRenderingControl2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_2)
	if err != nil {
		return nil, err
	}
	return newRenderingControl2ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl2 {
	clients := make([]*RenderingControl2, len(genericClients))
	for i := range genericClients {
		clients[i] = &RenderingControl2{genericClients[i]}
	}
	return clients
}

func NewRenderingControl2MultiClients(mutiClients goupnp.MultiServiceClient) []*RenderingControl2 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_RenderingControl_2))
	clients := make([]*RenderingControl2, len(services))
	for i := range services {
		clients[i] = &RenderingControl2{services[i]}
	}
	return clients
}

// Return values:
//
// * CurrentBlueVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBlueVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetBlueVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl2) GetBlueVideoBlackLevel(InstanceID uint32) (CurrentBlueVideoBlackLevel uint16, err error) {
	return client.GetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentBlueVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBlueVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetBlueVideoGain", request, response); err != nil {
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
func (client *RenderingControl2) GetBlueVideoGain(InstanceID uint32) (CurrentBlueVideoGain uint16, err error) {
	return client.GetBlueVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentBrightness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetBrightnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetBrightness", request, response); err != nil {
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
func (client *RenderingControl2) GetBrightness(InstanceID uint32) (CurrentBrightness uint16, err error) {
	return client.GetBrightnessCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentColorTemperature: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetColorTemperatureCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetColorTemperature", request, response); err != nil {
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
func (client *RenderingControl2) GetColorTemperature(InstanceID uint32) (CurrentColorTemperature uint16, err error) {
	return client.GetColorTemperatureCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentContrast: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetContrastCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetContrast", request, response); err != nil {
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
func (client *RenderingControl2) GetContrast(InstanceID uint32) (CurrentContrast uint16, err error) {
	return client.GetContrastCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentGreenVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetGreenVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetGreenVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl2) GetGreenVideoBlackLevel(InstanceID uint32) (CurrentGreenVideoBlackLevel uint16, err error) {
	return client.GetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentGreenVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetGreenVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetGreenVideoGain", request, response); err != nil {
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
func (client *RenderingControl2) GetGreenVideoGain(InstanceID uint32) (CurrentGreenVideoGain uint16, err error) {
	return client.GetGreenVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentHorizontalKeystone: allowed value range: step=1
func (client *RenderingControl2) GetHorizontalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetHorizontalKeystone", request, response); err != nil {
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
func (client *RenderingControl2) GetHorizontalKeystone(InstanceID uint32) (CurrentHorizontalKeystone int16, err error) {
	return client.GetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) GetLoudnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetLoudness", request, response); err != nil {
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
func (client *RenderingControl2) GetLoudness(InstanceID uint32, Channel string) (CurrentLoudness bool, err error) {
	return client.GetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) GetMuteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetMute", request, response); err != nil {
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
func (client *RenderingControl2) GetMute(InstanceID uint32, Channel string) (CurrentMute bool, err error) {
	return client.GetMuteCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

// Return values:
//
// * CurrentRedVideoBlackLevel: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetRedVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetRedVideoBlackLevel", request, response); err != nil {
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
func (client *RenderingControl2) GetRedVideoBlackLevel(InstanceID uint32) (CurrentRedVideoBlackLevel uint16, err error) {
	return client.GetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentRedVideoGain: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetRedVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetRedVideoGain", request, response); err != nil {
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
func (client *RenderingControl2) GetRedVideoGain(InstanceID uint32) (CurrentRedVideoGain uint16, err error) {
	return client.GetRedVideoGainCtx(context.Background(),
		InstanceID,
	)
}

// Return values:
//
// * CurrentSharpness: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetSharpnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetSharpness", request, response); err != nil {
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
func (client *RenderingControl2) GetSharpness(InstanceID uint32) (CurrentSharpness uint16, err error) {
	return client.GetSharpnessCtx(context.Background(),
		InstanceID,
	)
}

func (client *RenderingControl2) GetStateVariablesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetStateVariables", request, response); err != nil {
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
func (client *RenderingControl2) GetStateVariables(InstanceID uint32, StateVariableList string) (StateVariableValuePairs string, err error) {
	return client.GetStateVariablesCtx(context.Background(),
		InstanceID,
		StateVariableList,
	)
}

// Return values:
//
// * CurrentVerticalKeystone: allowed value range: step=1
func (client *RenderingControl2) GetVerticalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetVerticalKeystone", request, response); err != nil {
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
func (client *RenderingControl2) GetVerticalKeystone(InstanceID uint32) (CurrentVerticalKeystone int16, err error) {
	return client.GetVerticalKeystoneCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, step=1
func (client *RenderingControl2) GetVolumeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetVolume", request, response); err != nil {
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
func (client *RenderingControl2) GetVolume(InstanceID uint32, Channel string) (CurrentVolume uint16, err error) {
	return client.GetVolumeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) GetVolumeDBCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetVolumeDB", request, response); err != nil {
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
func (client *RenderingControl2) GetVolumeDB(InstanceID uint32, Channel string) (CurrentVolume int16, err error) {
	return client.GetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) GetVolumeDBRangeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "GetVolumeDBRange", request, response); err != nil {
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
func (client *RenderingControl2) GetVolumeDBRange(InstanceID uint32, Channel string) (MinValue int16, MaxValue int16, err error) {
	return client.GetVolumeDBRangeCtx(context.Background(),
		InstanceID,
		Channel,
	)
}

func (client *RenderingControl2) ListPresetsCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "ListPresets", request, response); err != nil {
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
func (client *RenderingControl2) ListPresets(InstanceID uint32) (CurrentPresetNameList string, err error) {
	return client.ListPresetsCtx(context.Background(),
		InstanceID,
	)
}

//
// Arguments:
//
// * PresetName: allowed values: FactoryDefaults

func (client *RenderingControl2) SelectPresetCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SelectPreset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SelectPreset is the legacy version of SelectPresetCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SelectPreset(InstanceID uint32, PresetName string) (err error) {
	return client.SelectPresetCtx(context.Background(),
		InstanceID,
		PresetName,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetBlueVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetBlueVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoBlackLevel is the legacy version of SetBlueVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetBlueVideoBlackLevel(InstanceID uint32, DesiredBlueVideoBlackLevel uint16) (err error) {
	return client.SetBlueVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoBlackLevel,
	)
}

//
// Arguments:
//
// * DesiredBlueVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetBlueVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetBlueVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBlueVideoGain is the legacy version of SetBlueVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetBlueVideoGain(InstanceID uint32, DesiredBlueVideoGain uint16) (err error) {
	return client.SetBlueVideoGainCtx(context.Background(),
		InstanceID,
		DesiredBlueVideoGain,
	)
}

//
// Arguments:
//
// * DesiredBrightness: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetBrightnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetBrightness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBrightness is the legacy version of SetBrightnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetBrightness(InstanceID uint32, DesiredBrightness uint16) (err error) {
	return client.SetBrightnessCtx(context.Background(),
		InstanceID,
		DesiredBrightness,
	)
}

//
// Arguments:
//
// * DesiredColorTemperature: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetColorTemperatureCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetColorTemperature", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetColorTemperature is the legacy version of SetColorTemperatureCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetColorTemperature(InstanceID uint32, DesiredColorTemperature uint16) (err error) {
	return client.SetColorTemperatureCtx(context.Background(),
		InstanceID,
		DesiredColorTemperature,
	)
}

//
// Arguments:
//
// * DesiredContrast: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetContrastCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetContrast", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetContrast is the legacy version of SetContrastCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetContrast(InstanceID uint32, DesiredContrast uint16) (err error) {
	return client.SetContrastCtx(context.Background(),
		InstanceID,
		DesiredContrast,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetGreenVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetGreenVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoBlackLevel is the legacy version of SetGreenVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetGreenVideoBlackLevel(InstanceID uint32, DesiredGreenVideoBlackLevel uint16) (err error) {
	return client.SetGreenVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoBlackLevel,
	)
}

//
// Arguments:
//
// * DesiredGreenVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetGreenVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetGreenVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetGreenVideoGain is the legacy version of SetGreenVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetGreenVideoGain(InstanceID uint32, DesiredGreenVideoGain uint16) (err error) {
	return client.SetGreenVideoGainCtx(context.Background(),
		InstanceID,
		DesiredGreenVideoGain,
	)
}

//
// Arguments:
//
// * DesiredHorizontalKeystone: allowed value range: step=1

func (client *RenderingControl2) SetHorizontalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetHorizontalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetHorizontalKeystone is the legacy version of SetHorizontalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetHorizontalKeystone(InstanceID uint32, DesiredHorizontalKeystone int16) (err error) {
	return client.SetHorizontalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredHorizontalKeystone,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) SetLoudnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetLoudness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetLoudness is the legacy version of SetLoudnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetLoudness(InstanceID uint32, Channel string, DesiredLoudness bool) (err error) {
	return client.SetLoudnessCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredLoudness,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master

func (client *RenderingControl2) SetMuteCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetMute(InstanceID uint32, Channel string, DesiredMute bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredMute,
	)
}

//
// Arguments:
//
// * DesiredRedVideoBlackLevel: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetRedVideoBlackLevelCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetRedVideoBlackLevel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoBlackLevel is the legacy version of SetRedVideoBlackLevelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetRedVideoBlackLevel(InstanceID uint32, DesiredRedVideoBlackLevel uint16) (err error) {
	return client.SetRedVideoBlackLevelCtx(context.Background(),
		InstanceID,
		DesiredRedVideoBlackLevel,
	)
}

//
// Arguments:
//
// * DesiredRedVideoGain: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetRedVideoGainCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetRedVideoGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRedVideoGain is the legacy version of SetRedVideoGainCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetRedVideoGain(InstanceID uint32, DesiredRedVideoGain uint16) (err error) {
	return client.SetRedVideoGainCtx(context.Background(),
		InstanceID,
		DesiredRedVideoGain,
	)
}

//
// Arguments:
//
// * DesiredSharpness: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetSharpnessCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetSharpness", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSharpness is the legacy version of SetSharpnessCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetSharpness(InstanceID uint32, DesiredSharpness uint16) (err error) {
	return client.SetSharpnessCtx(context.Background(),
		InstanceID,
		DesiredSharpness,
	)
}

func (client *RenderingControl2) SetStateVariablesCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetStateVariables", request, response); err != nil {
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
func (client *RenderingControl2) SetStateVariables(InstanceID uint32, RenderingControlUDN string, ServiceType string, ServiceId string, StateVariableValuePairs string) (StateVariableList string, err error) {
	return client.SetStateVariablesCtx(context.Background(),
		InstanceID,
		RenderingControlUDN,
		ServiceType,
		ServiceId,
		StateVariableValuePairs,
	)
}

//
// Arguments:
//
// * DesiredVerticalKeystone: allowed value range: step=1

func (client *RenderingControl2) SetVerticalKeystoneCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetVerticalKeystone", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVerticalKeystone is the legacy version of SetVerticalKeystoneCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetVerticalKeystone(InstanceID uint32, DesiredVerticalKeystone int16) (err error) {
	return client.SetVerticalKeystoneCtx(context.Background(),
		InstanceID,
		DesiredVerticalKeystone,
	)
}

//
// Arguments:
//
// * Channel: allowed values: Master
//
// * DesiredVolume: allowed value range: minimum=0, step=1

func (client *RenderingControl2) SetVolumeCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetVolume(InstanceID uint32, Channel string, DesiredVolume uint16) (err error) {
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

func (client *RenderingControl2) SetVolumeDBCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_2, "SetVolumeDB", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolumeDB is the legacy version of SetVolumeDBCtx, but uses
// context.Background() as the context.
func (client *RenderingControl2) SetVolumeDB(InstanceID uint32, Channel string, DesiredVolume int16) (err error) {
	return client.SetVolumeDBCtx(context.Background(),
		InstanceID,
		Channel,
		DesiredVolume,
	)
}

// ScheduledRecording1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ScheduledRecording:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ScheduledRecording1 struct {
	goupnp.ServiceClient
}

// NewScheduledRecording1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording1ClientsCtx(ctx context.Context) (clients []*ScheduledRecording1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ScheduledRecording_1); err != nil {
		return
	}
	clients = newScheduledRecording1ClientsFromGenericClients(genericClients)
	return
}

// NewScheduledRecording1Clients is the legacy version of NewScheduledRecording1ClientsCtx, but uses
// context.Background() as the context.
func NewScheduledRecording1Clients() (clients []*ScheduledRecording1, errors []error, err error) {
	return NewScheduledRecording1ClientsCtx(context.Background())
}

// NewScheduledRecording1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ScheduledRecording1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ScheduledRecording_1)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording1ClientsFromGenericClients(genericClients), nil
}

// NewScheduledRecording1ClientsByURL is the legacy version of NewScheduledRecording1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewScheduledRecording1ClientsByURL(loc *url.URL) ([]*ScheduledRecording1, error) {
	return NewScheduledRecording1ClientsByURLCtx(context.Background(), loc)
}

// NewScheduledRecording1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ScheduledRecording1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ScheduledRecording_1)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording1ClientsFromGenericClients(genericClients), nil
}

func newScheduledRecording1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ScheduledRecording1 {
	clients := make([]*ScheduledRecording1, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording1{genericClients[i]}
	}
	return clients
}

func NewScheduledRecording1MultiClients(mutiClients goupnp.MultiServiceClient) []*ScheduledRecording1 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ScheduledRecording_1))
	clients := make([]*ScheduledRecording1, len(services))
	for i := range services {
		clients[i] = &ScheduledRecording1{services[i]}
	}
	return clients
}

func (client *ScheduledRecording1) BrowseRecordSchedulesCtx(
	ctx context.Context,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}
	// BEGIN Marshal arguments into request.

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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "BrowseRecordSchedules", request, response); err != nil {
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

// BrowseRecordSchedules is the legacy version of BrowseRecordSchedulesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) BrowseRecordSchedules(Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseRecordSchedulesCtx(context.Background(),
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ScheduledRecording1) BrowseRecordTasksCtx(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
		StartingIndex    string
		RequestedCount   string
		SortCriteria     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "BrowseRecordTasks", request, response); err != nil {
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

// BrowseRecordTasks is the legacy version of BrowseRecordTasksCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) BrowseRecordTasks(RecordScheduleID string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseRecordTasksCtx(context.Background(),
		RecordScheduleID,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ScheduledRecording1) CreateRecordScheduleCtx(
	ctx context.Context,
	Elements string,
) (RecordScheduleID string, Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		Elements string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Elements, err = soap.MarshalString(Elements); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordScheduleID string
		Result           string
		UpdateID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "CreateRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordScheduleID, err = soap.UnmarshalString(response.RecordScheduleID); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// CreateRecordSchedule is the legacy version of CreateRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) CreateRecordSchedule(Elements string) (RecordScheduleID string, Result string, UpdateID uint32, err error) {
	return client.CreateRecordScheduleCtx(context.Background(),
		Elements,
	)
}

func (client *ScheduledRecording1) DeleteRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "DeleteRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteRecordSchedule is the legacy version of DeleteRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) DeleteRecordSchedule(RecordScheduleID string) (err error) {
	return client.DeleteRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording1) DeleteRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "DeleteRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteRecordTask is the legacy version of DeleteRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) DeleteRecordTask(RecordTaskID string) (err error) {
	return client.DeleteRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording1) DisableRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "DisableRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DisableRecordSchedule is the legacy version of DisableRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) DisableRecordSchedule(RecordScheduleID string) (err error) {
	return client.DisableRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording1) DisableRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "DisableRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DisableRecordTask is the legacy version of DisableRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) DisableRecordTask(RecordTaskID string) (err error) {
	return client.DisableRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording1) EnableRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "EnableRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// EnableRecordSchedule is the legacy version of EnableRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) EnableRecordSchedule(RecordScheduleID string) (err error) {
	return client.EnableRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording1) EnableRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "EnableRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// EnableRecordTask is the legacy version of EnableRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) EnableRecordTask(RecordTaskID string) (err error) {
	return client.EnableRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

//
// Arguments:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts

func (client *ScheduledRecording1) GetAllowedValuesCtx(
	ctx context.Context,
	DataTypeID string,
	Filter string,
) (PropertyInfo string, err error) {
	// Request structure.
	request := &struct {
		DataTypeID string
		Filter     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DataTypeID, err = soap.MarshalString(DataTypeID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PropertyInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetAllowedValues", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PropertyInfo, err = soap.UnmarshalString(response.PropertyInfo); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllowedValues is the legacy version of GetAllowedValuesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetAllowedValues(DataTypeID string, Filter string) (PropertyInfo string, err error) {
	return client.GetAllowedValuesCtx(context.Background(),
		DataTypeID,
		Filter,
	)
}

//
// Arguments:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts

func (client *ScheduledRecording1) GetPropertyListCtx(
	ctx context.Context,
	DataTypeID string,
) (PropertyList string, err error) {
	// Request structure.
	request := &struct {
		DataTypeID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DataTypeID, err = soap.MarshalString(DataTypeID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PropertyList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetPropertyList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PropertyList, err = soap.UnmarshalString(response.PropertyList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPropertyList is the legacy version of GetPropertyListCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetPropertyList(DataTypeID string) (PropertyList string, err error) {
	return client.GetPropertyListCtx(context.Background(),
		DataTypeID,
	)
}

func (client *ScheduledRecording1) GetRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
) (Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordSchedule is the legacy version of GetRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetRecordSchedule(RecordScheduleID string, Filter string) (Result string, UpdateID uint32, err error) {
	return client.GetRecordScheduleCtx(context.Background(),
		RecordScheduleID,
		Filter,
	)
}

func (client *ScheduledRecording1) GetRecordScheduleConflictsCtx(
	ctx context.Context,
	RecordScheduleID string,
) (RecordScheduleConflictIDList string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordScheduleConflictIDList string
		UpdateID                     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetRecordScheduleConflicts", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordScheduleConflictIDList, err = soap.UnmarshalString(response.RecordScheduleConflictIDList); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordScheduleConflicts is the legacy version of GetRecordScheduleConflictsCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetRecordScheduleConflicts(RecordScheduleID string) (RecordScheduleConflictIDList string, UpdateID uint32, err error) {
	return client.GetRecordScheduleConflictsCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording1) GetRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
	Filter string,
) (Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
		Filter       string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordTask is the legacy version of GetRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetRecordTask(RecordTaskID string, Filter string) (Result string, UpdateID uint32, err error) {
	return client.GetRecordTaskCtx(context.Background(),
		RecordTaskID,
		Filter,
	)
}

func (client *ScheduledRecording1) GetRecordTaskConflictsCtx(
	ctx context.Context,
	RecordTaskID string,
) (RecordTaskConflictIDList string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordTaskConflictIDList string
		UpdateID                 string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetRecordTaskConflicts", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordTaskConflictIDList, err = soap.UnmarshalString(response.RecordTaskConflictIDList); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordTaskConflicts is the legacy version of GetRecordTaskConflictsCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetRecordTaskConflicts(RecordTaskID string) (RecordTaskConflictIDList string, UpdateID uint32, err error) {
	return client.GetRecordTaskConflictsCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording1) GetSortCapabilitiesCtx(
	ctx context.Context,
) (SortCaps string, SortLevelCap uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SortCaps     string
		SortLevelCap string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetSortCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SortCaps, err = soap.UnmarshalString(response.SortCaps); err != nil {
		return
	}
	if SortLevelCap, err = soap.UnmarshalUi4(response.SortLevelCap); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSortCapabilities is the legacy version of GetSortCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetSortCapabilities() (SortCaps string, SortLevelCap uint32, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ScheduledRecording1) GetStateUpdateIDCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "GetStateUpdateID", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Id, err = soap.UnmarshalUi4(response.Id); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetStateUpdateID is the legacy version of GetStateUpdateIDCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) GetStateUpdateID() (Id uint32, err error) {
	return client.GetStateUpdateIDCtx(context.Background())
}

func (client *ScheduledRecording1) ResetRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_1, "ResetRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ResetRecordTask is the legacy version of ResetRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording1) ResetRecordTask(RecordTaskID string) (err error) {
	return client.ResetRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

// ScheduledRecording2 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ScheduledRecording:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ScheduledRecording2 struct {
	goupnp.ServiceClient
}

// NewScheduledRecording2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewScheduledRecording2ClientsCtx(ctx context.Context) (clients []*ScheduledRecording2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ScheduledRecording_2); err != nil {
		return
	}
	clients = newScheduledRecording2ClientsFromGenericClients(genericClients)
	return
}

// NewScheduledRecording2Clients is the legacy version of NewScheduledRecording2ClientsCtx, but uses
// context.Background() as the context.
func NewScheduledRecording2Clients() (clients []*ScheduledRecording2, errors []error, err error) {
	return NewScheduledRecording2ClientsCtx(context.Background())
}

// NewScheduledRecording2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewScheduledRecording2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ScheduledRecording2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ScheduledRecording_2)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording2ClientsFromGenericClients(genericClients), nil
}

// NewScheduledRecording2ClientsByURL is the legacy version of NewScheduledRecording2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewScheduledRecording2ClientsByURL(loc *url.URL) ([]*ScheduledRecording2, error) {
	return NewScheduledRecording2ClientsByURLCtx(context.Background(), loc)
}

// NewScheduledRecording2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewScheduledRecording2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ScheduledRecording2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ScheduledRecording_2)
	if err != nil {
		return nil, err
	}
	return newScheduledRecording2ClientsFromGenericClients(genericClients), nil
}

func newScheduledRecording2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ScheduledRecording2 {
	clients := make([]*ScheduledRecording2, len(genericClients))
	for i := range genericClients {
		clients[i] = &ScheduledRecording2{genericClients[i]}
	}
	return clients
}

func NewScheduledRecording2MultiClients(mutiClients goupnp.MultiServiceClient) []*ScheduledRecording2 {
	services := mutiClients.Filter(goupnp.ServiceType(URN_ScheduledRecording_2))
	clients := make([]*ScheduledRecording2, len(services))
	for i := range services {
		clients[i] = &ScheduledRecording2{services[i]}
	}
	return clients
}

func (client *ScheduledRecording2) BrowseRecordSchedulesCtx(
	ctx context.Context,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		Filter         string
		StartingIndex  string
		RequestedCount string
		SortCriteria   string
	}{}
	// BEGIN Marshal arguments into request.

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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "BrowseRecordSchedules", request, response); err != nil {
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

// BrowseRecordSchedules is the legacy version of BrowseRecordSchedulesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) BrowseRecordSchedules(Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseRecordSchedulesCtx(context.Background(),
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ScheduledRecording2) BrowseRecordTasksCtx(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
	StartingIndex uint32,
	RequestedCount uint32,
	SortCriteria string,
) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
		StartingIndex    string
		RequestedCount   string
		SortCriteria     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "BrowseRecordTasks", request, response); err != nil {
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

// BrowseRecordTasks is the legacy version of BrowseRecordTasksCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) BrowseRecordTasks(RecordScheduleID string, Filter string, StartingIndex uint32, RequestedCount uint32, SortCriteria string) (Result string, NumberReturned uint32, TotalMatches uint32, UpdateID uint32, err error) {
	return client.BrowseRecordTasksCtx(context.Background(),
		RecordScheduleID,
		Filter,
		StartingIndex,
		RequestedCount,
		SortCriteria,
	)
}

func (client *ScheduledRecording2) CreateRecordScheduleCtx(
	ctx context.Context,
	Elements string,
) (RecordScheduleID string, Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		Elements string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Elements, err = soap.MarshalString(Elements); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordScheduleID string
		Result           string
		UpdateID         string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "CreateRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordScheduleID, err = soap.UnmarshalString(response.RecordScheduleID); err != nil {
		return
	}
	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// CreateRecordSchedule is the legacy version of CreateRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) CreateRecordSchedule(Elements string) (RecordScheduleID string, Result string, UpdateID uint32, err error) {
	return client.CreateRecordScheduleCtx(context.Background(),
		Elements,
	)
}

func (client *ScheduledRecording2) DeleteRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "DeleteRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteRecordSchedule is the legacy version of DeleteRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) DeleteRecordSchedule(RecordScheduleID string) (err error) {
	return client.DeleteRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording2) DeleteRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "DeleteRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteRecordTask is the legacy version of DeleteRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) DeleteRecordTask(RecordTaskID string) (err error) {
	return client.DeleteRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording2) DisableRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "DisableRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DisableRecordSchedule is the legacy version of DisableRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) DisableRecordSchedule(RecordScheduleID string) (err error) {
	return client.DisableRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording2) DisableRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "DisableRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DisableRecordTask is the legacy version of DisableRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) DisableRecordTask(RecordTaskID string) (err error) {
	return client.DisableRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording2) EnableRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "EnableRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// EnableRecordSchedule is the legacy version of EnableRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) EnableRecordSchedule(RecordScheduleID string) (err error) {
	return client.EnableRecordScheduleCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording2) EnableRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "EnableRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// EnableRecordTask is the legacy version of EnableRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) EnableRecordTask(RecordTaskID string) (err error) {
	return client.EnableRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}

//
// Arguments:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts

func (client *ScheduledRecording2) GetAllowedValuesCtx(
	ctx context.Context,
	DataTypeID string,
	Filter string,
) (PropertyInfo string, err error) {
	// Request structure.
	request := &struct {
		DataTypeID string
		Filter     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DataTypeID, err = soap.MarshalString(DataTypeID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PropertyInfo string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetAllowedValues", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PropertyInfo, err = soap.UnmarshalString(response.PropertyInfo); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetAllowedValues is the legacy version of GetAllowedValuesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetAllowedValues(DataTypeID string, Filter string) (PropertyInfo string, err error) {
	return client.GetAllowedValuesCtx(context.Background(),
		DataTypeID,
		Filter,
	)
}

//
// Arguments:
//
// * DataTypeID: allowed values: A_ARG_TYPE_RecordSchedule, A_ARG_TYPE_RecordTask, A_ARG_TYPE_RecordScheduleParts

func (client *ScheduledRecording2) GetPropertyListCtx(
	ctx context.Context,
	DataTypeID string,
) (PropertyList string, err error) {
	// Request structure.
	request := &struct {
		DataTypeID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DataTypeID, err = soap.MarshalString(DataTypeID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PropertyList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetPropertyList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PropertyList, err = soap.UnmarshalString(response.PropertyList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPropertyList is the legacy version of GetPropertyListCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetPropertyList(DataTypeID string) (PropertyList string, err error) {
	return client.GetPropertyListCtx(context.Background(),
		DataTypeID,
	)
}

func (client *ScheduledRecording2) GetRecordScheduleCtx(
	ctx context.Context,
	RecordScheduleID string,
	Filter string,
) (Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
		Filter           string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetRecordSchedule", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordSchedule is the legacy version of GetRecordScheduleCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetRecordSchedule(RecordScheduleID string, Filter string) (Result string, UpdateID uint32, err error) {
	return client.GetRecordScheduleCtx(context.Background(),
		RecordScheduleID,
		Filter,
	)
}

func (client *ScheduledRecording2) GetRecordScheduleConflictsCtx(
	ctx context.Context,
	RecordScheduleID string,
) (RecordScheduleConflictIDList string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordScheduleID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordScheduleID, err = soap.MarshalString(RecordScheduleID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordScheduleConflictIDList string
		UpdateID                     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetRecordScheduleConflicts", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordScheduleConflictIDList, err = soap.UnmarshalString(response.RecordScheduleConflictIDList); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordScheduleConflicts is the legacy version of GetRecordScheduleConflictsCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetRecordScheduleConflicts(RecordScheduleID string) (RecordScheduleConflictIDList string, UpdateID uint32, err error) {
	return client.GetRecordScheduleConflictsCtx(context.Background(),
		RecordScheduleID,
	)
}

func (client *ScheduledRecording2) GetRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
	Filter string,
) (Result string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
		Filter       string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	if request.Filter, err = soap.MarshalString(Filter); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Result   string
		UpdateID string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Result, err = soap.UnmarshalString(response.Result); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordTask is the legacy version of GetRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetRecordTask(RecordTaskID string, Filter string) (Result string, UpdateID uint32, err error) {
	return client.GetRecordTaskCtx(context.Background(),
		RecordTaskID,
		Filter,
	)
}

func (client *ScheduledRecording2) GetRecordTaskConflictsCtx(
	ctx context.Context,
	RecordTaskID string,
) (RecordTaskConflictIDList string, UpdateID uint32, err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		RecordTaskConflictIDList string
		UpdateID                 string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetRecordTaskConflicts", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if RecordTaskConflictIDList, err = soap.UnmarshalString(response.RecordTaskConflictIDList); err != nil {
		return
	}
	if UpdateID, err = soap.UnmarshalUi4(response.UpdateID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetRecordTaskConflicts is the legacy version of GetRecordTaskConflictsCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetRecordTaskConflicts(RecordTaskID string) (RecordTaskConflictIDList string, UpdateID uint32, err error) {
	return client.GetRecordTaskConflictsCtx(context.Background(),
		RecordTaskID,
	)
}

func (client *ScheduledRecording2) GetSortCapabilitiesCtx(
	ctx context.Context,
) (SortCaps string, SortLevelCap uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SortCaps     string
		SortLevelCap string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetSortCapabilities", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SortCaps, err = soap.UnmarshalString(response.SortCaps); err != nil {
		return
	}
	if SortLevelCap, err = soap.UnmarshalUi4(response.SortLevelCap); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSortCapabilities is the legacy version of GetSortCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetSortCapabilities() (SortCaps string, SortLevelCap uint32, err error) {
	return client.GetSortCapabilitiesCtx(context.Background())
}

func (client *ScheduledRecording2) GetStateUpdateIDCtx(
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
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "GetStateUpdateID", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Id, err = soap.UnmarshalUi4(response.Id); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetStateUpdateID is the legacy version of GetStateUpdateIDCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) GetStateUpdateID() (Id uint32, err error) {
	return client.GetStateUpdateIDCtx(context.Background())
}

func (client *ScheduledRecording2) ResetRecordTaskCtx(
	ctx context.Context,
	RecordTaskID string,
) (err error) {
	// Request structure.
	request := &struct {
		RecordTaskID string
	}{}
	// BEGIN Marshal arguments into request.

	if request.RecordTaskID, err = soap.MarshalString(RecordTaskID); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_ScheduledRecording_2, "ResetRecordTask", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ResetRecordTask is the legacy version of ResetRecordTaskCtx, but uses
// context.Background() as the context.
func (client *ScheduledRecording2) ResetRecordTask(RecordTaskID string) (err error) {
	return client.ResetRecordTaskCtx(context.Background(),
		RecordTaskID,
	)
}
