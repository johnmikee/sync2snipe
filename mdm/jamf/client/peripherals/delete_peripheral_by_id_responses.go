// Code generated by go-swagger; DO NOT EDIT.

package peripherals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePeripheralByIDReader is a Reader for the DeletePeripheralByID structure.
type DeletePeripheralByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeripheralByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePeripheralByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /peripherals/id/{id}] deletePeripheralById", response, response.Code())
	}
}

// NewDeletePeripheralByIDOK creates a DeletePeripheralByIDOK with default headers values
func NewDeletePeripheralByIDOK() *DeletePeripheralByIDOK {
	return &DeletePeripheralByIDOK{}
}

/*
DeletePeripheralByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeletePeripheralByIDOK struct {
}

// IsSuccess returns true when this delete peripheral by Id o k response has a 2xx status code
func (o *DeletePeripheralByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete peripheral by Id o k response has a 3xx status code
func (o *DeletePeripheralByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peripheral by Id o k response has a 4xx status code
func (o *DeletePeripheralByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete peripheral by Id o k response has a 5xx status code
func (o *DeletePeripheralByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peripheral by Id o k response a status code equal to that given
func (o *DeletePeripheralByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete peripheral by Id o k response
func (o *DeletePeripheralByIDOK) Code() int {
	return 200
}

func (o *DeletePeripheralByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /peripherals/id/{id}][%d] deletePeripheralByIdOK ", 200)
}

func (o *DeletePeripheralByIDOK) String() string {
	return fmt.Sprintf("[DELETE /peripherals/id/{id}][%d] deletePeripheralByIdOK ", 200)
}

func (o *DeletePeripheralByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
