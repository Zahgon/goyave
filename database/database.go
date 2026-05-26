package database

import (
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/slog"
)

// New create a new connection pool using the settings defined in the given configuration.
//
// In order to use a specific driver / dialect ("mysql", "sqlite3", ...), you must not
// forget to blank-import it in your main file.
//
//	import _ "goyave.dev/goyave/v5/database/dialect/mysql"
//	import _ "goyave.dev/goyave/v5/database/dialect/postgres"
//	import _ "goyave.dev/goyave/v5/database/dialect/sqlite"
//	import _ "goyave.dev/goyave/v5/database/dialect/mssql"
//	import _ "goyave.dev/goyave/v5/database/dialect/clickhouse"
//	import _ "goyave.dev/goyave/v5/database/dialect/bigquery"
func New(cfg *config.Config, logger func() *slog.Logger) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFromDialector create a new connection pool from a gorm dialector and using the settings
// defined in the given configuration.
//
// This can be used in tests to create a mock connection pool.
func NewFromDialector(cfg *config.Config, logger func() *slog.Logger, dialector gorm.Dialector) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newConfig(cfg *config.Config, logger func() *slog.Logger) *gorm.Config {
	_ = "STUB: not implemented"
	return nil
}

// Stay silent about DB operations when not in debug mode

func initTimeoutPlugin(cfg *config.Config, db *gorm.DB) error {
	_ = "STUB: not implemented"
	return nil
}

func initSQLDB(cfg *config.Config, db *gorm.DB) error { _ = "STUB: not implemented"; return nil }
