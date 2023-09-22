// Code generated by go-swagger; DO NOT EDIT.

package logflush

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

// NewCreateLogFlushForIntervalParams creates a new CreateLogFlushForIntervalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLogFlushForIntervalParams() *CreateLogFlushForIntervalParams {
	return &CreateLogFlushForIntervalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLogFlushForIntervalParamsWithTimeout creates a new CreateLogFlushForIntervalParams object
// with the ability to set a timeout on a request.
func NewCreateLogFlushForIntervalParamsWithTimeout(timeout time.Duration) *CreateLogFlushForIntervalParams {
	return &CreateLogFlushForIntervalParams{
		timeout: timeout,
	}
}

// NewCreateLogFlushForIntervalParamsWithContext creates a new CreateLogFlushForIntervalParams object
// with the ability to set a context for a request.
func NewCreateLogFlushForIntervalParamsWithContext(ctx context.Context) *CreateLogFlushForIntervalParams {
	return &CreateLogFlushForIntervalParams{
		Context: ctx,
	}
}

// NewCreateLogFlushForIntervalParamsWithHTTPClient creates a new CreateLogFlushForIntervalParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLogFlushForIntervalParamsWithHTTPClient(client *http.Client) *CreateLogFlushForIntervalParams {
	return &CreateLogFlushForIntervalParams{
		HTTPClient: client,
	}
}

/*
CreateLogFlushForIntervalParams contains all the parameters to send to the API endpoint

	for the create log flush for interval operation.

	Typically these are written to a http.Request.
*/
type CreateLogFlushForIntervalParams struct {

	/* Interval.

	   Supported values are a combination of [Zero, One, Two, Three, Six] and [Days, Weeks, Months, Years]. For example, JSSResource/logflush/policies/interval/Three+Months
	*/
	Interval string

	/* Log.

	   At this time, only 'policy' logs are supported
	*/
	Log string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create log flush for interval params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogFlushForIntervalParams) WithDefaults() *CreateLogFlushForIntervalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create log flush for interval params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogFlushForIntervalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) WithTimeout(timeout time.Duration) *CreateLogFlushForIntervalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) WithContext(ctx context.Context) *CreateLogFlushForIntervalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) WithHTTPClient(client *http.Client) *CreateLogFlushForIntervalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterval adds the interval to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) WithInterval(interval string) *CreateLogFlushForIntervalParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) SetInterval(interval string) {
	o.Interval = interval
}

// WithLog adds the log to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) WithLog(log string) *CreateLogFlushForIntervalParams {
	o.SetLog(log)
	return o
}

// SetLog adds the log to the create log flush for interval params
func (o *CreateLogFlushForIntervalParams) SetLog(log string) {
	o.Log = log
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLogFlushForIntervalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interval
	if err := r.SetPathParam("interval", o.Interval); err != nil {
		return err
	}

	// path param log
	if err := r.SetPathParam("log", o.Log); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
