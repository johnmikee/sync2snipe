// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUpdateUserByIDParams creates a new UpdateUserByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserByIDParams() *UpdateUserByIDParams {
	return &UpdateUserByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserByIDParamsWithTimeout creates a new UpdateUserByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateUserByIDParamsWithTimeout(timeout time.Duration) *UpdateUserByIDParams {
	return &UpdateUserByIDParams{
		timeout: timeout,
	}
}

// NewUpdateUserByIDParamsWithContext creates a new UpdateUserByIDParams object
// with the ability to set a context for a request.
func NewUpdateUserByIDParamsWithContext(ctx context.Context) *UpdateUserByIDParams {
	return &UpdateUserByIDParams{
		Context: ctx,
	}
}

// NewUpdateUserByIDParamsWithHTTPClient creates a new UpdateUserByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserByIDParamsWithHTTPClient(client *http.Client) *UpdateUserByIDParams {
	return &UpdateUserByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateUserByIDParams contains all the parameters to send to the API endpoint

	for the update user by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateUserByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserByIDParams) WithDefaults() *UpdateUserByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user by Id params
func (o *UpdateUserByIDParams) WithTimeout(timeout time.Duration) *UpdateUserByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user by Id params
func (o *UpdateUserByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user by Id params
func (o *UpdateUserByIDParams) WithContext(ctx context.Context) *UpdateUserByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user by Id params
func (o *UpdateUserByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user by Id params
func (o *UpdateUserByIDParams) WithHTTPClient(client *http.Client) *UpdateUserByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user by Id params
func (o *UpdateUserByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update user by Id params
func (o *UpdateUserByIDParams) WithID(id int64) *UpdateUserByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update user by Id params
func (o *UpdateUserByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
