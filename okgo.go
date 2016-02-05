package main

import (
	"./cmd"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "okgo"
	app.Version = "0.1.0"
	app.Usage = "Scheduling service and API"
	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		cmd.CmdServer,
		cmd.CmdJob,
	}

	app.Run(os.Args)
}
