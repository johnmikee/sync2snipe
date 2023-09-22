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
	"github.com/go-openapi/swag"
)

// NewDeleteLDAPServerByIDParams creates a new DeleteLDAPServerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLDAPServerByIDParams() *DeleteLDAPServerByIDParams {
	return &DeleteLDAPServerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLDAPServerByIDParamsWithTimeout creates a new DeleteLDAPServerByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteLDAPServerByIDParamsWithTimeout(timeout time.Duration) *DeleteLDAPServerByIDParams {
	return &DeleteLDAPServerByIDParams{
		timeout: timeout,
	}
}

// NewDeleteLDAPServerByIDParamsWithContext creates a new DeleteLDAPServerByIDParams object
// with the ability to set a context for a request.
func NewDeleteLDAPServerByIDParamsWithContext(ctx context.Context) *DeleteLDAPServerByIDParams {
	return &DeleteLDAPServerByIDParams{
		Context: ctx,
	}
}

// NewDeleteLDAPServerByIDParamsWithHTTPClient creates a new DeleteLDAPServerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLDAPServerByIDParamsWithHTTPClient(client *http.Client) *DeleteLDAPServerByIDParams {
	return &DeleteLDAPServerByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteLDAPServerByIDParams contains all the parameters to send to the API endpoint

	for the delete l d a p server by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteLDAPServerByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete l d a p server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLDAPServerByIDParams) WithDefaults() *DeleteLDAPServerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete l d a p server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLDAPServerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) WithTimeout(timeout time.Duration) *DeleteLDAPServerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) WithContext(ctx context.Context) *DeleteLDAPServerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) WithHTTPClient(client *http.Client) *DeleteLDAPServerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) WithID(id int64) *DeleteLDAPServerByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete l d a p server by Id params
func (o *DeleteLDAPServerByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLDAPServerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}