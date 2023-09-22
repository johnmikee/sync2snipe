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
)

// NewFindUsersByNameParams creates a new FindUsersByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindUsersByNameParams() *FindUsersByNameParams {
	return &FindUsersByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindUsersByNameParamsWithTimeout creates a new FindUsersByNameParams object
// with the ability to set a timeout on a request.
func NewFindUsersByNameParamsWithTimeout(timeout time.Duration) *FindUsersByNameParams {
	return &FindUsersByNameParams{
		timeout: timeout,
	}
}

// NewFindUsersByNameParamsWithContext creates a new FindUsersByNameParams object
// with the ability to set a context for a request.
func NewFindUsersByNameParamsWithContext(ctx context.Context) *FindUsersByNameParams {
	return &FindUsersByNameParams{
		Context: ctx,
	}
}

// NewFindUsersByNameParamsWithHTTPClient creates a new FindUsersByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindUsersByNameParamsWithHTTPClient(client *http.Client) *FindUsersByNameParams {
	return &FindUsersByNameParams{
		HTTPClient: client,
	}
}

/*
FindUsersByNameParams contains all the parameters to send to the API endpoint

	for the find users by name operation.

	Typically these are written to a http.Request.
*/
type FindUsersByNameParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find users by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUsersByNameParams) WithDefaults() *FindUsersByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find users by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUsersByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find users by name params
func (o *FindUsersByNameParams) WithTimeout(timeout time.Duration) *FindUsersByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find users by name params
func (o *FindUsersByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find users by name params
func (o *FindUsersByNameParams) WithContext(ctx context.Context) *FindUsersByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find users by name params
func (o *FindUsersByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find users by name params
func (o *FindUsersByNameParams) WithHTTPClient(client *http.Client) *FindUsersByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find users by name params
func (o *FindUsersByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find users by name params
func (o *FindUsersByNameParams) WithName(name string) *FindUsersByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find users by name params
func (o *FindUsersByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *FindUsersByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}