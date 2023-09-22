// Code generated by go-swagger; DO NOT EDIT.

package vppinvitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateInvitationByIDReader is a Reader for the UpdateInvitationByID structure.
type UpdateInvitationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInvitationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateInvitationByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /vppinvitations/id/{id}] updateInvitationById", response, response.Code())
	}
}

// NewUpdateInvitationByIDCreated creates a UpdateInvitationByIDCreated with default headers values
func NewUpdateInvitationByIDCreated() *UpdateInvitationByIDCreated {
	return &UpdateInvitationByIDCreated{}
}

/*
UpdateInvitationByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateInvitationByIDCreated struct {
}

// IsSuccess returns true when this update invitation by Id created response has a 2xx status code
func (o *UpdateInvitationByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update invitation by Id created response has a 3xx status code
func (o *UpdateInvitationByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update invitation by Id created response has a 4xx status code
func (o *UpdateInvitationByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update invitation by Id created response has a 5xx status code
func (o *UpdateInvitationByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update invitation by Id created response a status code equal to that given
func (o *UpdateInvitationByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update invitation by Id created response
func (o *UpdateInvitationByIDCreated) Code() int {
	return 201
}

func (o *UpdateInvitationByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /vppinvitations/id/{id}][%d] updateInvitationByIdCreated ", 201)
}

func (o *UpdateInvitationByIDCreated) String() string {
	return fmt.Sprintf("[PUT /vppinvitations/id/{id}][%d] updateInvitationByIdCreated ", 201)
}

func (o *UpdateInvitationByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
