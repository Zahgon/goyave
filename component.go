package goyave

import (
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/lang"
	"goyave.dev/goyave/v5/slog"
)

// Composable defines all the functions every component of the presentation
// layer (HTTP/REST) of your application must implement. These functions are
// accessors to the essential server resources.
// A component can be parent of several sub-components.
type Composable interface {
	Init(server *Server)
	Server() *Server
	DB() *gorm.DB
	Config() *config.Config
	Lang() *lang.Languages
	Logger() *slog.Logger
	Service(name string) Service
	LookupService(name string) (Service, bool)
}

// Registrer qualifies a controller that registers its routes itself.
// It is required for controllers to implement this interface if you want
// to use `router.Controller()`.
type Registrer interface {
	Composable
	RegisterRoutes(router *Router)
}

// Component base implementation of `Composable` to easily make a
// custom component using structure composition.
type Component struct {
	server *Server
}

var _ Composable = (*Component)(nil)

// Init the component using the given server.
func (c *Component) Init(server *Server) {
	_ = "STUB: not implemented"

	// Server returns the parent server.
	return
}

func (c *Component) Server() *Server {
	_ = "STUB: not implemented"

	// Service returns the service identified by the given name.
	// Panics if no service could be found with the given name.
	return nil
}

func (c *Component) Service(name string) Service { _ = "STUB: not implemented"; return *new(Service) }

// LookupService search for a service by its name. If the service
// identified by the given name exists, it is returned with the `true` boolean.
// Otherwise returns `nil` and `false`.
func (c *Component) LookupService(name string) (Service, bool) {
	_ = "STUB: not implemented"
	return *new(Service), false
}

// Logger returns the server's logger.
func (c *Component) Logger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// DB returns the root database instance. Panics if no
// database connection is set up.
func (c *Component) DB() *gorm.DB { _ = "STUB: not implemented"; return nil }

// Config returns the server's config.
func (c *Component) Config() *config.Config { _ = "STUB: not implemented"; return nil }

// Lang returns the languages loaded by the server.
func (c *Component) Lang() *lang.Languages { _ = "STUB: not implemented"; return nil }
