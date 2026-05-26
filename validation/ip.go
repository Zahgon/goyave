package validation

// IPValidator the field under validation must be a string representing
// a valid IPv4 or IPv6.
// If validation passes, the value is converted to `net.IP`.
type IPValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *IPValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *IPValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *IPValidator) IsType() bool {
	_ = "STUB: not implemented"

	// IP the field under validation must be a string representing
	// a valid IPv4 or IPv6.
	// If validation passes, the value is converted to `net.IP`.
	return false
}

func IP() *IPValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// IPv4Validator the field under validation must be a string representing
// a valid IPv4.
// If validation passes, the value is converted to `net.IP`.
type IPv4Validator struct{ IPValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *IPv4Validator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *IPv4Validator) Name() string {
	_ = "STUB: not implemented"

	// IPv4 the field under validation must be a string representing a valid IPv4.
	// If validation passes, the value is converted to `net.IP`.
	return ""
}

func IPv4() *IPv4Validator { _ = "STUB: not implemented"; return nil }

//------------------------------

// IPv6Validator the field under validation must be a string representing
// a valid IPv6.
// If validation passes, the value is converted to `net.IP`.
type IPv6Validator struct{ IPValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *IPv6Validator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *IPv6Validator) Name() string {
	_ = "STUB: not implemented"

	// IPv6 the field under validation must be a string representing a valid IPv6.
	// If validation passes, the value is converted to `net.IP`.
	return ""
}

func IPv6() *IPv6Validator { _ = "STUB: not implemented"; return nil }
