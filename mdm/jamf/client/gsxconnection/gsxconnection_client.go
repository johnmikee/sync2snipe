// Code generated by go-swagger; DO NOT EDIT.

package gsxconnection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gsxconnection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gsxconnection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FindGSXConnection(params *FindGSXConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindGSXConnectionOK, error)

	UpdateGSXConnection(params *UpdateGSXConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGSXConnectionCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
FindGSXConnection finds the jamf pro g s x connection information
*/
func (a *Client) FindGSXConnection(params *FindGSXConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindGSXConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindGSXConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findGSXConnection",
		Method:             "GET",
		PathPattern:        "/gsxconnection",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindGSXConnectionReader{formats: a.formats},
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
	success, ok := result.(*FindGSXConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findGSXConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateGSXConnection updates the jamf pro g s x connection information
*/
func (a *Client) UpdateGSXConnection(params *UpdateGSXConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGSXConnectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGSXConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateGSXConnection",
		Method:             "PUT",
		PathPattern:        "/gsxconnection",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateGSXConnectionReader{formats: a.formats},
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
	success, ok := result.(*UpdateGSXConnectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateGSXConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
