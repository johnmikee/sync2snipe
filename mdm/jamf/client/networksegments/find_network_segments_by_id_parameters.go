// Code generated by go-swagger; DO NOT EDIT.

package networksegments

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

// NewFindNetworkSegmentsByIDParams creates a new FindNetworkSegmentsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindNetworkSegmentsByIDParams() *FindNetworkSegmentsByIDParams {
	return &FindNetworkSegmentsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindNetworkSegmentsByIDParamsWithTimeout creates a new FindNetworkSegmentsByIDParams object
// with the ability to set a timeout on a request.
func NewFindNetworkSegmentsByIDParamsWithTimeout(timeout time.Duration) *FindNetworkSegmentsByIDParams {
	return &FindNetworkSegmentsByIDParams{
		timeout: timeout,
	}
}

// NewFindNetworkSegmentsByIDParamsWithContext creates a new FindNetworkSegmentsByIDParams object
// with the ability to set a context for a request.
func NewFindNetworkSegmentsByIDParamsWithContext(ctx context.Context) *FindNetworkSegmentsByIDParams {
	return &FindNetworkSegmentsByIDParams{
		Context: ctx,
	}
}

// NewFindNetworkSegmentsByIDParamsWithHTTPClient creates a new FindNetworkSegmentsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindNetworkSegmentsByIDParamsWithHTTPClient(client *http.Client) *FindNetworkSegmentsByIDParams {
	return &FindNetworkSegmentsByIDParams{
		HTTPClient: client,
	}
}

/*
FindNetworkSegmentsByIDParams contains all the parameters to send to the API endpoint

	for the find network segments by Id operation.

	Typically these are written to a http.Request.
*/
type FindNetworkSegmentsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find network segments by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindNetworkSegmentsByIDParams) WithDefaults() *FindNetworkSegmentsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find network segments by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindNetworkSegmentsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) WithTimeout(timeout time.Duration) *FindNetworkSegmentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) WithContext(ctx context.Context) *FindNetworkSegmentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) WithHTTPClient(client *http.Client) *FindNetworkSegmentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) WithID(id int64) *FindNetworkSegmentsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find network segments by Id params
func (o *FindNetworkSegmentsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindNetworkSegmentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
