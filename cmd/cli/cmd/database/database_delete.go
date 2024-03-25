package geodeticdatabase

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
)

var databaseDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing geodetic database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseDeleteCmd)

	databaseDeleteCmd.Flags().StringP("name", "n", "", "database name to delete")
	geodetic.ViperBindFlag("database.delete.name", databaseDeleteCmd.Flags().Lookup("name"))
}

func deleteDatabase(ctx context.Context) error {
	// setup geodetic http client
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	dName := viper.GetString("database.delete.name")
	if dName == "" {
		return geodetic.NewRequiredFieldMissingError("name")
	}

	d, err := cli.Client.DeleteDatabase(ctx, dName, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(d.DeleteDatabase)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return geodetic.SingleRowTablePrint(d.DeleteDatabase)
}
