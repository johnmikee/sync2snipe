// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

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

// NewFindManagedPreferenceProfilesByNameSubsetParams creates a new FindManagedPreferenceProfilesByNameSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindManagedPreferenceProfilesByNameSubsetParams() *FindManagedPreferenceProfilesByNameSubsetParams {
	return &FindManagedPreferenceProfilesByNameSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindManagedPreferenceProfilesByNameSubsetParamsWithTimeout creates a new FindManagedPreferenceProfilesByNameSubsetParams object
// with the ability to set a timeout on a request.
func NewFindManagedPreferenceProfilesByNameSubsetParamsWithTimeout(timeout time.Duration) *FindManagedPreferenceProfilesByNameSubsetParams {
	return &FindManagedPreferenceProfilesByNameSubsetParams{
		timeout: timeout,
	}
}

// NewFindManagedPreferenceProfilesByNameSubsetParamsWithContext creates a new FindManagedPreferenceProfilesByNameSubsetParams object
// with the ability to set a context for a request.
func NewFindManagedPreferenceProfilesByNameSubsetParamsWithContext(ctx context.Context) *FindManagedPreferenceProfilesByNameSubsetParams {
	return &FindManagedPreferenceProfilesByNameSubsetParams{
		Context: ctx,
	}
}

// NewFindManagedPreferenceProfilesByNameSubsetParamsWithHTTPClient creates a new FindManagedPreferenceProfilesByNameSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindManagedPreferenceProfilesByNameSubsetParamsWithHTTPClient(client *http.Client) *FindManagedPreferenceProfilesByNameSubsetParams {
	return &FindManagedPreferenceProfilesByNameSubsetParams{
		HTTPClient: client,
	}
}

/*
FindManagedPreferenceProfilesByNameSubsetParams contains all the parameters to send to the API endpoint

	for the find managed preference profiles by name subset operation.

	Typically these are written to a http.Request.
*/
type FindManagedPreferenceProfilesByNameSubsetParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find managed preference profiles by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithDefaults() *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find managed preference profiles by name subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithTimeout(timeout time.Duration) *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithContext(ctx context.Context) *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithHTTPClient(client *http.Client) *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithName(name string) *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetName(name string) {
	o.Name = name
}

// WithSubset adds the subset to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WithSubset(subset string) *FindManagedPreferenceProfilesByNameSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find managed preference profiles by name subset params
func (o *FindManagedPreferenceProfilesByNameSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindManagedPreferenceProfilesByNameSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
