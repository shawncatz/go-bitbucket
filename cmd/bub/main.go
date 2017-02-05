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
	"os"

	"github.com/urfave/cli"
)

var (
	cmdList = []cli.Command{} // so that the individual files can append to the list
)

func main() {
	app := cli.NewApp()
	app.Name = "bub"
	app.Version = "0.1.0"
	app.Author = "Shawn Catanzarite"
	app.Usage = "adds BitBucket functions to git command"
	app.UsageText = "bub [global options] <command> [command options] [arguments...]"
	app.Email = "me@shawncatz.com"
	app.Description = "pass-through git replacement, adds BitBucket functionality to git command"
	app.Commands = cmdList

	app.Run(os.Args)
}
