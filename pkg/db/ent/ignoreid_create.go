// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/ignoreid"
	"github.com/google/uuid"
)

// IgnoreIDCreate is the builder for creating a IgnoreID entity.
type IgnoreIDCreate struct {
	config
	mutation *IgnoreIDMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (iic *IgnoreIDCreate) SetCreatedAt(u uint32) *IgnoreIDCreate {
	iic.mutation.SetCreatedAt(u)
	return iic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iic *IgnoreIDCreate) SetNillableCreatedAt(u *uint32) *IgnoreIDCreate {
	if u != nil {
		iic.SetCreatedAt(*u)
	}
	return iic
}

// SetUpdatedAt sets the "updated_at" field.
func (iic *IgnoreIDCreate) SetUpdatedAt(u uint32) *IgnoreIDCreate {
	iic.mutation.SetUpdatedAt(u)
	return iic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iic *IgnoreIDCreate) SetNillableUpdatedAt(u *uint32) *IgnoreIDCreate {
	if u != nil {
		iic.SetUpdatedAt(*u)
	}
	return iic
}

// SetDeletedAt sets the "deleted_at" field.
func (iic *IgnoreIDCreate) SetDeletedAt(u uint32) *IgnoreIDCreate {
	iic.mutation.SetDeletedAt(u)
	return iic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iic *IgnoreIDCreate) SetNillableDeletedAt(u *uint32) *IgnoreIDCreate {
	if u != nil {
		iic.SetDeletedAt(*u)
	}
	return iic
}

// SetAutoID sets the "auto_id" field.
func (iic *IgnoreIDCreate) SetAutoID(u uint32) *IgnoreIDCreate {
	iic.mutation.SetAutoID(u)
	return iic
}

// SetSampleCol sets the "sample_col" field.
func (iic *IgnoreIDCreate) SetSampleCol(s string) *IgnoreIDCreate {
	iic.mutation.SetSampleCol(s)
	return iic
}

// SetNillableSampleCol sets the "sample_col" field if the given value is not nil.
func (iic *IgnoreIDCreate) SetNillableSampleCol(s *string) *IgnoreIDCreate {
	if s != nil {
		iic.SetSampleCol(*s)
	}
	return iic
}

// SetID sets the "id" field.
func (iic *IgnoreIDCreate) SetID(u uuid.UUID) *IgnoreIDCreate {
	iic.mutation.SetID(u)
	return iic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (iic *IgnoreIDCreate) SetNillableID(u *uuid.UUID) *IgnoreIDCreate {
	if u != nil {
		iic.SetID(*u)
	}
	return iic
}

// Mutation returns the IgnoreIDMutation object of the builder.
func (iic *IgnoreIDCreate) Mutation() *IgnoreIDMutation {
	return iic.mutation
}

// Save creates the IgnoreID in the database.
func (iic *IgnoreIDCreate) Save(ctx context.Context) (*IgnoreID, error) {
	var (
		err  error
		node *IgnoreID
	)
	if err := iic.defaults(); err != nil {
		return nil, err
	}
	if len(iic.hooks) == 0 {
		if err = iic.check(); err != nil {
			return nil, err
		}
		node, err = iic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IgnoreIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iic.check(); err != nil {
				return nil, err
			}
			iic.mutation = mutation
			if node, err = iic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(iic.hooks) - 1; i >= 0; i-- {
			if iic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*IgnoreID)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IgnoreIDMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iic *IgnoreIDCreate) SaveX(ctx context.Context) *IgnoreID {
	v, err := iic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iic *IgnoreIDCreate) Exec(ctx context.Context) error {
	_, err := iic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iic *IgnoreIDCreate) ExecX(ctx context.Context) {
	if err := iic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iic *IgnoreIDCreate) defaults() error {
	if _, ok := iic.mutation.CreatedAt(); !ok {
		if ignoreid.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized ignoreid.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := ignoreid.DefaultCreatedAt()
		iic.mutation.SetCreatedAt(v)
	}
	if _, ok := iic.mutation.UpdatedAt(); !ok {
		if ignoreid.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ignoreid.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ignoreid.DefaultUpdatedAt()
		iic.mutation.SetUpdatedAt(v)
	}
	if _, ok := iic.mutation.DeletedAt(); !ok {
		if ignoreid.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized ignoreid.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := ignoreid.DefaultDeletedAt()
		iic.mutation.SetDeletedAt(v)
	}
	if _, ok := iic.mutation.SampleCol(); !ok {
		v := ignoreid.DefaultSampleCol
		iic.mutation.SetSampleCol(v)
	}
	if _, ok := iic.mutation.ID(); !ok {
		if ignoreid.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized ignoreid.DefaultID (forgotten import ent/runtime?)")
		}
		v := ignoreid.DefaultID()
		iic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (iic *IgnoreIDCreate) check() error {
	if _, ok := iic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "IgnoreID.created_at"`)}
	}
	if _, ok := iic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "IgnoreID.updated_at"`)}
	}
	if _, ok := iic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "IgnoreID.deleted_at"`)}
	}
	if _, ok := iic.mutation.AutoID(); !ok {
		return &ValidationError{Name: "auto_id", err: errors.New(`ent: missing required field "IgnoreID.auto_id"`)}
	}
	return nil
}

func (iic *IgnoreIDCreate) sqlSave(ctx context.Context) (*IgnoreID, error) {
	_node, _spec := iic.createSpec()
	if err := sqlgraph.CreateNode(ctx, iic.driver, _spec); err != nil {
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

func (iic *IgnoreIDCreate) createSpec() (*IgnoreID, *sqlgraph.CreateSpec) {
	var (
		_node = &IgnoreID{config: iic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ignoreid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ignoreid.FieldID,
			},
		}
	)
	_spec.OnConflict = iic.conflict
	if id, ok := iic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := iic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ignoreid.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := iic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ignoreid.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := iic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ignoreid.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := iic.mutation.AutoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ignoreid.FieldAutoID,
		})
		_node.AutoID = value
	}
	if value, ok := iic.mutation.SampleCol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ignoreid.FieldSampleCol,
		})
		_node.SampleCol = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IgnoreID.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IgnoreIDUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (iic *IgnoreIDCreate) OnConflict(opts ...sql.ConflictOption) *IgnoreIDUpsertOne {
	iic.conflict = opts
	return &IgnoreIDUpsertOne{
		create: iic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IgnoreID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (iic *IgnoreIDCreate) OnConflictColumns(columns ...string) *IgnoreIDUpsertOne {
	iic.conflict = append(iic.conflict, sql.ConflictColumns(columns...))
	return &IgnoreIDUpsertOne{
		create: iic,
	}
}

type (
	// IgnoreIDUpsertOne is the builder for "upsert"-ing
	//  one IgnoreID node.
	IgnoreIDUpsertOne struct {
		create *IgnoreIDCreate
	}

	// IgnoreIDUpsert is the "OnConflict" setter.
	IgnoreIDUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *IgnoreIDUpsert) SetCreatedAt(v uint32) *IgnoreIDUpsert {
	u.Set(ignoreid.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IgnoreIDUpsert) UpdateCreatedAt() *IgnoreIDUpsert {
	u.SetExcluded(ignoreid.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *IgnoreIDUpsert) AddCreatedAt(v uint32) *IgnoreIDUpsert {
	u.Add(ignoreid.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *IgnoreIDUpsert) SetUpdatedAt(v uint32) *IgnoreIDUpsert {
	u.Set(ignoreid.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *IgnoreIDUpsert) UpdateUpdatedAt() *IgnoreIDUpsert {
	u.SetExcluded(ignoreid.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *IgnoreIDUpsert) AddUpdatedAt(v uint32) *IgnoreIDUpsert {
	u.Add(ignoreid.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IgnoreIDUpsert) SetDeletedAt(v uint32) *IgnoreIDUpsert {
	u.Set(ignoreid.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IgnoreIDUpsert) UpdateDeletedAt() *IgnoreIDUpsert {
	u.SetExcluded(ignoreid.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *IgnoreIDUpsert) AddDeletedAt(v uint32) *IgnoreIDUpsert {
	u.Add(ignoreid.FieldDeletedAt, v)
	return u
}

// SetAutoID sets the "auto_id" field.
func (u *IgnoreIDUpsert) SetAutoID(v uint32) *IgnoreIDUpsert {
	u.Set(ignoreid.FieldAutoID, v)
	return u
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *IgnoreIDUpsert) UpdateAutoID() *IgnoreIDUpsert {
	u.SetExcluded(ignoreid.FieldAutoID)
	return u
}

// AddAutoID adds v to the "auto_id" field.
func (u *IgnoreIDUpsert) AddAutoID(v uint32) *IgnoreIDUpsert {
	u.Add(ignoreid.FieldAutoID, v)
	return u
}

// SetSampleCol sets the "sample_col" field.
func (u *IgnoreIDUpsert) SetSampleCol(v string) *IgnoreIDUpsert {
	u.Set(ignoreid.FieldSampleCol, v)
	return u
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *IgnoreIDUpsert) UpdateSampleCol() *IgnoreIDUpsert {
	u.SetExcluded(ignoreid.FieldSampleCol)
	return u
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *IgnoreIDUpsert) ClearSampleCol() *IgnoreIDUpsert {
	u.SetNull(ignoreid.FieldSampleCol)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.IgnoreID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ignoreid.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *IgnoreIDUpsertOne) UpdateNewValues() *IgnoreIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(ignoreid.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.IgnoreID.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *IgnoreIDUpsertOne) Ignore() *IgnoreIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IgnoreIDUpsertOne) DoNothing() *IgnoreIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IgnoreIDCreate.OnConflict
// documentation for more info.
func (u *IgnoreIDUpsertOne) Update(set func(*IgnoreIDUpsert)) *IgnoreIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IgnoreIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *IgnoreIDUpsertOne) SetCreatedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *IgnoreIDUpsertOne) AddCreatedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertOne) UpdateCreatedAt() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *IgnoreIDUpsertOne) SetUpdatedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *IgnoreIDUpsertOne) AddUpdatedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertOne) UpdateUpdatedAt() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IgnoreIDUpsertOne) SetDeletedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *IgnoreIDUpsertOne) AddDeletedAt(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertOne) UpdateDeletedAt() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAutoID sets the "auto_id" field.
func (u *IgnoreIDUpsertOne) SetAutoID(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetAutoID(v)
	})
}

// AddAutoID adds v to the "auto_id" field.
func (u *IgnoreIDUpsertOne) AddAutoID(v uint32) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddAutoID(v)
	})
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *IgnoreIDUpsertOne) UpdateAutoID() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateAutoID()
	})
}

// SetSampleCol sets the "sample_col" field.
func (u *IgnoreIDUpsertOne) SetSampleCol(v string) *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetSampleCol(v)
	})
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *IgnoreIDUpsertOne) UpdateSampleCol() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateSampleCol()
	})
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *IgnoreIDUpsertOne) ClearSampleCol() *IgnoreIDUpsertOne {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.ClearSampleCol()
	})
}

// Exec executes the query.
func (u *IgnoreIDUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IgnoreIDCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IgnoreIDUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *IgnoreIDUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: IgnoreIDUpsertOne.ID is not supported by MySQL driver. Use IgnoreIDUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *IgnoreIDUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// IgnoreIDCreateBulk is the builder for creating many IgnoreID entities in bulk.
type IgnoreIDCreateBulk struct {
	config
	builders []*IgnoreIDCreate
	conflict []sql.ConflictOption
}

// Save creates the IgnoreID entities in the database.
func (iicb *IgnoreIDCreateBulk) Save(ctx context.Context) ([]*IgnoreID, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iicb.builders))
	nodes := make([]*IgnoreID, len(iicb.builders))
	mutators := make([]Mutator, len(iicb.builders))
	for i := range iicb.builders {
		func(i int, root context.Context) {
			builder := iicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IgnoreIDMutation)
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
					_, err = mutators[i+1].Mutate(root, iicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = iicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iicb *IgnoreIDCreateBulk) SaveX(ctx context.Context) []*IgnoreID {
	v, err := iicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iicb *IgnoreIDCreateBulk) Exec(ctx context.Context) error {
	_, err := iicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iicb *IgnoreIDCreateBulk) ExecX(ctx context.Context) {
	if err := iicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IgnoreID.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IgnoreIDUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (iicb *IgnoreIDCreateBulk) OnConflict(opts ...sql.ConflictOption) *IgnoreIDUpsertBulk {
	iicb.conflict = opts
	return &IgnoreIDUpsertBulk{
		create: iicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IgnoreID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (iicb *IgnoreIDCreateBulk) OnConflictColumns(columns ...string) *IgnoreIDUpsertBulk {
	iicb.conflict = append(iicb.conflict, sql.ConflictColumns(columns...))
	return &IgnoreIDUpsertBulk{
		create: iicb,
	}
}

// IgnoreIDUpsertBulk is the builder for "upsert"-ing
// a bulk of IgnoreID nodes.
type IgnoreIDUpsertBulk struct {
	create *IgnoreIDCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.IgnoreID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ignoreid.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *IgnoreIDUpsertBulk) UpdateNewValues() *IgnoreIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(ignoreid.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.IgnoreID.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *IgnoreIDUpsertBulk) Ignore() *IgnoreIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IgnoreIDUpsertBulk) DoNothing() *IgnoreIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IgnoreIDCreateBulk.OnConflict
// documentation for more info.
func (u *IgnoreIDUpsertBulk) Update(set func(*IgnoreIDUpsert)) *IgnoreIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IgnoreIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *IgnoreIDUpsertBulk) SetCreatedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *IgnoreIDUpsertBulk) AddCreatedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertBulk) UpdateCreatedAt() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *IgnoreIDUpsertBulk) SetUpdatedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *IgnoreIDUpsertBulk) AddUpdatedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertBulk) UpdateUpdatedAt() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IgnoreIDUpsertBulk) SetDeletedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *IgnoreIDUpsertBulk) AddDeletedAt(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IgnoreIDUpsertBulk) UpdateDeletedAt() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAutoID sets the "auto_id" field.
func (u *IgnoreIDUpsertBulk) SetAutoID(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetAutoID(v)
	})
}

// AddAutoID adds v to the "auto_id" field.
func (u *IgnoreIDUpsertBulk) AddAutoID(v uint32) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.AddAutoID(v)
	})
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *IgnoreIDUpsertBulk) UpdateAutoID() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateAutoID()
	})
}

// SetSampleCol sets the "sample_col" field.
func (u *IgnoreIDUpsertBulk) SetSampleCol(v string) *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.SetSampleCol(v)
	})
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *IgnoreIDUpsertBulk) UpdateSampleCol() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.UpdateSampleCol()
	})
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *IgnoreIDUpsertBulk) ClearSampleCol() *IgnoreIDUpsertBulk {
	return u.Update(func(s *IgnoreIDUpsert) {
		s.ClearSampleCol()
	})
}

// Exec executes the query.
func (u *IgnoreIDUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the IgnoreIDCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IgnoreIDCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IgnoreIDUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
