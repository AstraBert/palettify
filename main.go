package main

import (
	"log"

	"github.com/AstraBert/palettify/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := Setup()

	// Start the Fiber server on port 8000
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func Setup() *fiber.App {
	// Create a new Fiber app
	app := fiber.New()

	// Define a route for handling image uploads
	app.Post("/colors", handlers.ExtractColorsImage)
	app.Get("/", handlers.HomeRoute)
	app.Static("/", "./static/")

	return app
}
