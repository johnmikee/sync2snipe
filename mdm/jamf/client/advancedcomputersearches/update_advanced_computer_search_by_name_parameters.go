// Code generated by go-swagger; DO NOT EDIT.

package advancedcomputersearches

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

// NewUpdateAdvancedComputerSearchByNameParams creates a new UpdateAdvancedComputerSearchByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAdvancedComputerSearchByNameParams() *UpdateAdvancedComputerSearchByNameParams {
	return &UpdateAdvancedComputerSearchByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdvancedComputerSearchByNameParamsWithTimeout creates a new UpdateAdvancedComputerSearchByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateAdvancedComputerSearchByNameParamsWithTimeout(timeout time.Duration) *UpdateAdvancedComputerSearchByNameParams {
	return &UpdateAdvancedComputerSearchByNameParams{
		timeout: timeout,
	}
}

// NewUpdateAdvancedComputerSearchByNameParamsWithContext creates a new UpdateAdvancedComputerSearchByNameParams object
// with the ability to set a context for a request.
func NewUpdateAdvancedComputerSearchByNameParamsWithContext(ctx context.Context) *UpdateAdvancedComputerSearchByNameParams {
	return &UpdateAdvancedComputerSearchByNameParams{
		Context: ctx,
	}
}

// NewUpdateAdvancedComputerSearchByNameParamsWithHTTPClient creates a new UpdateAdvancedComputerSearchByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAdvancedComputerSearchByNameParamsWithHTTPClient(client *http.Client) *UpdateAdvancedComputerSearchByNameParams {
	return &UpdateAdvancedComputerSearchByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateAdvancedComputerSearchByNameParams contains all the parameters to send to the API endpoint

	for the update advanced computer search by name operation.

	Typically these are written to a http.Request.
*/
type UpdateAdvancedComputerSearchByNameParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update advanced computer search by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedComputerSearchByNameParams) WithDefaults() *UpdateAdvancedComputerSearchByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update advanced computer search by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedComputerSearchByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) WithTimeout(timeout time.Duration) *UpdateAdvancedComputerSearchByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) WithContext(ctx context.Context) *UpdateAdvancedComputerSearchByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) WithHTTPClient(client *http.Client) *UpdateAdvancedComputerSearchByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) WithName(name string) *UpdateAdvancedComputerSearchByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update advanced computer search by name params
func (o *UpdateAdvancedComputerSearchByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdvancedComputerSearchByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
