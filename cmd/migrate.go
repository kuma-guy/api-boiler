package cmd

import (
	"api-boiler/store"
	"log"

	"github.com/codegangsta/cli"
)

var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrates the store schema to the latest available version",
	Flags: []cli.Flag{
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
	Action: func(c *cli.Context) error {
		store, err := store.NewMySQLStore(
			c.String("mysql-username"),
			c.String("mysql-password"),
			c.String("mysql-host"),
			c.String("mysql-port"),
			c.String("mysql-dbname"),
		)

		if err != nil {
			log.Fatal("Failed to create Store: ", err)
		}

		if err := store.Ping(); err != nil {
			log.Fatal("Failed to connect with store: ", err)
		}

		if err := store.Migrate(); err != nil {
			log.Fatal("Failed to migrate store schema: ", err)
		}

		return nil
	},
}
