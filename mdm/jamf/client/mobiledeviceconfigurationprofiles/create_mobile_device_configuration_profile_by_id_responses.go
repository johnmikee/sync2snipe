// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceconfigurationprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateMobileDeviceConfigurationProfileByIDReader is a Reader for the CreateMobileDeviceConfigurationProfileByID structure.
type CreateMobileDeviceConfigurationProfileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMobileDeviceConfigurationProfileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMobileDeviceConfigurationProfileByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /mobiledeviceconfigurationprofiles/id/{id}] createMobileDeviceConfigurationProfileById", response, response.Code())
	}
}

// NewCreateMobileDeviceConfigurationProfileByIDCreated creates a CreateMobileDeviceConfigurationProfileByIDCreated with default headers values
func NewCreateMobileDeviceConfigurationProfileByIDCreated() *CreateMobileDeviceConfigurationProfileByIDCreated {
	return &CreateMobileDeviceConfigurationProfileByIDCreated{}
}

/*
CreateMobileDeviceConfigurationProfileByIDCreated describes a response with status code 201, with default header values.

Created
*/
type CreateMobileDeviceConfigurationProfileByIDCreated struct {
}

// IsSuccess returns true when this create mobile device configuration profile by Id created response has a 2xx status code
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mobile device configuration profile by Id created response has a 3xx status code
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mobile device configuration profile by Id created response has a 4xx status code
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mobile device configuration profile by Id created response has a 5xx status code
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create mobile device configuration profile by Id created response a status code equal to that given
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create mobile device configuration profile by Id created response
func (o *CreateMobileDeviceConfigurationProfileByIDCreated) Code() int {
	return 201
}

func (o *CreateMobileDeviceConfigurationProfileByIDCreated) Error() string {
	return fmt.Sprintf("[POST /mobiledeviceconfigurationprofiles/id/{id}][%d] createMobileDeviceConfigurationProfileByIdCreated ", 201)
}

func (o *CreateMobileDeviceConfigurationProfileByIDCreated) String() string {
	return fmt.Sprintf("[POST /mobiledeviceconfigurationprofiles/id/{id}][%d] createMobileDeviceConfigurationProfileByIdCreated ", 201)
}

func (o *CreateMobileDeviceConfigurationProfileByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}