// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicehistory

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

// NewFindMobileDeviceHistoryByNameSubsetParams creates a new FindMobileDeviceHistoryByNameSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceHistoryByNameSubsetParams() *FindMobileDeviceHistoryByNameSubsetParams {
	return &FindMobileDeviceHistoryByNameSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceHistoryByNameSubsetParamsWithTimeout creates a new FindMobileDeviceHistoryByNameSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceHistoryByNameSubsetParamsWithTimeout(timeout time.Duration) *FindMobileDeviceHistoryByNameSubsetParams {
	return &FindMobileDeviceHistoryByNameSubsetParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceHistoryByNameSubsetParamsWithContext creates a new FindMobileDeviceHistoryByNameSubsetParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceHistoryByNameSubsetParamsWithContext(ctx context.Context) *FindMobileDeviceHistoryByNameSubsetParams {
	return &FindMobileDeviceHistoryByNameSubsetParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceHistoryByNameSubsetParamsWithHTTPClient creates a new FindMobileDeviceHistoryByNameSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceHistoryByNameSubsetParamsWithHTTPClient(client *http.Client) *FindMobileDeviceHistoryByNameSubsetParams {
	return &FindMobileDeviceHistoryByNameSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceHistoryByNameSubsetParams contains all the parameters to send to the API endpoint

	for the find mobile device history by name subset operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceHistoryByNameSubsetParams struct {

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

// WithDefaults hydrates default values in the find mobile device history by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithDefaults() *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device history by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithTimeout(timeout time.Duration) *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithContext(ctx context.Context) *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithHTTPClient(client *http.Client) *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithName(name string) *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetName(name string) {
	o.Name = name
}

// WithSubset adds the subset to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) WithSubset(subset string) *FindMobileDeviceHistoryByNameSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find mobile device history by name subset params
func (o *FindMobileDeviceHistoryByNameSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceHistoryByNameSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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