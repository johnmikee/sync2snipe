// Code generated by go-swagger; DO NOT EDIT.

package advancedmobiledevicesearches

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

// NewUpdateAdvancedMobileDeviceSearchByIDParams creates a new UpdateAdvancedMobileDeviceSearchByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAdvancedMobileDeviceSearchByIDParams() *UpdateAdvancedMobileDeviceSearchByIDParams {
	return &UpdateAdvancedMobileDeviceSearchByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdvancedMobileDeviceSearchByIDParamsWithTimeout creates a new UpdateAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateAdvancedMobileDeviceSearchByIDParamsWithTimeout(timeout time.Duration) *UpdateAdvancedMobileDeviceSearchByIDParams {
	return &UpdateAdvancedMobileDeviceSearchByIDParams{
		timeout: timeout,
	}
}

// NewUpdateAdvancedMobileDeviceSearchByIDParamsWithContext creates a new UpdateAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a context for a request.
func NewUpdateAdvancedMobileDeviceSearchByIDParamsWithContext(ctx context.Context) *UpdateAdvancedMobileDeviceSearchByIDParams {
	return &UpdateAdvancedMobileDeviceSearchByIDParams{
		Context: ctx,
	}
}

// NewUpdateAdvancedMobileDeviceSearchByIDParamsWithHTTPClient creates a new UpdateAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAdvancedMobileDeviceSearchByIDParamsWithHTTPClient(client *http.Client) *UpdateAdvancedMobileDeviceSearchByIDParams {
	return &UpdateAdvancedMobileDeviceSearchByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateAdvancedMobileDeviceSearchByIDParams contains all the parameters to send to the API endpoint

	for the update advanced mobile device search by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateAdvancedMobileDeviceSearchByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update advanced mobile device search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WithDefaults() *UpdateAdvancedMobileDeviceSearchByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update advanced mobile device search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WithTimeout(timeout time.Duration) *UpdateAdvancedMobileDeviceSearchByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WithContext(ctx context.Context) *UpdateAdvancedMobileDeviceSearchByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WithHTTPClient(client *http.Client) *UpdateAdvancedMobileDeviceSearchByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WithID(id int64) *UpdateAdvancedMobileDeviceSearchByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update advanced mobile device search by Id params
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdvancedMobileDeviceSearchByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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