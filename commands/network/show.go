package network

import (
	"fmt"
	"net"
	"strings"

	"github.com/ThingiverseIO/thingiverseio_tool/common"
	"github.com/codegangsta/cli"
)

var ShowCommand = cli.Command{
	Name:        "show",
	Usage:       "Shows available network interfaces",
	Description: "Shows the current configuration of the thingiverse.io network on this machine",
	Action:      runShow,
}

func runShow(c *cli.Context) {
	addr, err := net.InterfaceAddrs()

	if err != nil {
		common.PrintError(err)
		return
	}

	fmt.Println("")
	fmt.Println("Available Interfaces")
	fmt.Println("")

	for n, addr := range addr {
		paddr := strings.Split(addr.String(), "/")[0]
		fmt.Printf("\t[%d] %s\n", n+1, paddr)
	}

}
