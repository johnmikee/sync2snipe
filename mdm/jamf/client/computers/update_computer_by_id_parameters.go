// Code generated by go-swagger; DO NOT EDIT.

package computers

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

// NewUpdateComputerByIDParams creates a new UpdateComputerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateComputerByIDParams() *UpdateComputerByIDParams {
	return &UpdateComputerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateComputerByIDParamsWithTimeout creates a new UpdateComputerByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateComputerByIDParamsWithTimeout(timeout time.Duration) *UpdateComputerByIDParams {
	return &UpdateComputerByIDParams{
		timeout: timeout,
	}
}

// NewUpdateComputerByIDParamsWithContext creates a new UpdateComputerByIDParams object
// with the ability to set a context for a request.
func NewUpdateComputerByIDParamsWithContext(ctx context.Context) *UpdateComputerByIDParams {
	return &UpdateComputerByIDParams{
		Context: ctx,
	}
}

// NewUpdateComputerByIDParamsWithHTTPClient creates a new UpdateComputerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateComputerByIDParamsWithHTTPClient(client *http.Client) *UpdateComputerByIDParams {
	return &UpdateComputerByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateComputerByIDParams contains all the parameters to send to the API endpoint

	for the update computer by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateComputerByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update computer by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateComputerByIDParams) WithDefaults() *UpdateComputerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update computer by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateComputerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update computer by Id params
func (o *UpdateComputerByIDParams) WithTimeout(timeout time.Duration) *UpdateComputerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update computer by Id params
func (o *UpdateComputerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update computer by Id params
func (o *UpdateComputerByIDParams) WithContext(ctx context.Context) *UpdateComputerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update computer by Id params
func (o *UpdateComputerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update computer by Id params
func (o *UpdateComputerByIDParams) WithHTTPClient(client *http.Client) *UpdateComputerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update computer by Id params
func (o *UpdateComputerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update computer by Id params
func (o *UpdateComputerByIDParams) WithID(id int64) *UpdateComputerByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update computer by Id params
func (o *UpdateComputerByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateComputerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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