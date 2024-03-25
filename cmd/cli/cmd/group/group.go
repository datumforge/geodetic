package geodeticgroup

import (
	"github.com/spf13/cobra"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
	"github.com/datumforge/geodetic/internal/geodeticclient"
)

// groupCmd represents the base group command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with geodetic groups",
}

func init() {
	geodetic.RootCmd.AddCommand(groupCmd)
}

func groupsTablePrint(g geodeticclient.GetAllGroups_Groups) error {
	// check if there are any groups, otherwise we have nothing to print
	if len(g.Edges) > 0 {
		groups := g.Edges
		// get the headers for the table for each struct and substruct
		header := geodetic.GetHeaders(groups[0].Node, "")

		data := [][]string{}

		// get the field values for each struct and substruct per row
		for _, g := range groups {
			fields := geodetic.GetFields(*g.Node)

			// append the fields to the data slice
			data = append(data, fields)
		}

		// print ze data
		return geodetic.TablePrint(header, data)
	}

	return nil
}

func groupTablePrint(g interface{}) error {
	// get the headers for the table for each struct and substruct
	header := geodetic.GetHeaders(g, "")

	data := [][]string{}

	// get the field values for each struct and substruct per row
	fields := geodetic.GetFields(g)

	// append the fields to the data slice
	data = append(data, fields)

	// print ze data
	return geodetic.TablePrint(header, data)
}
