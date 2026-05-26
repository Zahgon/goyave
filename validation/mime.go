package validation

// MIMEValidator validates the field under validation must be a file
// and match one of the given MIME types.
// Multi-files are supported (all files must satisfy the criteria).
type MIMEValidator struct {
	BaseValidator
	MIMETypes []string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *MIMEValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Ignore MIME settings (example: "text/plain; charset=utf-8")

// Name returns the string name of the validator.
func (v *MIMEValidator) Name() string {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns the ":values" placeholder.
	return ""
}

func (v *MIMEValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// MIME the field under validation must be a file and match one of the given
// MIME types. Multi-files are supported (all files must satisfy the criteria).
func MIME(mimeTypes ...string) *MIMEValidator { _ = "STUB: not implemented"; return nil }

//------------------------------

// ImageMIMETypes MIME types accepted by `ImageValidator`.
var ImageMIMETypes = []string{"image/jpeg", "image/png", "image/gif", "image/bmp", "image/svg+xml", "image/webp"}

// ImageValidator validates the field under validation must be an image file.
// Multi-files are supported (all files must satisfy the criteria).
type ImageValidator struct {
	MIMEValidator
}

// Name returns the string name of the validator.
func (v *ImageValidator) Name() string {
	_ = "STUB: not implemented"

	// Image the field under validation must be an image file.
	// Multi-files are supported (all files must satisfy the criteria).
	//
	// Accepted MIME types are defined by `ImageMIMETypes`.
	return ""
}

func Image() *ImageValidator { _ = "STUB: not implemented"; return nil }
