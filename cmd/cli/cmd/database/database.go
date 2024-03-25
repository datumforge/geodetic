package geodeticdatabase

import (
	"github.com/spf13/cobra"

	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"
	"github.com/datumforge/geodetic/internal/geodeticclient"
)

// databaseCmd represents the base database command when called without any subcommands
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "The subcommands for working with geodetic databases",
}

func init() {
	geodetic.RootCmd.AddCommand(databaseCmd)
}

func databasesTablePrint(d geodeticclient.GetAllDatabases_Databases) error {
	// check if there are any database, otherwise we have nothing to print
	if len(d.Edges) > 0 {
		dbs := d.Edges

		// get the headers for the table for each struct and substruct
		header := geodetic.GetHeaders(dbs[0].Node, "")

		data := [][]string{}

		// get the field values for each struct and substruct per row
		for _, g := range dbs {
			fields := geodetic.GetFields(*g.Node)

			// append the fields to the data slice
			data = append(data, fields)
		}

		// print ze data
		return geodetic.TablePrint(header, data)
	}

	return nil
}

func databaseTablePrint(d interface{}) error {
	// get the headers for the table for each struct and substruct
	header := geodetic.GetHeaders(d, "")

	data := [][]string{}

	// get the field values for each struct and substruct per row
	fields := geodetic.GetFields(d)

	// append the fields to the data slice
	data = append(data, fields)

	// print ze data
	return geodetic.TablePrint(header, data)
}
