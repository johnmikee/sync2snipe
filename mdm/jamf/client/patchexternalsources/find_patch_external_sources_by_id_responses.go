// Code generated by go-swagger; DO NOT EDIT.

package patchexternalsources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindPatchExternalSourcesByIDReader is a Reader for the FindPatchExternalSourcesByID structure.
type FindPatchExternalSourcesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPatchExternalSourcesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPatchExternalSourcesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /patchexternalsources/id/{id}] findPatchExternalSourcesById", response, response.Code())
	}
}

// NewFindPatchExternalSourcesByIDOK creates a FindPatchExternalSourcesByIDOK with default headers values
func NewFindPatchExternalSourcesByIDOK() *FindPatchExternalSourcesByIDOK {
	return &FindPatchExternalSourcesByIDOK{}
}

/*
FindPatchExternalSourcesByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindPatchExternalSourcesByIDOK struct {
	Payload *models.PatchExternalSource
}

// IsSuccess returns true when this find patch external sources by Id o k response has a 2xx status code
func (o *FindPatchExternalSourcesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find patch external sources by Id o k response has a 3xx status code
func (o *FindPatchExternalSourcesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find patch external sources by Id o k response has a 4xx status code
func (o *FindPatchExternalSourcesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find patch external sources by Id o k response has a 5xx status code
func (o *FindPatchExternalSourcesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find patch external sources by Id o k response a status code equal to that given
func (o *FindPatchExternalSourcesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find patch external sources by Id o k response
func (o *FindPatchExternalSourcesByIDOK) Code() int {
	return 200
}

func (o *FindPatchExternalSourcesByIDOK) Error() string {
	return fmt.Sprintf("[GET /patchexternalsources/id/{id}][%d] findPatchExternalSourcesByIdOK  %+v", 200, o.Payload)
}

func (o *FindPatchExternalSourcesByIDOK) String() string {
	return fmt.Sprintf("[GET /patchexternalsources/id/{id}][%d] findPatchExternalSourcesByIdOK  %+v", 200, o.Payload)
}

func (o *FindPatchExternalSourcesByIDOK) GetPayload() *models.PatchExternalSource {
	return o.Payload
}

func (o *FindPatchExternalSourcesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchExternalSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}