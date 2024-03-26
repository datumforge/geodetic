package geodeticgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
	"github.com/datumforge/geodetic/internal/ent/enums"
	"github.com/datumforge/geodetic/internal/geodeticclient"
)

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new geodetic group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringP("name", "n", "", "name of the group")
	geodetic.ViperBindFlag("group.create.name", groupCreateCmd.Flags().Lookup("name"))

	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
	geodetic.ViperBindFlag("group.create.description", groupCreateCmd.Flags().Lookup("description"))

	groupCreateCmd.Flags().StringP("region", "r", "", "region of the group (AMER, EMEA, APAC)")
	geodetic.ViperBindFlag("group.create.region", groupCreateCmd.Flags().Lookup("region"))

	groupCreateCmd.Flags().StringP("primary-location", "l", "", "primary location of the group")
	geodetic.ViperBindFlag("group.create.location", groupCreateCmd.Flags().Lookup("primary-location"))
}

func createGroup(ctx context.Context) error {
	cli, err := geodetic.GetGraphClient()
	if err != nil {
		return err
	}

	name := viper.GetString("group.create.name")
	if name == "" {
		return geodetic.NewRequiredFieldMissingError("name")
	}

	description := viper.GetString("group.create.description")
	location := viper.GetString("group.create.location")
	region := viper.GetString("group.create.region")

	input := geodeticclient.CreateGroupInput{
		Name:            name,
		Description:     &description,
		PrimaryLocation: location,
		Region:          enums.ToRegion(region),
	}

	g, err := cli.Client.CreateGroup(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(g.CreateGroup.Group)
		if err != nil {
			return err
		}

		return geodetic.JSONPrint(s)
	}

	return geodetic.SingleRowTablePrint(g.CreateGroup.Group)
}
