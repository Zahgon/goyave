package validation

// JSONValidator validates the field under validation must be a valid JSON string.
type JSONValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *JSONValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *JSONValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *JSONValidator) IsType() bool {
	_ = "STUB: not implemented"

	// JSON the field under validation must be a valid JSON string.
	// Unmarshals the string and sets the field value to the unmarshalled result.
	return false
}

func JSON() *JSONValidator { _ = "STUB: not implemented"; return nil }
