// Code generated by go-swagger; DO NOT EDIT.

package mobiledevices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMobileDeviceBySerialNumberReader is a Reader for the UpdateMobileDeviceBySerialNumber structure.
type UpdateMobileDeviceBySerialNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMobileDeviceBySerialNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMobileDeviceBySerialNumberCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /mobiledevices/serialnumber/{serialnumber}] updateMobileDeviceBySerialNumber", response, response.Code())
	}
}

// NewUpdateMobileDeviceBySerialNumberCreated creates a UpdateMobileDeviceBySerialNumberCreated with default headers values
func NewUpdateMobileDeviceBySerialNumberCreated() *UpdateMobileDeviceBySerialNumberCreated {
	return &UpdateMobileDeviceBySerialNumberCreated{}
}

/*
UpdateMobileDeviceBySerialNumberCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMobileDeviceBySerialNumberCreated struct {
}

// IsSuccess returns true when this update mobile device by serial number created response has a 2xx status code
func (o *UpdateMobileDeviceBySerialNumberCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mobile device by serial number created response has a 3xx status code
func (o *UpdateMobileDeviceBySerialNumberCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mobile device by serial number created response has a 4xx status code
func (o *UpdateMobileDeviceBySerialNumberCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mobile device by serial number created response has a 5xx status code
func (o *UpdateMobileDeviceBySerialNumberCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update mobile device by serial number created response a status code equal to that given
func (o *UpdateMobileDeviceBySerialNumberCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update mobile device by serial number created response
func (o *UpdateMobileDeviceBySerialNumberCreated) Code() int {
	return 201
}

func (o *UpdateMobileDeviceBySerialNumberCreated) Error() string {
	return fmt.Sprintf("[PUT /mobiledevices/serialnumber/{serialnumber}][%d] updateMobileDeviceBySerialNumberCreated ", 201)
}

func (o *UpdateMobileDeviceBySerialNumberCreated) String() string {
	return fmt.Sprintf("[PUT /mobiledevices/serialnumber/{serialnumber}][%d] updateMobileDeviceBySerialNumberCreated ", 201)
}

func (o *UpdateMobileDeviceBySerialNumberCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
