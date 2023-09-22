// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindLDAPServerUserByIDReader is a Reader for the FindLDAPServerUserByID structure.
type FindLDAPServerUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindLDAPServerUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindLDAPServerUserByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ldapservers/id/{id}/user/{user}] findLDAPServerUserById", response, response.Code())
	}
}

// NewFindLDAPServerUserByIDOK creates a FindLDAPServerUserByIDOK with default headers values
func NewFindLDAPServerUserByIDOK() *FindLDAPServerUserByIDOK {
	return &FindLDAPServerUserByIDOK{}
}

/*
FindLDAPServerUserByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindLDAPServerUserByIDOK struct {
}

// IsSuccess returns true when this find l d a p server user by Id o k response has a 2xx status code
func (o *FindLDAPServerUserByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find l d a p server user by Id o k response has a 3xx status code
func (o *FindLDAPServerUserByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find l d a p server user by Id o k response has a 4xx status code
func (o *FindLDAPServerUserByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find l d a p server user by Id o k response has a 5xx status code
func (o *FindLDAPServerUserByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find l d a p server user by Id o k response a status code equal to that given
func (o *FindLDAPServerUserByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find l d a p server user by Id o k response
func (o *FindLDAPServerUserByIDOK) Code() int {
	return 200
}

func (o *FindLDAPServerUserByIDOK) Error() string {
	return fmt.Sprintf("[GET /ldapservers/id/{id}/user/{user}][%d] findLDAPServerUserByIdOK ", 200)
}

func (o *FindLDAPServerUserByIDOK) String() string {
	return fmt.Sprintf("[GET /ldapservers/id/{id}/user/{user}][%d] findLDAPServerUserByIdOK ", 200)
}

func (o *FindLDAPServerUserByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
