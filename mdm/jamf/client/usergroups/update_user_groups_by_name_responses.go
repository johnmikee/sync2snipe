// Code generated by go-swagger; DO NOT EDIT.

package usergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserGroupsByNameReader is a Reader for the UpdateUserGroupsByName structure.
type UpdateUserGroupsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGroupsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateUserGroupsByNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /usergroups/name/{name}] updateUserGroupsByName", response, response.Code())
	}
}

// NewUpdateUserGroupsByNameCreated creates a UpdateUserGroupsByNameCreated with default headers values
func NewUpdateUserGroupsByNameCreated() *UpdateUserGroupsByNameCreated {
	return &UpdateUserGroupsByNameCreated{}
}

/*
UpdateUserGroupsByNameCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateUserGroupsByNameCreated struct {
}

// IsSuccess returns true when this update user groups by name created response has a 2xx status code
func (o *UpdateUserGroupsByNameCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user groups by name created response has a 3xx status code
func (o *UpdateUserGroupsByNameCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user groups by name created response has a 4xx status code
func (o *UpdateUserGroupsByNameCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user groups by name created response has a 5xx status code
func (o *UpdateUserGroupsByNameCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update user groups by name created response a status code equal to that given
func (o *UpdateUserGroupsByNameCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update user groups by name created response
func (o *UpdateUserGroupsByNameCreated) Code() int {
	return 201
}

func (o *UpdateUserGroupsByNameCreated) Error() string {
	return fmt.Sprintf("[PUT /usergroups/name/{name}][%d] updateUserGroupsByNameCreated ", 201)
}

func (o *UpdateUserGroupsByNameCreated) String() string {
	return fmt.Sprintf("[PUT /usergroups/name/{name}][%d] updateUserGroupsByNameCreated ", 201)
}

func (o *UpdateUserGroupsByNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
