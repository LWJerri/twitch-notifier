// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/satont/twitch-notifier/ent/chat"
	"github.com/satont/twitch-notifier/ent/chatsettings"
	"github.com/satont/twitch-notifier/ent/predicate"
)

// ChatSettingsQuery is the builder for querying ChatSettings entities.
type ChatSettingsQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ChatSettings
	withChat   *ChatQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatSettingsQuery builder.
func (csq *ChatSettingsQuery) Where(ps ...predicate.ChatSettings) *ChatSettingsQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit the number of records to be returned by this query.
func (csq *ChatSettingsQuery) Limit(limit int) *ChatSettingsQuery {
	csq.ctx.Limit = &limit
	return csq
}

// Offset to start from.
func (csq *ChatSettingsQuery) Offset(offset int) *ChatSettingsQuery {
	csq.ctx.Offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *ChatSettingsQuery) Unique(unique bool) *ChatSettingsQuery {
	csq.ctx.Unique = &unique
	return csq
}

// Order specifies how the records should be ordered.
func (csq *ChatSettingsQuery) Order(o ...OrderFunc) *ChatSettingsQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryChat chains the current query on the "chat" edge.
func (csq *ChatSettingsQuery) QueryChat() *ChatQuery {
	query := (&ChatClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatsettings.Table, chatsettings.FieldID, selector),
			sqlgraph.To(chat.Table, chat.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, chatsettings.ChatTable, chatsettings.ChatColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChatSettings entity from the query.
// Returns a *NotFoundError when no ChatSettings was found.
func (csq *ChatSettingsQuery) First(ctx context.Context) (*ChatSettings, error) {
	nodes, err := csq.Limit(1).All(setContextOp(ctx, csq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatsettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *ChatSettingsQuery) FirstX(ctx context.Context) *ChatSettings {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatSettings ID from the query.
// Returns a *NotFoundError when no ChatSettings ID was found.
func (csq *ChatSettingsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = csq.Limit(1).IDs(setContextOp(ctx, csq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatsettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *ChatSettingsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChatSettings entity is found.
// Returns a *NotFoundError when no ChatSettings entities are found.
func (csq *ChatSettingsQuery) Only(ctx context.Context) (*ChatSettings, error) {
	nodes, err := csq.Limit(2).All(setContextOp(ctx, csq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatsettings.Label}
	default:
		return nil, &NotSingularError{chatsettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *ChatSettingsQuery) OnlyX(ctx context.Context) *ChatSettings {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatSettings ID in the query.
// Returns a *NotSingularError when more than one ChatSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *ChatSettingsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = csq.Limit(2).IDs(setContextOp(ctx, csq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatsettings.Label}
	default:
		err = &NotSingularError{chatsettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *ChatSettingsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatSettingsSlice.
func (csq *ChatSettingsQuery) All(ctx context.Context) ([]*ChatSettings, error) {
	ctx = setContextOp(ctx, csq.ctx, "All")
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChatSettings, *ChatSettingsQuery]()
	return withInterceptors[[]*ChatSettings](ctx, csq, qr, csq.inters)
}

// AllX is like All, but panics if an error occurs.
func (csq *ChatSettingsQuery) AllX(ctx context.Context) []*ChatSettings {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatSettings IDs.
func (csq *ChatSettingsQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if csq.ctx.Unique == nil && csq.path != nil {
		csq.Unique(true)
	}
	ctx = setContextOp(ctx, csq.ctx, "IDs")
	if err = csq.Select(chatsettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *ChatSettingsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *ChatSettingsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, csq.ctx, "Count")
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csq, querierCount[*ChatSettingsQuery](), csq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (csq *ChatSettingsQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *ChatSettingsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, csq.ctx, "Exist")
	switch _, err := csq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *ChatSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *ChatSettingsQuery) Clone() *ChatSettingsQuery {
	if csq == nil {
		return nil
	}
	return &ChatSettingsQuery{
		config:     csq.config,
		ctx:        csq.ctx.Clone(),
		order:      append([]OrderFunc{}, csq.order...),
		inters:     append([]Interceptor{}, csq.inters...),
		predicates: append([]predicate.ChatSettings{}, csq.predicates...),
		withChat:   csq.withChat.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithChat tells the query-builder to eager-load the nodes that are connected to
// the "chat" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *ChatSettingsQuery) WithChat(opts ...func(*ChatQuery)) *ChatSettingsQuery {
	query := (&ChatClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withChat = query
	return csq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GameChangeNotification bool `json:"game_change_notification,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatSettings.Query().
//		GroupBy(chatsettings.FieldGameChangeNotification).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (csq *ChatSettingsQuery) GroupBy(field string, fields ...string) *ChatSettingsGroupBy {
	csq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChatSettingsGroupBy{build: csq}
	grbuild.flds = &csq.ctx.Fields
	grbuild.label = chatsettings.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GameChangeNotification bool `json:"game_change_notification,omitempty"`
//	}
//
//	client.ChatSettings.Query().
//		Select(chatsettings.FieldGameChangeNotification).
//		Scan(ctx, &v)
func (csq *ChatSettingsQuery) Select(fields ...string) *ChatSettingsSelect {
	csq.ctx.Fields = append(csq.ctx.Fields, fields...)
	sbuild := &ChatSettingsSelect{ChatSettingsQuery: csq}
	sbuild.label = chatsettings.Label
	sbuild.flds, sbuild.scan = &csq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChatSettingsSelect configured with the given aggregations.
func (csq *ChatSettingsQuery) Aggregate(fns ...AggregateFunc) *ChatSettingsSelect {
	return csq.Select().Aggregate(fns...)
}

func (csq *ChatSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csq); err != nil {
				return err
			}
		}
	}
	for _, f := range csq.ctx.Fields {
		if !chatsettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *ChatSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChatSettings, error) {
	var (
		nodes       = []*ChatSettings{}
		_spec       = csq.querySpec()
		loadedTypes = [1]bool{
			csq.withChat != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChatSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChatSettings{config: csq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := csq.withChat; query != nil {
		if err := csq.loadChat(ctx, query, nodes, nil,
			func(n *ChatSettings, e *Chat) { n.Edges.Chat = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (csq *ChatSettingsQuery) loadChat(ctx context.Context, query *ChatQuery, nodes []*ChatSettings, init func(*ChatSettings), assign func(*ChatSettings, *Chat)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ChatSettings)
	for i := range nodes {
		fk := nodes[i].ChatID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(chat.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chat_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (csq *ChatSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Columns = csq.ctx.Fields
	if len(csq.ctx.Fields) > 0 {
		_spec.Unique = csq.ctx.Unique != nil && *csq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *ChatSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chatsettings.Table, chatsettings.Columns, sqlgraph.NewFieldSpec(chatsettings.FieldID, field.TypeUUID))
	_spec.From = csq.sql
	if unique := csq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csq.path != nil {
		_spec.Unique = true
	}
	if fields := csq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatsettings.FieldID)
		for i := range fields {
			if fields[i] != chatsettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *ChatSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(chatsettings.Table)
	columns := csq.ctx.Fields
	if len(columns) == 0 {
		columns = chatsettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.ctx.Unique != nil && *csq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatSettingsGroupBy is the group-by builder for ChatSettings entities.
type ChatSettingsGroupBy struct {
	selector
	build *ChatSettingsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *ChatSettingsGroupBy) Aggregate(fns ...AggregateFunc) *ChatSettingsGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the selector query and scans the result into the given value.
func (csgb *ChatSettingsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csgb.build.ctx, "GroupBy")
	if err := csgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatSettingsQuery, *ChatSettingsGroupBy](ctx, csgb.build, csgb, csgb.build.inters, v)
}

func (csgb *ChatSettingsGroupBy) sqlScan(ctx context.Context, root *ChatSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csgb.flds)+len(csgb.fns))
		for _, f := range *csgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChatSettingsSelect is the builder for selecting fields of ChatSettings entities.
type ChatSettingsSelect struct {
	*ChatSettingsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (css *ChatSettingsSelect) Aggregate(fns ...AggregateFunc) *ChatSettingsSelect {
	css.fns = append(css.fns, fns...)
	return css
}

// Scan applies the selector query and scans the result into the given value.
func (css *ChatSettingsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, css.ctx, "Select")
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatSettingsQuery, *ChatSettingsSelect](ctx, css.ChatSettingsQuery, css, css.inters, v)
}

func (css *ChatSettingsSelect) sqlScan(ctx context.Context, root *ChatSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(css.fns))
	for _, fn := range css.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*css.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}