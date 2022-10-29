package main

import (
	"log"

	"github.com/Goscord/DocGen/database"
	"github.com/Goscord/DocGen/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables :
	godotenv.Load()

	// Load database :
	err := database.Init()
	if err != nil {
		log.Fatalf("Cannot init database: %v", err)
	}

	// Load routes :
	routes.Init()
}
