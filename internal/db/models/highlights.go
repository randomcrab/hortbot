// Code generated by SQLBoiler v4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Highlight is an object representing the database table.
type Highlight struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ChannelID     int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	HighlightedAt time.Time `boil:"highlighted_at" json:"highlighted_at" toml:"highlighted_at" yaml:"highlighted_at"`
	StartedAt     null.Time `boil:"started_at" json:"started_at,omitempty" toml:"started_at" yaml:"started_at,omitempty"`
	Status        string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	Game          string    `boil:"game" json:"game" toml:"game" yaml:"game"`

	R *highlightR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L highlightL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HighlightColumns = struct {
	ID            string
	CreatedAt     string
	ChannelID     string
	HighlightedAt string
	StartedAt     string
	Status        string
	Game          string
}{
	ID:            "id",
	CreatedAt:     "created_at",
	ChannelID:     "channel_id",
	HighlightedAt: "highlighted_at",
	StartedAt:     "started_at",
	Status:        "status",
	Game:          "game",
}

// Generated where

var HighlightWhere = struct {
	ID            whereHelperint64
	CreatedAt     whereHelpertime_Time
	ChannelID     whereHelperint64
	HighlightedAt whereHelpertime_Time
	StartedAt     whereHelpernull_Time
	Status        whereHelperstring
	Game          whereHelperstring
}{
	ID:            whereHelperint64{field: "\"highlights\".\"id\""},
	CreatedAt:     whereHelpertime_Time{field: "\"highlights\".\"created_at\""},
	ChannelID:     whereHelperint64{field: "\"highlights\".\"channel_id\""},
	HighlightedAt: whereHelpertime_Time{field: "\"highlights\".\"highlighted_at\""},
	StartedAt:     whereHelpernull_Time{field: "\"highlights\".\"started_at\""},
	Status:        whereHelperstring{field: "\"highlights\".\"status\""},
	Game:          whereHelperstring{field: "\"highlights\".\"game\""},
}

// HighlightRels is where relationship names are stored.
var HighlightRels = struct {
	Channel string
}{
	Channel: "Channel",
}

// highlightR is where relationships are stored.
type highlightR struct {
	Channel *Channel
}

// NewStruct creates a new relationship struct
func (*highlightR) NewStruct() *highlightR {
	return &highlightR{}
}

// highlightL is where Load methods for each relationship are stored.
type highlightL struct{}

var (
	highlightAllColumns            = []string{"id", "created_at", "channel_id", "highlighted_at", "started_at", "status", "game"}
	highlightColumnsWithoutDefault = []string{"channel_id", "highlighted_at", "started_at", "status", "game"}
	highlightColumnsWithDefault    = []string{"id", "created_at"}
	highlightPrimaryKeyColumns     = []string{"id"}
)

type (
	// HighlightSlice is an alias for a slice of pointers to Highlight.
	// This should generally be used opposed to []Highlight.
	HighlightSlice []*Highlight

	highlightQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	highlightType                 = reflect.TypeOf(&Highlight{})
	highlightMapping              = queries.MakeStructMapping(highlightType)
	highlightPrimaryKeyMapping, _ = queries.BindMapping(highlightType, highlightMapping, highlightPrimaryKeyColumns)
	highlightInsertCacheMut       sync.RWMutex
	highlightInsertCache          = make(map[string]insertCache)
	highlightUpdateCacheMut       sync.RWMutex
	highlightUpdateCache          = make(map[string]updateCache)
	highlightUpsertCacheMut       sync.RWMutex
	highlightUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single highlight record from the query.
func (q highlightQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Highlight, error) {
	o := &Highlight{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for highlights")
	}

	return o, nil
}

// All returns all Highlight records from the query.
func (q highlightQuery) All(ctx context.Context, exec boil.ContextExecutor) (HighlightSlice, error) {
	var o []*Highlight

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Highlight slice")
	}

	return o, nil
}

// Count returns the count of all Highlight records in the query.
func (q highlightQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count highlights rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q highlightQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if highlights exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *Highlight) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (highlightL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHighlight interface{}, mods queries.Applicator) error {
	var slice []*Highlight
	var object *Highlight

	if singular {
		object = maybeHighlight.(*Highlight)
	} else {
		slice = *maybeHighlight.(*[]*Highlight)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &highlightR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &highlightR{}
			}

			for _, a := range args {
				if a == obj.ChannelID {
					continue Outer
				}
			}

			args = append(args, obj.ChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`channels`),
		qm.WhereIn(`channels.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for channels")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Channel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.Highlights = append(foreign.R.Highlights, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.Highlights = append(foreign.R.Highlights, local)
				break
			}
		}
	}

	return nil
}

// SetChannel of the highlight to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.Highlights.
func (o *Highlight) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"highlights\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, highlightPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID
	if o.R == nil {
		o.R = &highlightR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			Highlights: HighlightSlice{o},
		}
	} else {
		related.R.Highlights = append(related.R.Highlights, o)
	}

	return nil
}

// Highlights retrieves all the records using an executor.
func Highlights(mods ...qm.QueryMod) highlightQuery {
	mods = append(mods, qm.From("\"highlights\""))
	return highlightQuery{NewQuery(mods...)}
}

// FindHighlight retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHighlight(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Highlight, error) {
	highlightObj := &Highlight{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"highlights\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, highlightObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from highlights")
	}

	return highlightObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Highlight) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no highlights provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(highlightColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	highlightInsertCacheMut.RLock()
	cache, cached := highlightInsertCache[key]
	highlightInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			highlightAllColumns,
			highlightColumnsWithDefault,
			highlightColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(highlightType, highlightMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"highlights\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"highlights\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into highlights")
	}

	if !cached {
		highlightInsertCacheMut.Lock()
		highlightInsertCache[key] = cache
		highlightInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Highlight.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Highlight) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	highlightUpdateCacheMut.RLock()
	cache, cached := highlightUpdateCache[key]
	highlightUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			highlightAllColumns,
			highlightPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update highlights, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"highlights\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, highlightPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, append(wl, highlightPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update highlights row")
	}

	if !cached {
		highlightUpdateCacheMut.Lock()
		highlightUpdateCache[key] = cache
		highlightUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q highlightQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for highlights")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HighlightSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"highlights\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, highlightPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in highlight slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Highlight) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no highlights provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(highlightColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	highlightUpsertCacheMut.RLock()
	cache, cached := highlightUpsertCache[key]
	highlightUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			highlightAllColumns,
			highlightColumnsWithDefault,
			highlightColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			highlightAllColumns,
			highlightPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert highlights, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(highlightPrimaryKeyColumns))
			copy(conflict, highlightPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"highlights\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(highlightType, highlightMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert highlights")
	}

	if !cached {
		highlightUpsertCacheMut.Lock()
		highlightUpsertCache[key] = cache
		highlightUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Highlight record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Highlight) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Highlight provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), highlightPrimaryKeyMapping)
	sql := "DELETE FROM \"highlights\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from highlights")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q highlightQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no highlightQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from highlights")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HighlightSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"highlights\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, highlightPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from highlight slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Highlight) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHighlight(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HighlightSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HighlightSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"highlights\".* FROM \"highlights\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, highlightPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HighlightSlice")
	}

	*o = slice

	return nil
}

// HighlightExists checks if the Highlight row exists.
func HighlightExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"highlights\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if highlights exists")
	}

	return exists, nil
}
