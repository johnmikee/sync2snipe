// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserByIDReader is a Reader for the UpdateUserByID structure.
type UpdateUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateUserByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /users/id/{id}] updateUserById", response, response.Code())
	}
}

// NewUpdateUserByIDCreated creates a UpdateUserByIDCreated with default headers values
func NewUpdateUserByIDCreated() *UpdateUserByIDCreated {
	return &UpdateUserByIDCreated{}
}

/*
UpdateUserByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateUserByIDCreated struct {
}

// IsSuccess returns true when this update user by Id created response has a 2xx status code
func (o *UpdateUserByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user by Id created response has a 3xx status code
func (o *UpdateUserByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user by Id created response has a 4xx status code
func (o *UpdateUserByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user by Id created response has a 5xx status code
func (o *UpdateUserByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update user by Id created response a status code equal to that given
func (o *UpdateUserByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update user by Id created response
func (o *UpdateUserByIDCreated) Code() int {
	return 201
}

func (o *UpdateUserByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /users/id/{id}][%d] updateUserByIdCreated ", 201)
}

func (o *UpdateUserByIDCreated) String() string {
	return fmt.Sprintf("[PUT /users/id/{id}][%d] updateUserByIdCreated ", 201)
}

func (o *UpdateUserByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}