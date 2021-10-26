// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainTree domain tree
//
// swagger:model domain.Tree
type DomainTree struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this domain tree
func (m *DomainTree) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain tree based on context it is used
func (m *DomainTree) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainTree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainTree) UnmarshalBinary(b []byte) error {
	var res DomainTree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
