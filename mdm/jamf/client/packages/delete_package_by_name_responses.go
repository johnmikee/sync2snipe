// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePackageByNameReader is a Reader for the DeletePackageByName structure.
type DeletePackageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePackageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /packages/name/{name}] deletePackageByName", response, response.Code())
	}
}

// NewDeletePackageByNameOK creates a DeletePackageByNameOK with default headers values
func NewDeletePackageByNameOK() *DeletePackageByNameOK {
	return &DeletePackageByNameOK{}
}

/*
DeletePackageByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeletePackageByNameOK struct {
}

// IsSuccess returns true when this delete package by name o k response has a 2xx status code
func (o *DeletePackageByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete package by name o k response has a 3xx status code
func (o *DeletePackageByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete package by name o k response has a 4xx status code
func (o *DeletePackageByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete package by name o k response has a 5xx status code
func (o *DeletePackageByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete package by name o k response a status code equal to that given
func (o *DeletePackageByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete package by name o k response
func (o *DeletePackageByNameOK) Code() int {
	return 200
}

func (o *DeletePackageByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /packages/name/{name}][%d] deletePackageByNameOK ", 200)
}

func (o *DeletePackageByNameOK) String() string {
	return fmt.Sprintf("[DELETE /packages/name/{name}][%d] deletePackageByNameOK ", 200)
}

func (o *DeletePackageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}