package lang

import (
	"goyave.dev/goyave/v5/util/fsutil"
)

// Languages container for all loaded languages.
//
// This structure is not protected for concurrent usage. Therefore, don't load
// more languages when this instance is expected to receive reads.
type Languages struct {
	languages map[string]*Language
	Default   string
}

// New create a `Languages` with preloaded default language "en-US".
//
// The default language can be replaced by modifying the `Default` field
// in the returned struct.
func New() *Languages { _ = "STUB: not implemented"; return nil }

// LoadAllAvailableLanguages loads every available language directory.
// If the given FS implements `fsutil.WorkingDirFS`, the directory
// used will be "<working directory>/resources/lang".
func (l *Languages) LoadAllAvailableLanguages(fs fsutil.FS) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadDirectory loads every language directory
// in the given directory if it exists.
func (l *Languages) LoadDirectory(fs fsutil.FS, directory string) error {
	_ = "STUB: not implemented"
	return nil
}

// Load a language directory.
//
// Directory structure of a language directory:
//
//	en-UK
//	  ├─ locale.json     (contains the normal language lines)
//	  ├─ rules.json      (contains the validation messages)
//	  └─ fields.json     (contains the field names)
//
// Each file is optional.
func (l *Languages) Load(fs fsutil.FS, language, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Languages) load(fs fsutil.FS, lang string, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLanguage returns a language by its name.
// If the language is not available, returns a dummy language
// that will always return the entry name.
func (l *Languages) GetLanguage(lang string) *Language { _ = "STUB: not implemented"; return nil }

// GetDefault is an alias for `l.GetLanguage(l.Default)`
func (l *Languages) GetDefault() *Language { _ = "STUB: not implemented"; return nil }

// IsAvailable returns true if the language is available.
func (l *Languages) IsAvailable(lang string) bool { _ = "STUB: not implemented"; return false }

// GetAvailableLanguages returns a slice of all loaded languages.
// This can be used to generate different routes for all languages
// supported by your applications.
//
//	/en/products
//	/fr/produits
//	...
func (l *Languages) GetAvailableLanguages() []string { _ = "STUB: not implemented"; return nil }

// DetectLanguage detects the language to use based on the given lang string.
// The given lang string can use the HTTP "Accept-Language" header format.
//
// If "*" is provided, the default language will be used.
// If multiple languages are given, the first available language will be used,
// and if none are available, the default language will be used.
// If no variant is given (for example "en"), the first available variant will be used.
func (l *Languages) DetectLanguage(lang string) *Language { _ = "STUB: not implemented"; return nil }

// Accept anything, so return default language

// TODO priority for languages? The "first available variant" is random because keys are not ordered.
// Ordering alphabetically won't always produce the desired result (e.g. en-UK < en-US)
// Can create a slice of language names (so the order will be the order in which the languages have been added)

// Get a language line.
//
// For validation rules messages and field names, use a dot-separated path:
//   - "validation.rules.<rule_name>"
//   - "validation.fields.<field_name>"
//
// For normal lines, just use the name of the line. Note that if you have
// a line called "validation", it won't conflict with the dot-separated paths.
//
// If not found, returns the exact "line" argument.
//
// The placeholders parameter is a variadic associative slice of placeholders and their
// replacement. In the following example, the placeholder ":username" will be replaced
// with the Name field in the user struct.
//
//	lang.Get("en-US", "greetings", ":username", user.Name)
func (l *Languages) Get(lang string, line string, placeholders ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func readLangFile(fs fsutil.FS, path string, dest any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func mergeLang(dst *Language, src *Language) { _ = "STUB: not implemented"; return }
