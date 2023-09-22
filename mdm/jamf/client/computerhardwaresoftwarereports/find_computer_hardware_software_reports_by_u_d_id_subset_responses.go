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

// FindComputerHardwareSoftwareReportsByUDIDSubsetReader is a Reader for the FindComputerHardwareSoftwareReportsByUDIDSubset structure.
type FindComputerHardwareSoftwareReportsByUDIDSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerHardwareSoftwareReportsByUDIDSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerhardwaresoftwarereports/udid/{udid}/{start_date}_{end_date}/subset/{subset}] findComputerHardwareSoftwareReportsByUDIDSubset", response, response.Code())
	}
}

// NewFindComputerHardwareSoftwareReportsByUDIDSubsetOK creates a FindComputerHardwareSoftwareReportsByUDIDSubsetOK with default headers values
func NewFindComputerHardwareSoftwareReportsByUDIDSubsetOK() *FindComputerHardwareSoftwareReportsByUDIDSubsetOK {
	return &FindComputerHardwareSoftwareReportsByUDIDSubsetOK{}
}

/*
FindComputerHardwareSoftwareReportsByUDIDSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerHardwareSoftwareReportsByUDIDSubsetOK struct {
	Payload *models.ComputerHardwareSoftwareReports
}

// IsSuccess returns true when this find computer hardware software reports by u d Id subset o k response has a 2xx status code
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer hardware software reports by u d Id subset o k response has a 3xx status code
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer hardware software reports by u d Id subset o k response has a 4xx status code
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer hardware software reports by u d Id subset o k response has a 5xx status code
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer hardware software reports by u d Id subset o k response a status code equal to that given
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer hardware software reports by u d Id subset o k response
func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) Code() int {
	return 200
}

func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) Error() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/udid/{udid}/{start_date}_{end_date}/subset/{subset}][%d] findComputerHardwareSoftwareReportsByUDIdSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) String() string {
	return fmt.Sprintf("[GET /computerhardwaresoftwarereports/udid/{udid}/{start_date}_{end_date}/subset/{subset}][%d] findComputerHardwareSoftwareReportsByUDIdSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) GetPayload() *models.ComputerHardwareSoftwareReports {
	return o.Payload
}

func (o *FindComputerHardwareSoftwareReportsByUDIDSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerHardwareSoftwareReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
