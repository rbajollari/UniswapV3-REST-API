package main

import (
	"net/http"
	"log"
	"github.com/rbajollari/UniswapV3-REST-API/api"
)

func main() {
	api := api.API{}

	api.Initialize()

	log.Fatal(http.ListenAndServe(":8080", api.Router))
}