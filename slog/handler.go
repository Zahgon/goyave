package slog

import (
	"bytes"
	"context"
	"io"
	"sync"

	"log/slog"
)

// Colors and formats
const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Gray      = "\033[90m"
	WhiteBold = "\033[37;1m"
	GrayBold  = "\033[90;1m"
	BGYellow  = "\033[43m"
	BGRed     = "\033[41m"
	BGCyan    = "\033[46m"
	BGGray    = "\033[100m"
)

var (
	// Indent the string used to indent attribute groups in the dev mode handler.
	Indent = "  "
)

// DevModeHandlerOptions options for the dev mode handler.
type DevModeHandlerOptions struct {
	// Level reports the minimum record level that will be logged.
	// The handler discards records with lower levels.
	// If Level is nil, the handler assumes `LevelInfo`.
	// The handler calls `Level.Level()` for each record processed;
	// to adjust the minimum level dynamically, use a `slog.LevelVar`.
	Level slog.Leveler
}

// DevModeHandler is a `slog.Handler` that writes Records to an io.Writer.
// The records are formatted to be easily readable by humans.
// This handler is meant for development use only as it doesn't provide optimal
// performance and its output is not machine-readable.
type DevModeHandler struct {
	opts   *DevModeHandlerOptions
	mu     *sync.Mutex
	w      io.Writer
	attrs  []slog.Attr
	groups []string
}

// NewHandler creates a new `slog.Handler` with default options.
// If `devMode` is true, a `*DevModeHandler` is returned, else a `*slog.JSONHandler`.
func NewHandler(devMode bool, w io.Writer) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// NewDevModeHandler creates a new `DevModeHandler` that writes to w, using the given options.
// If `opts` is `nil`, the default options are used.
func NewDevModeHandler(w io.Writer, opts *DevModeHandlerOptions) *DevModeHandler {
	_ = "STUB: not implemented"
	return nil
}

// Handle formats its argument `Record` in an output easily readable by humans.
// The output contains multiple lines:
//   - The first one contains the log level, the time and the source
//   - The second one contains the message
//   - The next lines contain the attributes and groups, if any
//
// Each call to `Handle` results in a single serialized call to `io.Writer.Write()`.
func (h *DevModeHandler) Handle(_ context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// Change color depending on level

// levelColor return a color for the tag describing the level in the output.
// We use ranges so custom levels can be supported.
func levelColor(level slog.Level) string { _ = "STUB: not implemented"; return "" }

// Debug

// Info

// Warn

// Error

func messageColor(level slog.Level) string { _ = "STUB: not implemented"; return "" }

// Debug and Info

// Warn

// Error

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *DevModeHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

// WithAttrs returns a new `DevModeHandler` whose attributes consists
// of h's attributes followed by attrs.
func (h *DevModeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// WithGroup returns a new `DevModeHandler` whose attributes are wrapped
// into a named group. All the handler's attributes will be printed indented
// into the added group.
func (h *DevModeHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func printAttr(attr slog.Attr, buf *bytes.Buffer, indent int) { _ = "STUB: not implemented"; return }

// This may be a struct or map, convert it if needed

// Break line if the message is multi-line (such as stacktrace)
// Otherwise print it next to attr name so the log is more compact

func printGroup(group []slog.Attr, buf *bytes.Buffer, indent int) {
	_ = "STUB: not implemented"
	return
}
