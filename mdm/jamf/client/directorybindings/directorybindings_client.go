// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new directorybindings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for directorybindings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDirectoryBindingByID(params *CreateDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDirectoryBindingByIDCreated, error)

	DeleteDirectoryBindingByID(params *DeleteDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDirectoryBindingByIDOK, error)

	DeleteDirectoryBindingByName(params *DeleteDirectoryBindingByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDirectoryBindingByNameOK, error)

	FindDirectoryBindings(params *FindDirectoryBindingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsOK, error)

	FindDirectoryBindingsByID(params *FindDirectoryBindingsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsByIDOK, error)

	FindDirectoryBindingsByName(params *FindDirectoryBindingsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsByNameOK, error)

	UpdateDirectoryBindingByID(params *UpdateDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDirectoryBindingByIDCreated, error)

	UpdateDirectoryBindingByName(params *UpdateDirectoryBindingByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDirectoryBindingByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDirectoryBindingByID creates a new directory binding by ID
*/
func (a *Client) CreateDirectoryBindingByID(params *CreateDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDirectoryBindingByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDirectoryBindingByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDirectoryBindingById",
		Method:             "POST",
		PathPattern:        "/directorybindings/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDirectoryBindingByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateDirectoryBindingByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDirectoryBindingById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDirectoryBindingByID deletes a directory binding by ID
*/
func (a *Client) DeleteDirectoryBindingByID(params *DeleteDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDirectoryBindingByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDirectoryBindingByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDirectoryBindingById",
		Method:             "DELETE",
		PathPattern:        "/directorybindings/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDirectoryBindingByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteDirectoryBindingByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDirectoryBindingById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDirectoryBindingByName deletes a directory binding by name
*/
func (a *Client) DeleteDirectoryBindingByName(params *DeleteDirectoryBindingByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDirectoryBindingByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDirectoryBindingByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDirectoryBindingByName",
		Method:             "DELETE",
		PathPattern:        "/directorybindings/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDirectoryBindingByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteDirectoryBindingByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDirectoryBindingByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDirectoryBindings finds all directory bindings
*/
func (a *Client) FindDirectoryBindings(params *FindDirectoryBindingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDirectoryBindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDirectoryBindings",
		Method:             "GET",
		PathPattern:        "/directorybindings",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDirectoryBindingsReader{formats: a.formats},
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
	success, ok := result.(*FindDirectoryBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDirectoryBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDirectoryBindingsByID finds directory bindings by ID
*/
func (a *Client) FindDirectoryBindingsByID(params *FindDirectoryBindingsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDirectoryBindingsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDirectoryBindingsById",
		Method:             "GET",
		PathPattern:        "/directorybindings/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDirectoryBindingsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindDirectoryBindingsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDirectoryBindingsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDirectoryBindingsByName finds directory bindings by name
*/
func (a *Client) FindDirectoryBindingsByName(params *FindDirectoryBindingsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDirectoryBindingsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDirectoryBindingsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDirectoryBindingsByName",
		Method:             "GET",
		PathPattern:        "/directorybindings/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDirectoryBindingsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindDirectoryBindingsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDirectoryBindingsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDirectoryBindingByID updates an existing directory binding by ID
*/
func (a *Client) UpdateDirectoryBindingByID(params *UpdateDirectoryBindingByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDirectoryBindingByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDirectoryBindingByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDirectoryBindingById",
		Method:             "PUT",
		PathPattern:        "/directorybindings/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDirectoryBindingByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDirectoryBindingByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDirectoryBindingById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDirectoryBindingByName updates an existing directory binding by name
*/
func (a *Client) UpdateDirectoryBindingByName(params *UpdateDirectoryBindingByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDirectoryBindingByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDirectoryBindingByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDirectoryBindingByName",
		Method:             "PUT",
		PathPattern:        "/directorybindings/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDirectoryBindingByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateDirectoryBindingByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDirectoryBindingByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
