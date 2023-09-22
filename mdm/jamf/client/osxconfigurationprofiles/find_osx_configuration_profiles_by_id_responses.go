// Code generated by go-swagger; DO NOT EDIT.

package osxconfigurationprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindOsxConfigurationProfilesByIDReader is a Reader for the FindOsxConfigurationProfilesByID structure.
type FindOsxConfigurationProfilesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOsxConfigurationProfilesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOsxConfigurationProfilesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /osxconfigurationprofiles/id/{id}] findOsxConfigurationProfilesById", response, response.Code())
	}
}

// NewFindOsxConfigurationProfilesByIDOK creates a FindOsxConfigurationProfilesByIDOK with default headers values
func NewFindOsxConfigurationProfilesByIDOK() *FindOsxConfigurationProfilesByIDOK {
	return &FindOsxConfigurationProfilesByIDOK{}
}

/*
FindOsxConfigurationProfilesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindOsxConfigurationProfilesByIDOK struct {
	Payload *models.OsxConfigurationProfile
}

// IsSuccess returns true when this find osx configuration profiles by Id o k response has a 2xx status code
func (o *FindOsxConfigurationProfilesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find osx configuration profiles by Id o k response has a 3xx status code
func (o *FindOsxConfigurationProfilesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find osx configuration profiles by Id o k response has a 4xx status code
func (o *FindOsxConfigurationProfilesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find osx configuration profiles by Id o k response has a 5xx status code
func (o *FindOsxConfigurationProfilesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find osx configuration profiles by Id o k response a status code equal to that given
func (o *FindOsxConfigurationProfilesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find osx configuration profiles by Id o k response
func (o *FindOsxConfigurationProfilesByIDOK) Code() int {
	return 200
}

func (o *FindOsxConfigurationProfilesByIDOK) Error() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/id/{id}][%d] findOsxConfigurationProfilesByIdOK  %+v", 200, o.Payload)
}

func (o *FindOsxConfigurationProfilesByIDOK) String() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/id/{id}][%d] findOsxConfigurationProfilesByIdOK  %+v", 200, o.Payload)
}

func (o *FindOsxConfigurationProfilesByIDOK) GetPayload() *models.OsxConfigurationProfile {
	return o.Payload
}

func (o *FindOsxConfigurationProfilesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OsxConfigurationProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}