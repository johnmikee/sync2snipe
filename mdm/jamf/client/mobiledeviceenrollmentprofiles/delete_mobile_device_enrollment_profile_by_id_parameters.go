// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

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

// NewDeleteMobileDeviceEnrollmentProfileByIDParams creates a new DeleteMobileDeviceEnrollmentProfileByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMobileDeviceEnrollmentProfileByIDParams() *DeleteMobileDeviceEnrollmentProfileByIDParams {
	return &DeleteMobileDeviceEnrollmentProfileByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithTimeout creates a new DeleteMobileDeviceEnrollmentProfileByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithTimeout(timeout time.Duration) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	return &DeleteMobileDeviceEnrollmentProfileByIDParams{
		timeout: timeout,
	}
}

// NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithContext creates a new DeleteMobileDeviceEnrollmentProfileByIDParams object
// with the ability to set a context for a request.
func NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithContext(ctx context.Context) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	return &DeleteMobileDeviceEnrollmentProfileByIDParams{
		Context: ctx,
	}
}

// NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithHTTPClient creates a new DeleteMobileDeviceEnrollmentProfileByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMobileDeviceEnrollmentProfileByIDParamsWithHTTPClient(client *http.Client) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	return &DeleteMobileDeviceEnrollmentProfileByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteMobileDeviceEnrollmentProfileByIDParams contains all the parameters to send to the API endpoint

	for the delete mobile device enrollment profile by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteMobileDeviceEnrollmentProfileByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mobile device enrollment profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WithDefaults() *DeleteMobileDeviceEnrollmentProfileByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mobile device enrollment profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WithTimeout(timeout time.Duration) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WithContext(ctx context.Context) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WithHTTPClient(client *http.Client) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WithID(id int64) *DeleteMobileDeviceEnrollmentProfileByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete mobile device enrollment profile by Id params
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobileDeviceEnrollmentProfileByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
