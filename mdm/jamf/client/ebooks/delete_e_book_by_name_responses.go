// Code generated by go-swagger; DO NOT EDIT.

package ebooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEBookByNameReader is a Reader for the DeleteEBookByName structure.
type DeleteEBookByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEBookByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEBookByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ebooks/name/{name}] deleteEBookByName", response, response.Code())
	}
}

// NewDeleteEBookByNameOK creates a DeleteEBookByNameOK with default headers values
func NewDeleteEBookByNameOK() *DeleteEBookByNameOK {
	return &DeleteEBookByNameOK{}
}

/*
DeleteEBookByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteEBookByNameOK struct {
}

// IsSuccess returns true when this delete e book by name o k response has a 2xx status code
func (o *DeleteEBookByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete e book by name o k response has a 3xx status code
func (o *DeleteEBookByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete e book by name o k response has a 4xx status code
func (o *DeleteEBookByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete e book by name o k response has a 5xx status code
func (o *DeleteEBookByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete e book by name o k response a status code equal to that given
func (o *DeleteEBookByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete e book by name o k response
func (o *DeleteEBookByNameOK) Code() int {
	return 200
}

func (o *DeleteEBookByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /ebooks/name/{name}][%d] deleteEBookByNameOK ", 200)
}

func (o *DeleteEBookByNameOK) String() string {
	return fmt.Sprintf("[DELETE /ebooks/name/{name}][%d] deleteEBookByNameOK ", 200)
}

func (o *DeleteEBookByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}