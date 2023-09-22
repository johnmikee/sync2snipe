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

// NewFindMacappsByIDSubsetParams creates a new FindMacappsByIDSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMacappsByIDSubsetParams() *FindMacappsByIDSubsetParams {
	return &FindMacappsByIDSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMacappsByIDSubsetParamsWithTimeout creates a new FindMacappsByIDSubsetParams object
// with the ability to set a timeout on a request.
func NewFindMacappsByIDSubsetParamsWithTimeout(timeout time.Duration) *FindMacappsByIDSubsetParams {
	return &FindMacappsByIDSubsetParams{
		timeout: timeout,
	}
}

// NewFindMacappsByIDSubsetParamsWithContext creates a new FindMacappsByIDSubsetParams object
// with the ability to set a context for a request.
func NewFindMacappsByIDSubsetParamsWithContext(ctx context.Context) *FindMacappsByIDSubsetParams {
	return &FindMacappsByIDSubsetParams{
		Context: ctx,
	}
}

// NewFindMacappsByIDSubsetParamsWithHTTPClient creates a new FindMacappsByIDSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMacappsByIDSubsetParamsWithHTTPClient(client *http.Client) *FindMacappsByIDSubsetParams {
	return &FindMacappsByIDSubsetParams{
		HTTPClient: client,
	}
}

/*
FindMacappsByIDSubsetParams contains all the parameters to send to the API endpoint

	for the find macapps by Id subset operation.

	Typically these are written to a http.Request.
*/
type FindMacappsByIDSubsetParams struct {

	/* ID.

	   ID to filter by
	*/
	ID int64

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find macapps by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMacappsByIDSubsetParams) WithDefaults() *FindMacappsByIDSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find macapps by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMacappsByIDSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) WithTimeout(timeout time.Duration) *FindMacappsByIDSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) WithContext(ctx context.Context) *FindMacappsByIDSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) WithHTTPClient(client *http.Client) *FindMacappsByIDSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) WithID(id int64) *FindMacappsByIDSubsetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) SetID(id int64) {
	o.ID = id
}

// WithSubset adds the subset to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) WithSubset(subset string) *FindMacappsByIDSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find macapps by Id subset params
func (o *FindMacappsByIDSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindMacappsByIDSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param subset
	if err := r.SetPathParam("subset", o.Subset); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
