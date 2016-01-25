package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdTask = cli.Command{
	Name:        "task",
	Usage:       "list tasks",
	Description: `start okgo server`,
	Action: func(c *cli.Context) {
		cli.ShowSubcommandHelp(c)
	},
	Flags: []cli.Flag{
		stringFlag("config, c", "custom/conf/app.yml", "Custom configuration file path"),
	},
	Subcommands: []cli.Command{
		{
			Name:  "add",
			Usage: "add a new template",
			Action: func(c *cli.Context) {
				println("new task template: ", c.Args().First())
			},
		},
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) {
				println("removed task template: ", c.Args().First())
			},
		},
	},
}
