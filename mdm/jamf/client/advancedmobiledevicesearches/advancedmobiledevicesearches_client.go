// Code generated by go-swagger; DO NOT EDIT.

package advancedmobiledevicesearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new advancedmobiledevicesearches API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for advancedmobiledevicesearches API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAdvancedMobileDeviceSearchByID(params *CreateAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdvancedMobileDeviceSearchByIDCreated, error)

	DeleteAdvancedMobileDeviceSearchByID(params *DeleteAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedMobileDeviceSearchByIDOK, error)

	DeleteAdvancedMobileDeviceSearchByName(params *DeleteAdvancedMobileDeviceSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedMobileDeviceSearchByNameOK, error)

	FindAdvancedMobileDeviceSearches(params *FindAdvancedMobileDeviceSearchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedMobileDeviceSearchesOK, error)

	FindAdvancedMobileDeviceSearchesByID(params *FindAdvancedMobileDeviceSearchesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedMobileDeviceSearchesByIDOK, error)

	FindMobileDeviceSearchesByName(params *FindMobileDeviceSearchesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceSearchesByNameOK, error)

	UpdateAdvancedMobileDeviceSearchByID(params *UpdateAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedMobileDeviceSearchByIDCreated, error)

	UpdateAdvancedMobileDeviceSearchByName(params *UpdateAdvancedMobileDeviceSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedMobileDeviceSearchByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAdvancedMobileDeviceSearchByID creates a new advanced mobile device search
*/
func (a *Client) CreateAdvancedMobileDeviceSearchByID(params *CreateAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdvancedMobileDeviceSearchByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAdvancedMobileDeviceSearchByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAdvancedMobileDeviceSearchById",
		Method:             "POST",
		PathPattern:        "/advancedmobiledevicesearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAdvancedMobileDeviceSearchByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateAdvancedMobileDeviceSearchByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAdvancedMobileDeviceSearchById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAdvancedMobileDeviceSearchByID deletes a mobile device search by ID
*/
func (a *Client) DeleteAdvancedMobileDeviceSearchByID(params *DeleteAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedMobileDeviceSearchByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdvancedMobileDeviceSearchByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdvancedMobileDeviceSearchById",
		Method:             "DELETE",
		PathPattern:        "/advancedmobiledevicesearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdvancedMobileDeviceSearchByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAdvancedMobileDeviceSearchByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAdvancedMobileDeviceSearchById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAdvancedMobileDeviceSearchByName deletes a mobile device search by name
*/
func (a *Client) DeleteAdvancedMobileDeviceSearchByName(params *DeleteAdvancedMobileDeviceSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedMobileDeviceSearchByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdvancedMobileDeviceSearchByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdvancedMobileDeviceSearchByName",
		Method:             "DELETE",
		PathPattern:        "/advancedmobiledevicesearches/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdvancedMobileDeviceSearchByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteAdvancedMobileDeviceSearchByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAdvancedMobileDeviceSearchByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAdvancedMobileDeviceSearches finds all advanced mobile device searches
*/
func (a *Client) FindAdvancedMobileDeviceSearches(params *FindAdvancedMobileDeviceSearchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedMobileDeviceSearchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAdvancedMobileDeviceSearchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAdvancedMobileDeviceSearches",
		Method:             "GET",
		PathPattern:        "/advancedmobiledevicesearches",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindAdvancedMobileDeviceSearchesReader{formats: a.formats},
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
	success, ok := result.(*FindAdvancedMobileDeviceSearchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAdvancedMobileDeviceSearches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAdvancedMobileDeviceSearchesByID finds mobile device searches by ID
*/
func (a *Client) FindAdvancedMobileDeviceSearchesByID(params *FindAdvancedMobileDeviceSearchesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedMobileDeviceSearchesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAdvancedMobileDeviceSearchesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAdvancedMobileDeviceSearchesById",
		Method:             "GET",
		PathPattern:        "/advancedmobiledevicesearches/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindAdvancedMobileDeviceSearchesByIDReader{formats: a.formats},
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
	success, ok := result.(*FindAdvancedMobileDeviceSearchesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAdvancedMobileDeviceSearchesById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMobileDeviceSearchesByName finds advanced mobile device searches by name
*/
func (a *Client) FindMobileDeviceSearchesByName(params *FindMobileDeviceSearchesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMobileDeviceSearchesByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMobileDeviceSearchesByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMobileDeviceSearchesByName",
		Method:             "GET",
		PathPattern:        "/advancedmobiledevicesearches/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMobileDeviceSearchesByNameReader{formats: a.formats},
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
	success, ok := result.(*FindMobileDeviceSearchesByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMobileDeviceSearchesByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAdvancedMobileDeviceSearchByID updates an existing advanced mobile device search by ID
*/
func (a *Client) UpdateAdvancedMobileDeviceSearchByID(params *UpdateAdvancedMobileDeviceSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedMobileDeviceSearchByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdvancedMobileDeviceSearchByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdvancedMobileDeviceSearchById",
		Method:             "PUT",
		PathPattern:        "/advancedmobiledevicesearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdvancedMobileDeviceSearchByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdvancedMobileDeviceSearchByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdvancedMobileDeviceSearchById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAdvancedMobileDeviceSearchByName updates an existing advanced mobile device search by name
*/
func (a *Client) UpdateAdvancedMobileDeviceSearchByName(params *UpdateAdvancedMobileDeviceSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedMobileDeviceSearchByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdvancedMobileDeviceSearchByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdvancedMobileDeviceSearchByName",
		Method:             "PUT",
		PathPattern:        "/advancedmobiledevicesearches/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdvancedMobileDeviceSearchByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdvancedMobileDeviceSearchByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdvancedMobileDeviceSearchByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
