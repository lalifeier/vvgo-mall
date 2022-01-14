// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo/app/sys/service/internal/data/ent/dict"
)

// DictCreate is the builder for creating a Dict entity.
type DictCreate struct {
	config
	mutation *DictMutation
	hooks    []Hook
}

// SetDictTypeID sets the "dict_type_id" field.
func (dc *DictCreate) SetDictTypeID(i int64) *DictCreate {
	dc.mutation.SetDictTypeID(i)
	return dc
}

// SetType sets the "type" field.
func (dc *DictCreate) SetType(s string) *DictCreate {
	dc.mutation.SetType(s)
	return dc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dc *DictCreate) SetNillableType(s *string) *DictCreate {
	if s != nil {
		dc.SetType(*s)
	}
	return dc
}

// SetLabel sets the "label" field.
func (dc *DictCreate) SetLabel(s string) *DictCreate {
	dc.mutation.SetLabel(s)
	return dc
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (dc *DictCreate) SetNillableLabel(s *string) *DictCreate {
	if s != nil {
		dc.SetLabel(*s)
	}
	return dc
}

// SetValue sets the "value" field.
func (dc *DictCreate) SetValue(s string) *DictCreate {
	dc.mutation.SetValue(s)
	return dc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (dc *DictCreate) SetNillableValue(s *string) *DictCreate {
	if s != nil {
		dc.SetValue(*s)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DictCreate) SetStatus(i int8) *DictCreate {
	dc.mutation.SetStatus(i)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DictCreate) SetNillableStatus(i *int8) *DictCreate {
	if i != nil {
		dc.SetStatus(*i)
	}
	return dc
}

// SetRemark sets the "remark" field.
func (dc *DictCreate) SetRemark(s string) *DictCreate {
	dc.mutation.SetRemark(s)
	return dc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dc *DictCreate) SetNillableRemark(s *string) *DictCreate {
	if s != nil {
		dc.SetRemark(*s)
	}
	return dc
}

// SetSort sets the "sort" field.
func (dc *DictCreate) SetSort(i int8) *DictCreate {
	dc.mutation.SetSort(i)
	return dc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dc *DictCreate) SetNillableSort(i *int8) *DictCreate {
	if i != nil {
		dc.SetSort(*i)
	}
	return dc
}

// SetIsDefault sets the "is_default" field.
func (dc *DictCreate) SetIsDefault(i int8) *DictCreate {
	dc.mutation.SetIsDefault(i)
	return dc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (dc *DictCreate) SetNillableIsDefault(i *int8) *DictCreate {
	if i != nil {
		dc.SetIsDefault(*i)
	}
	return dc
}

// SetCreateAt sets the "create_at" field.
func (dc *DictCreate) SetCreateAt(i int32) *DictCreate {
	dc.mutation.SetCreateAt(i)
	return dc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (dc *DictCreate) SetNillableCreateAt(i *int32) *DictCreate {
	if i != nil {
		dc.SetCreateAt(*i)
	}
	return dc
}

// SetCreateBy sets the "create_by" field.
func (dc *DictCreate) SetCreateBy(i int32) *DictCreate {
	dc.mutation.SetCreateBy(i)
	return dc
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (dc *DictCreate) SetNillableCreateBy(i *int32) *DictCreate {
	if i != nil {
		dc.SetCreateBy(*i)
	}
	return dc
}

// SetUpdateAt sets the "update_at" field.
func (dc *DictCreate) SetUpdateAt(i int32) *DictCreate {
	dc.mutation.SetUpdateAt(i)
	return dc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (dc *DictCreate) SetNillableUpdateAt(i *int32) *DictCreate {
	if i != nil {
		dc.SetUpdateAt(*i)
	}
	return dc
}

// SetUpdateBy sets the "update_by" field.
func (dc *DictCreate) SetUpdateBy(i int32) *DictCreate {
	dc.mutation.SetUpdateBy(i)
	return dc
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (dc *DictCreate) SetNillableUpdateBy(i *int32) *DictCreate {
	if i != nil {
		dc.SetUpdateBy(*i)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DictCreate) SetID(i int64) *DictCreate {
	dc.mutation.SetID(i)
	return dc
}

// Mutation returns the DictMutation object of the builder.
func (dc *DictCreate) Mutation() *DictMutation {
	return dc.mutation
}

// Save creates the Dict in the database.
func (dc *DictCreate) Save(ctx context.Context) (*Dict, error) {
	var (
		err  error
		node *Dict
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DictCreate) SaveX(ctx context.Context) *Dict {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DictCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DictCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DictCreate) defaults() {
	if _, ok := dc.mutation.GetType(); !ok {
		v := dict.DefaultType
		dc.mutation.SetType(v)
	}
	if _, ok := dc.mutation.Label(); !ok {
		v := dict.DefaultLabel
		dc.mutation.SetLabel(v)
	}
	if _, ok := dc.mutation.Value(); !ok {
		v := dict.DefaultValue
		dc.mutation.SetValue(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := dict.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Remark(); !ok {
		v := dict.DefaultRemark
		dc.mutation.SetRemark(v)
	}
	if _, ok := dc.mutation.Sort(); !ok {
		v := dict.DefaultSort
		dc.mutation.SetSort(v)
	}
	if _, ok := dc.mutation.IsDefault(); !ok {
		v := dict.DefaultIsDefault
		dc.mutation.SetIsDefault(v)
	}
	if _, ok := dc.mutation.CreateAt(); !ok {
		v := dict.DefaultCreateAt
		dc.mutation.SetCreateAt(v)
	}
	if _, ok := dc.mutation.CreateBy(); !ok {
		v := dict.DefaultCreateBy
		dc.mutation.SetCreateBy(v)
	}
	if _, ok := dc.mutation.UpdateAt(); !ok {
		v := dict.DefaultUpdateAt
		dc.mutation.SetUpdateAt(v)
	}
	if _, ok := dc.mutation.UpdateBy(); !ok {
		v := dict.DefaultUpdateBy
		dc.mutation.SetUpdateBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DictCreate) check() error {
	if _, ok := dc.mutation.DictTypeID(); !ok {
		return &ValidationError{Name: "dict_type_id", err: errors.New(`ent: missing required field "dict_type_id"`)}
	}
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if _, ok := dc.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "label"`)}
	}
	if _, ok := dc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "value"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := dc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "remark"`)}
	}
	if _, ok := dc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "sort"`)}
	}
	if _, ok := dc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "is_default"`)}
	}
	if _, ok := dc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := dc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "create_by", err: errors.New(`ent: missing required field "create_by"`)}
	}
	if _, ok := dc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := dc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "update_by"`)}
	}
	return nil
}

func (dc *DictCreate) sqlSave(ctx context.Context) (*Dict, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (dc *DictCreate) createSpec() (*Dict, *sqlgraph.CreateSpec) {
	var (
		_node = &Dict{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dict.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: dict.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.DictTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: dict.FieldDictTypeID,
		})
		_node.DictTypeID = value
	}
	if value, ok := dc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dict.FieldType,
		})
		_node.Type = value
	}
	if value, ok := dc.mutation.Label(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dict.FieldLabel,
		})
		_node.Label = value
	}
	if value, ok := dc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dict.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dict.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := dc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dict.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := dc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dict.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := dc.mutation.IsDefault(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dict.FieldIsDefault,
		})
		_node.IsDefault = value
	}
	if value, ok := dc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: dict.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := dc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: dict.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := dc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: dict.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := dc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: dict.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	return _node, _spec
}

// DictCreateBulk is the builder for creating many Dict entities in bulk.
type DictCreateBulk struct {
	config
	builders []*DictCreate
}

// Save creates the Dict entities in the database.
func (dcb *DictCreateBulk) Save(ctx context.Context) ([]*Dict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dict, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DictCreateBulk) SaveX(ctx context.Context) []*Dict {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DictCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DictCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
