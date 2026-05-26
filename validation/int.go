package validation

type integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type intValidator[T integer] struct{ BaseValidator }

func (v *intValidator[T]) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (v *intValidator[T]) isUnsigned() bool { _ = "STUB: not implemented"; return false }

func (v *intValidator[T]) getBitSize() int { _ = "STUB: not implemented"; return 0 }

func (v *intValidator[T]) checkFloat64Range(ctx *Context, val float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *intValidator[T]) checkFloat32Range(ctx *Context, val float32) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *intValidator[T]) checkIntRange(ctx *Context, val int) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *intValidator[T]) checkUintRange(ctx *Context, val uint) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *intValidator[T]) min() int { _ = "STUB: not implemented"; return 0 }

func (v *intValidator[T]) max() uint { _ = "STUB: not implemented"; return 0 }

func (v *intValidator[T]) parseString(ctx *Context, val string) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *intValidator[T]) Name() string { _ = "STUB: not implemented"; return "" }

// IsType returns true
func (v *intValidator[T]) IsType() bool {
	_ = "STUB: not implemented"

	// IntValidator validator for the "int" rule.
	return false
}

type IntValidator struct{ intValidator[int] }

// Int the field under validation must be an integer
// and fit into Go's `int` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `int` if it passes.
func Int() *IntValidator { _ = "STUB: not implemented"; return nil }

// Int8Validator validator for the "int8" rule.
type Int8Validator struct{ intValidator[int8] }

// Int8 the field under validation must be an integer
// and fit into Go's `int8` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `int8` if it passes.
func Int8() *Int8Validator { _ = "STUB: not implemented"; return nil }

// Int16Validator validator for the "int16" rule.
type Int16Validator struct{ intValidator[int16] }

// Int16 the field under validation must be an integer
// and fit into Go's `int16` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `int16` if it passes.
func Int16() *Int16Validator { _ = "STUB: not implemented"; return nil }

// Int32Validator validator for the "int32" rule.
type Int32Validator struct{ intValidator[int32] }

// Int32 the field under validation must be an integer
// and fit into Go's `int32` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `int32` if it passes.
func Int32() *Int32Validator { _ = "STUB: not implemented"; return nil }

// Int64Validator validator for the "int64" rule.
type Int64Validator struct{ intValidator[int64] }

// Int64 the field under validation must be an integer
// and fit into Go's `int64` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `int64` if it passes.
func Int64() *Int64Validator { _ = "STUB: not implemented"; return nil }

// UintValidator validator for the "uint" rule.
type UintValidator struct{ intValidator[uint] }

// Uint the field under validation must be a positive integer
// and fit into Go's `uint` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `uint` if it passes.
func Uint() *UintValidator { _ = "STUB: not implemented"; return nil }

// Uint8Validator validator for the "uint8" rule.
type Uint8Validator struct{ intValidator[uint8] }

// Uint8 the field under validation must be a positive integer
// and fit into Go's `uint8` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `uint8` if it passes.
func Uint8() *Uint8Validator { _ = "STUB: not implemented"; return nil }

// Uint16Validator validator for the "uint16" rule.
type Uint16Validator struct{ intValidator[uint16] }

// Uint16 the field under validation must be a positive integer
// and fit into Go's `uint16` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `uint16` if it passes.
func Uint16() *Uint16Validator { _ = "STUB: not implemented"; return nil }

// Uint32Validator validator for the "uint32" rule.
type Uint32Validator struct{ intValidator[uint32] }

// Uint32 the field under validation must be a positive integer
// and fit into Go's `uint32` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `uint32` if it passes.
func Uint32() *Uint32Validator { _ = "STUB: not implemented"; return nil }

// Uint64Validator validator for the "uint64" rule.
type Uint64Validator struct{ intValidator[uint64] }

// Uint64 the field under validation must be a positive integer
// and fit into Go's `uint64` type. If the source number is
// a float, the validator makes sure the value is within
// the range of integers that the float can accurately represent.
//
// Floats are only accepted if they don't have a decimal.
// Strings that can be converted to the target type are accepted.
// This rule converts the field to `uint64` if it passes.
func Uint64() *Uint64Validator { _ = "STUB: not implemented"; return nil }
