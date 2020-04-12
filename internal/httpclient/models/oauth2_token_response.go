// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Oauth2TokenResponse Oauth2TokenResponse Oauth2TokenResponse Oauth2TokenResponse Oauth2TokenResponse The Access Token Response
//
// swagger:model oauth2TokenResponse
type Oauth2TokenResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// id token
	IDToken string `json:"id_token,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this oauth2 token response
func (m *Oauth2TokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Oauth2TokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Oauth2TokenResponse) UnmarshalBinary(b []byte) error {
	var res Oauth2TokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
