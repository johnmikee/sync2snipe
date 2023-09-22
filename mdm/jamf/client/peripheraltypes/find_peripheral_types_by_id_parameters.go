// Code generated by go-swagger; DO NOT EDIT.

package peripheraltypes

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

// NewFindPeripheralTypesByIDParams creates a new FindPeripheralTypesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPeripheralTypesByIDParams() *FindPeripheralTypesByIDParams {
	return &FindPeripheralTypesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPeripheralTypesByIDParamsWithTimeout creates a new FindPeripheralTypesByIDParams object
// with the ability to set a timeout on a request.
func NewFindPeripheralTypesByIDParamsWithTimeout(timeout time.Duration) *FindPeripheralTypesByIDParams {
	return &FindPeripheralTypesByIDParams{
		timeout: timeout,
	}
}

// NewFindPeripheralTypesByIDParamsWithContext creates a new FindPeripheralTypesByIDParams object
// with the ability to set a context for a request.
func NewFindPeripheralTypesByIDParamsWithContext(ctx context.Context) *FindPeripheralTypesByIDParams {
	return &FindPeripheralTypesByIDParams{
		Context: ctx,
	}
}

// NewFindPeripheralTypesByIDParamsWithHTTPClient creates a new FindPeripheralTypesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPeripheralTypesByIDParamsWithHTTPClient(client *http.Client) *FindPeripheralTypesByIDParams {
	return &FindPeripheralTypesByIDParams{
		HTTPClient: client,
	}
}

/*
FindPeripheralTypesByIDParams contains all the parameters to send to the API endpoint

	for the find peripheral types by Id operation.

	Typically these are written to a http.Request.
*/
type FindPeripheralTypesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find peripheral types by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralTypesByIDParams) WithDefaults() *FindPeripheralTypesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find peripheral types by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralTypesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) WithTimeout(timeout time.Duration) *FindPeripheralTypesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) WithContext(ctx context.Context) *FindPeripheralTypesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) WithHTTPClient(client *http.Client) *FindPeripheralTypesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) WithID(id int64) *FindPeripheralTypesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find peripheral types by Id params
func (o *FindPeripheralTypesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindPeripheralTypesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
