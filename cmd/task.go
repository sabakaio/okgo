package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"

	"github.com/evindor/okgo/models"
)

// CmdTask - task specific commands
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
			Name:  "create",
			Usage: "create a new task",
			Action: func(c *cli.Context) {
				task, err := models.CreateTask(c.String("name"), c.String("command"))
				if err != nil {
					println(err)
					return
				}
				fmt.Printf("Created task %v", task.Name)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: "default",
					Usage: "Task name",
				},
				cli.StringFlag{
					Name:  "command, c",
					Value: "echo \"ok\"",
					Usage: "Command to execute",
				},
			},
		},
		{
			Name:  "ls",
			Usage: "list all defined tasks",
			Action: func(c *cli.Context) {
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 0, 8, 0, '\t', 0)
				tasks, _ := models.ListTasks()
				fmt.Fprintln(w, "Name\tCommand")
				for _, task := range *tasks {
					fmt.Fprintf(w, "%v\t%v\n", task.Name, task.Command)
				}
				w.Flush()
			},
		},
		{
			Name:  "remove",
			Usage: "remove an existing task",
			Action: func(c *cli.Context) {
				println("removed task template: ", c.Args().First())
			},
		},
	},
}
