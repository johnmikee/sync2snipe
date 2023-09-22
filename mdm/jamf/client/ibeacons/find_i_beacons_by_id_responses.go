// Code generated by go-swagger; DO NOT EDIT.

package ibeacons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindIBeaconsByIDReader is a Reader for the FindIBeaconsByID structure.
type FindIBeaconsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindIBeaconsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindIBeaconsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ibeacons/id/{id}] findIBeaconsById", response, response.Code())
	}
}

// NewFindIBeaconsByIDOK creates a FindIBeaconsByIDOK with default headers values
func NewFindIBeaconsByIDOK() *FindIBeaconsByIDOK {
	return &FindIBeaconsByIDOK{}
}

/*
FindIBeaconsByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindIBeaconsByIDOK struct {
	Payload *models.Ibeacon
}

// IsSuccess returns true when this find i beacons by Id o k response has a 2xx status code
func (o *FindIBeaconsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find i beacons by Id o k response has a 3xx status code
func (o *FindIBeaconsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find i beacons by Id o k response has a 4xx status code
func (o *FindIBeaconsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find i beacons by Id o k response has a 5xx status code
func (o *FindIBeaconsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find i beacons by Id o k response a status code equal to that given
func (o *FindIBeaconsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find i beacons by Id o k response
func (o *FindIBeaconsByIDOK) Code() int {
	return 200
}

func (o *FindIBeaconsByIDOK) Error() string {
	return fmt.Sprintf("[GET /ibeacons/id/{id}][%d] findIBeaconsByIdOK  %+v", 200, o.Payload)
}

func (o *FindIBeaconsByIDOK) String() string {
	return fmt.Sprintf("[GET /ibeacons/id/{id}][%d] findIBeaconsByIdOK  %+v", 200, o.Payload)
}

func (o *FindIBeaconsByIDOK) GetPayload() *models.Ibeacon {
	return o.Payload
}

func (o *FindIBeaconsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ibeacon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}