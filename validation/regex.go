package validation

import "regexp"

// RegexValidator the field under validation must be a string matching
// the specified `*regexp.Regexp`.
type RegexValidator struct {
	Regexp *regexp.Regexp
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *RegexValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *RegexValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":regexp" placeholder.
	return ""
}

func (v *RegexValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// Regex the field under validation must be a string matching
// the specified `*regexp.Regexp`.
func Regex(regex *regexp.Regexp) *RegexValidator { _ = "STUB: not implemented"; return nil }
