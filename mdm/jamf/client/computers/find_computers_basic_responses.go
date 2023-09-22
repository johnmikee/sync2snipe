// Code generated by go-swagger; DO NOT EDIT.

package computers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputersBasicReader is a Reader for the FindComputersBasic structure.
type FindComputersBasicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputersBasicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputersBasicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computers/subset/basic] findComputersBasic", response, response.Code())
	}
}

// NewFindComputersBasicOK creates a FindComputersBasicOK with default headers values
func NewFindComputersBasicOK() *FindComputersBasicOK {
	return &FindComputersBasicOK{}
}

/*
FindComputersBasicOK describes a response with status code 200, with default header values.

OK
*/
type FindComputersBasicOK struct {
	Payload *models.ComputersBasic
}

// IsSuccess returns true when this find computers basic o k response has a 2xx status code
func (o *FindComputersBasicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computers basic o k response has a 3xx status code
func (o *FindComputersBasicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computers basic o k response has a 4xx status code
func (o *FindComputersBasicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computers basic o k response has a 5xx status code
func (o *FindComputersBasicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computers basic o k response a status code equal to that given
func (o *FindComputersBasicOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computers basic o k response
func (o *FindComputersBasicOK) Code() int {
	return 200
}

func (o *FindComputersBasicOK) Error() string {
	return fmt.Sprintf("[GET /computers/subset/basic][%d] findComputersBasicOK  %+v", 200, o.Payload)
}

func (o *FindComputersBasicOK) String() string {
	return fmt.Sprintf("[GET /computers/subset/basic][%d] findComputersBasicOK  %+v", 200, o.Payload)
}

func (o *FindComputersBasicOK) GetPayload() *models.ComputersBasic {
	return o.Payload
}

func (o *FindComputersBasicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputersBasic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
