package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/ent/hooks"
	"github.com/datumforge/geodetic/internal/ent/mixin"
)

// Group holds the schema definition for the Group entity
type Group struct {
	ent.Schema
}

// Fields of the Group
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the group in turso").
			NotEmpty(),
		field.String("description").
			Comment("the description of the group").
			Optional(),
		field.String("primary_location").
			Comment("the primary of the group").
			Optional(),
		field.Strings("locations").
			Comment("the replica locations of the group").
			Optional(),
		field.String("token").
			Sensitive().
			Comment("the auth token used to connect to the group").
			NotEmpty(),
		field.Enum("region").
			GoType(enums.Region("")).
			Comment("region the group").
			Default(string(enums.Amer)),
	}
}

// Mixin of the Group
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Hooks of the Group
func (Group) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookGroupCreate(),
		hooks.HookGroupUpdate(),
		hooks.HookGroupDelete(),
	}
}

// Annotations of the Group
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
