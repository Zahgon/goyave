package validation

// KeysInValidator the field under validation must be an object and all its keys must
// be equal to one of the given values.
type KeysInValidator struct {
	BaseValidator
	Keys []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *KeysInValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *KeysInValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values" placeholder.
	return ""
}

func (v *KeysInValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// KeysIn the field under validation must be an object and all its keys must
// be equal to one of the given values.
func KeysIn[T ~string](keys ...T) *KeysInValidator { _ = "STUB: not implemented"; return nil }
