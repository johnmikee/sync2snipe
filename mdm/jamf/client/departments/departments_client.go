// Code generated by go-swagger; DO NOT EDIT.

package departments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new departments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for departments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDepartmentByID(params *CreateDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDepartmentByIDCreated, error)

	DeleteDepartmentByID(params *DeleteDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDepartmentByIDOK, error)

	DeleteDepartmentByName(params *DeleteDepartmentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDepartmentByNameOK, error)

	FindDepartments(params *FindDepartmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsOK, error)

	FindDepartmentsByID(params *FindDepartmentsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsByIDOK, error)

	FindDepartmentsByName(params *FindDepartmentsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsByNameOK, error)

	UpdateDepartmentByID(params *UpdateDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDepartmentByIDCreated, error)

	UpdateDepartmentByName(params *UpdateDepartmentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDepartmentByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDepartmentByID creates a new department by ID
*/
func (a *Client) CreateDepartmentByID(params *CreateDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDepartmentByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDepartmentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDepartmentById",
		Method:             "POST",
		PathPattern:        "/departments/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDepartmentByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateDepartmentByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDepartmentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDepartmentByID deletes a department by ID
*/
func (a *Client) DeleteDepartmentByID(params *DeleteDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDepartmentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDepartmentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDepartmentById",
		Method:             "DELETE",
		PathPattern:        "/departments/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDepartmentByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteDepartmentByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDepartmentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDepartmentByName deletes a department by name
*/
func (a *Client) DeleteDepartmentByName(params *DeleteDepartmentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDepartmentByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDepartmentByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDepartmentByName",
		Method:             "DELETE",
		PathPattern:        "/departments/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDepartmentByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteDepartmentByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDepartmentByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDepartments finds all departments
*/
func (a *Client) FindDepartments(params *FindDepartmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDepartmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDepartments",
		Method:             "GET",
		PathPattern:        "/departments",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDepartmentsReader{formats: a.formats},
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
	success, ok := result.(*FindDepartmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDepartments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDepartmentsByID finds departments by ID
*/
func (a *Client) FindDepartmentsByID(params *FindDepartmentsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDepartmentsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDepartmentsById",
		Method:             "GET",
		PathPattern:        "/departments/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDepartmentsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindDepartmentsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDepartmentsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDepartmentsByName finds departments by name
*/
func (a *Client) FindDepartmentsByName(params *FindDepartmentsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDepartmentsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDepartmentsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDepartmentsByName",
		Method:             "GET",
		PathPattern:        "/departments/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDepartmentsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindDepartmentsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDepartmentsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDepartmentByID updates an existing department by ID
*/
func (a *Client) UpdateDepartmentByID(params *UpdateDepartmentByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDepartmentByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDepartmentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDepartmentById",
		Method:             "PUT",
		PathPattern:        "/departments/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDepartmentByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDepartmentByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDepartmentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDepartmentByName updates an existing department by name
*/
func (a *Client) UpdateDepartmentByName(params *UpdateDepartmentByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDepartmentByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDepartmentByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDepartmentByName",
		Method:             "PUT",
		PathPattern:        "/departments/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDepartmentByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateDepartmentByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDepartmentByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}