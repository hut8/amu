// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hut8/amu/ent/mailbox"
	"github.com/hut8/amu/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetMessageID sets the "message_id" field.
func (mc *MessageCreate) SetMessageID(s string) *MessageCreate {
	mc.mutation.SetMessageID(s)
	return mc
}

// SetImapUID sets the "imap_uid" field.
func (mc *MessageCreate) SetImapUID(u uint32) *MessageCreate {
	mc.mutation.SetImapUID(u)
	return mc
}

// SetNillableImapUID sets the "imap_uid" field if the given value is not nil.
func (mc *MessageCreate) SetNillableImapUID(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetImapUID(*u)
	}
	return mc
}

// SetHeader sets the "header" field.
func (mc *MessageCreate) SetHeader(s string) *MessageCreate {
	mc.mutation.SetHeader(s)
	return mc
}

// SetNillableHeader sets the "header" field if the given value is not nil.
func (mc *MessageCreate) SetNillableHeader(s *string) *MessageCreate {
	if s != nil {
		mc.SetHeader(*s)
	}
	return mc
}

// SetBody sets the "body" field.
func (mc *MessageCreate) SetBody(s string) *MessageCreate {
	mc.mutation.SetBody(s)
	return mc
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (mc *MessageCreate) SetNillableBody(s *string) *MessageCreate {
	if s != nil {
		mc.SetBody(*s)
	}
	return mc
}

// SetMailboxID sets the "mailbox" edge to the Mailbox entity by ID.
func (mc *MessageCreate) SetMailboxID(id int) *MessageCreate {
	mc.mutation.SetMailboxID(id)
	return mc
}

// SetMailbox sets the "mailbox" edge to the Mailbox entity.
func (mc *MessageCreate) SetMailbox(m *Mailbox) *MessageCreate {
	return mc.SetMailboxID(m.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message_id", err: errors.New(`ent: missing required field "Message.message_id"`)}
	}
	if _, ok := mc.mutation.MailboxID(); !ok {
		return &ValidationError{Name: "mailbox", err: errors.New(`ent: missing required edge "Message.mailbox"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.MessageID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageID,
		})
		_node.MessageID = value
	}
	if value, ok := mc.mutation.ImapUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldImapUID,
		})
		_node.ImapUID = &value
	}
	if value, ok := mc.mutation.Header(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldHeader,
		})
		_node.Header = &value
	}
	if value, ok := mc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldBody,
		})
		_node.Body = &value
	}
	if nodes := mc.mutation.MailboxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.MailboxTable,
			Columns: []string{message.MailboxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailbox.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.mailbox_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetMessageID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetMessageID(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetMessageID sets the "message_id" field.
func (u *MessageUpsert) SetMessageID(v string) *MessageUpsert {
	u.Set(message.FieldMessageID, v)
	return u
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMessageID() *MessageUpsert {
	u.SetExcluded(message.FieldMessageID)
	return u
}

// SetImapUID sets the "imap_uid" field.
func (u *MessageUpsert) SetImapUID(v uint32) *MessageUpsert {
	u.Set(message.FieldImapUID, v)
	return u
}

// UpdateImapUID sets the "imap_uid" field to the value that was provided on create.
func (u *MessageUpsert) UpdateImapUID() *MessageUpsert {
	u.SetExcluded(message.FieldImapUID)
	return u
}

// AddImapUID adds v to the "imap_uid" field.
func (u *MessageUpsert) AddImapUID(v uint32) *MessageUpsert {
	u.Add(message.FieldImapUID, v)
	return u
}

// ClearImapUID clears the value of the "imap_uid" field.
func (u *MessageUpsert) ClearImapUID() *MessageUpsert {
	u.SetNull(message.FieldImapUID)
	return u
}

// SetHeader sets the "header" field.
func (u *MessageUpsert) SetHeader(v string) *MessageUpsert {
	u.Set(message.FieldHeader, v)
	return u
}

// UpdateHeader sets the "header" field to the value that was provided on create.
func (u *MessageUpsert) UpdateHeader() *MessageUpsert {
	u.SetExcluded(message.FieldHeader)
	return u
}

// ClearHeader clears the value of the "header" field.
func (u *MessageUpsert) ClearHeader() *MessageUpsert {
	u.SetNull(message.FieldHeader)
	return u
}

// SetBody sets the "body" field.
func (u *MessageUpsert) SetBody(v string) *MessageUpsert {
	u.Set(message.FieldBody, v)
	return u
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MessageUpsert) UpdateBody() *MessageUpsert {
	u.SetExcluded(message.FieldBody)
	return u
}

// ClearBody clears the value of the "body" field.
func (u *MessageUpsert) ClearBody() *MessageUpsert {
	u.SetNull(message.FieldBody)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Message.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertOne) SetMessageID(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMessageID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// SetImapUID sets the "imap_uid" field.
func (u *MessageUpsertOne) SetImapUID(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetImapUID(v)
	})
}

// AddImapUID adds v to the "imap_uid" field.
func (u *MessageUpsertOne) AddImapUID(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddImapUID(v)
	})
}

// UpdateImapUID sets the "imap_uid" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateImapUID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateImapUID()
	})
}

// ClearImapUID clears the value of the "imap_uid" field.
func (u *MessageUpsertOne) ClearImapUID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearImapUID()
	})
}

// SetHeader sets the "header" field.
func (u *MessageUpsertOne) SetHeader(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetHeader(v)
	})
}

// UpdateHeader sets the "header" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateHeader() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateHeader()
	})
}

// ClearHeader clears the value of the "header" field.
func (u *MessageUpsertOne) ClearHeader() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearHeader()
	})
}

// SetBody sets the "body" field.
func (u *MessageUpsertOne) SetBody(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateBody() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *MessageUpsertOne) ClearBody() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearBody()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetMessageID(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertBulk) SetMessageID(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMessageID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// SetImapUID sets the "imap_uid" field.
func (u *MessageUpsertBulk) SetImapUID(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetImapUID(v)
	})
}

// AddImapUID adds v to the "imap_uid" field.
func (u *MessageUpsertBulk) AddImapUID(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddImapUID(v)
	})
}

// UpdateImapUID sets the "imap_uid" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateImapUID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateImapUID()
	})
}

// ClearImapUID clears the value of the "imap_uid" field.
func (u *MessageUpsertBulk) ClearImapUID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearImapUID()
	})
}

// SetHeader sets the "header" field.
func (u *MessageUpsertBulk) SetHeader(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetHeader(v)
	})
}

// UpdateHeader sets the "header" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateHeader() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateHeader()
	})
}

// ClearHeader clears the value of the "header" field.
func (u *MessageUpsertBulk) ClearHeader() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearHeader()
	})
}

// SetBody sets the "body" field.
func (u *MessageUpsertBulk) SetBody(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateBody() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *MessageUpsertBulk) ClearBody() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearBody()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
