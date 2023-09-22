// Code generated by go-swagger; DO NOT EDIT.

package advancedcomputersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAdvancedComputerSearchByIDReader is a Reader for the UpdateAdvancedComputerSearchByID structure.
type UpdateAdvancedComputerSearchByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdvancedComputerSearchByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAdvancedComputerSearchByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /advancedcomputersearches/id/{id}] updateAdvancedComputerSearchById", response, response.Code())
	}
}

// NewUpdateAdvancedComputerSearchByIDCreated creates a UpdateAdvancedComputerSearchByIDCreated with default headers values
func NewUpdateAdvancedComputerSearchByIDCreated() *UpdateAdvancedComputerSearchByIDCreated {
	return &UpdateAdvancedComputerSearchByIDCreated{}
}

/*
UpdateAdvancedComputerSearchByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAdvancedComputerSearchByIDCreated struct {
}

// IsSuccess returns true when this update advanced computer search by Id created response has a 2xx status code
func (o *UpdateAdvancedComputerSearchByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update advanced computer search by Id created response has a 3xx status code
func (o *UpdateAdvancedComputerSearchByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update advanced computer search by Id created response has a 4xx status code
func (o *UpdateAdvancedComputerSearchByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update advanced computer search by Id created response has a 5xx status code
func (o *UpdateAdvancedComputerSearchByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update advanced computer search by Id created response a status code equal to that given
func (o *UpdateAdvancedComputerSearchByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update advanced computer search by Id created response
func (o *UpdateAdvancedComputerSearchByIDCreated) Code() int {
	return 201
}

func (o *UpdateAdvancedComputerSearchByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /advancedcomputersearches/id/{id}][%d] updateAdvancedComputerSearchByIdCreated ", 201)
}

func (o *UpdateAdvancedComputerSearchByIDCreated) String() string {
	return fmt.Sprintf("[PUT /advancedcomputersearches/id/{id}][%d] updateAdvancedComputerSearchByIdCreated ", 201)
}

func (o *UpdateAdvancedComputerSearchByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
