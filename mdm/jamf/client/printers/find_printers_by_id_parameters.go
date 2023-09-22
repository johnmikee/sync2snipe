// Code generated by go-swagger; DO NOT EDIT.

package printers

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

// NewFindPrintersByIDParams creates a new FindPrintersByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPrintersByIDParams() *FindPrintersByIDParams {
	return &FindPrintersByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPrintersByIDParamsWithTimeout creates a new FindPrintersByIDParams object
// with the ability to set a timeout on a request.
func NewFindPrintersByIDParamsWithTimeout(timeout time.Duration) *FindPrintersByIDParams {
	return &FindPrintersByIDParams{
		timeout: timeout,
	}
}

// NewFindPrintersByIDParamsWithContext creates a new FindPrintersByIDParams object
// with the ability to set a context for a request.
func NewFindPrintersByIDParamsWithContext(ctx context.Context) *FindPrintersByIDParams {
	return &FindPrintersByIDParams{
		Context: ctx,
	}
}

// NewFindPrintersByIDParamsWithHTTPClient creates a new FindPrintersByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPrintersByIDParamsWithHTTPClient(client *http.Client) *FindPrintersByIDParams {
	return &FindPrintersByIDParams{
		HTTPClient: client,
	}
}

/*
FindPrintersByIDParams contains all the parameters to send to the API endpoint

	for the find printers by Id operation.

	Typically these are written to a http.Request.
*/
type FindPrintersByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find printers by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPrintersByIDParams) WithDefaults() *FindPrintersByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find printers by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPrintersByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find printers by Id params
func (o *FindPrintersByIDParams) WithTimeout(timeout time.Duration) *FindPrintersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find printers by Id params
func (o *FindPrintersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find printers by Id params
func (o *FindPrintersByIDParams) WithContext(ctx context.Context) *FindPrintersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find printers by Id params
func (o *FindPrintersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find printers by Id params
func (o *FindPrintersByIDParams) WithHTTPClient(client *http.Client) *FindPrintersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find printers by Id params
func (o *FindPrintersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find printers by Id params
func (o *FindPrintersByIDParams) WithID(id int64) *FindPrintersByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find printers by Id params
func (o *FindPrintersByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindPrintersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
