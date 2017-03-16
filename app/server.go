package main

import (
	"net/http"

	"api-boiler/router"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var serverCmd = cli.Command{
	Name:  "server",
	Usage: "starts the server daemon",
	Action: func(c *cli.Context) {
		if err := server(c); err != nil {
			logrus.Fatal(err)
		}
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "start the server in debug mode",
		},
		cli.StringFlag{
			EnvVar: "SERVER_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":8080",
		},
	},
}

func server(c *cli.Context) error {
	handler := router.Load()
	return http.ListenAndServe(c.String("server-addr"), handler)
}
