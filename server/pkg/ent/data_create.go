// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/manhrev/IOT/server/pkg/ent/data"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
)

// DataCreate is the builder for creating a Data entity.
type DataCreate struct {
	config
	mutation *DataMutation
	hooks    []Hook
}

// SetData sets the "data" field.
func (dc *DataCreate) SetData(i int64) *DataCreate {
	dc.mutation.SetData(i)
	return dc
}

// SetRecordedAt sets the "recorded_at" field.
func (dc *DataCreate) SetRecordedAt(t time.Time) *DataCreate {
	dc.mutation.SetRecordedAt(t)
	return dc
}

// SetNillableRecordedAt sets the "recorded_at" field if the given value is not nil.
func (dc *DataCreate) SetNillableRecordedAt(t *time.Time) *DataCreate {
	if t != nil {
		dc.SetRecordedAt(*t)
	}
	return dc
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (dc *DataCreate) SetFeedID(id int) *DataCreate {
	dc.mutation.SetFeedID(id)
	return dc
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (dc *DataCreate) SetNillableFeedID(id *int) *DataCreate {
	if id != nil {
		dc = dc.SetFeedID(*id)
	}
	return dc
}

// SetFeed sets the "feed" edge to the Feed entity.
func (dc *DataCreate) SetFeed(f *Feed) *DataCreate {
	return dc.SetFeedID(f.ID)
}

// Mutation returns the DataMutation object of the builder.
func (dc *DataCreate) Mutation() *DataMutation {
	return dc.mutation
}

// Save creates the Data in the database.
func (dc *DataCreate) Save(ctx context.Context) (*Data, error) {
	var (
		err  error
		node *Data
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DataCreate) SaveX(ctx context.Context) *Data {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DataCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DataCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DataCreate) defaults() {
	if _, ok := dc.mutation.RecordedAt(); !ok {
		v := data.DefaultRecordedAt
		dc.mutation.SetRecordedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DataCreate) check() error {
	if _, ok := dc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Data.data"`)}
	}
	if _, ok := dc.mutation.RecordedAt(); !ok {
		return &ValidationError{Name: "recorded_at", err: errors.New(`ent: missing required field "Data.recorded_at"`)}
	}
	return nil
}

func (dc *DataCreate) sqlSave(ctx context.Context) (*Data, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DataCreate) createSpec() (*Data, *sqlgraph.CreateSpec) {
	var (
		_node = &Data{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: data.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: data.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: data.FieldData,
		})
		_node.Data = value
	}
	if value, ok := dc.mutation.RecordedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: data.FieldRecordedAt,
		})
		_node.RecordedAt = value
	}
	if nodes := dc.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   data.FeedTable,
			Columns: []string{data.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.feed_data = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DataCreateBulk is the builder for creating many Data entities in bulk.
type DataCreateBulk struct {
	config
	builders []*DataCreate
}

// Save creates the Data entities in the database.
func (dcb *DataCreateBulk) Save(ctx context.Context) ([]*Data, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Data, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DataMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DataCreateBulk) SaveX(ctx context.Context) []*Data {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DataCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DataCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
