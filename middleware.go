package goyave

// Middleware are special handlers executed in a stack above the controller handler.
// They allow to inspect and filter requests, transform responses or provide additional
// information to the next handlers in the stack.
// Example uses are authentication, authorization, logging, panic recovery, CORS,
// validation, gzip compression.
type Middleware interface {
	Composable
	Handle(next Handler) Handler
}

type middlewareHolder struct {
	middleware []Middleware
}

func (h *middlewareHolder) applyMiddleware(handler Handler) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

// GetMiddleware returns a copy of the middleware applied on this holder.
func (h *middlewareHolder) GetMiddleware() []Middleware { _ = "STUB: not implemented"; return nil }

func findMiddleware[T Middleware](m []Middleware) T { _ = "STUB: not implemented"; return *new(T) }

func hasMiddleware[T Middleware](m []Middleware) bool { _ = "STUB: not implemented"; return false }

// routeHasMiddleware returns true if the given route or any of its
// parents has a middleware of the T type.
func routeHasMiddleware[T Middleware](route *Route) bool { _ = "STUB: not implemented"; return false }

// routerHasMiddleware returns true if the given route or any of its
// parents has a middleware of the T type. Also returns true if the middleware
// is present as global middleware.
func routerHasMiddleware[T Middleware](router *Router) bool {
	_ = "STUB: not implemented"
	return false
}

// recoveryMiddleware is a middleware that recovers from panic and sends a 500 error code.
// If debugging is enabled in the config and the default status handler for the 500 status code
// had not been changed, the error is also written in the response.
type recoveryMiddleware struct {
	Component
}

func (m *recoveryMiddleware) Handle(next Handler) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

// Skipped: runtime.Callers, NewSkip, this func, runtime.panic

// Force status override if the header hasn't been written yet.

// languageMiddleware is a middleware that sets the language of a request.
//
// Uses the "Accept-Language" header to determine which language to use. If
// the header is not set or the language is not available, uses the default
// language as fallback.
//
// If "*" is provided, the default language will be used.
// If multiple languages are given, the first available language will be used,
// and if none are available, the default language will be used.
// If no variant is given (for example "en"), the first available variant will be used.
// For example, if "en-US" and "en-UK" are available and the request accepts "en",
// "en-US" will be used.
type languageMiddleware struct {
	Component
}

func (m *languageMiddleware) Handle(next Handler) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

// validateRequestMiddleware is a middleware that validates the request.
// If validation is not rules are not met, sets the response status to 422 Unprocessable Entity
// or 400 Bad Request and the response error (which can be retrieved with `GetError()`) to the
// `validation.Errors` returned by the validator.
// This data can then be used in a status handler.
// This middleware requires the parse middleware.
type validateRequestMiddleware struct {
	Component
	BodyRules  RuleSetFunc
	QueryRules RuleSetFunc
}

func (m *validateRequestMiddleware) Handle(next Handler) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

type corsMiddleware struct {
	Component
}

func (m *corsMiddleware) Handle(next Handler) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}
