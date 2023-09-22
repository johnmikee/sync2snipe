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

// FindManagedPreferenceProfilesReader is a Reader for the FindManagedPreferenceProfiles structure.
type FindManagedPreferenceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindManagedPreferenceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindManagedPreferenceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /managedpreferenceprofiles] findManagedPreferenceProfiles", response, response.Code())
	}
}

// NewFindManagedPreferenceProfilesOK creates a FindManagedPreferenceProfilesOK with default headers values
func NewFindManagedPreferenceProfilesOK() *FindManagedPreferenceProfilesOK {
	return &FindManagedPreferenceProfilesOK{}
}

/*
FindManagedPreferenceProfilesOK describes a response with status code 200, with default header values.

OK
*/
type FindManagedPreferenceProfilesOK struct {
	Payload models.ManagedPreferenceProfiles
}

// IsSuccess returns true when this find managed preference profiles o k response has a 2xx status code
func (o *FindManagedPreferenceProfilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find managed preference profiles o k response has a 3xx status code
func (o *FindManagedPreferenceProfilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find managed preference profiles o k response has a 4xx status code
func (o *FindManagedPreferenceProfilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find managed preference profiles o k response has a 5xx status code
func (o *FindManagedPreferenceProfilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find managed preference profiles o k response a status code equal to that given
func (o *FindManagedPreferenceProfilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find managed preference profiles o k response
func (o *FindManagedPreferenceProfilesOK) Code() int {
	return 200
}

func (o *FindManagedPreferenceProfilesOK) Error() string {
	return fmt.Sprintf("[GET /managedpreferenceprofiles][%d] findManagedPreferenceProfilesOK  %+v", 200, o.Payload)
}

func (o *FindManagedPreferenceProfilesOK) String() string {
	return fmt.Sprintf("[GET /managedpreferenceprofiles][%d] findManagedPreferenceProfilesOK  %+v", 200, o.Payload)
}

func (o *FindManagedPreferenceProfilesOK) GetPayload() models.ManagedPreferenceProfiles {
	return o.Payload
}

func (o *FindManagedPreferenceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
