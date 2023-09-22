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

// FindComputerManagementByMacAddressPatchFilterReader is a Reader for the FindComputerManagementByMacAddressPatchFilter structure.
type FindComputerManagementByMacAddressPatchFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerManagementByMacAddressPatchFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerManagementByMacAddressPatchFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computermanagement/macaddress/{macaddress}/patchfilter/{filter}] findComputerManagementByMacAddressPatchFilter", response, response.Code())
	}
}

// NewFindComputerManagementByMacAddressPatchFilterOK creates a FindComputerManagementByMacAddressPatchFilterOK with default headers values
func NewFindComputerManagementByMacAddressPatchFilterOK() *FindComputerManagementByMacAddressPatchFilterOK {
	return &FindComputerManagementByMacAddressPatchFilterOK{}
}

/*
FindComputerManagementByMacAddressPatchFilterOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerManagementByMacAddressPatchFilterOK struct {
	Payload *models.ComputerManagement
}

// IsSuccess returns true when this find computer management by mac address patch filter o k response has a 2xx status code
func (o *FindComputerManagementByMacAddressPatchFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer management by mac address patch filter o k response has a 3xx status code
func (o *FindComputerManagementByMacAddressPatchFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer management by mac address patch filter o k response has a 4xx status code
func (o *FindComputerManagementByMacAddressPatchFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer management by mac address patch filter o k response has a 5xx status code
func (o *FindComputerManagementByMacAddressPatchFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer management by mac address patch filter o k response a status code equal to that given
func (o *FindComputerManagementByMacAddressPatchFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer management by mac address patch filter o k response
func (o *FindComputerManagementByMacAddressPatchFilterOK) Code() int {
	return 200
}

func (o *FindComputerManagementByMacAddressPatchFilterOK) Error() string {
	return fmt.Sprintf("[GET /computermanagement/macaddress/{macaddress}/patchfilter/{filter}][%d] findComputerManagementByMacAddressPatchFilterOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByMacAddressPatchFilterOK) String() string {
	return fmt.Sprintf("[GET /computermanagement/macaddress/{macaddress}/patchfilter/{filter}][%d] findComputerManagementByMacAddressPatchFilterOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByMacAddressPatchFilterOK) GetPayload() *models.ComputerManagement {
	return o.Payload
}

func (o *FindComputerManagementByMacAddressPatchFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerManagement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
