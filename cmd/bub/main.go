/* bub - BitBucket / Stash version of Hub

   Add the following functionality to the git command.

   * pull-request
   * create
   * browse
   * compare
   * fork ???
   * ci-status ???

   All other commands pass-through to the git command normally.

*/
package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/urfave/cli"

	"github.com/shawncatz/go-bitbucket/bitbucket"
)

var (
	cmdList                   = []cli.Command{} // so that the individual files can append to the list
	client  *bitbucket.Client = nil
	term    *Term             = nil
)

func main() {
	args := os.Args

	if len(args) < 2 || supported(args[1]) || args[1][0] == '-' {
		bub(args)
		return
	}

	git(args)
}

func git(args []string) {
	err := syscall.Exec("git", args[1:], os.Environ())
	if err != nil {
		fmt.Printf("git error: %s\n", err)
	}
}

func supported(cmd string) bool {
	for _, c := range cmdList {
		if c.Name == cmd {
			return true
		}
	}

	return false
}

func bub(args []string) {
	// setup client
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	client = bitbucket.NewClient(cfg.URL, cfg.User, cfg.password)
	term = NewTerm()

	app := cli.NewApp()
	app.Name = "bub"
	app.Version = "0.1.0"
	app.Author = "Shawn Catanzarite"
	app.Usage = "adds BitBucket functions to git command"
	app.UsageText = "bub [global options] <command> [command options] [arguments...]"
	app.Email = "me@shawncatz.com"
	app.Description = "pass-through git replacement, adds BitBucket functionality to git command"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "no-color",
			Usage: "disable color output",
		},
	}
	app.Commands = cmdList

	app.Run(args)
}
