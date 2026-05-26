package compress

import (
	"io"
)

// Gzip encoder for the gzip format using Go's standard `compress/gzip` package.
//
// Takes a compression level as parameter. Accepted values are defined by constants
// in the standard `compress/gzip` package.
type Gzip struct {
	Level int
}

// Encoding returns "gzip".
func (w *Gzip) Encoding() string {
	_ = "STUB: not implemented"

	// NewWriter returns a new `compress/gzip.Writer` using the compression level
	// defined in this Gzip encoder.
	return ""
}

func (w *Gzip) NewWriter(wr io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}
