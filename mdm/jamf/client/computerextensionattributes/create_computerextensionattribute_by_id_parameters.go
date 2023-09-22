// Code generated by go-swagger; DO NOT EDIT.

package computerextensionattributes

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

// NewCreateComputerextensionattributeByIDParams creates a new CreateComputerextensionattributeByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateComputerextensionattributeByIDParams() *CreateComputerextensionattributeByIDParams {
	return &CreateComputerextensionattributeByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateComputerextensionattributeByIDParamsWithTimeout creates a new CreateComputerextensionattributeByIDParams object
// with the ability to set a timeout on a request.
func NewCreateComputerextensionattributeByIDParamsWithTimeout(timeout time.Duration) *CreateComputerextensionattributeByIDParams {
	return &CreateComputerextensionattributeByIDParams{
		timeout: timeout,
	}
}

// NewCreateComputerextensionattributeByIDParamsWithContext creates a new CreateComputerextensionattributeByIDParams object
// with the ability to set a context for a request.
func NewCreateComputerextensionattributeByIDParamsWithContext(ctx context.Context) *CreateComputerextensionattributeByIDParams {
	return &CreateComputerextensionattributeByIDParams{
		Context: ctx,
	}
}

// NewCreateComputerextensionattributeByIDParamsWithHTTPClient creates a new CreateComputerextensionattributeByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateComputerextensionattributeByIDParamsWithHTTPClient(client *http.Client) *CreateComputerextensionattributeByIDParams {
	return &CreateComputerextensionattributeByIDParams{
		HTTPClient: client,
	}
}

/*
CreateComputerextensionattributeByIDParams contains all the parameters to send to the API endpoint

	for the create computerextensionattribute by Id operation.

	Typically these are written to a http.Request.
*/
type CreateComputerextensionattributeByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create computerextensionattribute by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComputerextensionattributeByIDParams) WithDefaults() *CreateComputerextensionattributeByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create computerextensionattribute by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComputerextensionattributeByIDParams) SetDefaults() {
	var (
		idDefault = int64(0)
	)

	val := CreateComputerextensionattributeByIDParams{
		ID: idDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) WithTimeout(timeout time.Duration) *CreateComputerextensionattributeByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) WithContext(ctx context.Context) *CreateComputerextensionattributeByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) WithHTTPClient(client *http.Client) *CreateComputerextensionattributeByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) WithID(id int64) *CreateComputerextensionattributeByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create computerextensionattribute by Id params
func (o *CreateComputerextensionattributeByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateComputerextensionattributeByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
