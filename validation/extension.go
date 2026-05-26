package validation

// ExtensionValidator validates the field under validation must be a file whose
// filename has one of the specified extensions as suffix.
// Multi-files are supported (all files must satisfy the criteria).
type ExtensionValidator struct {
	BaseValidator
	Extensions []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ExtensionValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *ExtensionValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values" placeholder.
	return ""
}

func (v *ExtensionValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// Extension the field under validation must be a file whose
// filename has one of the specified extensions as suffix.
// Don't include the dot in the extension.
// Composite extensions (e.g. "tar.gz") are supported.
//
// Multi-files are supported (all files must satisfy the criteria).
func Extension(extensions ...string) *ExtensionValidator { _ = "STUB: not implemented"; return nil }
