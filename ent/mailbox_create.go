// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hut8/amu/ent/account"
	"github.com/hut8/amu/ent/mailbox"
	"github.com/hut8/amu/ent/message"
)

// MailboxCreate is the builder for creating a Mailbox entity.
type MailboxCreate struct {
	config
	mutation *MailboxMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (mc *MailboxCreate) SetName(s string) *MailboxCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (mc *MailboxCreate) SetAccountID(id int) *MailboxCreate {
	mc.mutation.SetAccountID(id)
	return mc
}

// SetAccount sets the "account" edge to the Account entity.
func (mc *MailboxCreate) SetAccount(a *Account) *MailboxCreate {
	return mc.SetAccountID(a.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (mc *MailboxCreate) AddMessageIDs(ids ...int) *MailboxCreate {
	mc.mutation.AddMessageIDs(ids...)
	return mc
}

// AddMessages adds the "messages" edges to the Message entity.
func (mc *MailboxCreate) AddMessages(m ...*Message) *MailboxCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMessageIDs(ids...)
}

// Mutation returns the MailboxMutation object of the builder.
func (mc *MailboxCreate) Mutation() *MailboxMutation {
	return mc.mutation
}

// Save creates the Mailbox in the database.
func (mc *MailboxCreate) Save(ctx context.Context) (*Mailbox, error) {
	var (
		err  error
		node *Mailbox
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxMutation)
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
func (mc *MailboxCreate) SaveX(ctx context.Context) *Mailbox {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MailboxCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MailboxCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MailboxCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Mailbox.name"`)}
	}
	if _, ok := mc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Mailbox.account"`)}
	}
	return nil
}

func (mc *MailboxCreate) sqlSave(ctx context.Context) (*Mailbox, error) {
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

func (mc *MailboxCreate) createSpec() (*Mailbox, *sqlgraph.CreateSpec) {
	var (
		_node = &Mailbox{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mailbox.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mailbox.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldName,
		})
		_node.Name = value
	}
	if nodes := mc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mailbox.AccountTable,
			Columns: []string{mailbox.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_mailboxes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.MessagesTable,
			Columns: []string{mailbox.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mailbox.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MailboxUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MailboxCreate) OnConflict(opts ...sql.ConflictOption) *MailboxUpsertOne {
	mc.conflict = opts
	return &MailboxUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mailbox.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MailboxCreate) OnConflictColumns(columns ...string) *MailboxUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MailboxUpsertOne{
		create: mc,
	}
}

type (
	// MailboxUpsertOne is the builder for "upsert"-ing
	//  one Mailbox node.
	MailboxUpsertOne struct {
		create *MailboxCreate
	}

	// MailboxUpsert is the "OnConflict" setter.
	MailboxUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *MailboxUpsert) SetName(v string) *MailboxUpsert {
	u.Set(mailbox.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailboxUpsert) UpdateName() *MailboxUpsert {
	u.SetExcluded(mailbox.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Mailbox.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MailboxUpsertOne) UpdateNewValues() *MailboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Mailbox.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MailboxUpsertOne) Ignore() *MailboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MailboxUpsertOne) DoNothing() *MailboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MailboxCreate.OnConflict
// documentation for more info.
func (u *MailboxUpsertOne) Update(set func(*MailboxUpsert)) *MailboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MailboxUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MailboxUpsertOne) SetName(v string) *MailboxUpsertOne {
	return u.Update(func(s *MailboxUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailboxUpsertOne) UpdateName() *MailboxUpsertOne {
	return u.Update(func(s *MailboxUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MailboxUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MailboxCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MailboxUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MailboxUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MailboxUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MailboxCreateBulk is the builder for creating many Mailbox entities in bulk.
type MailboxCreateBulk struct {
	config
	builders []*MailboxCreate
	conflict []sql.ConflictOption
}

// Save creates the Mailbox entities in the database.
func (mcb *MailboxCreateBulk) Save(ctx context.Context) ([]*Mailbox, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mailbox, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MailboxMutation)
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
func (mcb *MailboxCreateBulk) SaveX(ctx context.Context) []*Mailbox {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MailboxCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MailboxCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mailbox.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MailboxUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MailboxCreateBulk) OnConflict(opts ...sql.ConflictOption) *MailboxUpsertBulk {
	mcb.conflict = opts
	return &MailboxUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mailbox.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MailboxCreateBulk) OnConflictColumns(columns ...string) *MailboxUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MailboxUpsertBulk{
		create: mcb,
	}
}

// MailboxUpsertBulk is the builder for "upsert"-ing
// a bulk of Mailbox nodes.
type MailboxUpsertBulk struct {
	create *MailboxCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Mailbox.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MailboxUpsertBulk) UpdateNewValues() *MailboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mailbox.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MailboxUpsertBulk) Ignore() *MailboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MailboxUpsertBulk) DoNothing() *MailboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MailboxCreateBulk.OnConflict
// documentation for more info.
func (u *MailboxUpsertBulk) Update(set func(*MailboxUpsert)) *MailboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MailboxUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MailboxUpsertBulk) SetName(v string) *MailboxUpsertBulk {
	return u.Update(func(s *MailboxUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailboxUpsertBulk) UpdateName() *MailboxUpsertBulk {
	return u.Update(func(s *MailboxUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MailboxUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MailboxCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MailboxCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MailboxUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
