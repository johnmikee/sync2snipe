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

// NewDeleteMobileDeviceProvisioningProfilesByUUIDParams creates a new DeleteMobileDeviceProvisioningProfilesByUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMobileDeviceProvisioningProfilesByUUIDParams() *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	return &DeleteMobileDeviceProvisioningProfilesByUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithTimeout creates a new DeleteMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a timeout on a request.
func NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithTimeout(timeout time.Duration) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	return &DeleteMobileDeviceProvisioningProfilesByUUIDParams{
		timeout: timeout,
	}
}

// NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithContext creates a new DeleteMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a context for a request.
func NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithContext(ctx context.Context) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	return &DeleteMobileDeviceProvisioningProfilesByUUIDParams{
		Context: ctx,
	}
}

// NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithHTTPClient creates a new DeleteMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMobileDeviceProvisioningProfilesByUUIDParamsWithHTTPClient(client *http.Client) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	return &DeleteMobileDeviceProvisioningProfilesByUUIDParams{
		HTTPClient: client,
	}
}

/*
DeleteMobileDeviceProvisioningProfilesByUUIDParams contains all the parameters to send to the API endpoint

	for the delete mobile device provisioning profiles by UUID operation.

	Typically these are written to a http.Request.
*/
type DeleteMobileDeviceProvisioningProfilesByUUIDParams struct {

	/* UUID.

	   Uuid value to filter by
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mobile device provisioning profiles by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WithDefaults() *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mobile device provisioning profiles by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WithTimeout(timeout time.Duration) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WithContext(ctx context.Context) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WithHTTPClient(client *http.Client) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WithUUID(uuid string) *DeleteMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete mobile device provisioning profiles by UUID params
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobileDeviceProvisioningProfilesByUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
