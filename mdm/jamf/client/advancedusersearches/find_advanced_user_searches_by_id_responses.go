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

// FindAdvancedUserSearchesByIDReader is a Reader for the FindAdvancedUserSearchesByID structure.
type FindAdvancedUserSearchesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindAdvancedUserSearchesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindAdvancedUserSearchesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /advancedusersearches/id/{id}] findAdvancedUserSearchesById", response, response.Code())
	}
}

// NewFindAdvancedUserSearchesByIDOK creates a FindAdvancedUserSearchesByIDOK with default headers values
func NewFindAdvancedUserSearchesByIDOK() *FindAdvancedUserSearchesByIDOK {
	return &FindAdvancedUserSearchesByIDOK{}
}

/*
FindAdvancedUserSearchesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindAdvancedUserSearchesByIDOK struct {
	Payload *models.AdvancedUserSearch
}

// IsSuccess returns true when this find advanced user searches by Id o k response has a 2xx status code
func (o *FindAdvancedUserSearchesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find advanced user searches by Id o k response has a 3xx status code
func (o *FindAdvancedUserSearchesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find advanced user searches by Id o k response has a 4xx status code
func (o *FindAdvancedUserSearchesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find advanced user searches by Id o k response has a 5xx status code
func (o *FindAdvancedUserSearchesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find advanced user searches by Id o k response a status code equal to that given
func (o *FindAdvancedUserSearchesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find advanced user searches by Id o k response
func (o *FindAdvancedUserSearchesByIDOK) Code() int {
	return 200
}

func (o *FindAdvancedUserSearchesByIDOK) Error() string {
	return fmt.Sprintf("[GET /advancedusersearches/id/{id}][%d] findAdvancedUserSearchesByIdOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedUserSearchesByIDOK) String() string {
	return fmt.Sprintf("[GET /advancedusersearches/id/{id}][%d] findAdvancedUserSearchesByIdOK  %+v", 200, o.Payload)
}

func (o *FindAdvancedUserSearchesByIDOK) GetPayload() *models.AdvancedUserSearch {
	return o.Payload
}

func (o *FindAdvancedUserSearchesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdvancedUserSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
