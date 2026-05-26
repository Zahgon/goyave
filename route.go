package goyave

import (
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/validation"
)

// Route stores information for route matching and serving and can be
// used to generate dynamic URLs/URIs. Routes can, just like routers,
// hold Meta information that can be used by generic middleware to
// alter their behavior depending on the route being served.
type Route struct {
	name    string
	uri     string
	methods []string
	parent  *Router
	Meta    map[string]any
	handler Handler
	middlewareHolder
	parameterizable
}

var _ routeMatcher = (*Route)(nil) // implements routeMatcher

// RuleSetFunc function generating a new validation rule set.
// This function is called for every validated request.
// The returned value is expected to be fresh, not re-used across
// multiple requests nor concurrently.
type RuleSetFunc func(*Request) validation.RuleSet

// newRoute create a new route without any settings except its handler.
// This is used to generate a fake route for the Method Not Allowed and Not Found handlers.
// This route has the core middleware enabled and can be used without a parent router.
// Thus, custom status handlers can use language and body.
func newRoute(handler Handler, name string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) match(method string, match *routeMatch) bool {
	_ = "STUB: not implemented"
	return false
}

// Don't override error if already set.
// Not nil error means it's either already errMatchNotFound
// or it's errMatchMethodNotAllowed, implying that a route has
// already been matched but with wrong method.

func (r *Route) checkMethod(method string) bool { _ = "STUB: not implemented"; return false }

func (r *Route) makeParameters(match []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Name set the name of the route.
// Panics if a route with the same name already exists.
// Returns itself.
func (r *Route) Name(name string) *Route { _ = "STUB: not implemented"; return nil }

// SetMeta attach a value to this route identified by the given key.
//
// This value can override a value inherited by the parent routers for this route only.
func (r *Route) SetMeta(key string, value any) *Route { _ = "STUB: not implemented"; return nil }

// RemoveMeta detach the meta value identified by the given key from this route.
// This doesn't remove meta using the same key from the parent routers.
func (r *Route) RemoveMeta(key string) *Route { _ = "STUB: not implemented"; return nil }

// LookupMeta value identified by the given key. If not found in this route,
// the value is recursively fetched in the parent routers.
//
// Returns the value and `true` if found in the current route or one of the
// parent routers, `nil` and `false` otherwise.
func (r *Route) LookupMeta(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// ValidateBody adds (or replace) validation rules for the request body.
func (r *Route) ValidateBody(validationRules RuleSetFunc) *Route {
	_ = "STUB: not implemented"
	return nil
}

// ValidateQuery adds (or replace) validation rules for the request query.
func (r *Route) ValidateQuery(validationRules RuleSetFunc) *Route {
	_ = "STUB: not implemented"
	return nil
}

// CORS set the CORS options for this route only.
// The "OPTIONS" method is added if this route doesn't already support it.
//
// If the options are not `nil`, the CORS middleware is automatically added globally.
// To disable CORS, give `nil` options. The "OPTIONS" method will be removed
// if it isn't the only method for this route.
func (r *Route) CORS(options *cors.Options) *Route { _ = "STUB: not implemented"; return nil }

// Middleware register middleware for this route only.
//
// Returns itself.
func (r *Route) Middleware(middleware ...Middleware) *Route { _ = "STUB: not implemented"; return nil }

// BuildURL build a full URL pointing to this route.
// Panics if the amount of parameters doesn't match the amount of
// actual parameters for this route.
func (r *Route) BuildURL(parameters ...string) string { _ = "STUB: not implemented"; return "" }

// BuildProxyURL build a full URL pointing to this route using the proxy base URL.
// Panics if the amount of parameters doesn't match the amount of
// actual parameters for this route.
func (r *Route) BuildProxyURL(parameters ...string) string { _ = "STUB: not implemented"; return "" }

// BuildURI build a full URI pointing to this route. The returned
// string doesn't include the protocol and domain. (e.g. "/user/login")
// Panics if the amount of parameters doesn't match the amount of
// actual parameters for this route.
func (r *Route) BuildURI(parameters ...string) string { _ = "STUB: not implemented"; return "" }

// Skip closing braces

// GetName get the name of this route.
func (r *Route) GetName() string {
	_ = "STUB: not implemented"

	// GetURI get the URI of this route.
	// The returned URI is relative to the parent router of this route, it is NOT
	// the full path to this route.
	//
	// Note that this URI may contain route parameters in their définition format.
	// Use the request's URI if you want to see the URI as it was requested by the client.
	return ""
}

func (r *Route) GetURI() string {
	_ = "STUB: not implemented"

	// GetFullURI get the full URI of this route.
	//
	// Note that this URI may contain route parameters in their définition format.
	// Use the request's URI if you want to see the URI as it was requested by the client.
	return ""
}

func (r *Route) GetFullURI() string { _ = "STUB: not implemented"; return "" }

// Revert segements

// GetMethods returns the methods the route matches against.
func (r *Route) GetMethods() []string { _ = "STUB: not implemented"; return nil }

// GetHandler returns the Handler associated with this route.
func (r *Route) GetHandler() Handler {
	_ = "STUB: not implemented"

	// GetParent returns the parent Router of this route.
	return *new(Handler)
}

func (r *Route) GetParent() *Router {
	_ = "STUB: not implemented"

	// GetFullURIAndParameters get the full uri and parameters for this route and all its parent routers.
	return nil
}

func (r *Route) GetFullURIAndParameters() (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

// Revert segements

// Revert parameters
