package goyave

// StatusHandler is a regular handler executed during the finalization step of the request's lifecycle
// if the response body is empty but a status code has been set.
// Status handlers are mainly used to implement a custom behavior for user or server errors (400 and 500 status codes).
type StatusHandler interface {
	Composable
	Handle(response *Response, request *Request)
}

// PanicStatusHandler for the HTTP 500 error.
// If debugging is enabled, writes the error details to the response and
// print stacktrace in the console.
// If debugging is not enabled, writes `{"error": "Internal Server Error"}`
// to the response.
type PanicStatusHandler struct {
	Component
}

// Handle internal server error responses.
func (*PanicStatusHandler) Handle(response *Response, _ *Request) {
	_ = "STUB: not implemented"
	return
}

// ErrorStatusHandler a generic status handler for non-success codes.
// Writes the corresponding status message to the response.
type ErrorStatusHandler struct {
	Component
}

// Handle generic error responses.
func (*ErrorStatusHandler) Handle(response *Response, _ *Request) {
	_ = "STUB: not implemented"
	return
}

// ParseErrorStatusHandler a generic (error) status handler for requests.
type ParseErrorStatusHandler struct {
	Component
}

// Handle generic request (error) responses.
func (h *ParseErrorStatusHandler) Handle(response *Response, request *Request) {
	_ = "STUB: not implemented"
	return
}

// ValidationStatusHandler for HTTP 422 errors.
// Writes the validation errors to the response.
type ValidationStatusHandler struct {
	Component
}

// Handle validation error responses.
func (*ValidationStatusHandler) Handle(response *Response, request *Request) {
	_ = "STUB: not implemented"
	return
}
