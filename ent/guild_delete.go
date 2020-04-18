// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/roleypoly/db/ent/guild"
	"github.com/roleypoly/db/ent/predicate"
)

// GuildDelete is the builder for deleting a Guild entity.
type GuildDelete struct {
	config
	hooks      []Hook
	mutation   *GuildMutation
	predicates []predicate.Guild
}

// Where adds a new predicate to the delete builder.
func (gd *GuildDelete) Where(ps ...predicate.Guild) *GuildDelete {
	gd.predicates = append(gd.predicates, ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GuildDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gd.hooks) == 0 {
		affected, err = gd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gd.mutation = mutation
			affected, err = gd.sqlExec(ctx)
			return affected, err
		})
		for i := len(gd.hooks) - 1; i >= 0; i-- {
			mut = gd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GuildDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GuildDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: guild.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guild.FieldID,
			},
		},
	}
	if ps := gd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GuildDeleteOne is the builder for deleting a single Guild entity.
type GuildDeleteOne struct {
	gd *GuildDelete
}

// Exec executes the deletion query.
func (gdo *GuildDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{guild.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GuildDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}
