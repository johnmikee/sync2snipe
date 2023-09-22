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
)

// NewFindMobileDeviceEnrollmentProfilesParams creates a new FindMobileDeviceEnrollmentProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceEnrollmentProfilesParams() *FindMobileDeviceEnrollmentProfilesParams {
	return &FindMobileDeviceEnrollmentProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesParamsWithTimeout creates a new FindMobileDeviceEnrollmentProfilesParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceEnrollmentProfilesParamsWithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesParams {
	return &FindMobileDeviceEnrollmentProfilesParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesParamsWithContext creates a new FindMobileDeviceEnrollmentProfilesParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceEnrollmentProfilesParamsWithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesParams {
	return &FindMobileDeviceEnrollmentProfilesParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceEnrollmentProfilesParamsWithHTTPClient creates a new FindMobileDeviceEnrollmentProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceEnrollmentProfilesParamsWithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesParams {
	return &FindMobileDeviceEnrollmentProfilesParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceEnrollmentProfilesParams contains all the parameters to send to the API endpoint

	for the find mobile device enrollment profiles operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceEnrollmentProfilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device enrollment profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesParams) WithDefaults() *FindMobileDeviceEnrollmentProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device enrollment profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) WithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) WithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) WithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device enrollment profiles params
func (o *FindMobileDeviceEnrollmentProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceEnrollmentProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}