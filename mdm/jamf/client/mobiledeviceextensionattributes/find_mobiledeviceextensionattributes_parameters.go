// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceextensionattributes

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

// NewFindMobiledeviceextensionattributesParams creates a new FindMobiledeviceextensionattributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobiledeviceextensionattributesParams() *FindMobiledeviceextensionattributesParams {
	return &FindMobiledeviceextensionattributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobiledeviceextensionattributesParamsWithTimeout creates a new FindMobiledeviceextensionattributesParams object
// with the ability to set a timeout on a request.
func NewFindMobiledeviceextensionattributesParamsWithTimeout(timeout time.Duration) *FindMobiledeviceextensionattributesParams {
	return &FindMobiledeviceextensionattributesParams{
		timeout: timeout,
	}
}

// NewFindMobiledeviceextensionattributesParamsWithContext creates a new FindMobiledeviceextensionattributesParams object
// with the ability to set a context for a request.
func NewFindMobiledeviceextensionattributesParamsWithContext(ctx context.Context) *FindMobiledeviceextensionattributesParams {
	return &FindMobiledeviceextensionattributesParams{
		Context: ctx,
	}
}

// NewFindMobiledeviceextensionattributesParamsWithHTTPClient creates a new FindMobiledeviceextensionattributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobiledeviceextensionattributesParamsWithHTTPClient(client *http.Client) *FindMobiledeviceextensionattributesParams {
	return &FindMobiledeviceextensionattributesParams{
		HTTPClient: client,
	}
}

/*
FindMobiledeviceextensionattributesParams contains all the parameters to send to the API endpoint

	for the find mobiledeviceextensionattributes operation.

	Typically these are written to a http.Request.
*/
type FindMobiledeviceextensionattributesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobiledeviceextensionattributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobiledeviceextensionattributesParams) WithDefaults() *FindMobiledeviceextensionattributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobiledeviceextensionattributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobiledeviceextensionattributesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) WithTimeout(timeout time.Duration) *FindMobiledeviceextensionattributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) WithContext(ctx context.Context) *FindMobiledeviceextensionattributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) WithHTTPClient(client *http.Client) *FindMobiledeviceextensionattributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobiledeviceextensionattributes params
func (o *FindMobiledeviceextensionattributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobiledeviceextensionattributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
