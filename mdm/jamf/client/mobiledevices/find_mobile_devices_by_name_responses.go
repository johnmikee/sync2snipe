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

// FindMobileDevicesByNameReader is a Reader for the FindMobileDevicesByName structure.
type FindMobileDevicesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDevicesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDevicesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledevices/name/{name}] findMobileDevicesByName", response, response.Code())
	}
}

// NewFindMobileDevicesByNameOK creates a FindMobileDevicesByNameOK with default headers values
func NewFindMobileDevicesByNameOK() *FindMobileDevicesByNameOK {
	return &FindMobileDevicesByNameOK{}
}

/*
FindMobileDevicesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDevicesByNameOK struct {
	Payload *models.MobileDevice
}

// IsSuccess returns true when this find mobile devices by name o k response has a 2xx status code
func (o *FindMobileDevicesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile devices by name o k response has a 3xx status code
func (o *FindMobileDevicesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile devices by name o k response has a 4xx status code
func (o *FindMobileDevicesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile devices by name o k response has a 5xx status code
func (o *FindMobileDevicesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile devices by name o k response a status code equal to that given
func (o *FindMobileDevicesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile devices by name o k response
func (o *FindMobileDevicesByNameOK) Code() int {
	return 200
}

func (o *FindMobileDevicesByNameOK) Error() string {
	return fmt.Sprintf("[GET /mobiledevices/name/{name}][%d] findMobileDevicesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByNameOK) String() string {
	return fmt.Sprintf("[GET /mobiledevices/name/{name}][%d] findMobileDevicesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByNameOK) GetPayload() *models.MobileDevice {
	return o.Payload
}

func (o *FindMobileDevicesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}