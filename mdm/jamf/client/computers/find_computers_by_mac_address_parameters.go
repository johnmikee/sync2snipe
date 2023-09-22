// Code generated by go-swagger; DO NOT EDIT.

package computers

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

// NewFindComputersByMacAddressParams creates a new FindComputersByMacAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputersByMacAddressParams() *FindComputersByMacAddressParams {
	return &FindComputersByMacAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputersByMacAddressParamsWithTimeout creates a new FindComputersByMacAddressParams object
// with the ability to set a timeout on a request.
func NewFindComputersByMacAddressParamsWithTimeout(timeout time.Duration) *FindComputersByMacAddressParams {
	return &FindComputersByMacAddressParams{
		timeout: timeout,
	}
}

// NewFindComputersByMacAddressParamsWithContext creates a new FindComputersByMacAddressParams object
// with the ability to set a context for a request.
func NewFindComputersByMacAddressParamsWithContext(ctx context.Context) *FindComputersByMacAddressParams {
	return &FindComputersByMacAddressParams{
		Context: ctx,
	}
}

// NewFindComputersByMacAddressParamsWithHTTPClient creates a new FindComputersByMacAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputersByMacAddressParamsWithHTTPClient(client *http.Client) *FindComputersByMacAddressParams {
	return &FindComputersByMacAddressParams{
		HTTPClient: client,
	}
}

/*
FindComputersByMacAddressParams contains all the parameters to send to the API endpoint

	for the find computers by mac address operation.

	Typically these are written to a http.Request.
*/
type FindComputersByMacAddressParams struct {

	/* Macaddress.

	   Mac address to filter by
	*/
	Macaddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computers by mac address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputersByMacAddressParams) WithDefaults() *FindComputersByMacAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computers by mac address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputersByMacAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computers by mac address params
func (o *FindComputersByMacAddressParams) WithTimeout(timeout time.Duration) *FindComputersByMacAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computers by mac address params
func (o *FindComputersByMacAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computers by mac address params
func (o *FindComputersByMacAddressParams) WithContext(ctx context.Context) *FindComputersByMacAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computers by mac address params
func (o *FindComputersByMacAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computers by mac address params
func (o *FindComputersByMacAddressParams) WithHTTPClient(client *http.Client) *FindComputersByMacAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computers by mac address params
func (o *FindComputersByMacAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacaddress adds the macaddress to the find computers by mac address params
func (o *FindComputersByMacAddressParams) WithMacaddress(macaddress string) *FindComputersByMacAddressParams {
	o.SetMacaddress(macaddress)
	return o
}

// SetMacaddress adds the macaddress to the find computers by mac address params
func (o *FindComputersByMacAddressParams) SetMacaddress(macaddress string) {
	o.Macaddress = macaddress
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputersByMacAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param macaddress
	if err := r.SetPathParam("macaddress", o.Macaddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
