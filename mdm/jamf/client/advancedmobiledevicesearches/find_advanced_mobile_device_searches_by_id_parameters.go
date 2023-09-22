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

// NewFindAdvancedMobileDeviceSearchesByIDParams creates a new FindAdvancedMobileDeviceSearchesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindAdvancedMobileDeviceSearchesByIDParams() *FindAdvancedMobileDeviceSearchesByIDParams {
	return &FindAdvancedMobileDeviceSearchesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindAdvancedMobileDeviceSearchesByIDParamsWithTimeout creates a new FindAdvancedMobileDeviceSearchesByIDParams object
// with the ability to set a timeout on a request.
func NewFindAdvancedMobileDeviceSearchesByIDParamsWithTimeout(timeout time.Duration) *FindAdvancedMobileDeviceSearchesByIDParams {
	return &FindAdvancedMobileDeviceSearchesByIDParams{
		timeout: timeout,
	}
}

// NewFindAdvancedMobileDeviceSearchesByIDParamsWithContext creates a new FindAdvancedMobileDeviceSearchesByIDParams object
// with the ability to set a context for a request.
func NewFindAdvancedMobileDeviceSearchesByIDParamsWithContext(ctx context.Context) *FindAdvancedMobileDeviceSearchesByIDParams {
	return &FindAdvancedMobileDeviceSearchesByIDParams{
		Context: ctx,
	}
}

// NewFindAdvancedMobileDeviceSearchesByIDParamsWithHTTPClient creates a new FindAdvancedMobileDeviceSearchesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindAdvancedMobileDeviceSearchesByIDParamsWithHTTPClient(client *http.Client) *FindAdvancedMobileDeviceSearchesByIDParams {
	return &FindAdvancedMobileDeviceSearchesByIDParams{
		HTTPClient: client,
	}
}

/*
FindAdvancedMobileDeviceSearchesByIDParams contains all the parameters to send to the API endpoint

	for the find advanced mobile device searches by Id operation.

	Typically these are written to a http.Request.
*/
type FindAdvancedMobileDeviceSearchesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find advanced mobile device searches by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WithDefaults() *FindAdvancedMobileDeviceSearchesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find advanced mobile device searches by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindAdvancedMobileDeviceSearchesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WithTimeout(timeout time.Duration) *FindAdvancedMobileDeviceSearchesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WithContext(ctx context.Context) *FindAdvancedMobileDeviceSearchesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WithHTTPClient(client *http.Client) *FindAdvancedMobileDeviceSearchesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WithID(id int64) *FindAdvancedMobileDeviceSearchesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find advanced mobile device searches by Id params
func (o *FindAdvancedMobileDeviceSearchesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindAdvancedMobileDeviceSearchesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
