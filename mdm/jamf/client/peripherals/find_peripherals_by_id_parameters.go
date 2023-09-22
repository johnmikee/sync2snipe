// Code generated by go-swagger; DO NOT EDIT.

package peripherals

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

// NewFindPeripheralsByIDParams creates a new FindPeripheralsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPeripheralsByIDParams() *FindPeripheralsByIDParams {
	return &FindPeripheralsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPeripheralsByIDParamsWithTimeout creates a new FindPeripheralsByIDParams object
// with the ability to set a timeout on a request.
func NewFindPeripheralsByIDParamsWithTimeout(timeout time.Duration) *FindPeripheralsByIDParams {
	return &FindPeripheralsByIDParams{
		timeout: timeout,
	}
}

// NewFindPeripheralsByIDParamsWithContext creates a new FindPeripheralsByIDParams object
// with the ability to set a context for a request.
func NewFindPeripheralsByIDParamsWithContext(ctx context.Context) *FindPeripheralsByIDParams {
	return &FindPeripheralsByIDParams{
		Context: ctx,
	}
}

// NewFindPeripheralsByIDParamsWithHTTPClient creates a new FindPeripheralsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPeripheralsByIDParamsWithHTTPClient(client *http.Client) *FindPeripheralsByIDParams {
	return &FindPeripheralsByIDParams{
		HTTPClient: client,
	}
}

/*
FindPeripheralsByIDParams contains all the parameters to send to the API endpoint

	for the find peripherals by Id operation.

	Typically these are written to a http.Request.
*/
type FindPeripheralsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find peripherals by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralsByIDParams) WithDefaults() *FindPeripheralsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find peripherals by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) WithTimeout(timeout time.Duration) *FindPeripheralsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) WithContext(ctx context.Context) *FindPeripheralsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) WithHTTPClient(client *http.Client) *FindPeripheralsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) WithID(id int64) *FindPeripheralsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find peripherals by Id params
func (o *FindPeripheralsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindPeripheralsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
