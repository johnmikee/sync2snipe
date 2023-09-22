// Code generated by go-swagger; DO NOT EDIT.

package jsonwebtokenconfigurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// CreateJSONWebTokenConfigurationByIDReader is a Reader for the CreateJSONWebTokenConfigurationByID structure.
type CreateJSONWebTokenConfigurationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJSONWebTokenConfigurationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateJSONWebTokenConfigurationByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /jsonwebtokenconfigurations/id/{id}] createJsonWebTokenConfigurationById", response, response.Code())
	}
}

// NewCreateJSONWebTokenConfigurationByIDCreated creates a CreateJSONWebTokenConfigurationByIDCreated with default headers values
func NewCreateJSONWebTokenConfigurationByIDCreated() *CreateJSONWebTokenConfigurationByIDCreated {
	return &CreateJSONWebTokenConfigurationByIDCreated{}
}

/*
CreateJSONWebTokenConfigurationByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateJSONWebTokenConfigurationByIDCreated struct {
	Payload *models.JSONWebTokenConfiguration
}

// IsSuccess returns true when this create Json web token configuration by Id created response has a 2xx status code
func (o *CreateJSONWebTokenConfigurationByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Json web token configuration by Id created response has a 3xx status code
func (o *CreateJSONWebTokenConfigurationByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Json web token configuration by Id created response has a 4xx status code
func (o *CreateJSONWebTokenConfigurationByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Json web token configuration by Id created response has a 5xx status code
func (o *CreateJSONWebTokenConfigurationByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Json web token configuration by Id created response a status code equal to that given
func (o *CreateJSONWebTokenConfigurationByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Json web token configuration by Id created response
func (o *CreateJSONWebTokenConfigurationByIDCreated) Code() int {
	return 201
}

func (o *CreateJSONWebTokenConfigurationByIDCreated) Error() string {
	return fmt.Sprintf("[POST /jsonwebtokenconfigurations/id/{id}][%d] createJsonWebTokenConfigurationByIdCreated  %+v", 201, o.Payload)
}

func (o *CreateJSONWebTokenConfigurationByIDCreated) String() string {
	return fmt.Sprintf("[POST /jsonwebtokenconfigurations/id/{id}][%d] createJsonWebTokenConfigurationByIdCreated  %+v", 201, o.Payload)
}

func (o *CreateJSONWebTokenConfigurationByIDCreated) GetPayload() *models.JSONWebTokenConfiguration {
	return o.Payload
}

func (o *CreateJSONWebTokenConfigurationByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebTokenConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
