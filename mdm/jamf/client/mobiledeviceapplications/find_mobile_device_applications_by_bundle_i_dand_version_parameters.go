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

// NewFindMobileDeviceApplicationsByBundleIDandVersionParams creates a new FindMobileDeviceApplicationsByBundleIDandVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceApplicationsByBundleIDandVersionParams() *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	return &FindMobileDeviceApplicationsByBundleIDandVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithTimeout creates a new FindMobileDeviceApplicationsByBundleIDandVersionParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	return &FindMobileDeviceApplicationsByBundleIDandVersionParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithContext creates a new FindMobileDeviceApplicationsByBundleIDandVersionParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithContext(ctx context.Context) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	return &FindMobileDeviceApplicationsByBundleIDandVersionParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithHTTPClient creates a new FindMobileDeviceApplicationsByBundleIDandVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceApplicationsByBundleIDandVersionParamsWithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	return &FindMobileDeviceApplicationsByBundleIDandVersionParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceApplicationsByBundleIDandVersionParams contains all the parameters to send to the API endpoint

	for the find mobile device applications by bundle i dand version operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceApplicationsByBundleIDandVersionParams struct {

	/* Bundleid.

	   Bundle ID to filter by
	*/
	Bundleid string

	/* Version.

	   Version to filter by
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device applications by bundle i dand version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithDefaults() *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device applications by bundle i dand version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithContext(ctx context.Context) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleid adds the bundleid to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithBundleid(bundleid string) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetBundleid(bundleid)
	return o
}

// SetBundleid adds the bundleid to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetBundleid(bundleid string) {
	o.Bundleid = bundleid
}

// WithVersion adds the version to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WithVersion(version string) *FindMobileDeviceApplicationsByBundleIDandVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the find mobile device applications by bundle i dand version params
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceApplicationsByBundleIDandVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundleid
	if err := r.SetPathParam("bundleid", o.Bundleid); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
