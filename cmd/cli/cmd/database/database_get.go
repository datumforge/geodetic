package geodeticdatabase

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
)

var databaseGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new geodetic database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseGetCmd)

	databaseGetCmd.Flags().StringP("name", "n", "", "database name to query")
	geodetic.ViperBindFlag("database.get.name", databaseGetCmd.Flags().Lookup("name"))
}

func getDatabase(ctx context.Context) error {
	// setup geodetic http client
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	if cli.Client == nil {
		log.Fatal("client is nil")
	}

	// filter options
	dName := viper.GetString("database.get.name")

	// if an db name is provided, filter on that db, otherwise get all
	if dName != "" {
		db, err := cli.Client.GetDatabase(ctx, dName, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(db)
			if err != nil {
				return err
			}

			return geodetic.JSONPrint(s)
		}

		return databaseTablePrint((*db).Database)
	}

	dbs, err := cli.Client.GetAllDatabases(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(dbs)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return databasesTablePrint(dbs.Databases)
}
