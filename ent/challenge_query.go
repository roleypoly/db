// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/roleypoly/db/ent/challenge"
	"github.com/roleypoly/db/ent/predicate"
)

// ChallengeQuery is the builder for querying Challenge entities.
type ChallengeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Challenge
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (cq *ChallengeQuery) Where(ps ...predicate.Challenge) *ChallengeQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ChallengeQuery) Limit(limit int) *ChallengeQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ChallengeQuery) Offset(offset int) *ChallengeQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *ChallengeQuery) Order(o ...OrderFunc) *ChallengeQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Challenge entity in the query. Returns *NotFoundError when no challenge was found.
func (cq *ChallengeQuery) First(ctx context.Context) (*Challenge, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &NotFoundError{challenge.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChallengeQuery) FirstX(ctx context.Context) *Challenge {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first Challenge id in the query. Returns *NotFoundError when no id was found.
func (cq *ChallengeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{challenge.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *ChallengeQuery) FirstXID(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Challenge entity in the query, returns an error if not exactly one entity was returned.
func (cq *ChallengeQuery) Only(ctx context.Context) (*Challenge, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &NotFoundError{challenge.Label}
	default:
		return nil, &NotSingularError{challenge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChallengeQuery) OnlyX(ctx context.Context) *Challenge {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only Challenge id in the query, returns an error if not exactly one id was returned.
func (cq *ChallengeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = &NotSingularError{challenge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChallengeQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Challenges.
func (cq *ChallengeQuery) All(ctx context.Context) ([]*Challenge, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChallengeQuery) AllX(ctx context.Context) []*Challenge {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of Challenge ids.
func (cq *ChallengeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(challenge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChallengeQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChallengeQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChallengeQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChallengeQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChallengeQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChallengeQuery) Clone() *ChallengeQuery {
	return &ChallengeQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.Challenge{}, cq.predicates...),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Challenge.Query().
//		GroupBy(challenge.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *ChallengeQuery) GroupBy(field string, fields ...string) *ChallengeGroupBy {
	group := &ChallengeGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Challenge.Query().
//		Select(challenge.FieldCreateTime).
//		Scan(ctx, &v)
//
func (cq *ChallengeQuery) Select(field string, fields ...string) *ChallengeSelect {
	selector := &ChallengeSelect{config: cq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return selector
}

func (cq *ChallengeQuery) prepareQuery(ctx context.Context) error {
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChallengeQuery) sqlAll(ctx context.Context) ([]*Challenge, error) {
	var (
		nodes = []*Challenge{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Challenge{config: cq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *ChallengeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChallengeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *ChallengeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: challenge.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChallengeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(challenge.Table)
	selector := builder.Select(t1.Columns(challenge.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(challenge.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChallengeGroupBy is the builder for group-by Challenge entities.
type ChallengeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChallengeGroupBy) Aggregate(fns ...AggregateFunc) *ChallengeGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *ChallengeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *ChallengeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChallengeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *ChallengeGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cgb *ChallengeGroupBy) StringX(ctx context.Context) string {
	v, err := cgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChallengeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *ChallengeGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cgb *ChallengeGroupBy) IntX(ctx context.Context) int {
	v, err := cgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChallengeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *ChallengeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cgb *ChallengeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChallengeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *ChallengeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (cgb *ChallengeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cgb *ChallengeGroupBy) BoolX(ctx context.Context) bool {
	v, err := cgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *ChallengeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ChallengeGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// ChallengeSelect is the builder for select fields of Challenge entities.
type ChallengeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (cs *ChallengeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := cs.path(ctx)
	if err != nil {
		return err
	}
	cs.sql = query
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *ChallengeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChallengeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *ChallengeSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cs *ChallengeSelect) StringX(ctx context.Context) string {
	v, err := cs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChallengeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *ChallengeSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cs *ChallengeSelect) IntX(ctx context.Context) int {
	v, err := cs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChallengeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *ChallengeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cs *ChallengeSelect) Float64X(ctx context.Context) float64 {
	v, err := cs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChallengeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *ChallengeSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (cs *ChallengeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = fmt.Errorf("ent: ChallengeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cs *ChallengeSelect) BoolX(ctx context.Context) bool {
	v, err := cs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *ChallengeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *ChallengeSelect) sqlQuery() sql.Querier {
	selector := cs.sql
	selector.Select(selector.Columns(cs.fields...)...)
	return selector
}
