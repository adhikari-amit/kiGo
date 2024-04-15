package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

const PORT = ":8080"

func main() {
	app := fiber.New()

	app.Get("/api/list", func(c fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	app.Post("/api/register", func(c fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	log.Fatal(app.Listen(PORT))
}
