// Code generated by go-swagger; DO NOT EDIT.

package dockitems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindDockItemsReader is a Reader for the FindDockItems structure.
type FindDockItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDockItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDockItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dockitems] findDockItems", response, response.Code())
	}
}

// NewFindDockItemsOK creates a FindDockItemsOK with default headers values
func NewFindDockItemsOK() *FindDockItemsOK {
	return &FindDockItemsOK{}
}

/*
FindDockItemsOK describes a response with status code 200, with default header values.

OK
*/
type FindDockItemsOK struct {
	Payload models.DockItems
}

// IsSuccess returns true when this find dock items o k response has a 2xx status code
func (o *FindDockItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find dock items o k response has a 3xx status code
func (o *FindDockItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find dock items o k response has a 4xx status code
func (o *FindDockItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find dock items o k response has a 5xx status code
func (o *FindDockItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find dock items o k response a status code equal to that given
func (o *FindDockItemsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find dock items o k response
func (o *FindDockItemsOK) Code() int {
	return 200
}

func (o *FindDockItemsOK) Error() string {
	return fmt.Sprintf("[GET /dockitems][%d] findDockItemsOK  %+v", 200, o.Payload)
}

func (o *FindDockItemsOK) String() string {
	return fmt.Sprintf("[GET /dockitems][%d] findDockItemsOK  %+v", 200, o.Payload)
}

func (o *FindDockItemsOK) GetPayload() models.DockItems {
	return o.Payload
}

func (o *FindDockItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}