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

// FindOsxConfigurationProfilesByNameSubsetReader is a Reader for the FindOsxConfigurationProfilesByNameSubset structure.
type FindOsxConfigurationProfilesByNameSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOsxConfigurationProfilesByNameSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOsxConfigurationProfilesByNameSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /osxconfigurationprofiles/name/{name}/subset/{subset}] findOsxConfigurationProfilesByNameSubset", response, response.Code())
	}
}

// NewFindOsxConfigurationProfilesByNameSubsetOK creates a FindOsxConfigurationProfilesByNameSubsetOK with default headers values
func NewFindOsxConfigurationProfilesByNameSubsetOK() *FindOsxConfigurationProfilesByNameSubsetOK {
	return &FindOsxConfigurationProfilesByNameSubsetOK{}
}

/*
FindOsxConfigurationProfilesByNameSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindOsxConfigurationProfilesByNameSubsetOK struct {
	Payload *models.OsxConfigurationProfile
}

// IsSuccess returns true when this find osx configuration profiles by name subset o k response has a 2xx status code
func (o *FindOsxConfigurationProfilesByNameSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find osx configuration profiles by name subset o k response has a 3xx status code
func (o *FindOsxConfigurationProfilesByNameSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find osx configuration profiles by name subset o k response has a 4xx status code
func (o *FindOsxConfigurationProfilesByNameSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find osx configuration profiles by name subset o k response has a 5xx status code
func (o *FindOsxConfigurationProfilesByNameSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find osx configuration profiles by name subset o k response a status code equal to that given
func (o *FindOsxConfigurationProfilesByNameSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find osx configuration profiles by name subset o k response
func (o *FindOsxConfigurationProfilesByNameSubsetOK) Code() int {
	return 200
}

func (o *FindOsxConfigurationProfilesByNameSubsetOK) Error() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/name/{name}/subset/{subset}][%d] findOsxConfigurationProfilesByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindOsxConfigurationProfilesByNameSubsetOK) String() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/name/{name}/subset/{subset}][%d] findOsxConfigurationProfilesByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindOsxConfigurationProfilesByNameSubsetOK) GetPayload() *models.OsxConfigurationProfile {
	return o.Payload
}

func (o *FindOsxConfigurationProfilesByNameSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OsxConfigurationProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
