package parse

import (
	"net/http"
	"net/url"

	"goyave.dev/goyave/v5"
)

// Middleware reading the raw request query and body.
//
// First, the query is parsed using Go's standard `url.ParseQuery()`. After being flattened
// (single value arrays converted to non-array), the result is put in the request's `Query`.
// If the parsing fails, returns "400 Bad request".
//
// The body is read only if the "Content-Type" header is set. If
// the body exceeds the configured max upload size (in MiB), "413 Request Entity Too Large"
// is returned.
// If the content type is "application/json", the middleware will attempt
// to unmarshal the body and put the result in the request's `Data`. If it fails, returns "400 Bad request".
// If the content-type has another value, Go's standard `ParseMultipartForm` is called. The result
// is put inside the request's `Data` after being flattened.
// If the form is not a multipart form, attempts `ParseForm`. If `ParseMultipartForm` or `ParseForm` return
// an error, returns "400 Bad request".
//
// This middleware depletes the request's body reader. You cannot read `request.Body()` to access the unparsed
// data afterwards.
//
// In `multipart/form-data`, all file parts are automatically converted to `[]fsutil.File`.
// Inside `request.Data`, a field of type "file" will therefore always be of type `[]fsutil.File`.
// It is a slice so it support multi-file uploads in a single field.
type Middleware struct {
	goyave.Component

	// MaxUpoadSize the maximum size of the request (in MiB).
	// Defaults to the value provided in the config "server.maxUploadSize".
	MaxUploadSize float64
}

// Handle reads the request query and body and parses it if necessary.
// If the request Data is not nil, the body is not parsed again and the
// middleware immediately passes after parsing the query.
// If the current route is the special "not found" or "method not allowed" route,
// the middleware is skipped and immediately passes.
func (m *Middleware) Handle(next goyave.Handler) goyave.Handler {
	_ = "STUB: not implemented"
	return *new(goyave.Handler)
}

func (m *Middleware) getMaxUploadSize() float64 { _ = "STUB: not implemented"; return 0 }

func parseQuery(request *goyave.Request) error { _ = "STUB: not implemented"; return nil }

func generateFlatMap(request *http.Request, maxSize int64) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prevent Form from being parsed because it would be redundant with our parsing

// Source form is not needed anymore, clear it.

func flatten(dst map[string]any, values url.Values) { _ = "STUB: not implemented"; return }
