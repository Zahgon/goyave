package goyave

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"net"
	"net/http"
	"sync"

	errorutil "goyave.dev/goyave/v5/util/errors"
)

var (
	// ErrNotHijackable returned by response.Hijack() if the underlying
	// http.ResponseWriter doesn't implement http.Hijacker. This can
	// happen with HTTP/2 connections.
	ErrNotHijackable = errors.New("underlying http.ResponseWriter doesn't implement http.Hijacker")
)

// PreWriter is a writter that needs to alter the response headers or status
// before they are written.
// If implemented, PreWrite will be called right before the first `Write` operation.
type PreWriter interface {
	PreWrite(b []byte)
}

// The Flusher interface is implemented by writers that allow
// handlers to flush buffered data to the client.
//
// Note that even for writers that support flushing, if the client
// is connected through an HTTP proxy, the buffered data may not reach
// the client until the response completes.
type Flusher interface {
	Flush() error
}

// CommonWriter is a component meant to be used with composition
// to avoid having to implement the base behavior of the common interfaces
// a chained writer has to implement (`PreWrite()`, `Write()`, `Close()`, `Flush()`)
type CommonWriter struct {
	Component
	wr io.Writer
}

// NewCommonWriter create a new common writer that will output to the given `io.Writer`.
func NewCommonWriter(wr io.Writer) CommonWriter {
	_ = "STUB: not implemented"
	return *new(CommonWriter)
}

// PreWrite calls PreWrite on the
// child writer if it implements PreWriter.
func (w CommonWriter) PreWrite(b []byte) { _ = "STUB: not implemented"; return }

func (w CommonWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close the underlying writer if it implements `io.Closer`.
func (w CommonWriter) Close() error { _ = "STUB: not implemented"; return nil }

// Flush the underlying writer if it implements `goyave.Flusher` or `http.Flusher`.
func (w *CommonWriter) Flush() error { _ = "STUB: not implemented"; return nil }

// Writer returns the underlying writer.
func (w *CommonWriter) Writer() io.Writer {
	_ = "STUB: not implemented"

	// Response implementation wrapping `http.ResponseWriter`. Writing an HTTP response without
	// using it is incorrect. This acts as a proxy to one or many `io.Writer` chained, with the original
	// `http.ResponseWriter` always last.
	return *new(io.Writer)
}

type Response struct {
	writer         io.Writer
	responseWriter http.ResponseWriter
	server         *Server
	request        *Request
	err            *errorutil.Error
	status         int

	// Used to check if controller didn't write anything so
	// core can write default 204 No Content.
	// See RFC 7231, 6.3.5
	empty       bool
	wroteHeader bool
	hijacked    bool
}

var responsePool = sync.Pool{
	New: func() any {
		return &Response{}
	},
}

// NewResponse create a new Response using the given `http.ResponseWriter` and request.
func NewResponse(server *Server, request *Request, writer http.ResponseWriter) *Response {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) reset(server *Server, request *Request, writer http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

// --------------------------------------
// PreWriter implementation

// PreWrite writes the response header after calling PreWrite on the
// child writer if it implements PreWriter.
func (r *Response) PreWrite(b []byte) { _ = "STUB: not implemented"; return }

// --------------------------------------
// http.ResponseWriter implementation

// Write writes the data as a response.
// See `http.ResponseWriter.Write`.
func (r *Response) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteHeader sends an HTTP response header with the provided
// status code.
// Prefer using "Status()" method instead.
// Calling this method a second time will have no effect.
func (r *Response) WriteHeader(status int) { _ = "STUB: not implemented"; return }

// Header returns the header map that will be sent.
func (r *Response) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

// Cookie add a Set-Cookie header to the response.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
func (r *Response) Cookie(cookie *http.Cookie) { _ = "STUB: not implemented"; return }

// Flush sends any buffered data to the client if the underlying
// writer implements `goyave.Flusher`.
//
// If the response headers have not been written already, `PreWrite()` will
// be called with an empty byte slice.
func (r *Response) Flush() { _ = "STUB: not implemented"; return }

// --------------------------------------
// http.Hijacker implementation

// Hijack implements the Hijacker.Hijack method.
// For more details, check http.Hijacker.
//
// Returns ErrNotHijackable if the underlying http.ResponseWriter doesn't
// implement http.Hijacker. This can happen with HTTP/2 connections.
//
// Middleware executed after controller handlers, as well as status handlers,
// keep working as usual after a connection has been hijacked.
// Callers should properly set the response status to ensure middleware and
// status handler execute correctly. Usually, callers of the Hijack method
// set the HTTP status to http.StatusSwitchingProtocols.
// If no status is set, the regular behavior will be kept and `204 No Content`
// will be set as the response status.
func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// Hijacked returns true if the underlying connection has been successfully hijacked
// via the Hijack method.
func (r *Response) Hijacked() bool {
	_ = "STUB: not implemented"

	// --------------------------------------
	// Chained writers
	return false
}

// Writer return the current writer used to write the response.
// Note that the returned writer is not necessarily a http.ResponseWriter, as
// it can be replaced using SetWriter.
func (r *Response) Writer() io.Writer {
	_ = "STUB: not implemented"

	// SetWriter set the writer used to write the response.
	// This can be used to chain writers, for example to enable
	// gzip compression, or for logging.
	//
	// The original http.ResponseWriter is always kept.
	return *new(io.Writer)
}

func (r *Response) SetWriter(writer io.Writer) { _ = "STUB: not implemented"; return }

func (r *Response) close() error { _ = "STUB: not implemented"; return nil }

// --------------------------------------
// Accessors

// GetStatus return the response code for this request or 0 if not yet set.
func (r *Response) GetStatus() int {
	_ = "STUB: not implemented"

	// IsEmpty return true if nothing has been written to the response body yet.
	return 0
}

func (r *Response) IsEmpty() bool {
	_ = "STUB: not implemented"

	// IsHeaderWritten return true if the response header has been written.
	// Once the response header is written, you cannot change the response status
	// and headers anymore.
	return false
}

func (r *Response) IsHeaderWritten() bool { _ = "STUB: not implemented"; return false }

// GetError return the `*errors.Error` that occurred in the process of this response, or `nil`.
// The error can be set by:
//   - Calling `Response.Error()`
//   - The recovery middleware
//   - The status handler for the 500 status code, if the error is not already set
func (r *Response) GetError() *errorutil.Error {
	_ = "STUB: not implemented"

	// --------------------------------------
	// Write methods
	return nil
}

// Status set the response status code.
// Calling this method a second time will have no effect.
func (r *Response) Status(status int) { _ = "STUB: not implemented"; return }

// JSON write json data as a response.
// Also sets the "Content-Type" header automatically.
func (r *Response) JSON(responseCode int, data any) { _ = "STUB: not implemented"; return }

// String write a string as a response.
func (r *Response) String(responseCode int, message string) { _ = "STUB: not implemented"; return }

func (r *Response) writeFile(fs fs.StatFS, file string, disposition string) {
	_ = "STUB: not implemented"
	return
}

// File write a file as an inline element.
// Automatically detects the file MIME type and sets the "Content-Type" header accordingly.
// If the file doesn't exist, respond with status 404 Not Found.
// The given path can be relative or absolute.
//
// If you want the file to be sent as a download ("Content-Disposition: attachment"), use the "Download" function instead.
func (r *Response) File(fs fs.StatFS, file string) { _ = "STUB: not implemented"; return }

// Download write a file as an attachment element.
// Automatically detects the file MIME type and sets the "Content-Type" header accordingly.
// If the file doesn't exist, respond with status 404 Not Found.
// The given path can be relative or absolute.
//
// The "fileName" parameter defines the name the client will see. In other words, it sets the header "Content-Disposition" to
// "attachment; filename="${fileName}""
//
// If you want the file to be sent as an inline element ("Content-Disposition: inline"), use the "File" function instead.
func (r *Response) Download(fs fs.StatFS, file string, fileName string) {
	_ = "STUB: not implemented"
	return
}

// Error print the error in the console and return it with an error code 500 (or previously defined
// status code using `response.Status()`).
// If debugging is enabled in the config, the error is also written in the response
// and the stacktrace is printed in the console.
// If debugging is not enabled, only the status code is set, which means you can still
// write to the response, or use your error status handler.
func (r *Response) Error(err any) { _ = "STUB: not implemented"; return }

// Skipped: runtime.Callers, NewSkip, this func

func (r *Response) error(err any) { _ = "STUB: not implemented"; return }

// Skipped: runtime.Callers, NewSkip, this func

// Don't set r.empty to false to let error status handler process the error

// WriteDBError takes an error and automatically writes HTTP status code 404 Not Found
// if the error is a `gorm.ErrRecordNotFound` error.
// Calls `Response.Error()` if there is another type of error.
//
// Returns true if there is an error. You can then safely `return` in you controller.
//
//	func (ctrl *ProductController) Show(response *goyave.Response, request *goyave.Request) {
//	    product := model.Product{}
//	    result := ctrl.DB().First(&product, request.RouteParams["id"])
//	    if response.WriteDBError(result.Error) {
//	        return
//	    }
//	    response.JSON(http.StatusOK, product)
//	}
func (r *Response) WriteDBError(err error) bool { _ = "STUB: not implemented"; return false }
