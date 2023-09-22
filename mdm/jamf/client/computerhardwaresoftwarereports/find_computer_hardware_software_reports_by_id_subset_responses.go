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

// FindComputerHardwareSoftwareReportsByIDSubsetReader is a Reader for the FindComputerHardwareSoftwareReportsByIDSubset structure.
type FindComputerHardwareSoftwareReportsByIDSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerHardwareSoftwareReportsByIDSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerHardwareSoftwareReportsByIDSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerhardwaresoftwarereports/id/{id}/{start_date}_{end_date}/subset/{subset}] findComputerHardwareSoftwareReportsByIdSubset", response, response.Code())
	}
}

// NewFindComputerHardwareSoftwareReportsByIDSubsetOK creates a FindComputerHardwareSoftwareReportsByIDSubsetOK with default headers values
func NewFindComputerHardwareSoftwareReportsByIDSubsetOK() *FindComputerHardwareSoftwareReportsByIDSubsetOK {
	return &FindComputerHardwareSoftwareReportsByIDSubsetOK{}
}

/*
FindComputerHardwareSoftwareReportsByIDSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerHardwareSoftwareReportsByIDSubsetOK struct {
	Payload *models.ComputerHardwareSoftwareReports
}

// IsSuccess returns true when this find computer hardware software reports by Id subset o k response has a 2xx status code
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer hardware software reports by Id subset o k response has a 3xx status code
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer hardware software reports by Id subset o k response has a 4xx status code
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer hardware software reports by Id subset o k response has a 5xx status code
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer hardware software reports by Id subset o k response a status code equal to that given
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer hardware software reports by Id subset o k response
func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) Code() int {
	return 200
}

func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) Error() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/id/{id}/{start_date}_{end_date}/subset/{subset}][%d] findComputerHardwareSoftwareReportsByIdSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) String() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/id/{id}/{start_date}_{end_date}/subset/{subset}][%d] findComputerHardwareSoftwareReportsByIdSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) GetPayload() *models.ComputerHardwareSoftwareReports {
	return o.Payload
}

func (o *FindComputerHardwareSoftwareReportsByIDSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerHardwareSoftwareReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}