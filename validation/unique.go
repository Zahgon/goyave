package validation

import (
	"math/big"
	"reflect"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Mapping of Go types to Clickhouse types can be found here:
// https://github.com/ClickHouse/clickhouse-go/blob/main/TYPES.md
// Go types uint and int are not specified, default to 'UInt64' and 'Int64', respectively
var clickhouseTypes = map[reflect.Type]string{
	reflect.TypeFor[uint64]():    "UInt64",
	reflect.TypeFor[uint32]():    "UInt32",
	reflect.TypeFor[uint16]():    "UInt16",
	reflect.TypeFor[uint8]():     "UInt8",
	reflect.TypeFor[uint]():      "UInt64",
	reflect.TypeFor[int64]():     "Int64",
	reflect.TypeFor[int32]():     "Int32",
	reflect.TypeFor[int16]():     "Int16",
	reflect.TypeFor[int8]():      "Int8",
	reflect.TypeFor[int]():       "Int64",
	reflect.TypeFor[float32]():   "Float32",
	reflect.TypeFor[float64]():   "Float64",
	reflect.TypeFor[string]():    "String",
	reflect.TypeFor[bool]():      "Bool",
	reflect.TypeFor[uuid.UUID](): "UUID",
	reflect.TypeFor[time.Time](): "DateTime64",
	reflect.TypeFor[*big.Int]():  "Int256",
}

// UniqueValidator validates the field under validation must have a unique value in database
// according to the provided database scope. Uniqueness is checked using a COUNT query.
type UniqueValidator struct {
	Scope func(db *gorm.DB, val any) *gorm.DB // TODO v6: change val to validation.Context
	BaseValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *UniqueValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *UniqueValidator) Name() string {
	_ = "STUB: not implemented"

	// Unique validates the field under validation must have a unique value in database
	// according to the provided database scope. Uniqueness is checked using a COUNT query.
	//
	//	 v.Unique(func(db *gorm.DB, val any) *gorm.DB {
	//		return db.Model(&model.User{}).Where(clause.PrimaryKey, val)
	//	 })
	//
	//	 v.Unique(func(db *gorm.DB, val any) *gorm.DB {
	//		// Unique email excluding the currently authenticated user
	//		return db.Model(&model.User{}).Where("email", val).Where("email != ?", request.User.(*model.User).Email)
	//	 })
	return ""
}

func Unique(scope func(db *gorm.DB, val any) *gorm.DB) *UniqueValidator {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// ExistsValidator validates the field under validation must exist in database
// according to the provided database scope. Existence is checked using a COUNT query.
type ExistsValidator struct {
	UniqueValidator
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ExistsValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *ExistsValidator) Name() string {
	_ = "STUB: not implemented"

	// Exists validates the field under validation must have exist database
	// according to the provided database scope. Existence is checked using a COUNT query.
	//
	//	 v.Exists(func(db *gorm.DB, val any) *gorm.DB {
	//		return db.Model(&model.User{}).Where(clause.PrimaryKey, val)
	//	 })
	return ""
}

func Exists(scope func(db *gorm.DB, val any) *gorm.DB) *ExistsValidator {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// ExistsArrayValidator validates the field under validation must be an array and all
// of its elements must exist. The type `T` is the type of the elements of the array
// under validation.
//
// This is preferable to use this validation rule on the array instead of `Exists` on
// each array element because this rule will only execute a single SQL query instead of
// as many as there are elements in the array.
//
// If provided, the `Transform` function is called on every array element to transform
// them into a raw expression. For example to transform a number into `(123::int)` for
// Postgres to prevent some type errors.
type ExistsArrayValidator[T any] struct {
	BaseValidator
	Transform func(val T) clause.Expr
	Table     string
	Column    string
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *ExistsArrayValidator[T]) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *ExistsArrayValidator[T]) buildQuery(values []T, condition bool) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *ExistsArrayValidator[T]) buildClickhouseQuery(values []T, condition bool) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *ExistsArrayValidator[T]) validate(ctx *Context, condition bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *ExistsArrayValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// ExistsArray validates the field under validation must be an array and all
	// of its elements must exist. The type `T` is the type of the elements of the array
	// under validation.
	//
	// This is preferable to use this validation rule on the array instead of `Exists` on
	// each array element because this rule will only execute a single SQL query instead of
	// as many as there are elements in the array.
	//
	// If provided, the `Transform` function is called on every array element to transform
	// them into a raw expression. For example to transform a number into `(123::int)` for
	// Postgres to prevent some type errors.
	return ""
}

func ExistsArray[T any](table, column string, transform func(val T) clause.Expr) *ExistsArrayValidator[T] {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------

// UniqueArrayValidator validates the field under validation must be an array and all
// of its elements must not already exist. The type `T` is the type of the elements of the array
// under validation.
//
// This is preferable to use this validation rule on the array instead of `Unique` on
// each array element because this rule will only execute a single SQL query instead of
// as many as there are elements in the array.
//
// If provided, the `Transform` function is called on every array element to transform
// them into a raw expression. For example to transform a number into `(123::int)` for
// Postgres to prevent some type errors.
type UniqueArrayValidator[T any] struct {
	ExistsArrayValidator[T]
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *UniqueArrayValidator[T]) Validate(ctx *Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Name returns the string name of the validator.
func (v *UniqueArrayValidator[T]) Name() string {
	_ = "STUB: not implemented"

	// UniqueArray validates the field under validation must be an array and all
	// of its elements must not already exist. The type `T` is the type of the elements of the array
	// under validation.
	//
	// This is preferable to use this validation rule on the array instead of `Unique` on
	// each array element because this rule will only execute a single SQL query instead of
	// as many as there are elements in the array.
	//
	// If provided, the `Transform` function is called on every array element to transform
	// them into a raw expression. For example to transform a number into `(123::int)` for
	// Postgres to prevent some type errors.
	return ""
}

func UniqueArray[T any](table, column string, transform func(val T) clause.Expr) *UniqueArrayValidator[T] {
	_ = "STUB: not implemented"
	return nil
}
