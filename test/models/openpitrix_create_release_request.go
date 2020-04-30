// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateReleaseRequest openpitrix create release request
// swagger:model openpitrixCreateReleaseRequest
type OpenpitrixCreateReleaseRequest struct {

	// required, id of app to run in cluster
	AppID string `json:"app_id,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// release name
	ReleaseName string `json:"release_name,omitempty"`

	// required, id of runtime
	RuntimeID string `json:"runtime_id,omitempty"`

	// required, id of app version
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix create release request
func (m *OpenpitrixCreateReleaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateReleaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateReleaseRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateReleaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}