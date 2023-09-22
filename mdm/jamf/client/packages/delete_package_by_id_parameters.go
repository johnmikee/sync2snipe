// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewDeletePackageByIDParams creates a new DeletePackageByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePackageByIDParams() *DeletePackageByIDParams {
	return &DeletePackageByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackageByIDParamsWithTimeout creates a new DeletePackageByIDParams object
// with the ability to set a timeout on a request.
func NewDeletePackageByIDParamsWithTimeout(timeout time.Duration) *DeletePackageByIDParams {
	return &DeletePackageByIDParams{
		timeout: timeout,
	}
}

// NewDeletePackageByIDParamsWithContext creates a new DeletePackageByIDParams object
// with the ability to set a context for a request.
func NewDeletePackageByIDParamsWithContext(ctx context.Context) *DeletePackageByIDParams {
	return &DeletePackageByIDParams{
		Context: ctx,
	}
}

// NewDeletePackageByIDParamsWithHTTPClient creates a new DeletePackageByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePackageByIDParamsWithHTTPClient(client *http.Client) *DeletePackageByIDParams {
	return &DeletePackageByIDParams{
		HTTPClient: client,
	}
}

/*
DeletePackageByIDParams contains all the parameters to send to the API endpoint

	for the delete package by Id operation.

	Typically these are written to a http.Request.
*/
type DeletePackageByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete package by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackageByIDParams) WithDefaults() *DeletePackageByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete package by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackageByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete package by Id params
func (o *DeletePackageByIDParams) WithTimeout(timeout time.Duration) *DeletePackageByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete package by Id params
func (o *DeletePackageByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete package by Id params
func (o *DeletePackageByIDParams) WithContext(ctx context.Context) *DeletePackageByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete package by Id params
func (o *DeletePackageByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete package by Id params
func (o *DeletePackageByIDParams) WithHTTPClient(client *http.Client) *DeletePackageByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete package by Id params
func (o *DeletePackageByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete package by Id params
func (o *DeletePackageByIDParams) WithID(id int64) *DeletePackageByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete package by Id params
func (o *DeletePackageByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackageByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
