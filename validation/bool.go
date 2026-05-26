package validation

// BoolValidator the field under validation must be a bool or one of the following values:
//   - "1" / "0"
//   - "true" / "false"
//   - "yes" / "no"
//   - "on" / "off"
//   - a number different from 0 is converetd to `true`, a number equals to 0 is converted to `false`
//
// This rule converts the field to `bool` if it passes.
type BoolValidator struct {
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *BoolValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *BoolValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true
	return ""
}

func (v *BoolValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Bool the field under validation must be a bool or one of the following values:
	//   - "1" / "0"
	//   - "true" / "false"
	//   - "yes" / "no"
	//   - "on" / "off"
	//   - a number different from 0 is converetd to `true`, a number equals to 0 is converted to `false`
	//
	// This rule converts the field to `bool` if it passes.
	return false
}

func Bool() *BoolValidator { _ = "STUB: not implemented"; return nil }
