// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

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

// NewFindMobileDeviceEnrollmentProfilesByIDParams creates a new FindMobileDeviceEnrollmentProfilesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceEnrollmentProfilesByIDParams() *FindMobileDeviceEnrollmentProfilesByIDParams {
	return &FindMobileDeviceEnrollmentProfilesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByIDParamsWithTimeout creates a new FindMobileDeviceEnrollmentProfilesByIDParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceEnrollmentProfilesByIDParamsWithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByIDParams {
	return &FindMobileDeviceEnrollmentProfilesByIDParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByIDParamsWithContext creates a new FindMobileDeviceEnrollmentProfilesByIDParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceEnrollmentProfilesByIDParamsWithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByIDParams {
	return &FindMobileDeviceEnrollmentProfilesByIDParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByIDParamsWithHTTPClient creates a new FindMobileDeviceEnrollmentProfilesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceEnrollmentProfilesByIDParamsWithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByIDParams {
	return &FindMobileDeviceEnrollmentProfilesByIDParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceEnrollmentProfilesByIDParams contains all the parameters to send to the API endpoint

	for the find mobile device enrollment profiles by Id operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceEnrollmentProfilesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device enrollment profiles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WithDefaults() *FindMobileDeviceEnrollmentProfilesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device enrollment profiles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WithID(id int64) *FindMobileDeviceEnrollmentProfilesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find mobile device enrollment profiles by Id params
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceEnrollmentProfilesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
