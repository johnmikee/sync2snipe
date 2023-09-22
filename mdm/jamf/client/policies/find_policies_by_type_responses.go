// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindPoliciesByTypeReader is a Reader for the FindPoliciesByType structure.
type FindPoliciesByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPoliciesByTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPoliciesByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /policies/createdBy/{createdBy}] findPoliciesByType", response, response.Code())
	}
}

// NewFindPoliciesByTypeOK creates a FindPoliciesByTypeOK with default headers values
func NewFindPoliciesByTypeOK() *FindPoliciesByTypeOK {
	return &FindPoliciesByTypeOK{}
}

/*
FindPoliciesByTypeOK describes a response with status code 200, with default header values.

OK
*/
type FindPoliciesByTypeOK struct {
}

// IsSuccess returns true when this find policies by type o k response has a 2xx status code
func (o *FindPoliciesByTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find policies by type o k response has a 3xx status code
func (o *FindPoliciesByTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find policies by type o k response has a 4xx status code
func (o *FindPoliciesByTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find policies by type o k response has a 5xx status code
func (o *FindPoliciesByTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find policies by type o k response a status code equal to that given
func (o *FindPoliciesByTypeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find policies by type o k response
func (o *FindPoliciesByTypeOK) Code() int {
	return 200
}

func (o *FindPoliciesByTypeOK) Error() string {
	return fmt.Sprintf("[GET /policies/createdBy/{createdBy}][%d] findPoliciesByTypeOK ", 200)
}

func (o *FindPoliciesByTypeOK) String() string {
	return fmt.Sprintf("[GET /policies/createdBy/{createdBy}][%d] findPoliciesByTypeOK ", 200)
}

func (o *FindPoliciesByTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
