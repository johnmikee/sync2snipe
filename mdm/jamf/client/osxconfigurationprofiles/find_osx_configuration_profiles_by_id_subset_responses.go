// Code generated by go-swagger; DO NOT EDIT.

package osxconfigurationprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindOsxConfigurationProfilesByIDSubsetReader is a Reader for the FindOsxConfigurationProfilesByIDSubset structure.
type FindOsxConfigurationProfilesByIDSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOsxConfigurationProfilesByIDSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOsxConfigurationProfilesByIDSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /osxconfigurationprofiles/id/{id}/subset/{subset}] findOsxConfigurationProfilesByIdSubset", response, response.Code())
	}
}

// NewFindOsxConfigurationProfilesByIDSubsetOK creates a FindOsxConfigurationProfilesByIDSubsetOK with default headers values
func NewFindOsxConfigurationProfilesByIDSubsetOK() *FindOsxConfigurationProfilesByIDSubsetOK {
	return &FindOsxConfigurationProfilesByIDSubsetOK{}
}

/*
FindOsxConfigurationProfilesByIDSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindOsxConfigurationProfilesByIDSubsetOK struct {
}

// IsSuccess returns true when this find osx configuration profiles by Id subset o k response has a 2xx status code
func (o *FindOsxConfigurationProfilesByIDSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find osx configuration profiles by Id subset o k response has a 3xx status code
func (o *FindOsxConfigurationProfilesByIDSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find osx configuration profiles by Id subset o k response has a 4xx status code
func (o *FindOsxConfigurationProfilesByIDSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find osx configuration profiles by Id subset o k response has a 5xx status code
func (o *FindOsxConfigurationProfilesByIDSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find osx configuration profiles by Id subset o k response a status code equal to that given
func (o *FindOsxConfigurationProfilesByIDSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find osx configuration profiles by Id subset o k response
func (o *FindOsxConfigurationProfilesByIDSubsetOK) Code() int {
	return 200
}

func (o *FindOsxConfigurationProfilesByIDSubsetOK) Error() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/id/{id}/subset/{subset}][%d] findOsxConfigurationProfilesByIdSubsetOK ", 200)
}

func (o *FindOsxConfigurationProfilesByIDSubsetOK) String() string {
	return fmt.Sprintf("[GET /osxconfigurationprofiles/id/{id}/subset/{subset}][%d] findOsxConfigurationProfilesByIdSubsetOK ", 200)
}

func (o *FindOsxConfigurationProfilesByIDSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
