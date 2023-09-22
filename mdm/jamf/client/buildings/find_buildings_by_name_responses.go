// Code generated by go-swagger; DO NOT EDIT.

package buildings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindBuildingsByNameReader is a Reader for the FindBuildingsByName structure.
type FindBuildingsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindBuildingsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindBuildingsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /buildings/name/{name}] findBuildingsByName", response, response.Code())
	}
}

// NewFindBuildingsByNameOK creates a FindBuildingsByNameOK with default headers values
func NewFindBuildingsByNameOK() *FindBuildingsByNameOK {
	return &FindBuildingsByNameOK{}
}

/*
FindBuildingsByNameOK describes a response with status code 200, with default header values.

OK
*/
type FindBuildingsByNameOK struct {
	Payload *models.Building
}

// IsSuccess returns true when this find buildings by name o k response has a 2xx status code
func (o *FindBuildingsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find buildings by name o k response has a 3xx status code
func (o *FindBuildingsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find buildings by name o k response has a 4xx status code
func (o *FindBuildingsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find buildings by name o k response has a 5xx status code
func (o *FindBuildingsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find buildings by name o k response a status code equal to that given
func (o *FindBuildingsByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find buildings by name o k response
func (o *FindBuildingsByNameOK) Code() int {
	return 200
}

func (o *FindBuildingsByNameOK) Error() string {
	return fmt.Sprintf("[GET /buildings/name/{name}][%d] findBuildingsByNameOK  %+v", 200, o.Payload)
}

func (o *FindBuildingsByNameOK) String() string {
	return fmt.Sprintf("[GET /buildings/name/{name}][%d] findBuildingsByNameOK  %+v", 200, o.Payload)
}

func (o *FindBuildingsByNameOK) GetPayload() *models.Building {
	return o.Payload
}

func (o *FindBuildingsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Building)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}