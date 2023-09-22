// Code generated by go-swagger; DO NOT EDIT.

package healthcarelistenerrule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindHealthcareListenerRuleReader is a Reader for the FindHealthcareListenerRule structure.
type FindHealthcareListenerRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindHealthcareListenerRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindHealthcareListenerRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /healthcarelistenerrule] findHealthcareListenerRule", response, response.Code())
	}
}

// NewFindHealthcareListenerRuleOK creates a FindHealthcareListenerRuleOK with default headers values
func NewFindHealthcareListenerRuleOK() *FindHealthcareListenerRuleOK {
	return &FindHealthcareListenerRuleOK{}
}

/*
FindHealthcareListenerRuleOK describes a response with status code 200, with default header values.

OK
*/
type FindHealthcareListenerRuleOK struct {
	Payload models.HealthcareListenerRules
}

// IsSuccess returns true when this find healthcare listener rule o k response has a 2xx status code
func (o *FindHealthcareListenerRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find healthcare listener rule o k response has a 3xx status code
func (o *FindHealthcareListenerRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find healthcare listener rule o k response has a 4xx status code
func (o *FindHealthcareListenerRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find healthcare listener rule o k response has a 5xx status code
func (o *FindHealthcareListenerRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find healthcare listener rule o k response a status code equal to that given
func (o *FindHealthcareListenerRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find healthcare listener rule o k response
func (o *FindHealthcareListenerRuleOK) Code() int {
	return 200
}

func (o *FindHealthcareListenerRuleOK) Error() string {
	return fmt.Sprintf("[GET /healthcarelistenerrule][%d] findHealthcareListenerRuleOK  %+v", 200, o.Payload)
}

func (o *FindHealthcareListenerRuleOK) String() string {
	return fmt.Sprintf("[GET /healthcarelistenerrule][%d] findHealthcareListenerRuleOK  %+v", 200, o.Payload)
}

func (o *FindHealthcareListenerRuleOK) GetPayload() models.HealthcareListenerRules {
	return o.Payload
}

func (o *FindHealthcareListenerRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
