// Code generated by go-swagger; DO NOT EDIT.

package removablemacaddresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRemovableMacAddressByIDReader is a Reader for the DeleteRemovableMacAddressByID structure.
type DeleteRemovableMacAddressByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRemovableMacAddressByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRemovableMacAddressByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /removablemacaddresses/id/{id}] deleteRemovableMacAddressById", response, response.Code())
	}
}

// NewDeleteRemovableMacAddressByIDOK creates a DeleteRemovableMacAddressByIDOK with default headers values
func NewDeleteRemovableMacAddressByIDOK() *DeleteRemovableMacAddressByIDOK {
	return &DeleteRemovableMacAddressByIDOK{}
}

/*
DeleteRemovableMacAddressByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRemovableMacAddressByIDOK struct {
}

// IsSuccess returns true when this delete removable mac address by Id o k response has a 2xx status code
func (o *DeleteRemovableMacAddressByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete removable mac address by Id o k response has a 3xx status code
func (o *DeleteRemovableMacAddressByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete removable mac address by Id o k response has a 4xx status code
func (o *DeleteRemovableMacAddressByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete removable mac address by Id o k response has a 5xx status code
func (o *DeleteRemovableMacAddressByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete removable mac address by Id o k response a status code equal to that given
func (o *DeleteRemovableMacAddressByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete removable mac address by Id o k response
func (o *DeleteRemovableMacAddressByIDOK) Code() int {
	return 200
}

func (o *DeleteRemovableMacAddressByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /removablemacaddresses/id/{id}][%d] deleteRemovableMacAddressByIdOK ", 200)
}

func (o *DeleteRemovableMacAddressByIDOK) String() string {
	return fmt.Sprintf("[DELETE /removablemacaddresses/id/{id}][%d] deleteRemovableMacAddressByIdOK ", 200)
}

func (o *DeleteRemovableMacAddressByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
