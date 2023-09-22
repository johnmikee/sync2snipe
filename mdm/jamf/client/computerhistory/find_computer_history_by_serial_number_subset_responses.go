// Code generated by go-swagger; DO NOT EDIT.

package computerhistory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindComputerHistoryBySerialNumberSubsetReader is a Reader for the FindComputerHistoryBySerialNumberSubset structure.
type FindComputerHistoryBySerialNumberSubsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindComputerHistoryBySerialNumberSubsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindComputerHistoryBySerialNumberSubsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /computerhistory/serialnumber/{serialnumber}/subset/{subset}] findComputerHistoryBySerialNumberSubset", response, response.Code())
	}
}

// NewFindComputerHistoryBySerialNumberSubsetOK creates a FindComputerHistoryBySerialNumberSubsetOK with default headers values
func NewFindComputerHistoryBySerialNumberSubsetOK() *FindComputerHistoryBySerialNumberSubsetOK {
	return &FindComputerHistoryBySerialNumberSubsetOK{}
}

/*
FindComputerHistoryBySerialNumberSubsetOK describes a response with status code 200, with default header values.

OK
*/
type FindComputerHistoryBySerialNumberSubsetOK struct {
	Payload *models.ComputerHistory
}

// IsSuccess returns true when this find computer history by serial number subset o k response has a 2xx status code
func (o *FindComputerHistoryBySerialNumberSubsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find computer history by serial number subset o k response has a 3xx status code
func (o *FindComputerHistoryBySerialNumberSubsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find computer history by serial number subset o k response has a 4xx status code
func (o *FindComputerHistoryBySerialNumberSubsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find computer history by serial number subset o k response has a 5xx status code
func (o *FindComputerHistoryBySerialNumberSubsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find computer history by serial number subset o k response a status code equal to that given
func (o *FindComputerHistoryBySerialNumberSubsetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find computer history by serial number subset o k response
func (o *FindComputerHistoryBySerialNumberSubsetOK) Code() int {
	return 200
}

func (o *FindComputerHistoryBySerialNumberSubsetOK) Error() string {
	return fmt.Sprintf("[GET /computerhistory/serialnumber/{serialnumber}/subset/{subset}][%d] findComputerHistoryBySerialNumberSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHistoryBySerialNumberSubsetOK) String() string {
	return fmt.Sprintf("[GET /computerhistory/serialnumber/{serialnumber}/subset/{subset}][%d] findComputerHistoryBySerialNumberSubsetOK  %+v", 200, o.Payload)
}

func (o *FindComputerHistoryBySerialNumberSubsetOK) GetPayload() *models.ComputerHistory {
	return o.Payload
}

func (o *FindComputerHistoryBySerialNumberSubsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
