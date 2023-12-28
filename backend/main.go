package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/m3phist/gobank/backend/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment variable")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("Error converting PORT to integer: %v", err)
	}

	api.NewServer(port)
}
