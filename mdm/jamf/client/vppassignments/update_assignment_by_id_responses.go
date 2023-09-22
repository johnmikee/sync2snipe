// Code generated by go-swagger; DO NOT EDIT.

package vppassignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAssignmentByIDReader is a Reader for the UpdateAssignmentByID structure.
type UpdateAssignmentByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAssignmentByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAssignmentByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /vppassignments/id/{id}] updateAssignmentById", response, response.Code())
	}
}

// NewUpdateAssignmentByIDCreated creates a UpdateAssignmentByIDCreated with default headers values
func NewUpdateAssignmentByIDCreated() *UpdateAssignmentByIDCreated {
	return &UpdateAssignmentByIDCreated{}
}

/*
UpdateAssignmentByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAssignmentByIDCreated struct {
}

// IsSuccess returns true when this update assignment by Id created response has a 2xx status code
func (o *UpdateAssignmentByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update assignment by Id created response has a 3xx status code
func (o *UpdateAssignmentByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update assignment by Id created response has a 4xx status code
func (o *UpdateAssignmentByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update assignment by Id created response has a 5xx status code
func (o *UpdateAssignmentByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update assignment by Id created response a status code equal to that given
func (o *UpdateAssignmentByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update assignment by Id created response
func (o *UpdateAssignmentByIDCreated) Code() int {
	return 201
}

func (o *UpdateAssignmentByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /vppassignments/id/{id}][%d] updateAssignmentByIdCreated ", 201)
}

func (o *UpdateAssignmentByIDCreated) String() string {
	return fmt.Sprintf("[PUT /vppassignments/id/{id}][%d] updateAssignmentByIdCreated ", 201)
}

func (o *UpdateAssignmentByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}