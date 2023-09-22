// Code generated by go-swagger; DO NOT EDIT.

package usergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new usergroups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for usergroups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUserGroupsByID(params *CreateUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserGroupsByIDCreated, error)

	DeleteUserGroupsByID(params *DeleteUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserGroupsByIDOK, error)

	DeleteUserGroupsByName(params *DeleteUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserGroupsByNameOK, error)

	FindUserGroups(params *FindUserGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsOK, error)

	FindUserGroupsByID(params *FindUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsByIDOK, error)

	FindUserGroupsByName(params *FindUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsByNameOK, error)

	UpdateUserGroupsByID(params *UpdateUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserGroupsByIDCreated, error)

	UpdateUserGroupsByName(params *UpdateUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserGroupsByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateUserGroupsByID creates user groups by ID
*/
func (a *Client) CreateUserGroupsByID(params *CreateUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserGroupsByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserGroupsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUserGroupsById",
		Method:             "POST",
		PathPattern:        "/usergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserGroupsByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateUserGroupsByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUserGroupsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserGroupsByID deletes user groups by ID
*/
func (a *Client) DeleteUserGroupsByID(params *DeleteUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserGroupsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserGroupsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserGroupsById",
		Method:             "DELETE",
		PathPattern:        "/usergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserGroupsByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserGroupsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserGroupsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserGroupsByName deletes user groups by name
*/
func (a *Client) DeleteUserGroupsByName(params *DeleteUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserGroupsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserGroupsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserGroupsByName",
		Method:             "DELETE",
		PathPattern:        "/usergroups/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserGroupsByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserGroupsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserGroupsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUserGroups finds all user groups
*/
func (a *Client) FindUserGroups(params *FindUserGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUserGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUserGroups",
		Method:             "GET",
		PathPattern:        "/usergroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUserGroupsReader{formats: a.formats},
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
	success, ok := result.(*FindUserGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUserGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUserGroupsByID finds user groups by ID
*/
func (a *Client) FindUserGroupsByID(params *FindUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUserGroupsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUserGroupsById",
		Method:             "GET",
		PathPattern:        "/usergroups/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUserGroupsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindUserGroupsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUserGroupsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUserGroupsByName finds user groups by name
*/
func (a *Client) FindUserGroupsByName(params *FindUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserGroupsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUserGroupsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUserGroupsByName",
		Method:             "GET",
		PathPattern:        "/usergroups/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUserGroupsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindUserGroupsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUserGroupsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserGroupsByID updates user groups by ID

One or more users can be added by using "user_additions" instead of "users". One or more users can be deleted by using "user_deletions" instead of "users".
*/
func (a *Client) UpdateUserGroupsByID(params *UpdateUserGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserGroupsByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserGroupsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserGroupsById",
		Method:             "PUT",
		PathPattern:        "/usergroups/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserGroupsByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserGroupsByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserGroupsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserGroupsByName updates user groups by name

One or more users can be added by using "user_additions" instead of "users". One or more users can be deleted by using "user_deletions" instead of "users".
*/
func (a *Client) UpdateUserGroupsByName(params *UpdateUserGroupsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserGroupsByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserGroupsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserGroupsByName",
		Method:             "PUT",
		PathPattern:        "/usergroups/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserGroupsByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserGroupsByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserGroupsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
