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

// NewUpdateMobileDeviceProvisioningProfilesByUUIDParams creates a new UpdateMobileDeviceProvisioningProfilesByUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMobileDeviceProvisioningProfilesByUUIDParams() *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	return &UpdateMobileDeviceProvisioningProfilesByUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithTimeout creates a new UpdateMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a timeout on a request.
func NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithTimeout(timeout time.Duration) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	return &UpdateMobileDeviceProvisioningProfilesByUUIDParams{
		timeout: timeout,
	}
}

// NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithContext creates a new UpdateMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a context for a request.
func NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithContext(ctx context.Context) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	return &UpdateMobileDeviceProvisioningProfilesByUUIDParams{
		Context: ctx,
	}
}

// NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithHTTPClient creates a new UpdateMobileDeviceProvisioningProfilesByUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMobileDeviceProvisioningProfilesByUUIDParamsWithHTTPClient(client *http.Client) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	return &UpdateMobileDeviceProvisioningProfilesByUUIDParams{
		HTTPClient: client,
	}
}

/*
UpdateMobileDeviceProvisioningProfilesByUUIDParams contains all the parameters to send to the API endpoint

	for the update mobile device provisioning profiles by UUID operation.

	Typically these are written to a http.Request.
*/
type UpdateMobileDeviceProvisioningProfilesByUUIDParams struct {

	/* UUID.

	   Uuid value to filter by
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mobile device provisioning profiles by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WithDefaults() *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mobile device provisioning profiles by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WithTimeout(timeout time.Duration) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WithContext(ctx context.Context) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WithHTTPClient(client *http.Client) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WithUUID(uuid string) *UpdateMobileDeviceProvisioningProfilesByUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update mobile device provisioning profiles by UUID params
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMobileDeviceProvisioningProfilesByUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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