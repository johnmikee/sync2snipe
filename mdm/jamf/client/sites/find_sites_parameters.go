// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewFindSitesParams creates a new FindSitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindSitesParams() *FindSitesParams {
	return &FindSitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindSitesParamsWithTimeout creates a new FindSitesParams object
// with the ability to set a timeout on a request.
func NewFindSitesParamsWithTimeout(timeout time.Duration) *FindSitesParams {
	return &FindSitesParams{
		timeout: timeout,
	}
}

// NewFindSitesParamsWithContext creates a new FindSitesParams object
// with the ability to set a context for a request.
func NewFindSitesParamsWithContext(ctx context.Context) *FindSitesParams {
	return &FindSitesParams{
		Context: ctx,
	}
}

// NewFindSitesParamsWithHTTPClient creates a new FindSitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindSitesParamsWithHTTPClient(client *http.Client) *FindSitesParams {
	return &FindSitesParams{
		HTTPClient: client,
	}
}

/*
FindSitesParams contains all the parameters to send to the API endpoint

	for the find sites operation.

	Typically these are written to a http.Request.
*/
type FindSitesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find sites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSitesParams) WithDefaults() *FindSitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find sites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find sites params
func (o *FindSitesParams) WithTimeout(timeout time.Duration) *FindSitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find sites params
func (o *FindSitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find sites params
func (o *FindSitesParams) WithContext(ctx context.Context) *FindSitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find sites params
func (o *FindSitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find sites params
func (o *FindSitesParams) WithHTTPClient(client *http.Client) *FindSitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find sites params
func (o *FindSitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindSitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
