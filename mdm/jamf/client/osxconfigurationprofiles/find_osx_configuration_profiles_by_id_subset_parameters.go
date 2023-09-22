// Code generated by go-swagger; DO NOT EDIT.

package osxconfigurationprofiles

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

// NewFindOsxConfigurationProfilesByIDSubsetParams creates a new FindOsxConfigurationProfilesByIDSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindOsxConfigurationProfilesByIDSubsetParams() *FindOsxConfigurationProfilesByIDSubsetParams {
	return &FindOsxConfigurationProfilesByIDSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindOsxConfigurationProfilesByIDSubsetParamsWithTimeout creates a new FindOsxConfigurationProfilesByIDSubsetParams object
// with the ability to set a timeout on a request.
func NewFindOsxConfigurationProfilesByIDSubsetParamsWithTimeout(timeout time.Duration) *FindOsxConfigurationProfilesByIDSubsetParams {
	return &FindOsxConfigurationProfilesByIDSubsetParams{
		timeout: timeout,
	}
}

// NewFindOsxConfigurationProfilesByIDSubsetParamsWithContext creates a new FindOsxConfigurationProfilesByIDSubsetParams object
// with the ability to set a context for a request.
func NewFindOsxConfigurationProfilesByIDSubsetParamsWithContext(ctx context.Context) *FindOsxConfigurationProfilesByIDSubsetParams {
	return &FindOsxConfigurationProfilesByIDSubsetParams{
		Context: ctx,
	}
}

// NewFindOsxConfigurationProfilesByIDSubsetParamsWithHTTPClient creates a new FindOsxConfigurationProfilesByIDSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindOsxConfigurationProfilesByIDSubsetParamsWithHTTPClient(client *http.Client) *FindOsxConfigurationProfilesByIDSubsetParams {
	return &FindOsxConfigurationProfilesByIDSubsetParams{
		HTTPClient: client,
	}
}

/*
FindOsxConfigurationProfilesByIDSubsetParams contains all the parameters to send to the API endpoint

	for the find osx configuration profiles by Id subset operation.

	Typically these are written to a http.Request.
*/
type FindOsxConfigurationProfilesByIDSubsetParams struct {

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

// WithDefaults hydrates default values in the find osx configuration profiles by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithDefaults() *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find osx configuration profiles by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithTimeout(timeout time.Duration) *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithContext(ctx context.Context) *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithHTTPClient(client *http.Client) *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithID(id int64) *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetID(id int64) {
	o.ID = id
}

// WithSubset adds the subset to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WithSubset(subset string) *FindOsxConfigurationProfilesByIDSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find osx configuration profiles by Id subset params
func (o *FindOsxConfigurationProfilesByIDSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindOsxConfigurationProfilesByIDSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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