package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/rbajollari/UniswapV3-REST-API/api"
)

func main() {
	api := api.API{}

	api.Initialize()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file, error: %v", err)
	}

	PORT := os.Getenv("API_PORT")

	log.Fatal(http.ListenAndServe(":" + PORT, api.Router))
}