package main

import (
	"os"

	// "github.com/jasonlvhit/gocron"
	// "github.com/boltdb/bolt"
	"github.com/codegangsta/cli"

	"github.com/evindor/okgo/cmd"
	"github.com/evindor/okgo/models"
)

func init() {
	// Register store to libkv
	models.CreateTask("more tasks", "docker run")
	models.CreateTask("try tasks", "ls -la")
}

func main() {
	app := cli.NewApp()
	app.Name = "okgo"
	app.Version = "0.0.1"
	app.Usage = "Scheduling service and API"
	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		cmd.CmdServer,
		cmd.CmdTask,
	}

	app.Run(os.Args)
}
