// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceprovisioningprofiles

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

// NewFindMobileDeviceProvisioningProfilesByNameParams creates a new FindMobileDeviceProvisioningProfilesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceProvisioningProfilesByNameParams() *FindMobileDeviceProvisioningProfilesByNameParams {
	return &FindMobileDeviceProvisioningProfilesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceProvisioningProfilesByNameParamsWithTimeout creates a new FindMobileDeviceProvisioningProfilesByNameParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceProvisioningProfilesByNameParamsWithTimeout(timeout time.Duration) *FindMobileDeviceProvisioningProfilesByNameParams {
	return &FindMobileDeviceProvisioningProfilesByNameParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceProvisioningProfilesByNameParamsWithContext creates a new FindMobileDeviceProvisioningProfilesByNameParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceProvisioningProfilesByNameParamsWithContext(ctx context.Context) *FindMobileDeviceProvisioningProfilesByNameParams {
	return &FindMobileDeviceProvisioningProfilesByNameParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceProvisioningProfilesByNameParamsWithHTTPClient creates a new FindMobileDeviceProvisioningProfilesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceProvisioningProfilesByNameParamsWithHTTPClient(client *http.Client) *FindMobileDeviceProvisioningProfilesByNameParams {
	return &FindMobileDeviceProvisioningProfilesByNameParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceProvisioningProfilesByNameParams contains all the parameters to send to the API endpoint

	for the find mobile device provisioning profiles by name operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceProvisioningProfilesByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device provisioning profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WithDefaults() *FindMobileDeviceProvisioningProfilesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device provisioning profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceProvisioningProfilesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WithTimeout(timeout time.Duration) *FindMobileDeviceProvisioningProfilesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WithContext(ctx context.Context) *FindMobileDeviceProvisioningProfilesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WithHTTPClient(client *http.Client) *FindMobileDeviceProvisioningProfilesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WithName(name string) *FindMobileDeviceProvisioningProfilesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find mobile device provisioning profiles by name params
func (o *FindMobileDeviceProvisioningProfilesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceProvisioningProfilesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
