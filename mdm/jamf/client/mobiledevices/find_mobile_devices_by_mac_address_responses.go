// Code generated by go-swagger; DO NOT EDIT.

package mobiledevices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDevicesByMacAddressReader is a Reader for the FindMobileDevicesByMacAddress structure.
type FindMobileDevicesByMacAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDevicesByMacAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDevicesByMacAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledevices/macaddress/{macaddress}] findMobileDevicesByMacAddress", response, response.Code())
	}
}

// NewFindMobileDevicesByMacAddressOK creates a FindMobileDevicesByMacAddressOK with default headers values
func NewFindMobileDevicesByMacAddressOK() *FindMobileDevicesByMacAddressOK {
	return &FindMobileDevicesByMacAddressOK{}
}

/*
FindMobileDevicesByMacAddressOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDevicesByMacAddressOK struct {
	Payload *models.MobileDevice
}

// IsSuccess returns true when this find mobile devices by mac address o k response has a 2xx status code
func (o *FindMobileDevicesByMacAddressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile devices by mac address o k response has a 3xx status code
func (o *FindMobileDevicesByMacAddressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile devices by mac address o k response has a 4xx status code
func (o *FindMobileDevicesByMacAddressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile devices by mac address o k response has a 5xx status code
func (o *FindMobileDevicesByMacAddressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile devices by mac address o k response a status code equal to that given
func (o *FindMobileDevicesByMacAddressOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile devices by mac address o k response
func (o *FindMobileDevicesByMacAddressOK) Code() int {
	return 200
}

func (o *FindMobileDevicesByMacAddressOK) Error() string {
	return fmt.Sprintf("[GET /mobiledevices/macaddress/{macaddress}][%d] findMobileDevicesByMacAddressOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByMacAddressOK) String() string {
	return fmt.Sprintf("[GET /mobiledevices/macaddress/{macaddress}][%d] findMobileDevicesByMacAddressOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByMacAddressOK) GetPayload() *models.MobileDevice {
	return o.Payload
}

func (o *FindMobileDevicesByMacAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
