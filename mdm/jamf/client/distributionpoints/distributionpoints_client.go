// Code generated by go-swagger; DO NOT EDIT.

package distributionpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new distributionpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for distributionpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDistributionPointByID(params *CreateDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDistributionPointByIDCreated, error)

	DeleteDistributionPointByID(params *DeleteDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistributionPointByIDOK, error)

	DeleteDistributionPointByName(params *DeleteDistributionPointByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistributionPointByNameOK, error)

	FindDistributionPoints(params *FindDistributionPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsOK, error)

	FindDistributionPointsByID(params *FindDistributionPointsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsByIDOK, error)

	FindDistributionPointsByName(params *FindDistributionPointsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsByNameOK, error)

	UpdateDistributionPointByID(params *UpdateDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistributionPointByIDCreated, error)

	UpdateDistributionPointByName(params *UpdateDistributionPointByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistributionPointByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDistributionPointByID creates a new distribution point by ID
*/
func (a *Client) CreateDistributionPointByID(params *CreateDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDistributionPointByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDistributionPointByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDistributionPointById",
		Method:             "POST",
		PathPattern:        "/distributionpoints/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDistributionPointByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateDistributionPointByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDistributionPointById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDistributionPointByID deletes a distribution point by ID
*/
func (a *Client) DeleteDistributionPointByID(params *DeleteDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistributionPointByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistributionPointByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDistributionPointById",
		Method:             "DELETE",
		PathPattern:        "/distributionpoints/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDistributionPointByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteDistributionPointByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDistributionPointById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDistributionPointByName deletes a distribution point by name
*/
func (a *Client) DeleteDistributionPointByName(params *DeleteDistributionPointByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDistributionPointByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDistributionPointByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDistributionPointByName",
		Method:             "DELETE",
		PathPattern:        "/distributionpoints/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDistributionPointByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteDistributionPointByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDistributionPointByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDistributionPoints finds all distribution points
*/
func (a *Client) FindDistributionPoints(params *FindDistributionPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDistributionPointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDistributionPoints",
		Method:             "GET",
		PathPattern:        "/distributionpoints",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDistributionPointsReader{formats: a.formats},
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
	success, ok := result.(*FindDistributionPointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDistributionPoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDistributionPointsByID finds distribution points by ID
*/
func (a *Client) FindDistributionPointsByID(params *FindDistributionPointsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDistributionPointsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDistributionPointsById",
		Method:             "GET",
		PathPattern:        "/distributionpoints/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDistributionPointsByIDReader{formats: a.formats},
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
	success, ok := result.(*FindDistributionPointsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDistributionPointsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindDistributionPointsByName finds distribution points by name
*/
func (a *Client) FindDistributionPointsByName(params *FindDistributionPointsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDistributionPointsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDistributionPointsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDistributionPointsByName",
		Method:             "GET",
		PathPattern:        "/distributionpoints/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindDistributionPointsByNameReader{formats: a.formats},
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
	success, ok := result.(*FindDistributionPointsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDistributionPointsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDistributionPointByID updates an existing distribution point by ID
*/
func (a *Client) UpdateDistributionPointByID(params *UpdateDistributionPointByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistributionPointByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDistributionPointByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDistributionPointById",
		Method:             "PUT",
		PathPattern:        "/distributionpoints/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDistributionPointByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDistributionPointByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDistributionPointById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDistributionPointByName updates an existing distribution point by name
*/
func (a *Client) UpdateDistributionPointByName(params *UpdateDistributionPointByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDistributionPointByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDistributionPointByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDistributionPointByName",
		Method:             "PUT",
		PathPattern:        "/distributionpoints/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDistributionPointByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateDistributionPointByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDistributionPointByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}