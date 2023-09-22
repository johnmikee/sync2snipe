// Code generated by go-swagger; DO NOT EDIT.

package patches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSoftwareTitlesByIDReader is a Reader for the UpdateSoftwareTitlesByID structure.
type UpdateSoftwareTitlesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSoftwareTitlesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateSoftwareTitlesByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /patches/id/{id}] updateSoftwareTitlesById", response, response.Code())
	}
}

// NewUpdateSoftwareTitlesByIDCreated creates a UpdateSoftwareTitlesByIDCreated with default headers values
func NewUpdateSoftwareTitlesByIDCreated() *UpdateSoftwareTitlesByIDCreated {
	return &UpdateSoftwareTitlesByIDCreated{}
}

/*
UpdateSoftwareTitlesByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateSoftwareTitlesByIDCreated struct {
}

// IsSuccess returns true when this update software titles by Id created response has a 2xx status code
func (o *UpdateSoftwareTitlesByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update software titles by Id created response has a 3xx status code
func (o *UpdateSoftwareTitlesByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update software titles by Id created response has a 4xx status code
func (o *UpdateSoftwareTitlesByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update software titles by Id created response has a 5xx status code
func (o *UpdateSoftwareTitlesByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update software titles by Id created response a status code equal to that given
func (o *UpdateSoftwareTitlesByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update software titles by Id created response
func (o *UpdateSoftwareTitlesByIDCreated) Code() int {
	return 201
}

func (o *UpdateSoftwareTitlesByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /patches/id/{id}][%d] updateSoftwareTitlesByIdCreated ", 201)
}

func (o *UpdateSoftwareTitlesByIDCreated) String() string {
	return fmt.Sprintf("[PUT /patches/id/{id}][%d] updateSoftwareTitlesByIdCreated ", 201)
}

func (o *UpdateSoftwareTitlesByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}