// Code generated by go-swagger; DO NOT EDIT.

package patchpolicies

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

// NewFindPatchPoliciesParams creates a new FindPatchPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPatchPoliciesParams() *FindPatchPoliciesParams {
	return &FindPatchPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPatchPoliciesParamsWithTimeout creates a new FindPatchPoliciesParams object
// with the ability to set a timeout on a request.
func NewFindPatchPoliciesParamsWithTimeout(timeout time.Duration) *FindPatchPoliciesParams {
	return &FindPatchPoliciesParams{
		timeout: timeout,
	}
}

// NewFindPatchPoliciesParamsWithContext creates a new FindPatchPoliciesParams object
// with the ability to set a context for a request.
func NewFindPatchPoliciesParamsWithContext(ctx context.Context) *FindPatchPoliciesParams {
	return &FindPatchPoliciesParams{
		Context: ctx,
	}
}

// NewFindPatchPoliciesParamsWithHTTPClient creates a new FindPatchPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPatchPoliciesParamsWithHTTPClient(client *http.Client) *FindPatchPoliciesParams {
	return &FindPatchPoliciesParams{
		HTTPClient: client,
	}
}

/*
FindPatchPoliciesParams contains all the parameters to send to the API endpoint

	for the find patch policies operation.

	Typically these are written to a http.Request.
*/
type FindPatchPoliciesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find patch policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPatchPoliciesParams) WithDefaults() *FindPatchPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find patch policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPatchPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find patch policies params
func (o *FindPatchPoliciesParams) WithTimeout(timeout time.Duration) *FindPatchPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find patch policies params
func (o *FindPatchPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find patch policies params
func (o *FindPatchPoliciesParams) WithContext(ctx context.Context) *FindPatchPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find patch policies params
func (o *FindPatchPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find patch policies params
func (o *FindPatchPoliciesParams) WithHTTPClient(client *http.Client) *FindPatchPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find patch policies params
func (o *FindPatchPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindPatchPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
