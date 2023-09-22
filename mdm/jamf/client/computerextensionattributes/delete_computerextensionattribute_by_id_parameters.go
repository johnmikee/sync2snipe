// Code generated by go-swagger; DO NOT EDIT.

package computerextensionattributes

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

// NewDeleteComputerextensionattributeByIDParams creates a new DeleteComputerextensionattributeByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteComputerextensionattributeByIDParams() *DeleteComputerextensionattributeByIDParams {
	return &DeleteComputerextensionattributeByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteComputerextensionattributeByIDParamsWithTimeout creates a new DeleteComputerextensionattributeByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteComputerextensionattributeByIDParamsWithTimeout(timeout time.Duration) *DeleteComputerextensionattributeByIDParams {
	return &DeleteComputerextensionattributeByIDParams{
		timeout: timeout,
	}
}

// NewDeleteComputerextensionattributeByIDParamsWithContext creates a new DeleteComputerextensionattributeByIDParams object
// with the ability to set a context for a request.
func NewDeleteComputerextensionattributeByIDParamsWithContext(ctx context.Context) *DeleteComputerextensionattributeByIDParams {
	return &DeleteComputerextensionattributeByIDParams{
		Context: ctx,
	}
}

// NewDeleteComputerextensionattributeByIDParamsWithHTTPClient creates a new DeleteComputerextensionattributeByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteComputerextensionattributeByIDParamsWithHTTPClient(client *http.Client) *DeleteComputerextensionattributeByIDParams {
	return &DeleteComputerextensionattributeByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteComputerextensionattributeByIDParams contains all the parameters to send to the API endpoint

	for the delete computerextensionattribute by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteComputerextensionattributeByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete computerextensionattribute by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteComputerextensionattributeByIDParams) WithDefaults() *DeleteComputerextensionattributeByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete computerextensionattribute by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteComputerextensionattributeByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) WithTimeout(timeout time.Duration) *DeleteComputerextensionattributeByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) WithContext(ctx context.Context) *DeleteComputerextensionattributeByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) WithHTTPClient(client *http.Client) *DeleteComputerextensionattributeByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) WithID(id int64) *DeleteComputerextensionattributeByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete computerextensionattribute by Id params
func (o *DeleteComputerextensionattributeByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteComputerextensionattributeByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
