// Code generated by go-swagger; DO NOT EDIT.

package savedsearches

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

// NewSavedsearchesGetParams creates a new SavedsearchesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSavedsearchesGetParams() *SavedsearchesGetParams {
	return &SavedsearchesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSavedsearchesGetParamsWithTimeout creates a new SavedsearchesGetParams object
// with the ability to set a timeout on a request.
func NewSavedsearchesGetParamsWithTimeout(timeout time.Duration) *SavedsearchesGetParams {
	return &SavedsearchesGetParams{
		timeout: timeout,
	}
}

// NewSavedsearchesGetParamsWithContext creates a new SavedsearchesGetParams object
// with the ability to set a context for a request.
func NewSavedsearchesGetParamsWithContext(ctx context.Context) *SavedsearchesGetParams {
	return &SavedsearchesGetParams{
		Context: ctx,
	}
}

// NewSavedsearchesGetParamsWithHTTPClient creates a new SavedsearchesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSavedsearchesGetParamsWithHTTPClient(client *http.Client) *SavedsearchesGetParams {
	return &SavedsearchesGetParams{
		HTTPClient: client,
	}
}

/*
SavedsearchesGetParams contains all the parameters to send to the API endpoint

	for the savedsearches get operation.

	Typically these are written to a http.Request.
*/
type SavedsearchesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the savedsearches get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SavedsearchesGetParams) WithDefaults() *SavedsearchesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the savedsearches get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SavedsearchesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the savedsearches get params
func (o *SavedsearchesGetParams) WithTimeout(timeout time.Duration) *SavedsearchesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the savedsearches get params
func (o *SavedsearchesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the savedsearches get params
func (o *SavedsearchesGetParams) WithContext(ctx context.Context) *SavedsearchesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the savedsearches get params
func (o *SavedsearchesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the savedsearches get params
func (o *SavedsearchesGetParams) WithHTTPClient(client *http.Client) *SavedsearchesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the savedsearches get params
func (o *SavedsearchesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SavedsearchesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
