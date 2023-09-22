// Code generated by go-swagger; DO NOT EDIT.

package computerinvitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// CreateComputerInvitationsByInvitationReader is a Reader for the CreateComputerInvitationsByInvitation structure.
type CreateComputerInvitationsByInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateComputerInvitationsByInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateComputerInvitationsByInvitationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /computerinvitations/invitation/{invitation}] createComputerInvitationsByInvitation", response, response.Code())
	}
}

// NewCreateComputerInvitationsByInvitationCreated creates a CreateComputerInvitationsByInvitationCreated with default headers values
func NewCreateComputerInvitationsByInvitationCreated() *CreateComputerInvitationsByInvitationCreated {
	return &CreateComputerInvitationsByInvitationCreated{}
}

/*
CreateComputerInvitationsByInvitationCreated describes a response with status code 201, with default header values.

Created
*/
type CreateComputerInvitationsByInvitationCreated struct {
	Payload *models.ComputerInvitation
}

// IsSuccess returns true when this create computer invitations by invitation created response has a 2xx status code
func (o *CreateComputerInvitationsByInvitationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create computer invitations by invitation created response has a 3xx status code
func (o *CreateComputerInvitationsByInvitationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create computer invitations by invitation created response has a 4xx status code
func (o *CreateComputerInvitationsByInvitationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create computer invitations by invitation created response has a 5xx status code
func (o *CreateComputerInvitationsByInvitationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create computer invitations by invitation created response a status code equal to that given
func (o *CreateComputerInvitationsByInvitationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create computer invitations by invitation created response
func (o *CreateComputerInvitationsByInvitationCreated) Code() int {
	return 201
}

func (o *CreateComputerInvitationsByInvitationCreated) Error() string {
	return fmt.Sprintf("[POST /computerinvitations/invitation/{invitation}][%d] createComputerInvitationsByInvitationCreated  %+v", 201, o.Payload)
}

func (o *CreateComputerInvitationsByInvitationCreated) String() string {
	return fmt.Sprintf("[POST /computerinvitations/invitation/{invitation}][%d] createComputerInvitationsByInvitationCreated  %+v", 201, o.Payload)
}

func (o *CreateComputerInvitationsByInvitationCreated) GetPayload() *models.ComputerInvitation {
	return o.Payload
}

func (o *CreateComputerInvitationsByInvitationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerInvitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
