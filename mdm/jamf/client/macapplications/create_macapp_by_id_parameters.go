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

// NewCreateMacappByIDParams creates a new CreateMacappByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMacappByIDParams() *CreateMacappByIDParams {
	return &CreateMacappByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMacappByIDParamsWithTimeout creates a new CreateMacappByIDParams object
// with the ability to set a timeout on a request.
func NewCreateMacappByIDParamsWithTimeout(timeout time.Duration) *CreateMacappByIDParams {
	return &CreateMacappByIDParams{
		timeout: timeout,
	}
}

// NewCreateMacappByIDParamsWithContext creates a new CreateMacappByIDParams object
// with the ability to set a context for a request.
func NewCreateMacappByIDParamsWithContext(ctx context.Context) *CreateMacappByIDParams {
	return &CreateMacappByIDParams{
		Context: ctx,
	}
}

// NewCreateMacappByIDParamsWithHTTPClient creates a new CreateMacappByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMacappByIDParamsWithHTTPClient(client *http.Client) *CreateMacappByIDParams {
	return &CreateMacappByIDParams{
		HTTPClient: client,
	}
}

/*
CreateMacappByIDParams contains all the parameters to send to the API endpoint

	for the create macapp by Id operation.

	Typically these are written to a http.Request.
*/
type CreateMacappByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create macapp by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMacappByIDParams) WithDefaults() *CreateMacappByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create macapp by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMacappByIDParams) SetDefaults() {
	var (
		idDefault = int64(0)
	)

	val := CreateMacappByIDParams{
		ID: idDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create macapp by Id params
func (o *CreateMacappByIDParams) WithTimeout(timeout time.Duration) *CreateMacappByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create macapp by Id params
func (o *CreateMacappByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create macapp by Id params
func (o *CreateMacappByIDParams) WithContext(ctx context.Context) *CreateMacappByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create macapp by Id params
func (o *CreateMacappByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create macapp by Id params
func (o *CreateMacappByIDParams) WithHTTPClient(client *http.Client) *CreateMacappByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create macapp by Id params
func (o *CreateMacappByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create macapp by Id params
func (o *CreateMacappByIDParams) WithID(id int64) *CreateMacappByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create macapp by Id params
func (o *CreateMacappByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMacappByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
