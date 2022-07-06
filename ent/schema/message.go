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
		field.String("message_id"),
		field.Uint32("imap_uid").Nillable().Optional(),
		field.String("subject").Nillable().Optional(),
		field.Time("timestamp").Nillable().Optional(),
		field.Text("header").Nillable().Optional(),
		field.Text("body").Nillable().Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mailbox", Mailbox.Type).
			Ref("messages").Unique().Required(),
	}
}
