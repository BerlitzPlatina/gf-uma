// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BerlitzPlatina/gf-uma/internal/ent/game"
	"github.com/BerlitzPlatina/gf-uma/internal/ent/predicate"
)

// GameDelete is the builder for deleting a Game entity.
type GameDelete struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameDelete builder.
func (gd *GameDelete) Where(ps ...predicate.Game) *GameDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GameDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gd.sqlExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GameDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GameDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(game.Table, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gd.mutation.done = true
	return affected, err
}

// GameDeleteOne is the builder for deleting a single Game entity.
type GameDeleteOne struct {
	gd *GameDelete
}

// Where appends a list predicates to the GameDelete builder.
func (gdo *GameDeleteOne) Where(ps ...predicate.Game) *GameDeleteOne {
	gdo.gd.mutation.Where(ps...)
	return gdo
}

// Exec executes the deletion query.
func (gdo *GameDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{game.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GameDeleteOne) ExecX(ctx context.Context) {
	if err := gdo.Exec(ctx); err != nil {
		panic(err)
	}
}
