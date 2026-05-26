package session

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// Session aims at facilitating business transactions while abstracting the underlying mechanism,
// be it a database transaction or another transaction mechanism. This allows services to execute
// multiple business use-cases and easily rollback changes in case of error, without creating a
// dependency to the database layer.
//
// Sessions should be constituted of a root session created with a "New"-type constructor and allow
// the creation of child sessions with `Begin()` and `Transaction()`. Nested transactions should be supported
// as well.
type Session interface {
	// Begin returns a new session with the given context and a started transaction.
	// Using the returned session should have no side-effect on the parent session.
	// The underlying transaction mechanism is injected as a value into the new session's context.
	Begin(ctx context.Context) (Session, error)

	// Transaction executes a transaction. If the given function returns an error, the transaction
	// is rolled back. Otherwise it is automatically committed before `Transaction()` returns.
	// The underlying transaction mechanism is injected into the context as a value.
	Transaction(ctx context.Context, f func(context.Context) error) error

	// Rollback the changes in the transaction. This action is final.
	Rollback() error

	// Commit the changes in the transaction. This action is final.
	Commit() error

	// Context returns the session's context. If it's the root session, `context.Background()` is returned.
	// If it's a child session started with `Begin()`, then the context will contain the associated
	// transaction mechanism as a value.
	Context() context.Context
}

// Gorm session implementation.
type Gorm struct {
	db        *gorm.DB
	TxOptions *sql.TxOptions
	ctx       context.Context
	savepoint string // Savepoint for manual nested transactions
}

// GORM create a new root session for Gorm.
// The transaction options are optional.
func GORM(db *gorm.DB, opt *sql.TxOptions) Gorm { _ = "STUB: not implemented"; return *new(Gorm) }

// Begin returns a new session with the given context and a started DB transaction.
// The returned session has manual controls. Make sure a call to `Rollback()` or `Commit()`
// is executed before the session is expired (eligible for garbage collection).
// The Gorm DB associated with this session is injected as a value into the new session's context.
//
// If the newly created session is nested, a savepoint is generated instead. Calls to the returned
// session's `Rollback()` will rollback to this savepoint.
// This behavior is disabled if gorm config `DisableNestedTransaction` is set to `true`.
func (s Gorm) Begin(ctx context.Context) (Session, error) {
	_ = "STUB: not implemented"
	return *new(Session), nil
}

func (s Gorm) nestedBegin(db *gorm.DB) (Session, error) {
	_ = "STUB: not implemented"
	return *new(Session), nil
}

// Rollback the changes in the transaction. This action is final.
//
// If the session is nested, rolls back to the session's savepoint.
func (s Gorm) Rollback() error { _ = "STUB: not implemented"; return nil }

// Commit the changes in the transaction. This action is final.
//
// If the session is nested, calling Rollback() is a no-op.
func (s Gorm) Commit() error { _ = "STUB: not implemented"; return nil }

// Context returns the session's context. If it's the root session, `context.Background()`
// is returned. If it's a child session started with `Begin()`, then the context will contain
// the associated Gorm DB and can be used in combination with `session.DB()`.
func (s Gorm) Context() context.Context {
	_ = "STUB: not implemented"

	// dbKey the key used to store the database in the context.
	return *new(context.Context)
}

type dbKey struct{}

// Transaction executes a transaction. If the given function returns an error, the transaction
// is rolled back. Otherwise it is automatically committed before `Transaction()` returns.
//
// The Gorm DB associated with this session is injected into the context as a value so `session.DB()`
// can be used to retrieve it.
func (s Gorm) Transaction(ctx context.Context, f func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Gorm) nestedTransaction(tx *gorm.DB, f func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// DB returns the Gorm instance stored in the given context.
// If no Gorm DB could be found in the context, calls `fallback.WithContext` and
// return the result.
func DB(ctx context.Context, fallback *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }
