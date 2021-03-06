// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/auth.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Tokens with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tokens) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tokens with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokensMultiError, or nil if none found.
func (m *Tokens) ValidateAll() error {
	return m.validate(true)
}

func (m *Tokens) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return TokensMultiError(errors)
	}

	return nil
}

// TokensMultiError is an error wrapping multiple validation errors returned by
// Tokens.ValidateAll() if the designated constraints aren't met.
type TokensMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokensMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokensMultiError) AllErrors() []error { return m }

// TokensValidationError is the validation error returned by Tokens.Validate if
// the designated constraints aren't met.
type TokensValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokensValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokensValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokensValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokensValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokensValidationError) ErrorName() string { return "TokensValidationError" }

// Error satisfies the builtin error interface
func (e TokensValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokens.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokensValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokensValidationError{}

// Validate checks the field values on SignUpRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpRequestMultiError, or
// nil if none found.
func (m *SignUpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 4 || l > 16 {
		err := SignUpRequestValidationError{
			field:  "Name",
			reason: "value length must be between 4 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetEmail()) < 6 {
		err := SignUpRequestValidationError{
			field:  "Email",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := SignUpRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SignUpRequestMultiError(errors)
	}

	return nil
}

// SignUpRequestMultiError is an error wrapping multiple validation errors
// returned by SignUpRequest.ValidateAll() if the designated constraints
// aren't met.
type SignUpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpRequestMultiError) AllErrors() []error { return m }

// SignUpRequestValidationError is the validation error returned by
// SignUpRequest.Validate if the designated constraints aren't met.
type SignUpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpRequestValidationError) ErrorName() string { return "SignUpRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignUpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpRequestValidationError{}

// Validate checks the field values on SignUpReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpReplyMultiError, or
// nil if none found.
func (m *SignUpReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetTokens()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignUpReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignUpReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokens()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignUpReplyValidationError{
				field:  "Tokens",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SignUpReplyMultiError(errors)
	}

	return nil
}

// SignUpReplyMultiError is an error wrapping multiple validation errors
// returned by SignUpReply.ValidateAll() if the designated constraints aren't met.
type SignUpReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpReplyMultiError) AllErrors() []error { return m }

// SignUpReplyValidationError is the validation error returned by
// SignUpReply.Validate if the designated constraints aren't met.
type SignUpReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpReplyValidationError) ErrorName() string { return "SignUpReplyValidationError" }

// Error satisfies the builtin error interface
func (e SignUpReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpReplyValidationError{}

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInRequestMultiError, or
// nil if none found.
func (m *SignInRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetEmail()) < 6 {
		err := SignInRequestValidationError{
			field:  "Email",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := SignInRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SignInRequestMultiError(errors)
	}

	return nil
}

// SignInRequestMultiError is an error wrapping multiple validation errors
// returned by SignInRequest.ValidateAll() if the designated constraints
// aren't met.
type SignInRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInRequestMultiError) AllErrors() []error { return m }

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

// Validate checks the field values on SignInReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInReplyMultiError, or
// nil if none found.
func (m *SignInReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetTokens()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignInReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignInReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokens()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignInReplyValidationError{
				field:  "Tokens",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SignInReplyMultiError(errors)
	}

	return nil
}

// SignInReplyMultiError is an error wrapping multiple validation errors
// returned by SignInReply.ValidateAll() if the designated constraints aren't met.
type SignInReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInReplyMultiError) AllErrors() []error { return m }

// SignInReplyValidationError is the validation error returned by
// SignInReply.Validate if the designated constraints aren't met.
type SignInReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInReplyValidationError) ErrorName() string { return "SignInReplyValidationError" }

// Error satisfies the builtin error interface
func (e SignInReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInReplyValidationError{}

// Validate checks the field values on RefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenRequestMultiError, or nil if none found.
func (m *RefreshTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return RefreshTokenRequestMultiError(errors)
	}

	return nil
}

// RefreshTokenRequestMultiError is an error wrapping multiple validation
// errors returned by RefreshTokenRequest.ValidateAll() if the designated
// constraints aren't met.
type RefreshTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenRequestMultiError) AllErrors() []error { return m }

// RefreshTokenRequestValidationError is the validation error returned by
// RefreshTokenRequest.Validate if the designated constraints aren't met.
type RefreshTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenRequestValidationError) ErrorName() string {
	return "RefreshTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenRequestValidationError{}

// Validate checks the field values on RefreshTokenReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenReplyMultiError, or nil if none found.
func (m *RefreshTokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTokens()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RefreshTokenReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RefreshTokenReplyValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokens()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RefreshTokenReplyValidationError{
				field:  "Tokens",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RefreshTokenReplyMultiError(errors)
	}

	return nil
}

// RefreshTokenReplyMultiError is an error wrapping multiple validation errors
// returned by RefreshTokenReply.ValidateAll() if the designated constraints
// aren't met.
type RefreshTokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenReplyMultiError) AllErrors() []error { return m }

// RefreshTokenReplyValidationError is the validation error returned by
// RefreshTokenReply.Validate if the designated constraints aren't met.
type RefreshTokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenReplyValidationError) ErrorName() string {
	return "RefreshTokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenReplyValidationError{}

// Validate checks the field values on IdentityRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdentityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdentityRequestMultiError, or nil if none found.
func (m *IdentityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return IdentityRequestMultiError(errors)
	}

	return nil
}

// IdentityRequestMultiError is an error wrapping multiple validation errors
// returned by IdentityRequest.ValidateAll() if the designated constraints
// aren't met.
type IdentityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentityRequestMultiError) AllErrors() []error { return m }

// IdentityRequestValidationError is the validation error returned by
// IdentityRequest.Validate if the designated constraints aren't met.
type IdentityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityRequestValidationError) ErrorName() string { return "IdentityRequestValidationError" }

// Error satisfies the builtin error interface
func (e IdentityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityRequestValidationError{}

// Validate checks the field values on IdentityReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdentityReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentityReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdentityReplyMultiError, or
// nil if none found.
func (m *IdentityReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentityReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdentityReplyMultiError(errors)
	}

	return nil
}

// IdentityReplyMultiError is an error wrapping multiple validation errors
// returned by IdentityReply.ValidateAll() if the designated constraints
// aren't met.
type IdentityReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentityReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentityReplyMultiError) AllErrors() []error { return m }

// IdentityReplyValidationError is the validation error returned by
// IdentityReply.Validate if the designated constraints aren't met.
type IdentityReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityReplyValidationError) ErrorName() string { return "IdentityReplyValidationError" }

// Error satisfies the builtin error interface
func (e IdentityReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentityReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityReplyValidationError{}
