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

// FindComputerManagementBySerialNumberUsernameSubsetReader is a Reader for the FindComputerManagementBySerialNumberUsernameSubset structure.
type FindComputerManagementBySerialNumberUsernameSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerManagementBySerialNumberUsernameSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerManagementBySerialNumberUsernameSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}/subset/{subset}] findComputerManagementBySerialNumberUsernameSubset", response, response.Code())
	}
}

// NewFindComputerManagementBySerialNumberUsernameSubsetOK creates a FindComputerManagementBySerialNumberUsernameSubsetOK with default headers values
func NewFindComputerManagementBySerialNumberUsernameSubsetOK() *FindComputerManagementBySerialNumberUsernameSubsetOK {
	return &FindComputerManagementBySerialNumberUsernameSubsetOK{}
}

/*
FindComputerManagementBySerialNumberUsernameSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerManagementBySerialNumberUsernameSubsetOK struct {
	Payload *models.ComputerManagement
}

// IsSuccess returns true when this find computer management by serial number username subset o k response has a 2xx status code
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer management by serial number username subset o k response has a 3xx status code
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer management by serial number username subset o k response has a 4xx status code
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer management by serial number username subset o k response has a 5xx status code
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer management by serial number username subset o k response a status code equal to that given
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer management by serial number username subset o k response
func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) Code() int {
	return 200
}

func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) Error() string {
	return fmt.Sprintf("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}/subset/{subset}][%d] findComputerManagementBySerialNumberUsernameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) String() string {
	return fmt.Sprintf("[GET /computermanagement/serialnumber/{serialnumber}/username/{username}/subset/{subset}][%d] findComputerManagementBySerialNumberUsernameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) GetPayload() *models.ComputerManagement {
	return o.Payload
}

func (o *FindComputerManagementBySerialNumberUsernameSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerManagement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
