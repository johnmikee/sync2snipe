// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteManagedPreferenceProfilesByIDReader is a Reader for the DeleteManagedPreferenceProfilesByID structure.
type DeleteManagedPreferenceProfilesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteManagedPreferenceProfilesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteManagedPreferenceProfilesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /managedpreferenceprofiles/id/{id}] deleteManagedPreferenceProfilesById", response, response.Code())
	}
}

// NewDeleteManagedPreferenceProfilesByIDOK creates a DeleteManagedPreferenceProfilesByIDOK with default headers values
func NewDeleteManagedPreferenceProfilesByIDOK() *DeleteManagedPreferenceProfilesByIDOK {
	return &DeleteManagedPreferenceProfilesByIDOK{}
}

/*
DeleteManagedPreferenceProfilesByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteManagedPreferenceProfilesByIDOK struct {
}

// IsSuccess returns true when this delete managed preference profiles by Id o k response has a 2xx status code
func (o *DeleteManagedPreferenceProfilesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete managed preference profiles by Id o k response has a 3xx status code
func (o *DeleteManagedPreferenceProfilesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete managed preference profiles by Id o k response has a 4xx status code
func (o *DeleteManagedPreferenceProfilesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete managed preference profiles by Id o k response has a 5xx status code
func (o *DeleteManagedPreferenceProfilesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete managed preference profiles by Id o k response a status code equal to that given
func (o *DeleteManagedPreferenceProfilesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete managed preference profiles by Id o k response
func (o *DeleteManagedPreferenceProfilesByIDOK) Code() int {
	return 200
}

func (o *DeleteManagedPreferenceProfilesByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /managedpreferenceprofiles/id/{id}][%d] deleteManagedPreferenceProfilesByIdOK ", 200)
}

func (o *DeleteManagedPreferenceProfilesByIDOK) String() string {
	return fmt.Sprintf("[DELETE /managedpreferenceprofiles/id/{id}][%d] deleteManagedPreferenceProfilesByIdOK ", 200)
}

func (o *DeleteManagedPreferenceProfilesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}