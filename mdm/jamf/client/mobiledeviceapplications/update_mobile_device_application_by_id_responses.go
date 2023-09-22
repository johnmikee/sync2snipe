// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceapplications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMobileDeviceApplicationByIDReader is a Reader for the UpdateMobileDeviceApplicationByID structure.
type UpdateMobileDeviceApplicationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMobileDeviceApplicationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMobileDeviceApplicationByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /mobiledeviceapplications/id/{id}] updateMobileDeviceApplicationById", response, response.Code())
	}
}

// NewUpdateMobileDeviceApplicationByIDCreated creates a UpdateMobileDeviceApplicationByIDCreated with default headers values
func NewUpdateMobileDeviceApplicationByIDCreated() *UpdateMobileDeviceApplicationByIDCreated {
	return &UpdateMobileDeviceApplicationByIDCreated{}
}

/*
UpdateMobileDeviceApplicationByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMobileDeviceApplicationByIDCreated struct {
}

// IsSuccess returns true when this update mobile device application by Id created response has a 2xx status code
func (o *UpdateMobileDeviceApplicationByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mobile device application by Id created response has a 3xx status code
func (o *UpdateMobileDeviceApplicationByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mobile device application by Id created response has a 4xx status code
func (o *UpdateMobileDeviceApplicationByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mobile device application by Id created response has a 5xx status code
func (o *UpdateMobileDeviceApplicationByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update mobile device application by Id created response a status code equal to that given
func (o *UpdateMobileDeviceApplicationByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update mobile device application by Id created response
func (o *UpdateMobileDeviceApplicationByIDCreated) Code() int {
	return 201
}

func (o *UpdateMobileDeviceApplicationByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /mobiledeviceapplications/id/{id}][%d] updateMobileDeviceApplicationByIdCreated ", 201)
}

func (o *UpdateMobileDeviceApplicationByIDCreated) String() string {
	return fmt.Sprintf("[PUT /mobiledeviceapplications/id/{id}][%d] updateMobileDeviceApplicationByIdCreated ", 201)
}

func (o *UpdateMobileDeviceApplicationByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}