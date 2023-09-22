// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateLDAPServerByIDReader is a Reader for the UpdateLDAPServerByID structure.
type UpdateLDAPServerByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLDAPServerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateLDAPServerByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ldapservers/id/{id}] updateLDAPServerById", response, response.Code())
	}
}

// NewUpdateLDAPServerByIDCreated creates a UpdateLDAPServerByIDCreated with default headers values
func NewUpdateLDAPServerByIDCreated() *UpdateLDAPServerByIDCreated {
	return &UpdateLDAPServerByIDCreated{}
}

/*
UpdateLDAPServerByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateLDAPServerByIDCreated struct {
}

// IsSuccess returns true when this update l d a p server by Id created response has a 2xx status code
func (o *UpdateLDAPServerByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update l d a p server by Id created response has a 3xx status code
func (o *UpdateLDAPServerByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update l d a p server by Id created response has a 4xx status code
func (o *UpdateLDAPServerByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update l d a p server by Id created response has a 5xx status code
func (o *UpdateLDAPServerByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update l d a p server by Id created response a status code equal to that given
func (o *UpdateLDAPServerByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update l d a p server by Id created response
func (o *UpdateLDAPServerByIDCreated) Code() int {
	return 201
}

func (o *UpdateLDAPServerByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /ldapservers/id/{id}][%d] updateLDAPServerByIdCreated ", 201)
}

func (o *UpdateLDAPServerByIDCreated) String() string {
	return fmt.Sprintf("[PUT /ldapservers/id/{id}][%d] updateLDAPServerByIdCreated ", 201)
}

func (o *UpdateLDAPServerByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}