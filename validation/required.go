package validation

// RequiredValidator the field under validation is required.
// If a field is absent from the input data, subsequent validators
// will not be executed.
//
// If a field is `nil` and has the `Nullable` validator, this validator passes.
// As non-nullable fields are removed if they have a `nil` value, this validator
// doesn't pass if a field is `nil` and doesn't have the `Nullable` validator.
type RequiredValidator struct{ BaseValidator }

// Validate checks the field under validation satisfies this validator's criteria.
func (v *RequiredValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *RequiredValidator) Name() string {
	_ = "STUB: not implemented"

	// Required the field under validation is required.
	// If a field is absent from the input data, subsequent validators
	// will not be executed.
	//
	// If a field is `nil` and has the `Nullable` validator, this validator passes.
	// As non-nullable fields are removed if they have a `nil` value, this validator
	// doesn't pass if a field is `nil` and doesn't have the `Nullable` validator.
	return ""
}

func Required() *RequiredValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// RequiredIfValidator is the same as `RequiredValidator` but only applies the behavior
// described if the specified `Condition` function returns true.
type RequiredIfValidator struct {
	Condition func(*Context) bool
	RequiredValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *RequiredIfValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// RequiredIf is the same as `Required` but only applies the behavior
// described if the specified condition function returns true.
func RequiredIf(condition func(*Context) bool) *RequiredIfValidator {
	_ = "STUB: not implemented"
	return nil
}
