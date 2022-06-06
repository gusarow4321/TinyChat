// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gusarow4321/TinyChat/pkg/ent/chat"
	"github.com/gusarow4321/TinyChat/pkg/ent/message"
	"github.com/gusarow4321/TinyChat/pkg/ent/predicate"
	"github.com/gusarow4321/TinyChat/pkg/ent/user"
)

// ChatUpdate is the builder for updating Chat entities.
type ChatUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cu *ChatUpdate) Where(ps ...predicate.Chat) *ChatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetOwnerID sets the "ownerID" field.
func (cu *ChatUpdate) SetOwnerID(i int64) *ChatUpdate {
	cu.mutation.SetOwnerID(i)
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *ChatUpdate) SetOwner(u *User) *ChatUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cu *ChatUpdate) AddMessageIDs(ids ...int64) *ChatUpdate {
	cu.mutation.AddMessageIDs(ids...)
	return cu
}

// AddMessages adds the "messages" edges to the Message entity.
func (cu *ChatUpdate) AddMessages(m ...*Message) *ChatUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMessageIDs(ids...)
}

// Mutation returns the ChatMutation object of the builder.
func (cu *ChatUpdate) Mutation() *ChatMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *ChatUpdate) ClearOwner() *ChatUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cu *ChatUpdate) ClearMessages() *ChatUpdate {
	cu.mutation.ClearMessages()
	return cu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cu *ChatUpdate) RemoveMessageIDs(ids ...int64) *ChatUpdate {
	cu.mutation.RemoveMessageIDs(ids...)
	return cu
}

// RemoveMessages removes "messages" edges to Message entities.
func (cu *ChatUpdate) RemoveMessages(m ...*Message) *ChatUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChatUpdate) check() error {
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chat.owner"`)
	}
	return nil
}

func (cu *ChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chat.Table,
			Columns: chat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: chat.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChatUpdateOne is the builder for updating a single Chat entity.
type ChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMutation
}

// SetOwnerID sets the "ownerID" field.
func (cuo *ChatUpdateOne) SetOwnerID(i int64) *ChatUpdateOne {
	cuo.mutation.SetOwnerID(i)
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *ChatUpdateOne) SetOwner(u *User) *ChatUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cuo *ChatUpdateOne) AddMessageIDs(ids ...int64) *ChatUpdateOne {
	cuo.mutation.AddMessageIDs(ids...)
	return cuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (cuo *ChatUpdateOne) AddMessages(m ...*Message) *ChatUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMessageIDs(ids...)
}

// Mutation returns the ChatMutation object of the builder.
func (cuo *ChatUpdateOne) Mutation() *ChatMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *ChatUpdateOne) ClearOwner() *ChatUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cuo *ChatUpdateOne) ClearMessages() *ChatUpdateOne {
	cuo.mutation.ClearMessages()
	return cuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cuo *ChatUpdateOne) RemoveMessageIDs(ids ...int64) *ChatUpdateOne {
	cuo.mutation.RemoveMessageIDs(ids...)
	return cuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (cuo *ChatUpdateOne) RemoveMessages(m ...*Message) *ChatUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMessageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChatUpdateOne) Select(field string, fields ...string) *ChatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chat entity.
func (cuo *ChatUpdateOne) Save(ctx context.Context) (*Chat, error) {
	var (
		err  error
		node *Chat
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChatUpdateOne) SaveX(ctx context.Context) *Chat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChatUpdateOne) check() error {
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chat.owner"`)
	}
	return nil
}

func (cuo *ChatUpdateOne) sqlSave(ctx context.Context) (_node *Chat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chat.Table,
			Columns: chat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: chat.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chat.FieldID)
		for _, f := range fields {
			if !chat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chat.OwnerTable,
			Columns: []string{chat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: []string{chat.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
