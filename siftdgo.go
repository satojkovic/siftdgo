package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "siftdgo"
	app.Version = Version
	app.Usage = "SIFT descriptor in Go "
	app.Author = "satojkovic"
	app.Email = "satojkovic@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
