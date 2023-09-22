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

// NewFindMobileDevicesByMacAddressSubsetParams creates a new FindMobileDevicesByMacAddressSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDevicesByMacAddressSubsetParams() *FindMobileDevicesByMacAddressSubsetParams {
	return &FindMobileDevicesByMacAddressSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDevicesByMacAddressSubsetParamsWithTimeout creates a new FindMobileDevicesByMacAddressSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMobileDevicesByMacAddressSubsetParamsWithTimeout(timeout time.Duration) *FindMobileDevicesByMacAddressSubsetParams {
	return &FindMobileDevicesByMacAddressSubsetParams{
		timeout: timeout,
	}
}

// NewFindMobileDevicesByMacAddressSubsetParamsWithContext creates a new FindMobileDevicesByMacAddressSubsetParams object
// with the ability to set a context for a request.
func NewFindMobileDevicesByMacAddressSubsetParamsWithContext(ctx context.Context) *FindMobileDevicesByMacAddressSubsetParams {
	return &FindMobileDevicesByMacAddressSubsetParams{
		Context: ctx,
	}
}

// NewFindMobileDevicesByMacAddressSubsetParamsWithHTTPClient creates a new FindMobileDevicesByMacAddressSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDevicesByMacAddressSubsetParamsWithHTTPClient(client *http.Client) *FindMobileDevicesByMacAddressSubsetParams {
	return &FindMobileDevicesByMacAddressSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMobileDevicesByMacAddressSubsetParams contains all the parameters to send to the API endpoint

	for the find mobile devices by mac address subset operation.

	Typically these are written to a http.Request.
*/
type FindMobileDevicesByMacAddressSubsetParams struct {

	/* Macaddress.

	   Mac address to filter by
	*/
	Macaddress string

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile devices by mac address subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByMacAddressSubsetParams) WithDefaults() *FindMobileDevicesByMacAddressSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile devices by mac address subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDevicesByMacAddressSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) WithTimeout(timeout time.Duration) *FindMobileDevicesByMacAddressSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) WithContext(ctx context.Context) *FindMobileDevicesByMacAddressSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) WithHTTPClient(client *http.Client) *FindMobileDevicesByMacAddressSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacaddress adds the macaddress to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) WithMacaddress(macaddress string) *FindMobileDevicesByMacAddressSubsetParams {
	o.SetMacaddress(macaddress)
	return o
}

// SetMacaddress adds the macaddress to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) SetMacaddress(macaddress string) {
	o.Macaddress = macaddress
}

// WithSubset adds the subset to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) WithSubset(subset string) *FindMobileDevicesByMacAddressSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find mobile devices by mac address subset params
func (o *FindMobileDevicesByMacAddressSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDevicesByMacAddressSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param macaddress
	if err := r.SetPathParam("macaddress", o.Macaddress); err != nil {
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
