// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/datumforge/geodetic/internal/ent/generated/database"
	"github.com/datumforge/geodetic/internal/ent/generated/group"
	"github.com/datumforge/geodetic/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	databaseMixin := schema.Database{}.Mixin()
	databaseMixinHooks0 := databaseMixin[0].Hooks()
	databaseHooks := schema.Database{}.Hooks()
	database.Hooks[0] = databaseMixinHooks0[0]
	database.Hooks[1] = databaseHooks[0]
	database.Hooks[2] = databaseHooks[1]
	databaseMixinFields0 := databaseMixin[0].Fields()
	_ = databaseMixinFields0
	databaseMixinFields2 := databaseMixin[2].Fields()
	_ = databaseMixinFields2
	databaseFields := schema.Database{}.Fields()
	_ = databaseFields
	// databaseDescCreatedAt is the schema descriptor for created_at field.
	databaseDescCreatedAt := databaseMixinFields0[0].Descriptor()
	// database.DefaultCreatedAt holds the default value on creation for the created_at field.
	database.DefaultCreatedAt = databaseDescCreatedAt.Default.(func() time.Time)
	// databaseDescUpdatedAt is the schema descriptor for updated_at field.
	databaseDescUpdatedAt := databaseMixinFields0[1].Descriptor()
	// database.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	database.DefaultUpdatedAt = databaseDescUpdatedAt.Default.(func() time.Time)
	// database.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	database.UpdateDefaultUpdatedAt = databaseDescUpdatedAt.UpdateDefault.(func() time.Time)
	// databaseDescOrganizationID is the schema descriptor for organization_id field.
	databaseDescOrganizationID := databaseFields[0].Descriptor()
	// database.OrganizationIDValidator is a validator for the "organization_id" field. It is called by the builders before save.
	database.OrganizationIDValidator = databaseDescOrganizationID.Validators[0].(func(string) error)
	// databaseDescName is the schema descriptor for name field.
	databaseDescName := databaseFields[1].Descriptor()
	// database.NameValidator is a validator for the "name" field. It is called by the builders before save.
	database.NameValidator = databaseDescName.Validators[0].(func(string) error)
	// databaseDescDsn is the schema descriptor for dsn field.
	databaseDescDsn := databaseFields[3].Descriptor()
	// database.DsnValidator is a validator for the "dsn" field. It is called by the builders before save.
	database.DsnValidator = databaseDescDsn.Validators[0].(func(string) error)
	// databaseDescID is the schema descriptor for id field.
	databaseDescID := databaseMixinFields2[0].Descriptor()
	// database.DefaultID holds the default value on creation for the id field.
	database.DefaultID = databaseDescID.Default.(func() string)
	groupMixin := schema.Group{}.Mixin()
	groupMixinHooks0 := groupMixin[0].Hooks()
	groupHooks := schema.Group{}.Hooks()
	group.Hooks[0] = groupMixinHooks0[0]
	group.Hooks[1] = groupHooks[0]
	group.Hooks[2] = groupHooks[1]
	group.Hooks[3] = groupHooks[2]
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupMixinFields1 := groupMixin[1].Fields()
	_ = groupMixinFields1
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupMixinFields0[0].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescPrimaryLocation is the schema descriptor for primary_location field.
	groupDescPrimaryLocation := groupFields[2].Descriptor()
	// group.PrimaryLocationValidator is a validator for the "primary_location" field. It is called by the builders before save.
	group.PrimaryLocationValidator = groupDescPrimaryLocation.Validators[0].(func(string) error)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupMixinFields1[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() string)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
