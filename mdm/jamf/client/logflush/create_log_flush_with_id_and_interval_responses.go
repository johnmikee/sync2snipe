// Code generated by go-swagger; DO NOT EDIT.

package logflush

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateLogFlushWithIDAndIntervalReader is a Reader for the CreateLogFlushWithIDAndInterval structure.
type CreateLogFlushWithIDAndIntervalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLogFlushWithIDAndIntervalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLogFlushWithIDAndIntervalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /logflush/{log}/id/{id}/interval/{interval}] createLogFlushWithIdAndInterval", response, response.Code())
	}
}

// NewCreateLogFlushWithIDAndIntervalOK creates a CreateLogFlushWithIDAndIntervalOK with default headers values
func NewCreateLogFlushWithIDAndIntervalOK() *CreateLogFlushWithIDAndIntervalOK {
	return &CreateLogFlushWithIDAndIntervalOK{}
}

/*
CreateLogFlushWithIDAndIntervalOK describes a response with status code 200, with default header values.

OK
*/
type CreateLogFlushWithIDAndIntervalOK struct {
}

// IsSuccess returns true when this create log flush with Id and interval o k response has a 2xx status code
func (o *CreateLogFlushWithIDAndIntervalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create log flush with Id and interval o k response has a 3xx status code
func (o *CreateLogFlushWithIDAndIntervalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create log flush with Id and interval o k response has a 4xx status code
func (o *CreateLogFlushWithIDAndIntervalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create log flush with Id and interval o k response has a 5xx status code
func (o *CreateLogFlushWithIDAndIntervalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create log flush with Id and interval o k response a status code equal to that given
func (o *CreateLogFlushWithIDAndIntervalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create log flush with Id and interval o k response
func (o *CreateLogFlushWithIDAndIntervalOK) Code() int {
	return 200
}

func (o *CreateLogFlushWithIDAndIntervalOK) Error() string {
	return fmt.Sprintf("[DELETE /logflush/{log}/id/{id}/interval/{interval}][%d] createLogFlushWithIdAndIntervalOK ", 200)
}

func (o *CreateLogFlushWithIDAndIntervalOK) String() string {
	return fmt.Sprintf("[DELETE /logflush/{log}/id/{id}/interval/{interval}][%d] createLogFlushWithIdAndIntervalOK ", 200)
}

func (o *CreateLogFlushWithIDAndIntervalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
