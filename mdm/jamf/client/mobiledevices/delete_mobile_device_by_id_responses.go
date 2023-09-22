// Code generated by go-swagger; DO NOT EDIT.

package mobiledevices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMobileDeviceByIDReader is a Reader for the DeleteMobileDeviceByID structure.
type DeleteMobileDeviceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMobileDeviceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMobileDeviceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /mobiledevices/id/{id}] deleteMobileDeviceById", response, response.Code())
	}
}

// NewDeleteMobileDeviceByIDOK creates a DeleteMobileDeviceByIDOK with default headers values
func NewDeleteMobileDeviceByIDOK() *DeleteMobileDeviceByIDOK {
	return &DeleteMobileDeviceByIDOK{}
}

/*
DeleteMobileDeviceByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteMobileDeviceByIDOK struct {
}

// IsSuccess returns true when this delete mobile device by Id o k response has a 2xx status code
func (o *DeleteMobileDeviceByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mobile device by Id o k response has a 3xx status code
func (o *DeleteMobileDeviceByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobile device by Id o k response has a 4xx status code
func (o *DeleteMobileDeviceByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobile device by Id o k response has a 5xx status code
func (o *DeleteMobileDeviceByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobile device by Id o k response a status code equal to that given
func (o *DeleteMobileDeviceByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete mobile device by Id o k response
func (o *DeleteMobileDeviceByIDOK) Code() int {
	return 200
}

func (o *DeleteMobileDeviceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /mobiledevices/id/{id}][%d] deleteMobileDeviceByIdOK ", 200)
}

func (o *DeleteMobileDeviceByIDOK) String() string {
	return fmt.Sprintf("[DELETE /mobiledevices/id/{id}][%d] deleteMobileDeviceByIdOK ", 200)
}

func (o *DeleteMobileDeviceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}