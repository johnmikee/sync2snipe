// Code generated by go-swagger; DO NOT EDIT.

package computerhistory

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

// NewFindComputerHistoryBySerialNumberSubsetParams creates a new FindComputerHistoryBySerialNumberSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerHistoryBySerialNumberSubsetParams() *FindComputerHistoryBySerialNumberSubsetParams {
	return &FindComputerHistoryBySerialNumberSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerHistoryBySerialNumberSubsetParamsWithTimeout creates a new FindComputerHistoryBySerialNumberSubsetParams object
// with the ability to set a timeout on a request.
func NewFindComputerHistoryBySerialNumberSubsetParamsWithTimeout(timeout time.Duration) *FindComputerHistoryBySerialNumberSubsetParams {
	return &FindComputerHistoryBySerialNumberSubsetParams{
		timeout: timeout,
	}
}

// NewFindComputerHistoryBySerialNumberSubsetParamsWithContext creates a new FindComputerHistoryBySerialNumberSubsetParams object
// with the ability to set a context for a request.
func NewFindComputerHistoryBySerialNumberSubsetParamsWithContext(ctx context.Context) *FindComputerHistoryBySerialNumberSubsetParams {
	return &FindComputerHistoryBySerialNumberSubsetParams{
		Context: ctx,
	}
}

// NewFindComputerHistoryBySerialNumberSubsetParamsWithHTTPClient creates a new FindComputerHistoryBySerialNumberSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerHistoryBySerialNumberSubsetParamsWithHTTPClient(client *http.Client) *FindComputerHistoryBySerialNumberSubsetParams {
	return &FindComputerHistoryBySerialNumberSubsetParams{
		HTTPClient: client,
	}
}

/*
FindComputerHistoryBySerialNumberSubsetParams contains all the parameters to send to the API endpoint

	for the find computer history by serial number subset operation.

	Typically these are written to a http.Request.
*/
type FindComputerHistoryBySerialNumberSubsetParams struct {

	/* Serialnumber.

	   Computer Serial Number to filter by
	*/
	Serialnumber string

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer history by serial number subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithDefaults() *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer history by serial number subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithTimeout(timeout time.Duration) *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithContext(ctx context.Context) *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithHTTPClient(client *http.Client) *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerialnumber adds the serialnumber to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithSerialnumber(serialnumber string) *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetSerialnumber(serialnumber)
	return o
}

// SetSerialnumber adds the serialnumber to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetSerialnumber(serialnumber string) {
	o.Serialnumber = serialnumber
}

// WithSubset adds the subset to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) WithSubset(subset string) *FindComputerHistoryBySerialNumberSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find computer history by serial number subset params
func (o *FindComputerHistoryBySerialNumberSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerHistoryBySerialNumberSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serialnumber
	if err := r.SetPathParam("serialnumber", o.Serialnumber); err != nil {
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
