// Code generated by go-swagger; DO NOT EDIT.

package usergroups

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

// NewFindUserGroupsParams creates a new FindUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindUserGroupsParams() *FindUserGroupsParams {
	return &FindUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindUserGroupsParamsWithTimeout creates a new FindUserGroupsParams object
// with the ability to set a timeout on a request.
func NewFindUserGroupsParamsWithTimeout(timeout time.Duration) *FindUserGroupsParams {
	return &FindUserGroupsParams{
		timeout: timeout,
	}
}

// NewFindUserGroupsParamsWithContext creates a new FindUserGroupsParams object
// with the ability to set a context for a request.
func NewFindUserGroupsParamsWithContext(ctx context.Context) *FindUserGroupsParams {
	return &FindUserGroupsParams{
		Context: ctx,
	}
}

// NewFindUserGroupsParamsWithHTTPClient creates a new FindUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindUserGroupsParamsWithHTTPClient(client *http.Client) *FindUserGroupsParams {
	return &FindUserGroupsParams{
		HTTPClient: client,
	}
}

/*
FindUserGroupsParams contains all the parameters to send to the API endpoint

	for the find user groups operation.

	Typically these are written to a http.Request.
*/
type FindUserGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserGroupsParams) WithDefaults() *FindUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find user groups params
func (o *FindUserGroupsParams) WithTimeout(timeout time.Duration) *FindUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find user groups params
func (o *FindUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find user groups params
func (o *FindUserGroupsParams) WithContext(ctx context.Context) *FindUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find user groups params
func (o *FindUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find user groups params
func (o *FindUserGroupsParams) WithHTTPClient(client *http.Client) *FindUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find user groups params
func (o *FindUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}