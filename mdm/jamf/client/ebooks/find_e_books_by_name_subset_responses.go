// Code generated by go-swagger; DO NOT EDIT.

package ebooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindEBooksByNameSubsetReader is a Reader for the FindEBooksByNameSubset structure.
type FindEBooksByNameSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindEBooksByNameSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindEBooksByNameSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ebooks/name/{name}/subset/{subset}] findEBooksByNameSubset", response, response.Code())
	}
}

// NewFindEBooksByNameSubsetOK creates a FindEBooksByNameSubsetOK with default headers values
func NewFindEBooksByNameSubsetOK() *FindEBooksByNameSubsetOK {
	return &FindEBooksByNameSubsetOK{}
}

/*
FindEBooksByNameSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindEBooksByNameSubsetOK struct {
	Payload *models.Ebook
}

// IsSuccess returns true when this find e books by name subset o k response has a 2xx status code
func (o *FindEBooksByNameSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find e books by name subset o k response has a 3xx status code
func (o *FindEBooksByNameSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find e books by name subset o k response has a 4xx status code
func (o *FindEBooksByNameSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find e books by name subset o k response has a 5xx status code
func (o *FindEBooksByNameSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find e books by name subset o k response a status code equal to that given
func (o *FindEBooksByNameSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find e books by name subset o k response
func (o *FindEBooksByNameSubsetOK) Code() int {
	return 200
}

func (o *FindEBooksByNameSubsetOK) Error() string {
	return fmt.Sprintf("[GET /ebooks/name/{name}/subset/{subset}][%d] findEBooksByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindEBooksByNameSubsetOK) String() string {
	return fmt.Sprintf("[GET /ebooks/name/{name}/subset/{subset}][%d] findEBooksByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindEBooksByNameSubsetOK) GetPayload() *models.Ebook {
	return o.Payload
}

func (o *FindEBooksByNameSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ebook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}