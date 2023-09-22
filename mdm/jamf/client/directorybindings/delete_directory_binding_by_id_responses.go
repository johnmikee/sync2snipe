// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDirectoryBindingByIDReader is a Reader for the DeleteDirectoryBindingByID structure.
type DeleteDirectoryBindingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDirectoryBindingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDirectoryBindingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /directorybindings/id/{id}] deleteDirectoryBindingById", response, response.Code())
	}
}

// NewDeleteDirectoryBindingByIDOK creates a DeleteDirectoryBindingByIDOK with default headers values
func NewDeleteDirectoryBindingByIDOK() *DeleteDirectoryBindingByIDOK {
	return &DeleteDirectoryBindingByIDOK{}
}

/*
DeleteDirectoryBindingByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDirectoryBindingByIDOK struct {
}

// IsSuccess returns true when this delete directory binding by Id o k response has a 2xx status code
func (o *DeleteDirectoryBindingByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete directory binding by Id o k response has a 3xx status code
func (o *DeleteDirectoryBindingByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete directory binding by Id o k response has a 4xx status code
func (o *DeleteDirectoryBindingByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete directory binding by Id o k response has a 5xx status code
func (o *DeleteDirectoryBindingByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete directory binding by Id o k response a status code equal to that given
func (o *DeleteDirectoryBindingByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete directory binding by Id o k response
func (o *DeleteDirectoryBindingByIDOK) Code() int {
	return 200
}

func (o *DeleteDirectoryBindingByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /directorybindings/id/{id}][%d] deleteDirectoryBindingByIdOK ", 200)
}

func (o *DeleteDirectoryBindingByIDOK) String() string {
	return fmt.Sprintf("[DELETE /directorybindings/id/{id}][%d] deleteDirectoryBindingByIdOK ", 200)
}

func (o *DeleteDirectoryBindingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
