// Code generated by go-swagger; DO NOT EDIT.

package advancedcomputersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateAdvancedComputerSearchgByIDReader is a Reader for the CreateAdvancedComputerSearchgByID structure.
type CreateAdvancedComputerSearchgByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAdvancedComputerSearchgByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAdvancedComputerSearchgByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /advancedcomputersearches/id/{id}] createAdvancedComputerSearchgById", response, response.Code())
	}
}

// NewCreateAdvancedComputerSearchgByIDCreated creates a CreateAdvancedComputerSearchgByIDCreated with default headers values
func NewCreateAdvancedComputerSearchgByIDCreated() *CreateAdvancedComputerSearchgByIDCreated {
	return &CreateAdvancedComputerSearchgByIDCreated{}
}

/*
CreateAdvancedComputerSearchgByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAdvancedComputerSearchgByIDCreated struct {
}

// IsSuccess returns true when this create advanced computer searchg by Id created response has a 2xx status code
func (o *CreateAdvancedComputerSearchgByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create advanced computer searchg by Id created response has a 3xx status code
func (o *CreateAdvancedComputerSearchgByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create advanced computer searchg by Id created response has a 4xx status code
func (o *CreateAdvancedComputerSearchgByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create advanced computer searchg by Id created response has a 5xx status code
func (o *CreateAdvancedComputerSearchgByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create advanced computer searchg by Id created response a status code equal to that given
func (o *CreateAdvancedComputerSearchgByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create advanced computer searchg by Id created response
func (o *CreateAdvancedComputerSearchgByIDCreated) Code() int {
	return 201
}

func (o *CreateAdvancedComputerSearchgByIDCreated) Error() string {
	return fmt.Sprintf("[POST /advancedcomputersearches/id/{id}][%d] createAdvancedComputerSearchgByIdCreated ", 201)
}

func (o *CreateAdvancedComputerSearchgByIDCreated) String() string {
	return fmt.Sprintf("[POST /advancedcomputersearches/id/{id}][%d] createAdvancedComputerSearchgByIdCreated ", 201)
}

func (o *CreateAdvancedComputerSearchgByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
