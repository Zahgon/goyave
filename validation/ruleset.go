package validation

// Ruler adapter interface to make allow both RuleSet and Rules to
// be used when calling `Validate()`.
type Ruler interface {
	AsRules() Rules
}

// Validator is a Component validating a field value.
// A validator should not be re-usable or usable concurrently. They are meant to be
// scoped to a single field validation in a single request.
type Validator interface {
	Composable

	// init unexported method to force compositing with `BaseValidator`.
	init(opts *Options)

	// Init the validator with the resources required by the `Composable` interface.
	Init(opts *Options)

	// Validate checks the field under validation satisfies this validator's criteria.
	// If necessary, replaces the `Context.Value` with a converted value (see `IsType()`).
	Validate(ctx *Context) bool

	// Name returns the string name of the validator.
	// This is used to generate the language entry for the
	// validation error message.
	Name() string

	// IsTypeDependent returns true if the validator is type-dependent.
	// Type-dependent validators can be used with different field types
	// (numeric, string, arrays, objects and files) and have a different validation messages
	// depending on the type.
	// The language entry used will be "validation.rules.rulename.type"
	IsTypeDependent() bool

	// IsType returns true if the validator if a type validator.
	// A type validator checks if a field has a certain type
	// and can convert the raw value to a value fitting. For example, the UUID
	// validator is a type validator because it takes a string as input, checks if it's a
	// valid UUID and converts it to a `uuid.UUID`.
	IsType() bool

	// MessagePlaceholders returns an associative slice of placeholders and their replacement.
	// This is use to generate the validation error message. An empty slice can be returned.
	// See `lang.Language.Get()` for more details.
	MessagePlaceholders(ctx *Context) []string

	overrideMessage(langEntry string)
	getMessageOverride() string
}

// BaseValidator composable structure that implements the basic functions required to
// satisfy the `Validator` interface.
type BaseValidator struct {
	component
	messageOverride string
}

func (v *BaseValidator) init(options *Options) {
	v.component = component{
		db:     options.DB,
		config: options.Config,
		lang:   options.Language,
		logger: options.Logger,
	}
}

// Init the validator with the resources required by the `Composable` interface.
func (v *BaseValidator) Init(options *Options) {
	_ = "STUB: not implemented"

	// IsTypeDependent returns false.
	return
}

func (v *BaseValidator) IsTypeDependent() bool {
	_ = "STUB: not implemented"

	// IsType returns false.
	return false
}

func (v *BaseValidator) IsType() bool {
	_ = "STUB: not implemented"

	// MessagePlaceholders returns an empty slice (no placeholders)
	return false
}

func (v *BaseValidator) MessagePlaceholders(_ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

func (v *BaseValidator) overrideMessage(langEntry string) { _ = "STUB: not implemented"; return }

func (v *BaseValidator) getMessageOverride() string { _ = "STUB: not implemented"; return "" }

// WithMessage set a custom language entry for the error message of a Validator.
// Original placeholders returned by the validator are still used to render the message.
// Type-dependent and "element" suffixes are not added when the message is overridden.
func WithMessage[V Validator](v V, langEntry string) V { _ = "STUB: not implemented"; return *new(V) }

// FieldRulesConverter types implementing this interface define their behavior
// when converting a `FieldRules` to `Rules`. This enables rule sets composition.
type FieldRulesConverter interface {
	convert(path string, field *FieldRules, prefixDepth uint) Rules
}

// List of validators which will be applied on the field. The validators are executed in the
// order of the slice.
type List []Validator

func (l List) convert(path string, field *FieldRules, prefixDepth uint) Rules {
	_ = "STUB: not implemented"
	return *new(Rules)
}

// FieldRules structure associating a path (see `walk.Path`) identifying a field
// with a `FieldRulesApplier` (a `List` of rules or another `RuleSet` via composition).
type FieldRules struct {
	Rules FieldRulesConverter
	Path  string
}

// RuleSet definition of the validation rules applied on each field in the request.
// RuleSets are not meant to be re-used across multiple requests nor used concurrently.
type RuleSet []*FieldRules

func (r RuleSet) convert(path string, _ *FieldRules, _ uint) Rules {
	_ = "STUB: not implemented"
	return *new(Rules)
}

// AsRules converts this RuleSet to a Rules structure.
func (r RuleSet) AsRules() Rules { _ = "STUB: not implemented"; return *new(Rules) }

func (r RuleSet) asRulesWithPrefix(prefix string) Rules {
	_ = "STUB: not implemented"
	return *new(Rules)
}

// Keep a map for array fields to easily assign their element field later

// Should never be false because we injected array parents and there are no duplicates.

// injectArrayParents makes sure all array elements in the RuleSet have a parent field.
func (r RuleSet) injectArrayParents() RuleSet { _ = "STUB: not implemented"; return *new(RuleSet) }

// len(r) MUST be re-evaluated each loop, using "range r" would break it
// because the length is only evaluated once at the start of the loop.

// No parent array found, inject it

func (r Rules) checkDuplicates() { _ = "STUB: not implemented"; return }

func includeElementsKeys(paths map[string]struct{}, path string, elementField *Field) {
	_ = "STUB: not implemented"
	return
}

// Rules is the result of the transformation of RuleSet using `AsRules()`.
// It is a format that is more easily machine-readable than RuleSet.
type Rules []*Field

// AsRules returns itself.
func (r Rules) AsRules() Rules { _ = "STUB: not implemented"; return *new(Rules) }
