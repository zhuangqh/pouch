// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkConnect contains the request for the remote API: POST /networks/{id:.*}/connect
// swagger:model NetworkConnect
type NetworkConnect struct {

	// The ID or name of the container to connect to the network.
	Container string `json:"Container,omitempty"`

	// endpoint config
	EndpointConfig *EndpointSettings `json:"EndpointConfig,omitempty"`
}

// Validate validates this network connect
func (m *NetworkConnect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConnect) validateEndpointConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointConfig) { // not required
		return nil
	}

	if m.EndpointConfig != nil {
		if err := m.EndpointConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkConnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkConnect) UnmarshalBinary(b []byte) error {
	var res NetworkConnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
