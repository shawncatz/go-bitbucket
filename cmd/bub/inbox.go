package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:        "inbox",
		Usage:       "fetch inbox information",
		Description: `get list of pull requests on which you are a reviewer`,
		Action:      cmdInbox,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "c",
				Usage: "just get the count",
			},
		},
	})
}

func cmdInbox(c *cli.Context) error {
	if c.Bool("c") {
		count, err := client.Inbox.Count()
		if err != nil {
			return err
		}
		fmt.Printf("%d\n", count)
		return nil
	}

	list, err := client.Inbox.List()
	if err != nil {
		return err
	}

	f := "%-15.15s %30s %-70.70s\n"
	for _, p := range list.Values {
		s := p.UserStatus(client.Username)
		switch s {
		case "UNAPPROVED":
			term.Infof(f, p.Author.Mention(), p.Name(), p.Title)
		case "NEEDS_WORK":
			term.Warnf(f, p.Author.Mention(), p.Name(), p.Title)
		default:
			term.Normalf(f, p.Author.Mention(), p.Name(), p.Title)
		}
	}

	return nil
}
