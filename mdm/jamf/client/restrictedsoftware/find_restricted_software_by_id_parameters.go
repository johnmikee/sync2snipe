// Code generated by go-swagger; DO NOT EDIT.

package restrictedsoftware

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

// NewFindRestrictedSoftwareByIDParams creates a new FindRestrictedSoftwareByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindRestrictedSoftwareByIDParams() *FindRestrictedSoftwareByIDParams {
	return &FindRestrictedSoftwareByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindRestrictedSoftwareByIDParamsWithTimeout creates a new FindRestrictedSoftwareByIDParams object
// with the ability to set a timeout on a request.
func NewFindRestrictedSoftwareByIDParamsWithTimeout(timeout time.Duration) *FindRestrictedSoftwareByIDParams {
	return &FindRestrictedSoftwareByIDParams{
		timeout: timeout,
	}
}

// NewFindRestrictedSoftwareByIDParamsWithContext creates a new FindRestrictedSoftwareByIDParams object
// with the ability to set a context for a request.
func NewFindRestrictedSoftwareByIDParamsWithContext(ctx context.Context) *FindRestrictedSoftwareByIDParams {
	return &FindRestrictedSoftwareByIDParams{
		Context: ctx,
	}
}

// NewFindRestrictedSoftwareByIDParamsWithHTTPClient creates a new FindRestrictedSoftwareByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindRestrictedSoftwareByIDParamsWithHTTPClient(client *http.Client) *FindRestrictedSoftwareByIDParams {
	return &FindRestrictedSoftwareByIDParams{
		HTTPClient: client,
	}
}

/*
FindRestrictedSoftwareByIDParams contains all the parameters to send to the API endpoint

	for the find restricted software by Id operation.

	Typically these are written to a http.Request.
*/
type FindRestrictedSoftwareByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find restricted software by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindRestrictedSoftwareByIDParams) WithDefaults() *FindRestrictedSoftwareByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find restricted software by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindRestrictedSoftwareByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) WithTimeout(timeout time.Duration) *FindRestrictedSoftwareByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) WithContext(ctx context.Context) *FindRestrictedSoftwareByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) WithHTTPClient(client *http.Client) *FindRestrictedSoftwareByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) WithID(id int64) *FindRestrictedSoftwareByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find restricted software by Id params
func (o *FindRestrictedSoftwareByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindRestrictedSoftwareByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
