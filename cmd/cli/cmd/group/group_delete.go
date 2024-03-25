package geodeticgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
)

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing geodetic group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringP("name", "n", "", "group name to delete")
	geodetic.ViperBindFlag("group.delete.name", groupDeleteCmd.Flags().Lookup("name"))
}

func deleteGroup(ctx context.Context) error {
	// setup geodetic http client
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	gName := viper.GetString("group.delete.name")
	if gName == "" {
		return geodetic.NewRequiredFieldMissingError("name")
	}

	g, err := cli.Client.DeleteGroup(ctx, gName, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(g)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return groupTablePrint(g.DeleteGroup)
}
