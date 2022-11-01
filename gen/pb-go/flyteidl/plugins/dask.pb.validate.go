// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/plugins/dask.proto

package plugins

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
var _dask_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DaskJob with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DaskJob) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Namespace

	if v, ok := interface{}(m.GetJobPodSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DaskJobValidationError{
				field:  "JobPodSpec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DaskJobValidationError{
				field:  "Cluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DaskJobValidationError is the validation error returned by DaskJob.Validate
// if the designated constraints aren't met.
type DaskJobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DaskJobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DaskJobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DaskJobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DaskJobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DaskJobValidationError) ErrorName() string { return "DaskJobValidationError" }

// Error satisfies the builtin error interface
func (e DaskJobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDaskJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DaskJobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DaskJobValidationError{}

// Validate checks the field values on JobPodSpec with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JobPodSpec) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Image

	if v, ok := interface{}(m.GetResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JobPodSpecValidationError{
				field:  "Resources",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// JobPodSpecValidationError is the validation error returned by
// JobPodSpec.Validate if the designated constraints aren't met.
type JobPodSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobPodSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobPodSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobPodSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobPodSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobPodSpecValidationError) ErrorName() string { return "JobPodSpecValidationError" }

// Error satisfies the builtin error interface
func (e JobPodSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobPodSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobPodSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobPodSpecValidationError{}

// Validate checks the field values on DaskCluster with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DaskCluster) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Image

	// no validation rules for NWorkers

	if v, ok := interface{}(m.GetResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DaskClusterValidationError{
				field:  "Resources",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAdditionalWorkerGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DaskClusterValidationError{
					field:  fmt.Sprintf("AdditionalWorkerGroups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetAutoscaler()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DaskClusterValidationError{
				field:  "Autoscaler",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DaskClusterValidationError is the validation error returned by
// DaskCluster.Validate if the designated constraints aren't met.
type DaskClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DaskClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DaskClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DaskClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DaskClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DaskClusterValidationError) ErrorName() string { return "DaskClusterValidationError" }

// Error satisfies the builtin error interface
func (e DaskClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDaskCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DaskClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DaskClusterValidationError{}

// Validate checks the field values on DaskAutoscaler with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DaskAutoscaler) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Minimum

	// no validation rules for Maximum

	return nil
}

// DaskAutoscalerValidationError is the validation error returned by
// DaskAutoscaler.Validate if the designated constraints aren't met.
type DaskAutoscalerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DaskAutoscalerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DaskAutoscalerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DaskAutoscalerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DaskAutoscalerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DaskAutoscalerValidationError) ErrorName() string { return "DaskAutoscalerValidationError" }

// Error satisfies the builtin error interface
func (e DaskAutoscalerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDaskAutoscaler.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DaskAutoscalerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DaskAutoscalerValidationError{}

// Validate checks the field values on DaskWorkerGroup with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DaskWorkerGroup) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Image

	// no validation rules for NWorkers

	if v, ok := interface{}(m.GetResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DaskWorkerGroupValidationError{
				field:  "Resources",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DaskWorkerGroupValidationError is the validation error returned by
// DaskWorkerGroup.Validate if the designated constraints aren't met.
type DaskWorkerGroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DaskWorkerGroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DaskWorkerGroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DaskWorkerGroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DaskWorkerGroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DaskWorkerGroupValidationError) ErrorName() string { return "DaskWorkerGroupValidationError" }

// Error satisfies the builtin error interface
func (e DaskWorkerGroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDaskWorkerGroup.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DaskWorkerGroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DaskWorkerGroupValidationError{}
