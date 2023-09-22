// Code generated by go-swagger; DO NOT EDIT.

package vppaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindVPPAdminAccountByIDReader is a Reader for the FindVPPAdminAccountByID structure.
type FindVPPAdminAccountByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindVPPAdminAccountByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindVPPAdminAccountByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /vppaccounts/id/{id}] findVPPAdminAccountById", response, response.Code())
	}
}

// NewFindVPPAdminAccountByIDOK creates a FindVPPAdminAccountByIDOK with default headers values
func NewFindVPPAdminAccountByIDOK() *FindVPPAdminAccountByIDOK {
	return &FindVPPAdminAccountByIDOK{}
}

/*
FindVPPAdminAccountByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindVPPAdminAccountByIDOK struct {
	Payload *models.VppAccount
}

// IsSuccess returns true when this find v p p admin account by Id o k response has a 2xx status code
func (o *FindVPPAdminAccountByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find v p p admin account by Id o k response has a 3xx status code
func (o *FindVPPAdminAccountByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find v p p admin account by Id o k response has a 4xx status code
func (o *FindVPPAdminAccountByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find v p p admin account by Id o k response has a 5xx status code
func (o *FindVPPAdminAccountByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find v p p admin account by Id o k response a status code equal to that given
func (o *FindVPPAdminAccountByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find v p p admin account by Id o k response
func (o *FindVPPAdminAccountByIDOK) Code() int {
	return 200
}

func (o *FindVPPAdminAccountByIDOK) Error() string {
	return fmt.Sprintf("[GET /vppaccounts/id/{id}][%d] findVPPAdminAccountByIdOK  %+v", 200, o.Payload)
}

func (o *FindVPPAdminAccountByIDOK) String() string {
	return fmt.Sprintf("[GET /vppaccounts/id/{id}][%d] findVPPAdminAccountByIdOK  %+v", 200, o.Payload)
}

func (o *FindVPPAdminAccountByIDOK) GetPayload() *models.VppAccount {
	return o.Payload
}

func (o *FindVPPAdminAccountByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VppAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}