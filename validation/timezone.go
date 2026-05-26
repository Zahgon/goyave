package validation

import (
	"sync"
)

var timezoneCache = sync.Map{}

// TimezoneValidator the field under validation must be a valid string
// reprensentation of a timezone.
// If validation passes, the value is converted to `*time.Location`
// using `time.LoadLocation()`.
// "Local" as an input is not accepted as a valid timezone.
//
// As `time.LoadLocation()` can be a slow operation, timezones are cached.
type TimezoneValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *TimezoneValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *TimezoneValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *TimezoneValidator) IsType() bool {
	_ = "STUB: not implemented"

	// Timezone the field under validation must be a valid string reprensentation of a timezone.
	// If validation passes, the value is converted to `*time.Location` using `time.LoadLocation()`.
	// "Local" as an input is not accepted as a valid timezone.
	//
	// As `time.LoadLocation()` can be a slow operation, timezones are cached.
	return false
}

func Timezone() *TimezoneValidator { _ = "STUB: not implemented"; return nil }
