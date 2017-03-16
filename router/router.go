package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-boiler/server"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middleware...)

	e.GET("/api/hello", server.GetHello)

	return e
}
