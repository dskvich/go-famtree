package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./client/build")
}

func main() {
	app := fiber.New()

	serveStatic(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app.Listen(":" + port)
}
