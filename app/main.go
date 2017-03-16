package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ianschenck/envflag"
)

func main() {
	envflag.Parse()

	app := cli.NewApp()
	app.Usage = "command line utility"

	app.Commands = []cli.Command{
		serverCmd,
	}

	app.Run(os.Args)
}
