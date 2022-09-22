// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aegis/ent/predicate"
	"aegis/ent/tgonft"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoNFTUpdate is the builder for updating TGoNFT entities.
type TGoNFTUpdate struct {
	config
	hooks    []Hook
	mutation *TGoNFTMutation
}

// Where appends a list predicates to the TGoNFTUpdate builder.
func (tnu *TGoNFTUpdate) Where(ps ...predicate.TGoNFT) *TGoNFTUpdate {
	tnu.mutation.Where(ps...)
	return tnu
}

// SetWalletPub sets the "wallet_pub" field.
func (tnu *TGoNFTUpdate) SetWalletPub(s string) *TGoNFTUpdate {
	tnu.mutation.SetWalletPub(s)
	return tnu
}

// SetNillableWalletPub sets the "wallet_pub" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableWalletPub(s *string) *TGoNFTUpdate {
	if s != nil {
		tnu.SetWalletPub(*s)
	}
	return tnu
}

// SetRankType sets the "rank_type" field.
func (tnu *TGoNFTUpdate) SetRankType(i int) *TGoNFTUpdate {
	tnu.mutation.ResetRankType()
	tnu.mutation.SetRankType(i)
	return tnu
}

// SetNillableRankType sets the "rank_type" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableRankType(i *int) *TGoNFTUpdate {
	if i != nil {
		tnu.SetRankType(*i)
	}
	return tnu
}

// AddRankType adds i to the "rank_type" field.
func (tnu *TGoNFTUpdate) AddRankType(i int) *TGoNFTUpdate {
	tnu.mutation.AddRankType(i)
	return tnu
}

// SetRankYear sets the "rank_year" field.
func (tnu *TGoNFTUpdate) SetRankYear(i int) *TGoNFTUpdate {
	tnu.mutation.ResetRankYear()
	tnu.mutation.SetRankYear(i)
	return tnu
}

// SetNillableRankYear sets the "rank_year" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableRankYear(i *int) *TGoNFTUpdate {
	if i != nil {
		tnu.SetRankYear(*i)
	}
	return tnu
}

// AddRankYear adds i to the "rank_year" field.
func (tnu *TGoNFTUpdate) AddRankYear(i int) *TGoNFTUpdate {
	tnu.mutation.AddRankYear(i)
	return tnu
}

// SetRankSeason sets the "rank_season" field.
func (tnu *TGoNFTUpdate) SetRankSeason(i int) *TGoNFTUpdate {
	tnu.mutation.ResetRankSeason()
	tnu.mutation.SetRankSeason(i)
	return tnu
}

// SetNillableRankSeason sets the "rank_season" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableRankSeason(i *int) *TGoNFTUpdate {
	if i != nil {
		tnu.SetRankSeason(*i)
	}
	return tnu
}

// AddRankSeason adds i to the "rank_season" field.
func (tnu *TGoNFTUpdate) AddRankSeason(i int) *TGoNFTUpdate {
	tnu.mutation.AddRankSeason(i)
	return tnu
}

// SetRank sets the "rank" field.
func (tnu *TGoNFTUpdate) SetRank(i int) *TGoNFTUpdate {
	tnu.mutation.ResetRank()
	tnu.mutation.SetRank(i)
	return tnu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableRank(i *int) *TGoNFTUpdate {
	if i != nil {
		tnu.SetRank(*i)
	}
	return tnu
}

// AddRank adds i to the "rank" field.
func (tnu *TGoNFTUpdate) AddRank(i int) *TGoNFTUpdate {
	tnu.mutation.AddRank(i)
	return tnu
}

// SetMintTx sets the "mint_tx" field.
func (tnu *TGoNFTUpdate) SetMintTx(s string) *TGoNFTUpdate {
	tnu.mutation.SetMintTx(s)
	return tnu
}

// SetNillableMintTx sets the "mint_tx" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableMintTx(s *string) *TGoNFTUpdate {
	if s != nil {
		tnu.SetMintTx(*s)
	}
	return tnu
}

// SetMtime sets the "mtime" field.
func (tnu *TGoNFTUpdate) SetMtime(t time.Time) *TGoNFTUpdate {
	tnu.mutation.SetMtime(t)
	return tnu
}

// SetCtime sets the "ctime" field.
func (tnu *TGoNFTUpdate) SetCtime(t time.Time) *TGoNFTUpdate {
	tnu.mutation.SetCtime(t)
	return tnu
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tnu *TGoNFTUpdate) SetNillableCtime(t *time.Time) *TGoNFTUpdate {
	if t != nil {
		tnu.SetCtime(*t)
	}
	return tnu
}

// Mutation returns the TGoNFTMutation object of the builder.
func (tnu *TGoNFTUpdate) Mutation() *TGoNFTMutation {
	return tnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tnu *TGoNFTUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tnu.defaults()
	if len(tnu.hooks) == 0 {
		affected, err = tnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoNFTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tnu.mutation = mutation
			affected, err = tnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tnu.hooks) - 1; i >= 0; i-- {
			if tnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tnu *TGoNFTUpdate) SaveX(ctx context.Context) int {
	affected, err := tnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tnu *TGoNFTUpdate) Exec(ctx context.Context) error {
	_, err := tnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tnu *TGoNFTUpdate) ExecX(ctx context.Context) {
	if err := tnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tnu *TGoNFTUpdate) defaults() {
	if _, ok := tnu.mutation.Mtime(); !ok {
		v := tgonft.UpdateDefaultMtime()
		tnu.mutation.SetMtime(v)
	}
}

func (tnu *TGoNFTUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgonft.Table,
			Columns: tgonft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgonft.FieldID,
			},
		},
	}
	if ps := tnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tnu.mutation.WalletPub(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldWalletPub,
		})
	}
	if value, ok := tnu.mutation.RankType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankType,
		})
	}
	if value, ok := tnu.mutation.AddedRankType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankType,
		})
	}
	if value, ok := tnu.mutation.RankYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankYear,
		})
	}
	if value, ok := tnu.mutation.AddedRankYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankYear,
		})
	}
	if value, ok := tnu.mutation.RankSeason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankSeason,
		})
	}
	if value, ok := tnu.mutation.AddedRankSeason(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankSeason,
		})
	}
	if value, ok := tnu.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRank,
		})
	}
	if value, ok := tnu.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRank,
		})
	}
	if value, ok := tnu.mutation.MintTx(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldMintTx,
		})
	}
	if value, ok := tnu.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldMtime,
		})
	}
	if value, ok := tnu.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldCtime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tgonft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TGoNFTUpdateOne is the builder for updating a single TGoNFT entity.
type TGoNFTUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TGoNFTMutation
}

// SetWalletPub sets the "wallet_pub" field.
func (tnuo *TGoNFTUpdateOne) SetWalletPub(s string) *TGoNFTUpdateOne {
	tnuo.mutation.SetWalletPub(s)
	return tnuo
}

// SetNillableWalletPub sets the "wallet_pub" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableWalletPub(s *string) *TGoNFTUpdateOne {
	if s != nil {
		tnuo.SetWalletPub(*s)
	}
	return tnuo
}

// SetRankType sets the "rank_type" field.
func (tnuo *TGoNFTUpdateOne) SetRankType(i int) *TGoNFTUpdateOne {
	tnuo.mutation.ResetRankType()
	tnuo.mutation.SetRankType(i)
	return tnuo
}

// SetNillableRankType sets the "rank_type" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableRankType(i *int) *TGoNFTUpdateOne {
	if i != nil {
		tnuo.SetRankType(*i)
	}
	return tnuo
}

// AddRankType adds i to the "rank_type" field.
func (tnuo *TGoNFTUpdateOne) AddRankType(i int) *TGoNFTUpdateOne {
	tnuo.mutation.AddRankType(i)
	return tnuo
}

// SetRankYear sets the "rank_year" field.
func (tnuo *TGoNFTUpdateOne) SetRankYear(i int) *TGoNFTUpdateOne {
	tnuo.mutation.ResetRankYear()
	tnuo.mutation.SetRankYear(i)
	return tnuo
}

// SetNillableRankYear sets the "rank_year" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableRankYear(i *int) *TGoNFTUpdateOne {
	if i != nil {
		tnuo.SetRankYear(*i)
	}
	return tnuo
}

// AddRankYear adds i to the "rank_year" field.
func (tnuo *TGoNFTUpdateOne) AddRankYear(i int) *TGoNFTUpdateOne {
	tnuo.mutation.AddRankYear(i)
	return tnuo
}

// SetRankSeason sets the "rank_season" field.
func (tnuo *TGoNFTUpdateOne) SetRankSeason(i int) *TGoNFTUpdateOne {
	tnuo.mutation.ResetRankSeason()
	tnuo.mutation.SetRankSeason(i)
	return tnuo
}

// SetNillableRankSeason sets the "rank_season" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableRankSeason(i *int) *TGoNFTUpdateOne {
	if i != nil {
		tnuo.SetRankSeason(*i)
	}
	return tnuo
}

// AddRankSeason adds i to the "rank_season" field.
func (tnuo *TGoNFTUpdateOne) AddRankSeason(i int) *TGoNFTUpdateOne {
	tnuo.mutation.AddRankSeason(i)
	return tnuo
}

// SetRank sets the "rank" field.
func (tnuo *TGoNFTUpdateOne) SetRank(i int) *TGoNFTUpdateOne {
	tnuo.mutation.ResetRank()
	tnuo.mutation.SetRank(i)
	return tnuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableRank(i *int) *TGoNFTUpdateOne {
	if i != nil {
		tnuo.SetRank(*i)
	}
	return tnuo
}

// AddRank adds i to the "rank" field.
func (tnuo *TGoNFTUpdateOne) AddRank(i int) *TGoNFTUpdateOne {
	tnuo.mutation.AddRank(i)
	return tnuo
}

// SetMintTx sets the "mint_tx" field.
func (tnuo *TGoNFTUpdateOne) SetMintTx(s string) *TGoNFTUpdateOne {
	tnuo.mutation.SetMintTx(s)
	return tnuo
}

// SetNillableMintTx sets the "mint_tx" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableMintTx(s *string) *TGoNFTUpdateOne {
	if s != nil {
		tnuo.SetMintTx(*s)
	}
	return tnuo
}

// SetMtime sets the "mtime" field.
func (tnuo *TGoNFTUpdateOne) SetMtime(t time.Time) *TGoNFTUpdateOne {
	tnuo.mutation.SetMtime(t)
	return tnuo
}

// SetCtime sets the "ctime" field.
func (tnuo *TGoNFTUpdateOne) SetCtime(t time.Time) *TGoNFTUpdateOne {
	tnuo.mutation.SetCtime(t)
	return tnuo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tnuo *TGoNFTUpdateOne) SetNillableCtime(t *time.Time) *TGoNFTUpdateOne {
	if t != nil {
		tnuo.SetCtime(*t)
	}
	return tnuo
}

// Mutation returns the TGoNFTMutation object of the builder.
func (tnuo *TGoNFTUpdateOne) Mutation() *TGoNFTMutation {
	return tnuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tnuo *TGoNFTUpdateOne) Select(field string, fields ...string) *TGoNFTUpdateOne {
	tnuo.fields = append([]string{field}, fields...)
	return tnuo
}

// Save executes the query and returns the updated TGoNFT entity.
func (tnuo *TGoNFTUpdateOne) Save(ctx context.Context) (*TGoNFT, error) {
	var (
		err  error
		node *TGoNFT
	)
	tnuo.defaults()
	if len(tnuo.hooks) == 0 {
		node, err = tnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoNFTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tnuo.mutation = mutation
			node, err = tnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tnuo.hooks) - 1; i >= 0; i-- {
			if tnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tnuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tnuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TGoNFT)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TGoNFTMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tnuo *TGoNFTUpdateOne) SaveX(ctx context.Context) *TGoNFT {
	node, err := tnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tnuo *TGoNFTUpdateOne) Exec(ctx context.Context) error {
	_, err := tnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tnuo *TGoNFTUpdateOne) ExecX(ctx context.Context) {
	if err := tnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tnuo *TGoNFTUpdateOne) defaults() {
	if _, ok := tnuo.mutation.Mtime(); !ok {
		v := tgonft.UpdateDefaultMtime()
		tnuo.mutation.SetMtime(v)
	}
}

func (tnuo *TGoNFTUpdateOne) sqlSave(ctx context.Context) (_node *TGoNFT, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgonft.Table,
			Columns: tgonft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgonft.FieldID,
			},
		},
	}
	id, ok := tnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TGoNFT.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tgonft.FieldID)
		for _, f := range fields {
			if !tgonft.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tgonft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tnuo.mutation.WalletPub(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldWalletPub,
		})
	}
	if value, ok := tnuo.mutation.RankType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankType,
		})
	}
	if value, ok := tnuo.mutation.AddedRankType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankType,
		})
	}
	if value, ok := tnuo.mutation.RankYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankYear,
		})
	}
	if value, ok := tnuo.mutation.AddedRankYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankYear,
		})
	}
	if value, ok := tnuo.mutation.RankSeason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankSeason,
		})
	}
	if value, ok := tnuo.mutation.AddedRankSeason(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankSeason,
		})
	}
	if value, ok := tnuo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRank,
		})
	}
	if value, ok := tnuo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRank,
		})
	}
	if value, ok := tnuo.mutation.MintTx(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldMintTx,
		})
	}
	if value, ok := tnuo.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldMtime,
		})
	}
	if value, ok := tnuo.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldCtime,
		})
	}
	_node = &TGoNFT{config: tnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tgonft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
