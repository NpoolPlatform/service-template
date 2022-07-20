package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/mixin"
	"github.com/google/uuid"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

func (Template) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").
			Unique(),
		field.Uint32("age").
			Default(0),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return nil
}
