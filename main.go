package main

import (
	"os"

	"github.com/HungTP-Play/hunterxhunter/infrastructure/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectAndMigrate()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
