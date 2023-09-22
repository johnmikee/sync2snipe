// Code generated by go-swagger; DO NOT EDIT.

package softwareupdateservers

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

// NewDeleteSoftwareUpdateServerByNameParams creates a new DeleteSoftwareUpdateServerByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSoftwareUpdateServerByNameParams() *DeleteSoftwareUpdateServerByNameParams {
	return &DeleteSoftwareUpdateServerByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSoftwareUpdateServerByNameParamsWithTimeout creates a new DeleteSoftwareUpdateServerByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteSoftwareUpdateServerByNameParamsWithTimeout(timeout time.Duration) *DeleteSoftwareUpdateServerByNameParams {
	return &DeleteSoftwareUpdateServerByNameParams{
		timeout: timeout,
	}
}

// NewDeleteSoftwareUpdateServerByNameParamsWithContext creates a new DeleteSoftwareUpdateServerByNameParams object
// with the ability to set a context for a request.
func NewDeleteSoftwareUpdateServerByNameParamsWithContext(ctx context.Context) *DeleteSoftwareUpdateServerByNameParams {
	return &DeleteSoftwareUpdateServerByNameParams{
		Context: ctx,
	}
}

// NewDeleteSoftwareUpdateServerByNameParamsWithHTTPClient creates a new DeleteSoftwareUpdateServerByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSoftwareUpdateServerByNameParamsWithHTTPClient(client *http.Client) *DeleteSoftwareUpdateServerByNameParams {
	return &DeleteSoftwareUpdateServerByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteSoftwareUpdateServerByNameParams contains all the parameters to send to the API endpoint

	for the delete software update server by name operation.

	Typically these are written to a http.Request.
*/
type DeleteSoftwareUpdateServerByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete software update server by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareUpdateServerByNameParams) WithDefaults() *DeleteSoftwareUpdateServerByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete software update server by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareUpdateServerByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) WithTimeout(timeout time.Duration) *DeleteSoftwareUpdateServerByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) WithContext(ctx context.Context) *DeleteSoftwareUpdateServerByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) WithHTTPClient(client *http.Client) *DeleteSoftwareUpdateServerByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) WithName(name string) *DeleteSoftwareUpdateServerByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete software update server by name params
func (o *DeleteSoftwareUpdateServerByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSoftwareUpdateServerByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
