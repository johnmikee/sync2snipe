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
)

// NewFindComputerManagementByUDIDParams creates a new FindComputerManagementByUDIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerManagementByUDIDParams() *FindComputerManagementByUDIDParams {
	return &FindComputerManagementByUDIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerManagementByUDIDParamsWithTimeout creates a new FindComputerManagementByUDIDParams object
// with the ability to set a timeout on a request.
func NewFindComputerManagementByUDIDParamsWithTimeout(timeout time.Duration) *FindComputerManagementByUDIDParams {
	return &FindComputerManagementByUDIDParams{
		timeout: timeout,
	}
}

// NewFindComputerManagementByUDIDParamsWithContext creates a new FindComputerManagementByUDIDParams object
// with the ability to set a context for a request.
func NewFindComputerManagementByUDIDParamsWithContext(ctx context.Context) *FindComputerManagementByUDIDParams {
	return &FindComputerManagementByUDIDParams{
		Context: ctx,
	}
}

// NewFindComputerManagementByUDIDParamsWithHTTPClient creates a new FindComputerManagementByUDIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerManagementByUDIDParamsWithHTTPClient(client *http.Client) *FindComputerManagementByUDIDParams {
	return &FindComputerManagementByUDIDParams{
		HTTPClient: client,
	}
}

/*
FindComputerManagementByUDIDParams contains all the parameters to send to the API endpoint

	for the find computer management by u d ID operation.

	Typically these are written to a http.Request.
*/
type FindComputerManagementByUDIDParams struct {

	/* Udid.

	   Computer UDID to filter by
	*/
	Udid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer management by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByUDIDParams) WithDefaults() *FindComputerManagementByUDIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer management by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByUDIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) WithTimeout(timeout time.Duration) *FindComputerManagementByUDIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) WithContext(ctx context.Context) *FindComputerManagementByUDIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) WithHTTPClient(client *http.Client) *FindComputerManagementByUDIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUdid adds the udid to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) WithUdid(udid string) *FindComputerManagementByUDIDParams {
	o.SetUdid(udid)
	return o
}

// SetUdid adds the udid to the find computer management by u d ID params
func (o *FindComputerManagementByUDIDParams) SetUdid(udid string) {
	o.Udid = udid
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerManagementByUDIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param udid
	if err := r.SetPathParam("udid", o.Udid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
