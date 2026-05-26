package compress

import (
	"io"
)

// Brotli encoder for the br compression format
type Brotli struct {
	// Quality controls the compression-speed vs compression-density trade-offs.
	// The higher the quality, the slower the compression. Range is 0 to 11.
	Quality int
	// LGWin is the base 2 logarithm of the sliding window size.
	// Range is 10 to 24. 0 indicates automatic configuration based on Quality.
	LGWin int
}

// Encoding returns "br".
func (w *Brotli) Encoding() string {
	_ = "STUB: not implemented"

	// NewWriter returns a new `brotli.Writer` using the
	// Compression Quality and LGWin provided in the Brotli encoder
	return ""
}

func (w *Brotli) NewWriter(wr io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}
