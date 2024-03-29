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
	Short: "get an existing geodetic group",
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
			s, err := json.Marshal(group.Group)
			if err != nil {
				return err
			}

			return geodetic.JSONPrint(s)
		}

		return geodetic.SingleRowTablePrint(group.Group)
	}

	groups, err := cli.Client.GetAllGroups(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err := json.Marshal(groups.Groups)
	if err != nil {
		return err
	}

	// print json output
	if viper.GetString("output.format") == "json" {
		return geodetic.JSONPrint(s)
	}

	// print table output
	var resp geodetic.GraphResponse

	err = json.Unmarshal(s, &resp)
	if err != nil {
		return err
	}

	return geodetic.RowsTablePrint(resp)
}
