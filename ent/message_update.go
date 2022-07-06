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
	"github.com/hut8/amu/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetMessageID sets the "message_id" field.
func (mu *MessageUpdate) SetMessageID(s string) *MessageUpdate {
	mu.mutation.SetMessageID(s)
	return mu
}

// SetImapUID sets the "imap_uid" field.
func (mu *MessageUpdate) SetImapUID(u uint32) *MessageUpdate {
	mu.mutation.ResetImapUID()
	mu.mutation.SetImapUID(u)
	return mu
}

// SetNillableImapUID sets the "imap_uid" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableImapUID(u *uint32) *MessageUpdate {
	if u != nil {
		mu.SetImapUID(*u)
	}
	return mu
}

// AddImapUID adds u to the "imap_uid" field.
func (mu *MessageUpdate) AddImapUID(u int32) *MessageUpdate {
	mu.mutation.AddImapUID(u)
	return mu
}

// ClearImapUID clears the value of the "imap_uid" field.
func (mu *MessageUpdate) ClearImapUID() *MessageUpdate {
	mu.mutation.ClearImapUID()
	return mu
}

// SetHeader sets the "header" field.
func (mu *MessageUpdate) SetHeader(s string) *MessageUpdate {
	mu.mutation.SetHeader(s)
	return mu
}

// SetNillableHeader sets the "header" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableHeader(s *string) *MessageUpdate {
	if s != nil {
		mu.SetHeader(*s)
	}
	return mu
}

// ClearHeader clears the value of the "header" field.
func (mu *MessageUpdate) ClearHeader() *MessageUpdate {
	mu.mutation.ClearHeader()
	return mu
}

// SetBody sets the "body" field.
func (mu *MessageUpdate) SetBody(s string) *MessageUpdate {
	mu.mutation.SetBody(s)
	return mu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableBody(s *string) *MessageUpdate {
	if s != nil {
		mu.SetBody(*s)
	}
	return mu
}

// ClearBody clears the value of the "body" field.
func (mu *MessageUpdate) ClearBody() *MessageUpdate {
	mu.mutation.ClearBody()
	return mu
}

// SetMailboxID sets the "mailbox" edge to the Mailbox entity by ID.
func (mu *MessageUpdate) SetMailboxID(id int) *MessageUpdate {
	mu.mutation.SetMailboxID(id)
	return mu
}

// SetMailbox sets the "mailbox" edge to the Mailbox entity.
func (mu *MessageUpdate) SetMailbox(m *Mailbox) *MessageUpdate {
	return mu.SetMailboxID(m.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearMailbox clears the "mailbox" edge to the Mailbox entity.
func (mu *MessageUpdate) ClearMailbox() *MessageUpdate {
	mu.mutation.ClearMailbox()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if _, ok := mu.mutation.MailboxID(); mu.mutation.MailboxCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.mailbox"`)
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageID,
		})
	}
	if value, ok := mu.mutation.ImapUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldImapUID,
		})
	}
	if value, ok := mu.mutation.AddedImapUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldImapUID,
		})
	}
	if mu.mutation.ImapUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: message.FieldImapUID,
		})
	}
	if value, ok := mu.mutation.Header(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldHeader,
		})
	}
	if mu.mutation.HeaderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldHeader,
		})
	}
	if value, ok := mu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldBody,
		})
	}
	if mu.mutation.BodyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldBody,
		})
	}
	if mu.mutation.MailboxCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MailboxIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetMessageID sets the "message_id" field.
func (muo *MessageUpdateOne) SetMessageID(s string) *MessageUpdateOne {
	muo.mutation.SetMessageID(s)
	return muo
}

// SetImapUID sets the "imap_uid" field.
func (muo *MessageUpdateOne) SetImapUID(u uint32) *MessageUpdateOne {
	muo.mutation.ResetImapUID()
	muo.mutation.SetImapUID(u)
	return muo
}

// SetNillableImapUID sets the "imap_uid" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableImapUID(u *uint32) *MessageUpdateOne {
	if u != nil {
		muo.SetImapUID(*u)
	}
	return muo
}

// AddImapUID adds u to the "imap_uid" field.
func (muo *MessageUpdateOne) AddImapUID(u int32) *MessageUpdateOne {
	muo.mutation.AddImapUID(u)
	return muo
}

// ClearImapUID clears the value of the "imap_uid" field.
func (muo *MessageUpdateOne) ClearImapUID() *MessageUpdateOne {
	muo.mutation.ClearImapUID()
	return muo
}

// SetHeader sets the "header" field.
func (muo *MessageUpdateOne) SetHeader(s string) *MessageUpdateOne {
	muo.mutation.SetHeader(s)
	return muo
}

// SetNillableHeader sets the "header" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableHeader(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetHeader(*s)
	}
	return muo
}

// ClearHeader clears the value of the "header" field.
func (muo *MessageUpdateOne) ClearHeader() *MessageUpdateOne {
	muo.mutation.ClearHeader()
	return muo
}

// SetBody sets the "body" field.
func (muo *MessageUpdateOne) SetBody(s string) *MessageUpdateOne {
	muo.mutation.SetBody(s)
	return muo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableBody(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetBody(*s)
	}
	return muo
}

// ClearBody clears the value of the "body" field.
func (muo *MessageUpdateOne) ClearBody() *MessageUpdateOne {
	muo.mutation.ClearBody()
	return muo
}

// SetMailboxID sets the "mailbox" edge to the Mailbox entity by ID.
func (muo *MessageUpdateOne) SetMailboxID(id int) *MessageUpdateOne {
	muo.mutation.SetMailboxID(id)
	return muo
}

// SetMailbox sets the "mailbox" edge to the Mailbox entity.
func (muo *MessageUpdateOne) SetMailbox(m *Mailbox) *MessageUpdateOne {
	return muo.SetMailboxID(m.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearMailbox clears the "mailbox" edge to the Mailbox entity.
func (muo *MessageUpdateOne) ClearMailbox() *MessageUpdateOne {
	muo.mutation.ClearMailbox()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if _, ok := muo.mutation.MailboxID(); muo.mutation.MailboxCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.mailbox"`)
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageID,
		})
	}
	if value, ok := muo.mutation.ImapUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldImapUID,
		})
	}
	if value, ok := muo.mutation.AddedImapUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldImapUID,
		})
	}
	if muo.mutation.ImapUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: message.FieldImapUID,
		})
	}
	if value, ok := muo.mutation.Header(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldHeader,
		})
	}
	if muo.mutation.HeaderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldHeader,
		})
	}
	if value, ok := muo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldBody,
		})
	}
	if muo.mutation.BodyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldBody,
		})
	}
	if muo.mutation.MailboxCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MailboxIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
