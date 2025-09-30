package main

import (
	"log"
	"time"

	"github.com/AstraBert/palettify/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/sqlite3"
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
	corsConfig := cors.Config{
		AllowOrigins: "https://palettify.nl",
		AllowMethods: "POST",
	}

	storage := sqlite3.New()
	limiterConfig := limiter.Config{
		Max: 10000,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Track limit per IP address
		},
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"message": "Service is currently unavailable due to server overload, retry soon!",
			})
		},
		Storage: storage,
	}

	app.Post("/html/colors", cors.New(corsConfig), handlers.ExtractColorsImage)
	app.Post("/json/colors", limiter.New(limiterConfig), handlers.ExtractColorsJSON)
	app.Get("/", handlers.HomeRoute)
	app.Static("/", "./static/")

	return app
}
