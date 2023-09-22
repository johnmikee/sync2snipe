// Code generated by go-swagger; DO NOT EDIT.

package classes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteClassByNameReader is a Reader for the DeleteClassByName structure.
type DeleteClassByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClassByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClassByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /classes/name/{name}] deleteClassByName", response, response.Code())
	}
}

// NewDeleteClassByNameOK creates a DeleteClassByNameOK with default headers values
func NewDeleteClassByNameOK() *DeleteClassByNameOK {
	return &DeleteClassByNameOK{}
}

/*
DeleteClassByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteClassByNameOK struct {
}

// IsSuccess returns true when this delete class by name o k response has a 2xx status code
func (o *DeleteClassByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete class by name o k response has a 3xx status code
func (o *DeleteClassByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete class by name o k response has a 4xx status code
func (o *DeleteClassByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete class by name o k response has a 5xx status code
func (o *DeleteClassByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete class by name o k response a status code equal to that given
func (o *DeleteClassByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete class by name o k response
func (o *DeleteClassByNameOK) Code() int {
	return 200
}

func (o *DeleteClassByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /classes/name/{name}][%d] deleteClassByNameOK ", 200)
}

func (o *DeleteClassByNameOK) String() string {
	return fmt.Sprintf("[DELETE /classes/name/{name}][%d] deleteClassByNameOK ", 200)
}

func (o *DeleteClassByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}