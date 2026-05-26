package testutil

import (
	"bytes"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/fsutil"
)

type copyRequestMiddleware struct {
	goyave.Component
	request *goyave.Request
}

func (m *copyRequestMiddleware) Handle(next goyave.Handler) goyave.Handler {
	_ = "STUB: not implemented"
	return *new(goyave.Handler)
}

// TestServer extension of `goyave.Server` providing useful functions for testing.
type TestServer struct {
	*goyave.Server
}

// NewTestServer creates a new server using the given config file. The config path is relative to
// the project's directory. If not nil, the given `routeRegistrer` function is called to register
// routes without starting the server.
//
// A default logger redirecting the output to `testing.T.Log()` is used.
//
// Automatically closes the DB connection (if there is one) using a test `Cleanup` function.
func NewTestServer(t *testing.T, configFileName string) *TestServer {
	_ = "STUB: not implemented"
	return nil
}

// NewTestServerWithOptions creates a new server using the given options.
// If not nil, the given `routeRegistrer` function is called to register
// routes without starting the server.
//
// By default, if no `Logger` is given in the options, a default logger redirecting the
// output to `io.Discard` is used.
//
// Automatically closes the DB connection (if there is one) using a test `Cleanup` function.
func NewTestServerWithOptions(t *testing.T, opts goyave.Options) *TestServer {
	_ = "STUB: not implemented"
	return nil
}

// TestRequest execute a request by calling the root Router's `ServeHTTP()` implementation.
func (s *TestServer) TestRequest(request *http.Request) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// TestMiddleware executes with the given request and returns the response.
// The `procedure` parameter is the `next` handler passed to the middleware and can be used to
// make assertions. Keep in mind that this procedure won't be executed if your middleware is blocking.
//
// The request will go through the entire lifecycle like a regular request.
//
// The given request is cloned. If the middleware alters the request object, these changes won't be reflected on the input request.
// You can do your assertions inside the `procedure`.
func (s *TestServer) TestMiddleware(middleware goyave.Middleware, request *goyave.Request, procedure goyave.Handler) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// CloseDB close the server DB if one is open. It is a good practice to always
// call this in a test `Cleanup` function when using a database.
func (s *TestServer) CloseDB() { _ = "STUB: not implemented"; return }

// FindRootDirectory find relative path to the project's root directory based on the
// existence of a `go.mod` file. The returned path is a rooted path.
// Returns an empty string if not found.
func FindRootDirectory() string { _ = "STUB: not implemented"; return "" }

// NewTestRequest create a new `goyave.Request` with an underlying HTTP request created
// usin the `httptest` package.
func NewTestRequest(method, uri string, body io.Reader) *goyave.Request {
	_ = "STUB: not implemented"
	return nil
}

// NewTestRequest create a new `goyave.Request` with an underlying HTTP request created
// usin the `httptest` package. This function sets the request language using the default
// language of the server.
func (s *TestServer) NewTestRequest(method, uri string, body io.Reader) *goyave.Request {
	_ = "STUB: not implemented"
	return nil
}

// NewTestResponse create a new `goyave.Response` with an underlying HTTP response recorder created
// using the `httptest` package. This function uses a temporary `goyave.Server` with all defaults values loaded
// so all functions of `*goyave.Response` can be used safely.
func NewTestResponse(request *goyave.Request) (*goyave.Response, *httptest.ResponseRecorder) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTestResponse create a new `goyave.Response` with an underlying HTTP response recorder created
// using the `httptest` package.
func (s *TestServer) NewTestResponse(request *goyave.Request) (*goyave.Response, *httptest.ResponseRecorder) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadJSONBody decodes the given body reader into a new variable of type `*T`.
func ReadJSONBody[T any](body io.Reader) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// WriteMultipartFile reads a file from the given FS and writes it to the given multipart writer.
func WriteMultipartFile(writer *multipart.Writer, filesystem fs.FS, path, fieldName, fileName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CreateTestFiles create a slice of "fsutil.File" from the given FS.
// To reproduce the way the files are obtained in real scenarios,
// files are first encoded in a multipart form, then decoded with
// a multipart form reader.
//
// Paths are relative to the caller, not relative to the project's root directory.
func CreateTestFiles(fs fs.FS, paths ...string) ([]fsutil.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToJSON marshals the given data and creates a bytes reader from the result.
// Panics on error.
func ToJSON(data any) *bytes.Reader { _ = "STUB: not implemented"; return nil }
