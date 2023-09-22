// Code generated by go-swagger; DO NOT EDIT.

package patchsoftwaretitles

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

// NewPatchsoftwaretitlesIDByIDPostParams creates a new PatchsoftwaretitlesIDByIDPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchsoftwaretitlesIDByIDPostParams() *PatchsoftwaretitlesIDByIDPostParams {
	return &PatchsoftwaretitlesIDByIDPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchsoftwaretitlesIDByIDPostParamsWithTimeout creates a new PatchsoftwaretitlesIDByIDPostParams object
// with the ability to set a timeout on a request.
func NewPatchsoftwaretitlesIDByIDPostParamsWithTimeout(timeout time.Duration) *PatchsoftwaretitlesIDByIDPostParams {
	return &PatchsoftwaretitlesIDByIDPostParams{
		timeout: timeout,
	}
}

// NewPatchsoftwaretitlesIDByIDPostParamsWithContext creates a new PatchsoftwaretitlesIDByIDPostParams object
// with the ability to set a context for a request.
func NewPatchsoftwaretitlesIDByIDPostParamsWithContext(ctx context.Context) *PatchsoftwaretitlesIDByIDPostParams {
	return &PatchsoftwaretitlesIDByIDPostParams{
		Context: ctx,
	}
}

// NewPatchsoftwaretitlesIDByIDPostParamsWithHTTPClient creates a new PatchsoftwaretitlesIDByIDPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchsoftwaretitlesIDByIDPostParamsWithHTTPClient(client *http.Client) *PatchsoftwaretitlesIDByIDPostParams {
	return &PatchsoftwaretitlesIDByIDPostParams{
		HTTPClient: client,
	}
}

/*
PatchsoftwaretitlesIDByIDPostParams contains all the parameters to send to the API endpoint

	for the patchsoftwaretitles Id by Id post operation.

	Typically these are written to a http.Request.
*/
type PatchsoftwaretitlesIDByIDPostParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patchsoftwaretitles Id by Id post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchsoftwaretitlesIDByIDPostParams) WithDefaults() *PatchsoftwaretitlesIDByIDPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patchsoftwaretitles Id by Id post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchsoftwaretitlesIDByIDPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) WithTimeout(timeout time.Duration) *PatchsoftwaretitlesIDByIDPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) WithContext(ctx context.Context) *PatchsoftwaretitlesIDByIDPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) WithHTTPClient(client *http.Client) *PatchsoftwaretitlesIDByIDPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) WithID(id string) *PatchsoftwaretitlesIDByIDPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patchsoftwaretitles Id by Id post params
func (o *PatchsoftwaretitlesIDByIDPostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchsoftwaretitlesIDByIDPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
