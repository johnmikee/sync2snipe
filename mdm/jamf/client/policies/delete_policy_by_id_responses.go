// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePolicyByIDReader is a Reader for the DeletePolicyByID structure.
type DeletePolicyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /policies/id/{id}] deletePolicyById", response, response.Code())
	}
}

// NewDeletePolicyByIDOK creates a DeletePolicyByIDOK with default headers values
func NewDeletePolicyByIDOK() *DeletePolicyByIDOK {
	return &DeletePolicyByIDOK{}
}

/*
DeletePolicyByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeletePolicyByIDOK struct {
}

// IsSuccess returns true when this delete policy by Id o k response has a 2xx status code
func (o *DeletePolicyByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete policy by Id o k response has a 3xx status code
func (o *DeletePolicyByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete policy by Id o k response has a 4xx status code
func (o *DeletePolicyByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete policy by Id o k response has a 5xx status code
func (o *DeletePolicyByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete policy by Id o k response a status code equal to that given
func (o *DeletePolicyByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete policy by Id o k response
func (o *DeletePolicyByIDOK) Code() int {
	return 200
}

func (o *DeletePolicyByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /policies/id/{id}][%d] deletePolicyByIdOK ", 200)
}

func (o *DeletePolicyByIDOK) String() string {
	return fmt.Sprintf("[DELETE /policies/id/{id}][%d] deletePolicyByIdOK ", 200)
}

func (o *DeletePolicyByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}