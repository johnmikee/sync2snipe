// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindGroupsByIDReader is a Reader for the FindGroupsByID structure.
type FindGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindGroupsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /accounts/groupid/{id}] findGroupsById", response, response.Code())
	}
}

// NewFindGroupsByIDOK creates a FindGroupsByIDOK with default headers values
func NewFindGroupsByIDOK() *FindGroupsByIDOK {
	return &FindGroupsByIDOK{}
}

/*
FindGroupsByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindGroupsByIDOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this find groups by Id o k response has a 2xx status code
func (o *FindGroupsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find groups by Id o k response has a 3xx status code
func (o *FindGroupsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find groups by Id o k response has a 4xx status code
func (o *FindGroupsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find groups by Id o k response has a 5xx status code
func (o *FindGroupsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find groups by Id o k response a status code equal to that given
func (o *FindGroupsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find groups by Id o k response
func (o *FindGroupsByIDOK) Code() int {
	return 200
}

func (o *FindGroupsByIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/groupid/{id}][%d] findGroupsByIdOK  %+v", 200, o.Payload)
}

func (o *FindGroupsByIDOK) String() string {
	return fmt.Sprintf("[GET /accounts/groupid/{id}][%d] findGroupsByIdOK  %+v", 200, o.Payload)
}

func (o *FindGroupsByIDOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *FindGroupsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
