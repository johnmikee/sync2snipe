// Code generated by go-swagger; DO NOT EDIT.

package mobiledevicegroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateMobileDeviceGroupByIDReader is a Reader for the UpdateMobileDeviceGroupByID structure.
type UpdateMobileDeviceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMobileDeviceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMobileDeviceGroupByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /mobiledevicegroups/id/{id}] updateMobileDeviceGroupById", response, response.Code())
	}
}

// NewUpdateMobileDeviceGroupByIDCreated creates a UpdateMobileDeviceGroupByIDCreated with default headers values
func NewUpdateMobileDeviceGroupByIDCreated() *UpdateMobileDeviceGroupByIDCreated {
	return &UpdateMobileDeviceGroupByIDCreated{}
}

/*
UpdateMobileDeviceGroupByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateMobileDeviceGroupByIDCreated struct {
}

// IsSuccess returns true when this update mobile device group by Id created response has a 2xx status code
func (o *UpdateMobileDeviceGroupByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update mobile device group by Id created response has a 3xx status code
func (o *UpdateMobileDeviceGroupByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update mobile device group by Id created response has a 4xx status code
func (o *UpdateMobileDeviceGroupByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update mobile device group by Id created response has a 5xx status code
func (o *UpdateMobileDeviceGroupByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update mobile device group by Id created response a status code equal to that given
func (o *UpdateMobileDeviceGroupByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update mobile device group by Id created response
func (o *UpdateMobileDeviceGroupByIDCreated) Code() int {
	return 201
}

func (o *UpdateMobileDeviceGroupByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /mobiledevicegroups/id/{id}][%d] updateMobileDeviceGroupByIdCreated ", 201)
}

func (o *UpdateMobileDeviceGroupByIDCreated) String() string {
	return fmt.Sprintf("[PUT /mobiledevicegroups/id/{id}][%d] updateMobileDeviceGroupByIdCreated ", 201)
}

func (o *UpdateMobileDeviceGroupByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
