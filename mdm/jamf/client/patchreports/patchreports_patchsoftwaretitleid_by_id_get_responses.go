// Code generated by go-swagger; DO NOT EDIT.

package patchreports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// PatchreportsPatchsoftwaretitleidByIDGetReader is a Reader for the PatchreportsPatchsoftwaretitleidByIDGet structure.
type PatchreportsPatchsoftwaretitleidByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchreportsPatchsoftwaretitleidByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchreportsPatchsoftwaretitleidByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /patchreports/patchsoftwaretitleid/{id}] PatchreportsPatchsoftwaretitleidByIdGet", response, response.Code())
	}
}

// NewPatchreportsPatchsoftwaretitleidByIDGetOK creates a PatchreportsPatchsoftwaretitleidByIDGetOK with default headers values
func NewPatchreportsPatchsoftwaretitleidByIDGetOK() *PatchreportsPatchsoftwaretitleidByIDGetOK {
	return &PatchreportsPatchsoftwaretitleidByIDGetOK{}
}

/*
PatchreportsPatchsoftwaretitleidByIDGetOK describes a response with status code 200, with default header values.

OK
*/
type PatchreportsPatchsoftwaretitleidByIDGetOK struct {
	Payload *models.PatchReport
}

// IsSuccess returns true when this patchreports patchsoftwaretitleid by Id get o k response has a 2xx status code
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patchreports patchsoftwaretitleid by Id get o k response has a 3xx status code
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patchreports patchsoftwaretitleid by Id get o k response has a 4xx status code
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patchreports patchsoftwaretitleid by Id get o k response has a 5xx status code
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patchreports patchsoftwaretitleid by Id get o k response a status code equal to that given
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patchreports patchsoftwaretitleid by Id get o k response
func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) Code() int {
	return 200
}

func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /patchreports/patchsoftwaretitleid/{id}][%d] patchreportsPatchsoftwaretitleidByIdGetOK  %+v", 200, o.Payload)
}

func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) String() string {
	return fmt.Sprintf("[GET /patchreports/patchsoftwaretitleid/{id}][%d] patchreportsPatchsoftwaretitleidByIdGetOK  %+v", 200, o.Payload)
}

func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) GetPayload() *models.PatchReport {
	return o.Payload
}

func (o *PatchreportsPatchsoftwaretitleidByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}