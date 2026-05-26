package compress

import (
	"io"
	"net/http"

	"goyave.dev/goyave/v5"
)

// Encoder is an interface that wraps the methods returning the information
// necessary for the compress middleware to work.
//
// `NewWriter` returns any `io.WriteCloser`, allowing the middleware to support
// any compression algorithm.
//
// `Encoding` returns the name of the compression algorithm. Using the returned value,
// the middleware:
//  1. detects the client's preferred encoding with the `Accept-Encoding` request header
//  2. replaces the response writer with the writer returned by `NewWriter`
//  3. sets the `Content-Encoding` response header
type Encoder interface {
	NewWriter(wr io.Writer) io.WriteCloser
	Encoding() string
}

type resettable interface {
	Reset(w io.Writer)
}

type compressWriter struct {
	goyave.CommonWriter
	responseWriter http.ResponseWriter
	childWriter    io.Writer
	encoding       string
	empty          bool
}

func (w *compressWriter) PreWrite(b []byte) { _ = "STUB: not implemented"; return }

func (w *compressWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *compressWriter) Close() error {
	_ = "STUB: not implemented"

	// Do not write gzip/br/... footer if nothing has been written to the response body.
	return nil
}

// Middleware compresses HTTP responses.
//
// This middleware supports multiple algorithms thanks to the `Encoders` slice.
// The encoder will be chosen depending on the request's `Accept-Encoding` header,
// and the value returned by the `Encoder`'s `Encoding()` method. Quality values in
// the headers are taken into account.
//
// In case of equal priority, the encoding that is the earliest in the slice is chosen.
// If the header's value is `*` and no encoding already matched,
// the first element of the slice is used.
//
// If none of the accepted encodings are available in the `Encoders` slice, then the
// response will not be compressed and the middleware immediately passes.
//
// If the middleware successfully replaces the response writer, the `Accept-Encoding`
// header is removed from the request to avoid potential clashes with potential other
// encoding middleware.
//
// If not set at the first call of `Write()`, the middleware will automatically detect
// and set the `Content-Type` header using `http.DetectContentType()`.
//
// The middleware ignores hijacked responses or requests containing the `Upgrade` header.
//
// **Example:**
//
//	compressMiddleware := &compress.Middleware{
//		Encoders: []compress.Encoder{
//			&compress.Gzip{Level: gzip.BestCompression},
//		},
//	}
type Middleware struct {
	goyave.Component
	Encoders []Encoder
}

// Handle implementation of `goyave.Middleware`.
func (m *Middleware) Handle(next goyave.Handler) goyave.Handler {
	_ = "STUB: not implemented"
	return *new(goyave.Handler)
}

func (m *Middleware) getEncoder(response *goyave.Response, request *goyave.Request) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}
