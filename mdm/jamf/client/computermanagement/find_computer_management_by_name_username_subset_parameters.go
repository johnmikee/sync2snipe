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

// NewFindComputerManagementByNameUsernameSubsetParams creates a new FindComputerManagementByNameUsernameSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerManagementByNameUsernameSubsetParams() *FindComputerManagementByNameUsernameSubsetParams {
	return &FindComputerManagementByNameUsernameSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerManagementByNameUsernameSubsetParamsWithTimeout creates a new FindComputerManagementByNameUsernameSubsetParams object
// with the ability to set a timeout on a request.
func NewFindComputerManagementByNameUsernameSubsetParamsWithTimeout(timeout time.Duration) *FindComputerManagementByNameUsernameSubsetParams {
	return &FindComputerManagementByNameUsernameSubsetParams{
		timeout: timeout,
	}
}

// NewFindComputerManagementByNameUsernameSubsetParamsWithContext creates a new FindComputerManagementByNameUsernameSubsetParams object
// with the ability to set a context for a request.
func NewFindComputerManagementByNameUsernameSubsetParamsWithContext(ctx context.Context) *FindComputerManagementByNameUsernameSubsetParams {
	return &FindComputerManagementByNameUsernameSubsetParams{
		Context: ctx,
	}
}

// NewFindComputerManagementByNameUsernameSubsetParamsWithHTTPClient creates a new FindComputerManagementByNameUsernameSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerManagementByNameUsernameSubsetParamsWithHTTPClient(client *http.Client) *FindComputerManagementByNameUsernameSubsetParams {
	return &FindComputerManagementByNameUsernameSubsetParams{
		HTTPClient: client,
	}
}

/*
FindComputerManagementByNameUsernameSubsetParams contains all the parameters to send to the API endpoint

	for the find computer management by name username subset operation.

	Typically these are written to a http.Request.
*/
type FindComputerManagementByNameUsernameSubsetParams struct {

	/* Name.

	   Computer name to filter by
	*/
	Name string

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	/* Username.

	   Username to filter by
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer management by name username subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByNameUsernameSubsetParams) WithDefaults() *FindComputerManagementByNameUsernameSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer management by name username subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerManagementByNameUsernameSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithTimeout(timeout time.Duration) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithContext(ctx context.Context) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithHTTPClient(client *http.Client) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithName(name string) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetName(name string) {
	o.Name = name
}

// WithSubset adds the subset to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithSubset(subset string) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WithUsername adds the username to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) WithUsername(username string) *FindComputerManagementByNameUsernameSubsetParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the find computer management by name username subset params
func (o *FindComputerManagementByNameUsernameSubsetParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerManagementByNameUsernameSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param subset
	if err := r.SetPathParam("subset", o.Subset); err != nil {
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
