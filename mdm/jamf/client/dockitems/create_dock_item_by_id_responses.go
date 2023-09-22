// Code generated by go-swagger; DO NOT EDIT.

package dockitems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateDockItemByIDReader is a Reader for the CreateDockItemByID structure.
type CreateDockItemByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDockItemByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDockItemByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dockitems/id/{id}] createDockItemById", response, response.Code())
	}
}

// NewCreateDockItemByIDCreated creates a CreateDockItemByIDCreated with default headers values
func NewCreateDockItemByIDCreated() *CreateDockItemByIDCreated {
	return &CreateDockItemByIDCreated{}
}

/*
CreateDockItemByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateDockItemByIDCreated struct {
}

// IsSuccess returns true when this create dock item by Id created response has a 2xx status code
func (o *CreateDockItemByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dock item by Id created response has a 3xx status code
func (o *CreateDockItemByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dock item by Id created response has a 4xx status code
func (o *CreateDockItemByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dock item by Id created response has a 5xx status code
func (o *CreateDockItemByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create dock item by Id created response a status code equal to that given
func (o *CreateDockItemByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create dock item by Id created response
func (o *CreateDockItemByIDCreated) Code() int {
	return 201
}

func (o *CreateDockItemByIDCreated) Error() string {
	return fmt.Sprintf("[POST /dockitems/id/{id}][%d] createDockItemByIdCreated ", 201)
}

func (o *CreateDockItemByIDCreated) String() string {
	return fmt.Sprintf("[POST /dockitems/id/{id}][%d] createDockItemByIdCreated ", 201)
}

func (o *CreateDockItemByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
