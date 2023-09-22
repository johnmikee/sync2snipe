// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDirectoryBindingByIDReader is a Reader for the UpdateDirectoryBindingByID structure.
type UpdateDirectoryBindingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDirectoryBindingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateDirectoryBindingByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /directorybindings/id/{id}] updateDirectoryBindingById", response, response.Code())
	}
}

// NewUpdateDirectoryBindingByIDCreated creates a UpdateDirectoryBindingByIDCreated with default headers values
func NewUpdateDirectoryBindingByIDCreated() *UpdateDirectoryBindingByIDCreated {
	return &UpdateDirectoryBindingByIDCreated{}
}

/*
UpdateDirectoryBindingByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateDirectoryBindingByIDCreated struct {
}

// IsSuccess returns true when this update directory binding by Id created response has a 2xx status code
func (o *UpdateDirectoryBindingByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update directory binding by Id created response has a 3xx status code
func (o *UpdateDirectoryBindingByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update directory binding by Id created response has a 4xx status code
func (o *UpdateDirectoryBindingByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update directory binding by Id created response has a 5xx status code
func (o *UpdateDirectoryBindingByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update directory binding by Id created response a status code equal to that given
func (o *UpdateDirectoryBindingByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update directory binding by Id created response
func (o *UpdateDirectoryBindingByIDCreated) Code() int {
	return 201
}

func (o *UpdateDirectoryBindingByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /directorybindings/id/{id}][%d] updateDirectoryBindingByIdCreated ", 201)
}

func (o *UpdateDirectoryBindingByIDCreated) String() string {
	return fmt.Sprintf("[PUT /directorybindings/id/{id}][%d] updateDirectoryBindingByIdCreated ", 201)
}

func (o *UpdateDirectoryBindingByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
