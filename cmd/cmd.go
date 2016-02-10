package cmd

import (
	"../registry"
	"github.com/codegangsta/cli"
)

var jobs *registry.JobsRegistry

func init() {
	jobs = registry.NewJobsRegistry()
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
