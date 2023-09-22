// Code generated by go-swagger; DO NOT EDIT.

package directorybindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDirectoryBindingByNameReader is a Reader for the DeleteDirectoryBindingByName structure.
type DeleteDirectoryBindingByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDirectoryBindingByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDirectoryBindingByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /directorybindings/name/{name}] deleteDirectoryBindingByName", response, response.Code())
	}
}

// NewDeleteDirectoryBindingByNameOK creates a DeleteDirectoryBindingByNameOK with default headers values
func NewDeleteDirectoryBindingByNameOK() *DeleteDirectoryBindingByNameOK {
	return &DeleteDirectoryBindingByNameOK{}
}

/*
DeleteDirectoryBindingByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDirectoryBindingByNameOK struct {
}

// IsSuccess returns true when this delete directory binding by name o k response has a 2xx status code
func (o *DeleteDirectoryBindingByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete directory binding by name o k response has a 3xx status code
func (o *DeleteDirectoryBindingByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete directory binding by name o k response has a 4xx status code
func (o *DeleteDirectoryBindingByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete directory binding by name o k response has a 5xx status code
func (o *DeleteDirectoryBindingByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete directory binding by name o k response a status code equal to that given
func (o *DeleteDirectoryBindingByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete directory binding by name o k response
func (o *DeleteDirectoryBindingByNameOK) Code() int {
	return 200
}

func (o *DeleteDirectoryBindingByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /directorybindings/name/{name}][%d] deleteDirectoryBindingByNameOK ", 200)
}

func (o *DeleteDirectoryBindingByNameOK) String() string {
	return fmt.Sprintf("[DELETE /directorybindings/name/{name}][%d] deleteDirectoryBindingByNameOK ", 200)
}

func (o *DeleteDirectoryBindingByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
