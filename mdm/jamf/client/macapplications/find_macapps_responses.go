// Code generated by go-swagger; DO NOT EDIT.

package macapplications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMacappsReader is a Reader for the FindMacapps structure.
type FindMacappsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMacappsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMacappsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /macapplications] findMacapps", response, response.Code())
	}
}

// NewFindMacappsOK creates a FindMacappsOK with default headers values
func NewFindMacappsOK() *FindMacappsOK {
	return &FindMacappsOK{}
}

/*
FindMacappsOK describes a response with status code 200, with default header values.

OK
*/
type FindMacappsOK struct {
	Payload models.MacApplications
}

// IsSuccess returns true when this find macapps o k response has a 2xx status code
func (o *FindMacappsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find macapps o k response has a 3xx status code
func (o *FindMacappsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find macapps o k response has a 4xx status code
func (o *FindMacappsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find macapps o k response has a 5xx status code
func (o *FindMacappsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find macapps o k response a status code equal to that given
func (o *FindMacappsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find macapps o k response
func (o *FindMacappsOK) Code() int {
	return 200
}

func (o *FindMacappsOK) Error() string {
	return fmt.Sprintf("[GET /macapplications][%d] findMacappsOK  %+v", 200, o.Payload)
}

func (o *FindMacappsOK) String() string {
	return fmt.Sprintf("[GET /macapplications][%d] findMacappsOK  %+v", 200, o.Payload)
}

func (o *FindMacappsOK) GetPayload() models.MacApplications {
	return o.Payload
}

func (o *FindMacappsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}