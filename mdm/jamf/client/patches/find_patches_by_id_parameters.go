// Code generated by go-swagger; DO NOT EDIT.

package patches

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

// NewFindPatchesByIDParams creates a new FindPatchesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPatchesByIDParams() *FindPatchesByIDParams {
	return &FindPatchesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPatchesByIDParamsWithTimeout creates a new FindPatchesByIDParams object
// with the ability to set a timeout on a request.
func NewFindPatchesByIDParamsWithTimeout(timeout time.Duration) *FindPatchesByIDParams {
	return &FindPatchesByIDParams{
		timeout: timeout,
	}
}

// NewFindPatchesByIDParamsWithContext creates a new FindPatchesByIDParams object
// with the ability to set a context for a request.
func NewFindPatchesByIDParamsWithContext(ctx context.Context) *FindPatchesByIDParams {
	return &FindPatchesByIDParams{
		Context: ctx,
	}
}

// NewFindPatchesByIDParamsWithHTTPClient creates a new FindPatchesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPatchesByIDParamsWithHTTPClient(client *http.Client) *FindPatchesByIDParams {
	return &FindPatchesByIDParams{
		HTTPClient: client,
	}
}

/*
FindPatchesByIDParams contains all the parameters to send to the API endpoint

	for the find patches by Id operation.

	Typically these are written to a http.Request.
*/
type FindPatchesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find patches by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPatchesByIDParams) WithDefaults() *FindPatchesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find patches by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPatchesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find patches by Id params
func (o *FindPatchesByIDParams) WithTimeout(timeout time.Duration) *FindPatchesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find patches by Id params
func (o *FindPatchesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find patches by Id params
func (o *FindPatchesByIDParams) WithContext(ctx context.Context) *FindPatchesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find patches by Id params
func (o *FindPatchesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find patches by Id params
func (o *FindPatchesByIDParams) WithHTTPClient(client *http.Client) *FindPatchesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find patches by Id params
func (o *FindPatchesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find patches by Id params
func (o *FindPatchesByIDParams) WithID(id int64) *FindPatchesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find patches by Id params
func (o *FindPatchesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindPatchesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
