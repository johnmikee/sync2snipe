// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceprovisioningprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateMobileDeviceProvisioningProfilesByIDReader is a Reader for the CreateMobileDeviceProvisioningProfilesByID structure.
type CreateMobileDeviceProvisioningProfilesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMobileDeviceProvisioningProfilesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMobileDeviceProvisioningProfilesByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /mobiledeviceprovisioningprofiles/id/{id}] createMobileDeviceProvisioningProfilesById", response, response.Code())
	}
}

// NewCreateMobileDeviceProvisioningProfilesByIDCreated creates a CreateMobileDeviceProvisioningProfilesByIDCreated with default headers values
func NewCreateMobileDeviceProvisioningProfilesByIDCreated() *CreateMobileDeviceProvisioningProfilesByIDCreated {
	return &CreateMobileDeviceProvisioningProfilesByIDCreated{}
}

/*
CreateMobileDeviceProvisioningProfilesByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateMobileDeviceProvisioningProfilesByIDCreated struct {
}

// IsSuccess returns true when this create mobile device provisioning profiles by Id created response has a 2xx status code
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mobile device provisioning profiles by Id created response has a 3xx status code
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mobile device provisioning profiles by Id created response has a 4xx status code
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mobile device provisioning profiles by Id created response has a 5xx status code
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create mobile device provisioning profiles by Id created response a status code equal to that given
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create mobile device provisioning profiles by Id created response
func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) Code() int {
	return 201
}

func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) Error() string {
	return fmt.Sprintf("[POST /mobiledeviceprovisioningprofiles/id/{id}][%d] createMobileDeviceProvisioningProfilesByIdCreated ", 201)
}

func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) String() string {
	return fmt.Sprintf("[POST /mobiledeviceprovisioningprofiles/id/{id}][%d] createMobileDeviceProvisioningProfilesByIdCreated ", 201)
}

func (o *CreateMobileDeviceProvisioningProfilesByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}