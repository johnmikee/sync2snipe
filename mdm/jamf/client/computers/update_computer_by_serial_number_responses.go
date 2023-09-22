// Code generated by go-swagger; DO NOT EDIT.

package computers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateComputerBySerialNumberReader is a Reader for the UpdateComputerBySerialNumber structure.
type UpdateComputerBySerialNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComputerBySerialNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateComputerBySerialNumberCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /computers/serialnumber/{serialnumber}] updateComputerBySerialNumber", response, response.Code())
	}
}

// NewUpdateComputerBySerialNumberCreated creates a UpdateComputerBySerialNumberCreated with default headers values
func NewUpdateComputerBySerialNumberCreated() *UpdateComputerBySerialNumberCreated {
	return &UpdateComputerBySerialNumberCreated{}
}

/*
UpdateComputerBySerialNumberCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateComputerBySerialNumberCreated struct {
}

// IsSuccess returns true when this update computer by serial number created response has a 2xx status code
func (o *UpdateComputerBySerialNumberCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update computer by serial number created response has a 3xx status code
func (o *UpdateComputerBySerialNumberCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update computer by serial number created response has a 4xx status code
func (o *UpdateComputerBySerialNumberCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update computer by serial number created response has a 5xx status code
func (o *UpdateComputerBySerialNumberCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update computer by serial number created response a status code equal to that given
func (o *UpdateComputerBySerialNumberCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update computer by serial number created response
func (o *UpdateComputerBySerialNumberCreated) Code() int {
	return 201
}

func (o *UpdateComputerBySerialNumberCreated) Error() string {
	return fmt.Sprintf("[PUT /computers/serialnumber/{serialnumber}][%d] updateComputerBySerialNumberCreated ", 201)
}

func (o *UpdateComputerBySerialNumberCreated) String() string {
	return fmt.Sprintf("[PUT /computers/serialnumber/{serialnumber}][%d] updateComputerBySerialNumberCreated ", 201)
}

func (o *UpdateComputerBySerialNumberCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
