package validation

import "regexp"

var digitsRegex = regexp.MustCompile(`^[0-9]*$`)

// DigitsValidator the field under validation must be a string that
// only contains digits.
type DigitsValidator struct {
	RegexValidator
}

// Name returns the string name of the validator.
func (v *DigitsValidator) Name() string {
	_ = "STUB: not implemented"

	// Digits the field under validation must be a string that only contains digits.
	return ""
}

func Digits() *DigitsValidator { _ = "STUB: not implemented"; return nil }
