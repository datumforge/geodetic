package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/geodetic/internal/ent/mixin"
)

// Database holds the example schema definition for the Database entity
type Database struct {
	ent.Schema
}

// Fields of the Database
func (Database) Fields() []ent.Field {
	return []ent.Field{
		field.String("organization_id").
			Comment("the ID of the organization").
			NotEmpty(),
		field.String("name").
			Comment("the name to the database").
			NotEmpty(),
		field.String("geo").
			Comment("the geo location of the database").
			Optional(),
		field.String("dsn").
			Comment("the DSN to the database").
			NotEmpty(),
	}
}

func (Database) Indexes() []ent.Index {
	return []ent.Index{
		// organization_id should be unique
		index.Fields("organization_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Mixin of the Database
func (Database) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
	}
}

// Annotations of the Organization
func (Database) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
