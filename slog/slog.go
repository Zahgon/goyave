package slog

import (
	"context"
	"reflect"

	"log/slog"

	"goyave.dev/goyave/v5/util/errors"
)

type unwrapper interface {
	Unwrap() []error
}

// Logger an extension of standard `*slog.Logger` overriding the `Error()` and `ErrorCtx()`
// functions so they take an error as parameter and handle `*errors.Error` gracefully.
type Logger struct {
	*slog.Logger
}

// New creates a new Logger with the given non-nil Handler and a nil context.
func New(h slog.Handler) *Logger { _ = "STUB: not implemented"; return nil }

// With returns a new Logger that includes the given arguments, converted to
// Attrs as in [Logger.Log].
// The Attrs will be added to each output from the Logger.
// The new Logger shares the old Logger's context.
// The new Logger's handler is the result of calling WithAttrs on the receiver's
// handler.
func (l *Logger) With(args ...any) *Logger { _ = "STUB: not implemented"; return nil }

// DebugWithSource logs at `LevelDebug`. The given source will be used instead of the automatically collecting it from the caller.
func (l *Logger) DebugWithSource(ctx context.Context, source uintptr, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// InfoWithSource logs at `LevelInfo`. The given source will be used instead of the automatically collecting it from the caller.
func (l *Logger) InfoWithSource(ctx context.Context, source uintptr, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// WarnWithSource logs at `LevelWarn`. The given source will be used instead of the automatically collecting it from the caller.
func (l *Logger) WarnWithSource(ctx context.Context, source uintptr, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Error logs the given error at `LevelError`.
func (l *Logger) Error(err error, args ...any) { _ = "STUB: not implemented"; return }

// ErrorCtx logs the given error at `LevelError` with the given context.
func (l *Logger) ErrorCtx(ctx context.Context, err error, args ...any) {
	_ = "STUB: not implemented"
	return
}

// ErrorWithSource logs at `LevelError`. The given source will be used instead of the automatically collecting it from the caller.
func (l *Logger) ErrorWithSource(ctx context.Context, source uintptr, err error, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) logError(ctx context.Context, source uintptr, err error, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) log(ctx context.Context, level slog.Level, source uintptr, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) makeRecord(level slog.Level, msg string, pc uintptr, args ...any) slog.Record {
	_ = "STUB: not implemented"
	return *new(slog.Record)
}

func (l *Logger) handleError(ctx context.Context, err *errors.Error, record slog.Record) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) handleReason(ctx context.Context, reason error, trace *slog.Attr, record slog.Record) {
	_ = "STUB: not implemented"
	return
}

// StructValue recursively convert a structure, structure pointer or map to a `slog.GroupValue`.
// If the given value implements `slog.LogValuer`, this value is returned instead.
// Returns AnyValue if the type is not supported.
func StructValue(v any) slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }

func structValue(v reflect.Value) slog.Value { _ = "STUB: not implemented"; return *new(slog.Value) }

// DiscardLogger returns a new Logger that discards all logs.
func DiscardLogger() *Logger { _ = "STUB: not implemented"; return nil }
