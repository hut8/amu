package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mailbox holds the schema definition for the Mailbox entity.
type Mailbox struct {
	ent.Schema
}

// Fields of the Mailbox.
func (Mailbox) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Mailbox.
func (Mailbox) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("mailboxes").
			Unique().Required(),
		edge.To("messages", Message.Type),
	}
}
