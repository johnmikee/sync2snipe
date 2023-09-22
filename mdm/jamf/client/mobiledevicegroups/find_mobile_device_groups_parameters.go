// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicegroups

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

// NewFindMobileDeviceGroupsParams creates a new FindMobileDeviceGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceGroupsParams() *FindMobileDeviceGroupsParams {
	return &FindMobileDeviceGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceGroupsParamsWithTimeout creates a new FindMobileDeviceGroupsParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceGroupsParamsWithTimeout(timeout time.Duration) *FindMobileDeviceGroupsParams {
	return &FindMobileDeviceGroupsParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceGroupsParamsWithContext creates a new FindMobileDeviceGroupsParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceGroupsParamsWithContext(ctx context.Context) *FindMobileDeviceGroupsParams {
	return &FindMobileDeviceGroupsParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceGroupsParamsWithHTTPClient creates a new FindMobileDeviceGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceGroupsParamsWithHTTPClient(client *http.Client) *FindMobileDeviceGroupsParams {
	return &FindMobileDeviceGroupsParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceGroupsParams contains all the parameters to send to the API endpoint

	for the find mobile device groups operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceGroupsParams) WithDefaults() *FindMobileDeviceGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) WithTimeout(timeout time.Duration) *FindMobileDeviceGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) WithContext(ctx context.Context) *FindMobileDeviceGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) WithHTTPClient(client *http.Client) *FindMobileDeviceGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device groups params
func (o *FindMobileDeviceGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}