// Code generated by go-swagger; DO NOT EDIT.

package computergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new computergroups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for computergroups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateComputerGroupByID(params *CreateComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateComputerGroupByIDCreated, error)

	DeleteComputerGroupByID(params *DeleteComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputerGroupByIDOK, error)

	DeleteComputerGroupByName(params *DeleteComputerGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputerGroupByNameOK, error)

	FindComputerGroups(params *FindComputerGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsOK, error)

	FindComputerGroupsByID(params *FindComputerGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsByIDOK, error)

	FindComputerGroupsByName(params *FindComputerGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsByNameOK, error)

	UpdateComputerGroupByID(params *UpdateComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerGroupByIDCreated, error)

	UpdateComputerGroupByName(params *UpdateComputerGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerGroupByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateComputerGroupByID creates a new computer group by ID
*/
func (a *Client) CreateComputerGroupByID(params *CreateComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateComputerGroupByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateComputerGroupByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createComputerGroupById",
		Method:             "POST",
		PathPattern:        "/computergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateComputerGroupByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateComputerGroupByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createComputerGroupById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteComputerGroupByID deletes a computer group by ID
*/
func (a *Client) DeleteComputerGroupByID(params *DeleteComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputerGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputerGroupByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComputerGroupById",
		Method:             "DELETE",
		PathPattern:        "/computergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputerGroupByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteComputerGroupByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComputerGroupById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteComputerGroupByName deletes a computer group by name
*/
func (a *Client) DeleteComputerGroupByName(params *DeleteComputerGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputerGroupByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputerGroupByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComputerGroupByName",
		Method:             "DELETE",
		PathPattern:        "/computergroups/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputerGroupByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteComputerGroupByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComputerGroupByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindComputerGroups finds all computer groups
*/
func (a *Client) FindComputerGroups(params *FindComputerGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindComputerGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findComputerGroups",
		Method:             "GET",
		PathPattern:        "/computergroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindComputerGroupsReader{formats: a.formats},
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
	success, ok := result.(*FindComputerGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findComputerGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindComputerGroupsByID finds computer groups by ID
*/
func (a *Client) FindComputerGroupsByID(params *FindComputerGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindComputerGroupsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findComputerGroupsById",
		Method:             "GET",
		PathPattern:        "/computergroups/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindComputerGroupsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindComputerGroupsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findComputerGroupsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindComputerGroupsByName finds computer groups by name
*/
func (a *Client) FindComputerGroupsByName(params *FindComputerGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerGroupsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindComputerGroupsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findComputerGroupsByName",
		Method:             "GET",
		PathPattern:        "/computergroups/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindComputerGroupsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindComputerGroupsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findComputerGroupsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateComputerGroupByID updates an existing computer group by ID

One or more computers can be added by using "computer_additions" instead of "computers". One or more computers can be deleted by using "computer_deletions" instead of "computers"
*/
func (a *Client) UpdateComputerGroupByID(params *UpdateComputerGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerGroupByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateComputerGroupByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateComputerGroupById",
		Method:             "PUT",
		PathPattern:        "/computergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateComputerGroupByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateComputerGroupByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateComputerGroupById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateComputerGroupByName updates an existing computer group by name

One or more computers can be added by using "computer_additions" instead of "computers". One or more computers can be deleted by using "computer_deletions" instead of "computers"
*/
func (a *Client) UpdateComputerGroupByName(params *UpdateComputerGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerGroupByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateComputerGroupByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateComputerGroupByName",
		Method:             "PUT",
		PathPattern:        "/computergroups/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateComputerGroupByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateComputerGroupByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateComputerGroupByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}