package database

import (
	"gorm.io/gorm"
)

// Factory an object used to generate records or seed the database.
type Factory[T any] struct {
	generator func() *T
	override  *T
	BatchSize int
}

// NewFactory create a new Factory.
// The given generator function will be used to generate records.
func NewFactory[T any](generator func() *T) *Factory[T] { _ = "STUB: not implemented"; return nil }

// Override set an override model for generated records.
// Values present in the override model will replace the ones
// in the generated records.
// This function expects a struct pointer as parameter.
// Returns the same instance of `Factory` so this method can be chained.
func (f *Factory[T]) Override(override *T) *Factory[T] { _ = "STUB: not implemented"; return nil }

// Generate a number of records using the given factory.
func (f *Factory[T]) Generate(count int) []*T { _ = "STUB: not implemented"; return nil }

// Save generate a number of records using the given factory,
// insert them in the database and return the inserted records.
func (f *Factory[T]) Save(db *gorm.DB, count int) []*T { _ = "STUB: not implemented"; return nil }
