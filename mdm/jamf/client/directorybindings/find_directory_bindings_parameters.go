// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

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

// NewFindDirectoryBindingsParams creates a new FindDirectoryBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindDirectoryBindingsParams() *FindDirectoryBindingsParams {
	return &FindDirectoryBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindDirectoryBindingsParamsWithTimeout creates a new FindDirectoryBindingsParams object
// with the ability to set a timeout on a request.
func NewFindDirectoryBindingsParamsWithTimeout(timeout time.Duration) *FindDirectoryBindingsParams {
	return &FindDirectoryBindingsParams{
		timeout: timeout,
	}
}

// NewFindDirectoryBindingsParamsWithContext creates a new FindDirectoryBindingsParams object
// with the ability to set a context for a request.
func NewFindDirectoryBindingsParamsWithContext(ctx context.Context) *FindDirectoryBindingsParams {
	return &FindDirectoryBindingsParams{
		Context: ctx,
	}
}

// NewFindDirectoryBindingsParamsWithHTTPClient creates a new FindDirectoryBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindDirectoryBindingsParamsWithHTTPClient(client *http.Client) *FindDirectoryBindingsParams {
	return &FindDirectoryBindingsParams{
		HTTPClient: client,
	}
}

/*
FindDirectoryBindingsParams contains all the parameters to send to the API endpoint

	for the find directory bindings operation.

	Typically these are written to a http.Request.
*/
type FindDirectoryBindingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find directory bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDirectoryBindingsParams) WithDefaults() *FindDirectoryBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find directory bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDirectoryBindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find directory bindings params
func (o *FindDirectoryBindingsParams) WithTimeout(timeout time.Duration) *FindDirectoryBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find directory bindings params
func (o *FindDirectoryBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find directory bindings params
func (o *FindDirectoryBindingsParams) WithContext(ctx context.Context) *FindDirectoryBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find directory bindings params
func (o *FindDirectoryBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find directory bindings params
func (o *FindDirectoryBindingsParams) WithHTTPClient(client *http.Client) *FindDirectoryBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find directory bindings params
func (o *FindDirectoryBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindDirectoryBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
