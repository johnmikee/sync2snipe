// Code generated by go-swagger; DO NOT EDIT.

package removablemacaddresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new removablemacaddresses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for removablemacaddresses API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRemovableMacAddressByID(params *CreateRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRemovableMacAddressByIDCreated, error)

	DeleteRemovableMacAddressByID(params *DeleteRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemovableMacAddressByIDOK, error)

	DeleteRemovableMacAddressByName(params *DeleteRemovableMacAddressByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemovableMacAddressByNameOK, error)

	FindRemovableMacAddresses(params *FindRemovableMacAddressesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesOK, error)

	FindRemovableMacAddressesByID(params *FindRemovableMacAddressesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesByIDOK, error)

	FindRemovableMacAddressesByName(params *FindRemovableMacAddressesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesByNameOK, error)

	UpdateRemovableMacAddressByID(params *UpdateRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemovableMacAddressByIDCreated, error)

	UpdateRemovableMacAddressByName(params *UpdateRemovableMacAddressByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemovableMacAddressByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRemovableMacAddressByID creates a new removable mac address by ID
*/
func (a *Client) CreateRemovableMacAddressByID(params *CreateRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRemovableMacAddressByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRemovableMacAddressByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRemovableMacAddressById",
		Method:             "POST",
		PathPattern:        "/removablemacaddresses/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRemovableMacAddressByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateRemovableMacAddressByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRemovableMacAddressById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRemovableMacAddressByID deletes a removable mac address by ID
*/
func (a *Client) DeleteRemovableMacAddressByID(params *DeleteRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemovableMacAddressByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRemovableMacAddressByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRemovableMacAddressById",
		Method:             "DELETE",
		PathPattern:        "/removablemacaddresses/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRemovableMacAddressByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteRemovableMacAddressByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRemovableMacAddressById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRemovableMacAddressByName deletes a removable mac address by name
*/
func (a *Client) DeleteRemovableMacAddressByName(params *DeleteRemovableMacAddressByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemovableMacAddressByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRemovableMacAddressByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRemovableMacAddressByName",
		Method:             "DELETE",
		PathPattern:        "/removablemacaddresses/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRemovableMacAddressByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteRemovableMacAddressByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRemovableMacAddressByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindRemovableMacAddresses finds all removable mac addresses
*/
func (a *Client) FindRemovableMacAddresses(params *FindRemovableMacAddressesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindRemovableMacAddressesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findRemovableMacAddresses",
		Method:             "GET",
		PathPattern:        "/removablemacaddresses",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindRemovableMacAddressesReader{formats: a.formats},
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
	success, ok := result.(*FindRemovableMacAddressesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findRemovableMacAddresses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindRemovableMacAddressesByID finds removable mac addresses by ID
*/
func (a *Client) FindRemovableMacAddressesByID(params *FindRemovableMacAddressesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindRemovableMacAddressesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findRemovableMacAddressesById",
		Method:             "GET",
		PathPattern:        "/removablemacaddresses/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindRemovableMacAddressesByIDReader{formats: a.formats},
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
	success, ok := result.(*FindRemovableMacAddressesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findRemovableMacAddressesById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindRemovableMacAddressesByName finds removable mac addresses by name
*/
func (a *Client) FindRemovableMacAddressesByName(params *FindRemovableMacAddressesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRemovableMacAddressesByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindRemovableMacAddressesByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findRemovableMacAddressesByName",
		Method:             "GET",
		PathPattern:        "/removablemacaddresses/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindRemovableMacAddressesByNameReader{formats: a.formats},
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
	success, ok := result.(*FindRemovableMacAddressesByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findRemovableMacAddressesByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateRemovableMacAddressByID updates an existing removable mac address by ID
*/
func (a *Client) UpdateRemovableMacAddressByID(params *UpdateRemovableMacAddressByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemovableMacAddressByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRemovableMacAddressByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRemovableMacAddressById",
		Method:             "PUT",
		PathPattern:        "/removablemacaddresses/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRemovableMacAddressByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateRemovableMacAddressByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRemovableMacAddressById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateRemovableMacAddressByName updates an existing removable mac address by name
*/
func (a *Client) UpdateRemovableMacAddressByName(params *UpdateRemovableMacAddressByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemovableMacAddressByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRemovableMacAddressByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRemovableMacAddressByName",
		Method:             "PUT",
		PathPattern:        "/removablemacaddresses/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRemovableMacAddressByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateRemovableMacAddressByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRemovableMacAddressByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}