package typeutil

// Convert anything into the desired type using JSON marshaling and unmarshaling.
func Convert[T any](data any) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// MustConvert anything into the desired type using JSON marshaling and unmarshaling.
// Panics if it fails.
func MustConvert[T any](data any) T { _ = "STUB: not implemented"; return *new(T) }

// Copy deep-copy a DTO's non-zero fields to the given model. The model is updated in-place and returned.
// Field names are matched in a case sensitive way.
// If you need to copy a zero-value (empty string, `false`, 0, etc) into the destination model, your DTO
// can take advantage of `typeutil.Undefined`.
// Panics if an error occurs.
func Copy[T, D any](model *T, dto D) *T { _ = "STUB: not implemented"; return nil }
