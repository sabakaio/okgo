package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/tabwriter"

	"github.com/codegangsta/cli"

	"github.com/evindor/okgo/models"
)

// CmdTask - task specific commands
var CmdTask = cli.Command{
	Name:        "task",
	Usage:       "list tasks",
	Description: `start okgo server`,
	Action:      cli.ShowSubcommandHelp,
	Flags: []cli.Flag{
		stringFlag("config, c", "custom/conf/app.yml", "Custom configuration file path"),
	},
	Subcommands: []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new task",
			Action:  createAction,
			Flags: []cli.Flag{
				stringFlag("name, n", "default", "Task name"),
				stringFlag("command, c", "echo \"ok\"", "Command to execute"),
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list all defined tasks",
			Action:  listAction,
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove an existing task",
			Action:  removeAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run a task",
			Action:  runAction,
		},
		{
			Name:   "purge",
			Usage:  "remove all tasks",
			Action: purgeAction,
		},
	},
}

func createAction(c *cli.Context) {
	task, err := models.CreateTask(c.String("name"), c.String("command"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created task", task.Name)
}

func listAction(c *cli.Context) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	tasks, err := models.ListTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Name\tCommand")
	for _, task := range *tasks {
		fmt.Fprintf(w, "%v\t%v\n", task.Name, task.Command)
	}
	w.Flush()
}

func removeAction(c *cli.Context) {
	models.RemoveTask(c.Args().First())
	println("removed task ", c.Args().First())
}

func purgeAction(c *cli.Context) {
	models.PurgeTasks()
	println("removed task ", c.Args().First())
}

func runAction(c *cli.Context) {
	task, err := models.GetTask(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("/bin/bash", "-c", task.Command).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}
