// Code generated by go-swagger; DO NOT EDIT.

package advancedcomputersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindAdvancedComputerSearchesReader is a Reader for the FindAdvancedComputerSearches structure.
type FindAdvancedComputerSearchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindAdvancedComputerSearchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindAdvancedComputerSearchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindAdvancedComputerSearchesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /advancedcomputersearches] findAdvancedComputerSearches", response, response.Code())
	}
}

// NewFindAdvancedComputerSearchesOK creates a FindAdvancedComputerSearchesOK with default headers values
func NewFindAdvancedComputerSearchesOK() *FindAdvancedComputerSearchesOK {
	return &FindAdvancedComputerSearchesOK{}
}

/*
FindAdvancedComputerSearchesOK describes a response with status code 200, with default header values.

OK
*/
type FindAdvancedComputerSearchesOK struct {
	Payload models.AdvancedComputerSearches
}

// IsSuccess returns true when this find advanced computer searches o k response has a 2xx status code
func (o *FindAdvancedComputerSearchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find advanced computer searches o k response has a 3xx status code
func (o *FindAdvancedComputerSearchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find advanced computer searches o k response has a 4xx status code
func (o *FindAdvancedComputerSearchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find advanced computer searches o k response has a 5xx status code
func (o *FindAdvancedComputerSearchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find advanced computer searches o k response a status code equal to that given
func (o *FindAdvancedComputerSearchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find advanced computer searches o k response
func (o *FindAdvancedComputerSearchesOK) Code() int {
	return 200
}

func (o *FindAdvancedComputerSearchesOK) Error() string {
	return fmt.Sprintf("[GET /advancedcomputersearches][%d] findAdvancedComputerSearchesOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedComputerSearchesOK) String() string {
	return fmt.Sprintf("[GET /advancedcomputersearches][%d] findAdvancedComputerSearchesOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedComputerSearchesOK) GetPayload() models.AdvancedComputerSearches {
	return o.Payload
}

func (o *FindAdvancedComputerSearchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindAdvancedComputerSearchesUnauthorized creates a FindAdvancedComputerSearchesUnauthorized with default headers values
func NewFindAdvancedComputerSearchesUnauthorized() *FindAdvancedComputerSearchesUnauthorized {
	return &FindAdvancedComputerSearchesUnauthorized{}
}

/*
FindAdvancedComputerSearchesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FindAdvancedComputerSearchesUnauthorized struct {
}

// IsSuccess returns true when this find advanced computer searches unauthorized response has a 2xx status code
func (o *FindAdvancedComputerSearchesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find advanced computer searches unauthorized response has a 3xx status code
func (o *FindAdvancedComputerSearchesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find advanced computer searches unauthorized response has a 4xx status code
func (o *FindAdvancedComputerSearchesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this find advanced computer searches unauthorized response has a 5xx status code
func (o *FindAdvancedComputerSearchesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this find advanced computer searches unauthorized response a status code equal to that given
func (o *FindAdvancedComputerSearchesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the find advanced computer searches unauthorized response
func (o *FindAdvancedComputerSearchesUnauthorized) Code() int {
	return 401
}

func (o *FindAdvancedComputerSearchesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /advancedcomputersearches][%d] findAdvancedComputerSearchesUnauthorized ", 401)
}

func (o *FindAdvancedComputerSearchesUnauthorized) String() string {
	return fmt.Sprintf("[GET /advancedcomputersearches][%d] findAdvancedComputerSearchesUnauthorized ", 401)
}

func (o *FindAdvancedComputerSearchesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
