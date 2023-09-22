// Code generated by go-swagger; DO NOT EDIT.

package peripherals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindPeripheralsReader is a Reader for the FindPeripherals structure.
type FindPeripheralsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPeripheralsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPeripheralsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /peripherals] findPeripherals", response, response.Code())
	}
}

// NewFindPeripheralsOK creates a FindPeripheralsOK with default headers values
func NewFindPeripheralsOK() *FindPeripheralsOK {
	return &FindPeripheralsOK{}
}

/*
FindPeripheralsOK describes a response with status code 200, with default header values.

OK
*/
type FindPeripheralsOK struct {
	Payload models.Peripherals
}

// IsSuccess returns true when this find peripherals o k response has a 2xx status code
func (o *FindPeripheralsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find peripherals o k response has a 3xx status code
func (o *FindPeripheralsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find peripherals o k response has a 4xx status code
func (o *FindPeripheralsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find peripherals o k response has a 5xx status code
func (o *FindPeripheralsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find peripherals o k response a status code equal to that given
func (o *FindPeripheralsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find peripherals o k response
func (o *FindPeripheralsOK) Code() int {
	return 200
}

func (o *FindPeripheralsOK) Error() string {
	return fmt.Sprintf("[GET /peripherals][%d] findPeripheralsOK  %+v", 200, o.Payload)
}

func (o *FindPeripheralsOK) String() string {
	return fmt.Sprintf("[GET /peripherals][%d] findPeripheralsOK  %+v", 200, o.Payload)
}

func (o *FindPeripheralsOK) GetPayload() models.Peripherals {
	return o.Payload
}

func (o *FindPeripheralsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}