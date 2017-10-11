// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetClient get client
// swagger:model getClient

type GetClient struct {

	// Name of the company
	// Required: true
	CompanyName *string `json:"companyName"`

	// Login Email
	// Required: true
	Email *strfmt.Email `json:"email"`

	// First Name
	// Required: true
	FirstName *string `json:"firstName"`

	// Last Name
	// Required: true
	LastName *string `json:"lastName"`
}

/* polymorph getClient companyName false */

/* polymorph getClient email false */

/* polymorph getClient firstName false */

/* polymorph getClient lastName false */

// Validate validates this get client
func (m *GetClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanyName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClient) validateCompanyName(formats strfmt.Registry) error {

	if err := validate.Required("companyName", "body", m.CompanyName); err != nil {
		return err
	}

	return nil
}

func (m *GetClient) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetClient) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *GetClient) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClient) UnmarshalBinary(b []byte) error {
	var res GetClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}