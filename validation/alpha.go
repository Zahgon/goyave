package validation

import "regexp"

var (
	alphaRegex     = regexp.MustCompile(`^[\pL\pM]+$`)
	alphaNumRegex  = regexp.MustCompile(`^[\pL\pM0-9]+$`)
	alphaDashRegex = regexp.MustCompile(`^[\pL\pM0-9_-]+$`)
)

// AlphaValidator the field under validation must be an alphabetic string.
type AlphaValidator struct {
	RegexValidator
}

// Name returns the string name of the validator.
func (v *AlphaValidator) Name() string {
	_ = "STUB: not implemented"

	// Alpha the field under validation must be an alphabetic string.
	return ""
}

func Alpha() *AlphaValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// AlphaNumValidator the field under validation must an alphabetic-numeric string.
type AlphaNumValidator struct {
	RegexValidator
}

// Name returns the string name of the validator.
func (v *AlphaNumValidator) Name() string {
	_ = "STUB: not implemented"

	// AlphaNum the field under validation must an alphabetic-numeric string.
	return ""
}

func AlphaNum() *AlphaNumValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// AlphaDashValidator the field under validation must a string made
// of alphabetic-numeric characters, dashes or underscores.
type AlphaDashValidator struct {
	RegexValidator
}

// Name returns the string name of the validator.
func (v *AlphaDashValidator) Name() string {
	_ = "STUB: not implemented"

	// AlphaDash the field under validation must be a string made
	// of alphabetic-numeric characters, dashes or underscores.
	return ""
}

func AlphaDash() *AlphaDashValidator { _ = "STUB: not implemented"; return nil }
