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
)

// NewFindMobileDevicesBySerialNumberParams creates a new FindMobileDevicesBySerialNumberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDevicesBySerialNumberParams() *FindMobileDevicesBySerialNumberParams {
	return &FindMobileDevicesBySerialNumberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDevicesBySerialNumberParamsWithTimeout creates a new FindMobileDevicesBySerialNumberParams object
// with the ability to set a timeout on a request.
func NewFindMobileDevicesBySerialNumberParamsWithTimeout(timeout time.Duration) *FindMobileDevicesBySerialNumberParams {
	return &FindMobileDevicesBySerialNumberParams{
		timeout: timeout,
	}
}

// NewFindMobileDevicesBySerialNumberParamsWithContext creates a new FindMobileDevicesBySerialNumberParams object
// with the ability to set a context for a request.
func NewFindMobileDevicesBySerialNumberParamsWithContext(ctx context.Context) *FindMobileDevicesBySerialNumberParams {
	return &FindMobileDevicesBySerialNumberParams{
		Context: ctx,
	}
}

// NewFindMobileDevicesBySerialNumberParamsWithHTTPClient creates a new FindMobileDevicesBySerialNumberParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDevicesBySerialNumberParamsWithHTTPClient(client *http.Client) *FindMobileDevicesBySerialNumberParams {
	return &FindMobileDevicesBySerialNumberParams{
		HTTPClient: client,
	}
}

/*
FindMobileDevicesBySerialNumberParams contains all the parameters to send to the API endpoint

	for the find mobile devices by serial number operation.

	Typically these are written to a http.Request.
*/
type FindMobileDevicesBySerialNumberParams struct {

	/* Serialnumber.

	   Serial number to filter by
	*/
	Serialnumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile devices by serial number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesBySerialNumberParams) WithDefaults() *FindMobileDevicesBySerialNumberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile devices by serial number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesBySerialNumberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) WithTimeout(timeout time.Duration) *FindMobileDevicesBySerialNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) WithContext(ctx context.Context) *FindMobileDevicesBySerialNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) WithHTTPClient(client *http.Client) *FindMobileDevicesBySerialNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerialnumber adds the serialnumber to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) WithSerialnumber(serialnumber string) *FindMobileDevicesBySerialNumberParams {
	o.SetSerialnumber(serialnumber)
	return o
}

// SetSerialnumber adds the serialnumber to the find mobile devices by serial number params
func (o *FindMobileDevicesBySerialNumberParams) SetSerialnumber(serialnumber string) {
	o.Serialnumber = serialnumber
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDevicesBySerialNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serialnumber
	if err := r.SetPathParam("serialnumber", o.Serialnumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
