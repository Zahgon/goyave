package validation

import (
	"goyave.dev/goyave/v5/util/walk"
)

// InValidator validates the field under validation must be a one of the given values.
type InValidator[T comparable] struct {
	BaseValidator
	Values []T
}

// Validate checks the field under validation satisfies this validator's criteria.
// Always return false if the validated value is not of type `T`.
func (v *InValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *InValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values placeholder.
	return ""
}

func (v *InValidator[T]) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// In the field under validation must be a one of the given values.
func In[T comparable](values []T) *InValidator[T] { _ = "STUB: not implemented"; return nil }

//------------------------------

// NotInValidator validates the field undervalidation must not be a one of the given values.
type NotInValidator[T comparable] struct {
	BaseValidator
	Values []T
}

// Validate checks the field under validation satisfies this validator's criteria.
// Always return false if the validated value is not of type `T`or the matched arrays
// are not of type `[]T`.
func (v *NotInValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *NotInValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values placeholder.
	return ""
}

func (v *NotInValidator[T]) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// NotIn the field under validation must not be a one of the given values.
func NotIn[T comparable](values []T) *NotInValidator[T] { _ = "STUB: not implemented"; return nil }

//------------------------------

// InFieldValidator validates the field under validation must be in at least one
// of the arrays matched by the specified path.
type InFieldValidator[T comparable] struct {
	Path *walk.Path
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
// Always return false if the validated value is not of type `T` or the matched arrays
// are not of type `[]T`.
func (v *InFieldValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *InFieldValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":other" placeholder.
	return ""
}

func (v *InFieldValidator[T]) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// InField the field under validation must be in at least one
// of the arrays matched by the specified path.
func InField[T comparable](path string) *InFieldValidator[T] { _ = "STUB: not implemented"; return nil }

//------------------------------

// NotInFieldValidator validates the field under validation must not be in any
// of the arrays matched by the specified path.
type NotInFieldValidator[T comparable] struct {
	InFieldValidator[T]
}

// Validate checks the field under validation satisfies this validator's criteria.
// Always return false if the validated value is not of type `T`.
func (v *NotInFieldValidator[T]) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *NotInFieldValidator[T]) Name() string { _ = "STUB: not implemented"; return "" }

// NotInField the field under validation must not be in any
// of the arrays matched by the specified path.
func NotInField[T comparable](path string) *NotInFieldValidator[T] {
	_ = "STUB: not implemented"
	return nil
}
