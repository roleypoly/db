// Code generated by entc, DO NOT EDIT.

package session

import (
	"fmt"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/roleypoly/db/ent/schema"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSessionID holds the string denoting the session_id vertex property in the database.
	FieldSessionID = "session_id"
	// FieldUserID holds the string denoting the user_id vertex property in the database.
	FieldUserID = "user_id"
	// FieldSource holds the string denoting the source vertex property in the database.
	FieldSource = "source"
	// FieldExpiresAt holds the string denoting the expires_at vertex property in the database.
	FieldExpiresAt = "expires_at"

	// Table holds the table name of the session in the database.
	Table = "sessions"
)

// Columns holds all SQL columns are session fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSessionID,
	FieldUserID,
	FieldSource,
	FieldExpiresAt,
}

var (
	mixin       = schema.Session{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
	}
	fields = schema.Session{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = mixinFields[0][0].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = mixinFields[0][1].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)

	// descExpiresAt is the schema descriptor for expires_at field.
	descExpiresAt = fields[3].Descriptor()
	// DefaultExpiresAt holds the default value on creation for the expires_at field.
	DefaultExpiresAt = descExpiresAt.Default.(func() time.Time)
)

// Source defines the type for the source enum field.
type Source string

const (
	SourceOauth Source = "oauth"
	SourceDm    Source = "dm"
)

func (s Source) String() string {
	return string(s)
}

// SourceValidator is a validator for the "source" field enum values. It is called by the builders before save.
func SourceValidator(source Source) error {
	switch source {
	case SourceOauth, SourceDm:
		return nil
	default:
		return fmt.Errorf("session: invalid enum value for source field: %q", source)
	}
}
