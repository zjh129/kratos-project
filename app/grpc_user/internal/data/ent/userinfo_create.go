// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-project/app/grpc_user/internal/data/ent/userinfo"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfoCreate is the builder for creating a UserInfo entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// SetAccount sets the "account" field.
func (uic *UserInfoCreate) SetAccount(s string) *UserInfoCreate {
	uic.mutation.SetAccount(s)
	return uic
}

// SetPassword sets the "password" field.
func (uic *UserInfoCreate) SetPassword(s string) *UserInfoCreate {
	uic.mutation.SetPassword(s)
	return uic
}

// SetName sets the "name" field.
func (uic *UserInfoCreate) SetName(s string) *UserInfoCreate {
	uic.mutation.SetName(s)
	return uic
}

// SetAvatar sets the "avatar" field.
func (uic *UserInfoCreate) SetAvatar(s string) *UserInfoCreate {
	uic.mutation.SetAvatar(s)
	return uic
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableAvatar(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetAvatar(*s)
	}
	return uic
}

// SetIsEnable sets the "is_enable" field.
func (uic *UserInfoCreate) SetIsEnable(b bool) *UserInfoCreate {
	uic.mutation.SetIsEnable(b)
	return uic
}

// SetCreatedAt sets the "created_at" field.
func (uic *UserInfoCreate) SetCreatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetCreatedAt(t)
	return uic
}

// SetUpdatedAt sets the "updated_at" field.
func (uic *UserInfoCreate) SetUpdatedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetUpdatedAt(t)
	return uic
}

// SetDeletedAt sets the "deleted_at" field.
func (uic *UserInfoCreate) SetDeletedAt(t time.Time) *UserInfoCreate {
	uic.mutation.SetDeletedAt(t)
	return uic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableDeletedAt(t *time.Time) *UserInfoCreate {
	if t != nil {
		uic.SetDeletedAt(*t)
	}
	return uic
}

// SetID sets the "id" field.
func (uic *UserInfoCreate) SetID(i int) *UserInfoCreate {
	uic.mutation.SetID(i)
	return uic
}

// Mutation returns the UserInfoMutation object of the builder.
func (uic *UserInfoCreate) Mutation() *UserInfoMutation {
	return uic.mutation
}

// Save creates the UserInfo in the database.
func (uic *UserInfoCreate) Save(ctx context.Context) (*UserInfo, error) {
	return withHooks(ctx, uic.sqlSave, uic.mutation, uic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserInfoCreate) SaveX(ctx context.Context) *UserInfo {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uic *UserInfoCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserInfoCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	if _, ok := uic.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required field "UserInfo.account"`)}
	}
	if _, ok := uic.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "UserInfo.password"`)}
	}
	if _, ok := uic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserInfo.name"`)}
	}
	if _, ok := uic.mutation.IsEnable(); !ok {
		return &ValidationError{Name: "is_enable", err: errors.New(`ent: missing required field "UserInfo.is_enable"`)}
	}
	if _, ok := uic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserInfo.created_at"`)}
	}
	if _, ok := uic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserInfo.updated_at"`)}
	}
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*UserInfo, error) {
	if err := uic.check(); err != nil {
		return nil, err
	}
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uic.mutation.id = &_node.ID
	uic.mutation.done = true
	return _node, nil
}

func (uic *UserInfoCreate) createSpec() (*UserInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserInfo{config: uic.config}
		_spec = sqlgraph.NewCreateSpec(userinfo.Table, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	)
	if id, ok := uic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uic.mutation.Account(); ok {
		_spec.SetField(userinfo.FieldAccount, field.TypeString, value)
		_node.Account = value
	}
	if value, ok := uic.mutation.Password(); ok {
		_spec.SetField(userinfo.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uic.mutation.Name(); ok {
		_spec.SetField(userinfo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uic.mutation.Avatar(); ok {
		_spec.SetField(userinfo.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uic.mutation.IsEnable(); ok {
		_spec.SetField(userinfo.FieldIsEnable, field.TypeBool, value)
		_node.IsEnable = value
	}
	if value, ok := uic.mutation.CreatedAt(); ok {
		_spec.SetField(userinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uic.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uic.mutation.DeletedAt(); ok {
		_spec.SetField(userinfo.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating many UserInfo entities in bulk.
type UserInfoCreateBulk struct {
	config
	err      error
	builders []*UserInfoCreate
}

// Save creates the UserInfo entities in the database.
func (uicb *UserInfoCreateBulk) Save(ctx context.Context) ([]*UserInfo, error) {
	if uicb.err != nil {
		return nil, uicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*UserInfo, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*UserInfo {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}
