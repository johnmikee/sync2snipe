// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewFindGroupsByIDParams creates a new FindGroupsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindGroupsByIDParams() *FindGroupsByIDParams {
	return &FindGroupsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindGroupsByIDParamsWithTimeout creates a new FindGroupsByIDParams object
// with the ability to set a timeout on a request.
func NewFindGroupsByIDParamsWithTimeout(timeout time.Duration) *FindGroupsByIDParams {
	return &FindGroupsByIDParams{
		timeout: timeout,
	}
}

// NewFindGroupsByIDParamsWithContext creates a new FindGroupsByIDParams object
// with the ability to set a context for a request.
func NewFindGroupsByIDParamsWithContext(ctx context.Context) *FindGroupsByIDParams {
	return &FindGroupsByIDParams{
		Context: ctx,
	}
}

// NewFindGroupsByIDParamsWithHTTPClient creates a new FindGroupsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindGroupsByIDParamsWithHTTPClient(client *http.Client) *FindGroupsByIDParams {
	return &FindGroupsByIDParams{
		HTTPClient: client,
	}
}

/*
FindGroupsByIDParams contains all the parameters to send to the API endpoint

	for the find groups by Id operation.

	Typically these are written to a http.Request.
*/
type FindGroupsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindGroupsByIDParams) WithDefaults() *FindGroupsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindGroupsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find groups by Id params
func (o *FindGroupsByIDParams) WithTimeout(timeout time.Duration) *FindGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find groups by Id params
func (o *FindGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find groups by Id params
func (o *FindGroupsByIDParams) WithContext(ctx context.Context) *FindGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find groups by Id params
func (o *FindGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find groups by Id params
func (o *FindGroupsByIDParams) WithHTTPClient(client *http.Client) *FindGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find groups by Id params
func (o *FindGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find groups by Id params
func (o *FindGroupsByIDParams) WithID(id int64) *FindGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find groups by Id params
func (o *FindGroupsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
