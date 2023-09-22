// Code generated by go-swagger; DO NOT EDIT.

package peripheraltypes

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

// NewCreatePeripheralTypeByIDParams creates a new CreatePeripheralTypeByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePeripheralTypeByIDParams() *CreatePeripheralTypeByIDParams {
	return &CreatePeripheralTypeByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePeripheralTypeByIDParamsWithTimeout creates a new CreatePeripheralTypeByIDParams object
// with the ability to set a timeout on a request.
func NewCreatePeripheralTypeByIDParamsWithTimeout(timeout time.Duration) *CreatePeripheralTypeByIDParams {
	return &CreatePeripheralTypeByIDParams{
		timeout: timeout,
	}
}

// NewCreatePeripheralTypeByIDParamsWithContext creates a new CreatePeripheralTypeByIDParams object
// with the ability to set a context for a request.
func NewCreatePeripheralTypeByIDParamsWithContext(ctx context.Context) *CreatePeripheralTypeByIDParams {
	return &CreatePeripheralTypeByIDParams{
		Context: ctx,
	}
}

// NewCreatePeripheralTypeByIDParamsWithHTTPClient creates a new CreatePeripheralTypeByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePeripheralTypeByIDParamsWithHTTPClient(client *http.Client) *CreatePeripheralTypeByIDParams {
	return &CreatePeripheralTypeByIDParams{
		HTTPClient: client,
	}
}

/*
CreatePeripheralTypeByIDParams contains all the parameters to send to the API endpoint

	for the create peripheral type by Id operation.

	Typically these are written to a http.Request.
*/
type CreatePeripheralTypeByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create peripheral type by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePeripheralTypeByIDParams) WithDefaults() *CreatePeripheralTypeByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create peripheral type by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePeripheralTypeByIDParams) SetDefaults() {
	var (
		idDefault = int64(0)
	)

	val := CreatePeripheralTypeByIDParams{
		ID: idDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) WithTimeout(timeout time.Duration) *CreatePeripheralTypeByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) WithContext(ctx context.Context) *CreatePeripheralTypeByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) WithHTTPClient(client *http.Client) *CreatePeripheralTypeByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) WithID(id int64) *CreatePeripheralTypeByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create peripheral type by Id params
func (o *CreatePeripheralTypeByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePeripheralTypeByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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