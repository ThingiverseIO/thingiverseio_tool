package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ThingiverseIO/thingiverseio/config"
	"github.com/ThingiverseIO/thingiverseio_tool/common"
	"github.com/codegangsta/cli"
)

var useGlobal bool

var SetCommand = cli.Command{
	Name:        "set",
	Usage:       "Sets the network interface to the address or interface number",
	Description: "Sets the interface the thingiverse.io network on this machine to the given address or interface number. If global flag is not set, the configuration is written to the user config file.",
	Action:      runSet,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "global, g",
			Usage:       "Writes the value to the global config file",
			Destination: &useGlobal,
		},
	},
}

func runSet(c *cli.Context) {
	iface := c.Args().Get(0)

	if iface == "" {
		common.PrintError(fmt.Errorf("Not enough Arguments."))
		return
	}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		common.PrintError(err)
		return
	}

	n, err := strconv.Atoi(iface)

	if err == nil {
		if n > len(addrs) {
			common.PrintError(fmt.Errorf("Invalid interface number %d, must be between 0 and %d", n, len(addrs)))
			return
		}

		iface = strings.Split(addrs[n].String(), "/")[0]

	} else {

		var valid bool

		for _, addr := range addrs {
			valid = strings.Contains(addr.String(), iface)
			if valid {
				break
			}
		}

		if !valid {
			common.PrintError(fmt.Errorf("Invalid interface address %s", iface))
			return
		}
	}

	if useGlobal {
		var cfgf config.CfgFile
		var err error
		if config.CfgFileGlobalPresent() {
			cfgf, err = config.ReadCfgFile(config.CfgFileGlobal())
			if err != nil {
				common.PrintError(err)
				return
			}
		}
		cfgf.Network.Interface = []string{iface}
		err = config.WriteCfgFile(cfgf, config.CfgFileGlobal())
		if err != nil {
			common.PrintError(err)
			return
		}
	} else {
		var cfgf config.CfgFile
		var err error
		if config.CfgFileUserPresent() {
			cfgf, err = config.ReadCfgFile(config.CfgFileUser())
			if err != nil {
				common.PrintError(err)
				return
			}
		}
		cfgf.Network.Interface = []string{iface}
		err = config.WriteCfgFile(cfgf, config.CfgFileUser())
		if err != nil {
			common.PrintError(err)
			return
		}
	}

}
