package validation

import (
	"github.com/google/uuid"
)

// UUIDValidator the field under validation must be a string representing
// a valid UUID.
// If one or more `accepterVersions` are provided, the parsed UUID must
// be a UUID of one of these versions. If none are given, all versions are
// accepted.
//
// If validation passes, the value is converted to `uuid.UUID`.
type UUIDValidator struct {
	BaseValidator
	AcceptedVersions []uuid.Version
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *UUIDValidator) Validate(ctx *Context) bool { _ = "STUB: not implemented"; return false }

func (v *UUIDValidator) checkVersion(uid uuid.UUID) bool { _ = "STUB: not implemented"; return false }

// Name returns the string name of the validator.
func (v *UUIDValidator) Name() string {
	_ = "STUB: not implemented"

	// IsType returns true.
	return ""
}

func (v *UUIDValidator) IsType() bool {
	_ = "STUB: not implemented"

	// TODO specify accepted versions in validation message?
	return false
}

// UUID the field under validation must be a string representing
// a valid UUID.
// If one or more `accepterVersions` are provided, the parsed UUID must
// be a UUID of one of these versions. If none are given, all versions are
// accepted.
//
// If validation passes, the value is converted to `uuid.UUID`.
func UUID(acceptedVersions ...uuid.Version) *UUIDValidator { _ = "STUB: not implemented"; return nil }
