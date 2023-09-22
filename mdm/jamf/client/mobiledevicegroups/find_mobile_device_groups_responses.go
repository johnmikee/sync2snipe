// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicegroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceGroupsReader is a Reader for the FindMobileDeviceGroups structure.
type FindMobileDeviceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledevicegroups] findMobileDeviceGroups", response, response.Code())
	}
}

// NewFindMobileDeviceGroupsOK creates a FindMobileDeviceGroupsOK with default headers values
func NewFindMobileDeviceGroupsOK() *FindMobileDeviceGroupsOK {
	return &FindMobileDeviceGroupsOK{}
}

/*
FindMobileDeviceGroupsOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceGroupsOK struct {
	Payload models.MobileDeviceGroups
}

// IsSuccess returns true when this find mobile device groups o k response has a 2xx status code
func (o *FindMobileDeviceGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device groups o k response has a 3xx status code
func (o *FindMobileDeviceGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device groups o k response has a 4xx status code
func (o *FindMobileDeviceGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device groups o k response has a 5xx status code
func (o *FindMobileDeviceGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device groups o k response a status code equal to that given
func (o *FindMobileDeviceGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device groups o k response
func (o *FindMobileDeviceGroupsOK) Code() int {
	return 200
}

func (o *FindMobileDeviceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /mobiledevicegroups][%d] findMobileDeviceGroupsOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceGroupsOK) String() string {
	return fmt.Sprintf("[GET /mobiledevicegroups][%d] findMobileDeviceGroupsOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceGroupsOK) GetPayload() models.MobileDeviceGroups {
	return o.Payload
}

func (o *FindMobileDeviceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
