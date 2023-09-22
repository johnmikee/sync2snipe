// Code generated by go-swagger; DO NOT EDIT.

package computerhistory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerHistoryByNameReader is a Reader for the FindComputerHistoryByName structure.
type FindComputerHistoryByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerHistoryByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerHistoryByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerhistory/name/{name}] findComputerHistoryByName", response, response.Code())
	}
}

// NewFindComputerHistoryByNameOK creates a FindComputerHistoryByNameOK with default headers values
func NewFindComputerHistoryByNameOK() *FindComputerHistoryByNameOK {
	return &FindComputerHistoryByNameOK{}
}

/*
FindComputerHistoryByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerHistoryByNameOK struct {
	Payload *models.ComputerHistory
}

// IsSuccess returns true when this find computer history by name o k response has a 2xx status code
func (o *FindComputerHistoryByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer history by name o k response has a 3xx status code
func (o *FindComputerHistoryByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer history by name o k response has a 4xx status code
func (o *FindComputerHistoryByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer history by name o k response has a 5xx status code
func (o *FindComputerHistoryByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer history by name o k response a status code equal to that given
func (o *FindComputerHistoryByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer history by name o k response
func (o *FindComputerHistoryByNameOK) Code() int {
	return 200
}

func (o *FindComputerHistoryByNameOK) Error() string {
	return fmt.Sprintf("[GET /computerhistory/name/{name}][%d] findComputerHistoryByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerHistoryByNameOK) String() string {
	return fmt.Sprintf("[GET /computerhistory/name/{name}][%d] findComputerHistoryByNameOK  %+v", 200, o.Payload)
}

func (o *FindComputerHistoryByNameOK) GetPayload() *models.ComputerHistory {
	return o.Payload
}

func (o *FindComputerHistoryByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}