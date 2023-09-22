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

// NewFindComputerManagementByIDUsernameParams creates a new FindComputerManagementByIDUsernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerManagementByIDUsernameParams() *FindComputerManagementByIDUsernameParams {
	return &FindComputerManagementByIDUsernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerManagementByIDUsernameParamsWithTimeout creates a new FindComputerManagementByIDUsernameParams object
// with the ability to set a timeout on a request.
func NewFindComputerManagementByIDUsernameParamsWithTimeout(timeout time.Duration) *FindComputerManagementByIDUsernameParams {
	return &FindComputerManagementByIDUsernameParams{
		timeout: timeout,
	}
}

// NewFindComputerManagementByIDUsernameParamsWithContext creates a new FindComputerManagementByIDUsernameParams object
// with the ability to set a context for a request.
func NewFindComputerManagementByIDUsernameParamsWithContext(ctx context.Context) *FindComputerManagementByIDUsernameParams {
	return &FindComputerManagementByIDUsernameParams{
		Context: ctx,
	}
}

// NewFindComputerManagementByIDUsernameParamsWithHTTPClient creates a new FindComputerManagementByIDUsernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerManagementByIDUsernameParamsWithHTTPClient(client *http.Client) *FindComputerManagementByIDUsernameParams {
	return &FindComputerManagementByIDUsernameParams{
		HTTPClient: client,
	}
}

/*
FindComputerManagementByIDUsernameParams contains all the parameters to send to the API endpoint

	for the find computer management by Id username operation.

	Typically these are written to a http.Request.
*/
type FindComputerManagementByIDUsernameParams struct {

	/* ID.

	   Computer ID to filter by
	*/
	ID int64

	/* Username.

	   Username to filter by
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer management by Id username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByIDUsernameParams) WithDefaults() *FindComputerManagementByIDUsernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer management by Id username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByIDUsernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) WithTimeout(timeout time.Duration) *FindComputerManagementByIDUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) WithContext(ctx context.Context) *FindComputerManagementByIDUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) WithHTTPClient(client *http.Client) *FindComputerManagementByIDUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) WithID(id int64) *FindComputerManagementByIDUsernameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) SetID(id int64) {
	o.ID = id
}

// WithUsername adds the username to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) WithUsername(username string) *FindComputerManagementByIDUsernameParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the find computer management by Id username params
func (o *FindComputerManagementByIDUsernameParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerManagementByIDUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}