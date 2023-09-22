// Code generated by go-swagger; DO NOT EDIT.

package healthcarelistenerrule

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

// NewFindHealthcareListenerRuleParams creates a new FindHealthcareListenerRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindHealthcareListenerRuleParams() *FindHealthcareListenerRuleParams {
	return &FindHealthcareListenerRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindHealthcareListenerRuleParamsWithTimeout creates a new FindHealthcareListenerRuleParams object
// with the ability to set a timeout on a request.
func NewFindHealthcareListenerRuleParamsWithTimeout(timeout time.Duration) *FindHealthcareListenerRuleParams {
	return &FindHealthcareListenerRuleParams{
		timeout: timeout,
	}
}

// NewFindHealthcareListenerRuleParamsWithContext creates a new FindHealthcareListenerRuleParams object
// with the ability to set a context for a request.
func NewFindHealthcareListenerRuleParamsWithContext(ctx context.Context) *FindHealthcareListenerRuleParams {
	return &FindHealthcareListenerRuleParams{
		Context: ctx,
	}
}

// NewFindHealthcareListenerRuleParamsWithHTTPClient creates a new FindHealthcareListenerRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindHealthcareListenerRuleParamsWithHTTPClient(client *http.Client) *FindHealthcareListenerRuleParams {
	return &FindHealthcareListenerRuleParams{
		HTTPClient: client,
	}
}

/*
FindHealthcareListenerRuleParams contains all the parameters to send to the API endpoint

	for the find healthcare listener rule operation.

	Typically these are written to a http.Request.
*/
type FindHealthcareListenerRuleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find healthcare listener rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindHealthcareListenerRuleParams) WithDefaults() *FindHealthcareListenerRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find healthcare listener rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindHealthcareListenerRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) WithTimeout(timeout time.Duration) *FindHealthcareListenerRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) WithContext(ctx context.Context) *FindHealthcareListenerRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) WithHTTPClient(client *http.Client) *FindHealthcareListenerRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find healthcare listener rule params
func (o *FindHealthcareListenerRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindHealthcareListenerRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
