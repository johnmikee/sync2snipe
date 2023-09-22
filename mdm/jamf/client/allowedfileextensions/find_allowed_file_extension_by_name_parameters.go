// Code generated by go-swagger; DO NOT EDIT.

package allowedfileextensions

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

// NewFindAllowedFileExtensionByNameParams creates a new FindAllowedFileExtensionByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindAllowedFileExtensionByNameParams() *FindAllowedFileExtensionByNameParams {
	return &FindAllowedFileExtensionByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindAllowedFileExtensionByNameParamsWithTimeout creates a new FindAllowedFileExtensionByNameParams object
// with the ability to set a timeout on a request.
func NewFindAllowedFileExtensionByNameParamsWithTimeout(timeout time.Duration) *FindAllowedFileExtensionByNameParams {
	return &FindAllowedFileExtensionByNameParams{
		timeout: timeout,
	}
}

// NewFindAllowedFileExtensionByNameParamsWithContext creates a new FindAllowedFileExtensionByNameParams object
// with the ability to set a context for a request.
func NewFindAllowedFileExtensionByNameParamsWithContext(ctx context.Context) *FindAllowedFileExtensionByNameParams {
	return &FindAllowedFileExtensionByNameParams{
		Context: ctx,
	}
}

// NewFindAllowedFileExtensionByNameParamsWithHTTPClient creates a new FindAllowedFileExtensionByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindAllowedFileExtensionByNameParamsWithHTTPClient(client *http.Client) *FindAllowedFileExtensionByNameParams {
	return &FindAllowedFileExtensionByNameParams{
		HTTPClient: client,
	}
}

/*
FindAllowedFileExtensionByNameParams contains all the parameters to send to the API endpoint

	for the find allowed file extension by name operation.

	Typically these are written to a http.Request.
*/
type FindAllowedFileExtensionByNameParams struct {

	/* Extension.

	   String value of extension
	*/
	Extension string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find allowed file extension by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindAllowedFileExtensionByNameParams) WithDefaults() *FindAllowedFileExtensionByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find allowed file extension by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindAllowedFileExtensionByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) WithTimeout(timeout time.Duration) *FindAllowedFileExtensionByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) WithContext(ctx context.Context) *FindAllowedFileExtensionByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) WithHTTPClient(client *http.Client) *FindAllowedFileExtensionByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtension adds the extension to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) WithExtension(extension string) *FindAllowedFileExtensionByNameParams {
	o.SetExtension(extension)
	return o
}

// SetExtension adds the extension to the find allowed file extension by name params
func (o *FindAllowedFileExtensionByNameParams) SetExtension(extension string) {
	o.Extension = extension
}

// WriteToRequest writes these params to a swagger request
func (o *FindAllowedFileExtensionByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extension
	if err := r.SetPathParam("extension", o.Extension); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}