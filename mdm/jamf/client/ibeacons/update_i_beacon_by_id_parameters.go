// Code generated by go-swagger; DO NOT EDIT.

package ibeacons

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

// NewUpdateIBeaconByIDParams creates a new UpdateIBeaconByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIBeaconByIDParams() *UpdateIBeaconByIDParams {
	return &UpdateIBeaconByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIBeaconByIDParamsWithTimeout creates a new UpdateIBeaconByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateIBeaconByIDParamsWithTimeout(timeout time.Duration) *UpdateIBeaconByIDParams {
	return &UpdateIBeaconByIDParams{
		timeout: timeout,
	}
}

// NewUpdateIBeaconByIDParamsWithContext creates a new UpdateIBeaconByIDParams object
// with the ability to set a context for a request.
func NewUpdateIBeaconByIDParamsWithContext(ctx context.Context) *UpdateIBeaconByIDParams {
	return &UpdateIBeaconByIDParams{
		Context: ctx,
	}
}

// NewUpdateIBeaconByIDParamsWithHTTPClient creates a new UpdateIBeaconByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIBeaconByIDParamsWithHTTPClient(client *http.Client) *UpdateIBeaconByIDParams {
	return &UpdateIBeaconByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateIBeaconByIDParams contains all the parameters to send to the API endpoint

	for the update i beacon by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateIBeaconByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update i beacon by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIBeaconByIDParams) WithDefaults() *UpdateIBeaconByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update i beacon by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIBeaconByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) WithTimeout(timeout time.Duration) *UpdateIBeaconByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) WithContext(ctx context.Context) *UpdateIBeaconByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) WithHTTPClient(client *http.Client) *UpdateIBeaconByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) WithID(id int64) *UpdateIBeaconByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update i beacon by Id params
func (o *UpdateIBeaconByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIBeaconByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
