package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/tabwriter"

	"github.com/codegangsta/cli"
)

// CmdJob - job specific commands
var CmdJob = cli.Command{
	Name:        "job",
	Aliases:     []string{"j"},
	Usage:       "manage jobs - create, list, remove, run, purge",
	Description: `Subcommands set for managing jobs`,
	Action:      cli.ShowSubcommandHelp,
	Subcommands: []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new job",
			Action:  createAction,
			Flags: []cli.Flag{
				stringFlag("name, n", "default", "Job name"),
				stringFlag("command, c", "echo \"ok\"", "Command to execute"),
				stringFlag("schedule, s", "once", "Execution schedule (e.g. 'every 2 hours'"),
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list all defined jobs",
			Action:  listAction,
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove an existing job",
			Action:  removeAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run a job",
			Action:  runAction,
		},
		{
			Name:    "purge",
			Aliases: []string{"p"},
			Usage:   "remove all jobs",
			Action:  purgeAction,
		},
	},
}

func createAction(c *cli.Context) {
	job, err := jobs.Create(c.String("name"), c.String("command"), "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created job", job.Name)
}

func listAction(c *cli.Context) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	jobs, err := jobs.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Name\tCommand")
	for _, job := range jobs {
		fmt.Fprintf(w, "%v\t%v\n", job.Name, job.Command)
	}
	w.Flush()
}

func removeAction(c *cli.Context) {
	jobs.Delete(c.Args().First())
	println("removed job ", c.Args().First())
}

func purgeAction(c *cli.Context) {
	jobs.DeleteAll()
	println("removed job ", c.Args().First())
}

func runAction(c *cli.Context) {
	job, err := jobs.Get(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("/bin/bash", "-c", job.Command).CombinedOutput()
	if err != nil {
		fmt.Printf(string(out))
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}
