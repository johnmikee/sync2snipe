// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindMobileDeviceEnrollmentProfilesByIDSubsetReader is a Reader for the FindMobileDeviceEnrollmentProfilesByIDSubset structure.
type FindMobileDeviceEnrollmentProfilesByIDSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceEnrollmentProfilesByIDSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}] findMobileDeviceEnrollmentProfilesByIdSubset", response, response.Code())
	}
}

// NewFindMobileDeviceEnrollmentProfilesByIDSubsetOK creates a FindMobileDeviceEnrollmentProfilesByIDSubsetOK with default headers values
func NewFindMobileDeviceEnrollmentProfilesByIDSubsetOK() *FindMobileDeviceEnrollmentProfilesByIDSubsetOK {
	return &FindMobileDeviceEnrollmentProfilesByIDSubsetOK{}
}

/*
FindMobileDeviceEnrollmentProfilesByIDSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceEnrollmentProfilesByIDSubsetOK struct {
}

// IsSuccess returns true when this find mobile device enrollment profiles by Id subset o k response has a 2xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device enrollment profiles by Id subset o k response has a 3xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device enrollment profiles by Id subset o k response has a 4xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device enrollment profiles by Id subset o k response has a 5xx status code
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device enrollment profiles by Id subset o k response a status code equal to that given
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device enrollment profiles by Id subset o k response
func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) Code() int {
	return 200
}

func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}][%d] findMobileDeviceEnrollmentProfilesByIdSubsetOK ", 200)
}

func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}][%d] findMobileDeviceEnrollmentProfilesByIdSubsetOK ", 200)
}

func (o *FindMobileDeviceEnrollmentProfilesByIDSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}