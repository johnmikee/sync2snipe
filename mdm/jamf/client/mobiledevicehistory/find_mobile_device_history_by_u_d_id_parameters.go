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

// NewFindMobileDeviceHistoryByUDIDParams creates a new FindMobileDeviceHistoryByUDIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceHistoryByUDIDParams() *FindMobileDeviceHistoryByUDIDParams {
	return &FindMobileDeviceHistoryByUDIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceHistoryByUDIDParamsWithTimeout creates a new FindMobileDeviceHistoryByUDIDParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceHistoryByUDIDParamsWithTimeout(timeout time.Duration) *FindMobileDeviceHistoryByUDIDParams {
	return &FindMobileDeviceHistoryByUDIDParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceHistoryByUDIDParamsWithContext creates a new FindMobileDeviceHistoryByUDIDParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceHistoryByUDIDParamsWithContext(ctx context.Context) *FindMobileDeviceHistoryByUDIDParams {
	return &FindMobileDeviceHistoryByUDIDParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceHistoryByUDIDParamsWithHTTPClient creates a new FindMobileDeviceHistoryByUDIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceHistoryByUDIDParamsWithHTTPClient(client *http.Client) *FindMobileDeviceHistoryByUDIDParams {
	return &FindMobileDeviceHistoryByUDIDParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceHistoryByUDIDParams contains all the parameters to send to the API endpoint

	for the find mobile device history by u d ID operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceHistoryByUDIDParams struct {

	/* Udid.

	   UDID to filter by
	*/
	Udid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device history by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceHistoryByUDIDParams) WithDefaults() *FindMobileDeviceHistoryByUDIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device history by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceHistoryByUDIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) WithTimeout(timeout time.Duration) *FindMobileDeviceHistoryByUDIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) WithContext(ctx context.Context) *FindMobileDeviceHistoryByUDIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) WithHTTPClient(client *http.Client) *FindMobileDeviceHistoryByUDIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUdid adds the udid to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) WithUdid(udid string) *FindMobileDeviceHistoryByUDIDParams {
	o.SetUdid(udid)
	return o
}

// SetUdid adds the udid to the find mobile device history by u d ID params
func (o *FindMobileDeviceHistoryByUDIDParams) SetUdid(udid string) {
	o.Udid = udid
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceHistoryByUDIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param udid
	if err := r.SetPathParam("udid", o.Udid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}