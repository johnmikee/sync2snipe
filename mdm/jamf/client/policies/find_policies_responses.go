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

// FindPoliciesReader is a Reader for the FindPolicies structure.
type FindPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /policies] findPolicies", response, response.Code())
	}
}

// NewFindPoliciesOK creates a FindPoliciesOK with default headers values
func NewFindPoliciesOK() *FindPoliciesOK {
	return &FindPoliciesOK{}
}

/*
FindPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type FindPoliciesOK struct {
	Payload models.Policies
}

// IsSuccess returns true when this find policies o k response has a 2xx status code
func (o *FindPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find policies o k response has a 3xx status code
func (o *FindPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find policies o k response has a 4xx status code
func (o *FindPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find policies o k response has a 5xx status code
func (o *FindPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find policies o k response a status code equal to that given
func (o *FindPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find policies o k response
func (o *FindPoliciesOK) Code() int {
	return 200
}

func (o *FindPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policies][%d] findPoliciesOK  %+v", 200, o.Payload)
}

func (o *FindPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policies][%d] findPoliciesOK  %+v", 200, o.Payload)
}

func (o *FindPoliciesOK) GetPayload() models.Policies {
	return o.Payload
}

func (o *FindPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}