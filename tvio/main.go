package main

import (
	"os"

	"github.com/ThingiverseIO/thingiverseio_tool/commands/config"
	"github.com/ThingiverseIO/thingiverseio_tool/commands/network"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tvio"
	app.Usage = "The swiss army knife for thingiverse.io"
	app.Commands = []cli.Command{
		config.Command,
		network.Command,
	}

	app.Run(os.Args)
}
