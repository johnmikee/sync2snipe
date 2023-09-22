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

// FindMobileDevicesByIDReader is a Reader for the FindMobileDevicesByID structure.
type FindMobileDevicesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDevicesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDevicesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledevices/id/{id}] findMobileDevicesById", response, response.Code())
	}
}

// NewFindMobileDevicesByIDOK creates a FindMobileDevicesByIDOK with default headers values
func NewFindMobileDevicesByIDOK() *FindMobileDevicesByIDOK {
	return &FindMobileDevicesByIDOK{}
}

/*
FindMobileDevicesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDevicesByIDOK struct {
	Payload *models.MobileDevice
}

// IsSuccess returns true when this find mobile devices by Id o k response has a 2xx status code
func (o *FindMobileDevicesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile devices by Id o k response has a 3xx status code
func (o *FindMobileDevicesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile devices by Id o k response has a 4xx status code
func (o *FindMobileDevicesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile devices by Id o k response has a 5xx status code
func (o *FindMobileDevicesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile devices by Id o k response a status code equal to that given
func (o *FindMobileDevicesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile devices by Id o k response
func (o *FindMobileDevicesByIDOK) Code() int {
	return 200
}

func (o *FindMobileDevicesByIDOK) Error() string {
	return fmt.Sprintf("[GET /mobiledevices/id/{id}][%d] findMobileDevicesByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByIDOK) String() string {
	return fmt.Sprintf("[GET /mobiledevices/id/{id}][%d] findMobileDevicesByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDevicesByIDOK) GetPayload() *models.MobileDevice {
	return o.Payload
}

func (o *FindMobileDevicesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
