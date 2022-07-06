package schema

import "entgo.io/ent"

// Mailbox holds the schema definition for the Mailbox entity.
type Mailbox struct {
	ent.Schema
}

// Fields of the Mailbox.
func (Mailbox) Fields() []ent.Field {
	return nil
}

// Edges of the Mailbox.
func (Mailbox) Edges() []ent.Edge {
	return nil
}
