// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DatabasesColumns holds the columns for the "databases" table.
	DatabasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "organization_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "geo", Type: field.TypeString, Nullable: true},
		{Name: "dsn", Type: field.TypeString},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "CREATING", "DELETING", "DELETED"}, Default: "CREATING"},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"LOCAL", "TURSO"}, Default: "LOCAL"},
		{Name: "group_id", Type: field.TypeString},
	}
	// DatabasesTable holds the schema information for the "databases" table.
	DatabasesTable = &schema.Table{
		Name:       "databases",
		Columns:    DatabasesColumns,
		PrimaryKey: []*schema.Column{DatabasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "databases_groups_databases",
				Columns:    []*schema.Column{DatabasesColumns[14]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "database_organization_id",
				Unique:  true,
				Columns: []*schema.Column{DatabasesColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
			{
				Name:    "database_name",
				Unique:  true,
				Columns: []*schema.Column{DatabasesColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "primary_location", Type: field.TypeString},
		{Name: "locations", Type: field.TypeJSON, Nullable: true},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "region", Type: field.TypeEnum, Enums: []string{"AMER", "EMEA", "APAC"}, Default: "AMER"},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "group_name",
				Unique:  true,
				Columns: []*schema.Column{GroupsColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DatabasesTable,
		GroupsTable,
	}
)

func init() {
	DatabasesTable.ForeignKeys[0].RefTable = GroupsTable
}
