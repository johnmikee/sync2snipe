// Code generated by go-swagger; DO NOT EDIT.

package printers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new printers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for printers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePrinterByID(params *CreatePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePrinterByIDCreated, error)

	DeletePrinterByID(params *DeletePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePrinterByIDOK, error)

	DeletePrinterByName(params *DeletePrinterByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePrinterByNameOK, error)

	FindPrinters(params *FindPrintersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersOK, error)

	FindPrintersByID(params *FindPrintersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersByIDOK, error)

	FindPrintersByName(params *FindPrintersByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersByNameOK, error)

	UpdatePrinterByID(params *UpdatePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrinterByIDCreated, error)

	UpdatePrinterByName(params *UpdatePrinterByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrinterByNameCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePrinterByID creates a new printer by ID
*/
func (a *Client) CreatePrinterByID(params *CreatePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePrinterByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrinterByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrinterById",
		Method:             "POST",
		PathPattern:        "/printers/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePrinterByIDReader{formats: a.formats},
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
	success, ok := result.(*CreatePrinterByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPrinterById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePrinterByID deletes a printer by ID
*/
func (a *Client) DeletePrinterByID(params *DeletePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePrinterByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrinterByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePrinterById",
		Method:             "DELETE",
		PathPattern:        "/printers/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePrinterByIDReader{formats: a.formats},
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
	success, ok := result.(*DeletePrinterByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePrinterById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePrinterByName deletes a printer by name
*/
func (a *Client) DeletePrinterByName(params *DeletePrinterByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePrinterByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrinterByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePrinterByName",
		Method:             "DELETE",
		PathPattern:        "/printers/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePrinterByNameReader{formats: a.formats},
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
	success, ok := result.(*DeletePrinterByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePrinterByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPrinters finds all printers
*/
func (a *Client) FindPrinters(params *FindPrintersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPrintersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPrinters",
		Method:             "GET",
		PathPattern:        "/printers",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPrintersReader{formats: a.formats},
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
	success, ok := result.(*FindPrintersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPrinters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPrintersByID finds printers by ID
*/
func (a *Client) FindPrintersByID(params *FindPrintersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPrintersByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPrintersById",
		Method:             "GET",
		PathPattern:        "/printers/id/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPrintersByIDReader{formats: a.formats},
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
	success, ok := result.(*FindPrintersByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPrintersById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindPrintersByName finds printers by name
*/
func (a *Client) FindPrintersByName(params *FindPrintersByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPrintersByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPrintersByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPrintersByName",
		Method:             "GET",
		PathPattern:        "/printers/name/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPrintersByNameReader{formats: a.formats},
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
	success, ok := result.(*FindPrintersByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPrintersByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePrinterByID updates an existing printer by ID
*/
func (a *Client) UpdatePrinterByID(params *UpdatePrinterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrinterByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrinterByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrinterById",
		Method:             "PUT",
		PathPattern:        "/printers/id/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePrinterByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdatePrinterByIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePrinterById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePrinterByName updates an existing printer by name
*/
func (a *Client) UpdatePrinterByName(params *UpdatePrinterByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrinterByNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrinterByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrinterByName",
		Method:             "PUT",
		PathPattern:        "/printers/name/{name}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePrinterByNameReader{formats: a.formats},
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
	success, ok := result.(*UpdatePrinterByNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePrinterByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}