// Code generated by go-swagger; DO NOT EDIT.

package advancedusersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateAdvancedUserSearchgByIDReader is a Reader for the CreateAdvancedUserSearchgByID structure.
type CreateAdvancedUserSearchgByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAdvancedUserSearchgByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAdvancedUserSearchgByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /advancedusersearches/id/{id}] createAdvancedUserSearchgById", response, response.Code())
	}
}

// NewCreateAdvancedUserSearchgByIDCreated creates a CreateAdvancedUserSearchgByIDCreated with default headers values
func NewCreateAdvancedUserSearchgByIDCreated() *CreateAdvancedUserSearchgByIDCreated {
	return &CreateAdvancedUserSearchgByIDCreated{}
}

/*
CreateAdvancedUserSearchgByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAdvancedUserSearchgByIDCreated struct {
}

// IsSuccess returns true when this create advanced user searchg by Id created response has a 2xx status code
func (o *CreateAdvancedUserSearchgByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create advanced user searchg by Id created response has a 3xx status code
func (o *CreateAdvancedUserSearchgByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create advanced user searchg by Id created response has a 4xx status code
func (o *CreateAdvancedUserSearchgByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create advanced user searchg by Id created response has a 5xx status code
func (o *CreateAdvancedUserSearchgByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create advanced user searchg by Id created response a status code equal to that given
func (o *CreateAdvancedUserSearchgByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create advanced user searchg by Id created response
func (o *CreateAdvancedUserSearchgByIDCreated) Code() int {
	return 201
}

func (o *CreateAdvancedUserSearchgByIDCreated) Error() string {
	return fmt.Sprintf("[POST /advancedusersearches/id/{id}][%d] createAdvancedUserSearchgByIdCreated ", 201)
}

func (o *CreateAdvancedUserSearchgByIDCreated) String() string {
	return fmt.Sprintf("[POST /advancedusersearches/id/{id}][%d] createAdvancedUserSearchgByIdCreated ", 201)
}

func (o *CreateAdvancedUserSearchgByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}