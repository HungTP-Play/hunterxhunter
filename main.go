package main

import (
	"os"

	"github.com/HungTP-Play/hunterxhunter/infrastructure/controller"
	"github.com/HungTP-Play/hunterxhunter/infrastructure/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectAndMigrate()
	app := fiber.New()

	apiv1 := app.Group("/api/v1")

	controller.ApplyController(apiv1)

	app.Listen(":" + os.Getenv("PORT"))
}
