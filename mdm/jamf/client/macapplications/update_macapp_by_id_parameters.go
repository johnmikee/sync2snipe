// Code generated by go-swagger; DO NOT EDIT.

package macapplications

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

// NewUpdateMacappByIDParams creates a new UpdateMacappByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMacappByIDParams() *UpdateMacappByIDParams {
	return &UpdateMacappByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMacappByIDParamsWithTimeout creates a new UpdateMacappByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateMacappByIDParamsWithTimeout(timeout time.Duration) *UpdateMacappByIDParams {
	return &UpdateMacappByIDParams{
		timeout: timeout,
	}
}

// NewUpdateMacappByIDParamsWithContext creates a new UpdateMacappByIDParams object
// with the ability to set a context for a request.
func NewUpdateMacappByIDParamsWithContext(ctx context.Context) *UpdateMacappByIDParams {
	return &UpdateMacappByIDParams{
		Context: ctx,
	}
}

// NewUpdateMacappByIDParamsWithHTTPClient creates a new UpdateMacappByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMacappByIDParamsWithHTTPClient(client *http.Client) *UpdateMacappByIDParams {
	return &UpdateMacappByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateMacappByIDParams contains all the parameters to send to the API endpoint

	for the update macapp by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateMacappByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update macapp by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMacappByIDParams) WithDefaults() *UpdateMacappByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update macapp by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMacappByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update macapp by Id params
func (o *UpdateMacappByIDParams) WithTimeout(timeout time.Duration) *UpdateMacappByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update macapp by Id params
func (o *UpdateMacappByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update macapp by Id params
func (o *UpdateMacappByIDParams) WithContext(ctx context.Context) *UpdateMacappByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update macapp by Id params
func (o *UpdateMacappByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update macapp by Id params
func (o *UpdateMacappByIDParams) WithHTTPClient(client *http.Client) *UpdateMacappByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update macapp by Id params
func (o *UpdateMacappByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update macapp by Id params
func (o *UpdateMacappByIDParams) WithID(id int64) *UpdateMacappByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update macapp by Id params
func (o *UpdateMacappByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMacappByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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