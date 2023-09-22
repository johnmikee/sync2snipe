// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

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

// NewFindManagedPreferenceProfilesByIDParams creates a new FindManagedPreferenceProfilesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindManagedPreferenceProfilesByIDParams() *FindManagedPreferenceProfilesByIDParams {
	return &FindManagedPreferenceProfilesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindManagedPreferenceProfilesByIDParamsWithTimeout creates a new FindManagedPreferenceProfilesByIDParams object
// with the ability to set a timeout on a request.
func NewFindManagedPreferenceProfilesByIDParamsWithTimeout(timeout time.Duration) *FindManagedPreferenceProfilesByIDParams {
	return &FindManagedPreferenceProfilesByIDParams{
		timeout: timeout,
	}
}

// NewFindManagedPreferenceProfilesByIDParamsWithContext creates a new FindManagedPreferenceProfilesByIDParams object
// with the ability to set a context for a request.
func NewFindManagedPreferenceProfilesByIDParamsWithContext(ctx context.Context) *FindManagedPreferenceProfilesByIDParams {
	return &FindManagedPreferenceProfilesByIDParams{
		Context: ctx,
	}
}

// NewFindManagedPreferenceProfilesByIDParamsWithHTTPClient creates a new FindManagedPreferenceProfilesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindManagedPreferenceProfilesByIDParamsWithHTTPClient(client *http.Client) *FindManagedPreferenceProfilesByIDParams {
	return &FindManagedPreferenceProfilesByIDParams{
		HTTPClient: client,
	}
}

/*
FindManagedPreferenceProfilesByIDParams contains all the parameters to send to the API endpoint

	for the find managed preference profiles by Id operation.

	Typically these are written to a http.Request.
*/
type FindManagedPreferenceProfilesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find managed preference profiles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindManagedPreferenceProfilesByIDParams) WithDefaults() *FindManagedPreferenceProfilesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find managed preference profiles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindManagedPreferenceProfilesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) WithTimeout(timeout time.Duration) *FindManagedPreferenceProfilesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) WithContext(ctx context.Context) *FindManagedPreferenceProfilesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) WithHTTPClient(client *http.Client) *FindManagedPreferenceProfilesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) WithID(id int64) *FindManagedPreferenceProfilesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find managed preference profiles by Id params
func (o *FindManagedPreferenceProfilesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindManagedPreferenceProfilesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
