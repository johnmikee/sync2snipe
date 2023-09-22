// Code generated by go-swagger; DO NOT EDIT.

package fileuploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fileuploads API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fileuploads API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UploadFiles(params *UploadFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFilesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
UploadFiles creates file attachments in jamf pro

Here is a sample command curl -k -u user:password https://my.server.com:8443/JSSResource/fileuploads/computers/id/2 -F name=@/Users/admin/Documents/Sample.doc -X POST
*/
func (a *Client) UploadFiles(params *UploadFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFilesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadFiles",
		Method:             "POST",
		PathPattern:        "/fileuploads/{resource}/{idType}/{id}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadFilesReader{formats: a.formats},
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
	success, ok := result.(*UploadFilesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadFiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
