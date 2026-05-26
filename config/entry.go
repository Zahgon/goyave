package config

import (
	"reflect"
)

// Entry is the internal reprensentation of a config entry.
// It contains the entry value, its expected type (for validation)
// and a slice of authorized values (for validation too). If this slice
// is empty, it means any value can be used, provided it is of the correct type.
type Entry struct {
	Value            any
	AuthorizedValues []any // Leave empty for "any"
	Type             reflect.Kind
	IsSlice          bool
	Required         bool
}

func makeEntryFromValue(value any) *Entry { _ = "STUB: not implemented"; return nil }

func (e *Entry) validate(key string) error { _ = "STUB: not implemented"; return nil }

// Can't determine type, is 'zero' value.

// Accepted values for slices define the values that can be used inside the slice
// It doesn't represent the value of the slice itself (content and order)

func (e *Entry) tryConversion(kind reflect.Kind) bool { _ = "STUB: not implemented"; return false }

func convertSlice[T any](slice []any) ([]T, bool) { _ = "STUB: not implemented"; return nil, false }

func convertInt(value any) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func convertIntSlice(original []any) ([]int, bool) { _ = "STUB: not implemented"; return nil, false }

func (e *Entry) tryEnvVarConversion(key string) error { _ = "STUB: not implemented"; return nil }

func (e *Entry) convertEnvVar(str, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Keep value as string if type is not supported and let validation do its job
