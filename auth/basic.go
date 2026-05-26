package auth

import (
	"reflect"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/config"
)

// BasicAuthenticator implementation of Authenticator with the Basic
// authentication method.
//
// The T parameter represents the user DTO and should not be a pointer. The DTO used should be
// different from the DTO returned to clients as a response because it needs to contain the user's password.
type BasicAuthenticator[T any] struct {
	goyave.Component

	UserService UserService[T]

	// PasswordField the name of T's struct field that holds the user's hashed password.
	// It will be used to compare the password hash with the user input.
	PasswordField string

	// Optional defines if the authenticator allows requests that
	// don't provide credentials. Handlers should therefore check
	// if `request.User` is not `nil` before accessing it.
	Optional bool
}

// NewBasicAuthenticator create a new authenticator for the Basic authentication flow.
//
// The T parameter represents the user DTO and should not be a pointer. The DTO used should be
// different from the DTO returned to clients as a response because it needs to contain the user's password.
//
// The `passwordField` corresponds to the name of T's struct field that holds the user's hashed password.
// It will be used to compare the password hash with the user input.
func NewBasicAuthenticator[T any](userService UserService[T], passwordField string) *BasicAuthenticator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Authenticate fetch the user corresponding to the credentials
// found in the given request and returns it.
// If no user can be authenticated, returns an error.
// The password is checked using bcrypt.
func (a *BasicAuthenticator[T]) Authenticate(request *goyave.Request) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *BasicAuthenticator[T]) Scheme() string {
	_ = "STUB: not implemented"

	// --------------------------------------------
	return ""
}

func init() {
	config.Register("auth.basic.username", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []any{},
	})
	config.Register("auth.basic.password", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []any{},
	})
}

// BasicUser a simple user for config-based basic authentication.
type BasicUser struct {
	Name string
}

// ConfigBasicAuthenticator implementation of Authenticator with the Basic
// authentication method, using username and password from the configuration.
type ConfigBasicAuthenticator struct {
	goyave.Component
}

// Authenticate check if the request basic auth header matches the
// "auth.basic.username" and "auth.basic.password" config entries.
func (a *ConfigBasicAuthenticator) Authenticate(request *goyave.Request) (*BasicUser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ConfigBasicAuthenticator) Scheme() string {
	_ = "STUB: not implemented"

	// ConfigBasicAuth create a new authenticator middleware for
	// config-based Basic authentication. On auth success, the request
	// user is set to a `*BasicUser`.
	// The user is authenticated if the "auth.basic.username" and "auth.basic.password" config entries
	// match the request's Authorization header.
	return ""
}

func ConfigBasicAuth() *Handler[BasicUser] { _ = "STUB: not implemented"; return nil }

// ConfigBasicAuthWithRealm is the same as ConfigBasicAuth but with a custom realm description.
// The realm describes the protected area and is returned in the `WWW-Authenticate` header
// when the authentication fails.
func ConfigBasicAuthWithRealm(realm string) *Handler[BasicUser] {
	_ = "STUB: not implemented"
	return nil
}
