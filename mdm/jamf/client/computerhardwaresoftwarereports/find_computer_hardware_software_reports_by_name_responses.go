// Code generated by go-swagger; DO NOT EDIT.

package computerhardwaresoftwarereports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerHardwareSoftwareReportsByNameReader is a Reader for the FindComputerHardwareSoftwareReportsByName structure.
type FindComputerHardwareSoftwareReportsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerHardwareSoftwareReportsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerHardwareSoftwareReportsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerhardwaresoftwarereports/name/{name}/{start_date}_{end_date}] findComputerHardwareSoftwareReportsByName", response, response.Code())
	}
}

// NewFindComputerHardwareSoftwareReportsByNameOK creates a FindComputerHardwareSoftwareReportsByNameOK with default headers values
func NewFindComputerHardwareSoftwareReportsByNameOK() *FindComputerHardwareSoftwareReportsByNameOK {
	return &FindComputerHardwareSoftwareReportsByNameOK{}
}

/*
FindComputerHardwareSoftwareReportsByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerHardwareSoftwareReportsByNameOK struct {
	Payload *models.ComputerHardwareSoftwareReports
}

// IsSuccess returns true when this find computer hardware software reports by name o k response has a 2xx status code
func (o *FindComputerHardwareSoftwareReportsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer hardware software reports by name o k response has a 3xx status code
func (o *FindComputerHardwareSoftwareReportsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer hardware software reports by name o k response has a 4xx status code
func (o *FindComputerHardwareSoftwareReportsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer hardware software reports by name o k response has a 5xx status code
func (o *FindComputerHardwareSoftwareReportsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer hardware software reports by name o k response a status code equal to that given
func (o *FindComputerHardwareSoftwareReportsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer hardware software reports by name o k response
func (o *FindComputerHardwareSoftwareReportsByNameOK) Code() int {
	return 200
}

func (o *FindComputerHardwareSoftwareReportsByNameOK) Error() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/name/{name}/{start_date}_{end_date}][%d] findComputerHardwareSoftwareReportsByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByNameOK) String() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/name/{name}/{start_date}_{end_date}][%d] findComputerHardwareSoftwareReportsByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByNameOK) GetPayload() *models.ComputerHardwareSoftwareReports {
	return o.Payload
}

func (o *FindComputerHardwareSoftwareReportsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerHardwareSoftwareReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
