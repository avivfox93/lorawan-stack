// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: lorawan-stack/api/cluster.proto

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
var _cluster_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PeerInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PeerInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GrpcPort

	// no validation rules for Tls

	// no validation rules for Tags

	return nil
}

// PeerInfoValidationError is the validation error returned by
// PeerInfo.Validate if the designated constraints aren't met.
type PeerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerInfoValidationError) ErrorName() string { return "PeerInfoValidationError" }

// Error satisfies the builtin error interface
func (e PeerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerInfoValidationError{}
