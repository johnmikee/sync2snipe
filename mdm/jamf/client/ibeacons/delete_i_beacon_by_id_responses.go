// Code generated by go-swagger; DO NOT EDIT.

package ibeacons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteIBeaconByIDReader is a Reader for the DeleteIBeaconByID structure.
type DeleteIBeaconByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIBeaconByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIBeaconByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ibeacons/id/{id}] deleteIBeaconById", response, response.Code())
	}
}

// NewDeleteIBeaconByIDOK creates a DeleteIBeaconByIDOK with default headers values
func NewDeleteIBeaconByIDOK() *DeleteIBeaconByIDOK {
	return &DeleteIBeaconByIDOK{}
}

/*
DeleteIBeaconByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteIBeaconByIDOK struct {
}

// IsSuccess returns true when this delete i beacon by Id o k response has a 2xx status code
func (o *DeleteIBeaconByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete i beacon by Id o k response has a 3xx status code
func (o *DeleteIBeaconByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete i beacon by Id o k response has a 4xx status code
func (o *DeleteIBeaconByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete i beacon by Id o k response has a 5xx status code
func (o *DeleteIBeaconByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete i beacon by Id o k response a status code equal to that given
func (o *DeleteIBeaconByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete i beacon by Id o k response
func (o *DeleteIBeaconByIDOK) Code() int {
	return 200
}

func (o *DeleteIBeaconByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /ibeacons/id/{id}][%d] deleteIBeaconByIdOK ", 200)
}

func (o *DeleteIBeaconByIDOK) String() string {
	return fmt.Sprintf("[DELETE /ibeacons/id/{id}][%d] deleteIBeaconByIdOK ", 200)
}

func (o *DeleteIBeaconByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}