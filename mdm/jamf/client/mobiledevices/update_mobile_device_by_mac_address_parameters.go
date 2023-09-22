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

// NewUpdateMobileDeviceByMacAddressParams creates a new UpdateMobileDeviceByMacAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMobileDeviceByMacAddressParams() *UpdateMobileDeviceByMacAddressParams {
	return &UpdateMobileDeviceByMacAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMobileDeviceByMacAddressParamsWithTimeout creates a new UpdateMobileDeviceByMacAddressParams object
// with the ability to set a timeout on a request.
func NewUpdateMobileDeviceByMacAddressParamsWithTimeout(timeout time.Duration) *UpdateMobileDeviceByMacAddressParams {
	return &UpdateMobileDeviceByMacAddressParams{
		timeout: timeout,
	}
}

// NewUpdateMobileDeviceByMacAddressParamsWithContext creates a new UpdateMobileDeviceByMacAddressParams object
// with the ability to set a context for a request.
func NewUpdateMobileDeviceByMacAddressParamsWithContext(ctx context.Context) *UpdateMobileDeviceByMacAddressParams {
	return &UpdateMobileDeviceByMacAddressParams{
		Context: ctx,
	}
}

// NewUpdateMobileDeviceByMacAddressParamsWithHTTPClient creates a new UpdateMobileDeviceByMacAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMobileDeviceByMacAddressParamsWithHTTPClient(client *http.Client) *UpdateMobileDeviceByMacAddressParams {
	return &UpdateMobileDeviceByMacAddressParams{
		HTTPClient: client,
	}
}

/*
UpdateMobileDeviceByMacAddressParams contains all the parameters to send to the API endpoint

	for the update mobile device by mac address operation.

	Typically these are written to a http.Request.
*/
type UpdateMobileDeviceByMacAddressParams struct {

	/* Macaddress.

	   Mac address value to filter by
	*/
	Macaddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mobile device by mac address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceByMacAddressParams) WithDefaults() *UpdateMobileDeviceByMacAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mobile device by mac address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceByMacAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) WithTimeout(timeout time.Duration) *UpdateMobileDeviceByMacAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) WithContext(ctx context.Context) *UpdateMobileDeviceByMacAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) WithHTTPClient(client *http.Client) *UpdateMobileDeviceByMacAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacaddress adds the macaddress to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) WithMacaddress(macaddress string) *UpdateMobileDeviceByMacAddressParams {
	o.SetMacaddress(macaddress)
	return o
}

// SetMacaddress adds the macaddress to the update mobile device by mac address params
func (o *UpdateMobileDeviceByMacAddressParams) SetMacaddress(macaddress string) {
	o.Macaddress = macaddress
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMobileDeviceByMacAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param macaddress
	if err := r.SetPathParam("macaddress", o.Macaddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}