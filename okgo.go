package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/evindor/okgo/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "okgo"
	app.Version = "0.1.0"
	app.Usage = "Scheduling service and API"
	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		cmd.CmdServer,
		cmd.CmdTask,
	}

	app.Run(os.Args)
}
