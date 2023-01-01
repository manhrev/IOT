// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/user"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/usersetting"
)

// UserSettingCreate is the builder for creating a UserSetting entity.
type UserSettingCreate struct {
	config
	mutation *UserSettingMutation
	hooks    []Hook
}

// SetRegion sets the "region" field.
func (usc *UserSettingCreate) SetRegion(s string) *UserSettingCreate {
	usc.mutation.SetRegion(s)
	return usc
}

// SetLanguage sets the "language" field.
func (usc *UserSettingCreate) SetLanguage(s string) *UserSettingCreate {
	usc.mutation.SetLanguage(s)
	return usc
}

// SetIsNotification sets the "is_notification" field.
func (usc *UserSettingCreate) SetIsNotification(s string) *UserSettingCreate {
	usc.mutation.SetIsNotification(s)
	return usc
}

// SetDateModified sets the "date_modified" field.
func (usc *UserSettingCreate) SetDateModified(t time.Time) *UserSettingCreate {
	usc.mutation.SetDateModified(t)
	return usc
}

// SetNillableDateModified sets the "date_modified" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableDateModified(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetDateModified(*t)
	}
	return usc
}

// SetID sets the "id" field.
func (usc *UserSettingCreate) SetID(i int64) *UserSettingCreate {
	usc.mutation.SetID(i)
	return usc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usc *UserSettingCreate) SetUserID(id int64) *UserSettingCreate {
	usc.mutation.SetUserID(id)
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSettingCreate) SetUser(u *User) *UserSettingCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSettingMutation object of the builder.
func (usc *UserSettingCreate) Mutation() *UserSettingMutation {
	return usc.mutation
}

// Save creates the UserSetting in the database.
func (usc *UserSettingCreate) Save(ctx context.Context) (*UserSetting, error) {
	var (
		err  error
		node *UserSetting
	)
	usc.defaults()
	if len(usc.hooks) == 0 {
		if err = usc.check(); err != nil {
			return nil, err
		}
		node, err = usc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usc.check(); err != nil {
				return nil, err
			}
			usc.mutation = mutation
			if node, err = usc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(usc.hooks) - 1; i >= 0; i-- {
			if usc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, usc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSettingCreate) SaveX(ctx context.Context) *UserSetting {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSettingCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSettingCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSettingCreate) defaults() {
	if _, ok := usc.mutation.DateModified(); !ok {
		v := usersetting.DefaultDateModified
		usc.mutation.SetDateModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSettingCreate) check() error {
	if _, ok := usc.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`ent: missing required field "UserSetting.region"`)}
	}
	if _, ok := usc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "UserSetting.language"`)}
	}
	if _, ok := usc.mutation.IsNotification(); !ok {
		return &ValidationError{Name: "is_notification", err: errors.New(`ent: missing required field "UserSetting.is_notification"`)}
	}
	if _, ok := usc.mutation.DateModified(); !ok {
		return &ValidationError{Name: "date_modified", err: errors.New(`ent: missing required field "UserSetting.date_modified"`)}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserSetting.user"`)}
	}
	return nil
}

func (usc *UserSettingCreate) sqlSave(ctx context.Context) (*UserSetting, error) {
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (usc *UserSettingCreate) createSpec() (*UserSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSetting{config: usc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usersetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usersetting.FieldID,
			},
		}
	)
	if id, ok := usc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := usc.mutation.Region(); ok {
		_spec.SetField(usersetting.FieldRegion, field.TypeString, value)
		_node.Region = value
	}
	if value, ok := usc.mutation.Language(); ok {
		_spec.SetField(usersetting.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := usc.mutation.IsNotification(); ok {
		_spec.SetField(usersetting.FieldIsNotification, field.TypeString, value)
		_node.IsNotification = value
	}
	if value, ok := usc.mutation.DateModified(); ok {
		_spec.SetField(usersetting.FieldDateModified, field.TypeTime, value)
		_node.DateModified = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersetting.UserTable,
			Columns: []string{usersetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_setting = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSettingCreateBulk is the builder for creating many UserSetting entities in bulk.
type UserSettingCreateBulk struct {
	config
	builders []*UserSettingCreate
}

// Save creates the UserSetting entities in the database.
func (uscb *UserSettingCreateBulk) Save(ctx context.Context) ([]*UserSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSetting, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) SaveX(ctx context.Context) []*UserSetting {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}