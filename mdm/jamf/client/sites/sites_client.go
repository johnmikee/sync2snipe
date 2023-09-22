// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sites API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sites API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSiteByID(params *CreateSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSiteByIDCreated, error)

	DeleteSiteByID(params *DeleteSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSiteByIDOK, error)

	DeleteSiteByName(params *DeleteSiteByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSiteByNameOK, error)

	FindSites(params *FindSitesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesOK, error)

	FindSitesByID(params *FindSitesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesByIDOK, error)

	FindSitesByName(params *FindSitesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesByNameOK, error)

	UpdateSiteByID(params *UpdateSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSiteByIDCreated, error)

	UpdateSiteByName(params *UpdateSiteByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSiteByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSiteByID creates a new site by ID
*/
func (a *Client) CreateSiteByID(params *CreateSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSiteByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSiteByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSiteById",
		Method:             "POST",
		PathPattern:        "/sites/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSiteByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateSiteByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSiteById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSiteByID deletes a site by ID
*/
func (a *Client) DeleteSiteByID(params *DeleteSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSiteByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSiteByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSiteById",
		Method:             "DELETE",
		PathPattern:        "/sites/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSiteByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteSiteByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSiteById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSiteByName deletes a site by name
*/
func (a *Client) DeleteSiteByName(params *DeleteSiteByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSiteByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSiteByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSiteByName",
		Method:             "DELETE",
		PathPattern:        "/sites/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSiteByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteSiteByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSiteByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindSites finds all sites
*/
func (a *Client) FindSites(params *FindSitesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSitesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSites",
		Method:             "GET",
		PathPattern:        "/sites",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindSitesReader{formats: a.formats},
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
	success, ok := result.(*FindSitesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSites: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindSitesByID finds sites by ID
*/
func (a *Client) FindSitesByID(params *FindSitesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSitesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSitesById",
		Method:             "GET",
		PathPattern:        "/sites/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindSitesByIDReader{formats: a.formats},
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
	success, ok := result.(*FindSitesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSitesById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindSitesByName finds sites by name
*/
func (a *Client) FindSitesByName(params *FindSitesByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSitesByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSitesByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSitesByName",
		Method:             "GET",
		PathPattern:        "/sites/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindSitesByNameReader{formats: a.formats},
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
	success, ok := result.(*FindSitesByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSitesByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSiteByID updates an existing site by ID
*/
func (a *Client) UpdateSiteByID(params *UpdateSiteByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSiteByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSiteByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSiteById",
		Method:             "PUT",
		PathPattern:        "/sites/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSiteByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateSiteByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSiteById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSiteByName updates an existing site by name
*/
func (a *Client) UpdateSiteByName(params *UpdateSiteByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSiteByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSiteByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSiteByName",
		Method:             "PUT",
		PathPattern:        "/sites/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSiteByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdateSiteByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSiteByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}