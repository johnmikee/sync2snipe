// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePackageByIDReader is a Reader for the DeletePackageByID structure.
type DeletePackageByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackageByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePackageByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /packages/id/{id}] deletePackageById", response, response.Code())
	}
}

// NewDeletePackageByIDOK creates a DeletePackageByIDOK with default headers values
func NewDeletePackageByIDOK() *DeletePackageByIDOK {
	return &DeletePackageByIDOK{}
}

/*
DeletePackageByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeletePackageByIDOK struct {
}

// IsSuccess returns true when this delete package by Id o k response has a 2xx status code
func (o *DeletePackageByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete package by Id o k response has a 3xx status code
func (o *DeletePackageByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete package by Id o k response has a 4xx status code
func (o *DeletePackageByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete package by Id o k response has a 5xx status code
func (o *DeletePackageByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete package by Id o k response a status code equal to that given
func (o *DeletePackageByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete package by Id o k response
func (o *DeletePackageByIDOK) Code() int {
	return 200
}

func (o *DeletePackageByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /packages/id/{id}][%d] deletePackageByIdOK ", 200)
}

func (o *DeletePackageByIDOK) String() string {
	return fmt.Sprintf("[DELETE /packages/id/{id}][%d] deletePackageByIdOK ", 200)
}

func (o *DeletePackageByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
