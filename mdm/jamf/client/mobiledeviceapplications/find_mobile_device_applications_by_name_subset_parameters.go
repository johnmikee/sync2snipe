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

// NewFindMobileDeviceApplicationsByNameSubsetParams creates a new FindMobileDeviceApplicationsByNameSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceApplicationsByNameSubsetParams() *FindMobileDeviceApplicationsByNameSubsetParams {
	return &FindMobileDeviceApplicationsByNameSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceApplicationsByNameSubsetParamsWithTimeout creates a new FindMobileDeviceApplicationsByNameSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceApplicationsByNameSubsetParamsWithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByNameSubsetParams {
	return &FindMobileDeviceApplicationsByNameSubsetParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceApplicationsByNameSubsetParamsWithContext creates a new FindMobileDeviceApplicationsByNameSubsetParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceApplicationsByNameSubsetParamsWithContext(ctx context.Context) *FindMobileDeviceApplicationsByNameSubsetParams {
	return &FindMobileDeviceApplicationsByNameSubsetParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceApplicationsByNameSubsetParamsWithHTTPClient creates a new FindMobileDeviceApplicationsByNameSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceApplicationsByNameSubsetParamsWithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByNameSubsetParams {
	return &FindMobileDeviceApplicationsByNameSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceApplicationsByNameSubsetParams contains all the parameters to send to the API endpoint

	for the find mobile device applications by name subset operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceApplicationsByNameSubsetParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device applications by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithDefaults() *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device applications by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithTimeout(timeout time.Duration) *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithContext(ctx context.Context) *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithHTTPClient(client *http.Client) *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithName(name string) *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetName(name string) {
	o.Name = name
}

// WithSubset adds the subset to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WithSubset(subset string) *FindMobileDeviceApplicationsByNameSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find mobile device applications by name subset params
func (o *FindMobileDeviceApplicationsByNameSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceApplicationsByNameSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param subset
	if err := r.SetPathParam("subset", o.Subset); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
