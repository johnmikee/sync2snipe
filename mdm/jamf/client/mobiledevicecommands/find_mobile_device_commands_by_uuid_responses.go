// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicecommands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceCommandsByUUIDReader is a Reader for the FindMobileDeviceCommandsByUUID structure.
type FindMobileDeviceCommandsByUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceCommandsByUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceCommandsByUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /mobiledevicecommands/uuid/{uuid}] findMobileDeviceCommandsByUuid", response, response.Code())
	}
}

// NewFindMobileDeviceCommandsByUUIDOK creates a FindMobileDeviceCommandsByUUIDOK with default headers values
func NewFindMobileDeviceCommandsByUUIDOK() *FindMobileDeviceCommandsByUUIDOK {
	return &FindMobileDeviceCommandsByUUIDOK{}
}

/*
FindMobileDeviceCommandsByUUIDOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceCommandsByUUIDOK struct {
	Payload *models.MobileDeviceCommand
}

// IsSuccess returns true when this find mobile device commands by Uuid o k response has a 2xx status code
func (o *FindMobileDeviceCommandsByUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device commands by Uuid o k response has a 3xx status code
func (o *FindMobileDeviceCommandsByUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device commands by Uuid o k response has a 4xx status code
func (o *FindMobileDeviceCommandsByUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device commands by Uuid o k response has a 5xx status code
func (o *FindMobileDeviceCommandsByUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device commands by Uuid o k response a status code equal to that given
func (o *FindMobileDeviceCommandsByUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device commands by Uuid o k response
func (o *FindMobileDeviceCommandsByUUIDOK) Code() int {
	return 200
}

func (o *FindMobileDeviceCommandsByUUIDOK) Error() string {
	return fmt.Sprintf("[GET /mobiledevicecommands/uuid/{uuid}][%d] findMobileDeviceCommandsByUuidOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceCommandsByUUIDOK) String() string {
	return fmt.Sprintf("[GET /mobiledevicecommands/uuid/{uuid}][%d] findMobileDeviceCommandsByUuidOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceCommandsByUUIDOK) GetPayload() *models.MobileDeviceCommand {
	return o.Payload
}

func (o *FindMobileDeviceCommandsByUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MobileDeviceCommand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
