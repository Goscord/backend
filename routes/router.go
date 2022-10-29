package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/Goscord/DocGen/routes/webhook"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func Init() {
	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// /webhook
	app.Post("/webhook", webhook.POSTHandler)

	port := os.Getenv("PORT")

	if port != "" {
		log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
	} else {
		log.Fatal(app.Listen(fmt.Sprintf(":%d", 3001)))
	}
}
