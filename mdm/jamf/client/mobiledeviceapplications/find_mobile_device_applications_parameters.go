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
)

// NewFindMobileDeviceApplicationsParams creates a new FindMobileDeviceApplicationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceApplicationsParams() *FindMobileDeviceApplicationsParams {
	return &FindMobileDeviceApplicationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceApplicationsParamsWithTimeout creates a new FindMobileDeviceApplicationsParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceApplicationsParamsWithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsParams {
	return &FindMobileDeviceApplicationsParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceApplicationsParamsWithContext creates a new FindMobileDeviceApplicationsParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceApplicationsParamsWithContext(ctx context.Context) *FindMobileDeviceApplicationsParams {
	return &FindMobileDeviceApplicationsParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceApplicationsParamsWithHTTPClient creates a new FindMobileDeviceApplicationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceApplicationsParamsWithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsParams {
	return &FindMobileDeviceApplicationsParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceApplicationsParams contains all the parameters to send to the API endpoint

	for the find mobile device applications operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceApplicationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsParams) WithDefaults() *FindMobileDeviceApplicationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) WithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) WithContext(ctx context.Context) *FindMobileDeviceApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) WithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device applications params
func (o *FindMobileDeviceApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
