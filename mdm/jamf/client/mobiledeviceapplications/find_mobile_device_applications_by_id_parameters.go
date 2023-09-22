// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceapplications

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

// NewFindMobileDeviceApplicationsByIDParams creates a new FindMobileDeviceApplicationsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceApplicationsByIDParams() *FindMobileDeviceApplicationsByIDParams {
	return &FindMobileDeviceApplicationsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceApplicationsByIDParamsWithTimeout creates a new FindMobileDeviceApplicationsByIDParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceApplicationsByIDParamsWithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByIDParams {
	return &FindMobileDeviceApplicationsByIDParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceApplicationsByIDParamsWithContext creates a new FindMobileDeviceApplicationsByIDParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceApplicationsByIDParamsWithContext(ctx context.Context) *FindMobileDeviceApplicationsByIDParams {
	return &FindMobileDeviceApplicationsByIDParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceApplicationsByIDParamsWithHTTPClient creates a new FindMobileDeviceApplicationsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceApplicationsByIDParamsWithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByIDParams {
	return &FindMobileDeviceApplicationsByIDParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceApplicationsByIDParams contains all the parameters to send to the API endpoint

	for the find mobile device applications by Id operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceApplicationsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device applications by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByIDParams) WithDefaults() *FindMobileDeviceApplicationsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device applications by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) WithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) WithContext(ctx context.Context) *FindMobileDeviceApplicationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) WithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) WithID(id int64) *FindMobileDeviceApplicationsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find mobile device applications by Id params
func (o *FindMobileDeviceApplicationsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceApplicationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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