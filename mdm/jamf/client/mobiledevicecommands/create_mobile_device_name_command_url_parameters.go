// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicecommands

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

// NewCreateMobileDeviceNameCommandURLParams creates a new CreateMobileDeviceNameCommandURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMobileDeviceNameCommandURLParams() *CreateMobileDeviceNameCommandURLParams {
	return &CreateMobileDeviceNameCommandURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMobileDeviceNameCommandURLParamsWithTimeout creates a new CreateMobileDeviceNameCommandURLParams object
// with the ability to set a timeout on a request.
func NewCreateMobileDeviceNameCommandURLParamsWithTimeout(timeout time.Duration) *CreateMobileDeviceNameCommandURLParams {
	return &CreateMobileDeviceNameCommandURLParams{
		timeout: timeout,
	}
}

// NewCreateMobileDeviceNameCommandURLParamsWithContext creates a new CreateMobileDeviceNameCommandURLParams object
// with the ability to set a context for a request.
func NewCreateMobileDeviceNameCommandURLParamsWithContext(ctx context.Context) *CreateMobileDeviceNameCommandURLParams {
	return &CreateMobileDeviceNameCommandURLParams{
		Context: ctx,
	}
}

// NewCreateMobileDeviceNameCommandURLParamsWithHTTPClient creates a new CreateMobileDeviceNameCommandURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMobileDeviceNameCommandURLParamsWithHTTPClient(client *http.Client) *CreateMobileDeviceNameCommandURLParams {
	return &CreateMobileDeviceNameCommandURLParams{
		HTTPClient: client,
	}
}

/*
CreateMobileDeviceNameCommandURLParams contains all the parameters to send to the API endpoint

	for the create mobile device name command URL operation.

	Typically these are written to a http.Request.
*/
type CreateMobileDeviceNameCommandURLParams struct {

	/* DeviceName.

	   Device name to set
	*/
	DeviceName string

	/* IDList.

	   Mobile device ID values, multiple IDs may be separated by commas (e.g. /id/13,14,15)
	*/
	IDList string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create mobile device name command URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMobileDeviceNameCommandURLParams) WithDefaults() *CreateMobileDeviceNameCommandURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create mobile device name command URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMobileDeviceNameCommandURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) WithTimeout(timeout time.Duration) *CreateMobileDeviceNameCommandURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) WithContext(ctx context.Context) *CreateMobileDeviceNameCommandURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) WithHTTPClient(client *http.Client) *CreateMobileDeviceNameCommandURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceName adds the deviceName to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) WithDeviceName(deviceName string) *CreateMobileDeviceNameCommandURLParams {
	o.SetDeviceName(deviceName)
	return o
}

// SetDeviceName adds the deviceName to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) SetDeviceName(deviceName string) {
	o.DeviceName = deviceName
}

// WithIDList adds the iDList to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) WithIDList(iDList string) *CreateMobileDeviceNameCommandURLParams {
	o.SetIDList(iDList)
	return o
}

// SetIDList adds the idList to the create mobile device name command URL params
func (o *CreateMobileDeviceNameCommandURLParams) SetIDList(iDList string) {
	o.IDList = iDList
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMobileDeviceNameCommandURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_name
	if err := r.SetPathParam("device_name", o.DeviceName); err != nil {
		return err
	}

	// path param id_list
	if err := r.SetPathParam("id_list", o.IDList); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
