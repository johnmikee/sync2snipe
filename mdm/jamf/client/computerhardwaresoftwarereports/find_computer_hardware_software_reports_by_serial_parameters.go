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

// NewFindComputerHardwareSoftwareReportsBySerialParams creates a new FindComputerHardwareSoftwareReportsBySerialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerHardwareSoftwareReportsBySerialParams() *FindComputerHardwareSoftwareReportsBySerialParams {
	return &FindComputerHardwareSoftwareReportsBySerialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerHardwareSoftwareReportsBySerialParamsWithTimeout creates a new FindComputerHardwareSoftwareReportsBySerialParams object
// with the ability to set a timeout on a request.
func NewFindComputerHardwareSoftwareReportsBySerialParamsWithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsBySerialParams {
	return &FindComputerHardwareSoftwareReportsBySerialParams{
		timeout: timeout,
	}
}

// NewFindComputerHardwareSoftwareReportsBySerialParamsWithContext creates a new FindComputerHardwareSoftwareReportsBySerialParams object
// with the ability to set a context for a request.
func NewFindComputerHardwareSoftwareReportsBySerialParamsWithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsBySerialParams {
	return &FindComputerHardwareSoftwareReportsBySerialParams{
		Context: ctx,
	}
}

// NewFindComputerHardwareSoftwareReportsBySerialParamsWithHTTPClient creates a new FindComputerHardwareSoftwareReportsBySerialParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerHardwareSoftwareReportsBySerialParamsWithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsBySerialParams {
	return &FindComputerHardwareSoftwareReportsBySerialParams{
		HTTPClient: client,
	}
}

/*
FindComputerHardwareSoftwareReportsBySerialParams contains all the parameters to send to the API endpoint

	for the find computer hardware software reports by serial operation.

	Typically these are written to a http.Request.
*/
type FindComputerHardwareSoftwareReportsBySerialParams struct {

	/* EndDate.

	   End date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	EndDate strfmt.Date

	/* Serialnumber.

	   Serial number to filter by
	*/
	Serialnumber string

	/* StartDate.

	   Start date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	StartDate strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer hardware software reports by serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithDefaults() *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer hardware software reports by serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithEndDate(endDate strfmt.Date) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetEndDate(endDate strfmt.Date) {
	o.EndDate = endDate
}

// WithSerialnumber adds the serialnumber to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithSerialnumber(serialnumber string) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetSerialnumber(serialnumber)
	return o
}

// SetSerialnumber adds the serialnumber to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetSerialnumber(serialnumber string) {
	o.Serialnumber = serialnumber
}

// WithStartDate adds the startDate to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WithStartDate(startDate strfmt.Date) *FindComputerHardwareSoftwareReportsBySerialParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the find computer hardware software reports by serial params
func (o *FindComputerHardwareSoftwareReportsBySerialParams) SetStartDate(startDate strfmt.Date) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerHardwareSoftwareReportsBySerialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param end_date
	if err := r.SetPathParam("end_date", o.EndDate.String()); err != nil {
		return err
	}

	// path param serialnumber
	if err := r.SetPathParam("serialnumber", o.Serialnumber); err != nil {
		return err
	}

	// path param start_date
	if err := r.SetPathParam("start_date", o.StartDate.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}