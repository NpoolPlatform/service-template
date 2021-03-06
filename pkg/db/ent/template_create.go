// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/template"
	"github.com/google/uuid"
)

// TemplateCreate is the builder for creating a Template entity.
type TemplateCreate struct {
	config
	mutation *TemplateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TemplateCreate) SetCreatedAt(u uint32) *TemplateCreate {
	tc.mutation.SetCreatedAt(u)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreatedAt(u *uint32) *TemplateCreate {
	if u != nil {
		tc.SetCreatedAt(*u)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TemplateCreate) SetUpdatedAt(u uint32) *TemplateCreate {
	tc.mutation.SetUpdatedAt(u)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableUpdatedAt(u *uint32) *TemplateCreate {
	if u != nil {
		tc.SetUpdatedAt(*u)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TemplateCreate) SetDeletedAt(u uint32) *TemplateCreate {
	tc.mutation.SetDeletedAt(u)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableDeletedAt(u *uint32) *TemplateCreate {
	if u != nil {
		tc.SetDeletedAt(*u)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TemplateCreate) SetName(s string) *TemplateCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetAge sets the "age" field.
func (tc *TemplateCreate) SetAge(u uint32) *TemplateCreate {
	tc.mutation.SetAge(u)
	return tc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableAge(u *uint32) *TemplateCreate {
	if u != nil {
		tc.SetAge(*u)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TemplateCreate) SetID(u uuid.UUID) *TemplateCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableID(u *uuid.UUID) *TemplateCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TemplateMutation object of the builder.
func (tc *TemplateCreate) Mutation() *TemplateMutation {
	return tc.mutation
}

// Save creates the Template in the database.
func (tc *TemplateCreate) Save(ctx context.Context) (*Template, error) {
	var (
		err  error
		node *Template
	)
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
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
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemplateCreate) SaveX(ctx context.Context) *Template {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TemplateCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TemplateCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TemplateCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if template.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized template.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := template.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if template.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized template.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := template.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		if template.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized template.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := template.DefaultDeletedAt()
		tc.mutation.SetDeletedAt(v)
	}
	if _, ok := tc.mutation.Age(); !ok {
		v := template.DefaultAge
		tc.mutation.SetAge(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if template.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized template.DefaultID (forgotten import ent/runtime?)")
		}
		v := template.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TemplateCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Template.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Template.updated_at"`)}
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Template.deleted_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Template.name"`)}
	}
	if _, ok := tc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Template.age"`)}
	}
	return nil
}

func (tc *TemplateCreate) sqlSave(ctx context.Context) (*Template, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (tc *TemplateCreate) createSpec() (*Template, *sqlgraph.CreateSpec) {
	var (
		_node = &Template{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: template.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: template.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: template.FieldAge,
		})
		_node.Age = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Template.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TemplateCreate) OnConflict(opts ...sql.ConflictOption) *TemplateUpsertOne {
	tc.conflict = opts
	return &TemplateUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TemplateCreate) OnConflictColumns(columns ...string) *TemplateUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TemplateUpsertOne{
		create: tc,
	}
}

type (
	// TemplateUpsertOne is the builder for "upsert"-ing
	//  one Template node.
	TemplateUpsertOne struct {
		create *TemplateCreate
	}

	// TemplateUpsert is the "OnConflict" setter.
	TemplateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TemplateUpsert) SetCreatedAt(v uint32) *TemplateUpsert {
	u.Set(template.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateCreatedAt() *TemplateUpsert {
	u.SetExcluded(template.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TemplateUpsert) AddCreatedAt(v uint32) *TemplateUpsert {
	u.Add(template.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TemplateUpsert) SetUpdatedAt(v uint32) *TemplateUpsert {
	u.Set(template.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateUpdatedAt() *TemplateUpsert {
	u.SetExcluded(template.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TemplateUpsert) AddUpdatedAt(v uint32) *TemplateUpsert {
	u.Add(template.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TemplateUpsert) SetDeletedAt(v uint32) *TemplateUpsert {
	u.Set(template.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateDeletedAt() *TemplateUpsert {
	u.SetExcluded(template.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TemplateUpsert) AddDeletedAt(v uint32) *TemplateUpsert {
	u.Add(template.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *TemplateUpsert) SetName(v string) *TemplateUpsert {
	u.Set(template.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateName() *TemplateUpsert {
	u.SetExcluded(template.FieldName)
	return u
}

// SetAge sets the "age" field.
func (u *TemplateUpsert) SetAge(v uint32) *TemplateUpsert {
	u.Set(template.FieldAge, v)
	return u
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TemplateUpsert) UpdateAge() *TemplateUpsert {
	u.SetExcluded(template.FieldAge)
	return u
}

// AddAge adds v to the "age" field.
func (u *TemplateUpsert) AddAge(v uint32) *TemplateUpsert {
	u.Add(template.FieldAge, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(template.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TemplateUpsertOne) UpdateNewValues() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(template.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Template.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TemplateUpsertOne) Ignore() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateUpsertOne) DoNothing() *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateCreate.OnConflict
// documentation for more info.
func (u *TemplateUpsertOne) Update(set func(*TemplateUpsert)) *TemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TemplateUpsertOne) SetCreatedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TemplateUpsertOne) AddCreatedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateCreatedAt() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TemplateUpsertOne) SetUpdatedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TemplateUpsertOne) AddUpdatedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateUpdatedAt() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TemplateUpsertOne) SetDeletedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TemplateUpsertOne) AddDeletedAt(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateDeletedAt() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *TemplateUpsertOne) SetName(v string) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateName() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateName()
	})
}

// SetAge sets the "age" field.
func (u *TemplateUpsertOne) SetAge(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *TemplateUpsertOne) AddAge(v uint32) *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TemplateUpsertOne) UpdateAge() *TemplateUpsertOne {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateAge()
	})
}

// Exec executes the query.
func (u *TemplateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TemplateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TemplateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TemplateUpsertOne.ID is not supported by MySQL driver. Use TemplateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TemplateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TemplateCreateBulk is the builder for creating many Template entities in bulk.
type TemplateCreateBulk struct {
	config
	builders []*TemplateCreate
	conflict []sql.ConflictOption
}

// Save creates the Template entities in the database.
func (tcb *TemplateCreateBulk) Save(ctx context.Context) ([]*Template, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Template, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateMutation)
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
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
func (tcb *TemplateCreateBulk) SaveX(ctx context.Context) []*Template {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TemplateCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Template.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TemplateCreateBulk) OnConflict(opts ...sql.ConflictOption) *TemplateUpsertBulk {
	tcb.conflict = opts
	return &TemplateUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TemplateCreateBulk) OnConflictColumns(columns ...string) *TemplateUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TemplateUpsertBulk{
		create: tcb,
	}
}

// TemplateUpsertBulk is the builder for "upsert"-ing
// a bulk of Template nodes.
type TemplateUpsertBulk struct {
	create *TemplateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(template.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TemplateUpsertBulk) UpdateNewValues() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(template.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Template.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TemplateUpsertBulk) Ignore() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateUpsertBulk) DoNothing() *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateCreateBulk.OnConflict
// documentation for more info.
func (u *TemplateUpsertBulk) Update(set func(*TemplateUpsert)) *TemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TemplateUpsertBulk) SetCreatedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TemplateUpsertBulk) AddCreatedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateCreatedAt() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TemplateUpsertBulk) SetUpdatedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TemplateUpsertBulk) AddUpdatedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateUpdatedAt() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TemplateUpsertBulk) SetDeletedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TemplateUpsertBulk) AddDeletedAt(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateDeletedAt() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *TemplateUpsertBulk) SetName(v string) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateName() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateName()
	})
}

// SetAge sets the "age" field.
func (u *TemplateUpsertBulk) SetAge(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *TemplateUpsertBulk) AddAge(v uint32) *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TemplateUpsertBulk) UpdateAge() *TemplateUpsertBulk {
	return u.Update(func(s *TemplateUpsert) {
		s.UpdateAge()
	})
}

// Exec executes the query.
func (u *TemplateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TemplateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TemplateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
