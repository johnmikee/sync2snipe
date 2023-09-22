// Code generated by go-swagger; DO NOT EDIT.

package diskencryptionconfigurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDiskEncryptionConfigurationByNameReader is a Reader for the DeleteDiskEncryptionConfigurationByName structure.
type DeleteDiskEncryptionConfigurationByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDiskEncryptionConfigurationByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDiskEncryptionConfigurationByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /diskencryptionconfigurations/name/{name}] deleteDiskEncryptionConfigurationByName", response, response.Code())
	}
}

// NewDeleteDiskEncryptionConfigurationByNameOK creates a DeleteDiskEncryptionConfigurationByNameOK with default headers values
func NewDeleteDiskEncryptionConfigurationByNameOK() *DeleteDiskEncryptionConfigurationByNameOK {
	return &DeleteDiskEncryptionConfigurationByNameOK{}
}

/*
DeleteDiskEncryptionConfigurationByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDiskEncryptionConfigurationByNameOK struct {
}

// IsSuccess returns true when this delete disk encryption configuration by name o k response has a 2xx status code
func (o *DeleteDiskEncryptionConfigurationByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete disk encryption configuration by name o k response has a 3xx status code
func (o *DeleteDiskEncryptionConfigurationByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete disk encryption configuration by name o k response has a 4xx status code
func (o *DeleteDiskEncryptionConfigurationByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete disk encryption configuration by name o k response has a 5xx status code
func (o *DeleteDiskEncryptionConfigurationByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete disk encryption configuration by name o k response a status code equal to that given
func (o *DeleteDiskEncryptionConfigurationByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete disk encryption configuration by name o k response
func (o *DeleteDiskEncryptionConfigurationByNameOK) Code() int {
	return 200
}

func (o *DeleteDiskEncryptionConfigurationByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /diskencryptionconfigurations/name/{name}][%d] deleteDiskEncryptionConfigurationByNameOK ", 200)
}

func (o *DeleteDiskEncryptionConfigurationByNameOK) String() string {
	return fmt.Sprintf("[DELETE /diskencryptionconfigurations/name/{name}][%d] deleteDiskEncryptionConfigurationByNameOK ", 200)
}

func (o *DeleteDiskEncryptionConfigurationByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}