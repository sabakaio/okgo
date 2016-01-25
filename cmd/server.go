package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdServer = cli.Command{
	Name:        "server",
	Usage:       "start okgo server",
	Description: `start okgo server`,
	Action:      runServ,
	Flags: []cli.Flag{
		stringFlag("config, c", "custom/conf/app.yml", "Custom configuration file path"),
	},
}

func runServ(c *cli.Context) {
	println("Started server on port 3000")
}
