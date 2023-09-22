// Code generated by go-swagger; DO NOT EDIT.

package patches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePatchSoftwareTitlesByIDReader is a Reader for the CreatePatchSoftwareTitlesByID structure.
type CreatePatchSoftwareTitlesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePatchSoftwareTitlesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePatchSoftwareTitlesByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /patches/id/{id}] createPatchSoftwareTitlesById", response, response.Code())
	}
}

// NewCreatePatchSoftwareTitlesByIDCreated creates a CreatePatchSoftwareTitlesByIDCreated with default headers values
func NewCreatePatchSoftwareTitlesByIDCreated() *CreatePatchSoftwareTitlesByIDCreated {
	return &CreatePatchSoftwareTitlesByIDCreated{}
}

/*
CreatePatchSoftwareTitlesByIDCreated describes a response with status code 201, with default header values.

created
*/
type CreatePatchSoftwareTitlesByIDCreated struct {
}

// IsSuccess returns true when this create patch software titles by Id created response has a 2xx status code
func (o *CreatePatchSoftwareTitlesByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create patch software titles by Id created response has a 3xx status code
func (o *CreatePatchSoftwareTitlesByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create patch software titles by Id created response has a 4xx status code
func (o *CreatePatchSoftwareTitlesByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create patch software titles by Id created response has a 5xx status code
func (o *CreatePatchSoftwareTitlesByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create patch software titles by Id created response a status code equal to that given
func (o *CreatePatchSoftwareTitlesByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create patch software titles by Id created response
func (o *CreatePatchSoftwareTitlesByIDCreated) Code() int {
	return 201
}

func (o *CreatePatchSoftwareTitlesByIDCreated) Error() string {
	return fmt.Sprintf("[POST /patches/id/{id}][%d] createPatchSoftwareTitlesByIdCreated ", 201)
}

func (o *CreatePatchSoftwareTitlesByIDCreated) String() string {
	return fmt.Sprintf("[POST /patches/id/{id}][%d] createPatchSoftwareTitlesByIdCreated ", 201)
}

func (o *CreatePatchSoftwareTitlesByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}