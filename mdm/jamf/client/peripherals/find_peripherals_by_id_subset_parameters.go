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

// NewFindPeripheralsByIDSubsetParams creates a new FindPeripheralsByIDSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPeripheralsByIDSubsetParams() *FindPeripheralsByIDSubsetParams {
	return &FindPeripheralsByIDSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPeripheralsByIDSubsetParamsWithTimeout creates a new FindPeripheralsByIDSubsetParams object
// with the ability to set a timeout on a request.
func NewFindPeripheralsByIDSubsetParamsWithTimeout(timeout time.Duration) *FindPeripheralsByIDSubsetParams {
	return &FindPeripheralsByIDSubsetParams{
		timeout: timeout,
	}
}

// NewFindPeripheralsByIDSubsetParamsWithContext creates a new FindPeripheralsByIDSubsetParams object
// with the ability to set a context for a request.
func NewFindPeripheralsByIDSubsetParamsWithContext(ctx context.Context) *FindPeripheralsByIDSubsetParams {
	return &FindPeripheralsByIDSubsetParams{
		Context: ctx,
	}
}

// NewFindPeripheralsByIDSubsetParamsWithHTTPClient creates a new FindPeripheralsByIDSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPeripheralsByIDSubsetParamsWithHTTPClient(client *http.Client) *FindPeripheralsByIDSubsetParams {
	return &FindPeripheralsByIDSubsetParams{
		HTTPClient: client,
	}
}

/*
FindPeripheralsByIDSubsetParams contains all the parameters to send to the API endpoint

	for the find peripherals by Id subset operation.

	Typically these are written to a http.Request.
*/
type FindPeripheralsByIDSubsetParams struct {

	/* ID.

	   ID to filter by
	*/
	ID int64

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find peripherals by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralsByIDSubsetParams) WithDefaults() *FindPeripheralsByIDSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find peripherals by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPeripheralsByIDSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) WithTimeout(timeout time.Duration) *FindPeripheralsByIDSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) WithContext(ctx context.Context) *FindPeripheralsByIDSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) WithHTTPClient(client *http.Client) *FindPeripheralsByIDSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) WithID(id int64) *FindPeripheralsByIDSubsetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) SetID(id int64) {
	o.ID = id
}

// WithSubset adds the subset to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) WithSubset(subset string) *FindPeripheralsByIDSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find peripherals by Id subset params
func (o *FindPeripheralsByIDSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindPeripheralsByIDSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param subset
	if err := r.SetPathParam("subset", o.Subset); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}