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

	for _, p := range list.Values {
		s := p.UserStatus(client.Username)
		if s == "UNAPPROVED" {
			term.Infof("%-15.15s %30s \"%s\"\n", p.Author.Mention(), p.Name(), p.Title)
		} else {
			term.Normalf("%-15.15s %30s \"%s\"\n", p.Author.Mention(), p.Name(), p.Title)
		}
		//term.Normalf("%#v\n", p)
	}

	return nil
}
