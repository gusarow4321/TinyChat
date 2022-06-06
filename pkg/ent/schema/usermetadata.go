package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserMetadata holds the schema definition for the UserMetadata entity.
type UserMetadata struct {
	ent.Schema
}

// Fields of the UserMetadata.
func (UserMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("userID"),
		field.Int32("color"),
	}
}

// Edges of the UserMetadata.
func (UserMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("metadata").Unique().Required().Field("userID"),
	}
}
