// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/predicate"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/role"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RoleUpdate) SetCreatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCreatedAt(t *time.Time) *RoleUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetCreatedBy sets the "created_by" field.
func (ru *RoleUpdate) SetCreatedBy(i int64) *RoleUpdate {
	ru.mutation.ResetCreatedBy()
	ru.mutation.SetCreatedBy(i)
	return ru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCreatedBy(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetCreatedBy(*i)
	}
	return ru
}

// AddCreatedBy adds i to the "created_by" field.
func (ru *RoleUpdate) AddCreatedBy(i int64) *RoleUpdate {
	ru.mutation.AddCreatedBy(i)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RoleUpdate) SetUpdatedBy(i int64) *RoleUpdate {
	ru.mutation.ResetUpdatedBy()
	ru.mutation.SetUpdatedBy(i)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdatedBy(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetUpdatedBy(*i)
	}
	return ru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ru *RoleUpdate) AddUpdatedBy(i int64) *RoleUpdate {
	ru.mutation.AddUpdatedBy(i)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetSort sets the "sort" field.
func (ru *RoleUpdate) SetSort(i int8) *RoleUpdate {
	ru.mutation.ResetSort()
	ru.mutation.SetSort(i)
	return ru
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableSort(i *int8) *RoleUpdate {
	if i != nil {
		ru.SetSort(*i)
	}
	return ru
}

// AddSort adds i to the "sort" field.
func (ru *RoleUpdate) AddSort(i int8) *RoleUpdate {
	ru.mutation.AddSort(i)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(i int8) *RoleUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(i)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(i *int8) *RoleUpdate {
	if i != nil {
		ru.SetStatus(*i)
	}
	return ru
}

// AddStatus adds i to the "status" field.
func (ru *RoleUpdate) AddStatus(i int8) *RoleUpdate {
	ru.mutation.AddStatus(i)
	return ru
}

// SetRemark sets the "remark" field.
func (ru *RoleUpdate) SetRemark(s string) *RoleUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableRemark(s *string) *RoleUpdate {
	if s != nil {
		ru.SetRemark(*s)
	}
	return ru
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if role.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := role.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: role.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldCreatedBy,
		})
	}
	if value, ok := ru.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldCreatedBy,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldUpdatedBy,
		})
	}
	if value, ok := ru.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldUpdatedBy,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ru.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldSort,
		})
	}
	if value, ok := ru.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldSort,
		})
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldRemark,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RoleUpdateOne) SetCreatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCreatedAt(t *time.Time) *RoleUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetCreatedBy sets the "created_by" field.
func (ruo *RoleUpdateOne) SetCreatedBy(i int64) *RoleUpdateOne {
	ruo.mutation.ResetCreatedBy()
	ruo.mutation.SetCreatedBy(i)
	return ruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCreatedBy(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetCreatedBy(*i)
	}
	return ruo
}

// AddCreatedBy adds i to the "created_by" field.
func (ruo *RoleUpdateOne) AddCreatedBy(i int64) *RoleUpdateOne {
	ruo.mutation.AddCreatedBy(i)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RoleUpdateOne) SetUpdatedBy(i int64) *RoleUpdateOne {
	ruo.mutation.ResetUpdatedBy()
	ruo.mutation.SetUpdatedBy(i)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdatedBy(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetUpdatedBy(*i)
	}
	return ruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ruo *RoleUpdateOne) AddUpdatedBy(i int64) *RoleUpdateOne {
	ruo.mutation.AddUpdatedBy(i)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetSort sets the "sort" field.
func (ruo *RoleUpdateOne) SetSort(i int8) *RoleUpdateOne {
	ruo.mutation.ResetSort()
	ruo.mutation.SetSort(i)
	return ruo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableSort(i *int8) *RoleUpdateOne {
	if i != nil {
		ruo.SetSort(*i)
	}
	return ruo
}

// AddSort adds i to the "sort" field.
func (ruo *RoleUpdateOne) AddSort(i int8) *RoleUpdateOne {
	ruo.mutation.AddSort(i)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(i int8) *RoleUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(i)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(i *int8) *RoleUpdateOne {
	if i != nil {
		ruo.SetStatus(*i)
	}
	return ruo
}

// AddStatus adds i to the "status" field.
func (ruo *RoleUpdateOne) AddStatus(i int8) *RoleUpdateOne {
	ruo.mutation.AddStatus(i)
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *RoleUpdateOne) SetRemark(s string) *RoleUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableRemark(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetRemark(*s)
	}
	return ruo
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	var (
		err  error
		node *Role
	)
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if role.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized role.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := role.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: role.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldCreatedBy,
		})
	}
	if value, ok := ruo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldCreatedBy,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldUpdatedBy,
		})
	}
	if value, ok := ruo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: role.FieldUpdatedBy,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ruo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldSort,
		})
	}
	if value, ok := ruo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldSort,
		})
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldRemark,
		})
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
