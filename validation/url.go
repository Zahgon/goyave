package validation

// URLValidator the field under validation must be a string representing
// a valid URL as per `url.ParseRequestURI()`.
// If validation passes, the value is converted to `*url.URL`.
type URLValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *URLValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *URLValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *URLValidator) IsType() bool {
	_ = "STUB: not implemented"

	// URL the field under validation must be a representing
	// a valid URL as per `url.ParseRequestURI()`.
	// If validation passes, the value is converted to `*url.URL`.
	return false
}

func URL() *URLValidator { _ = "STUB: not implemented"; return nil }
