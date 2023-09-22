// Code generated by go-swagger; DO NOT EDIT.

package computergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateComputerGroupByIDReader is a Reader for the CreateComputerGroupByID structure.
type CreateComputerGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateComputerGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateComputerGroupByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /computergroups/id/{id}] createComputerGroupById", response, response.Code())
	}
}

// NewCreateComputerGroupByIDCreated creates a CreateComputerGroupByIDCreated with default headers values
func NewCreateComputerGroupByIDCreated() *CreateComputerGroupByIDCreated {
	return &CreateComputerGroupByIDCreated{}
}

/*
CreateComputerGroupByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateComputerGroupByIDCreated struct {
}

// IsSuccess returns true when this create computer group by Id created response has a 2xx status code
func (o *CreateComputerGroupByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create computer group by Id created response has a 3xx status code
func (o *CreateComputerGroupByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create computer group by Id created response has a 4xx status code
func (o *CreateComputerGroupByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create computer group by Id created response has a 5xx status code
func (o *CreateComputerGroupByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create computer group by Id created response a status code equal to that given
func (o *CreateComputerGroupByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create computer group by Id created response
func (o *CreateComputerGroupByIDCreated) Code() int {
	return 201
}

func (o *CreateComputerGroupByIDCreated) Error() string {
	return fmt.Sprintf("[POST /computergroups/id/{id}][%d] createComputerGroupByIdCreated ", 201)
}

func (o *CreateComputerGroupByIDCreated) String() string {
	return fmt.Sprintf("[POST /computergroups/id/{id}][%d] createComputerGroupByIdCreated ", 201)
}

func (o *CreateComputerGroupByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
