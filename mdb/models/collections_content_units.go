// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package mdbmodels

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// CollectionsContentUnit is an object representing the database table.
type CollectionsContentUnit struct {
	CollectionID  int64  `boil:"collection_id" json:"collection_id" toml:"collection_id" yaml:"collection_id"`
	ContentUnitID int64  `boil:"content_unit_id" json:"content_unit_id" toml:"content_unit_id" yaml:"content_unit_id"`
	Name          string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *collectionsContentUnitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L collectionsContentUnitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// collectionsContentUnitR is where relationships are stored.
type collectionsContentUnitR struct {
	Collection  *Collection
	ContentUnit *ContentUnit
}

// collectionsContentUnitL is where Load methods for each relationship are stored.
type collectionsContentUnitL struct{}

var (
	collectionsContentUnitColumns               = []string{"collection_id", "content_unit_id", "name"}
	collectionsContentUnitColumnsWithoutDefault = []string{"collection_id", "content_unit_id", "name"}
	collectionsContentUnitColumnsWithDefault    = []string{}
	collectionsContentUnitPrimaryKeyColumns     = []string{"collection_id", "content_unit_id"}
)

type (
	// CollectionsContentUnitSlice is an alias for a slice of pointers to CollectionsContentUnit.
	// This should generally be used opposed to []CollectionsContentUnit.
	CollectionsContentUnitSlice []*CollectionsContentUnit

	collectionsContentUnitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	collectionsContentUnitType                 = reflect.TypeOf(&CollectionsContentUnit{})
	collectionsContentUnitMapping              = queries.MakeStructMapping(collectionsContentUnitType)
	collectionsContentUnitPrimaryKeyMapping, _ = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, collectionsContentUnitPrimaryKeyColumns)
	collectionsContentUnitInsertCacheMut       sync.RWMutex
	collectionsContentUnitInsertCache          = make(map[string]insertCache)
	collectionsContentUnitUpdateCacheMut       sync.RWMutex
	collectionsContentUnitUpdateCache          = make(map[string]updateCache)
	collectionsContentUnitUpsertCacheMut       sync.RWMutex
	collectionsContentUnitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single collectionsContentUnit record from the query, and panics on error.
func (q collectionsContentUnitQuery) OneP() *CollectionsContentUnit {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single collectionsContentUnit record from the query.
func (q collectionsContentUnitQuery) One() (*CollectionsContentUnit, error) {
	o := &CollectionsContentUnit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mdbmodels: failed to execute a one query for collections_content_units")
	}

	return o, nil
}

// AllP returns all CollectionsContentUnit records from the query, and panics on error.
func (q collectionsContentUnitQuery) AllP() CollectionsContentUnitSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CollectionsContentUnit records from the query.
func (q collectionsContentUnitQuery) All() (CollectionsContentUnitSlice, error) {
	var o []*CollectionsContentUnit

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "mdbmodels: failed to assign all query results to CollectionsContentUnit slice")
	}

	return o, nil
}

// CountP returns the count of all CollectionsContentUnit records in the query, and panics on error.
func (q collectionsContentUnitQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CollectionsContentUnit records in the query.
func (q collectionsContentUnitQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mdbmodels: failed to count collections_content_units rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q collectionsContentUnitQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q collectionsContentUnitQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mdbmodels: failed to check if collections_content_units exists")
	}

	return count > 0, nil
}

// CollectionG pointed to by the foreign key.
func (o *CollectionsContentUnit) CollectionG(mods ...qm.QueryMod) collectionQuery {
	return o.Collection(boil.GetDB(), mods...)
}

// Collection pointed to by the foreign key.
func (o *CollectionsContentUnit) Collection(exec boil.Executor, mods ...qm.QueryMod) collectionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CollectionID),
	}

	queryMods = append(queryMods, mods...)

	query := Collections(exec, queryMods...)
	queries.SetFrom(query.Query, "\"collections\"")

	return query
}

// ContentUnitG pointed to by the foreign key.
func (o *CollectionsContentUnit) ContentUnitG(mods ...qm.QueryMod) contentUnitQuery {
	return o.ContentUnit(boil.GetDB(), mods...)
}

// ContentUnit pointed to by the foreign key.
func (o *CollectionsContentUnit) ContentUnit(exec boil.Executor, mods ...qm.QueryMod) contentUnitQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ContentUnitID),
	}

	queryMods = append(queryMods, mods...)

	query := ContentUnits(exec, queryMods...)
	queries.SetFrom(query.Query, "\"content_units\"")

	return query
} // LoadCollection allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (collectionsContentUnitL) LoadCollection(e boil.Executor, singular bool, maybeCollectionsContentUnit interface{}) error {
	var slice []*CollectionsContentUnit
	var object *CollectionsContentUnit

	count := 1
	if singular {
		object = maybeCollectionsContentUnit.(*CollectionsContentUnit)
	} else {
		slice = *maybeCollectionsContentUnit.(*[]*CollectionsContentUnit)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &collectionsContentUnitR{}
		}
		args[0] = object.CollectionID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &collectionsContentUnitR{}
			}
			args[i] = obj.CollectionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"collections\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Collection")
	}
	defer results.Close()

	var resultSlice []*Collection
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Collection")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Collection = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CollectionID == foreign.ID {
				local.R.Collection = foreign
				break
			}
		}
	}

	return nil
}

// LoadContentUnit allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (collectionsContentUnitL) LoadContentUnit(e boil.Executor, singular bool, maybeCollectionsContentUnit interface{}) error {
	var slice []*CollectionsContentUnit
	var object *CollectionsContentUnit

	count := 1
	if singular {
		object = maybeCollectionsContentUnit.(*CollectionsContentUnit)
	} else {
		slice = *maybeCollectionsContentUnit.(*[]*CollectionsContentUnit)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &collectionsContentUnitR{}
		}
		args[0] = object.ContentUnitID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &collectionsContentUnitR{}
			}
			args[i] = obj.ContentUnitID
		}
	}

	query := fmt.Sprintf(
		"select * from \"content_units\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ContentUnit")
	}
	defer results.Close()

	var resultSlice []*ContentUnit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ContentUnit")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.ContentUnit = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ContentUnitID == foreign.ID {
				local.R.ContentUnit = foreign
				break
			}
		}
	}

	return nil
}

// SetCollectionG of the collections_content_unit to the related item.
// Sets o.R.Collection to related.
// Adds o to related.R.CollectionsContentUnits.
// Uses the global database handle.
func (o *CollectionsContentUnit) SetCollectionG(insert bool, related *Collection) error {
	return o.SetCollection(boil.GetDB(), insert, related)
}

// SetCollectionP of the collections_content_unit to the related item.
// Sets o.R.Collection to related.
// Adds o to related.R.CollectionsContentUnits.
// Panics on error.
func (o *CollectionsContentUnit) SetCollectionP(exec boil.Executor, insert bool, related *Collection) {
	if err := o.SetCollection(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCollectionGP of the collections_content_unit to the related item.
// Sets o.R.Collection to related.
// Adds o to related.R.CollectionsContentUnits.
// Uses the global database handle and panics on error.
func (o *CollectionsContentUnit) SetCollectionGP(insert bool, related *Collection) {
	if err := o.SetCollection(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCollection of the collections_content_unit to the related item.
// Sets o.R.Collection to related.
// Adds o to related.R.CollectionsContentUnits.
func (o *CollectionsContentUnit) SetCollection(exec boil.Executor, insert bool, related *Collection) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"collections_content_units\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"collection_id"}),
		strmangle.WhereClause("\"", "\"", 2, collectionsContentUnitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.CollectionID, o.ContentUnitID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CollectionID = related.ID

	if o.R == nil {
		o.R = &collectionsContentUnitR{
			Collection: related,
		}
	} else {
		o.R.Collection = related
	}

	if related.R == nil {
		related.R = &collectionR{
			CollectionsContentUnits: CollectionsContentUnitSlice{o},
		}
	} else {
		related.R.CollectionsContentUnits = append(related.R.CollectionsContentUnits, o)
	}

	return nil
}

// SetContentUnitG of the collections_content_unit to the related item.
// Sets o.R.ContentUnit to related.
// Adds o to related.R.CollectionsContentUnits.
// Uses the global database handle.
func (o *CollectionsContentUnit) SetContentUnitG(insert bool, related *ContentUnit) error {
	return o.SetContentUnit(boil.GetDB(), insert, related)
}

// SetContentUnitP of the collections_content_unit to the related item.
// Sets o.R.ContentUnit to related.
// Adds o to related.R.CollectionsContentUnits.
// Panics on error.
func (o *CollectionsContentUnit) SetContentUnitP(exec boil.Executor, insert bool, related *ContentUnit) {
	if err := o.SetContentUnit(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetContentUnitGP of the collections_content_unit to the related item.
// Sets o.R.ContentUnit to related.
// Adds o to related.R.CollectionsContentUnits.
// Uses the global database handle and panics on error.
func (o *CollectionsContentUnit) SetContentUnitGP(insert bool, related *ContentUnit) {
	if err := o.SetContentUnit(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetContentUnit of the collections_content_unit to the related item.
// Sets o.R.ContentUnit to related.
// Adds o to related.R.CollectionsContentUnits.
func (o *CollectionsContentUnit) SetContentUnit(exec boil.Executor, insert bool, related *ContentUnit) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"collections_content_units\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"content_unit_id"}),
		strmangle.WhereClause("\"", "\"", 2, collectionsContentUnitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.CollectionID, o.ContentUnitID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ContentUnitID = related.ID

	if o.R == nil {
		o.R = &collectionsContentUnitR{
			ContentUnit: related,
		}
	} else {
		o.R.ContentUnit = related
	}

	if related.R == nil {
		related.R = &contentUnitR{
			CollectionsContentUnits: CollectionsContentUnitSlice{o},
		}
	} else {
		related.R.CollectionsContentUnits = append(related.R.CollectionsContentUnits, o)
	}

	return nil
}

// CollectionsContentUnitsG retrieves all records.
func CollectionsContentUnitsG(mods ...qm.QueryMod) collectionsContentUnitQuery {
	return CollectionsContentUnits(boil.GetDB(), mods...)
}

// CollectionsContentUnits retrieves all the records using an executor.
func CollectionsContentUnits(exec boil.Executor, mods ...qm.QueryMod) collectionsContentUnitQuery {
	mods = append(mods, qm.From("\"collections_content_units\""))
	return collectionsContentUnitQuery{NewQuery(exec, mods...)}
}

// FindCollectionsContentUnitG retrieves a single record by ID.
func FindCollectionsContentUnitG(collectionID int64, contentUnitID int64, selectCols ...string) (*CollectionsContentUnit, error) {
	return FindCollectionsContentUnit(boil.GetDB(), collectionID, contentUnitID, selectCols...)
}

// FindCollectionsContentUnitGP retrieves a single record by ID, and panics on error.
func FindCollectionsContentUnitGP(collectionID int64, contentUnitID int64, selectCols ...string) *CollectionsContentUnit {
	retobj, err := FindCollectionsContentUnit(boil.GetDB(), collectionID, contentUnitID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCollectionsContentUnit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCollectionsContentUnit(exec boil.Executor, collectionID int64, contentUnitID int64, selectCols ...string) (*CollectionsContentUnit, error) {
	collectionsContentUnitObj := &CollectionsContentUnit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"collections_content_units\" where \"collection_id\"=$1 AND \"content_unit_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, collectionID, contentUnitID)

	err := q.Bind(collectionsContentUnitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mdbmodels: unable to select from collections_content_units")
	}

	return collectionsContentUnitObj, nil
}

// FindCollectionsContentUnitP retrieves a single record by ID with an executor, and panics on error.
func FindCollectionsContentUnitP(exec boil.Executor, collectionID int64, contentUnitID int64, selectCols ...string) *CollectionsContentUnit {
	retobj, err := FindCollectionsContentUnit(exec, collectionID, contentUnitID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CollectionsContentUnit) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CollectionsContentUnit) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CollectionsContentUnit) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CollectionsContentUnit) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("mdbmodels: no collections_content_units provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(collectionsContentUnitColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	collectionsContentUnitInsertCacheMut.RLock()
	cache, cached := collectionsContentUnitInsertCache[key]
	collectionsContentUnitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			collectionsContentUnitColumns,
			collectionsContentUnitColumnsWithDefault,
			collectionsContentUnitColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"collections_content_units\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"collections_content_units\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to insert into collections_content_units")
	}

	if !cached {
		collectionsContentUnitInsertCacheMut.Lock()
		collectionsContentUnitInsertCache[key] = cache
		collectionsContentUnitInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CollectionsContentUnit record. See Update for
// whitelist behavior description.
func (o *CollectionsContentUnit) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CollectionsContentUnit record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CollectionsContentUnit) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CollectionsContentUnit, and panics on error.
// See Update for whitelist behavior description.
func (o *CollectionsContentUnit) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CollectionsContentUnit.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CollectionsContentUnit) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	collectionsContentUnitUpdateCacheMut.RLock()
	cache, cached := collectionsContentUnitUpdateCache[key]
	collectionsContentUnitUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			collectionsContentUnitColumns,
			collectionsContentUnitPrimaryKeyColumns,
			whitelist,
		)

		if len(wl) == 0 {
			return errors.New("mdbmodels: unable to update collections_content_units, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"collections_content_units\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, collectionsContentUnitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, append(wl, collectionsContentUnitPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to update collections_content_units row")
	}

	if !cached {
		collectionsContentUnitUpdateCacheMut.Lock()
		collectionsContentUnitUpdateCache[key] = cache
		collectionsContentUnitUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q collectionsContentUnitQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q collectionsContentUnitQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to update all for collections_content_units")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CollectionsContentUnitSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CollectionsContentUnitSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CollectionsContentUnitSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CollectionsContentUnitSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("mdbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionsContentUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"collections_content_units\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, collectionsContentUnitPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to update all in collectionsContentUnit slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CollectionsContentUnit) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CollectionsContentUnit) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CollectionsContentUnit) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CollectionsContentUnit) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("mdbmodels: no collections_content_units provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(collectionsContentUnitColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	collectionsContentUnitUpsertCacheMut.RLock()
	cache, cached := collectionsContentUnitUpsertCache[key]
	collectionsContentUnitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			collectionsContentUnitColumns,
			collectionsContentUnitColumnsWithDefault,
			collectionsContentUnitColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			collectionsContentUnitColumns,
			collectionsContentUnitPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("mdbmodels: unable to upsert collections_content_units, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(collectionsContentUnitPrimaryKeyColumns))
			copy(conflict, collectionsContentUnitPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"collections_content_units\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(collectionsContentUnitType, collectionsContentUnitMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to upsert collections_content_units")
	}

	if !cached {
		collectionsContentUnitUpsertCacheMut.Lock()
		collectionsContentUnitUpsertCache[key] = cache
		collectionsContentUnitUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single CollectionsContentUnit record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CollectionsContentUnit) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CollectionsContentUnit record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CollectionsContentUnit) DeleteG() error {
	if o == nil {
		return errors.New("mdbmodels: no CollectionsContentUnit provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CollectionsContentUnit record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CollectionsContentUnit) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CollectionsContentUnit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CollectionsContentUnit) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("mdbmodels: no CollectionsContentUnit provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), collectionsContentUnitPrimaryKeyMapping)
	sql := "DELETE FROM \"collections_content_units\" WHERE \"collection_id\"=$1 AND \"content_unit_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to delete from collections_content_units")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q collectionsContentUnitQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q collectionsContentUnitQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("mdbmodels: no collectionsContentUnitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to delete all from collections_content_units")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CollectionsContentUnitSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CollectionsContentUnitSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("mdbmodels: no CollectionsContentUnit slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CollectionsContentUnitSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CollectionsContentUnitSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("mdbmodels: no CollectionsContentUnit slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionsContentUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"collections_content_units\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, collectionsContentUnitPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to delete all from collectionsContentUnit slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CollectionsContentUnit) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CollectionsContentUnit) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CollectionsContentUnit) ReloadG() error {
	if o == nil {
		return errors.New("mdbmodels: no CollectionsContentUnit provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CollectionsContentUnit) Reload(exec boil.Executor) error {
	ret, err := FindCollectionsContentUnit(exec, o.CollectionID, o.ContentUnitID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CollectionsContentUnitSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CollectionsContentUnitSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CollectionsContentUnitSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("mdbmodels: empty CollectionsContentUnitSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CollectionsContentUnitSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	collectionsContentUnits := CollectionsContentUnitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), collectionsContentUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"collections_content_units\".* FROM \"collections_content_units\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, collectionsContentUnitPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&collectionsContentUnits)
	if err != nil {
		return errors.Wrap(err, "mdbmodels: unable to reload all in CollectionsContentUnitSlice")
	}

	*o = collectionsContentUnits

	return nil
}

// CollectionsContentUnitExists checks if the CollectionsContentUnit row exists.
func CollectionsContentUnitExists(exec boil.Executor, collectionID int64, contentUnitID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"collections_content_units\" where \"collection_id\"=$1 AND \"content_unit_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, collectionID, contentUnitID)
	}

	row := exec.QueryRow(sql, collectionID, contentUnitID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mdbmodels: unable to check if collections_content_units exists")
	}

	return exists, nil
}

// CollectionsContentUnitExistsG checks if the CollectionsContentUnit row exists.
func CollectionsContentUnitExistsG(collectionID int64, contentUnitID int64) (bool, error) {
	return CollectionsContentUnitExists(boil.GetDB(), collectionID, contentUnitID)
}

// CollectionsContentUnitExistsGP checks if the CollectionsContentUnit row exists. Panics on error.
func CollectionsContentUnitExistsGP(collectionID int64, contentUnitID int64) bool {
	e, err := CollectionsContentUnitExists(boil.GetDB(), collectionID, contentUnitID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CollectionsContentUnitExistsP checks if the CollectionsContentUnit row exists. Panics on error.
func CollectionsContentUnitExistsP(exec boil.Executor, collectionID int64, contentUnitID int64) bool {
	e, err := CollectionsContentUnitExists(exec, collectionID, contentUnitID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}