// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/ignoreid"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// IgnoreIDQuery is the builder for querying IgnoreID entities.
type IgnoreIDQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.IgnoreID
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IgnoreIDQuery builder.
func (iiq *IgnoreIDQuery) Where(ps ...predicate.IgnoreID) *IgnoreIDQuery {
	iiq.predicates = append(iiq.predicates, ps...)
	return iiq
}

// Limit adds a limit step to the query.
func (iiq *IgnoreIDQuery) Limit(limit int) *IgnoreIDQuery {
	iiq.limit = &limit
	return iiq
}

// Offset adds an offset step to the query.
func (iiq *IgnoreIDQuery) Offset(offset int) *IgnoreIDQuery {
	iiq.offset = &offset
	return iiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iiq *IgnoreIDQuery) Unique(unique bool) *IgnoreIDQuery {
	iiq.unique = &unique
	return iiq
}

// Order adds an order step to the query.
func (iiq *IgnoreIDQuery) Order(o ...OrderFunc) *IgnoreIDQuery {
	iiq.order = append(iiq.order, o...)
	return iiq
}

// First returns the first IgnoreID entity from the query.
// Returns a *NotFoundError when no IgnoreID was found.
func (iiq *IgnoreIDQuery) First(ctx context.Context) (*IgnoreID, error) {
	nodes, err := iiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ignoreid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iiq *IgnoreIDQuery) FirstX(ctx context.Context) *IgnoreID {
	node, err := iiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IgnoreID ID from the query.
// Returns a *NotFoundError when no IgnoreID ID was found.
func (iiq *IgnoreIDQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ignoreid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iiq *IgnoreIDQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IgnoreID entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IgnoreID entity is found.
// Returns a *NotFoundError when no IgnoreID entities are found.
func (iiq *IgnoreIDQuery) Only(ctx context.Context) (*IgnoreID, error) {
	nodes, err := iiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ignoreid.Label}
	default:
		return nil, &NotSingularError{ignoreid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iiq *IgnoreIDQuery) OnlyX(ctx context.Context) *IgnoreID {
	node, err := iiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IgnoreID ID in the query.
// Returns a *NotSingularError when more than one IgnoreID ID is found.
// Returns a *NotFoundError when no entities are found.
func (iiq *IgnoreIDQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ignoreid.Label}
	default:
		err = &NotSingularError{ignoreid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iiq *IgnoreIDQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IgnoreIDs.
func (iiq *IgnoreIDQuery) All(ctx context.Context) ([]*IgnoreID, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iiq *IgnoreIDQuery) AllX(ctx context.Context) []*IgnoreID {
	nodes, err := iiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IgnoreID IDs.
func (iiq *IgnoreIDQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iiq.Select(ignoreid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iiq *IgnoreIDQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iiq *IgnoreIDQuery) Count(ctx context.Context) (int, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iiq *IgnoreIDQuery) CountX(ctx context.Context) int {
	count, err := iiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iiq *IgnoreIDQuery) Exist(ctx context.Context) (bool, error) {
	if err := iiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iiq *IgnoreIDQuery) ExistX(ctx context.Context) bool {
	exist, err := iiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IgnoreIDQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iiq *IgnoreIDQuery) Clone() *IgnoreIDQuery {
	if iiq == nil {
		return nil
	}
	return &IgnoreIDQuery{
		config:     iiq.config,
		limit:      iiq.limit,
		offset:     iiq.offset,
		order:      append([]OrderFunc{}, iiq.order...),
		predicates: append([]predicate.IgnoreID{}, iiq.predicates...),
		// clone intermediate query.
		sql:    iiq.sql.Clone(),
		path:   iiq.path,
		unique: iiq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IgnoreID.Query().
//		GroupBy(ignoreid.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iiq *IgnoreIDQuery) GroupBy(field string, fields ...string) *IgnoreIDGroupBy {
	grbuild := &IgnoreIDGroupBy{config: iiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iiq.sqlQuery(ctx), nil
	}
	grbuild.label = ignoreid.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.IgnoreID.Query().
//		Select(ignoreid.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (iiq *IgnoreIDQuery) Select(fields ...string) *IgnoreIDSelect {
	iiq.fields = append(iiq.fields, fields...)
	selbuild := &IgnoreIDSelect{IgnoreIDQuery: iiq}
	selbuild.label = ignoreid.Label
	selbuild.flds, selbuild.scan = &iiq.fields, selbuild.Scan
	return selbuild
}

func (iiq *IgnoreIDQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iiq.fields {
		if !ignoreid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iiq.path != nil {
		prev, err := iiq.path(ctx)
		if err != nil {
			return err
		}
		iiq.sql = prev
	}
	if ignoreid.Policy == nil {
		return errors.New("ent: uninitialized ignoreid.Policy (forgotten import ent/runtime?)")
	}
	if err := ignoreid.Policy.EvalQuery(ctx, iiq); err != nil {
		return err
	}
	return nil
}

func (iiq *IgnoreIDQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IgnoreID, error) {
	var (
		nodes = []*IgnoreID{}
		_spec = iiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*IgnoreID).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &IgnoreID{config: iiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(iiq.modifiers) > 0 {
		_spec.Modifiers = iiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iiq *IgnoreIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iiq.querySpec()
	if len(iiq.modifiers) > 0 {
		_spec.Modifiers = iiq.modifiers
	}
	_spec.Node.Columns = iiq.fields
	if len(iiq.fields) > 0 {
		_spec.Unique = iiq.unique != nil && *iiq.unique
	}
	return sqlgraph.CountNodes(ctx, iiq.driver, _spec)
}

func (iiq *IgnoreIDQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iiq *IgnoreIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ignoreid.Table,
			Columns: ignoreid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ignoreid.FieldID,
			},
		},
		From:   iiq.sql,
		Unique: true,
	}
	if unique := iiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ignoreid.FieldID)
		for i := range fields {
			if fields[i] != ignoreid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iiq *IgnoreIDQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iiq.driver.Dialect())
	t1 := builder.Table(ignoreid.Table)
	columns := iiq.fields
	if len(columns) == 0 {
		columns = ignoreid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iiq.sql != nil {
		selector = iiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iiq.unique != nil && *iiq.unique {
		selector.Distinct()
	}
	for _, m := range iiq.modifiers {
		m(selector)
	}
	for _, p := range iiq.predicates {
		p(selector)
	}
	for _, p := range iiq.order {
		p(selector)
	}
	if offset := iiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (iiq *IgnoreIDQuery) ForUpdate(opts ...sql.LockOption) *IgnoreIDQuery {
	if iiq.driver.Dialect() == dialect.Postgres {
		iiq.Unique(false)
	}
	iiq.modifiers = append(iiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return iiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (iiq *IgnoreIDQuery) ForShare(opts ...sql.LockOption) *IgnoreIDQuery {
	if iiq.driver.Dialect() == dialect.Postgres {
		iiq.Unique(false)
	}
	iiq.modifiers = append(iiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return iiq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (iiq *IgnoreIDQuery) Modify(modifiers ...func(s *sql.Selector)) *IgnoreIDSelect {
	iiq.modifiers = append(iiq.modifiers, modifiers...)
	return iiq.Select()
}

// IgnoreIDGroupBy is the group-by builder for IgnoreID entities.
type IgnoreIDGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iigb *IgnoreIDGroupBy) Aggregate(fns ...AggregateFunc) *IgnoreIDGroupBy {
	iigb.fns = append(iigb.fns, fns...)
	return iigb
}

// Scan applies the group-by query and scans the result into the given value.
func (iigb *IgnoreIDGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := iigb.path(ctx)
	if err != nil {
		return err
	}
	iigb.sql = query
	return iigb.sqlScan(ctx, v)
}

func (iigb *IgnoreIDGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range iigb.fields {
		if !ignoreid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := iigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (iigb *IgnoreIDGroupBy) sqlQuery() *sql.Selector {
	selector := iigb.sql.Select()
	aggregation := make([]string, 0, len(iigb.fns))
	for _, fn := range iigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(iigb.fields)+len(iigb.fns))
		for _, f := range iigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(iigb.fields...)...)
}

// IgnoreIDSelect is the builder for selecting fields of IgnoreID entities.
type IgnoreIDSelect struct {
	*IgnoreIDQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (iis *IgnoreIDSelect) Scan(ctx context.Context, v interface{}) error {
	if err := iis.prepareQuery(ctx); err != nil {
		return err
	}
	iis.sql = iis.IgnoreIDQuery.sqlQuery(ctx)
	return iis.sqlScan(ctx, v)
}

func (iis *IgnoreIDSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iis.sql.Query()
	if err := iis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (iis *IgnoreIDSelect) Modify(modifiers ...func(s *sql.Selector)) *IgnoreIDSelect {
	iis.modifiers = append(iis.modifiers, modifiers...)
	return iis
}
