// Code generated by go-swagger; DO NOT EDIT.

package jssuser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new jssuser API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jssuser API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	JssuserGet(params *JssuserGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JssuserGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
JssuserGet returns basic information about jamf pro as well as privileges of the person requesting the resource deprecated
*/
func (a *Client) JssuserGet(params *JssuserGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JssuserGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJssuserGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "JssuserGet",
		Method:             "GET",
		PathPattern:        "/jssuser",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JssuserGetReader{formats: a.formats},
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
	success, ok := result.(*JssuserGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for JssuserGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
