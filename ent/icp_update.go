// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshino-s/soar-helper/ent/icp"
	"github.com/yoshino-s/soar-helper/ent/predicate"
)

// IcpUpdate is the builder for updating Icp entities.
type IcpUpdate struct {
	config
	hooks    []Hook
	mutation *IcpMutation
}

// Where appends a list predicates to the IcpUpdate builder.
func (iu *IcpUpdate) Where(ps ...predicate.Icp) *IcpUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetCity sets the "city" field.
func (iu *IcpUpdate) SetCity(s string) *IcpUpdate {
	iu.mutation.SetCity(s)
	return iu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableCity(s *string) *IcpUpdate {
	if s != nil {
		iu.SetCity(*s)
	}
	return iu
}

// ClearCity clears the value of the "city" field.
func (iu *IcpUpdate) ClearCity() *IcpUpdate {
	iu.mutation.ClearCity()
	return iu
}

// SetProvince sets the "province" field.
func (iu *IcpUpdate) SetProvince(s string) *IcpUpdate {
	iu.mutation.SetProvince(s)
	return iu
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableProvince(s *string) *IcpUpdate {
	if s != nil {
		iu.SetProvince(*s)
	}
	return iu
}

// ClearProvince clears the value of the "province" field.
func (iu *IcpUpdate) ClearProvince() *IcpUpdate {
	iu.mutation.ClearProvince()
	return iu
}

// SetCompany sets the "company" field.
func (iu *IcpUpdate) SetCompany(s string) *IcpUpdate {
	iu.mutation.SetCompany(s)
	return iu
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableCompany(s *string) *IcpUpdate {
	if s != nil {
		iu.SetCompany(*s)
	}
	return iu
}

// ClearCompany clears the value of the "company" field.
func (iu *IcpUpdate) ClearCompany() *IcpUpdate {
	iu.mutation.ClearCompany()
	return iu
}

// SetOwner sets the "owner" field.
func (iu *IcpUpdate) SetOwner(s string) *IcpUpdate {
	iu.mutation.SetOwner(s)
	return iu
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableOwner(s *string) *IcpUpdate {
	if s != nil {
		iu.SetOwner(*s)
	}
	return iu
}

// ClearOwner clears the value of the "owner" field.
func (iu *IcpUpdate) ClearOwner() *IcpUpdate {
	iu.mutation.ClearOwner()
	return iu
}

// SetType sets the "type" field.
func (iu *IcpUpdate) SetType(s string) *IcpUpdate {
	iu.mutation.SetType(s)
	return iu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableType(s *string) *IcpUpdate {
	if s != nil {
		iu.SetType(*s)
	}
	return iu
}

// SetHomepage sets the "homepage" field.
func (iu *IcpUpdate) SetHomepage(s string) *IcpUpdate {
	iu.mutation.SetHomepage(s)
	return iu
}

// SetNillableHomepage sets the "homepage" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableHomepage(s *string) *IcpUpdate {
	if s != nil {
		iu.SetHomepage(*s)
	}
	return iu
}

// ClearHomepage clears the value of the "homepage" field.
func (iu *IcpUpdate) ClearHomepage() *IcpUpdate {
	iu.mutation.ClearHomepage()
	return iu
}

// SetPermit sets the "permit" field.
func (iu *IcpUpdate) SetPermit(s string) *IcpUpdate {
	iu.mutation.SetPermit(s)
	return iu
}

// SetNillablePermit sets the "permit" field if the given value is not nil.
func (iu *IcpUpdate) SetNillablePermit(s *string) *IcpUpdate {
	if s != nil {
		iu.SetPermit(*s)
	}
	return iu
}

// ClearPermit clears the value of the "permit" field.
func (iu *IcpUpdate) ClearPermit() *IcpUpdate {
	iu.mutation.ClearPermit()
	return iu
}

// SetWebName sets the "webName" field.
func (iu *IcpUpdate) SetWebName(s string) *IcpUpdate {
	iu.mutation.SetWebName(s)
	return iu
}

// SetNillableWebName sets the "webName" field if the given value is not nil.
func (iu *IcpUpdate) SetNillableWebName(s *string) *IcpUpdate {
	if s != nil {
		iu.SetWebName(*s)
	}
	return iu
}

// ClearWebName clears the value of the "webName" field.
func (iu *IcpUpdate) ClearWebName() *IcpUpdate {
	iu.mutation.ClearWebName()
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IcpUpdate) SetUpdatedAt(t time.Time) *IcpUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// Mutation returns the IcpMutation object of the builder.
func (iu *IcpUpdate) Mutation() *IcpMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IcpUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IcpUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IcpUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IcpUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IcpUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := icp.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

func (iu *IcpUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(icp.Table, icp.Columns, sqlgraph.NewFieldSpec(icp.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.City(); ok {
		_spec.SetField(icp.FieldCity, field.TypeString, value)
	}
	if iu.mutation.CityCleared() {
		_spec.ClearField(icp.FieldCity, field.TypeString)
	}
	if value, ok := iu.mutation.Province(); ok {
		_spec.SetField(icp.FieldProvince, field.TypeString, value)
	}
	if iu.mutation.ProvinceCleared() {
		_spec.ClearField(icp.FieldProvince, field.TypeString)
	}
	if value, ok := iu.mutation.Company(); ok {
		_spec.SetField(icp.FieldCompany, field.TypeString, value)
	}
	if iu.mutation.CompanyCleared() {
		_spec.ClearField(icp.FieldCompany, field.TypeString)
	}
	if value, ok := iu.mutation.Owner(); ok {
		_spec.SetField(icp.FieldOwner, field.TypeString, value)
	}
	if iu.mutation.OwnerCleared() {
		_spec.ClearField(icp.FieldOwner, field.TypeString)
	}
	if value, ok := iu.mutation.GetType(); ok {
		_spec.SetField(icp.FieldType, field.TypeString, value)
	}
	if value, ok := iu.mutation.Homepage(); ok {
		_spec.SetField(icp.FieldHomepage, field.TypeString, value)
	}
	if iu.mutation.HomepageCleared() {
		_spec.ClearField(icp.FieldHomepage, field.TypeString)
	}
	if value, ok := iu.mutation.Permit(); ok {
		_spec.SetField(icp.FieldPermit, field.TypeString, value)
	}
	if iu.mutation.PermitCleared() {
		_spec.ClearField(icp.FieldPermit, field.TypeString)
	}
	if value, ok := iu.mutation.WebName(); ok {
		_spec.SetField(icp.FieldWebName, field.TypeString, value)
	}
	if iu.mutation.WebNameCleared() {
		_spec.ClearField(icp.FieldWebName, field.TypeString)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(icp.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IcpUpdateOne is the builder for updating a single Icp entity.
type IcpUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IcpMutation
}

// SetCity sets the "city" field.
func (iuo *IcpUpdateOne) SetCity(s string) *IcpUpdateOne {
	iuo.mutation.SetCity(s)
	return iuo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableCity(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetCity(*s)
	}
	return iuo
}

// ClearCity clears the value of the "city" field.
func (iuo *IcpUpdateOne) ClearCity() *IcpUpdateOne {
	iuo.mutation.ClearCity()
	return iuo
}

// SetProvince sets the "province" field.
func (iuo *IcpUpdateOne) SetProvince(s string) *IcpUpdateOne {
	iuo.mutation.SetProvince(s)
	return iuo
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableProvince(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetProvince(*s)
	}
	return iuo
}

// ClearProvince clears the value of the "province" field.
func (iuo *IcpUpdateOne) ClearProvince() *IcpUpdateOne {
	iuo.mutation.ClearProvince()
	return iuo
}

// SetCompany sets the "company" field.
func (iuo *IcpUpdateOne) SetCompany(s string) *IcpUpdateOne {
	iuo.mutation.SetCompany(s)
	return iuo
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableCompany(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetCompany(*s)
	}
	return iuo
}

// ClearCompany clears the value of the "company" field.
func (iuo *IcpUpdateOne) ClearCompany() *IcpUpdateOne {
	iuo.mutation.ClearCompany()
	return iuo
}

// SetOwner sets the "owner" field.
func (iuo *IcpUpdateOne) SetOwner(s string) *IcpUpdateOne {
	iuo.mutation.SetOwner(s)
	return iuo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableOwner(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetOwner(*s)
	}
	return iuo
}

// ClearOwner clears the value of the "owner" field.
func (iuo *IcpUpdateOne) ClearOwner() *IcpUpdateOne {
	iuo.mutation.ClearOwner()
	return iuo
}

// SetType sets the "type" field.
func (iuo *IcpUpdateOne) SetType(s string) *IcpUpdateOne {
	iuo.mutation.SetType(s)
	return iuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableType(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetType(*s)
	}
	return iuo
}

// SetHomepage sets the "homepage" field.
func (iuo *IcpUpdateOne) SetHomepage(s string) *IcpUpdateOne {
	iuo.mutation.SetHomepage(s)
	return iuo
}

// SetNillableHomepage sets the "homepage" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableHomepage(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetHomepage(*s)
	}
	return iuo
}

// ClearHomepage clears the value of the "homepage" field.
func (iuo *IcpUpdateOne) ClearHomepage() *IcpUpdateOne {
	iuo.mutation.ClearHomepage()
	return iuo
}

// SetPermit sets the "permit" field.
func (iuo *IcpUpdateOne) SetPermit(s string) *IcpUpdateOne {
	iuo.mutation.SetPermit(s)
	return iuo
}

// SetNillablePermit sets the "permit" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillablePermit(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetPermit(*s)
	}
	return iuo
}

// ClearPermit clears the value of the "permit" field.
func (iuo *IcpUpdateOne) ClearPermit() *IcpUpdateOne {
	iuo.mutation.ClearPermit()
	return iuo
}

// SetWebName sets the "webName" field.
func (iuo *IcpUpdateOne) SetWebName(s string) *IcpUpdateOne {
	iuo.mutation.SetWebName(s)
	return iuo
}

// SetNillableWebName sets the "webName" field if the given value is not nil.
func (iuo *IcpUpdateOne) SetNillableWebName(s *string) *IcpUpdateOne {
	if s != nil {
		iuo.SetWebName(*s)
	}
	return iuo
}

// ClearWebName clears the value of the "webName" field.
func (iuo *IcpUpdateOne) ClearWebName() *IcpUpdateOne {
	iuo.mutation.ClearWebName()
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IcpUpdateOne) SetUpdatedAt(t time.Time) *IcpUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// Mutation returns the IcpMutation object of the builder.
func (iuo *IcpUpdateOne) Mutation() *IcpMutation {
	return iuo.mutation
}

// Where appends a list predicates to the IcpUpdate builder.
func (iuo *IcpUpdateOne) Where(ps ...predicate.Icp) *IcpUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IcpUpdateOne) Select(field string, fields ...string) *IcpUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Icp entity.
func (iuo *IcpUpdateOne) Save(ctx context.Context) (*Icp, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IcpUpdateOne) SaveX(ctx context.Context) *Icp {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IcpUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IcpUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IcpUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := icp.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

func (iuo *IcpUpdateOne) sqlSave(ctx context.Context) (_node *Icp, err error) {
	_spec := sqlgraph.NewUpdateSpec(icp.Table, icp.Columns, sqlgraph.NewFieldSpec(icp.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Icp.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, icp.FieldID)
		for _, f := range fields {
			if !icp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != icp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.City(); ok {
		_spec.SetField(icp.FieldCity, field.TypeString, value)
	}
	if iuo.mutation.CityCleared() {
		_spec.ClearField(icp.FieldCity, field.TypeString)
	}
	if value, ok := iuo.mutation.Province(); ok {
		_spec.SetField(icp.FieldProvince, field.TypeString, value)
	}
	if iuo.mutation.ProvinceCleared() {
		_spec.ClearField(icp.FieldProvince, field.TypeString)
	}
	if value, ok := iuo.mutation.Company(); ok {
		_spec.SetField(icp.FieldCompany, field.TypeString, value)
	}
	if iuo.mutation.CompanyCleared() {
		_spec.ClearField(icp.FieldCompany, field.TypeString)
	}
	if value, ok := iuo.mutation.Owner(); ok {
		_spec.SetField(icp.FieldOwner, field.TypeString, value)
	}
	if iuo.mutation.OwnerCleared() {
		_spec.ClearField(icp.FieldOwner, field.TypeString)
	}
	if value, ok := iuo.mutation.GetType(); ok {
		_spec.SetField(icp.FieldType, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Homepage(); ok {
		_spec.SetField(icp.FieldHomepage, field.TypeString, value)
	}
	if iuo.mutation.HomepageCleared() {
		_spec.ClearField(icp.FieldHomepage, field.TypeString)
	}
	if value, ok := iuo.mutation.Permit(); ok {
		_spec.SetField(icp.FieldPermit, field.TypeString, value)
	}
	if iuo.mutation.PermitCleared() {
		_spec.ClearField(icp.FieldPermit, field.TypeString)
	}
	if value, ok := iuo.mutation.WebName(); ok {
		_spec.SetField(icp.FieldWebName, field.TypeString, value)
	}
	if iuo.mutation.WebNameCleared() {
		_spec.ClearField(icp.FieldWebName, field.TypeString)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(icp.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Icp{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
