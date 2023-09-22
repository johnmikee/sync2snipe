// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewUpdateAccountByIDParams creates a new UpdateAccountByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccountByIDParams() *UpdateAccountByIDParams {
	return &UpdateAccountByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccountByIDParamsWithTimeout creates a new UpdateAccountByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateAccountByIDParamsWithTimeout(timeout time.Duration) *UpdateAccountByIDParams {
	return &UpdateAccountByIDParams{
		timeout: timeout,
	}
}

// NewUpdateAccountByIDParamsWithContext creates a new UpdateAccountByIDParams object
// with the ability to set a context for a request.
func NewUpdateAccountByIDParamsWithContext(ctx context.Context) *UpdateAccountByIDParams {
	return &UpdateAccountByIDParams{
		Context: ctx,
	}
}

// NewUpdateAccountByIDParamsWithHTTPClient creates a new UpdateAccountByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccountByIDParamsWithHTTPClient(client *http.Client) *UpdateAccountByIDParams {
	return &UpdateAccountByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateAccountByIDParams contains all the parameters to send to the API endpoint

	for the update account by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateAccountByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update account by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountByIDParams) WithDefaults() *UpdateAccountByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update account by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update account by Id params
func (o *UpdateAccountByIDParams) WithTimeout(timeout time.Duration) *UpdateAccountByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update account by Id params
func (o *UpdateAccountByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update account by Id params
func (o *UpdateAccountByIDParams) WithContext(ctx context.Context) *UpdateAccountByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update account by Id params
func (o *UpdateAccountByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update account by Id params
func (o *UpdateAccountByIDParams) WithHTTPClient(client *http.Client) *UpdateAccountByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update account by Id params
func (o *UpdateAccountByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update account by Id params
func (o *UpdateAccountByIDParams) WithID(id int64) *UpdateAccountByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update account by Id params
func (o *UpdateAccountByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccountByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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