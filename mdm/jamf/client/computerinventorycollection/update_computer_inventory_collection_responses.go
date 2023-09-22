// Code generated by go-swagger; DO NOT EDIT.

package computerinventorycollection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateComputerInventoryCollectionReader is a Reader for the UpdateComputerInventoryCollection structure.
type UpdateComputerInventoryCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComputerInventoryCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateComputerInventoryCollectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /computerinventorycollection] updateComputerInventoryCollection", response, response.Code())
	}
}

// NewUpdateComputerInventoryCollectionCreated creates a UpdateComputerInventoryCollectionCreated with default headers values
func NewUpdateComputerInventoryCollectionCreated() *UpdateComputerInventoryCollectionCreated {
	return &UpdateComputerInventoryCollectionCreated{}
}

/*
UpdateComputerInventoryCollectionCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateComputerInventoryCollectionCreated struct {
}

// IsSuccess returns true when this update computer inventory collection created response has a 2xx status code
func (o *UpdateComputerInventoryCollectionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update computer inventory collection created response has a 3xx status code
func (o *UpdateComputerInventoryCollectionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update computer inventory collection created response has a 4xx status code
func (o *UpdateComputerInventoryCollectionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update computer inventory collection created response has a 5xx status code
func (o *UpdateComputerInventoryCollectionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update computer inventory collection created response a status code equal to that given
func (o *UpdateComputerInventoryCollectionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update computer inventory collection created response
func (o *UpdateComputerInventoryCollectionCreated) Code() int {
	return 201
}

func (o *UpdateComputerInventoryCollectionCreated) Error() string {
	return fmt.Sprintf("[PUT /computerinventorycollection][%d] updateComputerInventoryCollectionCreated ", 201)
}

func (o *UpdateComputerInventoryCollectionCreated) String() string {
	return fmt.Sprintf("[PUT /computerinventorycollection][%d] updateComputerInventoryCollectionCreated ", 201)
}

func (o *UpdateComputerInventoryCollectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
