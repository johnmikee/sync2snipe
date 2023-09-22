// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePolicyByIDReader is a Reader for the CreatePolicyByID structure.
type CreatePolicyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePolicyByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /policies/id/{id}] createPolicyById", response, response.Code())
	}
}

// NewCreatePolicyByIDCreated creates a CreatePolicyByIDCreated with default headers values
func NewCreatePolicyByIDCreated() *CreatePolicyByIDCreated {
	return &CreatePolicyByIDCreated{}
}

/*
CreatePolicyByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePolicyByIDCreated struct {
}

// IsSuccess returns true when this create policy by Id created response has a 2xx status code
func (o *CreatePolicyByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create policy by Id created response has a 3xx status code
func (o *CreatePolicyByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy by Id created response has a 4xx status code
func (o *CreatePolicyByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy by Id created response has a 5xx status code
func (o *CreatePolicyByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy by Id created response a status code equal to that given
func (o *CreatePolicyByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create policy by Id created response
func (o *CreatePolicyByIDCreated) Code() int {
	return 201
}

func (o *CreatePolicyByIDCreated) Error() string {
	return fmt.Sprintf("[POST /policies/id/{id}][%d] createPolicyByIdCreated ", 201)
}

func (o *CreatePolicyByIDCreated) String() string {
	return fmt.Sprintf("[POST /policies/id/{id}][%d] createPolicyByIdCreated ", 201)
}

func (o *CreatePolicyByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
