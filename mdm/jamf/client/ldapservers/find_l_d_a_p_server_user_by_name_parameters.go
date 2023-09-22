// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

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

// NewFindLDAPServerUserByNameParams creates a new FindLDAPServerUserByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindLDAPServerUserByNameParams() *FindLDAPServerUserByNameParams {
	return &FindLDAPServerUserByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindLDAPServerUserByNameParamsWithTimeout creates a new FindLDAPServerUserByNameParams object
// with the ability to set a timeout on a request.
func NewFindLDAPServerUserByNameParamsWithTimeout(timeout time.Duration) *FindLDAPServerUserByNameParams {
	return &FindLDAPServerUserByNameParams{
		timeout: timeout,
	}
}

// NewFindLDAPServerUserByNameParamsWithContext creates a new FindLDAPServerUserByNameParams object
// with the ability to set a context for a request.
func NewFindLDAPServerUserByNameParamsWithContext(ctx context.Context) *FindLDAPServerUserByNameParams {
	return &FindLDAPServerUserByNameParams{
		Context: ctx,
	}
}

// NewFindLDAPServerUserByNameParamsWithHTTPClient creates a new FindLDAPServerUserByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindLDAPServerUserByNameParamsWithHTTPClient(client *http.Client) *FindLDAPServerUserByNameParams {
	return &FindLDAPServerUserByNameParams{
		HTTPClient: client,
	}
}

/*
FindLDAPServerUserByNameParams contains all the parameters to send to the API endpoint

	for the find l d a p server user by name operation.

	Typically these are written to a http.Request.
*/
type FindLDAPServerUserByNameParams struct {

	/* Name.

	   Server name to filter by
	*/
	Name string

	/* User.

	   User to filter by
	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find l d a p server user by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindLDAPServerUserByNameParams) WithDefaults() *FindLDAPServerUserByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find l d a p server user by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindLDAPServerUserByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) WithTimeout(timeout time.Duration) *FindLDAPServerUserByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) WithContext(ctx context.Context) *FindLDAPServerUserByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) WithHTTPClient(client *http.Client) *FindLDAPServerUserByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) WithName(name string) *FindLDAPServerUserByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) SetName(name string) {
	o.Name = name
}

// WithUser adds the user to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) WithUser(user string) *FindLDAPServerUserByNameParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the find l d a p server user by name params
func (o *FindLDAPServerUserByNameParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *FindLDAPServerUserByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}