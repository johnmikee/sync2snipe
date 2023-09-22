// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

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

// NewDeleteDirectoryBindingByNameParams creates a new DeleteDirectoryBindingByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDirectoryBindingByNameParams() *DeleteDirectoryBindingByNameParams {
	return &DeleteDirectoryBindingByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDirectoryBindingByNameParamsWithTimeout creates a new DeleteDirectoryBindingByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteDirectoryBindingByNameParamsWithTimeout(timeout time.Duration) *DeleteDirectoryBindingByNameParams {
	return &DeleteDirectoryBindingByNameParams{
		timeout: timeout,
	}
}

// NewDeleteDirectoryBindingByNameParamsWithContext creates a new DeleteDirectoryBindingByNameParams object
// with the ability to set a context for a request.
func NewDeleteDirectoryBindingByNameParamsWithContext(ctx context.Context) *DeleteDirectoryBindingByNameParams {
	return &DeleteDirectoryBindingByNameParams{
		Context: ctx,
	}
}

// NewDeleteDirectoryBindingByNameParamsWithHTTPClient creates a new DeleteDirectoryBindingByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDirectoryBindingByNameParamsWithHTTPClient(client *http.Client) *DeleteDirectoryBindingByNameParams {
	return &DeleteDirectoryBindingByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteDirectoryBindingByNameParams contains all the parameters to send to the API endpoint

	for the delete directory binding by name operation.

	Typically these are written to a http.Request.
*/
type DeleteDirectoryBindingByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete directory binding by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDirectoryBindingByNameParams) WithDefaults() *DeleteDirectoryBindingByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete directory binding by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDirectoryBindingByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) WithTimeout(timeout time.Duration) *DeleteDirectoryBindingByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) WithContext(ctx context.Context) *DeleteDirectoryBindingByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) WithHTTPClient(client *http.Client) *DeleteDirectoryBindingByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) WithName(name string) *DeleteDirectoryBindingByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete directory binding by name params
func (o *DeleteDirectoryBindingByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDirectoryBindingByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
