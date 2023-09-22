// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceapplications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMobileDeviceApplicationByIDReader is a Reader for the DeleteMobileDeviceApplicationByID structure.
type DeleteMobileDeviceApplicationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMobileDeviceApplicationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMobileDeviceApplicationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /mobiledeviceapplications/id/{id}] deleteMobileDeviceApplicationById", response, response.Code())
	}
}

// NewDeleteMobileDeviceApplicationByIDOK creates a DeleteMobileDeviceApplicationByIDOK with default headers values
func NewDeleteMobileDeviceApplicationByIDOK() *DeleteMobileDeviceApplicationByIDOK {
	return &DeleteMobileDeviceApplicationByIDOK{}
}

/*
DeleteMobileDeviceApplicationByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteMobileDeviceApplicationByIDOK struct {
}

// IsSuccess returns true when this delete mobile device application by Id o k response has a 2xx status code
func (o *DeleteMobileDeviceApplicationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mobile device application by Id o k response has a 3xx status code
func (o *DeleteMobileDeviceApplicationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobile device application by Id o k response has a 4xx status code
func (o *DeleteMobileDeviceApplicationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobile device application by Id o k response has a 5xx status code
func (o *DeleteMobileDeviceApplicationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobile device application by Id o k response a status code equal to that given
func (o *DeleteMobileDeviceApplicationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete mobile device application by Id o k response
func (o *DeleteMobileDeviceApplicationByIDOK) Code() int {
	return 200
}

func (o *DeleteMobileDeviceApplicationByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /mobiledeviceapplications/id/{id}][%d] deleteMobileDeviceApplicationByIdOK ", 200)
}

func (o *DeleteMobileDeviceApplicationByIDOK) String() string {
	return fmt.Sprintf("[DELETE /mobiledeviceapplications/id/{id}][%d] deleteMobileDeviceApplicationByIdOK ", 200)
}

func (o *DeleteMobileDeviceApplicationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
