package main

import (
	geodetic "github.com/datumforge/geodetic/cmd/cli/cmd"

	// since the cmds are no longer part of the same package
	// they must all be imported in main
	_ "github.com/datumforge/geodetic/cmd/cli/cmd/version"
)

func main() {
	geodetic.Execute()
}
