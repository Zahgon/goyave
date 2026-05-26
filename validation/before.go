package validation

import (
	"time"
)

// BeforeValidator validates the field under validation must be a date (`time.Time`) before
// the specified date.
type BeforeValidator struct {
	DateComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *BeforeValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *BeforeValidator) Name() string {
	_ = "STUB: not implemented"

	// Before the field under validation must be a date (`time.Time`) before the given date.
	return ""
}

func Before(date time.Time) *BeforeValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// BeforeEqualValidator validates the field under validation must be a date (`time.Time`) before
// or equal to the specified date.
type BeforeEqualValidator struct {
	DateComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *BeforeEqualValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *BeforeEqualValidator) Name() string { _ = "STUB: not implemented"; return "" }

// BeforeEqual the field under validation must be a date (`time.Time`) before or equal to the given date.
func BeforeEqual(date time.Time) *BeforeEqualValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// BeforeFieldValidator validates the field under validation must be a date (`time.Time`) before
// all the other dates matched by the specified path.
type BeforeFieldValidator struct {
	DateFieldComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *BeforeFieldValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *BeforeFieldValidator) Name() string {
	_ = "STUB: not implemented"

	// BeforeField the field under validation must be a date (`time.Time`) before the date field identified
	// by the given path.
	return ""
}

func BeforeField(path string) *BeforeFieldValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// BeforeEqualFieldValidator validates the field under validation must be a date (`time.Time`) before
// or equal to all the other dates matched by the specified path.
type BeforeEqualFieldValidator struct {
	DateFieldComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *BeforeEqualFieldValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *BeforeEqualFieldValidator) Name() string { _ = "STUB: not implemented"; return "" }

// BeforeEqualField the field under validation must be a date (`time.Time`) before or equal to the date field identified
// by the given path.
func BeforeEqualField(path string) *BeforeEqualFieldValidator {
	_ = "STUB: not implemented"
	return nil
}
