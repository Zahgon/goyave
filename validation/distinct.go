package validation

// DistinctValidator validates the field under validation must be an array having
// distinct values.
type DistinctValidator[T comparable] struct {
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DistinctValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// The array will stay `[]any` even after recursive validation if it's empty.
// We don't want to check distinct elements for empty arrays.

// Name returns the string name of the validator.
func (v *DistinctValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// Distinct the field under validation must be an array having distinct values.
	return ""
}

func Distinct[T comparable]() *DistinctValidator[T] { _ = "STUB: not implemented"; return nil }
