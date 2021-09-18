package main

import (
	"github.com/gofiber/fiber/v2"

	"os"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./build")
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	app := fiber.New()

	serveStatic(app)

	app.Get("/api/users", func(c *fiber.Ctx) error {
		res := []User{
			{"John", "Doe"},
			{"Mary", "Jane"},
		}

		return c.Status(fiber.StatusOK).JSON(res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)
}
