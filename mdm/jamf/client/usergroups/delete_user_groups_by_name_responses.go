// Code generated by go-swagger; DO NOT EDIT.

package usergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserGroupsByNameReader is a Reader for the DeleteUserGroupsByName structure.
type DeleteUserGroupsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserGroupsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserGroupsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /usergroups/name/{name}] deleteUserGroupsByName", response, response.Code())
	}
}

// NewDeleteUserGroupsByNameOK creates a DeleteUserGroupsByNameOK with default headers values
func NewDeleteUserGroupsByNameOK() *DeleteUserGroupsByNameOK {
	return &DeleteUserGroupsByNameOK{}
}

/*
DeleteUserGroupsByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserGroupsByNameOK struct {
}

// IsSuccess returns true when this delete user groups by name o k response has a 2xx status code
func (o *DeleteUserGroupsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user groups by name o k response has a 3xx status code
func (o *DeleteUserGroupsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user groups by name o k response has a 4xx status code
func (o *DeleteUserGroupsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user groups by name o k response has a 5xx status code
func (o *DeleteUserGroupsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user groups by name o k response a status code equal to that given
func (o *DeleteUserGroupsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete user groups by name o k response
func (o *DeleteUserGroupsByNameOK) Code() int {
	return 200
}

func (o *DeleteUserGroupsByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/name/{name}][%d] deleteUserGroupsByNameOK ", 200)
}

func (o *DeleteUserGroupsByNameOK) String() string {
	return fmt.Sprintf("[DELETE /usergroups/name/{name}][%d] deleteUserGroupsByNameOK ", 200)
}

func (o *DeleteUserGroupsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}