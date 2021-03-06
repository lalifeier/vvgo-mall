// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/dicttype"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/predicate"
)

// DictTypeQuery is the builder for querying DictType entities.
type DictTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DictType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictTypeQuery builder.
func (dtq *DictTypeQuery) Where(ps ...predicate.DictType) *DictTypeQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit adds a limit step to the query.
func (dtq *DictTypeQuery) Limit(limit int) *DictTypeQuery {
	dtq.limit = &limit
	return dtq
}

// Offset adds an offset step to the query.
func (dtq *DictTypeQuery) Offset(offset int) *DictTypeQuery {
	dtq.offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DictTypeQuery) Unique(unique bool) *DictTypeQuery {
	dtq.unique = &unique
	return dtq
}

// Order adds an order step to the query.
func (dtq *DictTypeQuery) Order(o ...OrderFunc) *DictTypeQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// First returns the first DictType entity from the query.
// Returns a *NotFoundError when no DictType was found.
func (dtq *DictTypeQuery) First(ctx context.Context) (*DictType, error) {
	nodes, err := dtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dicttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DictTypeQuery) FirstX(ctx context.Context) *DictType {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DictType ID from the query.
// Returns a *NotFoundError when no DictType ID was found.
func (dtq *DictTypeQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dicttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DictTypeQuery) FirstIDX(ctx context.Context) int64 {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DictType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DictType entity is found.
// Returns a *NotFoundError when no DictType entities are found.
func (dtq *DictTypeQuery) Only(ctx context.Context) (*DictType, error) {
	nodes, err := dtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dicttype.Label}
	default:
		return nil, &NotSingularError{dicttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DictTypeQuery) OnlyX(ctx context.Context) *DictType {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DictType ID in the query.
// Returns a *NotSingularError when more than one DictType ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DictTypeQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = &NotSingularError{dicttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DictTypeQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DictTypes.
func (dtq *DictTypeQuery) All(ctx context.Context) ([]*DictType, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DictTypeQuery) AllX(ctx context.Context) []*DictType {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DictType IDs.
func (dtq *DictTypeQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := dtq.Select(dicttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DictTypeQuery) IDsX(ctx context.Context) []int64 {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DictTypeQuery) Count(ctx context.Context) (int, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DictTypeQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DictTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DictTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DictTypeQuery) Clone() *DictTypeQuery {
	if dtq == nil {
		return nil
	}
	return &DictTypeQuery{
		config:     dtq.config,
		limit:      dtq.limit,
		offset:     dtq.offset,
		order:      append([]OrderFunc{}, dtq.order...),
		predicates: append([]predicate.DictType{}, dtq.predicates...),
		// clone intermediate query.
		sql:    dtq.sql.Clone(),
		path:   dtq.path,
		unique: dtq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DictType.Query().
//		GroupBy(dicttype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dtq *DictTypeQuery) GroupBy(field string, fields ...string) *DictTypeGroupBy {
	group := &DictTypeGroupBy{config: dtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dtq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DictType.Query().
//		Select(dicttype.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (dtq *DictTypeQuery) Select(fields ...string) *DictTypeSelect {
	dtq.fields = append(dtq.fields, fields...)
	return &DictTypeSelect{DictTypeQuery: dtq}
}

func (dtq *DictTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dtq.fields {
		if !dicttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DictTypeQuery) sqlAll(ctx context.Context) ([]*DictType, error) {
	var (
		nodes = []*DictType{}
		_spec = dtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DictType{config: dtq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dtq *DictTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.fields
	if len(dtq.fields) > 0 {
		_spec.Unique = dtq.unique != nil && *dtq.unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DictTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dtq *DictTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dicttype.Table,
			Columns: dicttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: dicttype.FieldID,
			},
		},
		From:   dtq.sql,
		Unique: true,
	}
	if unique := dtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dicttype.FieldID)
		for i := range fields {
			if fields[i] != dicttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DictTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(dicttype.Table)
	columns := dtq.fields
	if len(columns) == 0 {
		columns = dicttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.unique != nil && *dtq.unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DictTypeGroupBy is the group-by builder for DictType entities.
type DictTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DictTypeGroupBy) Aggregate(fns ...AggregateFunc) *DictTypeGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dtgb *DictTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dtgb.path(ctx)
	if err != nil {
		return err
	}
	dtgb.sql = query
	return dtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DictTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := dtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) StringX(ctx context.Context) string {
	v, err := dtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DictTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := dtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) IntX(ctx context.Context) int {
	v, err := dtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DictTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dtgb.fields) > 1 {
		return nil, errors.New("ent: DictTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dtgb *DictTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dtgb *DictTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := dtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dtgb *DictTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dtgb.fields {
		if !dicttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dtgb *DictTypeGroupBy) sqlQuery() *sql.Selector {
	selector := dtgb.sql.Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dtgb.fields)+len(dtgb.fns))
		for _, f := range dtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dtgb.fields...)...)
}

// DictTypeSelect is the builder for selecting fields of DictType entities.
type DictTypeSelect struct {
	*DictTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DictTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	dts.sql = dts.DictTypeQuery.sqlQuery(ctx)
	return dts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dts *DictTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DictTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dts *DictTypeSelect) StringsX(ctx context.Context) []string {
	v, err := dts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dts *DictTypeSelect) StringX(ctx context.Context) string {
	v, err := dts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DictTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dts *DictTypeSelect) IntsX(ctx context.Context) []int {
	v, err := dts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dts *DictTypeSelect) IntX(ctx context.Context) int {
	v, err := dts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DictTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dts *DictTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dts *DictTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := dts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dts.fields) > 1 {
		return nil, errors.New("ent: DictTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dts *DictTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := dts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dts *DictTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = fmt.Errorf("ent: DictTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dts *DictTypeSelect) BoolX(ctx context.Context) bool {
	v, err := dts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dts *DictTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dts.sql.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
