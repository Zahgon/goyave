package compress

import (
	"io"
)

// Zlib encoder for the deflate format using Go's standard `compress/zlib` package.
// Takes a compression level and "dict" ([]byte) as parameters. Accepted values are defined by constants
// in the standard `compress/zlib` package.
type Zlib struct {
	// The dictionary is optional and may be nil
	Dict  []byte
	Level int
}

// Encoding returns "deflate".
func (w *Zlib) Encoding() string {
	_ = "STUB: not implemented"

	// NewWriter returns a new `compress/zlib.Writer` using the compression level
	// defined in this Zlib encoder.
	// You may also provide a dict to compress with, or leave as nil
	return ""
}

func (w *Zlib) NewWriter(wr io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}
