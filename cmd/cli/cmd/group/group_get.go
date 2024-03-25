package geodeticgroup

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new geodetic group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringP("name", "n", "", "group name to query")
	geodetic.ViperBindFlag("group.get.name", groupGetCmd.Flags().Lookup("name"))
}

func getGroup(ctx context.Context) error {
	// setup geodetic http client
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	if cli.Client == nil {
		log.Fatal("client is nil")
	}

	// filter options
	gName := viper.GetString("group.get.name")

	// if an group name is provided, filter on that group, otherwise get all
	if gName != "" {
		group, err := cli.Client.GetGroup(ctx, gName, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(group)
			if err != nil {
				return err
			}

			return geodetic.JSONPrint(s)
		}

		return groupTablePrint((*group).Group)
	}

	groups, err := cli.Client.GetAllGroups(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(groups)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return groupsTablePrint(groups.Groups)
}
