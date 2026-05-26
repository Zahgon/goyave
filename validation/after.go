package validation

import (
	"time"
)

// AfterValidator validates the field under validation must be a date (`time.Time`) before
// the specified date.
type AfterValidator struct {
	DateComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *AfterValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *AfterValidator) Name() string {
	_ = "STUB: not implemented"

	// After the field under validation must be a date (`time.Time`) before the given date.
	return ""
}

func After(date time.Time) *AfterValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// AfterEqualValidator validates the field under validation must be a date (`time.Time`) after
// or equal to the specified date.
type AfterEqualValidator struct {
	DateComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *AfterEqualValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *AfterEqualValidator) Name() string { _ = "STUB: not implemented"; return "" }

// AfterEqual the field under validation must be a date (`time.Time`) after or equal to the given date.
func AfterEqual(date time.Time) *AfterEqualValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// AfterFieldValidator validates the field under validation must be a date (`time.Time`) before
// all the other dates matched by the specified path.
type AfterFieldValidator struct {
	DateFieldComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *AfterFieldValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *AfterFieldValidator) Name() string {
	_ = "STUB: not implemented"

	// AfterField the field under validation must be a date (`time.Time`) before the date field identified
	// by the given path.
	return ""
}

func AfterField(path string) *AfterFieldValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// AfterEqualFieldValidator validates the field under validation must be a date (`time.Time`) after
// or equal to all the other dates matched by the specified path.
type AfterEqualFieldValidator struct {
	DateFieldComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *AfterEqualFieldValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *AfterEqualFieldValidator) Name() string { _ = "STUB: not implemented"; return "" }

// AfterEqualField the field under validation must be a date (`time.Time`) after or equal to the date field identified
// by the given path.
func AfterEqualField(path string) *AfterEqualFieldValidator { _ = "STUB: not implemented"; return nil }
