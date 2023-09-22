// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceinvitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// CreateMobileDeviceInvitationsByInvitationReader is a Reader for the CreateMobileDeviceInvitationsByInvitation structure.
type CreateMobileDeviceInvitationsByInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMobileDeviceInvitationsByInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMobileDeviceInvitationsByInvitationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /mobiledeviceinvitations/invitation/{invitation}] createMobileDeviceInvitationsByInvitation", response, response.Code())
	}
}

// NewCreateMobileDeviceInvitationsByInvitationCreated creates a CreateMobileDeviceInvitationsByInvitationCreated with default headers values
func NewCreateMobileDeviceInvitationsByInvitationCreated() *CreateMobileDeviceInvitationsByInvitationCreated {
	return &CreateMobileDeviceInvitationsByInvitationCreated{}
}

/*
CreateMobileDeviceInvitationsByInvitationCreated describes a response with status code 201, with default header values.

Created
*/
type CreateMobileDeviceInvitationsByInvitationCreated struct {
	Payload *models.MobileDeviceInvitationPost
}

// IsSuccess returns true when this create mobile device invitations by invitation created response has a 2xx status code
func (o *CreateMobileDeviceInvitationsByInvitationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mobile device invitations by invitation created response has a 3xx status code
func (o *CreateMobileDeviceInvitationsByInvitationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mobile device invitations by invitation created response has a 4xx status code
func (o *CreateMobileDeviceInvitationsByInvitationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mobile device invitations by invitation created response has a 5xx status code
func (o *CreateMobileDeviceInvitationsByInvitationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create mobile device invitations by invitation created response a status code equal to that given
func (o *CreateMobileDeviceInvitationsByInvitationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create mobile device invitations by invitation created response
func (o *CreateMobileDeviceInvitationsByInvitationCreated) Code() int {
	return 201
}

func (o *CreateMobileDeviceInvitationsByInvitationCreated) Error() string {
	return fmt.Sprintf("[POST /mobiledeviceinvitations/invitation/{invitation}][%d] createMobileDeviceInvitationsByInvitationCreated  %+v", 201, o.Payload)
}

func (o *CreateMobileDeviceInvitationsByInvitationCreated) String() string {
	return fmt.Sprintf("[POST /mobiledeviceinvitations/invitation/{invitation}][%d] createMobileDeviceInvitationsByInvitationCreated  %+v", 201, o.Payload)
}

func (o *CreateMobileDeviceInvitationsByInvitationCreated) GetPayload() *models.MobileDeviceInvitationPost {
	return o.Payload
}

func (o *CreateMobileDeviceInvitationsByInvitationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceInvitationPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}