// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aegis/ent/atasklog"
	"aegis/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ATaskLogUpdate is the builder for updating ATaskLog entities.
type ATaskLogUpdate struct {
	config
	hooks    []Hook
	mutation *ATaskLogMutation
}

// Where appends a list predicates to the ATaskLogUpdate builder.
func (alu *ATaskLogUpdate) Where(ps ...predicate.ATaskLog) *ATaskLogUpdate {
	alu.mutation.Where(ps...)
	return alu
}

// SetQuestID sets the "quest_id" field.
func (alu *ATaskLogUpdate) SetQuestID(u uint64) *ATaskLogUpdate {
	alu.mutation.ResetQuestID()
	alu.mutation.SetQuestID(u)
	return alu
}

// SetNillableQuestID sets the "quest_id" field if the given value is not nil.
func (alu *ATaskLogUpdate) SetNillableQuestID(u *uint64) *ATaskLogUpdate {
	if u != nil {
		alu.SetQuestID(*u)
	}
	return alu
}

// AddQuestID adds u to the "quest_id" field.
func (alu *ATaskLogUpdate) AddQuestID(u int64) *ATaskLogUpdate {
	alu.mutation.AddQuestID(u)
	return alu
}

// SetTaskID sets the "task_id" field.
func (alu *ATaskLogUpdate) SetTaskID(u uint64) *ATaskLogUpdate {
	alu.mutation.ResetTaskID()
	alu.mutation.SetTaskID(u)
	return alu
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (alu *ATaskLogUpdate) SetNillableTaskID(u *uint64) *ATaskLogUpdate {
	if u != nil {
		alu.SetTaskID(*u)
	}
	return alu
}

// AddTaskID adds u to the "task_id" field.
func (alu *ATaskLogUpdate) AddTaskID(u int64) *ATaskLogUpdate {
	alu.mutation.AddTaskID(u)
	return alu
}

// SetMid sets the "mid" field.
func (alu *ATaskLogUpdate) SetMid(u uint64) *ATaskLogUpdate {
	alu.mutation.ResetMid()
	alu.mutation.SetMid(u)
	return alu
}

// SetNillableMid sets the "mid" field if the given value is not nil.
func (alu *ATaskLogUpdate) SetNillableMid(u *uint64) *ATaskLogUpdate {
	if u != nil {
		alu.SetMid(*u)
	}
	return alu
}

// AddMid adds u to the "mid" field.
func (alu *ATaskLogUpdate) AddMid(u int64) *ATaskLogUpdate {
	alu.mutation.AddMid(u)
	return alu
}

// SetMeta sets the "meta" field.
func (alu *ATaskLogUpdate) SetMeta(s string) *ATaskLogUpdate {
	alu.mutation.SetMeta(s)
	return alu
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (alu *ATaskLogUpdate) SetNillableMeta(s *string) *ATaskLogUpdate {
	if s != nil {
		alu.SetMeta(*s)
	}
	return alu
}

// SetMtime sets the "mtime" field.
func (alu *ATaskLogUpdate) SetMtime(t time.Time) *ATaskLogUpdate {
	alu.mutation.SetMtime(t)
	return alu
}

// SetCtime sets the "ctime" field.
func (alu *ATaskLogUpdate) SetCtime(t time.Time) *ATaskLogUpdate {
	alu.mutation.SetCtime(t)
	return alu
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (alu *ATaskLogUpdate) SetNillableCtime(t *time.Time) *ATaskLogUpdate {
	if t != nil {
		alu.SetCtime(*t)
	}
	return alu
}

// Mutation returns the ATaskLogMutation object of the builder.
func (alu *ATaskLogUpdate) Mutation() *ATaskLogMutation {
	return alu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alu *ATaskLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	alu.defaults()
	if len(alu.hooks) == 0 {
		affected, err = alu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ATaskLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			alu.mutation = mutation
			affected, err = alu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(alu.hooks) - 1; i >= 0; i-- {
			if alu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, alu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (alu *ATaskLogUpdate) SaveX(ctx context.Context) int {
	affected, err := alu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alu *ATaskLogUpdate) Exec(ctx context.Context) error {
	_, err := alu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alu *ATaskLogUpdate) ExecX(ctx context.Context) {
	if err := alu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alu *ATaskLogUpdate) defaults() {
	if _, ok := alu.mutation.Mtime(); !ok {
		v := atasklog.UpdateDefaultMtime()
		alu.mutation.SetMtime(v)
	}
}

func (alu *ATaskLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   atasklog.Table,
			Columns: atasklog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: atasklog.FieldID,
			},
		},
	}
	if ps := alu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alu.mutation.QuestID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldQuestID,
		})
	}
	if value, ok := alu.mutation.AddedQuestID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldQuestID,
		})
	}
	if value, ok := alu.mutation.TaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldTaskID,
		})
	}
	if value, ok := alu.mutation.AddedTaskID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldTaskID,
		})
	}
	if value, ok := alu.mutation.Mid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldMid,
		})
	}
	if value, ok := alu.mutation.AddedMid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldMid,
		})
	}
	if value, ok := alu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: atasklog.FieldMeta,
		})
	}
	if value, ok := alu.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: atasklog.FieldMtime,
		})
	}
	if value, ok := alu.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: atasklog.FieldCtime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, alu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{atasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ATaskLogUpdateOne is the builder for updating a single ATaskLog entity.
type ATaskLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ATaskLogMutation
}

// SetQuestID sets the "quest_id" field.
func (aluo *ATaskLogUpdateOne) SetQuestID(u uint64) *ATaskLogUpdateOne {
	aluo.mutation.ResetQuestID()
	aluo.mutation.SetQuestID(u)
	return aluo
}

// SetNillableQuestID sets the "quest_id" field if the given value is not nil.
func (aluo *ATaskLogUpdateOne) SetNillableQuestID(u *uint64) *ATaskLogUpdateOne {
	if u != nil {
		aluo.SetQuestID(*u)
	}
	return aluo
}

// AddQuestID adds u to the "quest_id" field.
func (aluo *ATaskLogUpdateOne) AddQuestID(u int64) *ATaskLogUpdateOne {
	aluo.mutation.AddQuestID(u)
	return aluo
}

// SetTaskID sets the "task_id" field.
func (aluo *ATaskLogUpdateOne) SetTaskID(u uint64) *ATaskLogUpdateOne {
	aluo.mutation.ResetTaskID()
	aluo.mutation.SetTaskID(u)
	return aluo
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (aluo *ATaskLogUpdateOne) SetNillableTaskID(u *uint64) *ATaskLogUpdateOne {
	if u != nil {
		aluo.SetTaskID(*u)
	}
	return aluo
}

// AddTaskID adds u to the "task_id" field.
func (aluo *ATaskLogUpdateOne) AddTaskID(u int64) *ATaskLogUpdateOne {
	aluo.mutation.AddTaskID(u)
	return aluo
}

// SetMid sets the "mid" field.
func (aluo *ATaskLogUpdateOne) SetMid(u uint64) *ATaskLogUpdateOne {
	aluo.mutation.ResetMid()
	aluo.mutation.SetMid(u)
	return aluo
}

// SetNillableMid sets the "mid" field if the given value is not nil.
func (aluo *ATaskLogUpdateOne) SetNillableMid(u *uint64) *ATaskLogUpdateOne {
	if u != nil {
		aluo.SetMid(*u)
	}
	return aluo
}

// AddMid adds u to the "mid" field.
func (aluo *ATaskLogUpdateOne) AddMid(u int64) *ATaskLogUpdateOne {
	aluo.mutation.AddMid(u)
	return aluo
}

// SetMeta sets the "meta" field.
func (aluo *ATaskLogUpdateOne) SetMeta(s string) *ATaskLogUpdateOne {
	aluo.mutation.SetMeta(s)
	return aluo
}

// SetNillableMeta sets the "meta" field if the given value is not nil.
func (aluo *ATaskLogUpdateOne) SetNillableMeta(s *string) *ATaskLogUpdateOne {
	if s != nil {
		aluo.SetMeta(*s)
	}
	return aluo
}

// SetMtime sets the "mtime" field.
func (aluo *ATaskLogUpdateOne) SetMtime(t time.Time) *ATaskLogUpdateOne {
	aluo.mutation.SetMtime(t)
	return aluo
}

// SetCtime sets the "ctime" field.
func (aluo *ATaskLogUpdateOne) SetCtime(t time.Time) *ATaskLogUpdateOne {
	aluo.mutation.SetCtime(t)
	return aluo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (aluo *ATaskLogUpdateOne) SetNillableCtime(t *time.Time) *ATaskLogUpdateOne {
	if t != nil {
		aluo.SetCtime(*t)
	}
	return aluo
}

// Mutation returns the ATaskLogMutation object of the builder.
func (aluo *ATaskLogUpdateOne) Mutation() *ATaskLogMutation {
	return aluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aluo *ATaskLogUpdateOne) Select(field string, fields ...string) *ATaskLogUpdateOne {
	aluo.fields = append([]string{field}, fields...)
	return aluo
}

// Save executes the query and returns the updated ATaskLog entity.
func (aluo *ATaskLogUpdateOne) Save(ctx context.Context) (*ATaskLog, error) {
	var (
		err  error
		node *ATaskLog
	)
	aluo.defaults()
	if len(aluo.hooks) == 0 {
		node, err = aluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ATaskLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aluo.mutation = mutation
			node, err = aluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aluo.hooks) - 1; i >= 0; i-- {
			if aluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ATaskLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ATaskLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aluo *ATaskLogUpdateOne) SaveX(ctx context.Context) *ATaskLog {
	node, err := aluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aluo *ATaskLogUpdateOne) Exec(ctx context.Context) error {
	_, err := aluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aluo *ATaskLogUpdateOne) ExecX(ctx context.Context) {
	if err := aluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aluo *ATaskLogUpdateOne) defaults() {
	if _, ok := aluo.mutation.Mtime(); !ok {
		v := atasklog.UpdateDefaultMtime()
		aluo.mutation.SetMtime(v)
	}
}

func (aluo *ATaskLogUpdateOne) sqlSave(ctx context.Context) (_node *ATaskLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   atasklog.Table,
			Columns: atasklog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: atasklog.FieldID,
			},
		},
	}
	id, ok := aluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ATaskLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, atasklog.FieldID)
		for _, f := range fields {
			if !atasklog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != atasklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aluo.mutation.QuestID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldQuestID,
		})
	}
	if value, ok := aluo.mutation.AddedQuestID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldQuestID,
		})
	}
	if value, ok := aluo.mutation.TaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldTaskID,
		})
	}
	if value, ok := aluo.mutation.AddedTaskID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldTaskID,
		})
	}
	if value, ok := aluo.mutation.Mid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldMid,
		})
	}
	if value, ok := aluo.mutation.AddedMid(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: atasklog.FieldMid,
		})
	}
	if value, ok := aluo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: atasklog.FieldMeta,
		})
	}
	if value, ok := aluo.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: atasklog.FieldMtime,
		})
	}
	if value, ok := aluo.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: atasklog.FieldCtime,
		})
	}
	_node = &ATaskLog{config: aluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{atasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
