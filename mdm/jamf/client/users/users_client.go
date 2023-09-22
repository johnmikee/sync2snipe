// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUserByID(params *CreateUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserByIDCreated, error)

	DeleteUserByEmailAddress(params *DeleteUserByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByEmailAddressOK, error)

	DeleteUserByID(params *DeleteUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByIDOK, error)

	DeleteUserByName(params *DeleteUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByNameOK, error)

	FindUsers(params *FindUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersOK, error)

	FindUsersByEmailAddress(params *FindUsersByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByEmailAddressOK, error)

	FindUsersByID(params *FindUsersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByIDOK, error)

	FindUsersByName(params *FindUsersByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByNameOK, error)

	UpdateUserByEmailAddress(params *UpdateUserByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByEmailAddressCreated, error)

	UpdateUserByID(params *UpdateUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByIDCreated, error)

	UpdateUserByName(params *UpdateUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateUserByID creates a new user by ID

This operation cannot be used to add computers, mobile devices, peripherals, vpp assignments to a user.
*/
func (a *Client) CreateUserByID(params *CreateUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUserById",
		Method:             "POST",
		PathPattern:        "/users/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateUserByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUserById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserByEmailAddress deletes a user by email address
*/
func (a *Client) DeleteUserByEmailAddress(params *DeleteUserByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByEmailAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserByEmailAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserByEmailAddress",
		Method:             "DELETE",
		PathPattern:        "/users/email/{email}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserByEmailAddressReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserByEmailAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserByEmailAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserByID deletes a user by ID
*/
func (a *Client) DeleteUserByID(params *DeleteUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserById",
		Method:             "DELETE",
		PathPattern:        "/users/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserByName deletes a user by name
*/
func (a *Client) DeleteUserByName(params *DeleteUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserByName",
		Method:             "DELETE",
		PathPattern:        "/users/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUsers finds all users
*/
func (a *Client) FindUsers(params *FindUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUsersReader{formats: a.formats},
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
	success, ok := result.(*FindUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUsersByEmailAddress finds users by email address

Because email addresses may not be unique, this operation may return a list of users.
*/
func (a *Client) FindUsersByEmailAddress(params *FindUsersByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByEmailAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUsersByEmailAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUsersByEmailAddress",
		Method:             "GET",
		PathPattern:        "/users/email/{email}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUsersByEmailAddressReader{formats: a.formats},
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
	success, ok := result.(*FindUsersByEmailAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUsersByEmailAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUsersByID finds users by ID
*/
func (a *Client) FindUsersByID(params *FindUsersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUsersByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUsersById",
		Method:             "GET",
		PathPattern:        "/users/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUsersByIDReader{formats: a.formats},
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
	success, ok := result.(*FindUsersByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUsersById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUsersByName finds users by name
*/
func (a *Client) FindUsersByName(params *FindUsersByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUsersByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUsersByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUsersByName",
		Method:             "GET",
		PathPattern:        "/users/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindUsersByNameReader{formats: a.formats},
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
	success, ok := result.(*FindUsersByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUsersByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserByEmailAddress updates an existing user by email address
*/
func (a *Client) UpdateUserByEmailAddress(params *UpdateUserByEmailAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByEmailAddressCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserByEmailAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserByEmailAddress",
		Method:             "PUT",
		PathPattern:        "/users/email/{email}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserByEmailAddressReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserByEmailAddressCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserByEmailAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserByID updates an existing user by ID

This operation cannot be used to add computers, mobile devices, peripherals, vpp assignments to a user.
*/
func (a *Client) UpdateUserByID(params *UpdateUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserById",
		Method:             "PUT",
		PathPattern:        "/users/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserByName updates an existing user by name

This operation cannot be used to add computers, mobile devices, peripherals, vpp assignments to a user.
*/
func (a *Client) UpdateUserByName(params *UpdateUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserByName",
		Method:             "PUT",
		PathPattern:        "/users/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}