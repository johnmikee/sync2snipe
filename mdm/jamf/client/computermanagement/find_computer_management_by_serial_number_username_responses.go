// Code generated by go-swagger; DO NOT EDIT.

package computermanagement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerManagementBySerialNumberUsernameReader is a Reader for the FindComputerManagementBySerialNumberUsername structure.
type FindComputerManagementBySerialNumberUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerManagementBySerialNumberUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerManagementBySerialNumberUsernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}] findComputerManagementBySerialNumberUsername", response, response.Code())
	}
}

// NewFindComputerManagementBySerialNumberUsernameOK creates a FindComputerManagementBySerialNumberUsernameOK with default headers values
func NewFindComputerManagementBySerialNumberUsernameOK() *FindComputerManagementBySerialNumberUsernameOK {
	return &FindComputerManagementBySerialNumberUsernameOK{}
}

/*
FindComputerManagementBySerialNumberUsernameOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerManagementBySerialNumberUsernameOK struct {
	Payload *models.ComputerManagement
}

// IsSuccess returns true when this find computer management by serial number username o k response has a 2xx status code
func (o *FindComputerManagementBySerialNumberUsernameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer management by serial number username o k response has a 3xx status code
func (o *FindComputerManagementBySerialNumberUsernameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer management by serial number username o k response has a 4xx status code
func (o *FindComputerManagementBySerialNumberUsernameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer management by serial number username o k response has a 5xx status code
func (o *FindComputerManagementBySerialNumberUsernameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer management by serial number username o k response a status code equal to that given
func (o *FindComputerManagementBySerialNumberUsernameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer management by serial number username o k response
func (o *FindComputerManagementBySerialNumberUsernameOK) Code() int {
	return 200
}

func (o *FindComputerManagementBySerialNumberUsernameOK) Error() string {
	return fmt.Sprintf("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}][%d] findComputerManagementBySerialNumberUsernameOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementBySerialNumberUsernameOK) String() string {
	return fmt.Sprintf("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}][%d] findComputerManagementBySerialNumberUsernameOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementBySerialNumberUsernameOK) GetPayload() *models.ComputerManagement {
	return o.Payload
}

func (o *FindComputerManagementBySerialNumberUsernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerManagement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
