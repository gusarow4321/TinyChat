package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("chatID"),
		field.Int64("userID"),
		field.Text("text"),
		field.Time("timestamp").Immutable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).Ref("messages").Unique().Required().Field("chatID"),
		edge.From("user", User.Type).Ref("messages").Unique().Required().Field("userID"),
	}
}
