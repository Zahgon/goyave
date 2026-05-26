package auth

import (
	"io/fs"
	"reflect"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/lang"
)

const (
	// JWTServiceName identifier for the `JWTService`.
	JWTServiceName = "goyave.jwt"
)

// ExtraJWTClaims when using the built-in `JWTAuthenticator`, this
// key can be used to retrieve the JWT claims in the request's `Extra`.
type ExtraJWTClaims struct{}

func init() {
	config.Register("auth.jwt.expiry", config.Entry{
		Value:            300,
		Type:             reflect.Int,
		IsSlice:          false,
		AuthorizedValues: []any{},
	})
	registerKeyConfigEntry("auth.jwt.secret")
	registerKeyConfigEntry("auth.jwt.rsa.public")
	registerKeyConfigEntry("auth.jwt.rsa.private")
	registerKeyConfigEntry("auth.jwt.ecdsa.public")
	registerKeyConfigEntry("auth.jwt.ecdsa.private")
}

func registerKeyConfigEntry(name string) { _ = "STUB: not implemented"; return }

// JWTService providing signature keys cache and JWT generation.
//
// This service is identified by `auth.JWTServiceName`.
type JWTService struct {
	fs     fs.FS
	config *config.Config
	cache  sync.Map
}

// NewJWTService create a new `JWTService` with the given config and file system.
// The file system is used to get the signing keys.
func NewJWTService(config *config.Config, fs fs.FS) *JWTService {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the service.
func (s *JWTService) Name() string { _ = "STUB: not implemented"; return "" }

// GenerateToken generate a new JWT.
// The token is created using the HMAC SHA256 method and signed using
// the `auth.jwt.secret` config entry.
// The token is set to expire in the amount of seconds defined by
// the `auth.jwt.expiry` config entry.
//
// The generated token will contain the following claims:
//   - `sub`: has the value of the `id` parameter
//   - `nbf`: "Not before", the current timestamp is used
//   - `exp`: "Expiry", the current timestamp plus the `auth.jwt.expiry` config entry.
func (s *JWTService) GenerateToken(username any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GenerateTokenWithClaims generates a new JWT with custom claims.
// The token is set to expire in the amount of seconds defined by
// the `auth.jwt.expiry` config entry.
// Depending on the given signing method, the following configuration entries
// will be used:
//   - RSA: `auth.jwt.rsa.private`: path to the private PEM-encoded RSA key.
//   - ECDSA: `auth.jwt.ecdsa.private`: path to the private PEM-encoded ECDSA key.
//   - HMAC: `auth.jwt.secret`: HMAC secret
//
// The generated token will also contain the following claims:
//   - `nbf`: "Not before", the current timestamp is used
//   - `exp`: "Expiry", the current timestamp plus the `auth.jwt.expiry` config entry.
//
// `nbf` and `exp` can be overridden if they are set in the `claims` parameter.
func (s *JWTService) GenerateTokenWithClaims(claims jwt.MapClaims, signingMethod jwt.SigningMethod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Not Before
// Expiry

// GetKey load a JWT signature key from the config.
// List of `entry` parameter possible values:
//
//   - `auth.jwt.rsa.public`
//   - `auth.jwt.rsa.private`
//   - `auth.jwt.ecdsa.public`
//   - `auth.jwt.ecdsa.private`
//   - `auth.jwt.secret`
//
// To optimize subsequent requests and avoid IO for keys that are stored on the
// disk, the keys are cached.
func (s *JWTService) GetKey(entry string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// GetPrivateKey loads the private key that corresponds to the given `signingMethod`.
func (s *JWTService) GetPrivateKey(signingMethod jwt.SigningMethod) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// JWTAuthenticator implementation of Authenticator using a JSON Web Token.
//
// The T parameter represents the user DTO and should not be a pointer.
type JWTAuthenticator[T any] struct {
	goyave.Component

	service *JWTService

	UserService UserService[T]

	// SigningMethod expected by this authenticator when parsing JWT.
	// Defaults to HMAC.
	SigningMethod jwt.SigningMethod

	// ClaimName the name of the claim used to retrieve the user.
	// Defaults to "sub".
	ClaimName string

	// Optional defines if the authenticator allows requests that
	// don't provide credentials. Handlers should therefore check
	// if `request.User` is not `nil` before accessing it.
	Optional bool
}

// NewJWTAuthenticator create a new authenticator for the JSON Web Token authentication flow.
//
// The T parameter represents the user DTO and should not be a pointer.
func NewJWTAuthenticator[T any](userService UserService[T]) *JWTAuthenticator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Init the authenticator. Automatically registers the `JWTService` if not already registered,
// using `osfs.FS` as file system for the keys.
func (a *JWTAuthenticator[T]) Init(server *goyave.Server) { _ = "STUB: not implemented"; return }

// Authenticate fetch the user corresponding to the token
// found in the given request and returns it.
// If no user can be authenticated, returns an error.
//
// If the token is valid and has claims, those claims will be added to `request.Extra` with the key "jwt_claims".
func (a *JWTAuthenticator[T]) Authenticate(request *goyave.Request) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *JWTAuthenticator[T]) keyFunc(token *jwt.Token) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (a *JWTAuthenticator[T]) makeError(language *lang.Language, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *JWTAuthenticator[T]) Scheme() string { _ = "STUB: not implemented"; return "" }
