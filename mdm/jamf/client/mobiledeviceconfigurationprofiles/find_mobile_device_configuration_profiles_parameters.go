// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceconfigurationprofiles

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

// NewFindMobileDeviceConfigurationProfilesParams creates a new FindMobileDeviceConfigurationProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceConfigurationProfilesParams() *FindMobileDeviceConfigurationProfilesParams {
	return &FindMobileDeviceConfigurationProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceConfigurationProfilesParamsWithTimeout creates a new FindMobileDeviceConfigurationProfilesParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceConfigurationProfilesParamsWithTimeout(timeout time.Duration) *FindMobileDeviceConfigurationProfilesParams {
	return &FindMobileDeviceConfigurationProfilesParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceConfigurationProfilesParamsWithContext creates a new FindMobileDeviceConfigurationProfilesParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceConfigurationProfilesParamsWithContext(ctx context.Context) *FindMobileDeviceConfigurationProfilesParams {
	return &FindMobileDeviceConfigurationProfilesParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceConfigurationProfilesParamsWithHTTPClient creates a new FindMobileDeviceConfigurationProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceConfigurationProfilesParamsWithHTTPClient(client *http.Client) *FindMobileDeviceConfigurationProfilesParams {
	return &FindMobileDeviceConfigurationProfilesParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceConfigurationProfilesParams contains all the parameters to send to the API endpoint

	for the find mobile device configuration profiles operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceConfigurationProfilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device configuration profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceConfigurationProfilesParams) WithDefaults() *FindMobileDeviceConfigurationProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device configuration profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceConfigurationProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) WithTimeout(timeout time.Duration) *FindMobileDeviceConfigurationProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) WithContext(ctx context.Context) *FindMobileDeviceConfigurationProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) WithHTTPClient(client *http.Client) *FindMobileDeviceConfigurationProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device configuration profiles params
func (o *FindMobileDeviceConfigurationProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceConfigurationProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
