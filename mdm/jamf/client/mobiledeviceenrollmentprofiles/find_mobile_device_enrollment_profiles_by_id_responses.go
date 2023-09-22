// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceEnrollmentProfilesByIDReader is a Reader for the FindMobileDeviceEnrollmentProfilesByID structure.
type FindMobileDeviceEnrollmentProfilesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceEnrollmentProfilesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceEnrollmentProfilesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceenrollmentprofiles/id/{id}] findMobileDeviceEnrollmentProfilesById", response, response.Code())
	}
}

// NewFindMobileDeviceEnrollmentProfilesByIDOK creates a FindMobileDeviceEnrollmentProfilesByIDOK with default headers values
func NewFindMobileDeviceEnrollmentProfilesByIDOK() *FindMobileDeviceEnrollmentProfilesByIDOK {
	return &FindMobileDeviceEnrollmentProfilesByIDOK{}
}

/*
FindMobileDeviceEnrollmentProfilesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceEnrollmentProfilesByIDOK struct {
	Payload *models.MobileDeviceEnrollmentProfile
}

// IsSuccess returns true when this find mobile device enrollment profiles by Id o k response has a 2xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device enrollment profiles by Id o k response has a 3xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device enrollment profiles by Id o k response has a 4xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device enrollment profiles by Id o k response has a 5xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device enrollment profiles by Id o k response a status code equal to that given
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device enrollment profiles by Id o k response
func (o *FindMobileDeviceEnrollmentProfilesByIDOK) Code() int {
	return 200
}

func (o *FindMobileDeviceEnrollmentProfilesByIDOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceenrollmentprofiles/id/{id}][%d] findMobileDeviceEnrollmentProfilesByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceEnrollmentProfilesByIDOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceenrollmentprofiles/id/{id}][%d] findMobileDeviceEnrollmentProfilesByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceEnrollmentProfilesByIDOK) GetPayload() *models.MobileDeviceEnrollmentProfile {
	return o.Payload
}

func (o *FindMobileDeviceEnrollmentProfilesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceEnrollmentProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}