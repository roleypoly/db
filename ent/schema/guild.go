package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/roleypoly/db/ent/mixin"
)

// Guild holds the schema definition for the Guild entity.
type Guild struct {
	ent.Schema
}

type Category struct {
	ID       string
	Name     string
	Hidden   bool
	Type     string
	Position int
	Roles    []string
}

// Fields of the Guild.
func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.Text("snowflake").
			Immutable().Unique(),
		field.Text("message"),
		field.JSON("categories", []Category{}),
		field.Strings("entitlements"),
	}
}

// Edges of the Guild.
func (Guild) Edges() []ent.Edge {
	return nil
}

// Mixin of the Guild.
func (Guild) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}
