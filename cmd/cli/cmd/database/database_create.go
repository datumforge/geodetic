package geodeticdatabase

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/geodeticclient"
)

var databaseCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new geodetic database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseCreateCmd)

	databaseCreateCmd.Flags().StringP("org-id", "o", "", "owning organization id of the database")
	geodetic.ViperBindFlag("database.create.orgid", databaseCreateCmd.Flags().Lookup("org-id"))

	databaseCreateCmd.Flags().StringP("provider", "p", "turso", "provider of the database (local, turso)")
	geodetic.ViperBindFlag("database.create.provider", databaseCreateCmd.Flags().Lookup("provider"))

	databaseCreateCmd.Flags().StringP("group-id", "g", "", "group name to assign to the database")
	geodetic.ViperBindFlag("database.create.groupid", databaseCreateCmd.Flags().Lookup("group-id"))
}

func createDatabase(ctx context.Context) error {
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	orgID := viper.GetString("database.create.orgid")
	if orgID == "" {
		return geodetic.NewRequiredFieldMissingError("organization_id")
	}

	provider := viper.GetString("database.create.provider")
	if provider == "" {
		return geodetic.NewRequiredFieldMissingError("provider")
	}

	groupID := viper.GetString("database.create.groupid")
	if groupID == "" {
		return geodetic.NewRequiredFieldMissingError("group_id")
	}

	input := geodeticclient.CreateDatabaseInput{
		OrganizationID: orgID,
		Provider:       enums.ToDatabaseProvider(provider),
		GroupID:        groupID,
	}

	d, err := cli.Client.CreateDatabase(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(d)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return databaseTablePrint(d.CreateDatabase.Database)
}
