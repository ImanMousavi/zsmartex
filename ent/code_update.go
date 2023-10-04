// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/zsmartex/zsmartex/ent/code"
	"github.com/zsmartex/zsmartex/ent/predicate"
)

// CodeUpdate is the builder for updating Code entities.
type CodeUpdate struct {
	config
	hooks    []Hook
	mutation *CodeMutation
}

// Where appends a list predicates to the CodeUpdate builder.
func (cu *CodeUpdate) Where(ps ...predicate.Code) *CodeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CodeUpdate) SetUserID(s string) *CodeUpdate {
	cu.mutation.SetUserID(s)
	return cu
}

// SetCode sets the "code" field.
func (cu *CodeUpdate) SetCode(s string) *CodeUpdate {
	cu.mutation.SetCode(s)
	return cu
}

// SetType sets the "type" field.
func (cu *CodeUpdate) SetType(c code.Type) *CodeUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CodeUpdate) SetNillableType(c *code.Type) *CodeUpdate {
	if c != nil {
		cu.SetType(*c)
	}
	return cu
}

// SetCategory sets the "category" field.
func (cu *CodeUpdate) SetCategory(c code.Category) *CodeUpdate {
	cu.mutation.SetCategory(c)
	return cu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (cu *CodeUpdate) SetNillableCategory(c *code.Category) *CodeUpdate {
	if c != nil {
		cu.SetCategory(*c)
	}
	return cu
}

// SetEmailIndex sets the "email_index" field.
func (cu *CodeUpdate) SetEmailIndex(i int64) *CodeUpdate {
	cu.mutation.ResetEmailIndex()
	cu.mutation.SetEmailIndex(i)
	return cu
}

// SetNillableEmailIndex sets the "email_index" field if the given value is not nil.
func (cu *CodeUpdate) SetNillableEmailIndex(i *int64) *CodeUpdate {
	if i != nil {
		cu.SetEmailIndex(*i)
	}
	return cu
}

// AddEmailIndex adds i to the "email_index" field.
func (cu *CodeUpdate) AddEmailIndex(i int64) *CodeUpdate {
	cu.mutation.AddEmailIndex(i)
	return cu
}

// ClearEmailIndex clears the value of the "email_index" field.
func (cu *CodeUpdate) ClearEmailIndex() *CodeUpdate {
	cu.mutation.ClearEmailIndex()
	return cu
}

// SetEmailEncrypted sets the "email_encrypted" field.
func (cu *CodeUpdate) SetEmailEncrypted(s string) *CodeUpdate {
	cu.mutation.SetEmailEncrypted(s)
	return cu
}

// SetNillableEmailEncrypted sets the "email_encrypted" field if the given value is not nil.
func (cu *CodeUpdate) SetNillableEmailEncrypted(s *string) *CodeUpdate {
	if s != nil {
		cu.SetEmailEncrypted(*s)
	}
	return cu
}

// ClearEmailEncrypted clears the value of the "email_encrypted" field.
func (cu *CodeUpdate) ClearEmailEncrypted() *CodeUpdate {
	cu.mutation.ClearEmailEncrypted()
	return cu
}

// SetPhoneIndex sets the "phone_index" field.
func (cu *CodeUpdate) SetPhoneIndex(i int64) *CodeUpdate {
	cu.mutation.ResetPhoneIndex()
	cu.mutation.SetPhoneIndex(i)
	return cu
}

// SetNillablePhoneIndex sets the "phone_index" field if the given value is not nil.
func (cu *CodeUpdate) SetNillablePhoneIndex(i *int64) *CodeUpdate {
	if i != nil {
		cu.SetPhoneIndex(*i)
	}
	return cu
}

// AddPhoneIndex adds i to the "phone_index" field.
func (cu *CodeUpdate) AddPhoneIndex(i int64) *CodeUpdate {
	cu.mutation.AddPhoneIndex(i)
	return cu
}

// ClearPhoneIndex clears the value of the "phone_index" field.
func (cu *CodeUpdate) ClearPhoneIndex() *CodeUpdate {
	cu.mutation.ClearPhoneIndex()
	return cu
}

// SetPhoneEncrypted sets the "phone_encrypted" field.
func (cu *CodeUpdate) SetPhoneEncrypted(s string) *CodeUpdate {
	cu.mutation.SetPhoneEncrypted(s)
	return cu
}

// SetNillablePhoneEncrypted sets the "phone_encrypted" field if the given value is not nil.
func (cu *CodeUpdate) SetNillablePhoneEncrypted(s *string) *CodeUpdate {
	if s != nil {
		cu.SetPhoneEncrypted(*s)
	}
	return cu
}

// ClearPhoneEncrypted clears the value of the "phone_encrypted" field.
func (cu *CodeUpdate) ClearPhoneEncrypted() *CodeUpdate {
	cu.mutation.ClearPhoneEncrypted()
	return cu
}

// SetData sets the "data" field.
func (cu *CodeUpdate) SetData(jm json.RawMessage) *CodeUpdate {
	cu.mutation.SetData(jm)
	return cu
}

// AppendData appends jm to the "data" field.
func (cu *CodeUpdate) AppendData(jm json.RawMessage) *CodeUpdate {
	cu.mutation.AppendData(jm)
	return cu
}

// ClearData clears the value of the "data" field.
func (cu *CodeUpdate) ClearData() *CodeUpdate {
	cu.mutation.ClearData()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CodeUpdate) SetUpdatedAt(t time.Time) *CodeUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// Mutation returns the CodeMutation object of the builder.
func (cu *CodeUpdate) Mutation() *CodeMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CodeUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CodeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CodeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CodeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CodeUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := code.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CodeUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := code.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Code.type": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Category(); ok {
		if err := code.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Code.category": %w`, err)}
		}
	}
	return nil
}

func (cu *CodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(code.Table, code.Columns, sqlgraph.NewFieldSpec(code.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(code.FieldUserID, field.TypeString, value)
	}
	if value, ok := cu.mutation.Code(); ok {
		_spec.SetField(code.FieldCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(code.FieldType, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Category(); ok {
		_spec.SetField(code.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.EmailIndex(); ok {
		_spec.SetField(code.FieldEmailIndex, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedEmailIndex(); ok {
		_spec.AddField(code.FieldEmailIndex, field.TypeInt64, value)
	}
	if cu.mutation.EmailIndexCleared() {
		_spec.ClearField(code.FieldEmailIndex, field.TypeInt64)
	}
	if value, ok := cu.mutation.EmailEncrypted(); ok {
		_spec.SetField(code.FieldEmailEncrypted, field.TypeString, value)
	}
	if cu.mutation.EmailEncryptedCleared() {
		_spec.ClearField(code.FieldEmailEncrypted, field.TypeString)
	}
	if value, ok := cu.mutation.PhoneIndex(); ok {
		_spec.SetField(code.FieldPhoneIndex, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedPhoneIndex(); ok {
		_spec.AddField(code.FieldPhoneIndex, field.TypeInt64, value)
	}
	if cu.mutation.PhoneIndexCleared() {
		_spec.ClearField(code.FieldPhoneIndex, field.TypeInt64)
	}
	if value, ok := cu.mutation.PhoneEncrypted(); ok {
		_spec.SetField(code.FieldPhoneEncrypted, field.TypeString, value)
	}
	if cu.mutation.PhoneEncryptedCleared() {
		_spec.ClearField(code.FieldPhoneEncrypted, field.TypeString)
	}
	if value, ok := cu.mutation.Data(); ok {
		_spec.SetField(code.FieldData, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, code.FieldData, value)
		})
	}
	if cu.mutation.DataCleared() {
		_spec.ClearField(code.FieldData, field.TypeJSON)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(code.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{code.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CodeUpdateOne is the builder for updating a single Code entity.
type CodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodeMutation
}

// SetUserID sets the "user_id" field.
func (cuo *CodeUpdateOne) SetUserID(s string) *CodeUpdateOne {
	cuo.mutation.SetUserID(s)
	return cuo
}

// SetCode sets the "code" field.
func (cuo *CodeUpdateOne) SetCode(s string) *CodeUpdateOne {
	cuo.mutation.SetCode(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CodeUpdateOne) SetType(c code.Type) *CodeUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillableType(c *code.Type) *CodeUpdateOne {
	if c != nil {
		cuo.SetType(*c)
	}
	return cuo
}

// SetCategory sets the "category" field.
func (cuo *CodeUpdateOne) SetCategory(c code.Category) *CodeUpdateOne {
	cuo.mutation.SetCategory(c)
	return cuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillableCategory(c *code.Category) *CodeUpdateOne {
	if c != nil {
		cuo.SetCategory(*c)
	}
	return cuo
}

// SetEmailIndex sets the "email_index" field.
func (cuo *CodeUpdateOne) SetEmailIndex(i int64) *CodeUpdateOne {
	cuo.mutation.ResetEmailIndex()
	cuo.mutation.SetEmailIndex(i)
	return cuo
}

// SetNillableEmailIndex sets the "email_index" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillableEmailIndex(i *int64) *CodeUpdateOne {
	if i != nil {
		cuo.SetEmailIndex(*i)
	}
	return cuo
}

// AddEmailIndex adds i to the "email_index" field.
func (cuo *CodeUpdateOne) AddEmailIndex(i int64) *CodeUpdateOne {
	cuo.mutation.AddEmailIndex(i)
	return cuo
}

// ClearEmailIndex clears the value of the "email_index" field.
func (cuo *CodeUpdateOne) ClearEmailIndex() *CodeUpdateOne {
	cuo.mutation.ClearEmailIndex()
	return cuo
}

// SetEmailEncrypted sets the "email_encrypted" field.
func (cuo *CodeUpdateOne) SetEmailEncrypted(s string) *CodeUpdateOne {
	cuo.mutation.SetEmailEncrypted(s)
	return cuo
}

// SetNillableEmailEncrypted sets the "email_encrypted" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillableEmailEncrypted(s *string) *CodeUpdateOne {
	if s != nil {
		cuo.SetEmailEncrypted(*s)
	}
	return cuo
}

// ClearEmailEncrypted clears the value of the "email_encrypted" field.
func (cuo *CodeUpdateOne) ClearEmailEncrypted() *CodeUpdateOne {
	cuo.mutation.ClearEmailEncrypted()
	return cuo
}

// SetPhoneIndex sets the "phone_index" field.
func (cuo *CodeUpdateOne) SetPhoneIndex(i int64) *CodeUpdateOne {
	cuo.mutation.ResetPhoneIndex()
	cuo.mutation.SetPhoneIndex(i)
	return cuo
}

// SetNillablePhoneIndex sets the "phone_index" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillablePhoneIndex(i *int64) *CodeUpdateOne {
	if i != nil {
		cuo.SetPhoneIndex(*i)
	}
	return cuo
}

// AddPhoneIndex adds i to the "phone_index" field.
func (cuo *CodeUpdateOne) AddPhoneIndex(i int64) *CodeUpdateOne {
	cuo.mutation.AddPhoneIndex(i)
	return cuo
}

// ClearPhoneIndex clears the value of the "phone_index" field.
func (cuo *CodeUpdateOne) ClearPhoneIndex() *CodeUpdateOne {
	cuo.mutation.ClearPhoneIndex()
	return cuo
}

// SetPhoneEncrypted sets the "phone_encrypted" field.
func (cuo *CodeUpdateOne) SetPhoneEncrypted(s string) *CodeUpdateOne {
	cuo.mutation.SetPhoneEncrypted(s)
	return cuo
}

// SetNillablePhoneEncrypted sets the "phone_encrypted" field if the given value is not nil.
func (cuo *CodeUpdateOne) SetNillablePhoneEncrypted(s *string) *CodeUpdateOne {
	if s != nil {
		cuo.SetPhoneEncrypted(*s)
	}
	return cuo
}

// ClearPhoneEncrypted clears the value of the "phone_encrypted" field.
func (cuo *CodeUpdateOne) ClearPhoneEncrypted() *CodeUpdateOne {
	cuo.mutation.ClearPhoneEncrypted()
	return cuo
}

// SetData sets the "data" field.
func (cuo *CodeUpdateOne) SetData(jm json.RawMessage) *CodeUpdateOne {
	cuo.mutation.SetData(jm)
	return cuo
}

// AppendData appends jm to the "data" field.
func (cuo *CodeUpdateOne) AppendData(jm json.RawMessage) *CodeUpdateOne {
	cuo.mutation.AppendData(jm)
	return cuo
}

// ClearData clears the value of the "data" field.
func (cuo *CodeUpdateOne) ClearData() *CodeUpdateOne {
	cuo.mutation.ClearData()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CodeUpdateOne) SetUpdatedAt(t time.Time) *CodeUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// Mutation returns the CodeMutation object of the builder.
func (cuo *CodeUpdateOne) Mutation() *CodeMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CodeUpdate builder.
func (cuo *CodeUpdateOne) Where(ps ...predicate.Code) *CodeUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CodeUpdateOne) Select(field string, fields ...string) *CodeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Code entity.
func (cuo *CodeUpdateOne) Save(ctx context.Context) (*Code, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CodeUpdateOne) SaveX(ctx context.Context) *Code {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CodeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CodeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CodeUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := code.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CodeUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := code.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Code.type": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Category(); ok {
		if err := code.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Code.category": %w`, err)}
		}
	}
	return nil
}

func (cuo *CodeUpdateOne) sqlSave(ctx context.Context) (_node *Code, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(code.Table, code.Columns, sqlgraph.NewFieldSpec(code.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Code.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, code.FieldID)
		for _, f := range fields {
			if !code.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != code.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(code.FieldUserID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Code(); ok {
		_spec.SetField(code.FieldCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(code.FieldType, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Category(); ok {
		_spec.SetField(code.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.EmailIndex(); ok {
		_spec.SetField(code.FieldEmailIndex, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedEmailIndex(); ok {
		_spec.AddField(code.FieldEmailIndex, field.TypeInt64, value)
	}
	if cuo.mutation.EmailIndexCleared() {
		_spec.ClearField(code.FieldEmailIndex, field.TypeInt64)
	}
	if value, ok := cuo.mutation.EmailEncrypted(); ok {
		_spec.SetField(code.FieldEmailEncrypted, field.TypeString, value)
	}
	if cuo.mutation.EmailEncryptedCleared() {
		_spec.ClearField(code.FieldEmailEncrypted, field.TypeString)
	}
	if value, ok := cuo.mutation.PhoneIndex(); ok {
		_spec.SetField(code.FieldPhoneIndex, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedPhoneIndex(); ok {
		_spec.AddField(code.FieldPhoneIndex, field.TypeInt64, value)
	}
	if cuo.mutation.PhoneIndexCleared() {
		_spec.ClearField(code.FieldPhoneIndex, field.TypeInt64)
	}
	if value, ok := cuo.mutation.PhoneEncrypted(); ok {
		_spec.SetField(code.FieldPhoneEncrypted, field.TypeString, value)
	}
	if cuo.mutation.PhoneEncryptedCleared() {
		_spec.ClearField(code.FieldPhoneEncrypted, field.TypeString)
	}
	if value, ok := cuo.mutation.Data(); ok {
		_spec.SetField(code.FieldData, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, code.FieldData, value)
		})
	}
	if cuo.mutation.DataCleared() {
		_spec.ClearField(code.FieldData, field.TypeJSON)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(code.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Code{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{code.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}