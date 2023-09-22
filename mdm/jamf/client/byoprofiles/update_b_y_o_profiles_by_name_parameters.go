// Code generated by go-swagger; DO NOT EDIT.

package byoprofiles

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

// NewUpdateBYOProfilesByNameParams creates a new UpdateBYOProfilesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBYOProfilesByNameParams() *UpdateBYOProfilesByNameParams {
	return &UpdateBYOProfilesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBYOProfilesByNameParamsWithTimeout creates a new UpdateBYOProfilesByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateBYOProfilesByNameParamsWithTimeout(timeout time.Duration) *UpdateBYOProfilesByNameParams {
	return &UpdateBYOProfilesByNameParams{
		timeout: timeout,
	}
}

// NewUpdateBYOProfilesByNameParamsWithContext creates a new UpdateBYOProfilesByNameParams object
// with the ability to set a context for a request.
func NewUpdateBYOProfilesByNameParamsWithContext(ctx context.Context) *UpdateBYOProfilesByNameParams {
	return &UpdateBYOProfilesByNameParams{
		Context: ctx,
	}
}

// NewUpdateBYOProfilesByNameParamsWithHTTPClient creates a new UpdateBYOProfilesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBYOProfilesByNameParamsWithHTTPClient(client *http.Client) *UpdateBYOProfilesByNameParams {
	return &UpdateBYOProfilesByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateBYOProfilesByNameParams contains all the parameters to send to the API endpoint

	for the update b y o profiles by name operation.

	Typically these are written to a http.Request.
*/
type UpdateBYOProfilesByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update b y o profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBYOProfilesByNameParams) WithDefaults() *UpdateBYOProfilesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update b y o profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBYOProfilesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) WithTimeout(timeout time.Duration) *UpdateBYOProfilesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) WithContext(ctx context.Context) *UpdateBYOProfilesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) WithHTTPClient(client *http.Client) *UpdateBYOProfilesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) WithName(name string) *UpdateBYOProfilesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update b y o profiles by name params
func (o *UpdateBYOProfilesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBYOProfilesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
