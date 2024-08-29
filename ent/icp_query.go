// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/ent/icp"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/ent/predicate"
)

// IcpQuery is the builder for querying Icp entities.
type IcpQuery struct {
	config
	ctx        *QueryContext
	order      []icp.OrderOption
	inters     []Interceptor
	predicates []predicate.Icp
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IcpQuery builder.
func (iq *IcpQuery) Where(ps ...predicate.Icp) *IcpQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IcpQuery) Limit(limit int) *IcpQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IcpQuery) Offset(offset int) *IcpQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IcpQuery) Unique(unique bool) *IcpQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IcpQuery) Order(o ...icp.OrderOption) *IcpQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Icp entity from the query.
// Returns a *NotFoundError when no Icp was found.
func (iq *IcpQuery) First(ctx context.Context) (*Icp, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{icp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IcpQuery) FirstX(ctx context.Context) *Icp {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Icp ID from the query.
// Returns a *NotFoundError when no Icp ID was found.
func (iq *IcpQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{icp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IcpQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Icp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Icp entity is found.
// Returns a *NotFoundError when no Icp entities are found.
func (iq *IcpQuery) Only(ctx context.Context) (*Icp, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{icp.Label}
	default:
		return nil, &NotSingularError{icp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IcpQuery) OnlyX(ctx context.Context) *Icp {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Icp ID in the query.
// Returns a *NotSingularError when more than one Icp ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IcpQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{icp.Label}
	default:
		err = &NotSingularError{icp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IcpQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Icps.
func (iq *IcpQuery) All(ctx context.Context) ([]*Icp, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Icp, *IcpQuery]()
	return withInterceptors[[]*Icp](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IcpQuery) AllX(ctx context.Context) []*Icp {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Icp IDs.
func (iq *IcpQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(icp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IcpQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IcpQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IcpQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IcpQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IcpQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IcpQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IcpQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IcpQuery) Clone() *IcpQuery {
	if iq == nil {
		return nil
	}
	return &IcpQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]icp.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Icp{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Host string `json:"host,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Icp.Query().
//		GroupBy(icp.FieldHost).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *IcpQuery) GroupBy(field string, fields ...string) *IcpGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IcpGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = icp.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Host string `json:"host,omitempty"`
//	}
//
//	client.Icp.Query().
//		Select(icp.FieldHost).
//		Scan(ctx, &v)
func (iq *IcpQuery) Select(fields ...string) *IcpSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IcpSelect{IcpQuery: iq}
	sbuild.label = icp.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IcpSelect configured with the given aggregations.
func (iq *IcpQuery) Aggregate(fns ...AggregateFunc) *IcpSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IcpQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !icp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IcpQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Icp, error) {
	var (
		nodes = []*Icp{}
		_spec = iq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Icp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Icp{config: iq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iq *IcpQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IcpQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(icp.Table, icp.Columns, sqlgraph.NewFieldSpec(icp.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, icp.FieldID)
		for i := range fields {
			if fields[i] != icp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IcpQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(icp.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = icp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IcpGroupBy is the group-by builder for Icp entities.
type IcpGroupBy struct {
	selector
	build *IcpQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IcpGroupBy) Aggregate(fns ...AggregateFunc) *IcpGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IcpGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IcpQuery, *IcpGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IcpGroupBy) sqlScan(ctx context.Context, root *IcpQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IcpSelect is the builder for selecting fields of Icp entities.
type IcpSelect struct {
	*IcpQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IcpSelect) Aggregate(fns ...AggregateFunc) *IcpSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IcpSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IcpQuery, *IcpSelect](ctx, is.IcpQuery, is, is.inters, v)
}

func (is *IcpSelect) sqlScan(ctx context.Context, root *IcpQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
