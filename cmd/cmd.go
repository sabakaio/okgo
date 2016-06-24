package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/sabakaio/okgo/registry"
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
