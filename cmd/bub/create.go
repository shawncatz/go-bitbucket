package main

import (
	"github.com/urfave/cli"
)

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:  "create",
		Usage: "create a pull request",
		Description: `Of the form 'PROJECT/REPO'
	 If omitted, this will assume the PROJECT is the name of the parent
	 directory, and the REPO is the name of the current working directory.
	 If NAME does not contain a '/' it will override the name of the REPO`,
		Action:    cmdCreate,
		ArgsUsage: "[[PROJECT/]REPO]",
	})
}

func cmdCreate(c *cli.Context) error {
	return nil
}
