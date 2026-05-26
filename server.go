package goyave

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"
	"sync/atomic"

	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/lang"
	"goyave.dev/goyave/v5/slog"
	"goyave.dev/goyave/v5/util/fsutil"
)

// serverKey is a context key used to store the server instance into its base context.
type serverKey struct{}

// Options represent server creation options.
type Options struct {

	// Config used by the server and propagated to all its components.
	// If no configuration is provided, automatically load
	// the default configuration using `config.Load()`.
	Config *config.Config

	// Logger used by the server and propagated to all its components.
	// If no logger is provided in the options, uses the default logger.
	Logger *slog.Logger

	// LangFS the file system from which the language files
	// will be loaded. This file system is expected to contain
	// a `resources/lang` directory.
	// If not provided, uses `osfs.FS` as a default.
	LangFS fsutil.FS

	// HTTP2 configures HTTP/2 connections.
	HTTP2 *http.HTTP2Config

	// ListenConfig optionally specifies the configuration for the network listener.
	// If not provided, the default net.ListenConfig is used.
	// This can be useful for customizing keep-alives and other network-level settings
	// for optimal performance in large traffic scenarios.
	ListenConfig *net.ListenConfig

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// `http.ConnState` type and associated constants for details.
	ConnState func(net.Conn, http.ConnState)

	// Context optionnally defines a function that returns the base context
	// for the server. It will be used as base context for all incoming requests.
	//
	// The provided `net.Listener` is the specific Listener that's
	// about to start accepting requests.
	//
	// If not given, the default is `context.Background()`.
	//
	// The context returned then has a the server instance added to it as a value.
	// The server can thus be retrieved using `goyave.ServerFromContext(ctx)`.
	//
	// If the context is canceled, the server won't shut down automatically, you are
	// responsible of calling `server.Stop()` if you want this to happen. Otherwise the
	// server will continue serving requests, at the risk of generating "context canceled" errors.
	BaseContext func(net.Listener) context.Context

	// ConnContext optionally specifies a function that modifies
	// the context used for a new connection `c`. The provided context
	// is derived from the base context and has the server instance value, which can
	// be retrieved using `goyave.ServerFromContext(ctx)`.
	ConnContext func(ctx context.Context, c net.Conn) context.Context

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, http.DefaultMaxHeaderBytes is used.
	MaxHeaderBytes int
}

// Server the central component of a Goyave application.
type Server struct {
	server *http.Server
	config *config.Config
	Lang   *lang.Languages

	router *Router
	db     *gorm.DB

	services map[string]Service

	// Logger the logger for default output
	// Writes to stderr by default.
	Logger *slog.Logger

	host         string
	baseURL      string
	proxyBaseURL string

	stopChannel chan struct{}
	sigChannel  chan os.Signal

	ctx           context.Context
	baseContext   func(net.Listener) context.Context
	listenConfig  *net.ListenConfig
	startupHooks  []func(*Server)
	shutdownHooks []func(*Server)

	port int

	state atomic.Uint32 // 0 -> created, 1 -> preparing, 2 -> ready, 3 -> stopped
}

// New create a new `Server` using the given options.
func New(opts Options) (*Server, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) internalBaseContext(_ net.Listener) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *Server) isIPv6(host string) bool { _ = "STUB: not implemented"; return false }

func (s *Server) getAddress(cfg *config.Config) string { _ = "STUB: not implemented"; return "" }

func (s *Server) getProxyAddress(cfg *config.Config) string { _ = "STUB: not implemented"; return "" }

func (s *Server) refreshURLs() { _ = "STUB: not implemented"; return }

// Service returns the service identified by the given name.
// Panics if no service could be found with the given name.
func (s *Server) Service(name string) Service { _ = "STUB: not implemented"; return *new(Service) }

// LookupService search for a service by its name. If the service
// identified by the given name exists, it is returned with the `true` boolean.
// Otherwise returns `nil` and `false`.
func (s *Server) LookupService(name string) (Service, bool) {
	_ = "STUB: not implemented"
	return *new(Service), false
}

// RegisterService on this server using its name (returned by `Service.Name()`).
// A service's name should be unique.
// `Service.Init(server)` is called on the given service upon registration.
func (s *Server) RegisterService(service Service) { _ = "STUB: not implemented"; return }

// Host returns the hostname and port the server is running on.
func (s *Server) Host() string { _ = "STUB: not implemented"; return "" }

// Port returns the port the server is running on.
func (s *Server) Port() int {
	_ = "STUB: not implemented"

	// BaseURL returns the base URL of your application.
	// If "server.domain" is set in the config, uses it instead
	// of an IP address.
	return 0
}

func (s *Server) BaseURL() string {
	_ = "STUB: not implemented"

	// ProxyBaseURL returns the base URL of your application based on the "server.proxy" configuration.
	// This is useful when you want to generate an URL when your application is served behind a reverse proxy.
	// If "server.proxy.host" configuration is not set, returns the same value as "BaseURL()".
	return ""
}

func (s *Server) ProxyBaseURL() string { _ = "STUB: not implemented"; return "" }

// IsReady returns true if the server has finished initializing and
// is ready to serve incoming requests.
// This operation is concurrently safe.
func (s *Server) IsReady() bool { _ = "STUB: not implemented"; return false }

// RegisterStartupHook to execute some code once the server is ready and running.
// All startup hooks are executed in a single goroutine and in order of registration.
func (s *Server) RegisterStartupHook(hook func(*Server)) { _ = "STUB: not implemented"; return }

// ClearStartupHooks removes all startup hooks.
func (s *Server) ClearStartupHooks() { _ = "STUB: not implemented"; return }

// RegisterShutdownHook to execute some code after the server stopped.
// Shutdown hooks are executed before `Start()` returns and are NOT executed
// in a goroutine, meaning that the shutdown process can be blocked by your
// shutdown hooks. It is your responsibility to implement a timeout mechanism
// inside your hook if necessary.
func (s *Server) RegisterShutdownHook(hook func(*Server)) { _ = "STUB: not implemented"; return }

// ClearShutdownHooks removes all shutdown hooks.
func (s *Server) ClearShutdownHooks() { _ = "STUB: not implemented"; return }

// Config returns the server's config.
func (s *Server) Config() *config.Config { _ = "STUB: not implemented"; return nil }

func (s *Server) HasDB() bool {
	_ = "STUB: not implemented"

	// DB returns the root database instance. Panics if no
	// database connection is set up.
	return false
}

func (s *Server) DB() *gorm.DB { _ = "STUB: not implemented"; return nil }

// Transaction makes it so all DB requests are run inside a transaction.
//
// Returns the rollback function. When you are done, call this function to
// complete the transaction and roll it back. This will also restore the original
// DB so it can be used again out of the transaction.
//
// This is used for tests. This operation is not concurrently safe.
func (s *Server) Transaction(opts ...*sql.TxOptions) func() { _ = "STUB: not implemented"; return nil }

// ReplaceDB manually replace the automatic DB connection.
// If a connection already exists, closes it before discarding it.
// This can be used to create a mock DB in tests. Using this function
// is not recommended outside of tests. Prefer using a custom dialect.
// This operation is not concurrently safe.
func (s *Server) ReplaceDB(dialector gorm.Dialector) error { _ = "STUB: not implemented"; return nil }

// CloseDB close the database connection if there is one.
// Does nothing and returns `nil` if there is no connection.
func (s *Server) CloseDB() error { _ = "STUB: not implemented"; return nil }

// Router returns the root router.
func (s *Server) Router() *Router {
	_ = "STUB: not implemented"

	// Start the server. This operation is blocking and returns when the server is closed.
	return nil
}

func (s *Server) Start() error { _ = "STUB: not implemented"; return nil }

// Notify the shutdown is complete so Stop() can return

// We check if the server is ready to prevent startup hook execution
// if `Serve` returned an error before the goroutine started

// RegisterRoutes runs the given `routeRegistrer` function with this Server and its router.
// The router's regex cache is cleared after the `routeRegistrer` function returns.
// This method should only be called once.
func (s *Server) RegisterRoutes(routeRegistrer func(*Server, *Router)) {
	_ = "STUB: not implemented"
	return
}

// Stop gracefully shuts down the server without interrupting any
// active connections.
//
// `Stop()` does not attempt to close nor wait for hijacked
// connections such as WebSockets. The caller of `Stop` should
// separately notify such long-lived connections of shutdown and wait
// for them to close, if desired. This can be done using shutdown hooks.
//
// If registered, the OS signal channel is closed.
//
// Make sure the program doesn't exit before `Stop()` returns.
//
// After being stopped, a `Server` is not meant to be re-used.
//
// This function can be called from any goroutine and is concurrently safe.
// Calling this function several times is safe. Calls after the first one are no-op.
func (s *Server) Stop() { _ = "STUB: not implemented"; return }

// Start has not been called or Stop has already been called, do nothing

// Wait for stop channel before returning

// RegisterSignalHook creates a channel listening on SIGINT and SIGTERM. When receiving such
// signal, the server is stopped automatically and the listener on these signals is removed.
func (s *Server) RegisterSignalHook() {
	_ = "STUB: not implemented"
	// Sometimes users may not want to have a sigChannel setup
	// also we don't want it in tests
	// users will have to manually call this function if they want the shutdown on signal feature
	return
}

// errLogWriter is a proxy io.Writer that pipes into the server logger.
// This is used so the error logger (type `*log.Logger`) of the underlying
// std HTTP server write to the same logger as the rest of the application.
type errLogWriter struct {
	server *Server
}

func (w errLogWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// ServerFromContext returns the `*goyave.Server` stored in the given context or `nil`.
// This is safe to call using any context retrieved from incoming HTTP requests as this value
// is automatically injected when the server is created.
func ServerFromContext(ctx context.Context) *Server { _ = "STUB: not implemented"; return nil }
