// Code generated by go-swagger; DO NOT EDIT.

package computerreports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerReportsReader is a Reader for the FindComputerReports structure.
type FindComputerReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerreports] findComputerReports", response, response.Code())
	}
}

// NewFindComputerReportsOK creates a FindComputerReportsOK with default headers values
func NewFindComputerReportsOK() *FindComputerReportsOK {
	return &FindComputerReportsOK{}
}

/*
FindComputerReportsOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerReportsOK struct {
	Payload models.ComputerReports
}

// IsSuccess returns true when this find computer reports o k response has a 2xx status code
func (o *FindComputerReportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer reports o k response has a 3xx status code
func (o *FindComputerReportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer reports o k response has a 4xx status code
func (o *FindComputerReportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer reports o k response has a 5xx status code
func (o *FindComputerReportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer reports o k response a status code equal to that given
func (o *FindComputerReportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer reports o k response
func (o *FindComputerReportsOK) Code() int {
	return 200
}

func (o *FindComputerReportsOK) Error() string {
	return fmt.Sprintf("[GET /computerreports][%d] findComputerReportsOK  %+v", 200, o.Payload)
}

func (o *FindComputerReportsOK) String() string {
	return fmt.Sprintf("[GET /computerreports][%d] findComputerReportsOK  %+v", 200, o.Payload)
}

func (o *FindComputerReportsOK) GetPayload() models.ComputerReports {
	return o.Payload
}

func (o *FindComputerReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
