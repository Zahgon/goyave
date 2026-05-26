package validation

// EmailValidator the field under validation must be a string that can be parsed
// using Go's standard `mail.ParseAddress` function.
//
// The email address format is defined by RFC 5322. For example:
//   - Barry Gibbs <bg@example.com>
//   - foo@example.com
//
// This validator is not enough in itself to properly validate an email address.
// The only way to ensure an email address is valid is by sending a confirmation email.
//
// On successful validation, converts the value to `string`.
type EmailValidator struct {
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *EmailValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// IsType returns true.
func (v *EmailValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Name returns the string name of the validator.
	return false
}

func (v *EmailValidator) Name() string {
	_ = "STUB: not implemented"

	// Email the field under validation must be a string that can be parsed using Go's standard
	// `mail.ParseAddress` function.
	//
	// The email address format is defined by RFC 5322. For example:
	//   - Barry Gibbs <bg@example.com>
	//   - foo@example.com
	//
	// This validator is not enough in itself to properly validate an email address.
	// The only way to ensure an email address is valid is by sending a confirmation email.
	//
	// On successful validation, converts the value to `string`.
	return ""
}

func Email() *EmailValidator { _ = "STUB: not implemented"; return nil }
