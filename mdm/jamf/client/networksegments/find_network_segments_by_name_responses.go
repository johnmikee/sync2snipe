// Code generated by go-swagger; DO NOT EDIT.

package networksegments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindNetworkSegmentsByNameReader is a Reader for the FindNetworkSegmentsByName structure.
type FindNetworkSegmentsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindNetworkSegmentsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindNetworkSegmentsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networksegments/name/{name}] findNetworkSegmentsByName", response, response.Code())
	}
}

// NewFindNetworkSegmentsByNameOK creates a FindNetworkSegmentsByNameOK with default headers values
func NewFindNetworkSegmentsByNameOK() *FindNetworkSegmentsByNameOK {
	return &FindNetworkSegmentsByNameOK{}
}

/*
FindNetworkSegmentsByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindNetworkSegmentsByNameOK struct {
	Payload *models.NetworkSegment
}

// IsSuccess returns true when this find network segments by name o k response has a 2xx status code
func (o *FindNetworkSegmentsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find network segments by name o k response has a 3xx status code
func (o *FindNetworkSegmentsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find network segments by name o k response has a 4xx status code
func (o *FindNetworkSegmentsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find network segments by name o k response has a 5xx status code
func (o *FindNetworkSegmentsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find network segments by name o k response a status code equal to that given
func (o *FindNetworkSegmentsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find network segments by name o k response
func (o *FindNetworkSegmentsByNameOK) Code() int {
	return 200
}

func (o *FindNetworkSegmentsByNameOK) Error() string {
	return fmt.Sprintf("[GET /networksegments/name/{name}][%d] findNetworkSegmentsByNameOK  %+v", 200, o.Payload)
}

func (o *FindNetworkSegmentsByNameOK) String() string {
	return fmt.Sprintf("[GET /networksegments/name/{name}][%d] findNetworkSegmentsByNameOK  %+v", 200, o.Payload)
}

func (o *FindNetworkSegmentsByNameOK) GetPayload() *models.NetworkSegment {
	return o.Payload
}

func (o *FindNetworkSegmentsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkSegment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}