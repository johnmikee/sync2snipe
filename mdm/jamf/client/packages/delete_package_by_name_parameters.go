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
)

// NewDeletePackageByNameParams creates a new DeletePackageByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePackageByNameParams() *DeletePackageByNameParams {
	return &DeletePackageByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackageByNameParamsWithTimeout creates a new DeletePackageByNameParams object
// with the ability to set a timeout on a request.
func NewDeletePackageByNameParamsWithTimeout(timeout time.Duration) *DeletePackageByNameParams {
	return &DeletePackageByNameParams{
		timeout: timeout,
	}
}

// NewDeletePackageByNameParamsWithContext creates a new DeletePackageByNameParams object
// with the ability to set a context for a request.
func NewDeletePackageByNameParamsWithContext(ctx context.Context) *DeletePackageByNameParams {
	return &DeletePackageByNameParams{
		Context: ctx,
	}
}

// NewDeletePackageByNameParamsWithHTTPClient creates a new DeletePackageByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePackageByNameParamsWithHTTPClient(client *http.Client) *DeletePackageByNameParams {
	return &DeletePackageByNameParams{
		HTTPClient: client,
	}
}

/*
DeletePackageByNameParams contains all the parameters to send to the API endpoint

	for the delete package by name operation.

	Typically these are written to a http.Request.
*/
type DeletePackageByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete package by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackageByNameParams) WithDefaults() *DeletePackageByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete package by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackageByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete package by name params
func (o *DeletePackageByNameParams) WithTimeout(timeout time.Duration) *DeletePackageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete package by name params
func (o *DeletePackageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete package by name params
func (o *DeletePackageByNameParams) WithContext(ctx context.Context) *DeletePackageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete package by name params
func (o *DeletePackageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete package by name params
func (o *DeletePackageByNameParams) WithHTTPClient(client *http.Client) *DeletePackageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete package by name params
func (o *DeletePackageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete package by name params
func (o *DeletePackageByNameParams) WithName(name string) *DeletePackageByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete package by name params
func (o *DeletePackageByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}