package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./web/build")
}

func main() {
	app := fiber.New()

	serveStatic(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)
}
