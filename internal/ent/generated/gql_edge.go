// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (d *Database) Group(ctx context.Context) (*Group, error) {
	result, err := d.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryGroup().Only(ctx)
	}
	return result, err
}

func (gr *Group) Databases(ctx context.Context) (result []*Database, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedDatabases(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.DatabasesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryDatabases().All(ctx)
	}
	return result, err
}
