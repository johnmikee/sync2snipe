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

// FindMobileDeviceProvisioningProfilesByUUIDReader is a Reader for the FindMobileDeviceProvisioningProfilesByUUID structure.
type FindMobileDeviceProvisioningProfilesByUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceProvisioningProfilesByUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceProvisioningProfilesByUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceprovisioningprofiles/uuid/{uuid}] findMobileDeviceProvisioningProfilesByUUID", response, response.Code())
	}
}

// NewFindMobileDeviceProvisioningProfilesByUUIDOK creates a FindMobileDeviceProvisioningProfilesByUUIDOK with default headers values
func NewFindMobileDeviceProvisioningProfilesByUUIDOK() *FindMobileDeviceProvisioningProfilesByUUIDOK {
	return &FindMobileDeviceProvisioningProfilesByUUIDOK{}
}

/*
FindMobileDeviceProvisioningProfilesByUUIDOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceProvisioningProfilesByUUIDOK struct {
	Payload *models.MobileDeviceProvisioningProfile
}

// IsSuccess returns true when this find mobile device provisioning profiles by Uuid o k response has a 2xx status code
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device provisioning profiles by Uuid o k response has a 3xx status code
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device provisioning profiles by Uuid o k response has a 4xx status code
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device provisioning profiles by Uuid o k response has a 5xx status code
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device provisioning profiles by Uuid o k response a status code equal to that given
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device provisioning profiles by Uuid o k response
func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) Code() int {
	return 200
}

func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceprovisioningprofiles/uuid/{uuid}][%d] findMobileDeviceProvisioningProfilesByUuidOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceprovisioningprofiles/uuid/{uuid}][%d] findMobileDeviceProvisioningProfilesByUuidOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) GetPayload() *models.MobileDeviceProvisioningProfile {
	return o.Payload
}

func (o *FindMobileDeviceProvisioningProfilesByUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceProvisioningProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}