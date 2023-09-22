// Code generated by go-swagger; DO NOT EDIT.

package advancedusersearches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAdvancedUserSearchByNameReader is a Reader for the DeleteAdvancedUserSearchByName structure.
type DeleteAdvancedUserSearchByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdvancedUserSearchByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdvancedUserSearchByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /advancedusersearches/name/{name}] deleteAdvancedUserSearchByName", response, response.Code())
	}
}

// NewDeleteAdvancedUserSearchByNameOK creates a DeleteAdvancedUserSearchByNameOK with default headers values
func NewDeleteAdvancedUserSearchByNameOK() *DeleteAdvancedUserSearchByNameOK {
	return &DeleteAdvancedUserSearchByNameOK{}
}

/*
DeleteAdvancedUserSearchByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAdvancedUserSearchByNameOK struct {
}

// IsSuccess returns true when this delete advanced user search by name o k response has a 2xx status code
func (o *DeleteAdvancedUserSearchByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete advanced user search by name o k response has a 3xx status code
func (o *DeleteAdvancedUserSearchByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete advanced user search by name o k response has a 4xx status code
func (o *DeleteAdvancedUserSearchByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete advanced user search by name o k response has a 5xx status code
func (o *DeleteAdvancedUserSearchByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete advanced user search by name o k response a status code equal to that given
func (o *DeleteAdvancedUserSearchByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete advanced user search by name o k response
func (o *DeleteAdvancedUserSearchByNameOK) Code() int {
	return 200
}

func (o *DeleteAdvancedUserSearchByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /advancedusersearches/name/{name}][%d] deleteAdvancedUserSearchByNameOK ", 200)
}

func (o *DeleteAdvancedUserSearchByNameOK) String() string {
	return fmt.Sprintf("[DELETE /advancedusersearches/name/{name}][%d] deleteAdvancedUserSearchByNameOK ", 200)
}

func (o *DeleteAdvancedUserSearchByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
