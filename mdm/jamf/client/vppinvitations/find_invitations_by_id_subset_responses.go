// Code generated by go-swagger; DO NOT EDIT.

package vppinvitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindInvitationsByIDSubsetReader is a Reader for the FindInvitationsByIDSubset structure.
type FindInvitationsByIDSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindInvitationsByIDSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindInvitationsByIDSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /vppinvitations/id/{id}/subset/{subset}] findInvitationsByIdSubset", response, response.Code())
	}
}

// NewFindInvitationsByIDSubsetOK creates a FindInvitationsByIDSubsetOK with default headers values
func NewFindInvitationsByIDSubsetOK() *FindInvitationsByIDSubsetOK {
	return &FindInvitationsByIDSubsetOK{}
}

/*
FindInvitationsByIDSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindInvitationsByIDSubsetOK struct {
}

// IsSuccess returns true when this find invitations by Id subset o k response has a 2xx status code
func (o *FindInvitationsByIDSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find invitations by Id subset o k response has a 3xx status code
func (o *FindInvitationsByIDSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find invitations by Id subset o k response has a 4xx status code
func (o *FindInvitationsByIDSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find invitations by Id subset o k response has a 5xx status code
func (o *FindInvitationsByIDSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find invitations by Id subset o k response a status code equal to that given
func (o *FindInvitationsByIDSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find invitations by Id subset o k response
func (o *FindInvitationsByIDSubsetOK) Code() int {
	return 200
}

func (o *FindInvitationsByIDSubsetOK) Error() string {
	return fmt.Sprintf("[GET /vppinvitations/id/{id}/subset/{subset}][%d] findInvitationsByIdSubsetOK ", 200)
}

func (o *FindInvitationsByIDSubsetOK) String() string {
	return fmt.Sprintf("[GET /vppinvitations/id/{id}/subset/{subset}][%d] findInvitationsByIdSubsetOK ", 200)
}

func (o *FindInvitationsByIDSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
