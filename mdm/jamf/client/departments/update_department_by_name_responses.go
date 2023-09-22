// Code generated by go-swagger; DO NOT EDIT.

package departments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDepartmentByNameReader is a Reader for the UpdateDepartmentByName structure.
type UpdateDepartmentByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDepartmentByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateDepartmentByNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /departments/name/{name}] updateDepartmentByName", response, response.Code())
	}
}

// NewUpdateDepartmentByNameCreated creates a UpdateDepartmentByNameCreated with default headers values
func NewUpdateDepartmentByNameCreated() *UpdateDepartmentByNameCreated {
	return &UpdateDepartmentByNameCreated{}
}

/*
UpdateDepartmentByNameCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateDepartmentByNameCreated struct {
}

// IsSuccess returns true when this update department by name created response has a 2xx status code
func (o *UpdateDepartmentByNameCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update department by name created response has a 3xx status code
func (o *UpdateDepartmentByNameCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update department by name created response has a 4xx status code
func (o *UpdateDepartmentByNameCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update department by name created response has a 5xx status code
func (o *UpdateDepartmentByNameCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update department by name created response a status code equal to that given
func (o *UpdateDepartmentByNameCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update department by name created response
func (o *UpdateDepartmentByNameCreated) Code() int {
	return 201
}

func (o *UpdateDepartmentByNameCreated) Error() string {
	return fmt.Sprintf("[PUT /departments/name/{name}][%d] updateDepartmentByNameCreated ", 201)
}

func (o *UpdateDepartmentByNameCreated) String() string {
	return fmt.Sprintf("[PUT /departments/name/{name}][%d] updateDepartmentByNameCreated ", 201)
}

func (o *UpdateDepartmentByNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}