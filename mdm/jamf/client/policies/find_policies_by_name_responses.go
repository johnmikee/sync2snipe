// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindPoliciesByNameReader is a Reader for the FindPoliciesByName structure.
type FindPoliciesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPoliciesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPoliciesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /policies/name/{name}] findPoliciesByName", response, response.Code())
	}
}

// NewFindPoliciesByNameOK creates a FindPoliciesByNameOK with default headers values
func NewFindPoliciesByNameOK() *FindPoliciesByNameOK {
	return &FindPoliciesByNameOK{}
}

/*
FindPoliciesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindPoliciesByNameOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this find policies by name o k response has a 2xx status code
func (o *FindPoliciesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find policies by name o k response has a 3xx status code
func (o *FindPoliciesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find policies by name o k response has a 4xx status code
func (o *FindPoliciesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find policies by name o k response has a 5xx status code
func (o *FindPoliciesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find policies by name o k response a status code equal to that given
func (o *FindPoliciesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find policies by name o k response
func (o *FindPoliciesByNameOK) Code() int {
	return 200
}

func (o *FindPoliciesByNameOK) Error() string {
	return fmt.Sprintf("[GET /policies/name/{name}][%d] findPoliciesByNameOK  %+v", 200, o.Payload)
}

func (o *FindPoliciesByNameOK) String() string {
	return fmt.Sprintf("[GET /policies/name/{name}][%d] findPoliciesByNameOK  %+v", 200, o.Payload)
}

func (o *FindPoliciesByNameOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *FindPoliciesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
