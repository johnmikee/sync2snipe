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

// FindComputerextensionattributesByNameReader is a Reader for the FindComputerextensionattributesByName structure.
type FindComputerextensionattributesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerextensionattributesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerextensionattributesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerextensionattributes/name/{name}] findComputerextensionattributesByName", response, response.Code())
	}
}

// NewFindComputerextensionattributesByNameOK creates a FindComputerextensionattributesByNameOK with default headers values
func NewFindComputerextensionattributesByNameOK() *FindComputerextensionattributesByNameOK {
	return &FindComputerextensionattributesByNameOK{}
}

/*
FindComputerextensionattributesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerextensionattributesByNameOK struct {
	Payload *models.ComputerExtensionAttribute
}

// IsSuccess returns true when this find computerextensionattributes by name o k response has a 2xx status code
func (o *FindComputerextensionattributesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computerextensionattributes by name o k response has a 3xx status code
func (o *FindComputerextensionattributesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computerextensionattributes by name o k response has a 4xx status code
func (o *FindComputerextensionattributesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computerextensionattributes by name o k response has a 5xx status code
func (o *FindComputerextensionattributesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computerextensionattributes by name o k response a status code equal to that given
func (o *FindComputerextensionattributesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computerextensionattributes by name o k response
func (o *FindComputerextensionattributesByNameOK) Code() int {
	return 200
}

func (o *FindComputerextensionattributesByNameOK) Error() string {
	return fmt.Sprintf("[GET /computerextensionattributes/name/{name}][%d] findComputerextensionattributesByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerextensionattributesByNameOK) String() string {
	return fmt.Sprintf("[GET /computerextensionattributes/name/{name}][%d] findComputerextensionattributesByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerextensionattributesByNameOK) GetPayload() *models.ComputerExtensionAttribute {
	return o.Payload
}

func (o *FindComputerextensionattributesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerExtensionAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
