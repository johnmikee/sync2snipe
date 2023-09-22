// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateDirectoryBindingByIDReader is a Reader for the CreateDirectoryBindingByID structure.
type CreateDirectoryBindingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectoryBindingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDirectoryBindingByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /directorybindings/id/{id}] createDirectoryBindingById", response, response.Code())
	}
}

// NewCreateDirectoryBindingByIDCreated creates a CreateDirectoryBindingByIDCreated with default headers values
func NewCreateDirectoryBindingByIDCreated() *CreateDirectoryBindingByIDCreated {
	return &CreateDirectoryBindingByIDCreated{}
}

/*
CreateDirectoryBindingByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateDirectoryBindingByIDCreated struct {
}

// IsSuccess returns true when this create directory binding by Id created response has a 2xx status code
func (o *CreateDirectoryBindingByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create directory binding by Id created response has a 3xx status code
func (o *CreateDirectoryBindingByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create directory binding by Id created response has a 4xx status code
func (o *CreateDirectoryBindingByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create directory binding by Id created response has a 5xx status code
func (o *CreateDirectoryBindingByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create directory binding by Id created response a status code equal to that given
func (o *CreateDirectoryBindingByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create directory binding by Id created response
func (o *CreateDirectoryBindingByIDCreated) Code() int {
	return 201
}

func (o *CreateDirectoryBindingByIDCreated) Error() string {
	return fmt.Sprintf("[POST /directorybindings/id/{id}][%d] createDirectoryBindingByIdCreated ", 201)
}

func (o *CreateDirectoryBindingByIDCreated) String() string {
	return fmt.Sprintf("[POST /directorybindings/id/{id}][%d] createDirectoryBindingByIdCreated ", 201)
}

func (o *CreateDirectoryBindingByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
