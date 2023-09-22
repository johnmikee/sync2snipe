// Code generated by go-swagger; DO NOT EDIT.

package diskencryptionconfigurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDiskEncryptionConfigurationByIDReader is a Reader for the DeleteDiskEncryptionConfigurationByID structure.
type DeleteDiskEncryptionConfigurationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDiskEncryptionConfigurationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDiskEncryptionConfigurationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /diskencryptionconfigurations/id/{id}] deleteDiskEncryptionConfigurationById", response, response.Code())
	}
}

// NewDeleteDiskEncryptionConfigurationByIDOK creates a DeleteDiskEncryptionConfigurationByIDOK with default headers values
func NewDeleteDiskEncryptionConfigurationByIDOK() *DeleteDiskEncryptionConfigurationByIDOK {
	return &DeleteDiskEncryptionConfigurationByIDOK{}
}

/*
DeleteDiskEncryptionConfigurationByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDiskEncryptionConfigurationByIDOK struct {
}

// IsSuccess returns true when this delete disk encryption configuration by Id o k response has a 2xx status code
func (o *DeleteDiskEncryptionConfigurationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete disk encryption configuration by Id o k response has a 3xx status code
func (o *DeleteDiskEncryptionConfigurationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete disk encryption configuration by Id o k response has a 4xx status code
func (o *DeleteDiskEncryptionConfigurationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete disk encryption configuration by Id o k response has a 5xx status code
func (o *DeleteDiskEncryptionConfigurationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete disk encryption configuration by Id o k response a status code equal to that given
func (o *DeleteDiskEncryptionConfigurationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete disk encryption configuration by Id o k response
func (o *DeleteDiskEncryptionConfigurationByIDOK) Code() int {
	return 200
}

func (o *DeleteDiskEncryptionConfigurationByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /diskencryptionconfigurations/id/{id}][%d] deleteDiskEncryptionConfigurationByIdOK ", 200)
}

func (o *DeleteDiskEncryptionConfigurationByIDOK) String() string {
	return fmt.Sprintf("[DELETE /diskencryptionconfigurations/id/{id}][%d] deleteDiskEncryptionConfigurationByIdOK ", 200)
}

func (o *DeleteDiskEncryptionConfigurationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
