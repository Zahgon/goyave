package goyave

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"goyave.dev/goyave/v5/lang"
)

type (
	// ExtraBodyValidationRules the key used in `Context.Extra` to
	// store the body validation rules.
	ExtraBodyValidationRules struct{}

	// ExtraQueryValidationRules the key used in `Context.Extra` to
	// store the query validation rules.
	ExtraQueryValidationRules struct{}

	// ExtraValidationError the key used in `Context.Extra` to
	// store the body validation errors.
	ExtraValidationError struct{}

	// ExtraQueryValidationError the key used in `Context.Extra` to
	// store the query validation errors.
	ExtraQueryValidationError struct{}

	// ExtraParseError the key used in `Context.Extra` to
	// store specific parsing errors.
	ExtraParseError struct{}
)

var (
	// ErrInvalidQuery error when an invalid query string is passed.
	ErrInvalidQuery = errors.New("parse middleware: could not parse query")

	// ErrInvalidJSONBody error when an empty or malformed JSON body is sent.
	ErrInvalidJSONBody = errors.New("parse middleware: could not JSON unmarshal body")

	// ErrInvalidContentForType error when e.g. a multipart form is not actually multipart, or empty.
	ErrInvalidContentForType = errors.New("parse middleware: could not parse form")

	// ErrErrorInRequestBody error when e.g. a incoming request is not received properly.
	ErrErrorInRequestBody = errors.New("parse middleware: could not read body")
)

// Request represents a http request received by the server.
type Request struct {
	httpRequest *http.Request
	Now         time.Time
	Data        any
	User        any
	Query       map[string]any
	Lang        *lang.Language

	// Extra can be used to store any extra information related to the request.
	// For example, the JWT middleware stores the token claim in the extras.
	//
	// The keys must be comparable and should not be of type
	// string or any other built-in type to avoid collisions.
	// To avoid allocating when assigning to an `interface{}`, context keys often have
	// concrete type `struct{}`. Alternatively, exported context key variables' static
	// type should be a pointer or interface.
	Extra       map[any]any
	Route       *Route
	RouteParams map[string]string
	cookies     []*http.Cookie
}

var requestPool = sync.Pool{
	New: func() any {
		return &Request{}
	},
}

// NewRequest create a new Request from the given raw http request.
// Initializes Now with the current time and Extra with a non-nil map.
func NewRequest(httpRequest *http.Request) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) reset(httpRequest *http.Request) { _ = "STUB: not implemented"; return }

// Request return the raw http request.
// Prefer using the "goyave.Request" accessors.
func (r *Request) Request() *http.Request { _ = "STUB: not implemented"; return nil }

// Method specifies the HTTP method (GET, POST, PUT, etc.).
func (r *Request) Method() string { _ = "STUB: not implemented"; return "" }

// Protocol the protocol used by this request, "HTTP/1.1" for example.
func (r *Request) Protocol() string { _ = "STUB: not implemented"; return "" }

// URL specifies the URL being requested.
func (r *Request) URL() *url.URL { _ = "STUB: not implemented"; return nil }

// Header contains the request header fields either received
// by the server or to be sent by the client.
// Header names are case-insensitive.
//
// If the raw request has the following header lines,
//
//	Host: example.com
//	accept-encoding: gzip, deflate
//	Accept-Language: en-us
//	fOO: Bar
//	foo: two
//
// then the header map will look like this:
//
//	Header = map[string][]string{
//		"Accept-Encoding": {"gzip, deflate"},
//		"Accept-Language": {"en-us"},
//		"Foo": {"Bar", "two"},
//	}
func (r *Request) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

// ContentLength records the length of the associated content.
// The value -1 indicates that the length is unknown.
func (r *Request) ContentLength() int64 { _ = "STUB: not implemented"; return 0 }

// RemoteAddress allows to record the network address that
// sent the request, usually for logging.
func (r *Request) RemoteAddress() string { _ = "STUB: not implemented"; return "" }

// Cookies returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*http.Cookie { _ = "STUB: not implemented"; return nil }

// Referrer returns the referring URL, if sent in the request.
func (r *Request) Referrer() string { _ = "STUB: not implemented"; return "" }

// UserAgent returns the client's User-Agent, if sent in the request.
func (r *Request) UserAgent() string { _ = "STUB: not implemented"; return "" }

// BasicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
func (r *Request) BasicAuth() (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false

	// BearerToken extract the auth token from the "Authorization" header.
	// Only takes tokens of type "Bearer".
	// Returns empty string if no token found or the header is invalid.
}

func (r *Request) BearerToken() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Body the request body.
// Always non-nil, but will return EOF immediately when no body is present.
// The server will close the request body so handlers don't need to.
func (r *Request) Body() io.ReadCloser {
	_ = "STUB: not implemented"
	return *

	// Context returns the request's context. To change the context, use `WithContext`.
	//
	// The returned context is always non-nil; it defaults to the
	// background context.
	//
	// The context is canceled when the client's connection closes, the request is canceled (with HTTP/2),
	// or when the `ServeHTTP` method returns (after the finalization step of the request lifecycle).
	new(io.ReadCloser)
}

func (r *Request) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithContext creates a shallow copy of the underlying `*http.Request` with
// its context changed to `ctx` then returns itself.
// The provided ctx must be non-nil.
func (r *Request) WithContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }
