// Code generated by go-swagger; DO NOT EDIT.

package osxconfigurationprofiles

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

// NewDeleteOsxConfigurationProfileByIDParams creates a new DeleteOsxConfigurationProfileByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOsxConfigurationProfileByIDParams() *DeleteOsxConfigurationProfileByIDParams {
	return &DeleteOsxConfigurationProfileByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOsxConfigurationProfileByIDParamsWithTimeout creates a new DeleteOsxConfigurationProfileByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOsxConfigurationProfileByIDParamsWithTimeout(timeout time.Duration) *DeleteOsxConfigurationProfileByIDParams {
	return &DeleteOsxConfigurationProfileByIDParams{
		timeout: timeout,
	}
}

// NewDeleteOsxConfigurationProfileByIDParamsWithContext creates a new DeleteOsxConfigurationProfileByIDParams object
// with the ability to set a context for a request.
func NewDeleteOsxConfigurationProfileByIDParamsWithContext(ctx context.Context) *DeleteOsxConfigurationProfileByIDParams {
	return &DeleteOsxConfigurationProfileByIDParams{
		Context: ctx,
	}
}

// NewDeleteOsxConfigurationProfileByIDParamsWithHTTPClient creates a new DeleteOsxConfigurationProfileByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOsxConfigurationProfileByIDParamsWithHTTPClient(client *http.Client) *DeleteOsxConfigurationProfileByIDParams {
	return &DeleteOsxConfigurationProfileByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteOsxConfigurationProfileByIDParams contains all the parameters to send to the API endpoint

	for the delete osx configuration profile by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteOsxConfigurationProfileByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete osx configuration profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOsxConfigurationProfileByIDParams) WithDefaults() *DeleteOsxConfigurationProfileByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete osx configuration profile by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOsxConfigurationProfileByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) WithTimeout(timeout time.Duration) *DeleteOsxConfigurationProfileByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) WithContext(ctx context.Context) *DeleteOsxConfigurationProfileByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) WithHTTPClient(client *http.Client) *DeleteOsxConfigurationProfileByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) WithID(id int64) *DeleteOsxConfigurationProfileByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete osx configuration profile by Id params
func (o *DeleteOsxConfigurationProfileByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOsxConfigurationProfileByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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