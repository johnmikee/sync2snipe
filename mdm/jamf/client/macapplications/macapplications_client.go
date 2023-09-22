// Code generated by go-swagger; DO NOT EDIT.

package macapplications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new macapplications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for macapplications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMacappByID(params *CreateMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMacappByIDCreated, error)

	DeleteMacappByID(params *DeleteMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMacappByIDOK, error)

	DeleteMacappByName(params *DeleteMacappByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMacappByNameOK, error)

	FindMacapps(params *FindMacappsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsOK, error)

	FindMacappsByID(params *FindMacappsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByIDOK, error)

	FindMacappsByIDSubset(params *FindMacappsByIDSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByIDSubsetOK, error)

	FindMacappsByName(params *FindMacappsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByNameOK, error)

	FindMacappsByNameSubset(params *FindMacappsByNameSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByNameSubsetOK, error)

	UpdateMacappByID(params *UpdateMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateMacappByIDCreated, error)

	UpdateMacappByName(params *UpdateMacappByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateMacappByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateMacappByID creates a new mac application by ID

Does not support vpp codes.
*/
func (a *Client) CreateMacappByID(params *CreateMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMacappByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMacappByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMacappById",
		Method:             "POST",
		PathPattern:        "/macapplications/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMacappByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateMacappByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMacappById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMacappByID deletes a mac application by ID
*/
func (a *Client) DeleteMacappByID(params *DeleteMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMacappByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMacappByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMacappById",
		Method:             "DELETE",
		PathPattern:        "/macapplications/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMacappByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteMacappByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMacappById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMacappByName deletes a mac application by name
*/
func (a *Client) DeleteMacappByName(params *DeleteMacappByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMacappByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMacappByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMacappByName",
		Method:             "DELETE",
		PathPattern:        "/macapplications/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMacappByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteMacappByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMacappByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMacapps finds all mac applications
*/
func (a *Client) FindMacapps(params *FindMacappsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMacappsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMacapps",
		Method:             "GET",
		PathPattern:        "/macapplications",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMacappsReader{formats: a.formats},
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
	success, ok := result.(*FindMacappsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMacapps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMacappsByID finds mac applications by ID
*/
func (a *Client) FindMacappsByID(params *FindMacappsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMacappsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMacappsById",
		Method:             "GET",
		PathPattern:        "/macapplications/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMacappsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindMacappsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMacappsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMacappsByIDSubset finds a subset of date for a mac application by ID

Subset values can also be appended using an ampersand to return multiple subsets (e.g. /subsets/General&Scope)
*/
func (a *Client) FindMacappsByIDSubset(params *FindMacappsByIDSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByIDSubsetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMacappsByIDSubsetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMacappsByIdSubset",
		Method:             "GET",
		PathPattern:        "/macapplications/id/{id}/subset/{subset}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMacappsByIDSubsetReader{formats: a.formats},
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
	success, ok := result.(*FindMacappsByIDSubsetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMacappsByIdSubset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMacappsByName finds mac applications by name
*/
func (a *Client) FindMacappsByName(params *FindMacappsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMacappsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMacappsByName",
		Method:             "GET",
		PathPattern:        "/macapplications/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMacappsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindMacappsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMacappsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindMacappsByNameSubset finds a subset of data for mac applications by name

Subset values can also be appended using an ampersand to return multiple subsets (e.g. /subsets/General&Scope)
*/
func (a *Client) FindMacappsByNameSubset(params *FindMacappsByNameSubsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindMacappsByNameSubsetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMacappsByNameSubsetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findMacappsByNameSubset",
		Method:             "GET",
		PathPattern:        "/macapplications/name/{name}/subset/{subset}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindMacappsByNameSubsetReader{formats: a.formats},
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
	success, ok := result.(*FindMacappsByNameSubsetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findMacappsByNameSubset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMacappByID updates an existing mac application by ID

Does not support vpp codes.
*/
func (a *Client) UpdateMacappByID(params *UpdateMacappByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateMacappByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMacappByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateMacappById",
		Method:             "PUT",
		PathPattern:        "/macapplications/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMacappByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateMacappByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMacappById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMacappByName updates an existing mac application by name

Does not support vpp codes.
*/
func (a *Client) UpdateMacappByName(params *UpdateMacappByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateMacappByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMacappByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateMacappByName",
		Method:             "PUT",
		PathPattern:        "/macapplications/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMacappByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateMacappByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMacappByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}