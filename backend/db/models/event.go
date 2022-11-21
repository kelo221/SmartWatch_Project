// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Event is an object representing the database table.
type Event struct {
	ID             int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Eventname      string `boil:"eventname" json:"eventname" toml:"eventname" yaml:"eventname"`
	Eventtime      int    `boil:"eventtime" json:"eventtime" toml:"eventtime" yaml:"eventtime"`
	Userid         int    `boil:"userid" json:"userid" toml:"userid" yaml:"userid"`
	Issilent       bool   `boil:"issilent" json:"issilent" toml:"issilent" yaml:"issilent"`
	Snoozedisabled bool   `boil:"snoozedisabled" json:"snoozedisabled" toml:"snoozedisabled" yaml:"snoozedisabled"`

	R *eventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventColumns = struct {
	ID             string
	Eventname      string
	Eventtime      string
	Userid         string
	Issilent       string
	Snoozedisabled string
}{
	ID:             "id",
	Eventname:      "eventname",
	Eventtime:      "eventtime",
	Userid:         "userid",
	Issilent:       "issilent",
	Snoozedisabled: "snoozedisabled",
}

var EventTableColumns = struct {
	ID             string
	Eventname      string
	Eventtime      string
	Userid         string
	Issilent       string
	Snoozedisabled string
}{
	ID:             "event.id",
	Eventname:      "event.eventname",
	Eventtime:      "event.eventtime",
	Userid:         "event.userid",
	Issilent:       "event.issilent",
	Snoozedisabled: "event.snoozedisabled",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var EventWhere = struct {
	ID             whereHelperint
	Eventname      whereHelperstring
	Eventtime      whereHelperint
	Userid         whereHelperint
	Issilent       whereHelperbool
	Snoozedisabled whereHelperbool
}{
	ID:             whereHelperint{field: "\"event\".\"id\""},
	Eventname:      whereHelperstring{field: "\"event\".\"eventname\""},
	Eventtime:      whereHelperint{field: "\"event\".\"eventtime\""},
	Userid:         whereHelperint{field: "\"event\".\"userid\""},
	Issilent:       whereHelperbool{field: "\"event\".\"issilent\""},
	Snoozedisabled: whereHelperbool{field: "\"event\".\"snoozedisabled\""},
}

// EventRels is where relationship names are stored.
var EventRels = struct {
	UseridUser string
}{
	UseridUser: "UseridUser",
}

// eventR is where relationships are stored.
type eventR struct {
	UseridUser *User `boil:"UseridUser" json:"UseridUser" toml:"UseridUser" yaml:"UseridUser"`
}

// NewStruct creates a new relationship struct
func (*eventR) NewStruct() *eventR {
	return &eventR{}
}

func (r *eventR) GetUseridUser() *User {
	if r == nil {
		return nil
	}
	return r.UseridUser
}

// eventL is where Load methods for each relationship are stored.
type eventL struct{}

var (
	eventAllColumns            = []string{"id", "eventname", "eventtime", "userid", "issilent", "snoozedisabled"}
	eventColumnsWithoutDefault = []string{"eventname", "eventtime", "userid"}
	eventColumnsWithDefault    = []string{"id", "issilent", "snoozedisabled"}
	eventPrimaryKeyColumns     = []string{"id"}
	eventGeneratedColumns      = []string{}
)

type (
	// EventSlice is an alias for a slice of pointers to Event.
	// This should almost always be used instead of []Event.
	EventSlice []*Event
	// EventHook is the signature for custom Event hook methods
	EventHook func(context.Context, boil.ContextExecutor, *Event) error

	eventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventType                 = reflect.TypeOf(&Event{})
	eventMapping              = queries.MakeStructMapping(eventType)
	eventPrimaryKeyMapping, _ = queries.BindMapping(eventType, eventMapping, eventPrimaryKeyColumns)
	eventInsertCacheMut       sync.RWMutex
	eventInsertCache          = make(map[string]insertCache)
	eventUpdateCacheMut       sync.RWMutex
	eventUpdateCache          = make(map[string]updateCache)
	eventUpsertCacheMut       sync.RWMutex
	eventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventAfterSelectHooks []EventHook

var eventBeforeInsertHooks []EventHook
var eventAfterInsertHooks []EventHook

var eventBeforeUpdateHooks []EventHook
var eventAfterUpdateHooks []EventHook

var eventBeforeDeleteHooks []EventHook
var eventAfterDeleteHooks []EventHook

var eventBeforeUpsertHooks []EventHook
var eventAfterUpsertHooks []EventHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Event) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Event) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Event) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Event) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Event) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Event) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Event) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Event) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Event) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventHook registers your hook function for all future operations.
func AddEventHook(hookPoint boil.HookPoint, eventHook EventHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventAfterSelectHooks = append(eventAfterSelectHooks, eventHook)
	case boil.BeforeInsertHook:
		eventBeforeInsertHooks = append(eventBeforeInsertHooks, eventHook)
	case boil.AfterInsertHook:
		eventAfterInsertHooks = append(eventAfterInsertHooks, eventHook)
	case boil.BeforeUpdateHook:
		eventBeforeUpdateHooks = append(eventBeforeUpdateHooks, eventHook)
	case boil.AfterUpdateHook:
		eventAfterUpdateHooks = append(eventAfterUpdateHooks, eventHook)
	case boil.BeforeDeleteHook:
		eventBeforeDeleteHooks = append(eventBeforeDeleteHooks, eventHook)
	case boil.AfterDeleteHook:
		eventAfterDeleteHooks = append(eventAfterDeleteHooks, eventHook)
	case boil.BeforeUpsertHook:
		eventBeforeUpsertHooks = append(eventBeforeUpsertHooks, eventHook)
	case boil.AfterUpsertHook:
		eventAfterUpsertHooks = append(eventAfterUpsertHooks, eventHook)
	}
}

// OneG returns a single event record from the query using the global executor.
func (q eventQuery) OneG(ctx context.Context) (*Event, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single event record from the query.
func (q eventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Event, error) {
	o := &Event{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for event")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Event records from the query using the global executor.
func (q eventQuery) AllG(ctx context.Context) (EventSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Event records from the query.
func (q eventQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventSlice, error) {
	var o []*Event

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Event slice")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Event records in the query using the global executor
func (q eventQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Event records in the query.
func (q eventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count event rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q eventQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q eventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if event exists")
	}

	return count > 0, nil
}

// UseridUser pointed to by the foreign key.
func (o *Event) UseridUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Userid),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUseridUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventL) LoadUseridUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEvent interface{}, mods queries.Applicator) error {
	var slice []*Event
	var object *Event

	if singular {
		var ok bool
		object, ok = maybeEvent.(*Event)
		if !ok {
			object = new(Event)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEvent))
			}
		}
	} else {
		s, ok := maybeEvent.(*[]*Event)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEvent))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventR{}
		}
		args = append(args, object.Userid)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventR{}
			}

			for _, a := range args {
				if a == obj.Userid {
					continue Outer
				}
			}

			args = append(args, obj.Userid)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.UseridUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UseridEvents = append(foreign.R.UseridEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Userid == foreign.ID {
				local.R.UseridUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UseridEvents = append(foreign.R.UseridEvents, local)
				break
			}
		}
	}

	return nil
}

// SetUseridUserG of the event to the related item.
// Sets o.R.UseridUser to related.
// Adds o to related.R.UseridEvents.
// Uses the global database handle.
func (o *Event) SetUseridUserG(ctx context.Context, insert bool, related *User) error {
	return o.SetUseridUser(ctx, boil.GetContextDB(), insert, related)
}

// SetUseridUser of the event to the related item.
// Sets o.R.UseridUser to related.
// Adds o to related.R.UseridEvents.
func (o *Event) SetUseridUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"event\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"userid"}),
		strmangle.WhereClause("\"", "\"", 2, eventPrimaryKeyColumns),
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

	o.Userid = related.ID
	if o.R == nil {
		o.R = &eventR{
			UseridUser: related,
		}
	} else {
		o.R.UseridUser = related
	}

	if related.R == nil {
		related.R = &userR{
			UseridEvents: EventSlice{o},
		}
	} else {
		related.R.UseridEvents = append(related.R.UseridEvents, o)
	}

	return nil
}

// Events retrieves all the records using an executor.
func Events(mods ...qm.QueryMod) eventQuery {
	mods = append(mods, qm.From("\"event\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"event\".*"})
	}

	return eventQuery{q}
}

// FindEventG retrieves a single record by ID.
func FindEventG(ctx context.Context, iD int, selectCols ...string) (*Event, error) {
	return FindEvent(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Event, error) {
	eventObj := &Event{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"event\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from event")
	}

	if err = eventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventObj, err
	}

	return eventObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Event) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Event) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no event provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventInsertCacheMut.RLock()
	cache, cached := eventInsertCache[key]
	eventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventType, eventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"event\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"event\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodels: unable to insert into event")
	}

	if !cached {
		eventInsertCacheMut.Lock()
		eventInsertCache[key] = cache
		eventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Event record using the global executor.
// See Update for more documentation.
func (o *Event) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Event.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Event) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventUpdateCacheMut.RLock()
	cache, cached := eventUpdateCache[key]
	eventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update event, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"event\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, append(wl, eventPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update event row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for event")
	}

	if !cached {
		eventUpdateCacheMut.Lock()
		eventUpdateCache[key] = cache
		eventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q eventQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q eventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for event")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EventSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"event\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all event")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Event) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Event) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no event provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

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

	eventUpsertCacheMut.RLock()
	cache, cached := eventUpsertCache[key]
	eventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert event, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(eventPrimaryKeyColumns))
			copy(conflict, eventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"event\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventType, eventMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert event")
	}

	if !cached {
		eventUpsertCacheMut.Lock()
		eventUpsertCache[key] = cache
		eventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Event record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Event) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Event record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Event) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Event provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventPrimaryKeyMapping)
	sql := "DELETE FROM \"event\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for event")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q eventQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q eventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no eventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for event")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EventSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for event")
	}

	if len(eventAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Event) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: no Event provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Event) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: empty EventSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"event\".* FROM \"event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in EventSlice")
	}

	*o = slice

	return nil
}

// EventExistsG checks if the Event row exists.
func EventExistsG(ctx context.Context, iD int) (bool, error) {
	return EventExists(ctx, boil.GetContextDB(), iD)
}

// EventExists checks if the Event row exists.
func EventExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"event\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if event exists")
	}

	return exists, nil
}
