// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/ent/generated/database"
	"github.com/datumforge/geodetic/internal/ent/generated/group"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetCreatedBy sets the "created_by" field.
func (gc *GroupCreate) SetCreatedBy(s string) *GroupCreate {
	gc.mutation.SetCreatedBy(s)
	return gc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetCreatedBy(*s)
	}
	return gc
}

// SetUpdatedBy sets the "updated_by" field.
func (gc *GroupCreate) SetUpdatedBy(s string) *GroupCreate {
	gc.mutation.SetUpdatedBy(s)
	return gc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetUpdatedBy(*s)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GroupCreate) SetDeletedAt(t time.Time) *GroupCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetDeletedBy sets the "deleted_by" field.
func (gc *GroupCreate) SetDeletedBy(s string) *GroupCreate {
	gc.mutation.SetDeletedBy(s)
	return gc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetDeletedBy(*s)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GroupCreate) SetDescription(s string) *GroupCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDescription(s *string) *GroupCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetPrimaryLocation sets the "primary_location" field.
func (gc *GroupCreate) SetPrimaryLocation(s string) *GroupCreate {
	gc.mutation.SetPrimaryLocation(s)
	return gc
}

// SetLocations sets the "locations" field.
func (gc *GroupCreate) SetLocations(s []string) *GroupCreate {
	gc.mutation.SetLocations(s)
	return gc
}

// SetToken sets the "token" field.
func (gc *GroupCreate) SetToken(s string) *GroupCreate {
	gc.mutation.SetToken(s)
	return gc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (gc *GroupCreate) SetNillableToken(s *string) *GroupCreate {
	if s != nil {
		gc.SetToken(*s)
	}
	return gc
}

// SetRegion sets the "region" field.
func (gc *GroupCreate) SetRegion(e enums.Region) *GroupCreate {
	gc.mutation.SetRegion(e)
	return gc
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (gc *GroupCreate) SetNillableRegion(e *enums.Region) *GroupCreate {
	if e != nil {
		gc.SetRegion(*e)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GroupCreate) SetID(s string) *GroupCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableID(s *string) *GroupCreate {
	if s != nil {
		gc.SetID(*s)
	}
	return gc
}

// AddDatabaseIDs adds the "databases" edge to the Database entity by IDs.
func (gc *GroupCreate) AddDatabaseIDs(ids ...string) *GroupCreate {
	gc.mutation.AddDatabaseIDs(ids...)
	return gc
}

// AddDatabases adds the "databases" edges to the Database entity.
func (gc *GroupCreate) AddDatabases(d ...*Database) *GroupCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gc.AddDatabaseIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if err := gc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		if group.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		if group.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := group.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.Region(); !ok {
		v := group.DefaultRegion
		gc.mutation.SetRegion(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		if group.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultID (forgotten import generated/runtime?)")
		}
		v := group.DefaultID()
		gc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Group.name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.PrimaryLocation(); !ok {
		return &ValidationError{Name: "primary_location", err: errors.New(`generated: missing required field "Group.primary_location"`)}
	}
	if v, ok := gc.mutation.PrimaryLocation(); ok {
		if err := group.PrimaryLocationValidator(v); err != nil {
			return &ValidationError{Name: "primary_location", err: fmt.Errorf(`generated: validator failed for field "Group.primary_location": %w`, err)}
		}
	}
	if _, ok := gc.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`generated: missing required field "Group.region"`)}
	}
	if v, ok := gc.mutation.Region(); ok {
		if err := group.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`generated: validator failed for field "Group.region": %w`, err)}
		}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Group.ID type: %T", _spec.ID.Value)
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	)
	_spec.Schema = gc.schemaConfig.Group
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(group.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.CreatedBy(); ok {
		_spec.SetField(group.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gc.mutation.UpdatedBy(); ok {
		_spec.SetField(group.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gc.mutation.DeletedBy(); ok {
		_spec.SetField(group.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.PrimaryLocation(); ok {
		_spec.SetField(group.FieldPrimaryLocation, field.TypeString, value)
		_node.PrimaryLocation = value
	}
	if value, ok := gc.mutation.Locations(); ok {
		_spec.SetField(group.FieldLocations, field.TypeJSON, value)
		_node.Locations = value
	}
	if value, ok := gc.mutation.Token(); ok {
		_spec.SetField(group.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := gc.mutation.Region(); ok {
		_spec.SetField(group.FieldRegion, field.TypeEnum, value)
		_node.Region = value
	}
	if nodes := gc.mutation.DatabasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.DatabasesTable,
			Columns: []string{group.DatabasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(database.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.Database
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
