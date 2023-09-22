// Code generated by go-swagger; DO NOT EDIT.

package removablemacaddresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindRemovableMacAddressesReader is a Reader for the FindRemovableMacAddresses structure.
type FindRemovableMacAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindRemovableMacAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindRemovableMacAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /removablemacaddresses] findRemovableMacAddresses", response, response.Code())
	}
}

// NewFindRemovableMacAddressesOK creates a FindRemovableMacAddressesOK with default headers values
func NewFindRemovableMacAddressesOK() *FindRemovableMacAddressesOK {
	return &FindRemovableMacAddressesOK{}
}

/*
FindRemovableMacAddressesOK describes a response with status code 200, with default header values.

OK
*/
type FindRemovableMacAddressesOK struct {
	Payload models.RemovableMacAddresses
}

// IsSuccess returns true when this find removable mac addresses o k response has a 2xx status code
func (o *FindRemovableMacAddressesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find removable mac addresses o k response has a 3xx status code
func (o *FindRemovableMacAddressesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find removable mac addresses o k response has a 4xx status code
func (o *FindRemovableMacAddressesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find removable mac addresses o k response has a 5xx status code
func (o *FindRemovableMacAddressesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find removable mac addresses o k response a status code equal to that given
func (o *FindRemovableMacAddressesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find removable mac addresses o k response
func (o *FindRemovableMacAddressesOK) Code() int {
	return 200
}

func (o *FindRemovableMacAddressesOK) Error() string {
	return fmt.Sprintf("[GET /removablemacaddresses][%d] findRemovableMacAddressesOK  %+v", 200, o.Payload)
}

func (o *FindRemovableMacAddressesOK) String() string {
	return fmt.Sprintf("[GET /removablemacaddresses][%d] findRemovableMacAddressesOK  %+v", 200, o.Payload)
}

func (o *FindRemovableMacAddressesOK) GetPayload() models.RemovableMacAddresses {
	return o.Payload
}

func (o *FindRemovableMacAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}