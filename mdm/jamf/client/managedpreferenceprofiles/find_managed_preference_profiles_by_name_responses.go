// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindManagedPreferenceProfilesByNameReader is a Reader for the FindManagedPreferenceProfilesByName structure.
type FindManagedPreferenceProfilesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindManagedPreferenceProfilesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindManagedPreferenceProfilesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /managedpreferenceprofiles/name/{name}] findManagedPreferenceProfilesByName", response, response.Code())
	}
}

// NewFindManagedPreferenceProfilesByNameOK creates a FindManagedPreferenceProfilesByNameOK with default headers values
func NewFindManagedPreferenceProfilesByNameOK() *FindManagedPreferenceProfilesByNameOK {
	return &FindManagedPreferenceProfilesByNameOK{}
}

/*
FindManagedPreferenceProfilesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindManagedPreferenceProfilesByNameOK struct {
	Payload *models.ManagedPreferenceProfile
}

// IsSuccess returns true when this find managed preference profiles by name o k response has a 2xx status code
func (o *FindManagedPreferenceProfilesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find managed preference profiles by name o k response has a 3xx status code
func (o *FindManagedPreferenceProfilesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find managed preference profiles by name o k response has a 4xx status code
func (o *FindManagedPreferenceProfilesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find managed preference profiles by name o k response has a 5xx status code
func (o *FindManagedPreferenceProfilesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find managed preference profiles by name o k response a status code equal to that given
func (o *FindManagedPreferenceProfilesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find managed preference profiles by name o k response
func (o *FindManagedPreferenceProfilesByNameOK) Code() int {
	return 200
}

func (o *FindManagedPreferenceProfilesByNameOK) Error() string {
	return fmt.Sprintf("[GET /managedpreferenceprofiles/name/{name}][%d] findManagedPreferenceProfilesByNameOK  %+v", 200, o.Payload)
}

func (o *FindManagedPreferenceProfilesByNameOK) String() string {
	return fmt.Sprintf("[GET /managedpreferenceprofiles/name/{name}][%d] findManagedPreferenceProfilesByNameOK  %+v", 200, o.Payload)
}

func (o *FindManagedPreferenceProfilesByNameOK) GetPayload() *models.ManagedPreferenceProfile {
	return o.Payload
}

func (o *FindManagedPreferenceProfilesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagedPreferenceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}