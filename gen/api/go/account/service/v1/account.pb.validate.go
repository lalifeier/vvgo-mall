// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/service/v1/account.proto

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

// Validate checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReqMultiError, or
// nil if none found.
func (m *RegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Email

	// no validation rules for Phone

	if len(errors) > 0 {
		return RegisterReqMultiError(errors)
	}

	return nil
}

// RegisterReqMultiError is an error wrapping multiple validation errors
// returned by RegisterReq.ValidateAll() if the designated constraints aren't met.
type RegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReqMultiError) AllErrors() []error { return m }

// RegisterReqValidationError is the validation error returned by
// RegisterReq.Validate if the designated constraints aren't met.
type RegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReqValidationError) ErrorName() string { return "RegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReqValidationError{}

// Validate checks the field values on RegisterResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterRespMultiError, or
// nil if none found.
func (m *RegisterResp) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RegisterRespMultiError(errors)
	}

	return nil
}

// RegisterRespMultiError is an error wrapping multiple validation errors
// returned by RegisterResp.ValidateAll() if the designated constraints aren't met.
type RegisterRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRespMultiError) AllErrors() []error { return m }

// RegisterRespValidationError is the validation error returned by
// RegisterResp.Validate if the designated constraints aren't met.
type RegisterRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRespValidationError) ErrorName() string { return "RegisterRespValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRespValidationError{}

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Code

	// no validation rules for Type

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}

	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRespMultiError, or nil
// if none found.
func (m *LoginResp) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	// no validation rules for Email

	// no validation rules for Phone

	if len(errors) > 0 {
		return LoginRespMultiError(errors)
	}

	return nil
}

// LoginRespMultiError is an error wrapping multiple validation errors returned
// by LoginResp.ValidateAll() if the designated constraints aren't met.
type LoginRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRespMultiError) AllErrors() []error { return m }

// LoginRespValidationError is the validation error returned by
// LoginResp.Validate if the designated constraints aren't met.
type LoginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRespValidationError) ErrorName() string { return "LoginRespValidationError" }

// Error satisfies the builtin error interface
func (e LoginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRespValidationError{}
