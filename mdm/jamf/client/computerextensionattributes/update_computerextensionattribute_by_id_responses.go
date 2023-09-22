// Code generated by go-swagger; DO NOT EDIT.

package computerextensionattributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateComputerextensionattributeByIDReader is a Reader for the UpdateComputerextensionattributeByID structure.
type UpdateComputerextensionattributeByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComputerextensionattributeByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateComputerextensionattributeByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /computerextensionattributes/id/{id}] updateComputerextensionattributeById", response, response.Code())
	}
}

// NewUpdateComputerextensionattributeByIDCreated creates a UpdateComputerextensionattributeByIDCreated with default headers values
func NewUpdateComputerextensionattributeByIDCreated() *UpdateComputerextensionattributeByIDCreated {
	return &UpdateComputerextensionattributeByIDCreated{}
}

/*
UpdateComputerextensionattributeByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateComputerextensionattributeByIDCreated struct {
}

// IsSuccess returns true when this update computerextensionattribute by Id created response has a 2xx status code
func (o *UpdateComputerextensionattributeByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update computerextensionattribute by Id created response has a 3xx status code
func (o *UpdateComputerextensionattributeByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update computerextensionattribute by Id created response has a 4xx status code
func (o *UpdateComputerextensionattributeByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update computerextensionattribute by Id created response has a 5xx status code
func (o *UpdateComputerextensionattributeByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update computerextensionattribute by Id created response a status code equal to that given
func (o *UpdateComputerextensionattributeByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update computerextensionattribute by Id created response
func (o *UpdateComputerextensionattributeByIDCreated) Code() int {
	return 201
}

func (o *UpdateComputerextensionattributeByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /computerextensionattributes/id/{id}][%d] updateComputerextensionattributeByIdCreated ", 201)
}

func (o *UpdateComputerextensionattributeByIDCreated) String() string {
	return fmt.Sprintf("[PUT /computerextensionattributes/id/{id}][%d] updateComputerextensionattributeByIdCreated ", 201)
}

func (o *UpdateComputerextensionattributeByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
