// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMobileDeviceEnrollmentProfileByInvitationReader is a Reader for the UpdateMobileDeviceEnrollmentProfileByInvitation structure.
type UpdateMobileDeviceEnrollmentProfileByInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMobileDeviceEnrollmentProfileByInvitationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /mobiledeviceenrollmentprofiles/invitation/{invitation}] updateMobileDeviceEnrollmentProfileByInvitation", response, response.Code())
	}
}

// NewUpdateMobileDeviceEnrollmentProfileByInvitationCreated creates a UpdateMobileDeviceEnrollmentProfileByInvitationCreated with default headers values
func NewUpdateMobileDeviceEnrollmentProfileByInvitationCreated() *UpdateMobileDeviceEnrollmentProfileByInvitationCreated {
	return &UpdateMobileDeviceEnrollmentProfileByInvitationCreated{}
}

/*
UpdateMobileDeviceEnrollmentProfileByInvitationCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMobileDeviceEnrollmentProfileByInvitationCreated struct {
}

// IsSuccess returns true when this update mobile device enrollment profile by invitation created response has a 2xx status code
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mobile device enrollment profile by invitation created response has a 3xx status code
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mobile device enrollment profile by invitation created response has a 4xx status code
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mobile device enrollment profile by invitation created response has a 5xx status code
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update mobile device enrollment profile by invitation created response a status code equal to that given
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update mobile device enrollment profile by invitation created response
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) Code() int {
	return 201
}

func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) Error() string {
	return fmt.Sprintf("[PUT /mobiledeviceenrollmentprofiles/invitation/{invitation}][%d] updateMobileDeviceEnrollmentProfileByInvitationCreated ", 201)
}

func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) String() string {
	return fmt.Sprintf("[PUT /mobiledeviceenrollmentprofiles/invitation/{invitation}][%d] updateMobileDeviceEnrollmentProfileByInvitationCreated ", 201)
}

func (o *UpdateMobileDeviceEnrollmentProfileByInvitationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}