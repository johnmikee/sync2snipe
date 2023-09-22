// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceconfigurationprofiles

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

// NewDeleteMobileDeviceConfigurationProfileByNameParams creates a new DeleteMobileDeviceConfigurationProfileByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMobileDeviceConfigurationProfileByNameParams() *DeleteMobileDeviceConfigurationProfileByNameParams {
	return &DeleteMobileDeviceConfigurationProfileByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByNameParamsWithTimeout creates a new DeleteMobileDeviceConfigurationProfileByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteMobileDeviceConfigurationProfileByNameParamsWithTimeout(timeout time.Duration) *DeleteMobileDeviceConfigurationProfileByNameParams {
	return &DeleteMobileDeviceConfigurationProfileByNameParams{
		timeout: timeout,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByNameParamsWithContext creates a new DeleteMobileDeviceConfigurationProfileByNameParams object
// with the ability to set a context for a request.
func NewDeleteMobileDeviceConfigurationProfileByNameParamsWithContext(ctx context.Context) *DeleteMobileDeviceConfigurationProfileByNameParams {
	return &DeleteMobileDeviceConfigurationProfileByNameParams{
		Context: ctx,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByNameParamsWithHTTPClient creates a new DeleteMobileDeviceConfigurationProfileByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMobileDeviceConfigurationProfileByNameParamsWithHTTPClient(client *http.Client) *DeleteMobileDeviceConfigurationProfileByNameParams {
	return &DeleteMobileDeviceConfigurationProfileByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteMobileDeviceConfigurationProfileByNameParams contains all the parameters to send to the API endpoint

	for the delete mobile device configuration profile by name operation.

	Typically these are written to a http.Request.
*/
type DeleteMobileDeviceConfigurationProfileByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mobile device configuration profile by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WithDefaults() *DeleteMobileDeviceConfigurationProfileByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mobile device configuration profile by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WithTimeout(timeout time.Duration) *DeleteMobileDeviceConfigurationProfileByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WithContext(ctx context.Context) *DeleteMobileDeviceConfigurationProfileByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WithHTTPClient(client *http.Client) *DeleteMobileDeviceConfigurationProfileByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WithName(name string) *DeleteMobileDeviceConfigurationProfileByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete mobile device configuration profile by name params
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobileDeviceConfigurationProfileByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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