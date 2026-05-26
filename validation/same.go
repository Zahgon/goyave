package validation

import (
	"goyave.dev/goyave/v5/util/walk"
)

// SameValidator validates the field under validation is strictly equal to the field identified
// by the given path. Values of different types are never equal. Files are not checked and will never pass this validator.
// For arrays, objects and numbers, the values are compared using `reflect.DeepEqual()`.
// For numbers, make sure the two compared numbers have the same type. A `uint` with value `1` will be considered
// different from an `int` with value `1`.
type SameValidator struct {
	Path *walk.Path
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *SameValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// We cannot validate this field

// Name returns the string name of the validator.
func (v *SameValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":other" placeholder.
	return ""
}

func (v *SameValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// Same validates the field under validation is strictly equal to the field identified
// by the given path. Values of different types are never equal. Files are not checked
// and will never pass this validator.
// For arrays, objects and numbers, the values are compared using `reflect.DeepEqual()`.
// For numbers, make sure the two compared numbers have the same type. A `uint` with value `1` will be considered
// different from an `int` with value `1`.
func Same(path string) *SameValidator { _ = "STUB: not implemented"; return nil }
