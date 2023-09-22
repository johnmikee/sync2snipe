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
	"github.com/go-openapi/swag"
)

// NewUpdateAdvancedComputerSearchByIDParams creates a new UpdateAdvancedComputerSearchByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAdvancedComputerSearchByIDParams() *UpdateAdvancedComputerSearchByIDParams {
	return &UpdateAdvancedComputerSearchByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdvancedComputerSearchByIDParamsWithTimeout creates a new UpdateAdvancedComputerSearchByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateAdvancedComputerSearchByIDParamsWithTimeout(timeout time.Duration) *UpdateAdvancedComputerSearchByIDParams {
	return &UpdateAdvancedComputerSearchByIDParams{
		timeout: timeout,
	}
}

// NewUpdateAdvancedComputerSearchByIDParamsWithContext creates a new UpdateAdvancedComputerSearchByIDParams object
// with the ability to set a context for a request.
func NewUpdateAdvancedComputerSearchByIDParamsWithContext(ctx context.Context) *UpdateAdvancedComputerSearchByIDParams {
	return &UpdateAdvancedComputerSearchByIDParams{
		Context: ctx,
	}
}

// NewUpdateAdvancedComputerSearchByIDParamsWithHTTPClient creates a new UpdateAdvancedComputerSearchByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAdvancedComputerSearchByIDParamsWithHTTPClient(client *http.Client) *UpdateAdvancedComputerSearchByIDParams {
	return &UpdateAdvancedComputerSearchByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateAdvancedComputerSearchByIDParams contains all the parameters to send to the API endpoint

	for the update advanced computer search by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateAdvancedComputerSearchByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update advanced computer search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedComputerSearchByIDParams) WithDefaults() *UpdateAdvancedComputerSearchByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update advanced computer search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdvancedComputerSearchByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) WithTimeout(timeout time.Duration) *UpdateAdvancedComputerSearchByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) WithContext(ctx context.Context) *UpdateAdvancedComputerSearchByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) WithHTTPClient(client *http.Client) *UpdateAdvancedComputerSearchByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) WithID(id int64) *UpdateAdvancedComputerSearchByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update advanced computer search by Id params
func (o *UpdateAdvancedComputerSearchByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdvancedComputerSearchByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
