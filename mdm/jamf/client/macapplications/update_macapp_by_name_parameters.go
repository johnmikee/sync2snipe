// Code generated by go-swagger; DO NOT EDIT.

package macapplications

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

// NewUpdateMacappByNameParams creates a new UpdateMacappByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMacappByNameParams() *UpdateMacappByNameParams {
	return &UpdateMacappByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMacappByNameParamsWithTimeout creates a new UpdateMacappByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateMacappByNameParamsWithTimeout(timeout time.Duration) *UpdateMacappByNameParams {
	return &UpdateMacappByNameParams{
		timeout: timeout,
	}
}

// NewUpdateMacappByNameParamsWithContext creates a new UpdateMacappByNameParams object
// with the ability to set a context for a request.
func NewUpdateMacappByNameParamsWithContext(ctx context.Context) *UpdateMacappByNameParams {
	return &UpdateMacappByNameParams{
		Context: ctx,
	}
}

// NewUpdateMacappByNameParamsWithHTTPClient creates a new UpdateMacappByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMacappByNameParamsWithHTTPClient(client *http.Client) *UpdateMacappByNameParams {
	return &UpdateMacappByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateMacappByNameParams contains all the parameters to send to the API endpoint

	for the update macapp by name operation.

	Typically these are written to a http.Request.
*/
type UpdateMacappByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update macapp by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMacappByNameParams) WithDefaults() *UpdateMacappByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update macapp by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMacappByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update macapp by name params
func (o *UpdateMacappByNameParams) WithTimeout(timeout time.Duration) *UpdateMacappByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update macapp by name params
func (o *UpdateMacappByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update macapp by name params
func (o *UpdateMacappByNameParams) WithContext(ctx context.Context) *UpdateMacappByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update macapp by name params
func (o *UpdateMacappByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update macapp by name params
func (o *UpdateMacappByNameParams) WithHTTPClient(client *http.Client) *UpdateMacappByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update macapp by name params
func (o *UpdateMacappByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update macapp by name params
func (o *UpdateMacappByNameParams) WithName(name string) *UpdateMacappByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update macapp by name params
func (o *UpdateMacappByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMacappByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
