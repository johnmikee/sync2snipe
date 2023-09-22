// Code generated by go-swagger; DO NOT EDIT.

package computerhardwaresoftwarereports

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

// NewFindComputerHardwareSoftwareReportsByUDIDParams creates a new FindComputerHardwareSoftwareReportsByUDIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerHardwareSoftwareReportsByUDIDParams() *FindComputerHardwareSoftwareReportsByUDIDParams {
	return &FindComputerHardwareSoftwareReportsByUDIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerHardwareSoftwareReportsByUDIDParamsWithTimeout creates a new FindComputerHardwareSoftwareReportsByUDIDParams object
// with the ability to set a timeout on a request.
func NewFindComputerHardwareSoftwareReportsByUDIDParamsWithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsByUDIDParams {
	return &FindComputerHardwareSoftwareReportsByUDIDParams{
		timeout: timeout,
	}
}

// NewFindComputerHardwareSoftwareReportsByUDIDParamsWithContext creates a new FindComputerHardwareSoftwareReportsByUDIDParams object
// with the ability to set a context for a request.
func NewFindComputerHardwareSoftwareReportsByUDIDParamsWithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsByUDIDParams {
	return &FindComputerHardwareSoftwareReportsByUDIDParams{
		Context: ctx,
	}
}

// NewFindComputerHardwareSoftwareReportsByUDIDParamsWithHTTPClient creates a new FindComputerHardwareSoftwareReportsByUDIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerHardwareSoftwareReportsByUDIDParamsWithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsByUDIDParams {
	return &FindComputerHardwareSoftwareReportsByUDIDParams{
		HTTPClient: client,
	}
}

/*
FindComputerHardwareSoftwareReportsByUDIDParams contains all the parameters to send to the API endpoint

	for the find computer hardware software reports by u d ID operation.

	Typically these are written to a http.Request.
*/
type FindComputerHardwareSoftwareReportsByUDIDParams struct {

	/* EndDate.

	   End date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	EndDate strfmt.Date

	/* StartDate.

	   Start date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	StartDate strfmt.Date

	/* Udid.

	   UDID to filter by
	*/
	Udid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer hardware software reports by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithDefaults() *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer hardware software reports by u d ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithEndDate(endDate strfmt.Date) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetEndDate(endDate strfmt.Date) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithStartDate(startDate strfmt.Date) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetStartDate(startDate strfmt.Date) {
	o.StartDate = startDate
}

// WithUdid adds the udid to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WithUdid(udid string) *FindComputerHardwareSoftwareReportsByUDIDParams {
	o.SetUdid(udid)
	return o
}

// SetUdid adds the udid to the find computer hardware software reports by u d ID params
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) SetUdid(udid string) {
	o.Udid = udid
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerHardwareSoftwareReportsByUDIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param end_date
	if err := r.SetPathParam("end_date", o.EndDate.String()); err != nil {
		return err
	}

	// path param start_date
	if err := r.SetPathParam("start_date", o.StartDate.String()); err != nil {
		return err
	}

	// path param udid
	if err := r.SetPathParam("udid", o.Udid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}