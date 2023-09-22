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

// NewFindMobileDeviceCommandsParams creates a new FindMobileDeviceCommandsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceCommandsParams() *FindMobileDeviceCommandsParams {
	return &FindMobileDeviceCommandsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceCommandsParamsWithTimeout creates a new FindMobileDeviceCommandsParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceCommandsParamsWithTimeout(timeout time.Duration) *FindMobileDeviceCommandsParams {
	return &FindMobileDeviceCommandsParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceCommandsParamsWithContext creates a new FindMobileDeviceCommandsParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceCommandsParamsWithContext(ctx context.Context) *FindMobileDeviceCommandsParams {
	return &FindMobileDeviceCommandsParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceCommandsParamsWithHTTPClient creates a new FindMobileDeviceCommandsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceCommandsParamsWithHTTPClient(client *http.Client) *FindMobileDeviceCommandsParams {
	return &FindMobileDeviceCommandsParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceCommandsParams contains all the parameters to send to the API endpoint

	for the find mobile device commands operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceCommandsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceCommandsParams) WithDefaults() *FindMobileDeviceCommandsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceCommandsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) WithTimeout(timeout time.Duration) *FindMobileDeviceCommandsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) WithContext(ctx context.Context) *FindMobileDeviceCommandsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) WithHTTPClient(client *http.Client) *FindMobileDeviceCommandsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device commands params
func (o *FindMobileDeviceCommandsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceCommandsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
