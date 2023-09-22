// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceprovisioningprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceProvisioningProfilesByNameReader is a Reader for the FindMobileDeviceProvisioningProfilesByName structure.
type FindMobileDeviceProvisioningProfilesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceProvisioningProfilesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceProvisioningProfilesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceprovisioningprofiles/name/{name}] findMobileDeviceProvisioningProfilesByName", response, response.Code())
	}
}

// NewFindMobileDeviceProvisioningProfilesByNameOK creates a FindMobileDeviceProvisioningProfilesByNameOK with default headers values
func NewFindMobileDeviceProvisioningProfilesByNameOK() *FindMobileDeviceProvisioningProfilesByNameOK {
	return &FindMobileDeviceProvisioningProfilesByNameOK{}
}

/*
FindMobileDeviceProvisioningProfilesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceProvisioningProfilesByNameOK struct {
	Payload *models.MobileDeviceProvisioningProfile
}

// IsSuccess returns true when this find mobile device provisioning profiles by name o k response has a 2xx status code
func (o *FindMobileDeviceProvisioningProfilesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device provisioning profiles by name o k response has a 3xx status code
func (o *FindMobileDeviceProvisioningProfilesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device provisioning profiles by name o k response has a 4xx status code
func (o *FindMobileDeviceProvisioningProfilesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device provisioning profiles by name o k response has a 5xx status code
func (o *FindMobileDeviceProvisioningProfilesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device provisioning profiles by name o k response a status code equal to that given
func (o *FindMobileDeviceProvisioningProfilesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device provisioning profiles by name o k response
func (o *FindMobileDeviceProvisioningProfilesByNameOK) Code() int {
	return 200
}

func (o *FindMobileDeviceProvisioningProfilesByNameOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceprovisioningprofiles/name/{name}][%d] findMobileDeviceProvisioningProfilesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceProvisioningProfilesByNameOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceprovisioningprofiles/name/{name}][%d] findMobileDeviceProvisioningProfilesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceProvisioningProfilesByNameOK) GetPayload() *models.MobileDeviceProvisioningProfile {
	return o.Payload
}

func (o *FindMobileDeviceProvisioningProfilesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceProvisioningProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}