// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_packages.proto

package api

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _applicationserver_packages_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ApplicationPackage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationPackage) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) > 36 {
		return ApplicationPackageValidationError{
			field:  "Name",
			reason: "value length must be at most 36 runes",
		}
	}

	if !_ApplicationPackage_Name_Pattern.MatchString(m.GetName()) {
		return ApplicationPackageValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
		}
	}

	if val := m.GetDefaultFPort(); val < 1 || val > 255 {
		return ApplicationPackageValidationError{
			field:  "DefaultFPort",
			reason: "value must be inside range [1, 255]",
		}
	}

	return nil
}

// ApplicationPackageValidationError is the validation error returned by
// ApplicationPackage.Validate if the designated constraints aren't met.
type ApplicationPackageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageValidationError) ErrorName() string {
	return "ApplicationPackageValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageValidationError{}

var _ApplicationPackage_Name_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// Validate checks the field values on ApplicationPackages with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationPackages) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPackages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationPackagesValidationError{
					field:  fmt.Sprintf("Packages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ApplicationPackagesValidationError is the validation error returned by
// ApplicationPackages.Validate if the designated constraints aren't met.
type ApplicationPackagesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackagesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackagesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackagesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackagesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackagesValidationError) ErrorName() string {
	return "ApplicationPackagesValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackagesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackages.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackagesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackagesValidationError{}

// Validate checks the field values on ApplicationPackageAssociationIdentifiers
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ApplicationPackageAssociationIdentifiers) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEndDeviceIds() == nil {
		return ApplicationPackageAssociationIdentifiersValidationError{
			field:  "EndDeviceIds",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEndDeviceIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageAssociationIdentifiersValidationError{
				field:  "EndDeviceIds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if val := m.GetFPort(); val < 1 || val > 255 {
		return ApplicationPackageAssociationIdentifiersValidationError{
			field:  "FPort",
			reason: "value must be inside range [1, 255]",
		}
	}

	return nil
}

// ApplicationPackageAssociationIdentifiersValidationError is the validation
// error returned by ApplicationPackageAssociationIdentifiers.Validate if the
// designated constraints aren't met.
type ApplicationPackageAssociationIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageAssociationIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageAssociationIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageAssociationIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageAssociationIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageAssociationIdentifiersValidationError) ErrorName() string {
	return "ApplicationPackageAssociationIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageAssociationIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageAssociationIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageAssociationIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageAssociationIdentifiersValidationError{}

// Validate checks the field values on ApplicationPackageAssociation with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationPackageAssociation) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return ApplicationPackageAssociationValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageAssociationValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageAssociationValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageAssociationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetPackageName()) > 36 {
		return ApplicationPackageAssociationValidationError{
			field:  "PackageName",
			reason: "value length must be at most 36 runes",
		}
	}

	if !_ApplicationPackageAssociation_PackageName_Pattern.MatchString(m.GetPackageName()) {
		return ApplicationPackageAssociationValidationError{
			field:  "PackageName",
			reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
		}
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageAssociationValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ApplicationPackageAssociationValidationError is the validation error
// returned by ApplicationPackageAssociation.Validate if the designated
// constraints aren't met.
type ApplicationPackageAssociationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageAssociationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageAssociationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageAssociationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageAssociationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageAssociationValidationError) ErrorName() string {
	return "ApplicationPackageAssociationValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageAssociationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageAssociation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageAssociationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageAssociationValidationError{}

var _ApplicationPackageAssociation_PackageName_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// Validate checks the field values on ApplicationPackageAssociations with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationPackageAssociations) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAssociations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationPackageAssociationsValidationError{
					field:  fmt.Sprintf("Associations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ApplicationPackageAssociationsValidationError is the validation error
// returned by ApplicationPackageAssociations.Validate if the designated
// constraints aren't met.
type ApplicationPackageAssociationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageAssociationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageAssociationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageAssociationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageAssociationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageAssociationsValidationError) ErrorName() string {
	return "ApplicationPackageAssociationsValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageAssociationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageAssociations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageAssociationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageAssociationsValidationError{}

// Validate checks the field values on GetApplicationPackageAssociationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GetApplicationPackageAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return GetApplicationPackageAssociationRequestValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApplicationPackageAssociationRequestValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApplicationPackageAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetApplicationPackageAssociationRequestValidationError is the validation
// error returned by GetApplicationPackageAssociationRequest.Validate if the
// designated constraints aren't met.
type GetApplicationPackageAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplicationPackageAssociationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplicationPackageAssociationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApplicationPackageAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplicationPackageAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplicationPackageAssociationRequestValidationError) ErrorName() string {
	return "GetApplicationPackageAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApplicationPackageAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplicationPackageAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplicationPackageAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplicationPackageAssociationRequestValidationError{}

// Validate checks the field values on ListApplicationPackageAssociationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ListApplicationPackageAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return ListApplicationPackageAssociationRequestValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListApplicationPackageAssociationRequestValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetLimit() > 1000 {
		return ListApplicationPackageAssociationRequestValidationError{
			field:  "Limit",
			reason: "value must be less than or equal to 1000",
		}
	}

	// no validation rules for Page

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListApplicationPackageAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListApplicationPackageAssociationRequestValidationError is the validation
// error returned by ListApplicationPackageAssociationRequest.Validate if the
// designated constraints aren't met.
type ListApplicationPackageAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplicationPackageAssociationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplicationPackageAssociationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplicationPackageAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplicationPackageAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplicationPackageAssociationRequestValidationError) ErrorName() string {
	return "ListApplicationPackageAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplicationPackageAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplicationPackageAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplicationPackageAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplicationPackageAssociationRequestValidationError{}

// Validate checks the field values on SetApplicationPackageAssociationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *SetApplicationPackageAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAssociation() == nil {
		return SetApplicationPackageAssociationRequestValidationError{
			field:  "Association",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAssociation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetApplicationPackageAssociationRequestValidationError{
				field:  "Association",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetApplicationPackageAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SetApplicationPackageAssociationRequestValidationError is the validation
// error returned by SetApplicationPackageAssociationRequest.Validate if the
// designated constraints aren't met.
type SetApplicationPackageAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetApplicationPackageAssociationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetApplicationPackageAssociationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetApplicationPackageAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetApplicationPackageAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetApplicationPackageAssociationRequestValidationError) ErrorName() string {
	return "SetApplicationPackageAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetApplicationPackageAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetApplicationPackageAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetApplicationPackageAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetApplicationPackageAssociationRequestValidationError{}

// Validate checks the field values on
// ApplicationPackageDefaultAssociationIdentifiers with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ApplicationPackageDefaultAssociationIdentifiers) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetApplicationIds() == nil {
		return ApplicationPackageDefaultAssociationIdentifiersValidationError{
			field:  "ApplicationIds",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetApplicationIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageDefaultAssociationIdentifiersValidationError{
				field:  "ApplicationIds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if val := m.GetFPort(); val < 1 || val > 255 {
		return ApplicationPackageDefaultAssociationIdentifiersValidationError{
			field:  "FPort",
			reason: "value must be inside range [1, 255]",
		}
	}

	return nil
}

// ApplicationPackageDefaultAssociationIdentifiersValidationError is the
// validation error returned by
// ApplicationPackageDefaultAssociationIdentifiers.Validate if the designated
// constraints aren't met.
type ApplicationPackageDefaultAssociationIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) ErrorName() string {
	return "ApplicationPackageDefaultAssociationIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageDefaultAssociationIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageDefaultAssociationIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageDefaultAssociationIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageDefaultAssociationIdentifiersValidationError{}

// Validate checks the field values on ApplicationPackageDefaultAssociation
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ApplicationPackageDefaultAssociation) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return ApplicationPackageDefaultAssociationValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageDefaultAssociationValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageDefaultAssociationValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageDefaultAssociationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetPackageName()) > 36 {
		return ApplicationPackageDefaultAssociationValidationError{
			field:  "PackageName",
			reason: "value length must be at most 36 runes",
		}
	}

	if !_ApplicationPackageDefaultAssociation_PackageName_Pattern.MatchString(m.GetPackageName()) {
		return ApplicationPackageDefaultAssociationValidationError{
			field:  "PackageName",
			reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
		}
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationPackageDefaultAssociationValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ApplicationPackageDefaultAssociationValidationError is the validation error
// returned by ApplicationPackageDefaultAssociation.Validate if the designated
// constraints aren't met.
type ApplicationPackageDefaultAssociationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageDefaultAssociationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageDefaultAssociationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageDefaultAssociationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageDefaultAssociationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageDefaultAssociationValidationError) ErrorName() string {
	return "ApplicationPackageDefaultAssociationValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageDefaultAssociationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageDefaultAssociation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageDefaultAssociationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageDefaultAssociationValidationError{}

var _ApplicationPackageDefaultAssociation_PackageName_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// Validate checks the field values on ApplicationPackageDefaultAssociations
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ApplicationPackageDefaultAssociations) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDefaults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationPackageDefaultAssociationsValidationError{
					field:  fmt.Sprintf("Defaults[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ApplicationPackageDefaultAssociationsValidationError is the validation error
// returned by ApplicationPackageDefaultAssociations.Validate if the
// designated constraints aren't met.
type ApplicationPackageDefaultAssociationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationPackageDefaultAssociationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationPackageDefaultAssociationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationPackageDefaultAssociationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationPackageDefaultAssociationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationPackageDefaultAssociationsValidationError) ErrorName() string {
	return "ApplicationPackageDefaultAssociationsValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationPackageDefaultAssociationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationPackageDefaultAssociations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationPackageDefaultAssociationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationPackageDefaultAssociationsValidationError{}

// Validate checks the field values on
// GetApplicationPackageDefaultAssociationRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetApplicationPackageDefaultAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return GetApplicationPackageDefaultAssociationRequestValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApplicationPackageDefaultAssociationRequestValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApplicationPackageDefaultAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetApplicationPackageDefaultAssociationRequestValidationError is the
// validation error returned by
// GetApplicationPackageDefaultAssociationRequest.Validate if the designated
// constraints aren't met.
type GetApplicationPackageDefaultAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplicationPackageDefaultAssociationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplicationPackageDefaultAssociationRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GetApplicationPackageDefaultAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplicationPackageDefaultAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplicationPackageDefaultAssociationRequestValidationError) ErrorName() string {
	return "GetApplicationPackageDefaultAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApplicationPackageDefaultAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplicationPackageDefaultAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplicationPackageDefaultAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplicationPackageDefaultAssociationRequestValidationError{}

// Validate checks the field values on
// ListApplicationPackageDefaultAssociationRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListApplicationPackageDefaultAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIds() == nil {
		return ListApplicationPackageDefaultAssociationRequestValidationError{
			field:  "Ids",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListApplicationPackageDefaultAssociationRequestValidationError{
				field:  "Ids",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetLimit() > 1000 {
		return ListApplicationPackageDefaultAssociationRequestValidationError{
			field:  "Limit",
			reason: "value must be less than or equal to 1000",
		}
	}

	// no validation rules for Page

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListApplicationPackageDefaultAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListApplicationPackageDefaultAssociationRequestValidationError is the
// validation error returned by
// ListApplicationPackageDefaultAssociationRequest.Validate if the designated
// constraints aren't met.
type ListApplicationPackageDefaultAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplicationPackageDefaultAssociationRequestValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e ListApplicationPackageDefaultAssociationRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e ListApplicationPackageDefaultAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplicationPackageDefaultAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplicationPackageDefaultAssociationRequestValidationError) ErrorName() string {
	return "ListApplicationPackageDefaultAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplicationPackageDefaultAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplicationPackageDefaultAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplicationPackageDefaultAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplicationPackageDefaultAssociationRequestValidationError{}

// Validate checks the field values on
// SetApplicationPackageDefaultAssociationRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SetApplicationPackageDefaultAssociationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDefault() == nil {
		return SetApplicationPackageDefaultAssociationRequestValidationError{
			field:  "Default",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDefault()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetApplicationPackageDefaultAssociationRequestValidationError{
				field:  "Default",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetApplicationPackageDefaultAssociationRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SetApplicationPackageDefaultAssociationRequestValidationError is the
// validation error returned by
// SetApplicationPackageDefaultAssociationRequest.Validate if the designated
// constraints aren't met.
type SetApplicationPackageDefaultAssociationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetApplicationPackageDefaultAssociationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetApplicationPackageDefaultAssociationRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e SetApplicationPackageDefaultAssociationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetApplicationPackageDefaultAssociationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetApplicationPackageDefaultAssociationRequestValidationError) ErrorName() string {
	return "SetApplicationPackageDefaultAssociationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetApplicationPackageDefaultAssociationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetApplicationPackageDefaultAssociationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetApplicationPackageDefaultAssociationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetApplicationPackageDefaultAssociationRequestValidationError{}
