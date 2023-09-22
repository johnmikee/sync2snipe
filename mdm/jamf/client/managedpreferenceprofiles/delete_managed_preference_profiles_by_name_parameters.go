// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

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

// NewDeleteManagedPreferenceProfilesByNameParams creates a new DeleteManagedPreferenceProfilesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteManagedPreferenceProfilesByNameParams() *DeleteManagedPreferenceProfilesByNameParams {
	return &DeleteManagedPreferenceProfilesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteManagedPreferenceProfilesByNameParamsWithTimeout creates a new DeleteManagedPreferenceProfilesByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteManagedPreferenceProfilesByNameParamsWithTimeout(timeout time.Duration) *DeleteManagedPreferenceProfilesByNameParams {
	return &DeleteManagedPreferenceProfilesByNameParams{
		timeout: timeout,
	}
}

// NewDeleteManagedPreferenceProfilesByNameParamsWithContext creates a new DeleteManagedPreferenceProfilesByNameParams object
// with the ability to set a context for a request.
func NewDeleteManagedPreferenceProfilesByNameParamsWithContext(ctx context.Context) *DeleteManagedPreferenceProfilesByNameParams {
	return &DeleteManagedPreferenceProfilesByNameParams{
		Context: ctx,
	}
}

// NewDeleteManagedPreferenceProfilesByNameParamsWithHTTPClient creates a new DeleteManagedPreferenceProfilesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteManagedPreferenceProfilesByNameParamsWithHTTPClient(client *http.Client) *DeleteManagedPreferenceProfilesByNameParams {
	return &DeleteManagedPreferenceProfilesByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteManagedPreferenceProfilesByNameParams contains all the parameters to send to the API endpoint

	for the delete managed preference profiles by name operation.

	Typically these are written to a http.Request.
*/
type DeleteManagedPreferenceProfilesByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete managed preference profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteManagedPreferenceProfilesByNameParams) WithDefaults() *DeleteManagedPreferenceProfilesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete managed preference profiles by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteManagedPreferenceProfilesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) WithTimeout(timeout time.Duration) *DeleteManagedPreferenceProfilesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) WithContext(ctx context.Context) *DeleteManagedPreferenceProfilesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) WithHTTPClient(client *http.Client) *DeleteManagedPreferenceProfilesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) WithName(name string) *DeleteManagedPreferenceProfilesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete managed preference profiles by name params
func (o *DeleteManagedPreferenceProfilesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteManagedPreferenceProfilesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
