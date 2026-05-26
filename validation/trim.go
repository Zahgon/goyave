package validation

// TrimValidator if the field under validation is a string, trims it using
// `strings.TrimSpace()`.
type TrimValidator struct{ BaseValidator }

// Validate always returns true. If the field under validation is a string,
// trims it using `strings.TrimSpace()`.
func (v *TrimValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// This rule is just transforming, so we always return true.

// Name returns the string name of the validator.
func (v *TrimValidator) Name() string {
	_ = "STUB: not implemented"

	// Trim if the field under validation is a string, trims it using `strings.TrimSpace()`.
	return ""
}

func Trim() *TrimValidator { _ = "STUB: not implemented"; return nil }
