package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/go-turso"

	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/ent/generated"
	"github.com/datumforge/geodetic/internal/ent/generated/hook"
)

func HookCreateDatabase() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.DatabaseFunc(func(ctx context.Context, mutation *generated.DatabaseMutation) (generated.Value, error) {
			// get organization and provider from the request
			orgID, _ := mutation.OrganizationID()
			provider, _ := mutation.Provider()

			// create a name for the database
			name := strings.ToLower(fmt.Sprintf("org-%s", orgID))
			mutation.SetName(name)

			// if the provider is turso, create a database
			if provider == enums.Turso {
				// get the group to assign the database to
				groupID, _ := mutation.GroupID()

				group, err := mutation.Client().Group.Get(ctx, groupID)
				if err != nil {
					mutation.Logger.Errorw("unable to get group", "error", err)

					return nil, err
				}

				// create a turso db
				body := turso.CreateDatabaseRequest{
					Group: group.Name,
					Name:  name,
				}

				db, err := mutation.Turso.Database.CreateDatabase(ctx, body)
				if err != nil {
					return nil, err
				}

				mutation.Logger.Infow("created turso db", "db", db.Database.DatabaseID, "hostname", db.Database.Hostname)

				mutation.SetDsn(db.Database.Hostname)
			} else {
				// set the dsn to the name
				mutation.SetDsn(fmt.Sprintf("file:%s.db", name))
			}

			mutation.SetStatus(enums.Active)

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

func HookDatabaseDelete() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.DatabaseFunc(func(ctx context.Context, mutation *generated.DatabaseMutation) (generated.Value, error) {
			if ok := graphql.HasOperationContext(ctx); ok {
				// TODO: this only works for a delete database and not on a cascade delete
				gtx := graphql.GetOperationContext(ctx)
				name := gtx.Variables["name"].(string)

				if name == "" {
					mutation.Logger.Errorw("unable to delete database, no name provided")

					return nil, fmt.Errorf("no name provided") //nolint:goerr113
				}

				db, err := mutation.Turso.Database.DeleteDatabase(ctx, name)
				if err != nil {
					return nil, err
				}

				mutation.Logger.Infow("deleted turso database", "database", db.Database)
			}

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpDelete|ent.OpDeleteOne)
}
