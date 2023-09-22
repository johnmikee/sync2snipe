// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteLDAPServerByNameReader is a Reader for the DeleteLDAPServerByName structure.
type DeleteLDAPServerByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLDAPServerByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLDAPServerByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ldapservers/name/{name}] deleteLDAPServerByName", response, response.Code())
	}
}

// NewDeleteLDAPServerByNameOK creates a DeleteLDAPServerByNameOK with default headers values
func NewDeleteLDAPServerByNameOK() *DeleteLDAPServerByNameOK {
	return &DeleteLDAPServerByNameOK{}
}

/*
DeleteLDAPServerByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteLDAPServerByNameOK struct {
}

// IsSuccess returns true when this delete l d a p server by name o k response has a 2xx status code
func (o *DeleteLDAPServerByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete l d a p server by name o k response has a 3xx status code
func (o *DeleteLDAPServerByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete l d a p server by name o k response has a 4xx status code
func (o *DeleteLDAPServerByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete l d a p server by name o k response has a 5xx status code
func (o *DeleteLDAPServerByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete l d a p server by name o k response a status code equal to that given
func (o *DeleteLDAPServerByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete l d a p server by name o k response
func (o *DeleteLDAPServerByNameOK) Code() int {
	return 200
}

func (o *DeleteLDAPServerByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /ldapservers/name/{name}][%d] deleteLDAPServerByNameOK ", 200)
}

func (o *DeleteLDAPServerByNameOK) String() string {
	return fmt.Sprintf("[DELETE /ldapservers/name/{name}][%d] deleteLDAPServerByNameOK ", 200)
}

func (o *DeleteLDAPServerByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}