// Code generated by go-swagger; DO NOT EDIT.

package managedpreferenceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateManagedPreferenceProfileByIDReader is a Reader for the CreateManagedPreferenceProfileByID structure.
type CreateManagedPreferenceProfileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateManagedPreferenceProfileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateManagedPreferenceProfileByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /managedpreferenceprofiles/id/{id}] createManagedPreferenceProfileById", response, response.Code())
	}
}

// NewCreateManagedPreferenceProfileByIDCreated creates a CreateManagedPreferenceProfileByIDCreated with default headers values
func NewCreateManagedPreferenceProfileByIDCreated() *CreateManagedPreferenceProfileByIDCreated {
	return &CreateManagedPreferenceProfileByIDCreated{}
}

/*
CreateManagedPreferenceProfileByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateManagedPreferenceProfileByIDCreated struct {
}

// IsSuccess returns true when this create managed preference profile by Id created response has a 2xx status code
func (o *CreateManagedPreferenceProfileByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create managed preference profile by Id created response has a 3xx status code
func (o *CreateManagedPreferenceProfileByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create managed preference profile by Id created response has a 4xx status code
func (o *CreateManagedPreferenceProfileByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create managed preference profile by Id created response has a 5xx status code
func (o *CreateManagedPreferenceProfileByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create managed preference profile by Id created response a status code equal to that given
func (o *CreateManagedPreferenceProfileByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create managed preference profile by Id created response
func (o *CreateManagedPreferenceProfileByIDCreated) Code() int {
	return 201
}

func (o *CreateManagedPreferenceProfileByIDCreated) Error() string {
	return fmt.Sprintf("[POST /managedpreferenceprofiles/id/{id}][%d] createManagedPreferenceProfileByIdCreated ", 201)
}

func (o *CreateManagedPreferenceProfileByIDCreated) String() string {
	return fmt.Sprintf("[POST /managedpreferenceprofiles/id/{id}][%d] createManagedPreferenceProfileByIdCreated ", 201)
}

func (o *CreateManagedPreferenceProfileByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
