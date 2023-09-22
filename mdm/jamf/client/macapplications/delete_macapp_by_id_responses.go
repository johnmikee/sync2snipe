// Code generated by go-swagger; DO NOT EDIT.

package macapplications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMacappByIDReader is a Reader for the DeleteMacappByID structure.
type DeleteMacappByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMacappByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMacappByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /macapplications/id/{id}] deleteMacappById", response, response.Code())
	}
}

// NewDeleteMacappByIDOK creates a DeleteMacappByIDOK with default headers values
func NewDeleteMacappByIDOK() *DeleteMacappByIDOK {
	return &DeleteMacappByIDOK{}
}

/*
DeleteMacappByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteMacappByIDOK struct {
}

// IsSuccess returns true when this delete macapp by Id o k response has a 2xx status code
func (o *DeleteMacappByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete macapp by Id o k response has a 3xx status code
func (o *DeleteMacappByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete macapp by Id o k response has a 4xx status code
func (o *DeleteMacappByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete macapp by Id o k response has a 5xx status code
func (o *DeleteMacappByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete macapp by Id o k response a status code equal to that given
func (o *DeleteMacappByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete macapp by Id o k response
func (o *DeleteMacappByIDOK) Code() int {
	return 200
}

func (o *DeleteMacappByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /macapplications/id/{id}][%d] deleteMacappByIdOK ", 200)
}

func (o *DeleteMacappByIDOK) String() string {
	return fmt.Sprintf("[DELETE /macapplications/id/{id}][%d] deleteMacappByIdOK ", 200)
}

func (o *DeleteMacappByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
