package walk

import (
	"bufio"
	"strings"

	"github.com/samber/lo"
)

// PathType type of the element being explored.
type PathType int

// FoundType adds extra information about not found elements whether
// what's not found is their parent or themselves.
type FoundType int

const (
	// PathTypeElement the explored element is used as a final element (leaf).
	PathTypeElement PathType = iota

	// PathTypeArray the explored element is used as an array and not a final element.
	// All elements in the array will be explored using the next Path.
	PathTypeArray

	// PathTypeObject the explored element is used as an object (`map[string]any`)
	// and not a final element.
	PathTypeObject
)

const (
	// Found indicates the element could be found.
	Found FoundType = iota
	// ParentNotFound indicates one of the parents of the element could no be found.
	ParentNotFound
	// ElementNotFound indicates all parents of the element were found but the element
	// itself could not.
	ElementNotFound
)

var wildcard = lo.ToPtr("*")

var (
	// EscapeChars the list of characters that can be escaped using a backslash `\` when
	// parsing a Path. This map is read-only.
	EscapeChars = map[rune]struct{}{
		'*':  {},
		'[':  {},
		']':  {},
		'.':  {},
		'\\': {},
	}

	escapeRemover = strings.NewReplacer(
		`\*`, `*`,
		`\[`, `[`,
		`\]`, `]`,
		`\.`, `.`,
		`\\`, `\`,
	)

	escapeReplacer = strings.NewReplacer(
		`*`, `\*`,
		`[`, `\[`,
		`]`, `\]`,
		`.`, `\.`,
		`\`, `\\`,
	)
)

// Path allows for complex untyped data structure exploration.
// An instance of this structure represents a step in exploration.
// Items NOT having `PathTypeElement` as a `Type` are expected to have a non-nil `Next`.
type Path struct {
	Next  *Path
	Index *int
	Name  *string
	Type  PathType
}

// Context information sent to walk function.
type Context struct {
	Value  any
	Parent any       // Either map[string]any or a slice
	Path   *Path     // Exact Path to the current element
	Name   string    // Name of the current element
	Index  int       // If parent is a slice, the index of the current element in the slice, else -1
	Found  FoundType // True if the path could not be completely explored

	stop bool
}

// Break when called, indicates the path walker to stop.
// This means the current call of the callback function will be the last.
func (c *Context) Break() {
	_ = "STUB: not implemented"

	// Walk this path and execute the given callback for each matching element. Elements are final,
	// meaning they are the deepest explorable element using this path.
	// Only `map[string]any` and n-dimensional slices parents are supported.
	// The given "f" function is executed for each final element matched. If the path
	// cannot be completed because the step's name doesn't exist in the currently explored map,
	// the function will be executed as well, with a the `Context`'s `NotFound` field set to `true`.
	return
}

func (p *Path) Walk(currentElement any, f func(*Context)) { _ = "STUB: not implemented"; return }

func (p *Path) walk(currentElement any, parent any, index int, trackPath *Path, lastPathElement *Path, f func(*Context)) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Path) walkArray(element any, parent any, index int, trackPath *Path, lastPathElement *Path, f func(*Context)) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Path) outOfBounds(length int) bool { _ = "STUB: not implemented"; return false }

func (p *Path) completePath(lastPathElement *Path) { _ = "STUB: not implemented"; return }

func newNotFoundContext(parent any, path *Path, name *string, index int, found FoundType) *Context {
	_ = "STUB: not implemented"
	return nil
}

// First returns the first final element matched by the Path.
// Note that the returned Context may indicate that the value could
// not be found, so you should always check `Context.Found` before using
// `Context.Value`.
//
// Bear in mind that map iteration order is not guaranteed. Using paths containing
// wildcards `*` will not always yield the same result.
func (p *Path) First(currentElement any) *Context { _ = "STUB: not implemented"; return nil }

// HasArray returns true if a least one step in the path involves an array.
func (p *Path) HasArray() bool { _ = "STUB: not implemented"; return false }

// LastParent returns the last step in the path that is not a PathTypeElement, excluding
// the first step in the path, or nil.
func (p *Path) LastParent() *Path { _ = "STUB: not implemented"; return nil }

// Tail returns the last step in the path.
func (p *Path) Tail() *Path { _ = "STUB: not implemented"; return nil }

// Depth returns the depth of the path. For each step in the path, increments the depth by one.
func (p *Path) Depth() uint { _ = "STUB: not implemented"; return 0 }

// Truncate returns a clone of the n first steps of the path so the returned path's depth
// equals the given depth.
func (p *Path) Truncate(depth uint) *Path { _ = "STUB: not implemented"; return nil }

// Clone returns a deep clone of this Path.
func (p *Path) Clone() *Path { _ = "STUB: not implemented"; return nil }

// IsWildcard returns true if the path has unescaped "*" as Name.
func (p *Path) IsWildcard() bool { _ = "STUB: not implemented"; return false }

// String returns a string representation of the Path.
// The result contains the escape characters, if any.
func (p *Path) String() string { _ = "STUB: not implemented"; return "" }

// UnescapedString returns a string representation of the Path
// without the escape characters.
func (p *Path) UnescapedString() string { _ = "STUB: not implemented"; return "" }

func (p *Path) toString(showEscapeChars bool) string { _ = "STUB: not implemented"; return "" }

// Unescape remove escape characters from a path without parsing it.
func Unescape(path string) string { _ = "STUB: not implemented"; return "" }

// setAllMissingIndexes set Index to -1 for all `PathTypeArray` steps in this path.
func (p *Path) setAllMissingIndexes() { _ = "STUB: not implemented"; return }

// Parse transform given path string representation into usable Path.
//
// The wildcard `*` can be used to match all keys of an object. It is only effective
// if it is an entire path segment. For example, the `*` in `field*name` won't be
// considered a wildcard and only the literal field will match.
//
// Special characters defined in the `Escape` map (by default `*`, `[`, `]`, `.` and `\`)
// can be escaped using a backslack `\`.
//
// Example paths:
//
//	name
//	object.field
//	object.subobject.field
//	object.*
//	object.\*
//	object.array[]
//	object.arrayOfObjects[].field
//	[]
//	[].field
//	field*name
//	object.field\[]
//	object.field\[text\]
//	abc\[\]def
//	path\\to\\element
func Parse(p string) (*Path, error) { _ = "STUB: not implemented"; return nil, nil }

// MustParse is the same as `Parse` but panics if there is an error.
func MustParse(p string) *Path { _ = "STUB: not implemented"; return nil }

// Depth calculate the path's depth without parsing it.
func Depth(p string) uint { _ = "STUB: not implemented"; return 0 }

func createPathScanner(path string) *bufio.Scanner { _ = "STUB: not implemented"; return nil }

// Skip the next character

func checkSyntax(r rune, next rune) error { _ = "STUB: not implemented"; return nil }

func removeEscapeChars(t string) string { _ = "STUB: not implemented"; return "" }
