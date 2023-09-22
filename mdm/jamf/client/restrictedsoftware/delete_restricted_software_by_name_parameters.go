// Code generated by go-swagger; DO NOT EDIT.

package restrictedsoftware

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

// NewDeleteRestrictedSoftwareByNameParams creates a new DeleteRestrictedSoftwareByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRestrictedSoftwareByNameParams() *DeleteRestrictedSoftwareByNameParams {
	return &DeleteRestrictedSoftwareByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRestrictedSoftwareByNameParamsWithTimeout creates a new DeleteRestrictedSoftwareByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteRestrictedSoftwareByNameParamsWithTimeout(timeout time.Duration) *DeleteRestrictedSoftwareByNameParams {
	return &DeleteRestrictedSoftwareByNameParams{
		timeout: timeout,
	}
}

// NewDeleteRestrictedSoftwareByNameParamsWithContext creates a new DeleteRestrictedSoftwareByNameParams object
// with the ability to set a context for a request.
func NewDeleteRestrictedSoftwareByNameParamsWithContext(ctx context.Context) *DeleteRestrictedSoftwareByNameParams {
	return &DeleteRestrictedSoftwareByNameParams{
		Context: ctx,
	}
}

// NewDeleteRestrictedSoftwareByNameParamsWithHTTPClient creates a new DeleteRestrictedSoftwareByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRestrictedSoftwareByNameParamsWithHTTPClient(client *http.Client) *DeleteRestrictedSoftwareByNameParams {
	return &DeleteRestrictedSoftwareByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteRestrictedSoftwareByNameParams contains all the parameters to send to the API endpoint

	for the delete restricted software by name operation.

	Typically these are written to a http.Request.
*/
type DeleteRestrictedSoftwareByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete restricted software by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRestrictedSoftwareByNameParams) WithDefaults() *DeleteRestrictedSoftwareByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete restricted software by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRestrictedSoftwareByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) WithTimeout(timeout time.Duration) *DeleteRestrictedSoftwareByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) WithContext(ctx context.Context) *DeleteRestrictedSoftwareByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) WithHTTPClient(client *http.Client) *DeleteRestrictedSoftwareByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) WithName(name string) *DeleteRestrictedSoftwareByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete restricted software by name params
func (o *DeleteRestrictedSoftwareByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRestrictedSoftwareByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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