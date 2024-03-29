// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: google/protobuf/compiler/plugin.proto

package pluginpb

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

// Validate checks the field values on Version with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Version) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Version with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VersionMultiError, or nil if none found.
func (m *Version) ValidateAll() error {
	return m.validate(true)
}

func (m *Version) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Major

	// no validation rules for Minor

	// no validation rules for Patch

	// no validation rules for Suffix

	if len(errors) > 0 {
		return VersionMultiError(errors)
	}

	return nil
}

// VersionMultiError is an error wrapping multiple validation errors returned
// by Version.ValidateAll() if the designated constraints aren't met.
type VersionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VersionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VersionMultiError) AllErrors() []error { return m }

// VersionValidationError is the validation error returned by Version.Validate
// if the designated constraints aren't met.
type VersionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VersionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VersionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VersionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VersionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VersionValidationError) ErrorName() string { return "VersionValidationError" }

// Error satisfies the builtin error interface
func (e VersionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVersion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VersionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VersionValidationError{}

// Validate checks the field values on CodeGeneratorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CodeGeneratorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CodeGeneratorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CodeGeneratorRequestMultiError, or nil if none found.
func (m *CodeGeneratorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CodeGeneratorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Parameter

	for idx, item := range m.GetProtoFile() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CodeGeneratorRequestValidationError{
						field:  fmt.Sprintf("ProtoFile[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CodeGeneratorRequestValidationError{
						field:  fmt.Sprintf("ProtoFile[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CodeGeneratorRequestValidationError{
					field:  fmt.Sprintf("ProtoFile[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCompilerVersion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CodeGeneratorRequestValidationError{
					field:  "CompilerVersion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CodeGeneratorRequestValidationError{
					field:  "CompilerVersion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompilerVersion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CodeGeneratorRequestValidationError{
				field:  "CompilerVersion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CodeGeneratorRequestMultiError(errors)
	}

	return nil
}

// CodeGeneratorRequestMultiError is an error wrapping multiple validation
// errors returned by CodeGeneratorRequest.ValidateAll() if the designated
// constraints aren't met.
type CodeGeneratorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CodeGeneratorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CodeGeneratorRequestMultiError) AllErrors() []error { return m }

// CodeGeneratorRequestValidationError is the validation error returned by
// CodeGeneratorRequest.Validate if the designated constraints aren't met.
type CodeGeneratorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeGeneratorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeGeneratorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeGeneratorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeGeneratorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeGeneratorRequestValidationError) ErrorName() string {
	return "CodeGeneratorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CodeGeneratorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeGeneratorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeGeneratorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeGeneratorRequestValidationError{}

// Validate checks the field values on CodeGeneratorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CodeGeneratorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CodeGeneratorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CodeGeneratorResponseMultiError, or nil if none found.
func (m *CodeGeneratorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CodeGeneratorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Error

	// no validation rules for SupportedFeatures

	for idx, item := range m.GetFile() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CodeGeneratorResponseValidationError{
						field:  fmt.Sprintf("File[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CodeGeneratorResponseValidationError{
						field:  fmt.Sprintf("File[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CodeGeneratorResponseValidationError{
					field:  fmt.Sprintf("File[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CodeGeneratorResponseMultiError(errors)
	}

	return nil
}

// CodeGeneratorResponseMultiError is an error wrapping multiple validation
// errors returned by CodeGeneratorResponse.ValidateAll() if the designated
// constraints aren't met.
type CodeGeneratorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CodeGeneratorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CodeGeneratorResponseMultiError) AllErrors() []error { return m }

// CodeGeneratorResponseValidationError is the validation error returned by
// CodeGeneratorResponse.Validate if the designated constraints aren't met.
type CodeGeneratorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeGeneratorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeGeneratorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeGeneratorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeGeneratorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeGeneratorResponseValidationError) ErrorName() string {
	return "CodeGeneratorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CodeGeneratorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeGeneratorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeGeneratorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeGeneratorResponseValidationError{}

// Validate checks the field values on CodeGeneratorResponse_File with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CodeGeneratorResponse_File) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CodeGeneratorResponse_File with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CodeGeneratorResponse_FileMultiError, or nil if none found.
func (m *CodeGeneratorResponse_File) ValidateAll() error {
	return m.validate(true)
}

func (m *CodeGeneratorResponse_File) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for InsertionPoint

	// no validation rules for Content

	if all {
		switch v := interface{}(m.GetGeneratedCodeInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CodeGeneratorResponse_FileValidationError{
					field:  "GeneratedCodeInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CodeGeneratorResponse_FileValidationError{
					field:  "GeneratedCodeInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGeneratedCodeInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CodeGeneratorResponse_FileValidationError{
				field:  "GeneratedCodeInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CodeGeneratorResponse_FileMultiError(errors)
	}

	return nil
}

// CodeGeneratorResponse_FileMultiError is an error wrapping multiple
// validation errors returned by CodeGeneratorResponse_File.ValidateAll() if
// the designated constraints aren't met.
type CodeGeneratorResponse_FileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CodeGeneratorResponse_FileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CodeGeneratorResponse_FileMultiError) AllErrors() []error { return m }

// CodeGeneratorResponse_FileValidationError is the validation error returned
// by CodeGeneratorResponse_File.Validate if the designated constraints aren't met.
type CodeGeneratorResponse_FileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeGeneratorResponse_FileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeGeneratorResponse_FileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeGeneratorResponse_FileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeGeneratorResponse_FileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeGeneratorResponse_FileValidationError) ErrorName() string {
	return "CodeGeneratorResponse_FileValidationError"
}

// Error satisfies the builtin error interface
func (e CodeGeneratorResponse_FileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeGeneratorResponse_File.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeGeneratorResponse_FileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeGeneratorResponse_FileValidationError{}
