package main

import (
	"fmt"
	"log"

	"github.com/Goscord/DocGen/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func main() {
	app = fiber.New()

	err := config.Load()

	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Get().Port)))
}
