// Code generated by go-swagger; DO NOT EDIT.

package patchsoftwaretitles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// PatchsoftwaretitlesIDByIDGetReader is a Reader for the PatchsoftwaretitlesIDByIDGet structure.
type PatchsoftwaretitlesIDByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchsoftwaretitlesIDByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchsoftwaretitlesIDByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /patchsoftwaretitles/id/{id}] PatchsoftwaretitlesIdByIdGet", response, response.Code())
	}
}

// NewPatchsoftwaretitlesIDByIDGetOK creates a PatchsoftwaretitlesIDByIDGetOK with default headers values
func NewPatchsoftwaretitlesIDByIDGetOK() *PatchsoftwaretitlesIDByIDGetOK {
	return &PatchsoftwaretitlesIDByIDGetOK{}
}

/*
PatchsoftwaretitlesIDByIDGetOK describes a response with status code 200, with default header values.

OK
*/
type PatchsoftwaretitlesIDByIDGetOK struct {
	Payload *models.PatchSoftwareTitle
}

// IsSuccess returns true when this patchsoftwaretitles Id by Id get o k response has a 2xx status code
func (o *PatchsoftwaretitlesIDByIDGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patchsoftwaretitles Id by Id get o k response has a 3xx status code
func (o *PatchsoftwaretitlesIDByIDGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchsoftwaretitles Id by Id get o k response has a 4xx status code
func (o *PatchsoftwaretitlesIDByIDGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patchsoftwaretitles Id by Id get o k response has a 5xx status code
func (o *PatchsoftwaretitlesIDByIDGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patchsoftwaretitles Id by Id get o k response a status code equal to that given
func (o *PatchsoftwaretitlesIDByIDGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patchsoftwaretitles Id by Id get o k response
func (o *PatchsoftwaretitlesIDByIDGetOK) Code() int {
	return 200
}

func (o *PatchsoftwaretitlesIDByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdGetOK  %+v", 200, o.Payload)
}

func (o *PatchsoftwaretitlesIDByIDGetOK) String() string {
	return fmt.Sprintf("[GET /patchsoftwaretitles/id/{id}][%d] patchsoftwaretitlesIdByIdGetOK  %+v", 200, o.Payload)
}

func (o *PatchsoftwaretitlesIDByIDGetOK) GetPayload() *models.PatchSoftwareTitle {
	return o.Payload
}

func (o *PatchsoftwaretitlesIDByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchSoftwareTitle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
