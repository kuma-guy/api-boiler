package main

import (
	"api-boiler/cmd"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ianschenck/envflag"
)

func main() {
	envflag.Parse()

	app := cli.NewApp()
	app.Usage = "command line utility"

	app.Commands = []cli.Command{
		cmd.Server,
		cmd.Migrate,
	}

	app.Run(os.Args)
}
