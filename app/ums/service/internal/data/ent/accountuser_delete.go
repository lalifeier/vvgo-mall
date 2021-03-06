// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo-mall/app/ums/service/internal/data/ent/predicate"
)

// AccountUserDelete is the builder for deleting a AccountUser entity.
type AccountUserDelete struct {
	config
	hooks    []Hook
	mutation *AccountUserMutation
}

// Where appends a list predicates to the AccountUserDelete builder.
func (aud *AccountUserDelete) Where(ps ...predicate.AccountUser) *AccountUserDelete {
	aud.mutation.Where(ps...)
	return aud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aud *AccountUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aud.hooks) == 0 {
		affected, err = aud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aud.mutation = mutation
			affected, err = aud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aud.hooks) - 1; i >= 0; i-- {
			if aud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aud *AccountUserDelete) ExecX(ctx context.Context) int {
	n, err := aud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aud *AccountUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: accountuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: accountuser.FieldID,
			},
		},
	}
	if ps := aud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aud.driver, _spec)
}

// AccountUserDeleteOne is the builder for deleting a single AccountUser entity.
type AccountUserDeleteOne struct {
	aud *AccountUserDelete
}

// Exec executes the deletion query.
func (audo *AccountUserDeleteOne) Exec(ctx context.Context) error {
	n, err := audo.aud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (audo *AccountUserDeleteOne) ExecX(ctx context.Context) {
	audo.aud.ExecX(ctx)
}
