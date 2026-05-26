package goyave

import (
	"io/fs"
)

func staticHandler(fs fs.StatFS, download bool) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

func cleanStaticPath(fs fs.StatFS, file string) string { _ = "STUB: not implemented"; return "" }

func checkStaticPath(path string) bool { _ = "STUB: not implemented"; return false }

// Force leading slash if the path is not empty

func isSlash(r rune) bool { _ = "STUB: not implemented"; return false }
