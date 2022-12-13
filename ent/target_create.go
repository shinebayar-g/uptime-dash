// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/target"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TargetCreate is the builder for creating a Target entity.
type TargetCreate struct {
	config
	mutation *TargetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TargetCreate) SetName(s string) *TargetCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetType sets the "type" field.
func (tc *TargetCreate) SetType(t target.Type) *TargetCreate {
	tc.mutation.SetType(t)
	return tc
}

// SetIntervalSeconds sets the "interval_seconds" field.
func (tc *TargetCreate) SetIntervalSeconds(u uint32) *TargetCreate {
	tc.mutation.SetIntervalSeconds(u)
	return tc
}

// SetTimeoutSeconds sets the "timeout_seconds" field.
func (tc *TargetCreate) SetTimeoutSeconds(u uint32) *TargetCreate {
	tc.mutation.SetTimeoutSeconds(u)
	return tc
}

// SetURL sets the "url" field.
func (tc *TargetCreate) SetURL(s string) *TargetCreate {
	tc.mutation.SetURL(s)
	return tc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (tc *TargetCreate) SetNillableURL(s *string) *TargetCreate {
	if s != nil {
		tc.SetURL(*s)
	}
	return tc
}

// SetHostname sets the "hostname" field.
func (tc *TargetCreate) SetHostname(s string) *TargetCreate {
	tc.mutation.SetHostname(s)
	return tc
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (tc *TargetCreate) SetNillableHostname(s *string) *TargetCreate {
	if s != nil {
		tc.SetHostname(*s)
	}
	return tc
}

// SetPort sets the "port" field.
func (tc *TargetCreate) SetPort(u uint32) *TargetCreate {
	tc.mutation.SetPort(u)
	return tc
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (tc *TargetCreate) SetNillablePort(u *uint32) *TargetCreate {
	if u != nil {
		tc.SetPort(*u)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TargetCreate) SetID(u uuid.UUID) *TargetCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TargetCreate) SetNillableID(u *uuid.UUID) *TargetCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TargetMutation object of the builder.
func (tc *TargetCreate) Mutation() *TargetMutation {
	return tc.mutation
}

// Save creates the Target in the database.
func (tc *TargetCreate) Save(ctx context.Context) (*Target, error) {
	var (
		err  error
		node *Target
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (tc *TargetCreate) SaveX(ctx context.Context) *Target {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TargetCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TargetCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TargetCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := target.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TargetCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Target.name"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Target.type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := target.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Target.type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.IntervalSeconds(); !ok {
		return &ValidationError{Name: "interval_seconds", err: errors.New(`ent: missing required field "Target.interval_seconds"`)}
	}
	if v, ok := tc.mutation.IntervalSeconds(); ok {
		if err := target.IntervalSecondsValidator(v); err != nil {
			return &ValidationError{Name: "interval_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.interval_seconds": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TimeoutSeconds(); !ok {
		return &ValidationError{Name: "timeout_seconds", err: errors.New(`ent: missing required field "Target.timeout_seconds"`)}
	}
	if v, ok := tc.mutation.TimeoutSeconds(); ok {
		if err := target.TimeoutSecondsValidator(v); err != nil {
			return &ValidationError{Name: "timeout_seconds", err: fmt.Errorf(`ent: validator failed for field "Target.timeout_seconds": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Port(); ok {
		if err := target.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Target.port": %w`, err)}
		}
	}
	return nil
}

func (tc *TargetCreate) sqlSave(ctx context.Context) (*Target, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TargetCreate) createSpec() (*Target, *sqlgraph.CreateSpec) {
	var (
		_node = &Target{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: target.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: target.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(target.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(target.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.IntervalSeconds(); ok {
		_spec.SetField(target.FieldIntervalSeconds, field.TypeUint32, value)
		_node.IntervalSeconds = value
	}
	if value, ok := tc.mutation.TimeoutSeconds(); ok {
		_spec.SetField(target.FieldTimeoutSeconds, field.TypeUint32, value)
		_node.TimeoutSeconds = value
	}
	if value, ok := tc.mutation.URL(); ok {
		_spec.SetField(target.FieldURL, field.TypeString, value)
		_node.URL = &value
	}
	if value, ok := tc.mutation.Hostname(); ok {
		_spec.SetField(target.FieldHostname, field.TypeString, value)
		_node.Hostname = &value
	}
	if value, ok := tc.mutation.Port(); ok {
		_spec.SetField(target.FieldPort, field.TypeUint32, value)
		_node.Port = &value
	}
	return _node, _spec
}

// TargetCreateBulk is the builder for creating many Target entities in bulk.
type TargetCreateBulk struct {
	config
	builders []*TargetCreate
}

// Save creates the Target entities in the database.
func (tcb *TargetCreateBulk) Save(ctx context.Context) ([]*Target, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Target, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TargetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TargetCreateBulk) SaveX(ctx context.Context) []*Target {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TargetCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TargetCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}