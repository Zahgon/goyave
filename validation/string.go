package validation

// StringValidator the field under validation must be a string.
type StringValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *StringValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *StringValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *StringValidator) IsType() bool {
	_ = "STUB: not implemented"

	// String the field under validation must be a string.
	return false
}

func String() *StringValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// StartsWithValidator the field under validation must be a string starting
// with at least one of the specified prefixes.
type StartsWithValidator struct {
	BaseValidator
	Prefix []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *StartsWithValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *StartsWithValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":values" placeholder.
func (v *StartsWithValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// StartsWith the field under validation must be a string starting
// with at least one of the specified prefixes.
func StartsWith(prefix ...string) *StartsWithValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// EndsWithValidator the field under validation must be a string ending
// with at least one of the specified suffixes.
type EndsWithValidator struct {
	BaseValidator
	Suffix []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *EndsWithValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *EndsWithValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values" placeholder.
	return ""
}

func (v *EndsWithValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// EndsWith the field under validation must be a string ending
// with at least one of the specified prefixes.
func EndsWith(suffix ...string) *EndsWithValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// DoesntStartWithValidator the field under validation must be a string not starting
// with any of the specified prefixes.
type DoesntStartWithValidator struct {
	BaseValidator
	Prefix []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DoesntStartWithValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *DoesntStartWithValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":values" placeholder.
func (v *DoesntStartWithValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// DoesntStartWith the field under validation must be a string not starting
// with any of the specified prefixes.
func DoesntStartWith(prefix ...string) *DoesntStartWithValidator {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// DoesntEndWithValidator the field under validation must be a string not ending
// with any of the specified prefixes.
type DoesntEndWithValidator struct {
	BaseValidator
	Suffix []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DoesntEndWithValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *DoesntEndWithValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":values" placeholder.
func (v *DoesntEndWithValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// DoesntEndWith the field under validation must be a string not ending
// with any of the specified prefixes.
func DoesntEndWith(suffix ...string) *DoesntEndWithValidator { _ = "STUB: not implemented"; return nil }
