// Code generated by entc, DO NOT EDIT.

package guild

import (
	"time"
)

const (
	// Label holds the string label denoting the guild type in the database.
	Label = "guild"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSnowflake holds the string denoting the snowflake field in the database.
	FieldSnowflake = "snowflake"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldCategories holds the string denoting the categories field in the database.
	FieldCategories = "categories"
	// FieldEntitlements holds the string denoting the entitlements field in the database.
	FieldEntitlements = "entitlements"

	// Table holds the table name of the guild in the database.
	Table = "guilds"
)

// Columns holds all SQL columns for guild fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSnowflake,
	FieldMessage,
	FieldCategories,
	FieldEntitlements,
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
)
