package goyave

import (
	"errors"
	"io/fs"
	"net/http"
	"regexp"

	"goyave.dev/goyave/v5/cors"
)

// Common route meta keys.
const (
	MetaCORS = "goyave.cors"
)

// Special route names.
const (
	RouteMethodNotAllowed = "goyave.method-not-allowed"
	RouteNotFound         = "goyave.not-found"
)

var (
	errMatchMethodNotAllowed = errors.New("method not allowed for this route")
	errMatchNotFound         = errors.New("no match for this URI")

	methodNotAllowedRoute = newRoute(func(response *Response, _ *Request) {
		response.Status(http.StatusMethodNotAllowed)
	}, RouteMethodNotAllowed)
	notFoundRoute = newRoute(func(response *Response, _ *Request) {
		response.Status(http.StatusNotFound)
	}, RouteNotFound)
)

// Handler responds to an HTTP request.
//
// The given `Response` and `Request` value should not
// be used outside of the context of an HTTP request. e.g.: passed to
// a goroutine or used after the finalization step in the request lifecycle.
type Handler func(response *Response, request *Request)

type routeMatcher interface {
	match(method string, match *routeMatch) bool
}

type routeMatch struct {
	route       *Route
	parameters  map[string]string
	err         error
	currentPath string
}

func (rm *routeMatch) mergeParams(params map[string]string) { _ = "STUB: not implemented"; return }

func (rm *routeMatch) trimCurrentPath(fullMatch string) { _ = "STUB: not implemented"; return }

// Router registers routes to be matched and executes a handler.
type Router struct {
	server         *Server
	parent         *Router
	statusHandlers map[int]StatusHandler
	namedRoutes    map[string]*Route
	regexCache     map[string]*regexp.Regexp
	Meta           map[string]any

	parameterizable
	middlewareHolder
	globalMiddleware *middlewareHolder

	prefix     string
	routes     []*Route
	subrouters []*Router

	slashCount int
}

var _ http.Handler = (*Router)(nil) // implements http.Handler
var _ routeMatcher = (*Router)(nil) // implements routeMatcher

// NewRouter create a new root-level Router that is pre-configured with core
// middleware (recovery and language), as well as status handlers
// for all standard HTTP status codes.
//
// You don't need to manually build your router using this function.
// This method can however be useful for external tooling that build
// routers without starting the HTTP server. Don't forget to call
// `router.ClearRegexCache()` when you are done registering routes.
func NewRouter(server *Server) *Router { _ = "STUB: not implemented"; return nil }

// 444 is a nginx status code. Indicates server to return no information to the client and close the connection immediately

// ClearRegexCache set internal router's regex cache used for route parameters optimisation to nil
// so it can be garbage collected.
// You don't need to call this function if you are using `server.RegisterRoutes`.
func (r *Router) ClearRegexCache() { _ = "STUB: not implemented"; return }

// GetParent returns the parent Router of this router (can be `nil`).
func (r *Router) GetParent() *Router {
	_ = "STUB: not implemented"

	// GetRoutes returns the list of routes belonging to this router.
	return nil
}

func (r *Router) GetRoutes() []*Route { _ = "STUB: not implemented"; return nil }

// GetSubrouters returns the list of subrouters belonging to this router.
func (r *Router) GetSubrouters() []*Router { _ = "STUB: not implemented"; return nil }

// GetRoute get a named route.
// Returns nil if the route doesn't exist.
func (r *Router) GetRoute(name string) *Route { _ = "STUB: not implemented"; return nil }

// SetMeta attach a value to this router identified by the given key.
//
// This value is inherited by all subrouters and routes, unless they override
// it at their level.
func (r *Router) SetMeta(key string, value any) *Router { _ = "STUB: not implemented"; return nil }

// RemoveMeta detach the meta value identified by the given key from this router.
// This doesn't remove meta using the same key from the parent routers.
func (r *Router) RemoveMeta(key string) *Router { _ = "STUB: not implemented"; return nil }

// LookupMeta value identified by the given key. If not found in this router,
// the value is recursively fetched in the parent routers.
//
// Returns the value and `true` if found in the current router or one of the
// parent routers, `nil` and `false` otherwise.
func (r *Router) LookupMeta(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// GlobalMiddleware apply one or more global middleware. Global middleware are
// executed for every request, including when the request doesn't match any route
// or if it results in "Method Not Allowed".
// These middleware are global to the main Router: they will also be executed for subrouters.
// Global Middleware are always executed first.
// Use global middleware for logging and rate limiting for example.
func (r *Router) GlobalMiddleware(middleware ...Middleware) *Router {
	_ = "STUB: not implemented"
	return nil
}

// Middleware apply one or more middleware to the route group.
func (r *Router) Middleware(middleware ...Middleware) *Router {
	_ = "STUB: not implemented"
	return nil
}

// CORS set the CORS options for this route group.
// If the options are not `nil`, the CORS middleware is automatically added globally.
// To disable CORS for this router, subrouters and routes, give `nil` options.
// CORS can be re-enabled for subrouters and routes on a case-by-case basis
// using non-nil options.
func (r *Router) CORS(options *cors.Options) *Router { _ = "STUB: not implemented"; return nil }

// StatusHandler set a handler for responses with an empty body.
// The handler will be automatically executed if the request's life-cycle reaches its end
// and nothing has been written in the response body.
//
// Multiple status codes can be given. The handler will be executed if one of them matches.
//
// This method can be used to define custom error handlers for example.
//
// Status handlers are inherited as a copy in sub-routers. Modifying a child's status handler
// will not modify its parent's.
//
// Codes in the 400 and 500 ranges have a default status handler.
func (r *Router) StatusHandler(handler StatusHandler, status int, additionalStatuses ...int) {
	_ = "STUB: not implemented"
	return
}

// ServeHTTP dispatches the handler registered in the matched route.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO export RouteMatch and add Match with string param function

func (r *Router) match(method string, match *routeMatch) bool {
	_ = "STUB: not implemented"
	// Check if router itself matches
	return false
}

// Ignore slashes in router prefix

// Check in subrouters first

// This allows route groups with subrouters having empty prefix.

// Check if any route matches

// Return true if the subrouter matched so we don't turn back and check other subrouters

func nthIndex(str, substr string, n int) int { _ = "STUB: not implemented"; return 0 }

func (r *Router) makeParameters(match []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Subrouter create a new sub-router from this router.
// Use subrouters to create route groups and to apply middleware to multiple routes.
// CORS options are also inherited.
//
// Subrouters are matched before routes. For example, if you have a subrouter with a
// prefix "/{name}" and a route "/route", the "/route" will never match.
func (r *Router) Subrouter(prefix string) *Router { _ = "STUB: not implemented"; return nil }

// Typical CRUD has 5 routes

// Group create a new sub-router with an empty prefix.
func (r *Router) Group() *Router { _ = "STUB: not implemented"; return nil }

// Route register a new route.
//
// Multiple methods can be passed.
//
// If the route matches the "GET" method, the "HEAD" method is automatically added
// to the matcher if it's missing.
//
// If the router has the CORS middleware, the "OPTIONS" method is automatically added
// to the matcher if it's missing, so it allows preflight requests.
//
// Returns the generated route.
func (r *Router) Route(methods []string, uri string, handler Handler) *Route {
	_ = "STUB: not implemented"
	return nil
}

// Get registers a new route with the GET and HEAD methods.
func (r *Router) Get(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Post registers a new route with the POST method.
func (r *Router) Post(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Put registers a new route with the PUT method.
func (r *Router) Put(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Patch registers a new route with the PATCH method.
func (r *Router) Patch(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Delete registers a new route with the DELETE method.
func (r *Router) Delete(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Options registers a new route wit the OPTIONS method.
func (r *Router) Options(uri string, handler Handler) *Route { _ = "STUB: not implemented"; return nil }

// Static serve a directory and its subdirectories of static resources.
// Set the "download" parameter to true if you want the files to be sent as an attachment
// instead of an inline element.
//
// If no file is given in the url, or if the given file is a directory, the handler will
// send the "index.html" file if it exists.
//
// As a precaution, all requests with a path containing a path segment in the form of ".", ".." are rejected with
// `http.StatusNotFound`. This ensures clients cannot access files outside of the given filesystem base directory.
// Paths containing  "\" or "//" are also rejected.
func (r *Router) Static(fs fs.StatFS, uri string, download bool) *Route {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) registerRoute(methods []string, uri string, handler Handler) *Route {
	_ = "STUB: not implemented"
	return nil
}

// Controller register all routes for a controller implementing the `Registrer` interface.
// Automatically calls `Init()` and `RegisterRoutes()` on the given controller.
func (r *Router) Controller(controller Registrer) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) requestHandler(match *routeMatch, w http.ResponseWriter, rawRequest *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Route-specific middleware is executed after router middleware

// finalize the request's life-cycle.
func (r *Router) finalize(match *routeMatch, response *Response, request *Request) error {
	_ = "STUB: not implemented"
	return nil
}

// If the response is empty, return status 204 to
// comply with RFC 7231, 6.3.5

// An error occurred after the response header has been written. We still want to execute the panic status handler.

// Status has been set but body is empty.
// Execute status handler if exists.

// An error has occurred after a write. We still want to execute the panic status handler.

func (r *Router) getStatusHandler(match *routeMatch, status int) (StatusHandler, bool) {
	_ = "STUB: not implemented"
	return *new(StatusHandler), false
}
