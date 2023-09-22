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

// FindMobileDeviceInvitationsByIDReader is a Reader for the FindMobileDeviceInvitationsByID structure.
type FindMobileDeviceInvitationsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceInvitationsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceInvitationsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceinvitations/id/{id}] findMobileDeviceInvitationsById", response, response.Code())
	}
}

// NewFindMobileDeviceInvitationsByIDOK creates a FindMobileDeviceInvitationsByIDOK with default headers values
func NewFindMobileDeviceInvitationsByIDOK() *FindMobileDeviceInvitationsByIDOK {
	return &FindMobileDeviceInvitationsByIDOK{}
}

/*
FindMobileDeviceInvitationsByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceInvitationsByIDOK struct {
	Payload *models.MobileDeviceInvitation
}

// IsSuccess returns true when this find mobile device invitations by Id o k response has a 2xx status code
func (o *FindMobileDeviceInvitationsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device invitations by Id o k response has a 3xx status code
func (o *FindMobileDeviceInvitationsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device invitations by Id o k response has a 4xx status code
func (o *FindMobileDeviceInvitationsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device invitations by Id o k response has a 5xx status code
func (o *FindMobileDeviceInvitationsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device invitations by Id o k response a status code equal to that given
func (o *FindMobileDeviceInvitationsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device invitations by Id o k response
func (o *FindMobileDeviceInvitationsByIDOK) Code() int {
	return 200
}

func (o *FindMobileDeviceInvitationsByIDOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceinvitations/id/{id}][%d] findMobileDeviceInvitationsByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceInvitationsByIDOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceinvitations/id/{id}][%d] findMobileDeviceInvitationsByIdOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceInvitationsByIDOK) GetPayload() *models.MobileDeviceInvitation {
	return o.Payload
}

func (o *FindMobileDeviceInvitationsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceInvitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}