package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"

	"github.com/datumforge/go-turso"

	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/ent/generated"
	"github.com/datumforge/geodetic/internal/ent/generated/hook"
)

func HookDatabase() ent.Hook {
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
				// create a turso db
				body := turso.CreateDatabaseRequest{
					Group: "default",
					Name:  name,
				}

				db, err := mutation.Turso.Database.CreateDatabase(ctx, body)
				if err != nil {
					return nil, err
				}

				mutation.Logger.Infow("created turso db", "db", db.Database.DatabaseID, "hostname", db.Database.Hostname)

				mutation.SetDsn(db.Database.Hostname)
			}

			mutation.SetStatus(enums.Active)

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
