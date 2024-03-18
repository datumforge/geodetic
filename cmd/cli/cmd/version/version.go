package version

import (
	"github.com/spf13/cobra"

	"github.com/datumforge/datum/pkg/utils/cli/useragent"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
	"github.com/datumforge/geodetic/internal/constants"
)

// VersionCmd is the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print geodetic CLI version",
	Long:  `The version command prints the version of the geodetic CLI`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println(constants.VerboseCLIVersion)
		cmd.Printf("User Agent: %s\n", useragent.GetUserAgent())
	},
}

func init() {
	geodetic.RootCmd.AddCommand(versionCmd)
}
