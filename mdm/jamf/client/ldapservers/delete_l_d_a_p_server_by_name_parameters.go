// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteLDAPServerByNameParams creates a new DeleteLDAPServerByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLDAPServerByNameParams() *DeleteLDAPServerByNameParams {
	return &DeleteLDAPServerByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLDAPServerByNameParamsWithTimeout creates a new DeleteLDAPServerByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteLDAPServerByNameParamsWithTimeout(timeout time.Duration) *DeleteLDAPServerByNameParams {
	return &DeleteLDAPServerByNameParams{
		timeout: timeout,
	}
}

// NewDeleteLDAPServerByNameParamsWithContext creates a new DeleteLDAPServerByNameParams object
// with the ability to set a context for a request.
func NewDeleteLDAPServerByNameParamsWithContext(ctx context.Context) *DeleteLDAPServerByNameParams {
	return &DeleteLDAPServerByNameParams{
		Context: ctx,
	}
}

// NewDeleteLDAPServerByNameParamsWithHTTPClient creates a new DeleteLDAPServerByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLDAPServerByNameParamsWithHTTPClient(client *http.Client) *DeleteLDAPServerByNameParams {
	return &DeleteLDAPServerByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteLDAPServerByNameParams contains all the parameters to send to the API endpoint

	for the delete l d a p server by name operation.

	Typically these are written to a http.Request.
*/
type DeleteLDAPServerByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete l d a p server by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLDAPServerByNameParams) WithDefaults() *DeleteLDAPServerByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete l d a p server by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLDAPServerByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) WithTimeout(timeout time.Duration) *DeleteLDAPServerByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) WithContext(ctx context.Context) *DeleteLDAPServerByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) WithHTTPClient(client *http.Client) *DeleteLDAPServerByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) WithName(name string) *DeleteLDAPServerByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete l d a p server by name params
func (o *DeleteLDAPServerByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLDAPServerByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}