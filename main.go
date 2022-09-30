package main

import (
	"log"

	"github.com/Goscord/DocGen/config"
	"github.com/Goscord/DocGen/database"
	"github.com/Goscord/DocGen/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	err = database.Init()
	if err != nil {
		log.Fatalf("Cannot init database: %v", err)
	}

	routes.Init()
}
