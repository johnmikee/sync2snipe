// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindLDAPServerUserByNameReader is a Reader for the FindLDAPServerUserByName structure.
type FindLDAPServerUserByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindLDAPServerUserByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindLDAPServerUserByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ldapservers/name/{name}/user/{user}] findLDAPServerUserByName", response, response.Code())
	}
}

// NewFindLDAPServerUserByNameOK creates a FindLDAPServerUserByNameOK with default headers values
func NewFindLDAPServerUserByNameOK() *FindLDAPServerUserByNameOK {
	return &FindLDAPServerUserByNameOK{}
}

/*
FindLDAPServerUserByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindLDAPServerUserByNameOK struct {
}

// IsSuccess returns true when this find l d a p server user by name o k response has a 2xx status code
func (o *FindLDAPServerUserByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find l d a p server user by name o k response has a 3xx status code
func (o *FindLDAPServerUserByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find l d a p server user by name o k response has a 4xx status code
func (o *FindLDAPServerUserByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find l d a p server user by name o k response has a 5xx status code
func (o *FindLDAPServerUserByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find l d a p server user by name o k response a status code equal to that given
func (o *FindLDAPServerUserByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find l d a p server user by name o k response
func (o *FindLDAPServerUserByNameOK) Code() int {
	return 200
}

func (o *FindLDAPServerUserByNameOK) Error() string {
	return fmt.Sprintf("[GET /ldapservers/name/{name}/user/{user}][%d] findLDAPServerUserByNameOK ", 200)
}

func (o *FindLDAPServerUserByNameOK) String() string {
	return fmt.Sprintf("[GET /ldapservers/name/{name}/user/{user}][%d] findLDAPServerUserByNameOK ", 200)
}

func (o *FindLDAPServerUserByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}