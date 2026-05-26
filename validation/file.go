package validation

// FileValidator validates the field under validation must be a file.
// Multi-files are supported.
type FileValidator struct {
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *FileValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *FileValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *FileValidator) IsType() bool {
	_ = "STUB: not implemented"

	// File the field under validation must be a file. Multi-files are supported.
	return false
}

func File() *FileValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// FileCountValidator validates the field under validation must be a multi-files
// with exactly the specified number of files.
type FileCountValidator struct {
	BaseValidator
	Count uint
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *FileCountValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *FileCountValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":value" placeholder.
	return ""
}

func (v *FileCountValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// FileCount the field under validation must be a multi-files
// with exactly the specified number of files.
func FileCount(count uint) *FileCountValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// MinFileCountValidator validates the field under validation must be a multi-files
// with at least the specified number of files.
type MinFileCountValidator struct {
	BaseValidator
	Min uint
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *MinFileCountValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *MinFileCountValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":min" placeholder.
func (v *MinFileCountValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// MinFileCount the field under validation must be a multi-files
// with at least the specified number of files.
func MinFileCount(min uint) *MinFileCountValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// MaxFileCountValidator validates the field under validation must be a multi-files
// with at most the specified number of files.
type MaxFileCountValidator struct {
	BaseValidator
	Max uint
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *MaxFileCountValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *MaxFileCountValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":max" placeholder.
func (v *MaxFileCountValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// MaxFileCount the field under validation must be a multi-files
// with at most the specified number of files.
func MaxFileCount(max uint) *MaxFileCountValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// FileCountBetweenValidator validates the field under validation must be a multi-files
// with a number of files between the specified min and max.
type FileCountBetweenValidator struct {
	BaseValidator
	Min uint
	Max uint
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *FileCountBetweenValidator) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *FileCountBetweenValidator) Name() string { _ = "STUB: not implemented"; return "" }

// MessagePlaceholders returns the ":min" and ":max" placeholders.
func (v *FileCountBetweenValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// FileCountBetween the field under validation must be a multi-files
// with a number of files between the specified min and max.
func FileCountBetween(min, max uint) *FileCountBetweenValidator {
	_ = "STUB: not implemented"
	return nil
}
