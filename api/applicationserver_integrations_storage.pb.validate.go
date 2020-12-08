// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

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
var _applicationserver_integrations_storage_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetStoredApplicationUpRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetStoredApplicationUpRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetApplicationIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "ApplicationIds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEndDeviceIds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "EndDeviceIds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _GetStoredApplicationUpRequest_Type_InLookup[m.GetType()]; !ok {
		return GetStoredApplicationUpRequestValidationError{
			field:  "Type",
			reason: "value must be in list [ uplink_message join_accept downlink_ack downlink_nack downlink_sent downlink_failed downlink_queued downlink_queue_invalidated location_solved service_data]",
		}
	}

	if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAfter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "After",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "Before",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoredApplicationUpRequestValidationError{
				field:  "FPort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _GetStoredApplicationUpRequest_Order_InLookup[m.GetOrder()]; !ok {
		return GetStoredApplicationUpRequestValidationError{
			field:  "Order",
			reason: "value must be in list [ -received_at received_at]",
		}
	}

	return nil
}

// GetStoredApplicationUpRequestValidationError is the validation error
// returned by GetStoredApplicationUpRequest.Validate if the designated
// constraints aren't met.
type GetStoredApplicationUpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStoredApplicationUpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStoredApplicationUpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStoredApplicationUpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStoredApplicationUpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStoredApplicationUpRequestValidationError) ErrorName() string {
	return "GetStoredApplicationUpRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStoredApplicationUpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStoredApplicationUpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStoredApplicationUpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStoredApplicationUpRequestValidationError{}

var _GetStoredApplicationUpRequest_Type_InLookup = map[string]struct{}{
	"":                           {},
	"uplink_message":             {},
	"join_accept":                {},
	"downlink_ack":               {},
	"downlink_nack":              {},
	"downlink_sent":              {},
	"downlink_failed":            {},
	"downlink_queued":            {},
	"downlink_queue_invalidated": {},
	"location_solved":            {},
	"service_data":               {},
}

var _GetStoredApplicationUpRequest_Order_InLookup = map[string]struct{}{
	"":             {},
	"-received_at": {},
	"received_at":  {},
}
