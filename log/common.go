package log

import (
	"log/slog"
)

const (
	// Format is the fmt format code for common logs
	Format string = "%s %s %s [%s] \"%s %s %s\" %d %d"

	// TimestampFormat is the time format used in common logs
	TimestampFormat string = "02/Jan/2006:15:04:05 -0700"
)

// CommonLogFormatter build a log entry using the Common Log Format.
func CommonLogFormatter(ctx *Context) (string, []slog.Attr) {
	_ = "STUB: not implemented"
	return "", nil
}

// Requests using the CONNECT method over HTTP/2.0 must use
// the authority field (aka r.Host) to identify the target.
// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT

// CombinedLogFormatter build a log entry using the Combined Log Format.
func CombinedLogFormatter(ctx *Context) (string, []slog.Attr) {
	_ = "STUB: not implemented"
	return "", nil
}
