package httputil

import (
	"regexp"
)

// HeaderValue represent a value and its quality value (priority)
// in a multi-values HTTP header.
type HeaderValue struct {
	Value    string
	Priority float64
}

var qualityValueRegex = regexp.MustCompile(`^q=([01]\.[0-9]{1,3})$`)

// ParseMultiValuesHeader parses multi-values HTTP headers, taking the
// quality values into account. The result is a slice of values sorted
// according to the order of priority.
//
// The input is trimmed. If the input is empty, returns an empty slice.
//
// See: https://developer.mozilla.org/en-US/docs/Glossary/Quality_values
//
// For the following header:
//
//	"text/html,text/*;q=0.5,*/*;q=0.7"
//
// returns
//
//	[{text/html 1} {*/* 0.7} {text/* 0.5}]
func ParseMultiValuesHeader(header string) []HeaderValue { _ = "STUB: not implemented"; return nil }

// Parse priority

// Priority set to 0 if the quality value cannot be parsed
