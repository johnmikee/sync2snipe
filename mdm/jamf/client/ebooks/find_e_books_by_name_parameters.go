// Code generated by go-swagger; DO NOT EDIT.

package ebooks

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

// NewFindEBooksByNameParams creates a new FindEBooksByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindEBooksByNameParams() *FindEBooksByNameParams {
	return &FindEBooksByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindEBooksByNameParamsWithTimeout creates a new FindEBooksByNameParams object
// with the ability to set a timeout on a request.
func NewFindEBooksByNameParamsWithTimeout(timeout time.Duration) *FindEBooksByNameParams {
	return &FindEBooksByNameParams{
		timeout: timeout,
	}
}

// NewFindEBooksByNameParamsWithContext creates a new FindEBooksByNameParams object
// with the ability to set a context for a request.
func NewFindEBooksByNameParamsWithContext(ctx context.Context) *FindEBooksByNameParams {
	return &FindEBooksByNameParams{
		Context: ctx,
	}
}

// NewFindEBooksByNameParamsWithHTTPClient creates a new FindEBooksByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindEBooksByNameParamsWithHTTPClient(client *http.Client) *FindEBooksByNameParams {
	return &FindEBooksByNameParams{
		HTTPClient: client,
	}
}

/*
FindEBooksByNameParams contains all the parameters to send to the API endpoint

	for the find e books by name operation.

	Typically these are written to a http.Request.
*/
type FindEBooksByNameParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find e books by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEBooksByNameParams) WithDefaults() *FindEBooksByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find e books by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEBooksByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find e books by name params
func (o *FindEBooksByNameParams) WithTimeout(timeout time.Duration) *FindEBooksByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find e books by name params
func (o *FindEBooksByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find e books by name params
func (o *FindEBooksByNameParams) WithContext(ctx context.Context) *FindEBooksByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find e books by name params
func (o *FindEBooksByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find e books by name params
func (o *FindEBooksByNameParams) WithHTTPClient(client *http.Client) *FindEBooksByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find e books by name params
func (o *FindEBooksByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find e books by name params
func (o *FindEBooksByNameParams) WithName(name string) *FindEBooksByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find e books by name params
func (o *FindEBooksByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *FindEBooksByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
