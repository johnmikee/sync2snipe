// Code generated by go-swagger; DO NOT EDIT.

package advancedmobiledevicesearches

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

// NewDeleteAdvancedMobileDeviceSearchByNameParams creates a new DeleteAdvancedMobileDeviceSearchByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdvancedMobileDeviceSearchByNameParams() *DeleteAdvancedMobileDeviceSearchByNameParams {
	return &DeleteAdvancedMobileDeviceSearchByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByNameParamsWithTimeout creates a new DeleteAdvancedMobileDeviceSearchByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteAdvancedMobileDeviceSearchByNameParamsWithTimeout(timeout time.Duration) *DeleteAdvancedMobileDeviceSearchByNameParams {
	return &DeleteAdvancedMobileDeviceSearchByNameParams{
		timeout: timeout,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByNameParamsWithContext creates a new DeleteAdvancedMobileDeviceSearchByNameParams object
// with the ability to set a context for a request.
func NewDeleteAdvancedMobileDeviceSearchByNameParamsWithContext(ctx context.Context) *DeleteAdvancedMobileDeviceSearchByNameParams {
	return &DeleteAdvancedMobileDeviceSearchByNameParams{
		Context: ctx,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByNameParamsWithHTTPClient creates a new DeleteAdvancedMobileDeviceSearchByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdvancedMobileDeviceSearchByNameParamsWithHTTPClient(client *http.Client) *DeleteAdvancedMobileDeviceSearchByNameParams {
	return &DeleteAdvancedMobileDeviceSearchByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteAdvancedMobileDeviceSearchByNameParams contains all the parameters to send to the API endpoint

	for the delete advanced mobile device search by name operation.

	Typically these are written to a http.Request.
*/
type DeleteAdvancedMobileDeviceSearchByNameParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete advanced mobile device search by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WithDefaults() *DeleteAdvancedMobileDeviceSearchByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete advanced mobile device search by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WithTimeout(timeout time.Duration) *DeleteAdvancedMobileDeviceSearchByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WithContext(ctx context.Context) *DeleteAdvancedMobileDeviceSearchByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WithHTTPClient(client *http.Client) *DeleteAdvancedMobileDeviceSearchByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WithName(name string) *DeleteAdvancedMobileDeviceSearchByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete advanced mobile device search by name params
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdvancedMobileDeviceSearchByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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