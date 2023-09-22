// Code generated by go-swagger; DO NOT EDIT.

package mobiledevices

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

// NewFindMobileDevicesByIDParams creates a new FindMobileDevicesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDevicesByIDParams() *FindMobileDevicesByIDParams {
	return &FindMobileDevicesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDevicesByIDParamsWithTimeout creates a new FindMobileDevicesByIDParams object
// with the ability to set a timeout on a request.
func NewFindMobileDevicesByIDParamsWithTimeout(timeout time.Duration) *FindMobileDevicesByIDParams {
	return &FindMobileDevicesByIDParams{
		timeout: timeout,
	}
}

// NewFindMobileDevicesByIDParamsWithContext creates a new FindMobileDevicesByIDParams object
// with the ability to set a context for a request.
func NewFindMobileDevicesByIDParamsWithContext(ctx context.Context) *FindMobileDevicesByIDParams {
	return &FindMobileDevicesByIDParams{
		Context: ctx,
	}
}

// NewFindMobileDevicesByIDParamsWithHTTPClient creates a new FindMobileDevicesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDevicesByIDParamsWithHTTPClient(client *http.Client) *FindMobileDevicesByIDParams {
	return &FindMobileDevicesByIDParams{
		HTTPClient: client,
	}
}

/*
FindMobileDevicesByIDParams contains all the parameters to send to the API endpoint

	for the find mobile devices by Id operation.

	Typically these are written to a http.Request.
*/
type FindMobileDevicesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile devices by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByIDParams) WithDefaults() *FindMobileDevicesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile devices by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) WithTimeout(timeout time.Duration) *FindMobileDevicesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) WithContext(ctx context.Context) *FindMobileDevicesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) WithHTTPClient(client *http.Client) *FindMobileDevicesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) WithID(id int64) *FindMobileDevicesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find mobile devices by Id params
func (o *FindMobileDevicesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDevicesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
