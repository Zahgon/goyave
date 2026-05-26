package validation

func validateSize(value any, v func(size int) bool) bool { _ = "STUB: not implemented"; return false }

// Pass if field type cannot be checked (bool, dates, ...)

// SizeValidator validates the field under validation depending on its type.
//   - Strings must have a length of n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have n elements
//   - Objects must have n keys
//   - Files must weight n KiB (for multi-files, all files must match this criteria). The number of KiB is rounded up (ceil).
type SizeValidator struct {
	BaseValidator
	Size int
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *SizeValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *SizeValidator) Name() string {
	_ = "STUB: not implemented"

	// IsTypeDependent returns true
	return ""
}

func (v *SizeValidator) IsTypeDependent() bool {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":value" placeholder.
	return false
}

func (v *SizeValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// Size validates the field under validation depending on its type.
//   - Strings must have a length of n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have n elements
//   - Objects must have n keys
//   - Files must weight n KiB (for multi-files, all files must match this criteria). The number of KiB is rounded up (ceil).
func Size(size int) *SizeValidator { _ = "STUB: not implemented"; return nil }
