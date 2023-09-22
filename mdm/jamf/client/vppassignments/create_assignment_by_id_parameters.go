// Code generated by go-swagger; DO NOT EDIT.

package vppassignments

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

// NewCreateAssignmentByIDParams creates a new CreateAssignmentByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAssignmentByIDParams() *CreateAssignmentByIDParams {
	return &CreateAssignmentByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAssignmentByIDParamsWithTimeout creates a new CreateAssignmentByIDParams object
// with the ability to set a timeout on a request.
func NewCreateAssignmentByIDParamsWithTimeout(timeout time.Duration) *CreateAssignmentByIDParams {
	return &CreateAssignmentByIDParams{
		timeout: timeout,
	}
}

// NewCreateAssignmentByIDParamsWithContext creates a new CreateAssignmentByIDParams object
// with the ability to set a context for a request.
func NewCreateAssignmentByIDParamsWithContext(ctx context.Context) *CreateAssignmentByIDParams {
	return &CreateAssignmentByIDParams{
		Context: ctx,
	}
}

// NewCreateAssignmentByIDParamsWithHTTPClient creates a new CreateAssignmentByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAssignmentByIDParamsWithHTTPClient(client *http.Client) *CreateAssignmentByIDParams {
	return &CreateAssignmentByIDParams{
		HTTPClient: client,
	}
}

/*
CreateAssignmentByIDParams contains all the parameters to send to the API endpoint

	for the create assignment by Id operation.

	Typically these are written to a http.Request.
*/
type CreateAssignmentByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create assignment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAssignmentByIDParams) WithDefaults() *CreateAssignmentByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create assignment by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAssignmentByIDParams) SetDefaults() {
	var (
		idDefault = int64(0)
	)

	val := CreateAssignmentByIDParams{
		ID: idDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create assignment by Id params
func (o *CreateAssignmentByIDParams) WithTimeout(timeout time.Duration) *CreateAssignmentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create assignment by Id params
func (o *CreateAssignmentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create assignment by Id params
func (o *CreateAssignmentByIDParams) WithContext(ctx context.Context) *CreateAssignmentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create assignment by Id params
func (o *CreateAssignmentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create assignment by Id params
func (o *CreateAssignmentByIDParams) WithHTTPClient(client *http.Client) *CreateAssignmentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create assignment by Id params
func (o *CreateAssignmentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create assignment by Id params
func (o *CreateAssignmentByIDParams) WithID(id int64) *CreateAssignmentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create assignment by Id params
func (o *CreateAssignmentByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAssignmentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
