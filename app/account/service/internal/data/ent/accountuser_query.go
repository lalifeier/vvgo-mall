// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/accountuser"
	"github.com/lalifeier/vvgo-mall/app/account/service/internal/data/ent/predicate"
)

// AccountUserQuery is the builder for querying AccountUser entities.
type AccountUserQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.AccountUser
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountUserQuery builder.
func (auq *AccountUserQuery) Where(ps ...predicate.AccountUser) *AccountUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit the number of records to be returned by this query.
func (auq *AccountUserQuery) Limit(limit int) *AccountUserQuery {
	auq.ctx.Limit = &limit
	return auq
}

// Offset to start from.
func (auq *AccountUserQuery) Offset(offset int) *AccountUserQuery {
	auq.ctx.Offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AccountUserQuery) Unique(unique bool) *AccountUserQuery {
	auq.ctx.Unique = &unique
	return auq
}

// Order specifies how the records should be ordered.
func (auq *AccountUserQuery) Order(o ...OrderFunc) *AccountUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AccountUser entity from the query.
// Returns a *NotFoundError when no AccountUser was found.
func (auq *AccountUserQuery) First(ctx context.Context) (*AccountUser, error) {
	nodes, err := auq.Limit(1).All(setContextOp(ctx, auq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AccountUserQuery) FirstX(ctx context.Context) *AccountUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountUser ID from the query.
// Returns a *NotFoundError when no AccountUser ID was found.
func (auq *AccountUserQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = auq.Limit(1).IDs(setContextOp(ctx, auq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AccountUserQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountUser entity is found.
// Returns a *NotFoundError when no AccountUser entities are found.
func (auq *AccountUserQuery) Only(ctx context.Context) (*AccountUser, error) {
	nodes, err := auq.Limit(2).All(setContextOp(ctx, auq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountuser.Label}
	default:
		return nil, &NotSingularError{accountuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AccountUserQuery) OnlyX(ctx context.Context) *AccountUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountUser ID in the query.
// Returns a *NotSingularError when more than one AccountUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (auq *AccountUserQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = auq.Limit(2).IDs(setContextOp(ctx, auq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountuser.Label}
	default:
		err = &NotSingularError{accountuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AccountUserQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountUsers.
func (auq *AccountUserQuery) All(ctx context.Context) ([]*AccountUser, error) {
	ctx = setContextOp(ctx, auq.ctx, "All")
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountUser, *AccountUserQuery]()
	return withInterceptors[[]*AccountUser](ctx, auq, qr, auq.inters)
}

// AllX is like All, but panics if an error occurs.
func (auq *AccountUserQuery) AllX(ctx context.Context) []*AccountUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountUser IDs.
func (auq *AccountUserQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if auq.ctx.Unique == nil && auq.path != nil {
		auq.Unique(true)
	}
	ctx = setContextOp(ctx, auq.ctx, "IDs")
	if err = auq.Select(accountuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AccountUserQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AccountUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, auq.ctx, "Count")
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, auq, querierCount[*AccountUserQuery](), auq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AccountUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AccountUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, auq.ctx, "Exist")
	switch _, err := auq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AccountUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AccountUserQuery) Clone() *AccountUserQuery {
	if auq == nil {
		return nil
	}
	return &AccountUserQuery{
		config:     auq.config,
		ctx:        auq.ctx.Clone(),
		order:      append([]OrderFunc{}, auq.order...),
		inters:     append([]Interceptor{}, auq.inters...),
		predicates: append([]predicate.AccountUser{}, auq.predicates...),
		// clone intermediate query.
		sql:  auq.sql.Clone(),
		path: auq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccountUser.Query().
//		GroupBy(accountuser.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (auq *AccountUserQuery) GroupBy(field string, fields ...string) *AccountUserGroupBy {
	auq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountUserGroupBy{build: auq}
	grbuild.flds = &auq.ctx.Fields
	grbuild.label = accountuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.AccountUser.Query().
//		Select(accountuser.FieldUsername).
//		Scan(ctx, &v)
func (auq *AccountUserQuery) Select(fields ...string) *AccountUserSelect {
	auq.ctx.Fields = append(auq.ctx.Fields, fields...)
	sbuild := &AccountUserSelect{AccountUserQuery: auq}
	sbuild.label = accountuser.Label
	sbuild.flds, sbuild.scan = &auq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountUserSelect configured with the given aggregations.
func (auq *AccountUserQuery) Aggregate(fns ...AggregateFunc) *AccountUserSelect {
	return auq.Select().Aggregate(fns...)
}

func (auq *AccountUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range auq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, auq); err != nil {
				return err
			}
		}
	}
	for _, f := range auq.ctx.Fields {
		if !accountuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AccountUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountUser, error) {
	var (
		nodes = []*AccountUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountUser{config: auq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AccountUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	_spec.Node.Columns = auq.ctx.Fields
	if len(auq.ctx.Fields) > 0 {
		_spec.Unique = auq.ctx.Unique != nil && *auq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AccountUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accountuser.Table, accountuser.Columns, sqlgraph.NewFieldSpec(accountuser.FieldID, field.TypeUint32))
	_spec.From = auq.sql
	if unique := auq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if auq.path != nil {
		_spec.Unique = true
	}
	if fields := auq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountuser.FieldID)
		for i := range fields {
			if fields[i] != accountuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AccountUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(accountuser.Table)
	columns := auq.ctx.Fields
	if len(columns) == 0 {
		columns = accountuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.ctx.Unique != nil && *auq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range auq.modifiers {
		m(selector)
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (auq *AccountUserQuery) Modify(modifiers ...func(s *sql.Selector)) *AccountUserSelect {
	auq.modifiers = append(auq.modifiers, modifiers...)
	return auq.Select()
}

// AccountUserGroupBy is the group-by builder for AccountUser entities.
type AccountUserGroupBy struct {
	selector
	build *AccountUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AccountUserGroupBy) Aggregate(fns ...AggregateFunc) *AccountUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the selector query and scans the result into the given value.
func (augb *AccountUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, augb.build.ctx, "GroupBy")
	if err := augb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountUserQuery, *AccountUserGroupBy](ctx, augb.build, augb, augb.build.inters, v)
}

func (augb *AccountUserGroupBy) sqlScan(ctx context.Context, root *AccountUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*augb.flds)+len(augb.fns))
		for _, f := range *augb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*augb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountUserSelect is the builder for selecting fields of AccountUser entities.
type AccountUserSelect struct {
	*AccountUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aus *AccountUserSelect) Aggregate(fns ...AggregateFunc) *AccountUserSelect {
	aus.fns = append(aus.fns, fns...)
	return aus
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AccountUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aus.ctx, "Select")
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountUserQuery, *AccountUserSelect](ctx, aus.AccountUserQuery, aus, aus.inters, v)
}

func (aus *AccountUserSelect) sqlScan(ctx context.Context, root *AccountUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aus.fns))
	for _, fn := range aus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aus *AccountUserSelect) Modify(modifiers ...func(s *sql.Selector)) *AccountUserSelect {
	aus.modifiers = append(aus.modifiers, modifiers...)
	return aus
}
