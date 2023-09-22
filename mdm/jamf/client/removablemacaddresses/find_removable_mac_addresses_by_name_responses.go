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

// FindRemovableMacAddressesByNameReader is a Reader for the FindRemovableMacAddressesByName structure.
type FindRemovableMacAddressesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindRemovableMacAddressesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindRemovableMacAddressesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /removablemacaddresses/name/{name}] findRemovableMacAddressesByName", response, response.Code())
	}
}

// NewFindRemovableMacAddressesByNameOK creates a FindRemovableMacAddressesByNameOK with default headers values
func NewFindRemovableMacAddressesByNameOK() *FindRemovableMacAddressesByNameOK {
	return &FindRemovableMacAddressesByNameOK{}
}

/*
FindRemovableMacAddressesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindRemovableMacAddressesByNameOK struct {
	Payload *models.RemovableMacAddress
}

// IsSuccess returns true when this find removable mac addresses by name o k response has a 2xx status code
func (o *FindRemovableMacAddressesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find removable mac addresses by name o k response has a 3xx status code
func (o *FindRemovableMacAddressesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find removable mac addresses by name o k response has a 4xx status code
func (o *FindRemovableMacAddressesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find removable mac addresses by name o k response has a 5xx status code
func (o *FindRemovableMacAddressesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find removable mac addresses by name o k response a status code equal to that given
func (o *FindRemovableMacAddressesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find removable mac addresses by name o k response
func (o *FindRemovableMacAddressesByNameOK) Code() int {
	return 200
}

func (o *FindRemovableMacAddressesByNameOK) Error() string {
	return fmt.Sprintf("[GET /removablemacaddresses/name/{name}][%d] findRemovableMacAddressesByNameOK  %+v", 200, o.Payload)
}

func (o *FindRemovableMacAddressesByNameOK) String() string {
	return fmt.Sprintf("[GET /removablemacaddresses/name/{name}][%d] findRemovableMacAddressesByNameOK  %+v", 200, o.Payload)
}

func (o *FindRemovableMacAddressesByNameOK) GetPayload() *models.RemovableMacAddress {
	return o.Payload
}

func (o *FindRemovableMacAddressesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemovableMacAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}