// Code generated by go-swagger; DO NOT EDIT.

package classes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateClassByIDReader is a Reader for the CreateClassByID structure.
type CreateClassByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClassByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClassByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /classes/id/{id}] createClassById", response, response.Code())
	}
}

// NewCreateClassByIDCreated creates a CreateClassByIDCreated with default headers values
func NewCreateClassByIDCreated() *CreateClassByIDCreated {
	return &CreateClassByIDCreated{}
}

/*
CreateClassByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateClassByIDCreated struct {
}

// IsSuccess returns true when this create class by Id created response has a 2xx status code
func (o *CreateClassByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create class by Id created response has a 3xx status code
func (o *CreateClassByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create class by Id created response has a 4xx status code
func (o *CreateClassByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create class by Id created response has a 5xx status code
func (o *CreateClassByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create class by Id created response a status code equal to that given
func (o *CreateClassByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create class by Id created response
func (o *CreateClassByIDCreated) Code() int {
	return 201
}

func (o *CreateClassByIDCreated) Error() string {
	return fmt.Sprintf("[POST /classes/id/{id}][%d] createClassByIdCreated ", 201)
}

func (o *CreateClassByIDCreated) String() string {
	return fmt.Sprintf("[POST /classes/id/{id}][%d] createClassByIdCreated ", 201)
}

func (o *CreateClassByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}