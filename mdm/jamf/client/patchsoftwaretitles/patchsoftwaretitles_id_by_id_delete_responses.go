// Code generated by go-swagger; DO NOT EDIT.

package patchsoftwaretitles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchsoftwaretitlesIDByIDDeleteReader is a Reader for the PatchsoftwaretitlesIDByIDDelete structure.
type PatchsoftwaretitlesIDByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchsoftwaretitlesIDByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchsoftwaretitlesIDByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /patchsoftwaretitles/id/{id}] PatchsoftwaretitlesIdByIdDelete", response, response.Code())
	}
}

// NewPatchsoftwaretitlesIDByIDDeleteOK creates a PatchsoftwaretitlesIDByIDDeleteOK with default headers values
func NewPatchsoftwaretitlesIDByIDDeleteOK() *PatchsoftwaretitlesIDByIDDeleteOK {
	return &PatchsoftwaretitlesIDByIDDeleteOK{}
}

/*
PatchsoftwaretitlesIDByIDDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PatchsoftwaretitlesIDByIDDeleteOK struct {
}

// IsSuccess returns true when this patchsoftwaretitles Id by Id delete o k response has a 2xx status code
func (o *PatchsoftwaretitlesIDByIDDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patchsoftwaretitles Id by Id delete o k response has a 3xx status code
func (o *PatchsoftwaretitlesIDByIDDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchsoftwaretitles Id by Id delete o k response has a 4xx status code
func (o *PatchsoftwaretitlesIDByIDDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patchsoftwaretitles Id by Id delete o k response has a 5xx status code
func (o *PatchsoftwaretitlesIDByIDDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patchsoftwaretitles Id by Id delete o k response a status code equal to that given
func (o *PatchsoftwaretitlesIDByIDDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patchsoftwaretitles Id by Id delete o k response
func (o *PatchsoftwaretitlesIDByIDDeleteOK) Code() int {
	return 200
}

func (o *PatchsoftwaretitlesIDByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdDeleteOK ", 200)
}

func (o *PatchsoftwaretitlesIDByIDDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdDeleteOK ", 200)
}

func (o *PatchsoftwaretitlesIDByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}