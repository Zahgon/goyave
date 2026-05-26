package database

import (
	"sync"

	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
)

var (
	mu sync.Mutex

	dialects = map[string]dialect{}

	optionPlaceholders = map[string]string{
		"{username}": "database.username",
		"{password}": "database.password",
		"{host}":     "database.host",
		"{name}":     "database.name",
		"{options}":  "database.options",
	}
)

// DialectorInitializer function initializing a GORM Dialector using the given
// data source name (DSN).
type DialectorInitializer func(dsn string) gorm.Dialector

type dialect struct {
	initializer DialectorInitializer
	template    string
}

func (d dialect) buildDSN(cfg *config.Config) string { _ = "STUB: not implemented"; return "" }

// RegisterDialect registers a connection string template for the given dialect.
//
// You cannot override a dialect that already exists.
//
// Template format accepts the following placeholders, which will be replaced with
// the corresponding configuration entries automatically:
//   - "{username}"
//   - "{password}"
//   - "{host}"
//   - "{port}"
//   - "{name}"
//   - "{options}"
//
// Example template for the "mysql" dialect:
//
//	{username}:{password}@({host}:{port})/{name}?{options}
func RegisterDialect(name, template string, initializer DialectorInitializer) {
	_ = "STUB: not implemented"
	return
}
