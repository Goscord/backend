package main

import (
	"fmt"
	"log"

	"github.com/Goscord/DocGen/config"
	"github.com/Goscord/DocGen/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	err = database.Init()
	if err != nil {
		log.Fatalf("Cannot init database: %v", err)
	}

	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Get().Port)))
}
