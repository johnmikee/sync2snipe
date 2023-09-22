// Code generated by go-swagger; DO NOT EDIT.

package computergroups

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

// NewFindComputerGroupsByIDParams creates a new FindComputerGroupsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerGroupsByIDParams() *FindComputerGroupsByIDParams {
	return &FindComputerGroupsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerGroupsByIDParamsWithTimeout creates a new FindComputerGroupsByIDParams object
// with the ability to set a timeout on a request.
func NewFindComputerGroupsByIDParamsWithTimeout(timeout time.Duration) *FindComputerGroupsByIDParams {
	return &FindComputerGroupsByIDParams{
		timeout: timeout,
	}
}

// NewFindComputerGroupsByIDParamsWithContext creates a new FindComputerGroupsByIDParams object
// with the ability to set a context for a request.
func NewFindComputerGroupsByIDParamsWithContext(ctx context.Context) *FindComputerGroupsByIDParams {
	return &FindComputerGroupsByIDParams{
		Context: ctx,
	}
}

// NewFindComputerGroupsByIDParamsWithHTTPClient creates a new FindComputerGroupsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerGroupsByIDParamsWithHTTPClient(client *http.Client) *FindComputerGroupsByIDParams {
	return &FindComputerGroupsByIDParams{
		HTTPClient: client,
	}
}

/*
FindComputerGroupsByIDParams contains all the parameters to send to the API endpoint

	for the find computer groups by Id operation.

	Typically these are written to a http.Request.
*/
type FindComputerGroupsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerGroupsByIDParams) WithDefaults() *FindComputerGroupsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerGroupsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) WithTimeout(timeout time.Duration) *FindComputerGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) WithContext(ctx context.Context) *FindComputerGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) WithHTTPClient(client *http.Client) *FindComputerGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) WithID(id int64) *FindComputerGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find computer groups by Id params
func (o *FindComputerGroupsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
