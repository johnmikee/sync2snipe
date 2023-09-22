// Code generated by go-swagger; DO NOT EDIT.

package vppaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteVPPAdminAccountByIDReader is a Reader for the DeleteVPPAdminAccountByID structure.
type DeleteVPPAdminAccountByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVPPAdminAccountByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVPPAdminAccountByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /vppaccounts/id/{id}] deleteVPPAdminAccountById", response, response.Code())
	}
}

// NewDeleteVPPAdminAccountByIDOK creates a DeleteVPPAdminAccountByIDOK with default headers values
func NewDeleteVPPAdminAccountByIDOK() *DeleteVPPAdminAccountByIDOK {
	return &DeleteVPPAdminAccountByIDOK{}
}

/*
DeleteVPPAdminAccountByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteVPPAdminAccountByIDOK struct {
}

// IsSuccess returns true when this delete v p p admin account by Id o k response has a 2xx status code
func (o *DeleteVPPAdminAccountByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v p p admin account by Id o k response has a 3xx status code
func (o *DeleteVPPAdminAccountByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v p p admin account by Id o k response has a 4xx status code
func (o *DeleteVPPAdminAccountByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v p p admin account by Id o k response has a 5xx status code
func (o *DeleteVPPAdminAccountByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v p p admin account by Id o k response a status code equal to that given
func (o *DeleteVPPAdminAccountByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v p p admin account by Id o k response
func (o *DeleteVPPAdminAccountByIDOK) Code() int {
	return 200
}

func (o *DeleteVPPAdminAccountByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /vppaccounts/id/{id}][%d] deleteVPPAdminAccountByIdOK ", 200)
}

func (o *DeleteVPPAdminAccountByIDOK) String() string {
	return fmt.Sprintf("[DELETE /vppaccounts/id/{id}][%d] deleteVPPAdminAccountByIdOK ", 200)
}

func (o *DeleteVPPAdminAccountByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}