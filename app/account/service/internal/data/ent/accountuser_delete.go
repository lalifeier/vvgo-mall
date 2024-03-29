// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/predicate"
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
	return withHooks[int, AccountUserMutation](ctx, aud.sqlExec, aud.mutation, aud.hooks)
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
	_spec := sqlgraph.NewDeleteSpec(accountuser.Table, sqlgraph.NewFieldSpec(accountuser.FieldID, field.TypeUint32))
	if ps := aud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aud.mutation.done = true
	return affected, err
}

// AccountUserDeleteOne is the builder for deleting a single AccountUser entity.
type AccountUserDeleteOne struct {
	aud *AccountUserDelete
}

// Where appends a list predicates to the AccountUserDelete builder.
func (audo *AccountUserDeleteOne) Where(ps ...predicate.AccountUser) *AccountUserDeleteOne {
	audo.aud.mutation.Where(ps...)
	return audo
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
	if err := audo.Exec(ctx); err != nil {
		panic(err)
	}
}
