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

// NewUpdateMobileDeviceGroupByNameParams creates a new UpdateMobileDeviceGroupByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMobileDeviceGroupByNameParams() *UpdateMobileDeviceGroupByNameParams {
	return &UpdateMobileDeviceGroupByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMobileDeviceGroupByNameParamsWithTimeout creates a new UpdateMobileDeviceGroupByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateMobileDeviceGroupByNameParamsWithTimeout(timeout time.Duration) *UpdateMobileDeviceGroupByNameParams {
	return &UpdateMobileDeviceGroupByNameParams{
		timeout: timeout,
	}
}

// NewUpdateMobileDeviceGroupByNameParamsWithContext creates a new UpdateMobileDeviceGroupByNameParams object
// with the ability to set a context for a request.
func NewUpdateMobileDeviceGroupByNameParamsWithContext(ctx context.Context) *UpdateMobileDeviceGroupByNameParams {
	return &UpdateMobileDeviceGroupByNameParams{
		Context: ctx,
	}
}

// NewUpdateMobileDeviceGroupByNameParamsWithHTTPClient creates a new UpdateMobileDeviceGroupByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMobileDeviceGroupByNameParamsWithHTTPClient(client *http.Client) *UpdateMobileDeviceGroupByNameParams {
	return &UpdateMobileDeviceGroupByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateMobileDeviceGroupByNameParams contains all the parameters to send to the API endpoint

	for the update mobile device group by name operation.

	Typically these are written to a http.Request.
*/
type UpdateMobileDeviceGroupByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mobile device group by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceGroupByNameParams) WithDefaults() *UpdateMobileDeviceGroupByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mobile device group by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceGroupByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) WithTimeout(timeout time.Duration) *UpdateMobileDeviceGroupByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) WithContext(ctx context.Context) *UpdateMobileDeviceGroupByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) WithHTTPClient(client *http.Client) *UpdateMobileDeviceGroupByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) WithName(name string) *UpdateMobileDeviceGroupByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update mobile device group by name params
func (o *UpdateMobileDeviceGroupByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMobileDeviceGroupByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
