// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicecommands

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

// NewFindMobileDeviceCommandsByCommandParams creates a new FindMobileDeviceCommandsByCommandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceCommandsByCommandParams() *FindMobileDeviceCommandsByCommandParams {
	return &FindMobileDeviceCommandsByCommandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceCommandsByCommandParamsWithTimeout creates a new FindMobileDeviceCommandsByCommandParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceCommandsByCommandParamsWithTimeout(timeout time.Duration) *FindMobileDeviceCommandsByCommandParams {
	return &FindMobileDeviceCommandsByCommandParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceCommandsByCommandParamsWithContext creates a new FindMobileDeviceCommandsByCommandParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceCommandsByCommandParamsWithContext(ctx context.Context) *FindMobileDeviceCommandsByCommandParams {
	return &FindMobileDeviceCommandsByCommandParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceCommandsByCommandParamsWithHTTPClient creates a new FindMobileDeviceCommandsByCommandParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceCommandsByCommandParamsWithHTTPClient(client *http.Client) *FindMobileDeviceCommandsByCommandParams {
	return &FindMobileDeviceCommandsByCommandParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceCommandsByCommandParams contains all the parameters to send to the API endpoint

	for the find mobile device commands by command operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceCommandsByCommandParams struct {

	/* Command.

	   Name to filter by
	*/
	Command string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device commands by command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceCommandsByCommandParams) WithDefaults() *FindMobileDeviceCommandsByCommandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device commands by command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceCommandsByCommandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) WithTimeout(timeout time.Duration) *FindMobileDeviceCommandsByCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) WithContext(ctx context.Context) *FindMobileDeviceCommandsByCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) WithHTTPClient(client *http.Client) *FindMobileDeviceCommandsByCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommand adds the command to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) WithCommand(command string) *FindMobileDeviceCommandsByCommandParams {
	o.SetCommand(command)
	return o
}

// SetCommand adds the command to the find mobile device commands by command params
func (o *FindMobileDeviceCommandsByCommandParams) SetCommand(command string) {
	o.Command = command
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceCommandsByCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param command
	if err := r.SetPathParam("command", o.Command); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
