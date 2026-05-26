package typeutil

import (
	"database/sql/driver"
)

// Undefined utility type wrapping a generic value used to differentiate
// between the absence of a field and its zero value, without using pointers.
//
// This is especially useful when using wrappers such as `sql.NullString`, which
// are structures that encode/decode to a non-struct value. When working with
// requests that may or may not contain a field that is a nullable value, you cannot
// use pointers to define the presence or absence of this kind of structure. Thus the
// case where the field is absent (zero-value) and where the field is present but has
// a null value are indistinguishable.
//
// This type implements:
//   - `encoding.TextUnmarshaler`
//   - `json.Unmarshaler`
//   - `json.Marshaler`
//   - `driver.Valuer`
//   - `sql.Scanner`
//
// It is recommended to use the json tag `omitzero` on struct fields of type `Undefined`
// to properly handle JSON marshaling, model mapping and DTO conversion.
//
// This type can be used in response DTOs or for scanning database results. This is useful when
// you don't always select all fields from the model and you don't want the unselected fields to
// show in the response.
type Undefined[T any] struct {
	Val     T
	Present bool
}

// NewUndefined creates a new `Undefined` wrapper with `Present` set to `true`.
func NewUndefined[T any](val T) Undefined[T] { _ = "STUB: not implemented"; return nil }

// Set the value. Automatically sets `Present` to `true`.
func (u *Undefined[T]) Set(value T) { _ = "STUB: not implemented"; return }

// Unset the value (reset to zero-value) and set `Present` to `false`.
// This effectively works like if the value was entirely removed from the struct.
func (u *Undefined[T]) Unset() { _ = "STUB: not implemented"; return }

// UnmarshalJSON implements json.Unmarshaler.
// On successful unmarshal of the underlying value, sets the `Present` field to `true`.
func (u *Undefined[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// Only the value is marshaled, even if the field is not present.
// Therefore, it is recommended to use the json tag `omitzero`.
func (u Undefined[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// If the input is a blank string, `Present` is set to `false`, otherwise `true`.
// This implementation will return an error if the underlying value doesn't implement
// `encoding.TextUnmarshaler`.
func (u *Undefined[T]) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// IsZero returns true for non-present values.
func (u Undefined[T]) IsZero() bool {
	_ = "STUB: not implemented"

	// IsPresent returns true for present values.
	return false
}

func (u Undefined[T]) IsPresent() bool {
	_ = "STUB: not implemented"

	// Value implements the `driver.Valuer` interface.
	return false
}

func (u Undefined[T]) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan implements the `sql.Scanner` interface.
//
// When called, always set `Present` to `true`.
//
// If `src` is of type `Undefined`, `*Undefined`, `T` or `*T`, the value is copied
// directly. In case of nil pointer, T's zero value is used instead.
//
// If the generic type T implements `sql.Scanner` and `src`'s type doesn't match any of
// the above, T's `Scan()` method will be used.
//
// This implementation is also useful in the case of model mapping with `typeutil.Copy`.
func (u *Undefined[T]) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// CopyValue implements the copier.Valuer interface.
func (u Undefined[T]) CopyValue() any { _ = "STUB: not implemented"; return *new(any) }

// Default return the value if present, otherwise returns the given default value.
func (u Undefined[T]) Default(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }
