// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindDirectoryBindingsReader is a Reader for the FindDirectoryBindings structure.
type FindDirectoryBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDirectoryBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDirectoryBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /directorybindings] findDirectoryBindings", response, response.Code())
	}
}

// NewFindDirectoryBindingsOK creates a FindDirectoryBindingsOK with default headers values
func NewFindDirectoryBindingsOK() *FindDirectoryBindingsOK {
	return &FindDirectoryBindingsOK{}
}

/*
FindDirectoryBindingsOK describes a response with status code 200, with default header values.

OK
*/
type FindDirectoryBindingsOK struct {
	Payload models.DirectoryBindings
}

// IsSuccess returns true when this find directory bindings o k response has a 2xx status code
func (o *FindDirectoryBindingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find directory bindings o k response has a 3xx status code
func (o *FindDirectoryBindingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find directory bindings o k response has a 4xx status code
func (o *FindDirectoryBindingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find directory bindings o k response has a 5xx status code
func (o *FindDirectoryBindingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find directory bindings o k response a status code equal to that given
func (o *FindDirectoryBindingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find directory bindings o k response
func (o *FindDirectoryBindingsOK) Code() int {
	return 200
}

func (o *FindDirectoryBindingsOK) Error() string {
	return fmt.Sprintf("[GET /directorybindings][%d] findDirectoryBindingsOK  %+v", 200, o.Payload)
}

func (o *FindDirectoryBindingsOK) String() string {
	return fmt.Sprintf("[GET /directorybindings][%d] findDirectoryBindingsOK  %+v", 200, o.Payload)
}

func (o *FindDirectoryBindingsOK) GetPayload() models.DirectoryBindings {
	return o.Payload
}

func (o *FindDirectoryBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}