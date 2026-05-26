package validation

const (
	maxIntFloat32 = 16777216
	maxIntFloat64 = 9007199254740992
)

func numberAsFloat64(n any) (float64, bool, error) { _ = "STUB: not implemented"; return 0, false, nil }

type float interface {
	float32 | float64
}

type floatValidator[T float] struct{ BaseValidator }

func (v *floatValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// float32 -> float64, no check needed

func (v *floatValidator[T]) parseString(ctx *Context, val string) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *floatValidator[T]) getBitSize() int { _ = "STUB: not implemented"; return 0 }

func (v *floatValidator[T]) min() float64 { _ = "STUB: not implemented"; return 0 }

func (v *floatValidator[T]) max() float64 { _ = "STUB: not implemented"; return 0 }

func (v *floatValidator[T]) checkFloatRange(ctx *Context, val float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *floatValidator[T]) checkIntRange(ctx *Context, val int) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *floatValidator[T]) checkUintRange(ctx *Context, val uint) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *floatValidator[T]) Name() string { _ = "STUB: not implemented"; return "" }

func (v *floatValidator[T]) IsType() bool {
	_ = "STUB: not implemented"

	// Float64Validator validator for the "float64" rule.
	return false
}

type Float64Validator struct{ floatValidator[float64] }

// Float64 the field under validation must be a number
// and fit into Go's `float64` type. If the source number
// is an integer, the validator makes sure `float64` is
// capable of representing it without loss or rounding.
//
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `float64` if it passes.
func Float64() *Float64Validator { _ = "STUB: not implemented"; return nil }

// Float32Validator validator for the "float32" rule.
type Float32Validator struct{ floatValidator[float32] }

// Float32 the field under validation must be a number
// and fit into Go's `float32` type. If the source number
// is an integer, the validator makes sure `float32` is
// capable of representing it without loss or rounding.
//
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `float32` if it passes.
func Float32() *Float32Validator { _ = "STUB: not implemented"; return nil }
