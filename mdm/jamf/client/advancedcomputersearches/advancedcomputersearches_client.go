// Code generated by go-swagger; DO NOT EDIT.

package advancedcomputersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new advancedcomputersearches API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for advancedcomputersearches API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAdvancedComputerSearchgByID(params *CreateAdvancedComputerSearchgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdvancedComputerSearchgByIDCreated, error)

	DeleteAdvancedComputerSearchByID(params *DeleteAdvancedComputerSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedComputerSearchByIDOK, error)

	DeleteAdvancedComputerSearchByName(params *DeleteAdvancedComputerSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedComputerSearchByNameOK, error)

	FindAdvancedComputerSearches(params *FindAdvancedComputerSearchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesOK, error)

	FindAdvancedComputerSearchesByID(params *FindAdvancedComputerSearchesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesByIDOK, error)

	FindAdvancedComputerSearchesByName(params *FindAdvancedComputerSearchesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesByNameOK, error)

	UpdateAdvancedComputerSearchByID(params *UpdateAdvancedComputerSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedComputerSearchByIDCreated, error)

	UpdateAdvancedComputerSearchByName(params *UpdateAdvancedComputerSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedComputerSearchByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAdvancedComputerSearchgByID creates a new advanced computer search
*/
func (a *Client) CreateAdvancedComputerSearchgByID(params *CreateAdvancedComputerSearchgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAdvancedComputerSearchgByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAdvancedComputerSearchgByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAdvancedComputerSearchgById",
		Method:             "POST",
		PathPattern:        "/advancedcomputersearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAdvancedComputerSearchgByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateAdvancedComputerSearchgByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAdvancedComputerSearchgById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAdvancedComputerSearchByID deletes a computer search by ID
*/
func (a *Client) DeleteAdvancedComputerSearchByID(params *DeleteAdvancedComputerSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedComputerSearchByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdvancedComputerSearchByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdvancedComputerSearchById",
		Method:             "DELETE",
		PathPattern:        "/advancedcomputersearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdvancedComputerSearchByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAdvancedComputerSearchByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAdvancedComputerSearchById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAdvancedComputerSearchByName deletes a computer search by name
*/
func (a *Client) DeleteAdvancedComputerSearchByName(params *DeleteAdvancedComputerSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAdvancedComputerSearchByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdvancedComputerSearchByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAdvancedComputerSearchByName",
		Method:             "DELETE",
		PathPattern:        "/advancedcomputersearches/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdvancedComputerSearchByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteAdvancedComputerSearchByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAdvancedComputerSearchByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAdvancedComputerSearches finds all advanced computer searches
*/
func (a *Client) FindAdvancedComputerSearches(params *FindAdvancedComputerSearchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAdvancedComputerSearchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAdvancedComputerSearches",
		Method:             "GET",
		PathPattern:        "/advancedcomputersearches",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindAdvancedComputerSearchesReader{formats: a.formats},
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
	success, ok := result.(*FindAdvancedComputerSearchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAdvancedComputerSearches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAdvancedComputerSearchesByID finds computer searches by ID

Additional display fields are returned within the `Computer` object
*/
func (a *Client) FindAdvancedComputerSearchesByID(params *FindAdvancedComputerSearchesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAdvancedComputerSearchesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAdvancedComputerSearchesById",
		Method:             "GET",
		PathPattern:        "/advancedcomputersearches/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindAdvancedComputerSearchesByIDReader{formats: a.formats},
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
	success, ok := result.(*FindAdvancedComputerSearchesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAdvancedComputerSearchesById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAdvancedComputerSearchesByName finds advanced computer searches by name

Additional display fields are returned within the `Computer` object
*/
func (a *Client) FindAdvancedComputerSearchesByName(params *FindAdvancedComputerSearchesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindAdvancedComputerSearchesByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAdvancedComputerSearchesByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAdvancedComputerSearchesByName",
		Method:             "GET",
		PathPattern:        "/advancedcomputersearches/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindAdvancedComputerSearchesByNameReader{formats: a.formats},
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
	success, ok := result.(*FindAdvancedComputerSearchesByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAdvancedComputerSearchesByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAdvancedComputerSearchByID updates an existing advanced computer search by ID
*/
func (a *Client) UpdateAdvancedComputerSearchByID(params *UpdateAdvancedComputerSearchByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedComputerSearchByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdvancedComputerSearchByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdvancedComputerSearchById",
		Method:             "PUT",
		PathPattern:        "/advancedcomputersearches/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdvancedComputerSearchByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdvancedComputerSearchByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdvancedComputerSearchById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAdvancedComputerSearchByName updates an existing advanced computer search by name
*/
func (a *Client) UpdateAdvancedComputerSearchByName(params *UpdateAdvancedComputerSearchByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAdvancedComputerSearchByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdvancedComputerSearchByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAdvancedComputerSearchByName",
		Method:             "PUT",
		PathPattern:        "/advancedcomputersearches/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdvancedComputerSearchByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateAdvancedComputerSearchByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdvancedComputerSearchByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
