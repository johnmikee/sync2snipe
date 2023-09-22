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

// NewUpdateDirectoryBindingByNameParams creates a new UpdateDirectoryBindingByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDirectoryBindingByNameParams() *UpdateDirectoryBindingByNameParams {
	return &UpdateDirectoryBindingByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDirectoryBindingByNameParamsWithTimeout creates a new UpdateDirectoryBindingByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateDirectoryBindingByNameParamsWithTimeout(timeout time.Duration) *UpdateDirectoryBindingByNameParams {
	return &UpdateDirectoryBindingByNameParams{
		timeout: timeout,
	}
}

// NewUpdateDirectoryBindingByNameParamsWithContext creates a new UpdateDirectoryBindingByNameParams object
// with the ability to set a context for a request.
func NewUpdateDirectoryBindingByNameParamsWithContext(ctx context.Context) *UpdateDirectoryBindingByNameParams {
	return &UpdateDirectoryBindingByNameParams{
		Context: ctx,
	}
}

// NewUpdateDirectoryBindingByNameParamsWithHTTPClient creates a new UpdateDirectoryBindingByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDirectoryBindingByNameParamsWithHTTPClient(client *http.Client) *UpdateDirectoryBindingByNameParams {
	return &UpdateDirectoryBindingByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateDirectoryBindingByNameParams contains all the parameters to send to the API endpoint

	for the update directory binding by name operation.

	Typically these are written to a http.Request.
*/
type UpdateDirectoryBindingByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update directory binding by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDirectoryBindingByNameParams) WithDefaults() *UpdateDirectoryBindingByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update directory binding by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDirectoryBindingByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) WithTimeout(timeout time.Duration) *UpdateDirectoryBindingByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) WithContext(ctx context.Context) *UpdateDirectoryBindingByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) WithHTTPClient(client *http.Client) *UpdateDirectoryBindingByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) WithName(name string) *UpdateDirectoryBindingByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update directory binding by name params
func (o *UpdateDirectoryBindingByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDirectoryBindingByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
