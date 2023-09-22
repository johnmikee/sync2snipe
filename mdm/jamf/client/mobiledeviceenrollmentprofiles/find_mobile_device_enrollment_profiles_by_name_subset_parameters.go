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

// NewFindMobileDeviceEnrollmentProfilesByNameSubsetParams creates a new FindMobileDeviceEnrollmentProfilesByNameSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceEnrollmentProfilesByNameSubsetParams() *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	return &FindMobileDeviceEnrollmentProfilesByNameSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithTimeout creates a new FindMobileDeviceEnrollmentProfilesByNameSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	return &FindMobileDeviceEnrollmentProfilesByNameSubsetParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithContext creates a new FindMobileDeviceEnrollmentProfilesByNameSubsetParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	return &FindMobileDeviceEnrollmentProfilesByNameSubsetParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithHTTPClient creates a new FindMobileDeviceEnrollmentProfilesByNameSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceEnrollmentProfilesByNameSubsetParamsWithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	return &FindMobileDeviceEnrollmentProfilesByNameSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceEnrollmentProfilesByNameSubsetParams contains all the parameters to send to the API endpoint

	for the find mobile device enrollment profiles by name subset operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceEnrollmentProfilesByNameSubsetParams struct {

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

// WithDefaults hydrates default values in the find mobile device enrollment profiles by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithDefaults() *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device enrollment profiles by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithName(name string) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetName(name string) {
	o.Name = name
}

// WithSubset adds the subset to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WithSubset(subset string) *FindMobileDeviceEnrollmentProfilesByNameSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find mobile device enrollment profiles by name subset params
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceEnrollmentProfilesByNameSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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