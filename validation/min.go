package validation

// MinValidator validates the field under validation depending on its type.
//   - Numbers are directly compared if they fit in `float64`. If they don't the rule doesn't pass.
//   - Strings must have a length of at least n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have at least n elements
//   - Objects must have at least n keys
//   - Files must weight at least n KiB (for multi-files, all files must match this criteria). The number of KiB of each file is rounded up (ceil).
type MinValidator struct {
	BaseValidator
	Min float64
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *MinValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *MinValidator) Name() string {
	_ = "STUB: not implemented"

	// IsTypeDependent returns true
	return ""
}

func (v *MinValidator) IsTypeDependent() bool {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":min" placeholder.
	return false
}

func (v *MinValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// Min validates the field under validation depending on its type.
//   - Numbers are directly compared if they fit in `float64`. If they don't the rule doesn't pass.
//   - Strings must have a length of at least n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have at least n elements
//   - Objects must have at least n keys
//   - Files must weight at least n KiB (for multi-files, all files must match this criteria). The number of KiB of each file is rounded up (ceil).
func Min(min float64) *MinValidator { _ = "STUB: not implemented"; return nil }
