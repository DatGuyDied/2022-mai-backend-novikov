package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserMessage holds the schema definition for the UserMessage entity.
type UserMessage struct {
	ent.Schema
}

// Fields of the UserMessage.
func (UserMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("from"),
		field.String("to"),
		field.String("text"),
		field.Time("ctreated_at").Default(time.Now),
	}
}

// Edges of the UserMessage.
func (UserMessage) Edges() []ent.Edge {
	return nil
}
