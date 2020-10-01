// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/dns/v4alpha/dns_table.proto

package envoy_data_dns_v4alpha

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
var _dns_table_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DnsTable with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DnsTable) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetExternalRetryCount() > 3 {
		return DnsTableValidationError{
			field:  "ExternalRetryCount",
			reason: "value must be less than or equal to 3",
		}
	}

	for idx, item := range m.GetVirtualDomains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("VirtualDomains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetKnownSuffixes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("KnownSuffixes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DnsTableValidationError is the validation error returned by
// DnsTable.Validate if the designated constraints aren't met.
type DnsTableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTableValidationError) ErrorName() string { return "DnsTableValidationError" }

// Error satisfies the builtin error interface
func (e DnsTableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTableValidationError{}

// Validate checks the field values on DnsTable_AddressList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_AddressList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetAddress()) < 1 {
		return DnsTable_AddressListValidationError{
			field:  "Address",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetAddress() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 3 {
			return DnsTable_AddressListValidationError{
				field:  fmt.Sprintf("Address[%v]", idx),
				reason: "value length must be at least 3 runes",
			}
		}

	}

	return nil
}

// DnsTable_AddressListValidationError is the validation error returned by
// DnsTable_AddressList.Validate if the designated constraints aren't met.
type DnsTable_AddressListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_AddressListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_AddressListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_AddressListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_AddressListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_AddressListValidationError) ErrorName() string {
	return "DnsTable_AddressListValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_AddressListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_AddressList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_AddressListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_AddressListValidationError{}

// Validate checks the field values on DnsTable_DnsServiceProtocol with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsServiceProtocol) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ProtocolConfig.(type) {

	case *DnsTable_DnsServiceProtocol_Number:

		if m.GetNumber() >= 255 {
			return DnsTable_DnsServiceProtocolValidationError{
				field:  "Number",
				reason: "value must be less than 255",
			}
		}

	case *DnsTable_DnsServiceProtocol_Name:

		if utf8.RuneCountInString(m.GetName()) < 1 {
			return DnsTable_DnsServiceProtocolValidationError{
				field:  "Name",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_DnsTable_DnsServiceProtocol_Name_Pattern.MatchString(m.GetName()) {
			return DnsTable_DnsServiceProtocolValidationError{
				field:  "Name",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
		}

	default:
		return DnsTable_DnsServiceProtocolValidationError{
			field:  "ProtocolConfig",
			reason: "value is required",
		}

	}

	return nil
}

// DnsTable_DnsServiceProtocolValidationError is the validation error returned
// by DnsTable_DnsServiceProtocol.Validate if the designated constraints
// aren't met.
type DnsTable_DnsServiceProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsServiceProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsServiceProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsServiceProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsServiceProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsServiceProtocolValidationError) ErrorName() string {
	return "DnsTable_DnsServiceProtocolValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsServiceProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsServiceProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsServiceProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsServiceProtocolValidationError{}

var _DnsTable_DnsServiceProtocol_Name_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

// Validate checks the field values on DnsTable_DnsServiceTarget with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsServiceTarget) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPriority() >= 65536 {
		return DnsTable_DnsServiceTargetValidationError{
			field:  "Priority",
			reason: "value must be less than 65536",
		}
	}

	if m.GetWeight() >= 65536 {
		return DnsTable_DnsServiceTargetValidationError{
			field:  "Weight",
			reason: "value must be less than 65536",
		}
	}

	if m.GetPort() >= 65536 {
		return DnsTable_DnsServiceTargetValidationError{
			field:  "Port",
			reason: "value must be less than 65536",
		}
	}

	switch m.EndpointType.(type) {

	case *DnsTable_DnsServiceTarget_HostName:

		if utf8.RuneCountInString(m.GetHostName()) < 1 {
			return DnsTable_DnsServiceTargetValidationError{
				field:  "HostName",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_DnsTable_DnsServiceTarget_HostName_Pattern.MatchString(m.GetHostName()) {
			return DnsTable_DnsServiceTargetValidationError{
				field:  "HostName",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
		}

	case *DnsTable_DnsServiceTarget_ClusterName:

		if utf8.RuneCountInString(m.GetClusterName()) < 1 {
			return DnsTable_DnsServiceTargetValidationError{
				field:  "ClusterName",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_DnsTable_DnsServiceTarget_ClusterName_Pattern.MatchString(m.GetClusterName()) {
			return DnsTable_DnsServiceTargetValidationError{
				field:  "ClusterName",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
		}

	default:
		return DnsTable_DnsServiceTargetValidationError{
			field:  "EndpointType",
			reason: "value is required",
		}

	}

	return nil
}

// DnsTable_DnsServiceTargetValidationError is the validation error returned by
// DnsTable_DnsServiceTarget.Validate if the designated constraints aren't met.
type DnsTable_DnsServiceTargetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsServiceTargetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsServiceTargetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsServiceTargetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsServiceTargetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsServiceTargetValidationError) ErrorName() string {
	return "DnsTable_DnsServiceTargetValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsServiceTargetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsServiceTarget.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsServiceTargetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsServiceTargetValidationError{}

var _DnsTable_DnsServiceTarget_HostName_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _DnsTable_DnsServiceTarget_ClusterName_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

// Validate checks the field values on DnsTable_DnsService with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsService) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		return DnsTable_DnsServiceValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_DnsTable_DnsService_ServiceName_Pattern.MatchString(m.GetServiceName()) {
		return DnsTable_DnsServiceValidationError{
			field:  "ServiceName",
			reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
		}
	}

	if v, ok := interface{}(m.GetProtocol()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsTable_DnsServiceValidationError{
				field:  "Protocol",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetTtl(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsTable_DnsServiceValidationError{
				field:  "Ttl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(1*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return DnsTable_DnsServiceValidationError{
				field:  "Ttl",
				reason: "value must be greater than or equal to 1s",
			}
		}

	}

	if len(m.GetTargets()) < 1 {
		return DnsTable_DnsServiceValidationError{
			field:  "Targets",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTargets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsServiceValidationError{
					field:  fmt.Sprintf("Targets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DnsTable_DnsServiceValidationError is the validation error returned by
// DnsTable_DnsService.Validate if the designated constraints aren't met.
type DnsTable_DnsServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsServiceValidationError) ErrorName() string {
	return "DnsTable_DnsServiceValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsServiceValidationError{}

var _DnsTable_DnsService_ServiceName_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

// Validate checks the field values on DnsTable_DnsServiceList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsServiceList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetServices()) < 1 {
		return DnsTable_DnsServiceListValidationError{
			field:  "Services",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsServiceListValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DnsTable_DnsServiceListValidationError is the validation error returned by
// DnsTable_DnsServiceList.Validate if the designated constraints aren't met.
type DnsTable_DnsServiceListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsServiceListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsServiceListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsServiceListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsServiceListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsServiceListValidationError) ErrorName() string {
	return "DnsTable_DnsServiceListValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsServiceListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsServiceList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsServiceListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsServiceListValidationError{}

// Validate checks the field values on DnsTable_DnsEndpoint with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsEndpoint) Validate() error {
	if m == nil {
		return nil
	}

	switch m.EndpointConfig.(type) {

	case *DnsTable_DnsEndpoint_AddressList:

		if v, ok := interface{}(m.GetAddressList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsEndpointValidationError{
					field:  "AddressList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DnsTable_DnsEndpoint_ClusterName:
		// no validation rules for ClusterName

	case *DnsTable_DnsEndpoint_ServiceList:

		if v, ok := interface{}(m.GetServiceList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsEndpointValidationError{
					field:  "ServiceList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DnsTable_DnsEndpointValidationError{
			field:  "EndpointConfig",
			reason: "value is required",
		}

	}

	return nil
}

// DnsTable_DnsEndpointValidationError is the validation error returned by
// DnsTable_DnsEndpoint.Validate if the designated constraints aren't met.
type DnsTable_DnsEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsEndpointValidationError) ErrorName() string {
	return "DnsTable_DnsEndpointValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsEndpointValidationError{}

// Validate checks the field values on DnsTable_DnsVirtualDomain with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsVirtualDomain) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_DnsTable_DnsVirtualDomain_Name_Pattern.MatchString(m.GetName()) {
		return DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
		}
	}

	if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetAnswerTtl(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "AnswerTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(30*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "AnswerTtl",
				reason: "value must be greater than or equal to 30s",
			}
		}

	}

	return nil
}

// DnsTable_DnsVirtualDomainValidationError is the validation error returned by
// DnsTable_DnsVirtualDomain.Validate if the designated constraints aren't met.
type DnsTable_DnsVirtualDomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsVirtualDomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsVirtualDomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsVirtualDomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsVirtualDomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsVirtualDomainValidationError) ErrorName() string {
	return "DnsTable_DnsVirtualDomainValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsVirtualDomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsVirtualDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsVirtualDomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsVirtualDomainValidationError{}

var _DnsTable_DnsVirtualDomain_Name_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
