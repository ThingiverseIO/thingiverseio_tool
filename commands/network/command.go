package network

import "github.com/codegangsta/cli"

var Command = cli.Command{
	Name:        "network",
	Aliases:     []string{"net"},
	Usage:       "Tools for configurating thingiverse.io networking",
	Description: "Provides tools to display available network interfaces and set the interface to use.",
	Subcommands: []cli.Command{
		ShowCommand,
		SetCommand,
	},
}
