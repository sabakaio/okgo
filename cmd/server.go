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
	Flags: []cli.Flag{
		stringFlag("port, p", "3000", "set port for HTTP API server"),
	},
}

func serverAction(c *cli.Context) {
	r := api.CreateServer()
	port := string(c.String("port"))
	r.Run(":" + port)
}
