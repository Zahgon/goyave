package errors

import (
	"runtime"
)

// MaxStackDepth the maximum number of frames collected when creating a new Error.
var MaxStackDepth = 50

// Error wraps an errors and attaches callers at the time of creation.
// This implementation provides information for debugging or error reporting.
// It is encouraged to use this type of error everywhere a function can return
// an unexpected error. Functions returning an error that is only indicative should
// not use this error type (e.g.: user errors).
type Error struct {
	reasons      []error
	callers      []uintptr
	callerFrames FrameStack
}

// New create a new `*Error`. Collects the function callers.
//
// If the given reason is already of type `*Error`, returns it without change.
// If the reason is a slice of `error`, joins them using std's `errors.Join`.
//
// If the given reason is `nil`, returns `nil`. If the reason is `[]error`, `[]*Error` or `[]any`,
// the `nil` elements are ignored. `nil` is returned if the reason is an empty slice.
//
// If the reason is anything other than an `error`, `[]error`, `*Error`, `[]*Error`,
// `[]any`, it will be wrapped in a `Reason` structure, allowing to preserve
// its JSON marshaling behavior.
func New(reason any) error { _ = "STUB: not implemented"; return nil }

// NewSkip create a new `*Error`. Collects the function callers, skipping the given
// amount of frames.
//
// If the given reason is already of type `*Error`, returns it without change.
// If the reason is a slice of `error`, joins them using std's `errors.Join`.
//
// If the given reason is `nil`, returns `nil`. If the reason is `[]error`, `[]*Error` or `[]any`,
// the `nil` elements are ignored. `nil` is returned if the reason is an empty slice.
//
// If the reason is anything other than an `error`, `[]error`, `*Error`, `[]*Error`,
// `[]any`, it will be wrapped in a `Reason` structure, allowing to preserve
// its JSON marshaling behavior.
func NewSkip(reason any, skip int) error { _ = "STUB: not implemented"; return nil }

// Errorf is a shortcut for `errors.New(fmt.Errorf("format", args))`.
// Be careful when using this, this will result in losing the callers of
// the original error if one of the `args` is of type `*errors.Error`.
func Errorf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func toErr(reason any) []error { _ = "STUB: not implemented"; return nil }

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e Error) String() string { _ = "STUB: not implemented"; return "" }

// FileLine returns the file path and line of the error.
func (e Error) FileLine() string { _ = "STUB: not implemented"; return "" }

func (e Error) Unwrap() []error {
	_ = "STUB: not implemented"

	// Len returns the number of underlying reasons.
	return nil
}

func (e Error) Len() int { _ = "STUB: not implemented"; return 0 }

// Callers returns the function callers collected at the time of creation of the `Error`.
func (e Error) Callers() []uintptr {
	_ = "STUB: not implemented"

	// StackFrames returns the parsed `FrameStack` for this error.
	return nil
}

func (e Error) StackFrames() FrameStack { _ = "STUB: not implemented"; return *new(FrameStack) }

// MarshalJSON marshals the error and its underlying reasons.
// The result will be:
// - a string if the reasons slice is empty
// - the marshaled first reason if the reasons slice contains only one error
// - an array of the marshaled reasons otherwise
func (e Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (Error) marshalReason(e error) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// FrameStack slice of frames containing information about the stack.
// Can be used to generate a stack trace for debugging, or for error reporting.
type FrameStack []runtime.Frame

func (s FrameStack) String() string { _ = "STUB: not implemented"; return "" }

// Reason wrapper around any type of error Reason. This allows json marshaling of the Reason
// instead of losing the original data using `%v` format.
// Calling `Error()` on this structure returns the original data formatted with `%v`.
type Reason struct {
	reason any
}

// Value returns the reason's value.
func (r Reason) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (r Reason) Error() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON marshals the wrapped reason.
func (r Reason) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
