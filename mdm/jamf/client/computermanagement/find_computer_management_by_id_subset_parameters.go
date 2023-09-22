// Code generated by go-swagger; DO NOT EDIT.

package computermanagement

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

// NewFindComputerManagementByIDSubsetParams creates a new FindComputerManagementByIDSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerManagementByIDSubsetParams() *FindComputerManagementByIDSubsetParams {
	return &FindComputerManagementByIDSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerManagementByIDSubsetParamsWithTimeout creates a new FindComputerManagementByIDSubsetParams object
// with the ability to set a timeout on a request.
func NewFindComputerManagementByIDSubsetParamsWithTimeout(timeout time.Duration) *FindComputerManagementByIDSubsetParams {
	return &FindComputerManagementByIDSubsetParams{
		timeout: timeout,
	}
}

// NewFindComputerManagementByIDSubsetParamsWithContext creates a new FindComputerManagementByIDSubsetParams object
// with the ability to set a context for a request.
func NewFindComputerManagementByIDSubsetParamsWithContext(ctx context.Context) *FindComputerManagementByIDSubsetParams {
	return &FindComputerManagementByIDSubsetParams{
		Context: ctx,
	}
}

// NewFindComputerManagementByIDSubsetParamsWithHTTPClient creates a new FindComputerManagementByIDSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerManagementByIDSubsetParamsWithHTTPClient(client *http.Client) *FindComputerManagementByIDSubsetParams {
	return &FindComputerManagementByIDSubsetParams{
		HTTPClient: client,
	}
}

/*
FindComputerManagementByIDSubsetParams contains all the parameters to send to the API endpoint

	for the find computer management by Id subset operation.

	Typically these are written to a http.Request.
*/
type FindComputerManagementByIDSubsetParams struct {

	/* ID.

	   Computer ID value to filter by
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

// WithDefaults hydrates default values in the find computer management by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByIDSubsetParams) WithDefaults() *FindComputerManagementByIDSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer management by Id subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByIDSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) WithTimeout(timeout time.Duration) *FindComputerManagementByIDSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) WithContext(ctx context.Context) *FindComputerManagementByIDSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) WithHTTPClient(client *http.Client) *FindComputerManagementByIDSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) WithID(id int64) *FindComputerManagementByIDSubsetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) SetID(id int64) {
	o.ID = id
}

// WithSubset adds the subset to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) WithSubset(subset string) *FindComputerManagementByIDSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find computer management by Id subset params
func (o *FindComputerManagementByIDSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerManagementByIDSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
