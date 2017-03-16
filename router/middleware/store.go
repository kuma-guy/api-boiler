package middleware

import (
	"api-boiler/store"
	"log"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

func Store(cli *cli.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		store, err := store.NewMySQLStore(
			cli.String("mysql-username"),
			cli.String("mysql-password"),
			cli.String("mysql-host"),
			cli.String("mysql-port"),
			cli.String("mysql-dbname"),
		)

		if err != nil {
			log.Fatal("Failed to create Store: ", err)
		}

		c.Set("Store", store)
		c.Next()
	}
}
