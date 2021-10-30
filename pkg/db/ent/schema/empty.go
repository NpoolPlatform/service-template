package schema

import "entgo.io/ent"

// Empty holds the schema definition for the Empty entity.
type Empty struct {
	ent.Schema
}

// Fields of the Empty.
func (Empty) Fields() []ent.Field {
	return nil
}

// Edges of the Empty.
func (Empty) Edges() []ent.Edge {
	return nil
}
