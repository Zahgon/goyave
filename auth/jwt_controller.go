package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/validation"
)

// TokenFunc is the function used by JWTController to generate tokens
// during login process.
type TokenFunc[T any] func(request *goyave.Request, user *T) (string, error)

// JWTController controller adding a login route returning a JWT for quick prototyping.
//
// The T parameter represents the user DTO and should not be a pointer. The DTO used should be
// different from the DTO returned to clients as a response because it needs to contain the user's password.
type JWTController[T any] struct { // TODO refresh token
	goyave.Component

	jwtService *JWTService

	UserService UserService[T]

	// SigningMethod used to generate the token using the default
	// TokenFunc. By default, uses `jwt.SigningMethodHS256`.
	SigningMethod jwt.SigningMethod

	// The function generating the token on a successful authentication.
	// Defaults to a JWT signed with HS256 and containing the username as the
	// "sub" claim.
	TokenFunc TokenFunc[T]

	// UsernameRequestField the name of the request's body field
	// used as username in the authentication process.
	// Defaults to "username"
	UsernameRequestField string
	// PasswordRequestField the name of the request's body field
	// used as password in the authentication process.
	// Defaults to "password"
	PasswordRequestField string
	// PasswordField the name of T's struct field that holds the user's hashed password.
	// It will be used to compare the password hash with the user input.
	PasswordField string
}

// NewJWTController create a new JWTController that registers a login route returning a JWT for quick prototyping.
//
// The `passwordField` corresponds to the name of T's struct field that holds the user's hashed password.
// It will be used to compare the password hash with the user input.
func NewJWTController[T any](userService UserService[T], passwordField string) *JWTController[T] {
	_ = "STUB: not implemented"
	return nil
}

// Init the controller. Automatically registers the `JWTService` if not already registered,
// using `osfs.FS` as file system for the signing keys.
func (c *JWTController[T]) Init(server *goyave.Server) { _ = "STUB: not implemented"; return }

// RegisterRoutes register the "/login" route (with validation) on the given router.
func (c *JWTController[T]) RegisterRoutes(router *goyave.Router) { _ = "STUB: not implemented"; return }

func (c *JWTController[T]) validationRules(_ *goyave.Request) validation.RuleSet {
	_ = "STUB: not implemented"
	return *new(validation.RuleSet)
}

// Login POST handler for token-based authentication.
// Creates a new token for the user authenticated with the body fields
// defined in the controller and returns it as a response.
// The password is checked using bcrypt.
func (c *JWTController[T]) Login(response *goyave.Response, request *goyave.Request) {
	_ = "STUB: not implemented"
	return
}

func (c *JWTController[T]) defaultTokenFunc(r *goyave.Request, _ *T) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
