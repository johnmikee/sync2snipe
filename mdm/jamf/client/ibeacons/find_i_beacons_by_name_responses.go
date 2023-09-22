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

// FindIBeaconsByNameReader is a Reader for the FindIBeaconsByName structure.
type FindIBeaconsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindIBeaconsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindIBeaconsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ibeacons/name/{name}] findIBeaconsByName", response, response.Code())
	}
}

// NewFindIBeaconsByNameOK creates a FindIBeaconsByNameOK with default headers values
func NewFindIBeaconsByNameOK() *FindIBeaconsByNameOK {
	return &FindIBeaconsByNameOK{}
}

/*
FindIBeaconsByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindIBeaconsByNameOK struct {
	Payload *models.Ibeacon
}

// IsSuccess returns true when this find i beacons by name o k response has a 2xx status code
func (o *FindIBeaconsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find i beacons by name o k response has a 3xx status code
func (o *FindIBeaconsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find i beacons by name o k response has a 4xx status code
func (o *FindIBeaconsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find i beacons by name o k response has a 5xx status code
func (o *FindIBeaconsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find i beacons by name o k response a status code equal to that given
func (o *FindIBeaconsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find i beacons by name o k response
func (o *FindIBeaconsByNameOK) Code() int {
	return 200
}

func (o *FindIBeaconsByNameOK) Error() string {
	return fmt.Sprintf("[GET /ibeacons/name/{name}][%d] findIBeaconsByNameOK  %+v", 200, o.Payload)
}

func (o *FindIBeaconsByNameOK) String() string {
	return fmt.Sprintf("[GET /ibeacons/name/{name}][%d] findIBeaconsByNameOK  %+v", 200, o.Payload)
}

func (o *FindIBeaconsByNameOK) GetPayload() *models.Ibeacon {
	return o.Payload
}

func (o *FindIBeaconsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ibeacon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
