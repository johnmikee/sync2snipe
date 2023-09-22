// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicecommands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mobiledevicecommands API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mobiledevicecommands API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMobileDeviceCommand(params *CreateMobileDeviceCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceCommandCreated, error)

	CreateMobileDeviceCommandURL(params *CreateMobileDeviceCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceCommandURLCreated, error)

	CreateMobileDeviceLockCommandURL(params *CreateMobileDeviceLockCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceLockCommandURLCreated, error)

	CreateMobileDeviceNameCommandURL(params *CreateMobileDeviceNameCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceNameCommandURLCreated, error)

	CreateMobileScheduleOSUpdateCommandURL(params *CreateMobileScheduleOSUpdateCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileScheduleOSUpdateCommandURLCreated, error)

	CreateMobileScheduleOSUpdateCommandWithProductVersionURL(params *CreateMobileScheduleOSUpdateCommandWithProductVersionURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileScheduleOSUpdateCommandWithProductVersionURLCreated, error)

	FindMobileDeviceCommands(params *FindMobileDeviceCommandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsOK, error)

	FindMobileDeviceCommandsByCommand(params *FindMobileDeviceCommandsByCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByCommandOK, error)

	FindMobileDeviceCommandsByName(params *FindMobileDeviceCommandsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByNameOK, error)

	FindMobileDeviceCommandsByUUID(params *FindMobileDeviceCommandsByUUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByUUIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateMobileDeviceCommand creates a new mobile device command

	The chart below includes additional requirements for usage of specific commands

| command | Parameters | Requirements | Values |
| ---------- | ------- | ------------ | ------ |
| DeviceName | device_name | Supervised Device | string |
| DeviceLocation | N/A | Supervised and in lost mode | N/A |
| DeviceLock | lock_message | optional | string |
| DisableLostMode | N/A | Supervised and in lost mode | N/A |
| EnableLostMode | lost_mode_message | Supervised device (required if lost_mode_phone is not set) | string |
| EnableLostMode | lost_mode_phone | Supervised device (required if lost_mode_message is not set) | string |
| EnableLostMode | lost_mode_footnote | optional | string |
| EnableLostMode | always_enforce_lost_mode | optional (defaults to 'true') | boolean |
| EnableLostMode | lost_mode_with_sound | optional (defaults to 'false') | boolean |
| EraseDevice | preserve_data_plan | optional (defaults to 'false') | boolean |
| EraseDevice | disallow_proximity_setup | optional (defaults to 'false') | boolean |
| EraseDevice | clear_activation_lock | optional (defaults to 'false') | boolean |
| PasscodeLockGracePeriod | passcode_lock_grace_period | Shared iPad | seconds (integer) |
| PlayLostModeSound | N/A | Supervised and in lost mode | N/A |
| RestartDevice | N/A | Supervised device | N/A |
| ScheduleOSUpdate (deprecated on 2022-10-17) | install_action | iOS 9 - 10.2, enrolled via a Prestage enrollment; and/or iOS 10.3 or later; tvOS 12 or later | 1 = Download the update for users to install, 2 = Download and install the update, and restart devices after installation |
| ScheduleOSUpdate (deprecated on 2022-10-17) | product_version | iOS 11.3 or later, tvOS 12.2 or later | string |
| SettingsDisableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
| SettingsEnableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
| ShutDownDevice | N/A | Supervised device | N/A |
| Wallpaper | wallpaper_settings | Supervised device | 1 = Lock screen, 2 = Home screen, 3 = both |
| Wallpaper | wallpaper_id | Supervised device (required if wallpaper_content is not set) | Jamf Pro icon ID (integer) |
| Wallpaper | wallpaper_content | Supervised device (required if wallpaper_id is not set) | base64 encoded PNG or JPEG |
| RefreshCellularPlans | e_sim_server_url | N/A | This URL is obtained from each carrier separately |
*/
func (a *Client) CreateMobileDeviceCommand(params *CreateMobileDeviceCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceCommandCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileDeviceCommandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileDeviceCommand",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileDeviceCommandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileDeviceCommandCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileDeviceCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateMobileDeviceCommandURL creates a new mobile device command

	The chart below includes additional requirements for usage of specific commands

| command | Parameters | Requirements | Values |
| ---------- | ------- | ------------ | ------ |
| DeviceLocation | N/A | Supervised and in lost mode | N/A |
| DisableLostMode | N/A | Supervised and in lost mode | N/A |
| EnableLostMode | lost_mode_message | Supervised device (required if lost_mode_phone is not set) | string |
| EnableLostMode | lost_mode_phone | Supervised device (required if lost_mode_message is not set) | string |
| EnableLostMode | lost_mode_footnote | optional | string |
| EnableLostMode | always_enforce_lost_mode | optional (defaults to 'true') | boolean |
| EnableLostMode | lost_mode_with_sound | optional (defaults to 'false') | boolean |
| EraseDevice | preserve_data_plan | optional (defaults to 'false') | boolean |
| EraseDevice | disallow_proximity_setup | optional (defaults to 'false') | boolean |
| PasscodeLockGracePeriod | passcode_lock_grace_period | Shared iPad | seconds (integer) |
| PlayLostModeSound | N/A | Supervised and in lost mode | N/A |
| RestartDevice | N/A | Supervised device | N/A |
| SettingsDisableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
| SettingsEnableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
| ShutDownDevice | N/A | Supervised device | N/A |
*/
func (a *Client) CreateMobileDeviceCommandURL(params *CreateMobileDeviceCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceCommandURLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileDeviceCommandURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileDeviceCommandURL",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command/{command}/id/{id_list}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileDeviceCommandURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileDeviceCommandURLCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileDeviceCommandURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMobileDeviceLockCommandURL sends a new lock command to a list of mobile devices
*/
func (a *Client) CreateMobileDeviceLockCommandURL(params *CreateMobileDeviceLockCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceLockCommandURLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileDeviceLockCommandURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileDeviceLockCommandURL",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command/DeviceLock/{lock_message}/id/{id_list}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileDeviceLockCommandURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileDeviceLockCommandURLCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileDeviceLockCommandURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMobileDeviceNameCommandURL creates a new command to set the name of a mobile device
*/
func (a *Client) CreateMobileDeviceNameCommandURL(params *CreateMobileDeviceNameCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileDeviceNameCommandURLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileDeviceNameCommandURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileDeviceNameCommandURL",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command/DeviceName/{device_name}/id/{id_list}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileDeviceNameCommandURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileDeviceNameCommandURLCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileDeviceNameCommandURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMobileScheduleOSUpdateCommandURL creates a new command to request that a mobile device update its o s command and mobile device list specified in URL device will be updated to the latest o s version based on device eligibility deprecated on 2022 10 17
*/
func (a *Client) CreateMobileScheduleOSUpdateCommandURL(params *CreateMobileScheduleOSUpdateCommandURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileScheduleOSUpdateCommandURLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileScheduleOSUpdateCommandURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileScheduleOSUpdateCommandURL",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command/ScheduleOSUpdate/{install_action}/id/{id_list}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileScheduleOSUpdateCommandURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileScheduleOSUpdateCommandURLCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileScheduleOSUpdateCommandURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMobileScheduleOSUpdateCommandWithProductVersionURL creates a new command to request that a mobile device update its o s command and mobile device list specified in URL mixing i o s and tv o s devices in ID list is not advised as product version is specific to o s type deprecated on 2022 10 17
*/
func (a *Client) CreateMobileScheduleOSUpdateCommandWithProductVersionURL(params *CreateMobileScheduleOSUpdateCommandWithProductVersionURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMobileScheduleOSUpdateCommandWithProductVersionURLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMobileScheduleOSUpdateCommandWithProductVersionURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMobileScheduleOSUpdateCommandWithProductVersionURL",
		Method:             "POST",
		PathPattern:        "/mobiledevicecommands/command/ScheduleOSUpdate/{install_action}/{product_version}/id/{id_list}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMobileScheduleOSUpdateCommandWithProductVersionURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMobileScheduleOSUpdateCommandWithProductVersionURLCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMobileScheduleOSUpdateCommandWithProductVersionURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMobileDeviceCommands finds all mobile device commands
*/
func (a *Client) FindMobileDeviceCommands(params *FindMobileDeviceCommandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMobileDeviceCommandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMobileDeviceCommands",
		Method:             "GET",
		PathPattern:        "/mobiledevicecommands",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMobileDeviceCommandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMobileDeviceCommandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMobileDeviceCommands: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMobileDeviceCommandsByCommand finds all mobile device commands for specified command
*/
func (a *Client) FindMobileDeviceCommandsByCommand(params *FindMobileDeviceCommandsByCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByCommandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMobileDeviceCommandsByCommandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMobileDeviceCommandsByCommand",
		Method:             "GET",
		PathPattern:        "/mobiledevicecommands/command/{command}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMobileDeviceCommandsByCommandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMobileDeviceCommandsByCommandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMobileDeviceCommandsByCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMobileDeviceCommandsByName finds all mobile device commands by command name
*/
func (a *Client) FindMobileDeviceCommandsByName(params *FindMobileDeviceCommandsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMobileDeviceCommandsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMobileDeviceCommandsByName",
		Method:             "GET",
		PathPattern:        "/mobiledevicecommands/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMobileDeviceCommandsByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMobileDeviceCommandsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMobileDeviceCommandsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMobileDeviceCommandsByUUID finds a mobile device command by UUID
*/
func (a *Client) FindMobileDeviceCommandsByUUID(params *FindMobileDeviceCommandsByUUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceCommandsByUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMobileDeviceCommandsByUUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMobileDeviceCommandsByUuid",
		Method:             "GET",
		PathPattern:        "/mobiledevicecommands/uuid/{uuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMobileDeviceCommandsByUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindMobileDeviceCommandsByUUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMobileDeviceCommandsByUuid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
