// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package model

import (
	"database/sql"
	"fmt"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewDevice returns a new instance of Device.
func NewDevice(name string, user *User) (record *Device) {
	return newDevice(name, user)
}

// GetID returns the primary key of the model.
func (r *Device) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Device) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "name":
		return &r.Name, nil
	case "user_id":
		return types.Nullable(kallax.VirtualColumn("user_id", r, new(kallax.ULID))), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Device: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Device) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "name":
		return r.Name, nil
	case "user_id":
		return r.Model.VirtualColumn(col), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Device: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Device) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "User":
		return new(User), nil

	}
	return nil, fmt.Errorf("kallax: model Device has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Device) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "User":
		val, ok := rel.(*User)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship User", rel)
		}
		if !val.GetID().IsEmpty() {
			r.User = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Device has no relationship %s", field)
}

// DeviceStore is the entity to access the records of the type Device
// in the database.
type DeviceStore struct {
	*kallax.Store
}

// NewDeviceStore creates a new instance of DeviceStore
// using a SQL database.
func NewDeviceStore(db *sql.DB) *DeviceStore {
	return &DeviceStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *DeviceStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *DeviceStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *DeviceStore) Debug() *DeviceStore {
	return &DeviceStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *DeviceStore) DebugWith(logger kallax.LoggerFunc) *DeviceStore {
	return &DeviceStore{s.Store.DebugWith(logger)}
}

func (s *DeviceStore) inverseRecords(record *Device) []kallax.RecordWithSchema {
	record.ClearVirtualColumns()
	var records []kallax.RecordWithSchema

	if record.User != nil {
		record.AddVirtualColumn("user_id", record.User.GetID())
		records = append(records, kallax.RecordWithSchema{
			Schema: Schema.User.BaseSchema,
			Record: record.User,
		})
	}

	return records
}

// Insert inserts a Device in the database. A non-persisted object is
// required for this operation.
func (s *DeviceStore) Insert(record *Device) error {

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {

			for _, r := range inverseRecords {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Device.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Device.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *DeviceStore) Update(record *Device, cols ...kallax.SchemaField) (updated int64, err error) {

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {

			for _, r := range inverseRecords {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Device.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Device.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *DeviceStore) Save(record *Device) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *DeviceStore) Delete(record *Device) error {

	return s.Store.Delete(Schema.Device.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *DeviceStore) Find(q *DeviceQuery) (*DeviceResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewDeviceResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *DeviceStore) MustFind(q *DeviceQuery) *DeviceResultSet {
	return NewDeviceResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *DeviceStore) Count(q *DeviceQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *DeviceStore) MustCount(q *DeviceQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *DeviceStore) FindOne(q *DeviceQuery) (*Device, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *DeviceStore) FindAll(q *DeviceQuery) ([]*Device, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *DeviceStore) MustFindOne(q *DeviceQuery) *Device {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Device with the data in the database and
// makes it writable.
func (s *DeviceStore) Reload(record *Device) error {
	return s.Store.Reload(Schema.Device.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *DeviceStore) Transaction(callback func(*DeviceStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&DeviceStore{store})
	})
}

// DeviceQuery is the object used to create queries for the Device
// entity.
type DeviceQuery struct {
	*kallax.BaseQuery
}

// NewDeviceQuery returns a new instance of DeviceQuery.
func NewDeviceQuery() *DeviceQuery {
	return &DeviceQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Device.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *DeviceQuery) Select(columns ...kallax.SchemaField) *DeviceQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *DeviceQuery) SelectNot(columns ...kallax.SchemaField) *DeviceQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *DeviceQuery) Copy() *DeviceQuery {
	return &DeviceQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *DeviceQuery) Order(cols ...kallax.ColumnOrder) *DeviceQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *DeviceQuery) BatchSize(size uint64) *DeviceQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *DeviceQuery) Limit(n uint64) *DeviceQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *DeviceQuery) Offset(n uint64) *DeviceQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *DeviceQuery) Where(cond kallax.Condition) *DeviceQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *DeviceQuery) WithUser() *DeviceQuery {
	q.AddRelation(Schema.User.BaseSchema, "User", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *DeviceQuery) FindByID(v ...kallax.ULID) *DeviceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Device.ID, values...))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *DeviceQuery) FindByName(v string) *DeviceQuery {
	return q.Where(kallax.Eq(Schema.Device.Name, v))
}

// FindByUser adds a new filter to the query that will require that
// the foreign key of User is equal to the passed value.
func (q *DeviceQuery) FindByUser(v kallax.ULID) *DeviceQuery {
	return q.Where(kallax.Eq(Schema.Device.UserFK, v))
}

// DeviceResultSet is the set of results returned by a query to the
// database.
type DeviceResultSet struct {
	ResultSet kallax.ResultSet
	last      *Device
	lastErr   error
}

// NewDeviceResultSet creates a new result set for rows of the type
// Device.
func NewDeviceResultSet(rs kallax.ResultSet) *DeviceResultSet {
	return &DeviceResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *DeviceResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Device.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Device)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Device")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *DeviceResultSet) Get() (*Device, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *DeviceResultSet) ForEach(fn func(*Device) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *DeviceResultSet) All() ([]*Device, error) {
	var result []*Device
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *DeviceResultSet) One() (*Device, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *DeviceResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *DeviceResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUser returns a new instance of User.
func NewUser() (record *User) {
	return newUser()
}

// GetID returns the primary key of the model.
func (r *User) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *User) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// Value returns the value of the given column.
func (r *User) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *User) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model User has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *User) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model User has no relationships")
}

// UserStore is the entity to access the records of the type User
// in the database.
type UserStore struct {
	*kallax.Store
}

// NewUserStore creates a new instance of UserStore
// using a SQL database.
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *UserStore) Debug() *UserStore {
	return &UserStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *UserStore) DebugWith(logger kallax.LoggerFunc) *UserStore {
	return &UserStore{s.Store.DebugWith(logger)}
}

// Insert inserts a User in the database. A non-persisted object is
// required for this operation.
func (s *UserStore) Insert(record *User) error {

	return s.Store.Insert(Schema.User.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserStore) Update(record *User, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.User.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserStore) Save(record *User) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *UserStore) Delete(record *User) error {

	return s.Store.Delete(Schema.User.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *UserStore) Find(q *UserQuery) (*UserResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserStore) MustFind(q *UserQuery) *UserResultSet {
	return NewUserResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserStore) Count(q *UserQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserStore) MustCount(q *UserQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserStore) FindOne(q *UserQuery) (*User, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *UserStore) FindAll(q *UserQuery) ([]*User, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserStore) MustFindOne(q *UserQuery) *User {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the User with the data in the database and
// makes it writable.
func (s *UserStore) Reload(record *User) error {
	return s.Store.Reload(Schema.User.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserStore) Transaction(callback func(*UserStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserStore{store})
	})
}

// UserQuery is the object used to create queries for the User
// entity.
type UserQuery struct {
	*kallax.BaseQuery
}

// NewUserQuery returns a new instance of UserQuery.
func NewUserQuery() *UserQuery {
	return &UserQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.User.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserQuery) Select(columns ...kallax.SchemaField) *UserQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserQuery) SelectNot(columns ...kallax.SchemaField) *UserQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserQuery) Copy() *UserQuery {
	return &UserQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserQuery) Order(cols ...kallax.ColumnOrder) *UserQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserQuery) BatchSize(size uint64) *UserQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserQuery) Limit(n uint64) *UserQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserQuery) Offset(n uint64) *UserQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserQuery) Where(cond kallax.Condition) *UserQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByID(v ...kallax.ULID) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.User.ID, values...))
}

// UserResultSet is the set of results returned by a query to the
// database.
type UserResultSet struct {
	ResultSet kallax.ResultSet
	last      *User
	lastErr   error
}

// NewUserResultSet creates a new result set for rows of the type
// User.
func NewUserResultSet(rs kallax.ResultSet) *UserResultSet {
	return &UserResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.User.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*User)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *User")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserResultSet) Get() (*User, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserResultSet) ForEach(fn func(*User) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *UserResultSet) All() ([]*User, error) {
	var result []*User
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *UserResultSet) One() (*User, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *UserResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Device *schemaDevice
	User   *schemaUser
}

type schemaDevice struct {
	*kallax.BaseSchema
	ID     kallax.SchemaField
	Name   kallax.SchemaField
	UserFK kallax.SchemaField
}

type schemaUser struct {
	*kallax.BaseSchema
	ID kallax.SchemaField
}

var Schema = &schema{
	Device: &schemaDevice{
		BaseSchema: kallax.NewBaseSchema(
			"devices",
			"__device",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"User": kallax.NewForeignKey("user_id", true),
			},
			func() kallax.Record {
				return new(Device)
			},
			false,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("user_id"),
		),
		ID:     kallax.NewSchemaField("id"),
		Name:   kallax.NewSchemaField("name"),
		UserFK: kallax.NewSchemaField("user_id"),
	},
	User: &schemaUser{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__user",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(User)
			},
			false,
			kallax.NewSchemaField("id"),
		),
		ID: kallax.NewSchemaField("id"),
	},
}
