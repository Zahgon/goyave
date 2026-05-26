package validation

import (
	"reflect"
)

// ArrayValidator validates the field under validation must be an array.
type ArrayValidator struct {
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ArrayValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *ArrayValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *ArrayValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Array the field under validation must be an array.
	// On successful validation and if possible, converts the array to its correct type
	// based on its elements' type. If all elements have the same type, the array is converted to
	// a slice of this type.
	return false
}

func Array() *ArrayValidator { _ = "STUB: not implemented"; return nil }

// convertArray to its correct type based on its elements' type.
// If all elements have the same type, the array is converted to
// a slice of this type, otherwise the array is returned as-is.
func convertArray(array any, parentType reflect.Type) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// The first element is probably `nil`, avoid "call of reflect.Value.Interface on zero Value" error.

// Not all elements have the same type, keep it []any
