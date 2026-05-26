package lang

type validationLines struct {
	// Default messages for rules
	rules map[string]string

	// Field names translations
	fields map[string]string
}

// Language represents a full Language.
type Language struct {
	lines      map[string]string
	validation validationLines
	name       string
}

// Name returns the name of the language. For example "en-US".
func (l *Language) Name() string { _ = "STUB: not implemented"; return "" }

func (l *Language) clone() *Language { _ = "STUB: not implemented"; return nil }

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
//	lang.Get("greetings", ":username", user.Name)
func (l *Language) Get(line string, placeholders ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func convertEmptyLine(entry, line string, placeholders []string) string {
	_ = "STUB: not implemented"
	return ""
}

func processPlaceholders(message string, values []string) string {
	_ = "STUB: not implemented"
	return ""
}
