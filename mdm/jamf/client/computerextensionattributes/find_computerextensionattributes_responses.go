// Code generated by go-swagger; DO NOT EDIT.

package computerextensionattributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerextensionattributesReader is a Reader for the FindComputerextensionattributes structure.
type FindComputerextensionattributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerextensionattributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerextensionattributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerextensionattributes] findComputerextensionattributes", response, response.Code())
	}
}

// NewFindComputerextensionattributesOK creates a FindComputerextensionattributesOK with default headers values
func NewFindComputerextensionattributesOK() *FindComputerextensionattributesOK {
	return &FindComputerextensionattributesOK{}
}

/*
FindComputerextensionattributesOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerextensionattributesOK struct {
	Payload models.ComputerExtensionAttributes
}

// IsSuccess returns true when this find computerextensionattributes o k response has a 2xx status code
func (o *FindComputerextensionattributesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computerextensionattributes o k response has a 3xx status code
func (o *FindComputerextensionattributesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computerextensionattributes o k response has a 4xx status code
func (o *FindComputerextensionattributesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computerextensionattributes o k response has a 5xx status code
func (o *FindComputerextensionattributesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computerextensionattributes o k response a status code equal to that given
func (o *FindComputerextensionattributesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computerextensionattributes o k response
func (o *FindComputerextensionattributesOK) Code() int {
	return 200
}

func (o *FindComputerextensionattributesOK) Error() string {
	return fmt.Sprintf("[GET /computerextensionattributes][%d] findComputerextensionattributesOK  %+v", 200, o.Payload)
}

func (o *FindComputerextensionattributesOK) String() string {
	return fmt.Sprintf("[GET /computerextensionattributes][%d] findComputerextensionattributesOK  %+v", 200, o.Payload)
}

func (o *FindComputerextensionattributesOK) GetPayload() models.ComputerExtensionAttributes {
	return o.Payload
}

func (o *FindComputerextensionattributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
