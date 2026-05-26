package goyave

import (
	"regexp"
)

// parameterizable represents a route or router accepting
// parameters in its URI.
type parameterizable struct {
	regex      *regexp.Regexp
	parameters []string
}

// compileParameters parse the route parameters and compiles their regexes if needed.
// If "ends" is set to true, the generated regex ends with "$", thus set "ends" to true
// if you're compiling route parameters, set to false if you're compiling router parameters.
func (p *parameterizable) compileParameters(uri string, ends bool, regexCache map[string]*regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

// Final regex will never be larger than src uri + 2 (for ^ and $)
// Make initial alloc to avoid the need for realloc

// default pattern

// Skip closing braces

// braceIndices returns the first level curly brace indices from a string.
// Returns an error in case of unbalanced braces.
func (p *parameterizable) braceIndices(s string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// makeParameters from a regex match and the given parameter names.
// The match parameter is expected to contain only the capturing groups.
//
// Given ["/product/33/param", "33", "param"] ["id", "name"]
// The returned map will be ["id": "33", "name": "param"]
func (p *parameterizable) makeParameters(match []string, names []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// GetParameters returns the URI parameters' names (in order of appearance).
func (p *parameterizable) GetParameters() []string { _ = "STUB: not implemented"; return nil }
