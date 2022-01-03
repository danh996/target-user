// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/target.proto

package pb

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

// Validate checks the field values on Target with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Target) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Target with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TargetMultiError, or nil if none found.
func (m *Target) ValidateAll() error {
	return m.validate(true)
}

func (m *Target) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Slug

	// no validation rules for Created

	// no validation rules for Updated

	// no validation rules for StartDate

	// no validation rules for EndDate

	// no validation rules for Description

	// no validation rules for UserId

	if len(errors) > 0 {
		return TargetMultiError(errors)
	}
	return nil
}

// TargetMultiError is an error wrapping multiple validation errors returned by
// Target.ValidateAll() if the designated constraints aren't met.
type TargetMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TargetMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TargetMultiError) AllErrors() []error { return m }

// TargetValidationError is the validation error returned by Target.Validate if
// the designated constraints aren't met.
type TargetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TargetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TargetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TargetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TargetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TargetValidationError) ErrorName() string { return "TargetValidationError" }

// Error satisfies the builtin error interface
func (e TargetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTarget.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TargetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TargetValidationError{}

// Validate checks the field values on CreateTargetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTargetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTargetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTargetRequestMultiError, or nil if none found.
func (m *CreateTargetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTargetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for StartDate

	// no validation rules for EndDate

	// no validation rules for Description

	// no validation rules for UserId

	if len(errors) > 0 {
		return CreateTargetRequestMultiError(errors)
	}
	return nil
}

// CreateTargetRequestMultiError is an error wrapping multiple validation
// errors returned by CreateTargetRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateTargetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTargetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTargetRequestMultiError) AllErrors() []error { return m }

// CreateTargetRequestValidationError is the validation error returned by
// CreateTargetRequest.Validate if the designated constraints aren't met.
type CreateTargetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTargetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTargetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTargetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTargetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTargetRequestValidationError) ErrorName() string {
	return "CreateTargetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTargetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTargetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTargetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTargetRequestValidationError{}

// Validate checks the field values on CreateTargetReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTargetReponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTargetReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTargetReponseMultiError, or nil if none found.
func (m *CreateTargetReponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTargetReponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTargetReponseValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTargetReponseValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTargetReponseValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTargetReponseMultiError(errors)
	}
	return nil
}

// CreateTargetReponseMultiError is an error wrapping multiple validation
// errors returned by CreateTargetReponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTargetReponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTargetReponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTargetReponseMultiError) AllErrors() []error { return m }

// CreateTargetReponseValidationError is the validation error returned by
// CreateTargetReponse.Validate if the designated constraints aren't met.
type CreateTargetReponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTargetReponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTargetReponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTargetReponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTargetReponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTargetReponseValidationError) ErrorName() string {
	return "CreateTargetReponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTargetReponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTargetReponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTargetReponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTargetReponseValidationError{}

// Validate checks the field values on GetTargetsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTargetsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTargetsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTargetsRequestMultiError, or nil if none found.
func (m *GetTargetsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTargetsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTargetsRequestMultiError(errors)
	}
	return nil
}

// GetTargetsRequestMultiError is an error wrapping multiple validation errors
// returned by GetTargetsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTargetsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTargetsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTargetsRequestMultiError) AllErrors() []error { return m }

// GetTargetsRequestValidationError is the validation error returned by
// GetTargetsRequest.Validate if the designated constraints aren't met.
type GetTargetsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTargetsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTargetsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTargetsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTargetsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTargetsRequestValidationError) ErrorName() string {
	return "GetTargetsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTargetsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTargetsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTargetsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTargetsRequestValidationError{}

// Validate checks the field values on GetTargetsReponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTargetsReponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTargetsReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTargetsReponseMultiError, or nil if none found.
func (m *GetTargetsReponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTargetsReponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTarget() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTargetsReponseValidationError{
						field:  fmt.Sprintf("Target[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTargetsReponseValidationError{
						field:  fmt.Sprintf("Target[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTargetsReponseValidationError{
					field:  fmt.Sprintf("Target[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTargetsReponseMultiError(errors)
	}
	return nil
}

// GetTargetsReponseMultiError is an error wrapping multiple validation errors
// returned by GetTargetsReponse.ValidateAll() if the designated constraints
// aren't met.
type GetTargetsReponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTargetsReponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTargetsReponseMultiError) AllErrors() []error { return m }

// GetTargetsReponseValidationError is the validation error returned by
// GetTargetsReponse.Validate if the designated constraints aren't met.
type GetTargetsReponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTargetsReponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTargetsReponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTargetsReponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTargetsReponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTargetsReponseValidationError) ErrorName() string {
	return "GetTargetsReponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTargetsReponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTargetsReponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTargetsReponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTargetsReponseValidationError{}

// Validate checks the field values on GetTargetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTargetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTargetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTargetRequestMultiError, or nil if none found.
func (m *GetTargetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTargetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TargetId

	if len(errors) > 0 {
		return GetTargetRequestMultiError(errors)
	}
	return nil
}

// GetTargetRequestMultiError is an error wrapping multiple validation errors
// returned by GetTargetRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTargetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTargetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTargetRequestMultiError) AllErrors() []error { return m }

// GetTargetRequestValidationError is the validation error returned by
// GetTargetRequest.Validate if the designated constraints aren't met.
type GetTargetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTargetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTargetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTargetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTargetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTargetRequestValidationError) ErrorName() string { return "GetTargetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTargetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTargetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTargetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTargetRequestValidationError{}

// Validate checks the field values on GetTargetReponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTargetReponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTargetReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTargetReponseMultiError, or nil if none found.
func (m *GetTargetReponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTargetReponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTargetReponseValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTargetReponseValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTargetReponseValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTargetReponseMultiError(errors)
	}
	return nil
}

// GetTargetReponseMultiError is an error wrapping multiple validation errors
// returned by GetTargetReponse.ValidateAll() if the designated constraints
// aren't met.
type GetTargetReponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTargetReponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTargetReponseMultiError) AllErrors() []error { return m }

// GetTargetReponseValidationError is the validation error returned by
// GetTargetReponse.Validate if the designated constraints aren't met.
type GetTargetReponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTargetReponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTargetReponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTargetReponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTargetReponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTargetReponseValidationError) ErrorName() string { return "GetTargetReponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTargetReponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTargetReponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTargetReponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTargetReponseValidationError{}

// Validate checks the field values on UpdateTargetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTargetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTargetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTargetRequestMultiError, or nil if none found.
func (m *UpdateTargetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTargetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for StartDate

	// no validation rules for EndDate

	// no validation rules for TargetId

	// no validation rules for Description

	if len(errors) > 0 {
		return UpdateTargetRequestMultiError(errors)
	}
	return nil
}

// UpdateTargetRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateTargetRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateTargetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTargetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTargetRequestMultiError) AllErrors() []error { return m }

// UpdateTargetRequestValidationError is the validation error returned by
// UpdateTargetRequest.Validate if the designated constraints aren't met.
type UpdateTargetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTargetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTargetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTargetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTargetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTargetRequestValidationError) ErrorName() string {
	return "UpdateTargetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTargetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTargetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTargetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTargetRequestValidationError{}

// Validate checks the field values on UpdateTargetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTargetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTargetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTargetResponseMultiError, or nil if none found.
func (m *UpdateTargetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTargetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return UpdateTargetResponseMultiError(errors)
	}
	return nil
}

// UpdateTargetResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateTargetResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateTargetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTargetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTargetResponseMultiError) AllErrors() []error { return m }

// UpdateTargetResponseValidationError is the validation error returned by
// UpdateTargetResponse.Validate if the designated constraints aren't met.
type UpdateTargetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTargetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTargetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTargetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTargetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTargetResponseValidationError) ErrorName() string {
	return "UpdateTargetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTargetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTargetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTargetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTargetResponseValidationError{}

// Validate checks the field values on DeleteTargetReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTargetReponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTargetReponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTargetReponseMultiError, or nil if none found.
func (m *DeleteTargetReponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTargetReponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return DeleteTargetReponseMultiError(errors)
	}
	return nil
}

// DeleteTargetReponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTargetReponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTargetReponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTargetReponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTargetReponseMultiError) AllErrors() []error { return m }

// DeleteTargetReponseValidationError is the validation error returned by
// DeleteTargetReponse.Validate if the designated constraints aren't met.
type DeleteTargetReponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTargetReponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTargetReponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTargetReponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTargetReponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTargetReponseValidationError) ErrorName() string {
	return "DeleteTargetReponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTargetReponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTargetReponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTargetReponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTargetReponseValidationError{}