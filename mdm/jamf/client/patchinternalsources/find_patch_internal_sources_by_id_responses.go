// Code generated by go-swagger; DO NOT EDIT.

package patchinternalsources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindPatchInternalSourcesByIDReader is a Reader for the FindPatchInternalSourcesByID structure.
type FindPatchInternalSourcesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPatchInternalSourcesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPatchInternalSourcesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /patchinternalsources/id/{id}] findPatchInternalSourcesById", response, response.Code())
	}
}

// NewFindPatchInternalSourcesByIDOK creates a FindPatchInternalSourcesByIDOK with default headers values
func NewFindPatchInternalSourcesByIDOK() *FindPatchInternalSourcesByIDOK {
	return &FindPatchInternalSourcesByIDOK{}
}

/*
FindPatchInternalSourcesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindPatchInternalSourcesByIDOK struct {
	Payload *models.PatchInternalSource
}

// IsSuccess returns true when this find patch internal sources by Id o k response has a 2xx status code
func (o *FindPatchInternalSourcesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find patch internal sources by Id o k response has a 3xx status code
func (o *FindPatchInternalSourcesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find patch internal sources by Id o k response has a 4xx status code
func (o *FindPatchInternalSourcesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find patch internal sources by Id o k response has a 5xx status code
func (o *FindPatchInternalSourcesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find patch internal sources by Id o k response a status code equal to that given
func (o *FindPatchInternalSourcesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find patch internal sources by Id o k response
func (o *FindPatchInternalSourcesByIDOK) Code() int {
	return 200
}

func (o *FindPatchInternalSourcesByIDOK) Error() string {
	return fmt.Sprintf("[GET /patchinternalsources/id/{id}][%d] findPatchInternalSourcesByIdOK  %+v", 200, o.Payload)
}

func (o *FindPatchInternalSourcesByIDOK) String() string {
	return fmt.Sprintf("[GET /patchinternalsources/id/{id}][%d] findPatchInternalSourcesByIdOK  %+v", 200, o.Payload)
}

func (o *FindPatchInternalSourcesByIDOK) GetPayload() *models.PatchInternalSource {
	return o.Payload
}

func (o *FindPatchInternalSourcesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchInternalSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}