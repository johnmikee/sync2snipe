// Code generated by go-swagger; DO NOT EDIT.

package patchsoftwaretitles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchsoftwaretitlesIDByIDPutReader is a Reader for the PatchsoftwaretitlesIDByIDPut structure.
type PatchsoftwaretitlesIDByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchsoftwaretitlesIDByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchsoftwaretitlesIDByIDPutCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /patchsoftwaretitles/id/{id}] PatchsoftwaretitlesIdByIdPut", response, response.Code())
	}
}

// NewPatchsoftwaretitlesIDByIDPutCreated creates a PatchsoftwaretitlesIDByIDPutCreated with default headers values
func NewPatchsoftwaretitlesIDByIDPutCreated() *PatchsoftwaretitlesIDByIDPutCreated {
	return &PatchsoftwaretitlesIDByIDPutCreated{}
}

/*
PatchsoftwaretitlesIDByIDPutCreated describes a response with status code 201, with default header values.

Created
*/
type PatchsoftwaretitlesIDByIDPutCreated struct {
}

// IsSuccess returns true when this patchsoftwaretitles Id by Id put created response has a 2xx status code
func (o *PatchsoftwaretitlesIDByIDPutCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patchsoftwaretitles Id by Id put created response has a 3xx status code
func (o *PatchsoftwaretitlesIDByIDPutCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchsoftwaretitles Id by Id put created response has a 4xx status code
func (o *PatchsoftwaretitlesIDByIDPutCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this patchsoftwaretitles Id by Id put created response has a 5xx status code
func (o *PatchsoftwaretitlesIDByIDPutCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this patchsoftwaretitles Id by Id put created response a status code equal to that given
func (o *PatchsoftwaretitlesIDByIDPutCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the patchsoftwaretitles Id by Id put created response
func (o *PatchsoftwaretitlesIDByIDPutCreated) Code() int {
	return 201
}

func (o *PatchsoftwaretitlesIDByIDPutCreated) Error() string {
	return fmt.Sprintf("[PUT /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdPutCreated ", 201)
}

func (o *PatchsoftwaretitlesIDByIDPutCreated) String() string {
	return fmt.Sprintf("[PUT /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdPutCreated ", 201)
}

func (o *PatchsoftwaretitlesIDByIDPutCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}