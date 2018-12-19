// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteAttachmentsResponse openpitrix delete attachments response
// swagger:model openpitrixDeleteAttachmentsResponse
type OpenpitrixDeleteAttachmentsResponse struct {

	// attachment id
	AttachmentID []string `json:"attachment_id"`

	// filename
	Filename []string `json:"filename"`
}

// Validate validates this openpitrix delete attachments response
func (m *OpenpitrixDeleteAttachmentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteAttachmentsResponse) validateAttachmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachmentID) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixDeleteAttachmentsResponse) validateFilename(formats strfmt.Registry) error {

	if swag.IsZero(m.Filename) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteAttachmentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteAttachmentsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteAttachmentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}