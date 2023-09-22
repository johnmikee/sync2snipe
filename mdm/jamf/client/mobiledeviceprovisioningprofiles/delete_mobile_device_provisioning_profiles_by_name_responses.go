// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceprovisioningprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMobileDeviceProvisioningProfilesByNameReader is a Reader for the DeleteMobileDeviceProvisioningProfilesByName structure.
type DeleteMobileDeviceProvisioningProfilesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMobileDeviceProvisioningProfilesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMobileDeviceProvisioningProfilesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /mobiledeviceprovisioningprofiles/name/{name}] deleteMobileDeviceProvisioningProfilesByName", response, response.Code())
	}
}

// NewDeleteMobileDeviceProvisioningProfilesByNameOK creates a DeleteMobileDeviceProvisioningProfilesByNameOK with default headers values
func NewDeleteMobileDeviceProvisioningProfilesByNameOK() *DeleteMobileDeviceProvisioningProfilesByNameOK {
	return &DeleteMobileDeviceProvisioningProfilesByNameOK{}
}

/*
DeleteMobileDeviceProvisioningProfilesByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteMobileDeviceProvisioningProfilesByNameOK struct {
}

// IsSuccess returns true when this delete mobile device provisioning profiles by name o k response has a 2xx status code
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mobile device provisioning profiles by name o k response has a 3xx status code
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobile device provisioning profiles by name o k response has a 4xx status code
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobile device provisioning profiles by name o k response has a 5xx status code
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobile device provisioning profiles by name o k response a status code equal to that given
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete mobile device provisioning profiles by name o k response
func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) Code() int {
	return 200
}

func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /mobiledeviceprovisioningprofiles/name/{name}][%d] deleteMobileDeviceProvisioningProfilesByNameOK ", 200)
}

func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) String() string {
	return fmt.Sprintf("[DELETE /mobiledeviceprovisioningprofiles/name/{name}][%d] deleteMobileDeviceProvisioningProfilesByNameOK ", 200)
}

func (o *DeleteMobileDeviceProvisioningProfilesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}