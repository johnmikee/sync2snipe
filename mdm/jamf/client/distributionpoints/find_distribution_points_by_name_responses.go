// Code generated by go-swagger; DO NOT EDIT.

package distributionpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindDistributionPointsByNameReader is a Reader for the FindDistributionPointsByName structure.
type FindDistributionPointsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDistributionPointsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDistributionPointsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /distributionpoints/name/{name}] findDistributionPointsByName", response, response.Code())
	}
}

// NewFindDistributionPointsByNameOK creates a FindDistributionPointsByNameOK with default headers values
func NewFindDistributionPointsByNameOK() *FindDistributionPointsByNameOK {
	return &FindDistributionPointsByNameOK{}
}

/*
FindDistributionPointsByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindDistributionPointsByNameOK struct {
	Payload *models.DistributionPoint
}

// IsSuccess returns true when this find distribution points by name o k response has a 2xx status code
func (o *FindDistributionPointsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find distribution points by name o k response has a 3xx status code
func (o *FindDistributionPointsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find distribution points by name o k response has a 4xx status code
func (o *FindDistributionPointsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find distribution points by name o k response has a 5xx status code
func (o *FindDistributionPointsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find distribution points by name o k response a status code equal to that given
func (o *FindDistributionPointsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find distribution points by name o k response
func (o *FindDistributionPointsByNameOK) Code() int {
	return 200
}

func (o *FindDistributionPointsByNameOK) Error() string {
	return fmt.Sprintf("[GET /distributionpoints/name/{name}][%d] findDistributionPointsByNameOK  %+v", 200, o.Payload)
}

func (o *FindDistributionPointsByNameOK) String() string {
	return fmt.Sprintf("[GET /distributionpoints/name/{name}][%d] findDistributionPointsByNameOK  %+v", 200, o.Payload)
}

func (o *FindDistributionPointsByNameOK) GetPayload() *models.DistributionPoint {
	return o.Payload
}

func (o *FindDistributionPointsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistributionPoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}