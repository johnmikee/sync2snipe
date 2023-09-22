// Code generated by go-swagger; DO NOT EDIT.

package computerinventorycollection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new computerinventorycollection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for computerinventorycollection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FindComputerInventoryCollection(params *FindComputerInventoryCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerInventoryCollectionOK, error)

	UpdateComputerInventoryCollection(params *UpdateComputerInventoryCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerInventoryCollectionCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
FindComputerInventoryCollection finds the jamf pro computer inventory collection information
*/
func (a *Client) FindComputerInventoryCollection(params *FindComputerInventoryCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindComputerInventoryCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindComputerInventoryCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findComputerInventoryCollection",
		Method:             "GET",
		PathPattern:        "/computerinventorycollection",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindComputerInventoryCollectionReader{formats: a.formats},
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
	success, ok := result.(*FindComputerInventoryCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findComputerInventoryCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateComputerInventoryCollection updates the jamf pro computer inventory collection information
*/
func (a *Client) UpdateComputerInventoryCollection(params *UpdateComputerInventoryCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateComputerInventoryCollectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateComputerInventoryCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateComputerInventoryCollection",
		Method:             "PUT",
		PathPattern:        "/computerinventorycollection",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateComputerInventoryCollectionReader{formats: a.formats},
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
	success, ok := result.(*UpdateComputerInventoryCollectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateComputerInventoryCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}