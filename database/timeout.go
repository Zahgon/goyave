package database

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	timeoutCallbackBeforeName = "goyave:timeout_before"
	timeoutCallbackAfterName  = "goyave:timeout_after"
)

type timeoutContext struct {
	context.Context

	parentContext context.Context

	// We store the pointer to the original statement
	// so we can cancel the context only if the original
	// statement is completely finished. This prevents
	// sub-statements (such as preloads) to cancel the context
	// when they are done, despite the parent statement not being
	// executed yet.
	statement *gorm.Statement

	cancel context.CancelFunc
}

// TimeoutPlugin GORM plugin adding a default timeout to SQL queries if none is applied
// on the statement already. It works by replacing the statement's context with a child
// context having the configured timeout. The context is replaced in a "before" callback
// on all GORM operations. In a "after" callback, the new context is canceled.
//
// The `ReadTimeout` is applied on the `Query` and `Raw` GORM callbacks. The `WriteTimeout`
// is applied on the rest of the callbacks.
//
// Supports all GORM operations except `Scan()`.
//
// A timeout duration inferior or equal to 0 disables the plugin for the relevant operations.
type TimeoutPlugin struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Name returns the name of the plugin
func (p *TimeoutPlugin) Name() string { _ = "STUB: not implemented"; return "" }

// Initialize registers the callbacks for all operations.
func (p *TimeoutPlugin) Initialize(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

// Cannot use it with `Row()` because context is canceled before the call of `rows.Next()`, causing an error.
// rowCallback := db.Callback().Row()
// if err := rowCallback.Before("*").Register(timeoutCallbackBeforeName, p.readTimeoutBefore); err != nil {
// 	return errors.New(err)
// }
// if err := rowCallback.After("*").Register(timeoutCallbackAfterName, p.timeoutAfter); err != nil {
// 	return errors.New(err)
// }

func (p *TimeoutPlugin) readTimeoutBefore(db *gorm.DB) { _ = "STUB: not implemented"; return }

func (p *TimeoutPlugin) writeTimeoutBefore(db *gorm.DB) { _ = "STUB: not implemented"; return }

func (p *TimeoutPlugin) timeoutBefore(db *gorm.DB, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// The statement is re-used, replace the context with a new one

func (p *TimeoutPlugin) timeoutAfter(db *gorm.DB) { _ = "STUB: not implemented"; return }
