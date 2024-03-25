package geodeticgroup

import (
	"github.com/spf13/cobra"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
)

// groupCmd represents the base group command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with geodetic groups",
}

func init() {
	geodetic.RootCmd.AddCommand(groupCmd)
}
