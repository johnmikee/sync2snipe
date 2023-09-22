// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePolicyByNameReader is a Reader for the UpdatePolicyByName structure.
type UpdatePolicyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePolicyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdatePolicyByNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /policies/name/{name}] updatePolicyByName", response, response.Code())
	}
}

// NewUpdatePolicyByNameCreated creates a UpdatePolicyByNameCreated with default headers values
func NewUpdatePolicyByNameCreated() *UpdatePolicyByNameCreated {
	return &UpdatePolicyByNameCreated{}
}

/*
UpdatePolicyByNameCreated describes a response with status code 201, with default header values.

Created
*/
type UpdatePolicyByNameCreated struct {
}

// IsSuccess returns true when this update policy by name created response has a 2xx status code
func (o *UpdatePolicyByNameCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update policy by name created response has a 3xx status code
func (o *UpdatePolicyByNameCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update policy by name created response has a 4xx status code
func (o *UpdatePolicyByNameCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update policy by name created response has a 5xx status code
func (o *UpdatePolicyByNameCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update policy by name created response a status code equal to that given
func (o *UpdatePolicyByNameCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update policy by name created response
func (o *UpdatePolicyByNameCreated) Code() int {
	return 201
}

func (o *UpdatePolicyByNameCreated) Error() string {
	return fmt.Sprintf("[PUT /policies/name/{name}][%d] updatePolicyByNameCreated ", 201)
}

func (o *UpdatePolicyByNameCreated) String() string {
	return fmt.Sprintf("[PUT /policies/name/{name}][%d] updatePolicyByNameCreated ", 201)
}

func (o *UpdatePolicyByNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}