// Code generated by go-swagger; DO NOT EDIT.

package computers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateComputerByUDIDReader is a Reader for the UpdateComputerByUDID structure.
type UpdateComputerByUDIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComputerByUDIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateComputerByUDIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /computers/udid/{udid}] updateComputerByUDID", response, response.Code())
	}
}

// NewUpdateComputerByUDIDCreated creates a UpdateComputerByUDIDCreated with default headers values
func NewUpdateComputerByUDIDCreated() *UpdateComputerByUDIDCreated {
	return &UpdateComputerByUDIDCreated{}
}

/*
UpdateComputerByUDIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateComputerByUDIDCreated struct {
}

// IsSuccess returns true when this update computer by u d Id created response has a 2xx status code
func (o *UpdateComputerByUDIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update computer by u d Id created response has a 3xx status code
func (o *UpdateComputerByUDIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update computer by u d Id created response has a 4xx status code
func (o *UpdateComputerByUDIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update computer by u d Id created response has a 5xx status code
func (o *UpdateComputerByUDIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update computer by u d Id created response a status code equal to that given
func (o *UpdateComputerByUDIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update computer by u d Id created response
func (o *UpdateComputerByUDIDCreated) Code() int {
	return 201
}

func (o *UpdateComputerByUDIDCreated) Error() string {
	return fmt.Sprintf("[PUT /computers/udid/{udid}][%d] updateComputerByUDIdCreated ", 201)
}

func (o *UpdateComputerByUDIDCreated) String() string {
	return fmt.Sprintf("[PUT /computers/udid/{udid}][%d] updateComputerByUDIdCreated ", 201)
}

func (o *UpdateComputerByUDIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}