// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/target"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TargetUpdate is the builder for updating Target entities.
type TargetUpdate struct {
	config
	hooks    []Hook
	mutation *TargetMutation
}

// Where appends a list predicates to the TargetUpdate builder.
func (tu *TargetUpdate) Where(ps ...predicate.Target) *TargetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TargetUpdate) SetName(s string) *TargetUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetType sets the "type" field.
func (tu *TargetUpdate) SetType(t target.Type) *TargetUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetIntervalSeconds sets the "interval_seconds" field.
func (tu *TargetUpdate) SetIntervalSeconds(u uint32) *TargetUpdate {
	tu.mutation.ResetIntervalSeconds()
	tu.mutation.SetIntervalSeconds(u)
	return tu
}

// AddIntervalSeconds adds u to the "interval_seconds" field.
func (tu *TargetUpdate) AddIntervalSeconds(u int32) *TargetUpdate {
	tu.mutation.AddIntervalSeconds(u)
	return tu
}

// SetTimeoutSeconds sets the "timeout_seconds" field.
func (tu *TargetUpdate) SetTimeoutSeconds(u uint32) *TargetUpdate {
	tu.mutation.ResetTimeoutSeconds()
	tu.mutation.SetTimeoutSeconds(u)
	return tu
}

// AddTimeoutSeconds adds u to the "timeout_seconds" field.
func (tu *TargetUpdate) AddTimeoutSeconds(u int32) *TargetUpdate {
	tu.mutation.AddTimeoutSeconds(u)
	return tu
}

// SetURL sets the "url" field.
func (tu *TargetUpdate) SetURL(s string) *TargetUpdate {
	tu.mutation.SetURL(s)
	return tu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (tu *TargetUpdate) SetNillableURL(s *string) *TargetUpdate {
	if s != nil {
		tu.SetURL(*s)
	}
	return tu
}

// ClearURL clears the value of the "url" field.
func (tu *TargetUpdate) ClearURL() *TargetUpdate {
	tu.mutation.ClearURL()
	return tu
}

// SetHostname sets the "hostname" field.
func (tu *TargetUpdate) SetHostname(s string) *TargetUpdate {
	tu.mutation.SetHostname(s)
	return tu
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (tu *TargetUpdate) SetNillableHostname(s *string) *TargetUpdate {
	if s != nil {
		tu.SetHostname(*s)
	}
	return tu
}

// ClearHostname clears the value of the "hostname" field.
func (tu *TargetUpdate) ClearHostname() *TargetUpdate {
	tu.mutation.ClearHostname()
	return tu
}

// SetPort sets the "port" field.
func (tu *TargetUpdate) SetPort(u uint32) *TargetUpdate {
	tu.mutation.ResetPort()
	tu.mutation.SetPort(u)
	return tu
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (tu *TargetUpdate) SetNillablePort(u *uint32) *TargetUpdate {
	if u != nil {
		tu.SetPort(*u)
	}
	return tu
}

// AddPort adds u to the "port" field.
func (tu *TargetUpdate) AddPort(u int32) *TargetUpdate {
	tu.mutation.AddPort(u)
	return tu
}

// ClearPort clears the value of the "port" field.
func (tu *TargetUpdate) ClearPort() *TargetUpdate {
	tu.mutation.ClearPort()
	return tu
}

// Mutation returns the TargetMutation object of the builder.
func (tu *TargetUpdate) Mutation() *TargetMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TargetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TargetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TargetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TargetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TargetUpdate) check() error {
	if v, ok := tu.mutation.GetType(); ok {
		if err := target.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Target.type": %w`, err)}
		}
	}
	if v, ok := tu.mutation.IntervalSeconds(); ok {
		if err := target.IntervalSecondsValidator(v); err != nil {
			return &ValidationError{Name: "interval_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.interval_seconds": %w`, err)}
		}
	}
	if v, ok := tu.mutation.TimeoutSeconds(); ok {
		if err := target.TimeoutSecondsValidator(v); err != nil {
			return &ValidationError{Name: "timeout_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.timeout_seconds": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Port(); ok {
		if err := target.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Target.port": %w`, err)}
		}
	}
	return nil
}

func (tu *TargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: target.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(target.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(target.FieldType, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.IntervalSeconds(); ok {
		_spec.SetField(target.FieldIntervalSeconds, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedIntervalSeconds(); ok {
		_spec.AddField(target.FieldIntervalSeconds, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.TimeoutSeconds(); ok {
		_spec.SetField(target.FieldTimeoutSeconds, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedTimeoutSeconds(); ok {
		_spec.AddField(target.FieldTimeoutSeconds, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.URL(); ok {
		_spec.SetField(target.FieldURL, field.TypeString, value)
	}
	if tu.mutation.URLCleared() {
		_spec.ClearField(target.FieldURL, field.TypeString)
	}
	if value, ok := tu.mutation.Hostname(); ok {
		_spec.SetField(target.FieldHostname, field.TypeString, value)
	}
	if tu.mutation.HostnameCleared() {
		_spec.ClearField(target.FieldHostname, field.TypeString)
	}
	if value, ok := tu.mutation.Port(); ok {
		_spec.SetField(target.FieldPort, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedPort(); ok {
		_spec.AddField(target.FieldPort, field.TypeUint32, value)
	}
	if tu.mutation.PortCleared() {
		_spec.ClearField(target.FieldPort, field.TypeUint32)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{target.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TargetUpdateOne is the builder for updating a single Target entity.
type TargetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TargetMutation
}

// SetName sets the "name" field.
func (tuo *TargetUpdateOne) SetName(s string) *TargetUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TargetUpdateOne) SetType(t target.Type) *TargetUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetIntervalSeconds sets the "interval_seconds" field.
func (tuo *TargetUpdateOne) SetIntervalSeconds(u uint32) *TargetUpdateOne {
	tuo.mutation.ResetIntervalSeconds()
	tuo.mutation.SetIntervalSeconds(u)
	return tuo
}

// AddIntervalSeconds adds u to the "interval_seconds" field.
func (tuo *TargetUpdateOne) AddIntervalSeconds(u int32) *TargetUpdateOne {
	tuo.mutation.AddIntervalSeconds(u)
	return tuo
}

// SetTimeoutSeconds sets the "timeout_seconds" field.
func (tuo *TargetUpdateOne) SetTimeoutSeconds(u uint32) *TargetUpdateOne {
	tuo.mutation.ResetTimeoutSeconds()
	tuo.mutation.SetTimeoutSeconds(u)
	return tuo
}

// AddTimeoutSeconds adds u to the "timeout_seconds" field.
func (tuo *TargetUpdateOne) AddTimeoutSeconds(u int32) *TargetUpdateOne {
	tuo.mutation.AddTimeoutSeconds(u)
	return tuo
}

// SetURL sets the "url" field.
func (tuo *TargetUpdateOne) SetURL(s string) *TargetUpdateOne {
	tuo.mutation.SetURL(s)
	return tuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableURL(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetURL(*s)
	}
	return tuo
}

// ClearURL clears the value of the "url" field.
func (tuo *TargetUpdateOne) ClearURL() *TargetUpdateOne {
	tuo.mutation.ClearURL()
	return tuo
}

// SetHostname sets the "hostname" field.
func (tuo *TargetUpdateOne) SetHostname(s string) *TargetUpdateOne {
	tuo.mutation.SetHostname(s)
	return tuo
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableHostname(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetHostname(*s)
	}
	return tuo
}

// ClearHostname clears the value of the "hostname" field.
func (tuo *TargetUpdateOne) ClearHostname() *TargetUpdateOne {
	tuo.mutation.ClearHostname()
	return tuo
}

// SetPort sets the "port" field.
func (tuo *TargetUpdateOne) SetPort(u uint32) *TargetUpdateOne {
	tuo.mutation.ResetPort()
	tuo.mutation.SetPort(u)
	return tuo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillablePort(u *uint32) *TargetUpdateOne {
	if u != nil {
		tuo.SetPort(*u)
	}
	return tuo
}

// AddPort adds u to the "port" field.
func (tuo *TargetUpdateOne) AddPort(u int32) *TargetUpdateOne {
	tuo.mutation.AddPort(u)
	return tuo
}

// ClearPort clears the value of the "port" field.
func (tuo *TargetUpdateOne) ClearPort() *TargetUpdateOne {
	tuo.mutation.ClearPort()
	return tuo
}

// Mutation returns the TargetMutation object of the builder.
func (tuo *TargetUpdateOne) Mutation() *TargetMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TargetUpdateOne) Select(field string, fields ...string) *TargetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Target entity.
func (tuo *TargetUpdateOne) Save(ctx context.Context) (*Target, error) {
	var (
		err  error
		node *Target
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Target)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TargetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TargetUpdateOne) SaveX(ctx context.Context) *Target {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TargetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TargetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TargetUpdateOne) check() error {
	if v, ok := tuo.mutation.GetType(); ok {
		if err := target.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Target.type": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.IntervalSeconds(); ok {
		if err := target.IntervalSecondsValidator(v); err != nil {
			return &ValidationError{Name: "interval_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.interval_seconds": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.TimeoutSeconds(); ok {
		if err := target.TimeoutSecondsValidator(v); err != nil {
			return &ValidationError{Name: "timeout_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.timeout_seconds": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Port(); ok {
		if err := target.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Target.port": %w`, err)}
		}
	}
	return nil
}

func (tuo *TargetUpdateOne) sqlSave(ctx context.Context) (_node *Target, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: target.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Target.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, target.FieldID)
		for _, f := range fields {
			if !target.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != target.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(target.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(target.FieldType, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.IntervalSeconds(); ok {
		_spec.SetField(target.FieldIntervalSeconds, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedIntervalSeconds(); ok {
		_spec.AddField(target.FieldIntervalSeconds, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.TimeoutSeconds(); ok {
		_spec.SetField(target.FieldTimeoutSeconds, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedTimeoutSeconds(); ok {
		_spec.AddField(target.FieldTimeoutSeconds, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.URL(); ok {
		_spec.SetField(target.FieldURL, field.TypeString, value)
	}
	if tuo.mutation.URLCleared() {
		_spec.ClearField(target.FieldURL, field.TypeString)
	}
	if value, ok := tuo.mutation.Hostname(); ok {
		_spec.SetField(target.FieldHostname, field.TypeString, value)
	}
	if tuo.mutation.HostnameCleared() {
		_spec.ClearField(target.FieldHostname, field.TypeString)
	}
	if value, ok := tuo.mutation.Port(); ok {
		_spec.SetField(target.FieldPort, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedPort(); ok {
		_spec.AddField(target.FieldPort, field.TypeUint32, value)
	}
	if tuo.mutation.PortCleared() {
		_spec.ClearField(target.FieldPort, field.TypeUint32)
	}
	_node = &Target{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{target.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}