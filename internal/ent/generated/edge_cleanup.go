// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/datumforge/geodetic/internal/ent/generated/database"
	"github.com/datumforge/geodetic/internal/ent/generated/group"
)

func DatabaseEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).Database.Query().Where((database.HasGroupWith(group.ID(id)))).Exist(ctx); err == nil && exists {
		if databaseCount, err := FromContext(ctx).Database.Delete().Where(database.HasGroupWith(group.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting database", "count", databaseCount, "err", err)
			return err
		}
	}

	return nil
}
