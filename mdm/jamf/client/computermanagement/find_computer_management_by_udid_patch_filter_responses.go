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

// FindComputerManagementByUdidPatchFilterReader is a Reader for the FindComputerManagementByUdidPatchFilter structure.
type FindComputerManagementByUdidPatchFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerManagementByUdidPatchFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerManagementByUdidPatchFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computermanagement/udid/{udid}/patchfilter/{filter}] findComputerManagementByUdidPatchFilter", response, response.Code())
	}
}

// NewFindComputerManagementByUdidPatchFilterOK creates a FindComputerManagementByUdidPatchFilterOK with default headers values
func NewFindComputerManagementByUdidPatchFilterOK() *FindComputerManagementByUdidPatchFilterOK {
	return &FindComputerManagementByUdidPatchFilterOK{}
}

/*
FindComputerManagementByUdidPatchFilterOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerManagementByUdidPatchFilterOK struct {
	Payload *models.ComputerManagement
}

// IsSuccess returns true when this find computer management by udid patch filter o k response has a 2xx status code
func (o *FindComputerManagementByUdidPatchFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer management by udid patch filter o k response has a 3xx status code
func (o *FindComputerManagementByUdidPatchFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer management by udid patch filter o k response has a 4xx status code
func (o *FindComputerManagementByUdidPatchFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer management by udid patch filter o k response has a 5xx status code
func (o *FindComputerManagementByUdidPatchFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer management by udid patch filter o k response a status code equal to that given
func (o *FindComputerManagementByUdidPatchFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer management by udid patch filter o k response
func (o *FindComputerManagementByUdidPatchFilterOK) Code() int {
	return 200
}

func (o *FindComputerManagementByUdidPatchFilterOK) Error() string {
	return fmt.Sprintf("[GET /computermanagement/udid/{udid}/patchfilter/{filter}][%d] findComputerManagementByUdidPatchFilterOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByUdidPatchFilterOK) String() string {
	return fmt.Sprintf("[GET /computermanagement/udid/{udid}/patchfilter/{filter}][%d] findComputerManagementByUdidPatchFilterOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByUdidPatchFilterOK) GetPayload() *models.ComputerManagement {
	return o.Payload
}

func (o *FindComputerManagementByUdidPatchFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerManagement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}