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
	"github.com/go-openapi/swag"
)

// NewDeleteUserGroupsByIDParams creates a new DeleteUserGroupsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserGroupsByIDParams() *DeleteUserGroupsByIDParams {
	return &DeleteUserGroupsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserGroupsByIDParamsWithTimeout creates a new DeleteUserGroupsByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteUserGroupsByIDParamsWithTimeout(timeout time.Duration) *DeleteUserGroupsByIDParams {
	return &DeleteUserGroupsByIDParams{
		timeout: timeout,
	}
}

// NewDeleteUserGroupsByIDParamsWithContext creates a new DeleteUserGroupsByIDParams object
// with the ability to set a context for a request.
func NewDeleteUserGroupsByIDParamsWithContext(ctx context.Context) *DeleteUserGroupsByIDParams {
	return &DeleteUserGroupsByIDParams{
		Context: ctx,
	}
}

// NewDeleteUserGroupsByIDParamsWithHTTPClient creates a new DeleteUserGroupsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserGroupsByIDParamsWithHTTPClient(client *http.Client) *DeleteUserGroupsByIDParams {
	return &DeleteUserGroupsByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteUserGroupsByIDParams contains all the parameters to send to the API endpoint

	for the delete user groups by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteUserGroupsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupsByIDParams) WithDefaults() *DeleteUserGroupsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user groups by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) WithTimeout(timeout time.Duration) *DeleteUserGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) WithContext(ctx context.Context) *DeleteUserGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) WithHTTPClient(client *http.Client) *DeleteUserGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) WithID(id int64) *DeleteUserGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete user groups by Id params
func (o *DeleteUserGroupsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
