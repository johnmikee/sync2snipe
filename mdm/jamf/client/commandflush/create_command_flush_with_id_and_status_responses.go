// Code generated by go-swagger; DO NOT EDIT.

package commandflush

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateCommandFlushWithIDAndStatusReader is a Reader for the CreateCommandFlushWithIDAndStatus structure.
type CreateCommandFlushWithIDAndStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCommandFlushWithIDAndStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCommandFlushWithIDAndStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /commandflush/{idtype}/id/{id}/status/{status}] createCommandFlushWithIdAndStatus", response, response.Code())
	}
}

// NewCreateCommandFlushWithIDAndStatusOK creates a CreateCommandFlushWithIDAndStatusOK with default headers values
func NewCreateCommandFlushWithIDAndStatusOK() *CreateCommandFlushWithIDAndStatusOK {
	return &CreateCommandFlushWithIDAndStatusOK{}
}

/*
CreateCommandFlushWithIDAndStatusOK describes a response with status code 200, with default header values.

OK
*/
type CreateCommandFlushWithIDAndStatusOK struct {
}

// IsSuccess returns true when this create command flush with Id and status o k response has a 2xx status code
func (o *CreateCommandFlushWithIDAndStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create command flush with Id and status o k response has a 3xx status code
func (o *CreateCommandFlushWithIDAndStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create command flush with Id and status o k response has a 4xx status code
func (o *CreateCommandFlushWithIDAndStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create command flush with Id and status o k response has a 5xx status code
func (o *CreateCommandFlushWithIDAndStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create command flush with Id and status o k response a status code equal to that given
func (o *CreateCommandFlushWithIDAndStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create command flush with Id and status o k response
func (o *CreateCommandFlushWithIDAndStatusOK) Code() int {
	return 200
}

func (o *CreateCommandFlushWithIDAndStatusOK) Error() string {
	return fmt.Sprintf("[DELETE /commandflush/{idtype}/id/{id}/status/{status}][%d] createCommandFlushWithIdAndStatusOK ", 200)
}

func (o *CreateCommandFlushWithIDAndStatusOK) String() string {
	return fmt.Sprintf("[DELETE /commandflush/{idtype}/id/{id}/status/{status}][%d] createCommandFlushWithIdAndStatusOK ", 200)
}

func (o *CreateCommandFlushWithIDAndStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}