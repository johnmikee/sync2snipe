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

// NewFindMobileDevicesByIDSubsetParams creates a new FindMobileDevicesByIDSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDevicesByIDSubsetParams() *FindMobileDevicesByIDSubsetParams {
	return &FindMobileDevicesByIDSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDevicesByIDSubsetParamsWithTimeout creates a new FindMobileDevicesByIDSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMobileDevicesByIDSubsetParamsWithTimeout(timeout time.Duration) *FindMobileDevicesByIDSubsetParams {
	return &FindMobileDevicesByIDSubsetParams{
		timeout: timeout,
	}
}

// NewFindMobileDevicesByIDSubsetParamsWithContext creates a new FindMobileDevicesByIDSubsetParams object
// with the ability to set a context for a request.
func NewFindMobileDevicesByIDSubsetParamsWithContext(ctx context.Context) *FindMobileDevicesByIDSubsetParams {
	return &FindMobileDevicesByIDSubsetParams{
		Context: ctx,
	}
}

// NewFindMobileDevicesByIDSubsetParamsWithHTTPClient creates a new FindMobileDevicesByIDSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDevicesByIDSubsetParamsWithHTTPClient(client *http.Client) *FindMobileDevicesByIDSubsetParams {
	return &FindMobileDevicesByIDSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMobileDevicesByIDSubsetParams contains all the parameters to send to the API endpoint

	for the find mobile devices by Id subset operation.

	Typically these are written to a http.Request.
*/
type FindMobileDevicesByIDSubsetParams struct {

	/* ID.

	   ID to filter by
	*/
	ID int64

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile devices by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByIDSubsetParams) WithDefaults() *FindMobileDevicesByIDSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile devices by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByIDSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) WithTimeout(timeout time.Duration) *FindMobileDevicesByIDSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) WithContext(ctx context.Context) *FindMobileDevicesByIDSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) WithHTTPClient(client *http.Client) *FindMobileDevicesByIDSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) WithID(id int64) *FindMobileDevicesByIDSubsetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) SetID(id int64) {
	o.ID = id
}

// WithSubset adds the subset to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) WithSubset(subset string) *FindMobileDevicesByIDSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find mobile devices by Id subset params
func (o *FindMobileDevicesByIDSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDevicesByIDSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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