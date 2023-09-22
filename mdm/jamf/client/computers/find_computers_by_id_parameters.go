// Code generated by go-swagger; DO NOT EDIT.

package computers

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

// NewFindComputersByIDParams creates a new FindComputersByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputersByIDParams() *FindComputersByIDParams {
	return &FindComputersByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputersByIDParamsWithTimeout creates a new FindComputersByIDParams object
// with the ability to set a timeout on a request.
func NewFindComputersByIDParamsWithTimeout(timeout time.Duration) *FindComputersByIDParams {
	return &FindComputersByIDParams{
		timeout: timeout,
	}
}

// NewFindComputersByIDParamsWithContext creates a new FindComputersByIDParams object
// with the ability to set a context for a request.
func NewFindComputersByIDParamsWithContext(ctx context.Context) *FindComputersByIDParams {
	return &FindComputersByIDParams{
		Context: ctx,
	}
}

// NewFindComputersByIDParamsWithHTTPClient creates a new FindComputersByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputersByIDParamsWithHTTPClient(client *http.Client) *FindComputersByIDParams {
	return &FindComputersByIDParams{
		HTTPClient: client,
	}
}

/*
FindComputersByIDParams contains all the parameters to send to the API endpoint

	for the find computers by Id operation.

	Typically these are written to a http.Request.
*/
type FindComputersByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computers by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputersByIDParams) WithDefaults() *FindComputersByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computers by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputersByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computers by Id params
func (o *FindComputersByIDParams) WithTimeout(timeout time.Duration) *FindComputersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computers by Id params
func (o *FindComputersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computers by Id params
func (o *FindComputersByIDParams) WithContext(ctx context.Context) *FindComputersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computers by Id params
func (o *FindComputersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computers by Id params
func (o *FindComputersByIDParams) WithHTTPClient(client *http.Client) *FindComputersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computers by Id params
func (o *FindComputersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find computers by Id params
func (o *FindComputersByIDParams) WithID(id int64) *FindComputersByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find computers by Id params
func (o *FindComputersByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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