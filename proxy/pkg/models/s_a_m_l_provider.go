// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SAMLProvider s a m l provider
//
// swagger:model SAMLProvider
type SAMLProvider struct {

	// ACS URL
	// Required: true
	// Max Length: 200
	// Min Length: 1
	// Format: uri
	AcsURL *strfmt.URI `json:"acs_url"`

	// Assertion valid not before
	//
	// Assertion valid not before current time + this value (Format: hours=-1;minutes=-2;seconds=-3).
	// Min Length: 1
	AssertionValidNotBefore string `json:"assertion_valid_not_before,omitempty"`

	// Assertion valid not on or after
	//
	// Assertion not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	// Min Length: 1
	AssertionValidNotOnOrAfter string `json:"assertion_valid_not_on_or_after,omitempty"`

	// Audience
	// Min Length: 1
	Audience string `json:"audience,omitempty"`

	// Digest algorithm
	// Enum: [sha1 sha256]
	DigestAlgorithm string `json:"digest_algorithm,omitempty"`

	// Issuer
	//
	// Also known as EntityID
	// Required: true
	// Min Length: 1
	Issuer *string `json:"issuer"`

	// Name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// ID
	// Read Only: true
	Pk int64 `json:"pk,omitempty"`

	// property mappings
	// Unique: true
	PropertyMappings []strfmt.UUID `json:"property_mappings"`

	// Require signing
	//
	// Require Requests to be signed by an X509 Certificate. Must match the Certificate selected in `Singing Keypair`.
	RequireSigning bool `json:"require_signing,omitempty"`

	// Session valid not on or after
	//
	// Session not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	// Min Length: 1
	SessionValidNotOnOrAfter string `json:"session_valid_not_on_or_after,omitempty"`

	// Signature algorithm
	// Enum: [rsa-sha1 rsa-sha256 ecdsa-sha256 dsa-sha1]
	SignatureAlgorithm string `json:"signature_algorithm,omitempty"`

	// Signing Keypair
	//
	// Singing is enabled upon selection of a Key Pair.
	// Format: uuid
	SigningKp *strfmt.UUID `json:"signing_kp,omitempty"`
}

// Validate validates this s a m l provider
func (m *SAMLProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcsURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssertionValidNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssertionValidNotOnOrAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigestAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionValidNotOnOrAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatureAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigningKp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAMLProvider) validateAcsURL(formats strfmt.Registry) error {

	if err := validate.Required("acs_url", "body", m.AcsURL); err != nil {
		return err
	}

	if err := validate.MinLength("acs_url", "body", string(*m.AcsURL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("acs_url", "body", string(*m.AcsURL), 200); err != nil {
		return err
	}

	if err := validate.FormatOf("acs_url", "body", "uri", m.AcsURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateAssertionValidNotBefore(formats strfmt.Registry) error {

	if swag.IsZero(m.AssertionValidNotBefore) { // not required
		return nil
	}

	if err := validate.MinLength("assertion_valid_not_before", "body", string(m.AssertionValidNotBefore), 1); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateAssertionValidNotOnOrAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.AssertionValidNotOnOrAfter) { // not required
		return nil
	}

	if err := validate.MinLength("assertion_valid_not_on_or_after", "body", string(m.AssertionValidNotOnOrAfter), 1); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateAudience(formats strfmt.Registry) error {

	if swag.IsZero(m.Audience) { // not required
		return nil
	}

	if err := validate.MinLength("audience", "body", string(m.Audience), 1); err != nil {
		return err
	}

	return nil
}

var sAMLProviderTypeDigestAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha1","sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sAMLProviderTypeDigestAlgorithmPropEnum = append(sAMLProviderTypeDigestAlgorithmPropEnum, v)
	}
}

const (

	// SAMLProviderDigestAlgorithmSha1 captures enum value "sha1"
	SAMLProviderDigestAlgorithmSha1 string = "sha1"

	// SAMLProviderDigestAlgorithmSha256 captures enum value "sha256"
	SAMLProviderDigestAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *SAMLProvider) validateDigestAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sAMLProviderTypeDigestAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SAMLProvider) validateDigestAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.DigestAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateDigestAlgorithmEnum("digest_algorithm", "body", m.DigestAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	if err := validate.MinLength("issuer", "body", string(*m.Issuer), 1); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validatePropertyMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyMappings) { // not required
		return nil
	}

	if err := validate.UniqueItems("property_mappings", "body", m.PropertyMappings); err != nil {
		return err
	}

	for i := 0; i < len(m.PropertyMappings); i++ {

		if err := validate.FormatOf("property_mappings"+"."+strconv.Itoa(i), "body", "uuid", m.PropertyMappings[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *SAMLProvider) validateSessionValidNotOnOrAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionValidNotOnOrAfter) { // not required
		return nil
	}

	if err := validate.MinLength("session_valid_not_on_or_after", "body", string(m.SessionValidNotOnOrAfter), 1); err != nil {
		return err
	}

	return nil
}

var sAMLProviderTypeSignatureAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rsa-sha1","rsa-sha256","ecdsa-sha256","dsa-sha1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sAMLProviderTypeSignatureAlgorithmPropEnum = append(sAMLProviderTypeSignatureAlgorithmPropEnum, v)
	}
}

const (

	// SAMLProviderSignatureAlgorithmRsaSha1 captures enum value "rsa-sha1"
	SAMLProviderSignatureAlgorithmRsaSha1 string = "rsa-sha1"

	// SAMLProviderSignatureAlgorithmRsaSha256 captures enum value "rsa-sha256"
	SAMLProviderSignatureAlgorithmRsaSha256 string = "rsa-sha256"

	// SAMLProviderSignatureAlgorithmEcdsaSha256 captures enum value "ecdsa-sha256"
	SAMLProviderSignatureAlgorithmEcdsaSha256 string = "ecdsa-sha256"

	// SAMLProviderSignatureAlgorithmDsaSha1 captures enum value "dsa-sha1"
	SAMLProviderSignatureAlgorithmDsaSha1 string = "dsa-sha1"
)

// prop value enum
func (m *SAMLProvider) validateSignatureAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sAMLProviderTypeSignatureAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SAMLProvider) validateSignatureAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.SignatureAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateSignatureAlgorithmEnum("signature_algorithm", "body", m.SignatureAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *SAMLProvider) validateSigningKp(formats strfmt.Registry) error {

	if swag.IsZero(m.SigningKp) { // not required
		return nil
	}

	if err := validate.FormatOf("signing_kp", "body", "uuid", m.SigningKp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAMLProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAMLProvider) UnmarshalBinary(b []byte) error {
	var res SAMLProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
