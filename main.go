package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func main() {
	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/:name?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("%s ✋", c.Params("name", "y a rien :("))
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))
}
