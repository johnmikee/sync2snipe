// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceconfigurationprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMobileDeviceConfigurationProfileByIDReader is a Reader for the UpdateMobileDeviceConfigurationProfileByID structure.
type UpdateMobileDeviceConfigurationProfileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMobileDeviceConfigurationProfileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMobileDeviceConfigurationProfileByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /mobiledeviceconfigurationprofiles/id/{id}] updateMobileDeviceConfigurationProfileById", response, response.Code())
	}
}

// NewUpdateMobileDeviceConfigurationProfileByIDCreated creates a UpdateMobileDeviceConfigurationProfileByIDCreated with default headers values
func NewUpdateMobileDeviceConfigurationProfileByIDCreated() *UpdateMobileDeviceConfigurationProfileByIDCreated {
	return &UpdateMobileDeviceConfigurationProfileByIDCreated{}
}

/*
UpdateMobileDeviceConfigurationProfileByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMobileDeviceConfigurationProfileByIDCreated struct {
}

// IsSuccess returns true when this update mobile device configuration profile by Id created response has a 2xx status code
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mobile device configuration profile by Id created response has a 3xx status code
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mobile device configuration profile by Id created response has a 4xx status code
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mobile device configuration profile by Id created response has a 5xx status code
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update mobile device configuration profile by Id created response a status code equal to that given
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update mobile device configuration profile by Id created response
func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) Code() int {
	return 201
}

func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /mobiledeviceconfigurationprofiles/id/{id}][%d] updateMobileDeviceConfigurationProfileByIdCreated ", 201)
}

func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) String() string {
	return fmt.Sprintf("[PUT /mobiledeviceconfigurationprofiles/id/{id}][%d] updateMobileDeviceConfigurationProfileByIdCreated ", 201)
}

func (o *UpdateMobileDeviceConfigurationProfileByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
