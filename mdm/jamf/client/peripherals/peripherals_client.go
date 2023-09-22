// Code generated by go-swagger; DO NOT EDIT.

package peripherals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new peripherals API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for peripherals API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePeripheralByID(params *CreatePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePeripheralByIDCreated, error)

	DeletePeripheralByID(params *DeletePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePeripheralByIDOK, error)

	FindPeripherals(params *FindPeripheralsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsOK, error)

	FindPeripheralsByID(params *FindPeripheralsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsByIDOK, error)

	FindPeripheralsByIDSubset(params *FindPeripheralsByIDSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsByIDSubsetOK, error)

	UpdatePeripheralByID(params *UpdatePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePeripheralByIDCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePeripheralByID creates a new peripheral by ID
*/
func (a *Client) CreatePeripheralByID(params *CreatePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePeripheralByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePeripheralByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPeripheralById",
		Method:             "POST",
		PathPattern:        "/peripherals/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePeripheralByIDReader{formats: a.formats},
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
	success, ok := result.(*CreatePeripheralByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPeripheralById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePeripheralByID deletes a peripheral by ID
*/
func (a *Client) DeletePeripheralByID(params *DeletePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePeripheralByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePeripheralByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePeripheralById",
		Method:             "DELETE",
		PathPattern:        "/peripherals/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePeripheralByIDReader{formats: a.formats},
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
	success, ok := result.(*DeletePeripheralByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePeripheralById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPeripherals finds all peripherals
*/
func (a *Client) FindPeripherals(params *FindPeripheralsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPeripheralsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPeripherals",
		Method:             "GET",
		PathPattern:        "/peripherals",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPeripheralsReader{formats: a.formats},
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
	success, ok := result.(*FindPeripheralsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPeripherals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPeripheralsByID finds peripherals by ID
*/
func (a *Client) FindPeripheralsByID(params *FindPeripheralsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPeripheralsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPeripheralsById",
		Method:             "GET",
		PathPattern:        "/peripherals/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPeripheralsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindPeripheralsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPeripheralsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPeripheralsByIDSubset finds a subset of data for a peripheral

Subset values can also be appended using an ampersand to return multiple subsets (e.g. /subsets/General&Location)
*/
func (a *Client) FindPeripheralsByIDSubset(params *FindPeripheralsByIDSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPeripheralsByIDSubsetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPeripheralsByIDSubsetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPeripheralsByIdSubset",
		Method:             "GET",
		PathPattern:        "/peripherals/id/{id}/subset/{subset}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPeripheralsByIDSubsetReader{formats: a.formats},
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
	success, ok := result.(*FindPeripheralsByIDSubsetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPeripheralsByIdSubset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePeripheralByID updates an existing peripheral by ID
*/
func (a *Client) UpdatePeripheralByID(params *UpdatePeripheralByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePeripheralByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePeripheralByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePeripheralById",
		Method:             "PUT",
		PathPattern:        "/peripherals/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePeripheralByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdatePeripheralByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePeripheralById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}