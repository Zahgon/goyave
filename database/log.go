package database

import (
	"context"
	"regexp"
	"time"

	"gorm.io/gorm/logger"
	"goyave.dev/goyave/v5/slog"
)

var regexGormPath = regexp.MustCompile(`gorm.io/(.*?)@`)

// Logger adapter between `*slog.Logger` and GORM's logger.
type Logger struct {
	slogger func() *slog.Logger

	// SlowThreshold defines the minimum query execution time to be considered "slow".
	// If a query takes more time than `SlowThreshold`, the query will be logged at the WARN level.
	// If 0, disables query execution time checking.
	SlowThreshold time.Duration
}

// NewLogger create a new `Logger` adapter between GORM and `*slog.Logger`.
// Use a `SlowThreshold` of 200ms.
func NewLogger(slogger func() *slog.Logger) *Logger { _ = "STUB: not implemented"; return nil }

// LogMode returns a copy of this logger. The level argument actually has
// no effect as it is handled by the underlying `*slog.Logger`.
func (l *Logger) LogMode(_ logger.LogLevel) logger.Interface {
	_ = "STUB: not implemented"
	return *new(logger.Interface)
}

// Info logs at `LevelInfo`.
func (l Logger) Info(ctx context.Context, msg string, data ...any) {
	_ = "STUB: not implemented"
	return
}

// Warn logs at `LevelWarn`.
func (l Logger) Warn(ctx context.Context, msg string, data ...any) {
	_ = "STUB: not implemented"
	return
}

// Error logs at `LevelError`.
func (l Logger) Error(ctx context.Context, msg string, data ...any) {
	_ = "STUB: not implemented"
	return
}

// Trace SQL logs at
//   - `LevelDebug`
//   - `LevelWarn` if the query is slow
//   - `LevelError` if the given error is not nil
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	_ = "STUB: not implemented"
	return
}

func getSourceCaller() uintptr {
	_ = "STUB: not implemented"
	// function copied from gorm/utils/utils.go
	// the second caller usually from gorm internal, so set i start from 2
	return 0
}
