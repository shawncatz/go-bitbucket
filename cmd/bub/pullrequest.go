package main

import (
	"github.com/urfave/cli"
)

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:   "pull-request",
		Usage:  "create a pull-request",
		Action: cmdPullRequest,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "message, m",
				Usage: "the `MESSAGE`",
			},
			cli.StringFlag{
				Name:  "base",
				Usage: "the `BASE` branch [default: 'master']",
			},
			cli.StringFlag{
				Name:  "head",
				Usage: "the `HEAD` branch [default: current branch]",
			},
		},
	})
}

func cmdPullRequest(c *cli.Context) error {
	return nil
}
