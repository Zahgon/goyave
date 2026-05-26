package validation

import (
	"goyave.dev/goyave/v5/util/walk"
)

// Field representation of a single field in the data being validated.
// Provides useful information based on its validators (if required, nullable, etc).
type Field struct {
	isRequired func(*Context) bool

	Path       *walk.Path
	Elements   *Field
	Validators []Validator

	// prefixDepth When using composition, `prefixDepth` allows to truncate the path to the
	// validated element in order to retrieve the root object or array relative to
	// the composed RuleSet.
	prefixDepth uint

	isArray    bool
	isObject   bool
	isNullable bool
}

func alwaysRequired(_ *Context) bool { _ = "STUB: not implemented"; return false }

func newField(path string, validators []Validator, prefixDepth uint) *Field {
	_ = "STUB: not implemented"
	return nil
}

// getErrorPath returns the path to use when appending the error message to the
// final validation errors.
//
// The given `parentPath` corresponds to the path to the parent array
// if the parent is an array, otherwise `nil`. If `nil`, returns the unmodified
// path from the `walk.Context`.
func (f *Field) getErrorPath(parentPath *walk.Path, c *walk.Context) *walk.Path {
	_ = "STUB: not implemented"
	return nil
}

// IsRequired check if a field has the "required" rule
func (f *Field) IsRequired(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// IsNullable check if a field has the "nullable" rule
func (f *Field) IsNullable() bool { _ = "STUB: not implemented"; return false }

// IsArray check if a field has the "array" rule
func (f *Field) IsArray() bool {
	_ = "STUB: not implemented"

	// IsObject check if a field has the "object" rule
	return false
}

func (f *Field) IsObject() bool {
	_ = "STUB: not implemented"

	// PrefixDepth When using composition, `prefixDepth` allows to truncate the path to the
	// validated element in order to retrieve the root object or array relative to
	// the composed RuleSet.
	return false
}

func (f *Field) PrefixDepth() uint { _ = "STUB: not implemented"; return 0 }
