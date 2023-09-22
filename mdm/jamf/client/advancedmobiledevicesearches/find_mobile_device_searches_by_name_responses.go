// Code generated by go-swagger; DO NOT EDIT.

package advancedmobiledevicesearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindMobileDeviceSearchesByNameReader is a Reader for the FindMobileDeviceSearchesByName structure.
type FindMobileDeviceSearchesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMobileDeviceSearchesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindMobileDeviceSearchesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /advancedmobiledevicesearches/name/{name}] findMobileDeviceSearchesByName", response, response.Code())
	}
}

// NewFindMobileDeviceSearchesByNameOK creates a FindMobileDeviceSearchesByNameOK with default headers values
func NewFindMobileDeviceSearchesByNameOK() *FindMobileDeviceSearchesByNameOK {
	return &FindMobileDeviceSearchesByNameOK{}
}

/*
FindMobileDeviceSearchesByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindMobileDeviceSearchesByNameOK struct {
	Payload *models.AdvancedMobileDeviceSearch
}

// IsSuccess returns true when this find mobile device searches by name o k response has a 2xx status code
func (o *FindMobileDeviceSearchesByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find mobile device searches by name o k response has a 3xx status code
func (o *FindMobileDeviceSearchesByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find mobile device searches by name o k response has a 4xx status code
func (o *FindMobileDeviceSearchesByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find mobile device searches by name o k response has a 5xx status code
func (o *FindMobileDeviceSearchesByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find mobile device searches by name o k response a status code equal to that given
func (o *FindMobileDeviceSearchesByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find mobile device searches by name o k response
func (o *FindMobileDeviceSearchesByNameOK) Code() int {
	return 200
}

func (o *FindMobileDeviceSearchesByNameOK) Error() string {
	return fmt.Sprintf("[GET /advancedmobiledevicesearches/name/{name}][%d] findMobileDeviceSearchesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceSearchesByNameOK) String() string {
	return fmt.Sprintf("[GET /advancedmobiledevicesearches/name/{name}][%d] findMobileDeviceSearchesByNameOK  %+v", 200, o.Payload)
}

func (o *FindMobileDeviceSearchesByNameOK) GetPayload() *models.AdvancedMobileDeviceSearch {
	return o.Payload
}

func (o *FindMobileDeviceSearchesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdvancedMobileDeviceSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
