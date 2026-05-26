package config

import (
	"io/fs"
	"sync"
)

type object map[string]any

type readFunc func(string) (object, error)

// Config structure holding a configuration that should be used for a single
// instance of `goyave.Server`.
//
// This structure is not protected for safe concurrent access in order to increase
// performance. Therefore, you should never use the `Set()` function when the configuration
// is in use by an already running server.
type Config struct {
	config object
}

// Error returned when the configuration could not
// be loaded or is invalid.
// Can be unwraped to get the original error.
type Error struct {
	err error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

type loader struct {
	defaults object
	mu       sync.RWMutex
}

var defaultLoader = &loader{
	defaults: configDefaults,
}

// Register a new config entry and its validation.
//
// Each module should register its config entries in an "init()"
// function, even if they don't have a default value, in order to
// ensure they will be validated.
// Each module should use its own category and use a name both expressive
// and unique to avoid collisions.
// For example, the "auth" package registers, among others, "auth.basic.username"
// and "auth.jwt.expiry", thus creating a category for its package, and two subcategories
// for its features.
//
// To register an entry without a default value (only specify how it
// will be validated), set "Entry.Value" to "nil".
//
// Panics if an entry already exists for this key and is not identical to the
// one passed as parameter of this function. On the other hand, if the entries
// are identical, no conflict is expected so the configuration is left in its
// current state.
func Register(key string, entry Entry) { _ = "STUB: not implemented"; return }

func (l *loader) register(key string, entry Entry) { _ = "STUB: not implemented"; return }

func (l *loader) loadFrom(fs fs.FS, path string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *loader) loadJSON(cfg string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *loader) load(readFunc readFunc, source string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load loads the config.json file in the current working directory.
// If the "GOYAVE_ENV" env variable is set, the config file will be picked like so:
//   - "production": "config.production.json"
//   - "test": "config.test.json"
//   - By default: "config.json"
func Load() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadDefault loads default config.
func LoadDefault() *Config { _ = "STUB: not implemented"; return nil }

// LoadFrom loads a config file from the given path.
func LoadFrom(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadJSON load a configuration file from raw JSON. Can be used in combination with
// Go's embed directive.
//
//	var (
//		//go:embed config.json
//		cfgJSON string
//	)
//
//	func main() {
//		cfg, err := config.LoadJSON(cfgJSON)
//		if err != nil {
//			fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
//			os.Exit(1)
//		}
//
//		server, err := goyave.New(goyave.Options{Config: cfg})
//		if err != nil {
//			fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
//			os.Exit(1)
//		}
//
//		// ...
//	}
func LoadJSON(cfg string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func getConfigFilePath() string { _ = "STUB: not implemented"; return "" }

func (l *loader) readConfigFile(filesystem fs.FS, file string) (o object, err error) {
	_ = "STUB: not implemented"
	return *new(object), nil
}

func (l *loader) readString(str string) (object, error) {
	_ = "STUB: not implemented"
	return *new(object), nil
}

// walk the config using the key. Returns the deepest category, the entry key
// with its path stripped ("app.name" -> "name") and true if the entry already
// exists, false if it's not registered.
func walk(currentCategory object, key string) (object, string, bool) {
	_ = "STUB: not implemented"
	return *new(object), "", false
}

// If categories are missing, create them

// Entry doesn't exist and is not registered

// Entry exists

// createMissingCategories based on the key path, starting at the given index.
// Doesn't create anything is not needed.
// Returns the deepest category created, or the provided object if nothing has
// been created.
func createMissingCategories(currentCategory object, path string) object {
	_ = "STUB: not implemented"
	return *new(object)
}

func override(src object, dst object) error { _ = "STUB: not implemented"; return nil }

// Conflict: destination is not a category

// Conflict: override category with an entry

// If entry doesn't exist (and is not registered),
// register it with the type of the type given here
// and "any" authorized values.

func (o object) validate(key string) error { _ = "STUB: not implemented"; return nil }

// Get a config entry using a dot-separated path.
// Panics if the entry doesn't exist.
func (c *Config) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (c *Config) get(key string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// nil means unset

// GetString a config entry as string.
// Panics if entry is not a string or if it doesn't exist.
func (c *Config) GetString(key string) string { _ = "STUB: not implemented"; return "" }

// GetBool a config entry as bool.
// Panics if entry is not a bool or if it doesn't exist.
func (c *Config) GetBool(key string) bool { _ = "STUB: not implemented"; return false }

// GetInt a config entry as int.
// Panics if entry is not an int or if it doesn't exist.
func (c *Config) GetInt(key string) int { _ = "STUB: not implemented"; return 0 }

// GetFloat a config entry as float64.
// Panics if entry is not a float64 or if it doesn't exist.
func (c *Config) GetFloat(key string) float64 { _ = "STUB: not implemented"; return 0 }

// GetStringSlice a config entry as []string.
// Panics if entry is not a string slice or if it doesn't exist.
func (c *Config) GetStringSlice(key string) []string { _ = "STUB: not implemented"; return nil }

// GetBoolSlice a config entry as []bool.
// Panics if entry is not a bool slice or if it doesn't exist.
func (c *Config) GetBoolSlice(key string) []bool { _ = "STUB: not implemented"; return nil }

// GetIntSlice a config entry as []int.
// Panics if entry is not an int slice or if it doesn't exist.
func (c *Config) GetIntSlice(key string) []int { _ = "STUB: not implemented"; return nil }

// GetFloatSlice a config entry as []float64.
// Panics if entry is not a float slice or if it doesn't exist.
func (c *Config) GetFloatSlice(key string) []float64 { _ = "STUB: not implemented"; return nil }

// Has check if a config entry exists.
func (c *Config) Has(key string) bool { _ = "STUB: not implemented"; return false }

// Set a config entry.
// The change is temporary and will not be saved for next boot.
// Use "nil" to unset a value.
//
//   - A category cannot be replaced with an entry.
//   - An entry cannot be replaced with a category.
//   - New categories can be created with they don't already exist.
//   - New entries can be created if they don't already exist. This new entry
//     will be subsequently validated using the type of its initial value and
//     have an empty slice as authorized values (meaning it can have any value of its type)
//
// Panics and revert changes in case of error.
//
// This operation is not concurrently safe and should not be used when the configuration
// is in use by an already running server.
func (c *Config) Set(key string, value any) { _ = "STUB: not implemented"; return }
