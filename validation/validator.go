package validation

import (
	"context"
	"reflect"
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/lang"
	"goyave.dev/goyave/v5/slog"
	"goyave.dev/goyave/v5/util/walk"
)

const (
	// CurrentElement special key for field name in composite rule sets.
	// Use it if you want to apply rules to the current object element.
	// You cannot apply rules on the root element, these rules will only
	// apply if the rule set is used with composition.
	CurrentElement = ""
)

// ExtraRequest extra key used when validating a request so the
// request's information is accessible to validation rules
type ExtraRequest struct{}

// FieldType returned by the GetFieldType function.
const (
	FieldTypeNumeric     = "numeric"
	FieldTypeString      = "string"
	FieldTypeBool        = "bool"
	FieldTypeFile        = "file"
	FieldTypeArray       = "array"
	FieldTypeObject      = "object"
	FieldTypeUnsupported = "unsupported"
)

// ErrorResponse HTTP response format for validation errors.
type ErrorResponse struct {
	Body  *Errors `json:"body,omitempty"`
	Query *Errors `json:"query,omitempty"`
}

// Composable is a partial clone of `goyave.Component`, only
// including the accessors necessary for validation.
// Validators must implement this interface so they
// have access to DB, Config, Language and Logger.
type Composable interface {
	DB() *gorm.DB
	Config() *config.Config
	Lang() *lang.Language
	Logger() *slog.Logger
}

type component struct {
	db     *gorm.DB
	config *config.Config
	lang   *lang.Language
	logger *slog.Logger
}

// DB get the database instance given through the validation Options.
// Panics if there is none.
func (c *component) DB() *gorm.DB { _ = "STUB: not implemented"; return nil }

// Config get the configuration given through the validation Options.
// Panics if there is none.
func (c *component) Config() *config.Config { _ = "STUB: not implemented"; return nil }

// Lang get the language given through the validation Options.
// Panics if there is none.
func (c *component) Lang() *lang.Language { _ = "STUB: not implemented"; return nil }

// Logger get the Logger given through the validation Options.
// Panics if there is none.
func (c *component) Logger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// Options all the parameters required by `Validate()`.
//
// Only `Data`, `Rules` and `Language` are mandatory. However, it is recommended
// to provide values for all the options in case a `Validator` requires them to function.
type Options struct {
	// Context defaults to `context.Background()` if not provided.
	Context context.Context
	Data    any
	Rules   Ruler

	Now time.Time

	// Extra can be used to store any extra information. It is passed to each `Validator`
	// via the validation `Context`.
	//
	// The keys must be comparable and should not be of type
	// string or any other built-in type to avoid collisions.
	// To avoid allocating when assigning to an `interface{}`, context keys often have
	// concrete type `struct{}`. Alternatively, exported context key variables' static
	// type should be a pointer or interface.
	Extra map[any]any

	// Language used for translating validation error messages.
	// Defaults to `lang.Default`.
	Language *lang.Language
	DB       *gorm.DB
	Config   *config.Config
	Logger   *slog.Logger

	// ConvertSingleValueArrays set to true to convert fields that are expected
	// to be an array into an array with a single value.
	//
	// It is recommended to set this option to `true` when validating url-encoded requests.
	// For example, if set to `false`:
	//  field=A         --> map[string]any{"field": "A"}
	//  field=A&field=B --> map[string]any{"field": []string{"A", "B"}}
	// If set to `true` and `field` has the `Array` rule:
	//  field=A         --> map[string]any{"field": []string{"A"}}
	//  field=A&field=B --> map[string]any{"field": []string{"A", "B"}}
	ConvertSingleValueArrays bool
}

type addedValidationErrorConstraint interface {
	string | *Errors
}

// AddedValidationError a simple association path/message or path/*Errors
// for use in `Context.AddValidationError` or `Context.AddValidationErrors`
type AddedValidationError[T addedValidationErrorConstraint] struct {
	Path  *walk.Path
	Error T
}

// Context is a structure unique per `Validator.Validate()` execution containing
// all the data required by a validator.
type Context struct {
	// Context is never nil. Defaults to `context.Background()`. Readonly.
	Context context.Context
	Data    any

	// Extra the map of Extra from the validation Options.
	Extra                 map[any]any
	Value                 any
	Parent                any
	Field                 *Field
	arrayElementErrors    []int
	addedValidationErrors []AddedValidationError[string]
	mergeErrors           []AddedValidationError[*Errors]
	fieldName             string
	Now                   time.Time

	path *walk.Path

	// The name of the field under validation
	Name string

	errors []error

	// Invalid is true if at least one validator prior to the current one didn't pass
	// on the field under validation. This field is readonly.
	Invalid bool
}

// AddError adds an error to the validation context. This is NOT supposed
// to be used when the field under validation doesn't match the rule, but rather
// when there has been an operation error (such as a database error).
func (c *Context) AddError(err ...error) { _ = "STUB: not implemented"; return }

// Skipped: runtime.Callers, NewSkip, this func

// AddArrayElementValidationErrors marks a child element to the field currently under validation
// as invalid. This is useful when a validation rule validates an array and wants to
// precisely mark which element in the array is invalid.
func (c *Context) AddArrayElementValidationErrors(index ...int) { _ = "STUB: not implemented"; return }

// ArrayElementErrors returns the indexes of the child eelements to the field currently under validation
// that were marked as invalid with `AddArrayElementValidationErrors`.
func (c *Context) ArrayElementErrors() []int { _ = "STUB: not implemented"; return nil }

// AddValidationError add a validation error message at the given path.
// The path is relative to the root element.
//
// This can be used when a validation rule uses nested validation or needs to add
// a message on another field than the one this validator is targeted at.
func (c *Context) AddValidationError(path *walk.Path, message string) {
	_ = "STUB: not implemented"
	return
}

// AddedValidationError returns the additional errors added with `AddValidationError`.
func (c *Context) AddedValidationError() []AddedValidationError[string] {
	_ = "STUB: not implemented"
	return nil
}

// AddValidationErrors add a `*Errors` to be merged into the errors bag of the current
// validation. The path is relative to the root element.
//
// This can be used when a validation rule uses nested validation needs to merge
// the results into the higher-level validation errors.
//
// See `*validation.Errors.Merge` for more details.
func (c *Context) AddValidationErrors(path *walk.Path, errors *Errors) {
	_ = "STUB: not implemented"
	return
}

// AddedValidationErrors returns the additional errors added with `AddValidationErrors`.
func (c *Context) AddedValidationErrors() []AddedValidationError[*Errors] {
	_ = "STUB: not implemented"
	return nil

	// Path returns the exact Path to the current element.
	// The path is relative to the root element. If you are compositing rule sets in your validation,
	// the path returned is NOT relative to the root of the current rule set.
	//
	// You can use this path to inject validation errors using AddValidationError and MergeValidationErrors.
}

func (c *Context) Path() *walk.Path {
	_ = "STUB: not implemented"

	// Errors returns this validation context's errors.
	// The errors returned are NOT validation errors but operation errors (such as database error).
	// Because each rule on each field has its own Context, the returned array will only contain
	// errors related to the current field and the current rule.
	return nil
}

func (c *Context) Errors() []error { _ = "STUB: not implemented"; return nil }

type validator struct {
	validationErrors *Errors
	options          *Options
	now              time.Time
	errors           []error
}

// Validate the given data using the given `Options`.
// If all validation rules pass and no error occurred, the first returned value will be `nil`.
//
// The second returned value is a slice of error that occurred during validation. These
// errors are not validation errors but error raised when a validator could not be executed correctly.
// For example if a validator using the database generated a DB error.
//
// The `Options.Data` may be modified thanks to type rules.
func Validate(options *Options) (*Errors, []error) { _ = "STUB: not implemented"; return nil, nil }

// Validate the root element

func (v *validator) validateField(fieldName string, field *Field, walkData any, parentPath *walk.Path) {
	_ = "STUB: not implemented"
	return
}

// This is an array, validate its elements first so it can be converted to correct type

// We can use `First` here because the path contains array indexes
// so we are sure there will be only one match.

// Value may be modified (converting rule), replace it in the parent element

func (v *validator) isRootElement(fieldName string, errorPath *walk.Path) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *validator) shouldDeleteFromParent(field *Field, parentIsObject bool, value any) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *validator) shouldConvertSingleValueArray(fieldName string) bool {
	_ = "STUB: not implemented"
	// It's OK not to check for escape characters for the brackets here since the fieldName is the
	// escaped representation of the field name: e.g.: "escapedArray\[\]" and "escapedArray\[\][]"
	return false
}

func (v *validator) convertSingleValueArray(field *Field, value any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *validator) isAbsent(field *Field, c *walk.Context, parentPath *walk.Path, data any) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *validator) processAddedErrors(ctx *Context, parentPath *walk.Path, c *walk.Context, validator Validator) {
	_ = "STUB: not implemented"
	return
}

func (v *validator) getLangEntry(ctx *Context, validator Validator) string {
	_ = "STUB: not implemented"
	return ""
}

func (v *validator) processPlaceholders(ctx *Context, translatedFieldName string, validator Validator) []string {
	_ = "STUB: not implemented"
	return nil
}

func (v *validator) getMessage(ctx *Context, translatedFieldName string, validator Validator) string {
	_ = "STUB: not implemented"
	return ""
}

// findTypeValidator find the expected type of a field for a given array dimension.
func (v *validator) findTypeValidator(validators []Validator) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

func replaceValue(value any, c *walk.Context) { _ = "STUB: not implemented"; return }

// Parent is slice

func makeGenericSlice(original any) ([]any, bool) { _ = "STUB: not implemented"; return nil, false }

func appendPath(parentPath, childPath *walk.Path, index int) *walk.Path {
	_ = "STUB: not implemented"
	return nil
}

// GetFieldType returns the non-technical type of the given "value" interface.
// This is used by validation rules to know if the input data is a candidate
// for validation or not and is especially useful for type-dependent rules.
//   - "numeric" (`lang.FieldTypeNumeric`) if the value is an int, uint or a float
//   - "string" (`lang.FieldTypeString`) if the value is a string
//   - "array" (`lang.FieldTypeArray`) if the value is a slice
//   - "file" (`lang.FieldTypeFile`) if the value is a slice of "fsutil.File"
//   - "bool" (`lang.FieldTypeBool`) if the value is a bool
//   - "unsupported" (`lang.FieldTypeUnsupported`) otherwise
func GetFieldType(value any) string { _ = "STUB: not implemented"; return "" }

func getFieldType(value reflect.Value) string { _ = "STUB: not implemented"; return "" }

// GetFieldName returns the localized name of the field identified
// by the given path.
func GetFieldName(lang *lang.Language, path *walk.Path) string {
	_ = "STUB: not implemented"
	return ""
}

func translateFieldName(lang *lang.Language, fieldName string) string {
	_ = "STUB: not implemented"
	return ""
}

func lastUnescapedDot(str string) int { _ = "STUB: not implemented"; return 0 }
