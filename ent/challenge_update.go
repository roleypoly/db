// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/roleypoly/db/ent/challenge"
	"github.com/roleypoly/db/ent/predicate"
)

// ChallengeUpdate is the builder for updating Challenge entities.
type ChallengeUpdate struct {
	config
	hooks      []Hook
	mutation   *ChallengeMutation
	predicates []predicate.Challenge
}

// Where adds a new predicate for the builder.
func (cu *ChallengeUpdate) Where(ps ...predicate.Challenge) *ChallengeUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// Mutation returns the ChallengeMutation object of the builder.
func (cu *ChallengeUpdate) Mutation() *ChallengeMutation {
	return cu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *ChallengeUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := challenge.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChallengeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChallengeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChallengeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChallengeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChallengeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: challenge.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ChallengeUpdateOne is the builder for updating a single Challenge entity.
type ChallengeUpdateOne struct {
	config
	hooks    []Hook
	mutation *ChallengeMutation
}

// Mutation returns the ChallengeMutation object of the builder.
func (cuo *ChallengeUpdateOne) Mutation() *ChallengeMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated entity.
func (cuo *ChallengeUpdateOne) Save(ctx context.Context) (*Challenge, error) {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := challenge.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
	var (
		err  error
		node *Challenge
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChallengeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) SaveX(ctx context.Context) *Challenge {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *ChallengeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChallengeUpdateOne) sqlSave(ctx context.Context) (c *Challenge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: challenge.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Challenge.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldUpdateTime,
		})
	}
	c = &Challenge{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
