// Code generated by go-swagger; DO NOT EDIT.

package computercheckin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerCheckinReader is a Reader for the FindComputerCheckin structure.
type FindComputerCheckinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerCheckinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerCheckinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computercheckin] findComputerCheckin", response, response.Code())
	}
}

// NewFindComputerCheckinOK creates a FindComputerCheckinOK with default headers values
func NewFindComputerCheckinOK() *FindComputerCheckinOK {
	return &FindComputerCheckinOK{}
}

/*
FindComputerCheckinOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerCheckinOK struct {
	Payload *models.ComputerCheckIn
}

// IsSuccess returns true when this find computer checkin o k response has a 2xx status code
func (o *FindComputerCheckinOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer checkin o k response has a 3xx status code
func (o *FindComputerCheckinOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer checkin o k response has a 4xx status code
func (o *FindComputerCheckinOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer checkin o k response has a 5xx status code
func (o *FindComputerCheckinOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer checkin o k response a status code equal to that given
func (o *FindComputerCheckinOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer checkin o k response
func (o *FindComputerCheckinOK) Code() int {
	return 200
}

func (o *FindComputerCheckinOK) Error() string {
	return fmt.Sprintf("[GET /computercheckin][%d] findComputerCheckinOK  %+v", 200, o.Payload)
}

func (o *FindComputerCheckinOK) String() string {
	return fmt.Sprintf("[GET /computercheckin][%d] findComputerCheckinOK  %+v", 200, o.Payload)
}

func (o *FindComputerCheckinOK) GetPayload() *models.ComputerCheckIn {
	return o.Payload
}

func (o *FindComputerCheckinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerCheckIn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}