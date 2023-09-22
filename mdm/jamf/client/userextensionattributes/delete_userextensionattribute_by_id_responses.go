// Code generated by go-swagger; DO NOT EDIT.

package userextensionattributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserextensionattributeByIDReader is a Reader for the DeleteUserextensionattributeByID structure.
type DeleteUserextensionattributeByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserextensionattributeByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserextensionattributeByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /userextensionattributes/id/{id}] deleteUserextensionattributeById", response, response.Code())
	}
}

// NewDeleteUserextensionattributeByIDOK creates a DeleteUserextensionattributeByIDOK with default headers values
func NewDeleteUserextensionattributeByIDOK() *DeleteUserextensionattributeByIDOK {
	return &DeleteUserextensionattributeByIDOK{}
}

/*
DeleteUserextensionattributeByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserextensionattributeByIDOK struct {
}

// IsSuccess returns true when this delete userextensionattribute by Id o k response has a 2xx status code
func (o *DeleteUserextensionattributeByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete userextensionattribute by Id o k response has a 3xx status code
func (o *DeleteUserextensionattributeByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userextensionattribute by Id o k response has a 4xx status code
func (o *DeleteUserextensionattributeByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete userextensionattribute by Id o k response has a 5xx status code
func (o *DeleteUserextensionattributeByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userextensionattribute by Id o k response a status code equal to that given
func (o *DeleteUserextensionattributeByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete userextensionattribute by Id o k response
func (o *DeleteUserextensionattributeByIDOK) Code() int {
	return 200
}

func (o *DeleteUserextensionattributeByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /userextensionattributes/id/{id}][%d] deleteUserextensionattributeByIdOK ", 200)
}

func (o *DeleteUserextensionattributeByIDOK) String() string {
	return fmt.Sprintf("[DELETE /userextensionattributes/id/{id}][%d] deleteUserextensionattributeByIdOK ", 200)
}

func (o *DeleteUserextensionattributeByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
