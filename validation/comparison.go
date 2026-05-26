package validation

import (
	"goyave.dev/goyave/v5/util/walk"
)

// ComparisonValidator validates the field under validation is greater than field identified
// by the given path. Mixed types are supported, meaning you can use this rule for the following (non-exhaustive) cases:
//   - Compare the length of two strings
//   - Compare the value of two numeric fields
//   - Compare a numeric field with the length of a string or a string length with a numeric field
//   - Compare a numeric field with the number of elements in an array
//   - Compare the number of keys in an object with a numeric field
//   - Compare a file (or multifile) size with a numeric field. The number of KiB of each file is rounded up (ceil).
type ComparisonValidator struct {
	Path *walk.Path
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ComparisonValidator) validate(ctx *Context, comparisonFunc func(size1, size2 float64) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// IsTypeDependent returns true
func (v *ComparisonValidator) IsTypeDependent() bool {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":other" placeholder.
	return false
}

func (v *ComparisonValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// GreaterThanValidator validates the field under validation is greater than the field identified
// by the given path. See `ComparisonValidator` for more details.
type GreaterThanValidator struct {
	ComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *GreaterThanValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *GreaterThanValidator) Name() string { _ = "STUB: not implemented"; return "" }

// GreaterThan validates the field under validation is greater than the field identified
// by the given path. Mixed types are supported, meaning you can use this rule for the following (non-exhaustive) cases:
//   - Compare the length of two strings
//   - Compare the value of two numeric fields
//   - Compare a numeric field with the length of a string or a string length with a numeric field
//   - Compare a numeric field with the number of elements in an array
//   - Compare the number of keys in an object with a numeric field
//   - Compare a file (or multifile) size with a numeric field. The number of KiB of each file is rounded up (ceil).
func GreaterThan(path string) *GreaterThanValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// GreaterThanEqualValidator validates the field under validation is greater than the field identified
// by the given path. See `ComparisonValidator` for more details.
type GreaterThanEqualValidator struct {
	ComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *GreaterThanEqualValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *GreaterThanEqualValidator) Name() string { _ = "STUB: not implemented"; return "" }

// GreaterThanEqual validates the field under validation is greater or equal to the field identified
// by the given path. Mixed types are supported, meaning you can use this rule for the following (non-exhaustive) cases:
//   - Compare the length of two strings
//   - Compare the value of two numeric fields
//   - Compare a numeric field with the length of a string or a string length with a numeric field
//   - Compare a numeric field with the number of elements in an array
//   - Compare the number of keys in an object with a numeric field
//   - Compare a file (or multifile) size with a numeric field. The number of KiB of each file is rounded up (ceil).
func GreaterThanEqual(path string) *GreaterThanEqualValidator {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// LowerThanValidator validates the field under validation is lower than the field identified
// by the given path. See `ComparisonValidator` for more details.
type LowerThanValidator struct {
	ComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *LowerThanValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *LowerThanValidator) Name() string {
	_ = "STUB: not implemented"

	// LowerThan validates the field under validation is lower than the field identified
	// by the given path. Mixed types are supported, meaning you can use this rule for the following (non-exhaustive) cases:
	//   - Compare the length of two strings
	//   - Compare the value of two numeric fields
	//   - Compare a numeric field with the length of a string or a string length with a numeric field
	//   - Compare a numeric field with the number of elements in an array
	//   - Compare the number of keys in an object with a numeric field
	//   - Compare a file (or multifile) size with a numeric field. The number of KiB of each file is rounded up (ceil).
	return ""
}

func LowerThan(path string) *LowerThanValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// LowerThanEqualValidator validates the field under validation is lower or equal to the field identified
// by the given path. See `ComparisonValidator` for more details.
type LowerThanEqualValidator struct {
	ComparisonValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *LowerThanEqualValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *LowerThanEqualValidator) Name() string { _ = "STUB: not implemented"; return "" }

// LowerThanEqual validates the field under validation is lower or equal to the field identified
// by the given path. Mixed types are supported, meaning you can use this rule for the following (non-exhaustive) cases:
//   - Compare the length of two strings
//   - Compare the value of two numeric fields
//   - Compare a numeric field with the length of a string or a string length with a numeric field
//   - Compare a numeric field with the number of elements in an array
//   - Compare the number of keys in an object with a numeric field
//   - Compare a file (or multifile) size with a numeric field. The number of KiB of each file is rounded up (ceil).
func LowerThanEqual(path string) *LowerThanEqualValidator { _ = "STUB: not implemented"; return nil }
