package main

import (
	"net/http"

	"api-boiler/router"
)

func init() {
	handler := router.Load()
	http.Handle("/", handler)
}
