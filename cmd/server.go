package cmd

import (
	"github.com/codegangsta/cli"

	"github.com/evindor/okgo/api"
)

// CmdServer - start a server
var CmdServer = cli.Command{
	Name:        "server",
	Usage:       "start okgo server",
	Description: `start okgo server`,
	Action:      serverAction,
	Aliases:     []string{"s"},
}

func serverAction(c *cli.Context) {
	r := api.CreateServer()
	r.Run(":3000")
	println("Started server on port 3000")
}
