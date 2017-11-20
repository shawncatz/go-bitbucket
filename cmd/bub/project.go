package main

import (
	"errors"

	"github.com/urfave/cli"

	"github.com/shawncatz/go-bitbucket/bitbucket"
)

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:        "project",
		Usage:       "checkout every repo in a project",
		Description: `checkout every repo in a project`,
		Action:      cmdProject,
		ArgsUsage:   "<project>",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "a",
				Usage: "get all projects, even if greater than limit",
			},
			cli.IntFlag{
				Name:  "limit",
				Usage: "max number of projects to process",
				Value: 250,
			},
		},
	})
}

func cmdProject(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return errors.New("must specify project name")
	}

	args := c.Args()
	project := args[0]
	limit := c.Int("limit")

	list, err := client.Repos.List(project, &bitbucket.ReposListOptions{ListOptions: bitbucket.ListOptions{Limit: limit}})
	if err != nil {
		return err
	}

	if !list.IsLastPage {
		term.Warnf("warning: there are more repos than the given limit (%d)\n", limit)
	}

	for _, repo := range list.Values {
		term.Infof("* %s\n", repo.Name)
	}

	return nil
}
