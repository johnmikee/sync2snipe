// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// NewUpdateClassByIDParams creates a new UpdateClassByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClassByIDParams() *UpdateClassByIDParams {
	return &UpdateClassByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClassByIDParamsWithTimeout creates a new UpdateClassByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateClassByIDParamsWithTimeout(timeout time.Duration) *UpdateClassByIDParams {
	return &UpdateClassByIDParams{
		timeout: timeout,
	}
}

// NewUpdateClassByIDParamsWithContext creates a new UpdateClassByIDParams object
// with the ability to set a context for a request.
func NewUpdateClassByIDParamsWithContext(ctx context.Context) *UpdateClassByIDParams {
	return &UpdateClassByIDParams{
		Context: ctx,
	}
}

// NewUpdateClassByIDParamsWithHTTPClient creates a new UpdateClassByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClassByIDParamsWithHTTPClient(client *http.Client) *UpdateClassByIDParams {
	return &UpdateClassByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateClassByIDParams contains all the parameters to send to the API endpoint

	for the update class by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateClassByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update class by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClassByIDParams) WithDefaults() *UpdateClassByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update class by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClassByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update class by Id params
func (o *UpdateClassByIDParams) WithTimeout(timeout time.Duration) *UpdateClassByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update class by Id params
func (o *UpdateClassByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update class by Id params
func (o *UpdateClassByIDParams) WithContext(ctx context.Context) *UpdateClassByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update class by Id params
func (o *UpdateClassByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update class by Id params
func (o *UpdateClassByIDParams) WithHTTPClient(client *http.Client) *UpdateClassByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update class by Id params
func (o *UpdateClassByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update class by Id params
func (o *UpdateClassByIDParams) WithID(id int64) *UpdateClassByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update class by Id params
func (o *UpdateClassByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClassByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
