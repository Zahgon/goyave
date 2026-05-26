package fsutil

import (
	"mime/multipart"
	"sync"
)

// marshalCache temporarily stores files' `*multipart.FileHeader`. This type
// cannot be marshaled, making the use of `fsutil.file` inconvenient with DTO conversion.
// The key should a be unique ID. The key is removed from the map.
// To avoid infinite growth of this cache, leading to potential memory problems, this map
// is reset every time its length goes back to 0.
var marshalCache = map[string]*multipart.FileHeader{}
var cacheMu sync.RWMutex

// File represents a file received from client.
//
// File implements `json.Marshaler` and `json.Unmarshaler` to be able
// to be used in DTO conversion (`typeutil.Convert()`). This works with a global
// concurrency-safe map that acts as a cache for the `*multipart.FileHeader`.
// When marshaling, a UUID v1 is generated and used as a key. This UUID is the actual value
// used when marhsaling the `Header` field. When unmarshaling, the `*multipart.FileHeader` is
// retrieved then deleted from the cache. To avoid orphans clogging up the cache, you should
// never JSON marshal this type outside of `typeutil.Convert()`: if a marshaled File never gets
// unmarshaled, its UUID would remain in the cache forever.
type File struct {
	Header   *multipart.FileHeader
	MIMEType string
}

type marshaledFile struct {
	MIMEType string
	Header   string
}

// MarshalJSON implementation of `json.Marhsaler`.
func (file File) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implementation of `json.Unmarhsaler`.
func (file *File) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Maps never shrink, let's allocate a new empty map to reset the cache capacity
// and allow garbage collecting.

// Save writes the file's content to a new file in the given file system.
// Appends a timestamp to the given file name to avoid duplicate file names.
// The file is not readable anymore once saved as its FileReader has already been
// closed.
//
// Creates directories if needed.
//
// Returns the actual file name.
func (file *File) Save(fs WritableFS, path string, name string) (filename string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseMultipartFiles parse a single file field in a request.
func ParseMultipartFiles(headers []*multipart.FileHeader) ([]File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
