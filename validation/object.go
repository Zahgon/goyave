package validation

// ObjectValidator the field under validation must be an object (`map[string]any`).
// If the value of the field under validation is a valid JSON string that can be unmarshalled
// into a `map[string]any`, converts the value to `map[string]any`.
type ObjectValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ObjectValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *ObjectValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *ObjectValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Object the field under validation must be an object (`map[string]any`).
	// If the value of the field under validation is a valid JSON string that can be unmarshalled
	// into a `map[string]any`, converts the value to `map[string]any`.
	return false
}

func Object() *ObjectValidator { _ = "STUB: not implemented"; return nil }
