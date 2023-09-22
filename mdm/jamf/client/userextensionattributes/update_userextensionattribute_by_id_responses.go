// Code generated by go-swagger; DO NOT EDIT.

package userextensionattributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserextensionattributeByIDReader is a Reader for the UpdateUserextensionattributeByID structure.
type UpdateUserextensionattributeByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserextensionattributeByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateUserextensionattributeByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /userextensionattributes/id/{id}] updateUserextensionattributeById", response, response.Code())
	}
}

// NewUpdateUserextensionattributeByIDCreated creates a UpdateUserextensionattributeByIDCreated with default headers values
func NewUpdateUserextensionattributeByIDCreated() *UpdateUserextensionattributeByIDCreated {
	return &UpdateUserextensionattributeByIDCreated{}
}

/*
UpdateUserextensionattributeByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateUserextensionattributeByIDCreated struct {
}

// IsSuccess returns true when this update userextensionattribute by Id created response has a 2xx status code
func (o *UpdateUserextensionattributeByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update userextensionattribute by Id created response has a 3xx status code
func (o *UpdateUserextensionattributeByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update userextensionattribute by Id created response has a 4xx status code
func (o *UpdateUserextensionattributeByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update userextensionattribute by Id created response has a 5xx status code
func (o *UpdateUserextensionattributeByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update userextensionattribute by Id created response a status code equal to that given
func (o *UpdateUserextensionattributeByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update userextensionattribute by Id created response
func (o *UpdateUserextensionattributeByIDCreated) Code() int {
	return 201
}

func (o *UpdateUserextensionattributeByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /userextensionattributes/id/{id}][%d] updateUserextensionattributeByIdCreated ", 201)
}

func (o *UpdateUserextensionattributeByIDCreated) String() string {
	return fmt.Sprintf("[PUT /userextensionattributes/id/{id}][%d] updateUserextensionattributeByIdCreated ", 201)
}

func (o *UpdateUserextensionattributeByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}