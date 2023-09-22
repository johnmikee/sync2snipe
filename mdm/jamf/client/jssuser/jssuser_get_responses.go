// Code generated by go-swagger; DO NOT EDIT.

package jssuser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// JssuserGetReader is a Reader for the JssuserGet structure.
type JssuserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JssuserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJssuserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /jssuser] JssuserGet", response, response.Code())
	}
}

// NewJssuserGetOK creates a JssuserGetOK with default headers values
func NewJssuserGetOK() *JssuserGetOK {
	return &JssuserGetOK{}
}

/*
JssuserGetOK describes a response with status code 200, with default header values.

OK
*/
type JssuserGetOK struct {
	Payload *models.JssUser
}

// IsSuccess returns true when this jssuser get o k response has a 2xx status code
func (o *JssuserGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this jssuser get o k response has a 3xx status code
func (o *JssuserGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jssuser get o k response has a 4xx status code
func (o *JssuserGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this jssuser get o k response has a 5xx status code
func (o *JssuserGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this jssuser get o k response a status code equal to that given
func (o *JssuserGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the jssuser get o k response
func (o *JssuserGetOK) Code() int {
	return 200
}

func (o *JssuserGetOK) Error() string {
	return fmt.Sprintf("[GET /jssuser][%d] jssuserGetOK  %+v", 200, o.Payload)
}

func (o *JssuserGetOK) String() string {
	return fmt.Sprintf("[GET /jssuser][%d] jssuserGetOK  %+v", 200, o.Payload)
}

func (o *JssuserGetOK) GetPayload() *models.JssUser {
	return o.Payload
}

func (o *JssuserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JssUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
