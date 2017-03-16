package cmd

import (
	"net/http"

	"api-boiler/router"
	"api-boiler/router/middleware"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Server = cli.Command{
	Name:  "server",
	Usage: "starts the server daemon",
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
		cli.StringFlag{
			EnvVar: "MYSQL_HOST",
			Name:   "mysql-host",
			Usage:  "MySQL host",
			Value:  "localhost",
		},
		cli.StringFlag{
			EnvVar: "MYSQL_PORT",
			Name:   "mysql-port",
			Usage:  "MySQL port",
			Value:  "3306",
		},
		cli.StringFlag{
			EnvVar: "MYSQL_USERNAME",
			Name:   "mysql-username",
			Usage:  "MySQL username",
			Value:  "root",
		},
		cli.StringFlag{
			EnvVar: "MYSQL_PASSWORD",
			Name:   "mysql-password",
			Usage:  "MySQL password",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "MYSQL_DBNAME",
			Name:   "mysql-dbname",
			Usage:  "MySQL dbname",
			Value:  "api-boiler",
		},
	},
	Action: func(c *cli.Context) {
		if err := server(c); err != nil {
			logrus.Fatal(err)
		}
	},
}

func server(c *cli.Context) error {
	handler := router.Load(
		middleware.Store(c),
	)
	return http.ListenAndServe(c.String("server-addr"), handler)
}
