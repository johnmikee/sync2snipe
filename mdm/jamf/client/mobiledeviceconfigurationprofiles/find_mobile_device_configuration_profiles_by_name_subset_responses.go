// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceconfigurationprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceConfigurationProfilesByNameSubsetReader is a Reader for the FindMobileDeviceConfigurationProfilesByNameSubset structure.
type FindMobileDeviceConfigurationProfilesByNameSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceConfigurationProfilesByNameSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}] findMobileDeviceConfigurationProfilesByNameSubset", response, response.Code())
	}
}

// NewFindMobileDeviceConfigurationProfilesByNameSubsetOK creates a FindMobileDeviceConfigurationProfilesByNameSubsetOK with default headers values
func NewFindMobileDeviceConfigurationProfilesByNameSubsetOK() *FindMobileDeviceConfigurationProfilesByNameSubsetOK {
	return &FindMobileDeviceConfigurationProfilesByNameSubsetOK{}
}

/*
FindMobileDeviceConfigurationProfilesByNameSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceConfigurationProfilesByNameSubsetOK struct {
	Payload *models.MobileDeviceConfigurationProfile
}

// IsSuccess returns true when this find mobile device configuration profiles by name subset o k response has a 2xx status code
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device configuration profiles by name subset o k response has a 3xx status code
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device configuration profiles by name subset o k response has a 4xx status code
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device configuration profiles by name subset o k response has a 5xx status code
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device configuration profiles by name subset o k response a status code equal to that given
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device configuration profiles by name subset o k response
func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) Code() int {
	return 200
}

func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) Error() string {
	return fmt.Sprintf("[GET /mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}][%d] findMobileDeviceConfigurationProfilesByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) String() string {
	return fmt.Sprintf("[GET /mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}][%d] findMobileDeviceConfigurationProfilesByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) GetPayload() *models.MobileDeviceConfigurationProfile {
	return o.Payload
}

func (o *FindMobileDeviceConfigurationProfilesByNameSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceConfigurationProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
