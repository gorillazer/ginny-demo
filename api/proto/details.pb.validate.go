// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: details.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 999 {
		return GetReqValidationError{
			field:  "Id",
			reason: "value must be greater than 999",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 3 {
		return GetReqValidationError{
			field:  "Name",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// GetReqValidationError is the validation error returned by GetReq.Validate if
// the designated constraints aren't met.
type GetReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReqValidationError) ErrorName() string { return "GetReqValidationError" }

// Error satisfies the builtin error interface
func (e GetReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReqValidationError{}

// Validate checks the field values on GetRes with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Price

	if v, ok := interface{}(m.GetCreatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResValidationError{
				field:  "CreatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetResValidationError is the validation error returned by GetRes.Validate if
// the designated constraints aren't met.
type GetResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResValidationError) ErrorName() string { return "GetResValidationError" }

// Error satisfies the builtin error interface
func (e GetResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResValidationError{}
