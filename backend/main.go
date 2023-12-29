package main

import (
	"log"

	"github.com/m3phist/gobank/backend/api"
	"github.com/m3phist/gobank/backend/utils"
)

func main() {
	// method 1: using dotenv()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// portString := os.Getenv("PORT")
	// if portString == "" {
	// 	log.Fatal("PORT is not found in the environment variable")
	// }

	// port, err := strconv.Atoi(portString)
	// if err != nil {
	// 	log.Fatalf("Error converting PORT to integer: %v", err)
	// }

	// method 2: using viper
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	port := config.Port
	if err != nil {
		log.Fatalf("PORT is not found in the env variable: %v", err)
	}

	server := api.NewServer(".")
	server.Start(port)
}
