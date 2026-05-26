package fsutil

import (
	"io"
	"io/fs"
)

var contentTypeByExtension = map[string]string{
	".ai":     "application/postscript",
	".apk":    "application/vnd.android.package-archive",
	".apng":   "image/apng",
	".avif":   "image/avif",
	".bin":    "application/octet-stream",
	".bmp":    "image/bmp",
	".com":    "application/octet-stream",
	".doc":    "application/msword",
	".css":    "text/css",
	".csv":    "text/csv",
	".docx":   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".ehtml":  "text/html",
	".eml":    "message/rfc822",
	".eps":    "application/postscript",
	".exe":    "application/octet-stream",
	".flac":   "audio/flac",
	".gif":    "image/gif",
	".gz":     "application/gzip",
	".htm":    "text/html",
	".html":   "text/html",
	".ico":    "image/vnd.microsoft.icon",
	".ics":    "text/calendar",
	".jfif":   "image/jpeg",
	".jpg":    "image/jpeg",
	".jpeg":   "image/jpeg",
	".js":     "text/javascript",
	".jsonld": "application/ld+json",
	".json":   "application/json",
	".m4a":    "audio/mp4",
	".mjs":    "text/javascript",
	".mp3":    "audio/mpeg",
	".mp4":    "video/mp4",
	".mpeg":   "audio/mpeg",
	".ods":    "application/vnd.oasis.opendocument.spreadsheet",
	".odt":    "application/vnd.oasis.opendocument.text",
	".oga":    "audio/ogg",
	".ogv":    "video/ogg",
	".ogx":    "application/ogg",
	".opus":   "audio/ogg",
	".otf":    "font/otf",
	".png":    "image/png",
	".pdf":    "application/pdf",
	".pjp":    "image/jpeg",
	".pjpeg":  "image/jpeg",
	".ppt":    "application/vnd.ms-powerpoint",
	".pptx":   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".ps":     "application/postscript",
	".rdf":    "application/rdf+xml",
	".rtf":    "application/rtf",
	".sh":     "application/x-sh",
	".shtml":  "text/html",
	".svg":    "image/svg+xml",
	".tar":    "application/x-tar",
	".text":   "text/plain",
	".tif":    "image/tiff",
	".tiff":   "image/tiff",
	".ts":     "video/mp2t",
	".ttf":    "font/ttf",
	".txt":    "text/plain",
	".vtt":    "text/vtt",
	".wasm":   "application/wasm",
	".wav":    "audio/wav",
	".weba":   "audio/webm",
	".webm":   "audio/webm",
	".webp":   "image/webp",
	".woff":   "font/woff",
	".woff2":  "font/woff2",
	".xbl":    "text/xml",
	".xbm":    "image/x-xbitmap",
	".xht":    "application/xhtml+xml",
	".xhtml":  "application/xhtml+xml",
	".xls":    "application/vnd.ms-excel",
	".xlsx":   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".xml":    "application/xml",
	".xsl":    "text/xml",
	".zip":    "application/zip",
	".7z":     "application/x-7z-compressed",
}

// AddExtensionType set the MIME type associated with the given extension.
// The extension should begin with a dot (e.g.: ".html").
// The mimeType should not include the charset parameter nd be written in lowercase.
//
// Passing an extension that is already registered overrides the previous value.
//
// This function is not safe for concurrent use.
func AddExtensionType(ext, mimeType string) error { _ = "STUB: not implemented"; return nil }

// GetFileExtension returns the last part of a file name, without the leading dot.
// If the file doesn't have an extension, returns an empty string.
// For files with multiple extensions like `.tar.gz`, only `gz` is returned.
func GetFileExtension(filename string) string { _ = "STUB: not implemented"; return "" }

// GetMIMEType get the mime type and size of the given file.
// This function opens the file, stats it and calls `fsutil.DetectContentType`.
// If the file is empty (size of 0), the content-type will be detected using `fsutil.DetectContentTypeByExtension`.
// If a specific MIME type cannot be determined, returns "application/octet-stream" as a fallback.
func GetMIMEType(filesystem fs.FS, file string) (contentType string, size int64, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// DetectContentType by sniffing the first 512 bytes of the given reader using `http.DetectContentType`.
//
// If the detected content type is `"application/octet-stream"` or `"text/plain"`, this function will attempt to
// find a more precise one using `fsutil.DetectContentTypeByExtension`, unless `fileName` is empty.
// If the detected content type is `"text/xml"` or `"application/xml"`, this function promotes it to
// `"image/svg+xml"` only if the content signature indicates SVG.
// The header parameter is retained (e.g: `charset=utf-8`).
//
// If there is no error, this function always returns a valid MIME type. If it cannot determine a more specific one,
// it returns `"application/octet-stream"`.
//
// If the given reader implements `io.Seeker`, the reader's offset is reset to the start.
func DetectContentType(r io.Reader, fileName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func hasSVGSignature(buffer []byte) bool { _ = "STUB: not implemented"; return false }

// Skip optional XML tag and possible comments

func detectContentTypeByExtension(fileName, contentType string) string {
	_ = "STUB: not implemented"
	return ""
}

// Keep the "charset" arguments

// DetectContentTypeByExtension returns a MIME type associated with the extension (suffix) of the given file name.
// Note that this function should be used as a fallback if sniffing using `fsutil.DetectContentType`
// isn't possible due to:
//   - the file being empty (size of 0)
//   - the file requiring to be read after the sniffing but its reader doesn't implement `io.Seeker` for rewinding.
//
// The local database covers the most common MIME types.
// If the extension is not known to the `fsutil` package, returns `"application/octet-stream"`.
func DetectContentTypeByExtension(fileName string) string { _ = "STUB: not implemented"; return "" }

// FileExists returns true if the file at the given path exists and is readable.
// Returns false if the given file is a directory.
func FileExists(fs fs.StatFS, file string) bool { _ = "STUB: not implemented"; return false }

// IsDirectory returns true if the file at the given path exists, is a directory and is readable.
func IsDirectory(fs fs.StatFS, path string) bool { _ = "STUB: not implemented"; return false }

func timestampFileName(name string) string { _ = "STUB: not implemented"; return "" }

// An FS provides access to a hierarchical file system
// and implements `io/fs`'s `FS`, `ReadDirFS` and `StatFS` interfaces.
type FS interface {
	fs.ReadDirFS
	fs.StatFS
}

// A WorkingDirFS is a file system with a `Getwd()` method.
type WorkingDirFS interface {
	// Getwd returns a rooted path name corresponding to the
	// current directory. If the current directory can be
	// reached via multiple paths (due to symbolic links),
	// Getwd may return any one of them.
	Getwd() (dir string, err error)
}

// A MkdirFS is a file system with a `Mkdir()` and a `MkdirAll()` methods.
type MkdirFS interface {
	// MkdirAll creates a directory named path,
	// along with any necessary parents, and returns `nil`,
	// or else returns an error.
	// The permission bits perm (before umask) are used for all
	// directories that `MkdirAll` creates.
	// If path is already a directory, `MkdirAll` does nothing
	// and returns `nil`.
	MkdirAll(path string, perm fs.FileMode) error

	// Mkdir creates a new directory with the specified name and permission
	// bits (before umask).
	// If there is an error, it will be of type `*PathError`.
	Mkdir(path string, perm fs.FileMode) error
}

// A WritableFS is a file system with a `OpenFile()` method.
type WritableFS interface {
	// OpenFile is the generalized open call. It opens the named file with specified flag
	// (`O_RDONLY` etc.). If the file does not exist, and the `O_CREATE` flag
	// is passed, it is created with mode perm (before umask). If successful,
	// methods on the returned file can be used for I/O.
	// If there is an error, it will be of type `*PathError`.
	OpenFile(path string, flag int, perm fs.FileMode) (io.ReadWriteCloser, error)
}

// A RemoveFS is a file system with a `Remove()` and a `RemoveAll()` methods.
type RemoveFS interface {
	// Remove removes the named file or (empty) directory.
	// If there is an error, it will be of type `*PathError`.
	Remove(path string) error

	// RemoveAll removes path and any children it contains.
	// It removes everything it can but returns the first error
	// it encounters. If the path does not exist, `RemoveAll`
	// returns `nil` (no error).
	// If there is an error, it will be of type `*PathError`.
	RemoveAll(path string) error
}

// Embed is an extension of aimed at improving `embed.FS` by
// implementing `fs.StatFS` and a `Sub()` function.
type Embed struct {
	FS fs.ReadDirFS
}

// NewEmbed returns a new Embed with the given FS.
func NewEmbed(fs fs.ReadDirFS) Embed { _ = "STUB: not implemented"; return *new(Embed) }

// Open opens the named file.
//
// When Open returns an error, it should be of type *PathError
// with the Op field set to "open", the Path field set to name,
// and the Err field describing the problem.
//
// Open should reject attempts to open names that do not satisfy
// ValidPath(name), returning a *PathError with Err set to
// ErrInvalid or ErrNotExist.
func (e Embed) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
func (e Embed) ReadDir(name string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stat returns a FileInfo describing the file.
func (e Embed) Stat(name string) (fileinfo fs.FileInfo, err error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

// Sub returns an Embed FS corresponding to the subtree rooted at dir.
// Returns and error if the underlying sub FS doesn't implement `fs.ReadDirFS`.
func (e Embed) Sub(dir string) (Embed, error) { _ = "STUB: not implemented"; return *new(Embed), nil }
