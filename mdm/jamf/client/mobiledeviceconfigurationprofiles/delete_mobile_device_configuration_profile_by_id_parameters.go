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
	"github.com/go-openapi/swag"
)

// NewDeleteMobileDeviceConfigurationProfileByIDParams creates a new DeleteMobileDeviceConfigurationProfileByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMobileDeviceConfigurationProfileByIDParams() *DeleteMobileDeviceConfigurationProfileByIDParams {
	return &DeleteMobileDeviceConfigurationProfileByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByIDParamsWithTimeout creates a new DeleteMobileDeviceConfigurationProfileByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteMobileDeviceConfigurationProfileByIDParamsWithTimeout(timeout time.Duration) *DeleteMobileDeviceConfigurationProfileByIDParams {
	return &DeleteMobileDeviceConfigurationProfileByIDParams{
		timeout: timeout,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByIDParamsWithContext creates a new DeleteMobileDeviceConfigurationProfileByIDParams object
// with the ability to set a context for a request.
func NewDeleteMobileDeviceConfigurationProfileByIDParamsWithContext(ctx context.Context) *DeleteMobileDeviceConfigurationProfileByIDParams {
	return &DeleteMobileDeviceConfigurationProfileByIDParams{
		Context: ctx,
	}
}

// NewDeleteMobileDeviceConfigurationProfileByIDParamsWithHTTPClient creates a new DeleteMobileDeviceConfigurationProfileByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMobileDeviceConfigurationProfileByIDParamsWithHTTPClient(client *http.Client) *DeleteMobileDeviceConfigurationProfileByIDParams {
	return &DeleteMobileDeviceConfigurationProfileByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteMobileDeviceConfigurationProfileByIDParams contains all the parameters to send to the API endpoint

	for the delete mobile device configuration profile by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteMobileDeviceConfigurationProfileByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mobile device configuration profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WithDefaults() *DeleteMobileDeviceConfigurationProfileByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mobile device configuration profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WithTimeout(timeout time.Duration) *DeleteMobileDeviceConfigurationProfileByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WithContext(ctx context.Context) *DeleteMobileDeviceConfigurationProfileByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WithHTTPClient(client *http.Client) *DeleteMobileDeviceConfigurationProfileByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WithID(id int64) *DeleteMobileDeviceConfigurationProfileByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete mobile device configuration profile by Id params
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobileDeviceConfigurationProfileByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
