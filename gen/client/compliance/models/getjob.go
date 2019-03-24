// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Getjob getjob
// swagger:model getjob
type Getjob struct {

	// boardurl
	Boardurl string `json:"boardurl,omitempty"`

	// compliancetype
	Compliancetype string `json:"compliancetype,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// progress
	Progress int64 `json:"progress,omitempty"`

	// result
	Result Ruleresultarray `json:"result,omitempty"`

	// scanstatus
	Scanstatus string `json:"scanstatus,omitempty"`
}

// Validate validates this getjob
func (m *Getjob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Getjob) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Getjob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Getjob) UnmarshalBinary(b []byte) error {
	var res Getjob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}