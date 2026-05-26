package validation

import (
	"time"

	"goyave.dev/goyave/v5/util/walk"
)

// DateValidator validates the field under validation must be a string representing a date.
type DateValidator struct {
	BaseValidator
	Formats []string
}

func (v *DateValidator) parseDate(date any) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DateValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *DateValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *DateValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Date the field under validation must be a string representing a date.
	// On successful validation, converts the value to `time.Time`.
	//
	// The date must match at least one of the provided date formats (by order of preference).
	// The format uses the same syntax as Go's standard datetime format.
	// If no format is given the "2006-01-02" format is used.
	return false
}

func Date(acceptedFormats ...string) *DateValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// DateComparisonValidator factorized date comparison validator for static dates (before, after, etc.)
type DateComparisonValidator struct {
	Date time.Time
	BaseValidator
}

func (v *DateComparisonValidator) validate(ctx *Context, comparisonFunc func(time.Time, time.Time) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// MessagePlaceholders returns the ":date" placeholder.
func (v *DateComparisonValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// DateFieldComparisonValidator factorized date comparison validator for field dates (before field, after field, etc.)
type DateFieldComparisonValidator struct {
	Path *walk.Path
	BaseValidator
}

func (v *DateFieldComparisonValidator) validate(ctx *Context, comparisonFunc func(time.Time, time.Time) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Can't compare two different types or missing field

// MessagePlaceholders returns the ":date" placeholder.
func (v *DateFieldComparisonValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// DateEqualsValidator validates the field under validation must be a date (`time.Time`)
// equal to the specified date.
type DateEqualsValidator struct {
	DateComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DateEqualsValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *DateEqualsValidator) Name() string { _ = "STUB: not implemented"; return "" }

// DateEquals the field under validation must be a date (`time.Time`) equal to the given date.
func DateEquals(date time.Time) *DateEqualsValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// DateEqualsFieldValidator validates the field under validation must be a date (`time.Time`)
// equal to all the other dates matched by the specified path.
type DateEqualsFieldValidator struct {
	DateFieldComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *DateEqualsFieldValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *DateEqualsFieldValidator) Name() string { _ = "STUB: not implemented"; return "" }

// DateEqualsField the field under validation must be a date (`time.Time`) equal to the date field identified
// by the given path.
func DateEqualsField(path string) *DateEqualsFieldValidator { _ = "STUB: not implemented"; return nil }
