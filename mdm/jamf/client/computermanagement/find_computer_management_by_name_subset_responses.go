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

// FindComputerManagementByNameSubsetReader is a Reader for the FindComputerManagementByNameSubset structure.
type FindComputerManagementByNameSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerManagementByNameSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerManagementByNameSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computermanagement/name/{name}/subset/{subset}] findComputerManagementByNameSubset", response, response.Code())
	}
}

// NewFindComputerManagementByNameSubsetOK creates a FindComputerManagementByNameSubsetOK with default headers values
func NewFindComputerManagementByNameSubsetOK() *FindComputerManagementByNameSubsetOK {
	return &FindComputerManagementByNameSubsetOK{}
}

/*
FindComputerManagementByNameSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerManagementByNameSubsetOK struct {
	Payload *models.ComputerManagement
}

// IsSuccess returns true when this find computer management by name subset o k response has a 2xx status code
func (o *FindComputerManagementByNameSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer management by name subset o k response has a 3xx status code
func (o *FindComputerManagementByNameSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer management by name subset o k response has a 4xx status code
func (o *FindComputerManagementByNameSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer management by name subset o k response has a 5xx status code
func (o *FindComputerManagementByNameSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer management by name subset o k response a status code equal to that given
func (o *FindComputerManagementByNameSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer management by name subset o k response
func (o *FindComputerManagementByNameSubsetOK) Code() int {
	return 200
}

func (o *FindComputerManagementByNameSubsetOK) Error() string {
	return fmt.Sprintf("[GET /computermanagement/name/{name}/subset/{subset}][%d] findComputerManagementByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByNameSubsetOK) String() string {
	return fmt.Sprintf("[GET /computermanagement/name/{name}/subset/{subset}][%d] findComputerManagementByNameSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerManagementByNameSubsetOK) GetPayload() *models.ComputerManagement {
	return o.Payload
}

func (o *FindComputerManagementByNameSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerManagement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
