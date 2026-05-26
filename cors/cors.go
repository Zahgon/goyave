package cors

import (
	"net/http"
	"time"
)

// Options holds the CORS configuration for a router.
type Options struct {
	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the first value in the slice is "*" or if the slice is empty, all origins will be allowed.
	// Default value is ["*"]
	AllowedOrigins []string

	// AllowedMethods is a list of methods the client is allowed to use with cross-domain requests.
	// Default value is ["HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"].
	AllowedMethods []string

	// AllowedHeaders is a list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the first value in the slice is "*", all headers will be allowed.
	// If the slice is empty, the request's headers will be reflected.
	// Default value is ["Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"].
	AllowedHeaders []string

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string

	// MaxAge indicates how long the results of a preflight request can be cached.
	// Default is 12 hours.
	MaxAge time.Duration

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool
}

// Default create new CORS options with default settings.
// The returned value can be used as a starting point for
// customized options.
func Default() *Options { _ = "STUB: not implemented"; return nil }

// ConfigureCommon configures common headers between regular and preflight requests:
// Origin, Credentials and Exposed Headers.
func (o *Options) ConfigureCommon(headers http.Header, requestHeaders http.Header) {
	_ = "STUB: not implemented"
	return
}

func (o *Options) configureOrigin(headers http.Header, requestHeaders http.Header) {
	_ = "STUB: not implemented"
	return
}

func (o *Options) configureCredentials(headers http.Header) { _ = "STUB: not implemented"; return }

func (o *Options) configureExposedHeaders(headers http.Header) { _ = "STUB: not implemented"; return }

func (o *Options) configureAllowedMethods(headers http.Header) { _ = "STUB: not implemented"; return }

func (o *Options) configureAllowedHeaders(headers http.Header, requestHeaders http.Header) {
	_ = "STUB: not implemented"
	return
}

func (o *Options) configureMaxAge(headers http.Header) { _ = "STUB: not implemented"; return }

// HandlePreflight configures headers for preflight requests:
// Allowed Methods, Allowed Headers and Max Age.
func (o *Options) HandlePreflight(headers http.Header, requestHeaders http.Header) {
	_ = "STUB: not implemented"
	return
}

func (o *Options) validateOrigin(requestHeaders http.Header) bool {
	_ = "STUB: not implemented"
	return false
}
