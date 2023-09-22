// Code generated by go-swagger; DO NOT EDIT.

package departments

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

// NewFindDepartmentsByIDParams creates a new FindDepartmentsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindDepartmentsByIDParams() *FindDepartmentsByIDParams {
	return &FindDepartmentsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindDepartmentsByIDParamsWithTimeout creates a new FindDepartmentsByIDParams object
// with the ability to set a timeout on a request.
func NewFindDepartmentsByIDParamsWithTimeout(timeout time.Duration) *FindDepartmentsByIDParams {
	return &FindDepartmentsByIDParams{
		timeout: timeout,
	}
}

// NewFindDepartmentsByIDParamsWithContext creates a new FindDepartmentsByIDParams object
// with the ability to set a context for a request.
func NewFindDepartmentsByIDParamsWithContext(ctx context.Context) *FindDepartmentsByIDParams {
	return &FindDepartmentsByIDParams{
		Context: ctx,
	}
}

// NewFindDepartmentsByIDParamsWithHTTPClient creates a new FindDepartmentsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindDepartmentsByIDParamsWithHTTPClient(client *http.Client) *FindDepartmentsByIDParams {
	return &FindDepartmentsByIDParams{
		HTTPClient: client,
	}
}

/*
FindDepartmentsByIDParams contains all the parameters to send to the API endpoint

	for the find departments by Id operation.

	Typically these are written to a http.Request.
*/
type FindDepartmentsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find departments by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDepartmentsByIDParams) WithDefaults() *FindDepartmentsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find departments by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDepartmentsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find departments by Id params
func (o *FindDepartmentsByIDParams) WithTimeout(timeout time.Duration) *FindDepartmentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find departments by Id params
func (o *FindDepartmentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find departments by Id params
func (o *FindDepartmentsByIDParams) WithContext(ctx context.Context) *FindDepartmentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find departments by Id params
func (o *FindDepartmentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find departments by Id params
func (o *FindDepartmentsByIDParams) WithHTTPClient(client *http.Client) *FindDepartmentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find departments by Id params
func (o *FindDepartmentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find departments by Id params
func (o *FindDepartmentsByIDParams) WithID(id int64) *FindDepartmentsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find departments by Id params
func (o *FindDepartmentsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindDepartmentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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