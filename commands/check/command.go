package check

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ThingiverseIO/thingiverseio"
	"github.com/ThingiverseIO/thingiverseio_tool/common"
	"github.com/codegangsta/cli"
)

var Command = cli.Command{
	Name:        "check",
	Usage:       "Checks a given descriptor. Argument can be path or string.",
	Description: "Checks a descriptor file or string.",
	Action:      runCheck,
}

func runCheck(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "check")
		return
	}
	desc := c.Args().Get(0)
	if isFile(desc) {
		b, err := ioutil.ReadFile(desc)
		if err != nil {
			common.PrintError(err)
			return
		}
		desc = string(b)
	}
	
	fmt.Printf("Checking Descriptor:\n\n%s\n\n",desc)

	_, err := thingiverseio.ParseDescriptor(desc)

	result := "OK"

	if err != nil {
		result = err.Error()
	}

	fmt.Println(result)
}

func isFile(fp string) bool {
	if _, err := os.Stat(fp); err == nil {
		return true
	}
	return false
}
