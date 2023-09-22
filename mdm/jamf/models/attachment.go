// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Attachment attachment
//
// swagger:model attachment
type Attachment struct {

	// filename
	// Example: icon.png
	Filename string `json:"filename,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// uri
	// Example: https://example.jamfcloud/attachment.html?id=1\u0026amp;o=r
	URI string `json:"uri,omitempty"`
}

// Validate validates this attachment
func (m *Attachment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attachment based on context it is used
func (m *Attachment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Attachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attachment) UnmarshalBinary(b []byte) error {
	var res Attachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
