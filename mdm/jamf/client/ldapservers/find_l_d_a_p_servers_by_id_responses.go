// Code generated by go-swagger; DO NOT EDIT.

package ldapservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindLDAPServersByIDReader is a Reader for the FindLDAPServersByID structure.
type FindLDAPServersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindLDAPServersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindLDAPServersByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ldapservers/id/{id}] findLDAPServersById", response, response.Code())
	}
}

// NewFindLDAPServersByIDOK creates a FindLDAPServersByIDOK with default headers values
func NewFindLDAPServersByIDOK() *FindLDAPServersByIDOK {
	return &FindLDAPServersByIDOK{}
}

/*
FindLDAPServersByIDOK describes a response with status code 200, with default header values.

OK
*/
type FindLDAPServersByIDOK struct {
	Payload *models.LdapServer
}

// IsSuccess returns true when this find l d a p servers by Id o k response has a 2xx status code
func (o *FindLDAPServersByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find l d a p servers by Id o k response has a 3xx status code
func (o *FindLDAPServersByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find l d a p servers by Id o k response has a 4xx status code
func (o *FindLDAPServersByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find l d a p servers by Id o k response has a 5xx status code
func (o *FindLDAPServersByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find l d a p servers by Id o k response a status code equal to that given
func (o *FindLDAPServersByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find l d a p servers by Id o k response
func (o *FindLDAPServersByIDOK) Code() int {
	return 200
}

func (o *FindLDAPServersByIDOK) Error() string {
	return fmt.Sprintf("[GET /ldapservers/id/{id}][%d] findLDAPServersByIdOK  %+v", 200, o.Payload)
}

func (o *FindLDAPServersByIDOK) String() string {
	return fmt.Sprintf("[GET /ldapservers/id/{id}][%d] findLDAPServersByIdOK  %+v", 200, o.Payload)
}

func (o *FindLDAPServersByIDOK) GetPayload() *models.LdapServer {
	return o.Payload
}

func (o *FindLDAPServersByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LdapServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}