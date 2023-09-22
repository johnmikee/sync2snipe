// Code generated by go-swagger; DO NOT EDIT.

package advancedusersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindAdvancedUserSearchesReader is a Reader for the FindAdvancedUserSearches structure.
type FindAdvancedUserSearchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindAdvancedUserSearchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindAdvancedUserSearchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /advancedusersearches] findAdvancedUserSearches", response, response.Code())
	}
}

// NewFindAdvancedUserSearchesOK creates a FindAdvancedUserSearchesOK with default headers values
func NewFindAdvancedUserSearchesOK() *FindAdvancedUserSearchesOK {
	return &FindAdvancedUserSearchesOK{}
}

/*
FindAdvancedUserSearchesOK describes a response with status code 200, with default header values.

OK
*/
type FindAdvancedUserSearchesOK struct {
	Payload models.AdvancedUserSearches
}

// IsSuccess returns true when this find advanced user searches o k response has a 2xx status code
func (o *FindAdvancedUserSearchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find advanced user searches o k response has a 3xx status code
func (o *FindAdvancedUserSearchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find advanced user searches o k response has a 4xx status code
func (o *FindAdvancedUserSearchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find advanced user searches o k response has a 5xx status code
func (o *FindAdvancedUserSearchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find advanced user searches o k response a status code equal to that given
func (o *FindAdvancedUserSearchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find advanced user searches o k response
func (o *FindAdvancedUserSearchesOK) Code() int {
	return 200
}

func (o *FindAdvancedUserSearchesOK) Error() string {
	return fmt.Sprintf("[GET /advancedusersearches][%d] findAdvancedUserSearchesOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedUserSearchesOK) String() string {
	return fmt.Sprintf("[GET /advancedusersearches][%d] findAdvancedUserSearchesOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedUserSearchesOK) GetPayload() models.AdvancedUserSearches {
	return o.Payload
}

func (o *FindAdvancedUserSearchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}