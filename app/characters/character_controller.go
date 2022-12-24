package characters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CharacterController struct{}

func (c *CharacterController) ApplyController(r fiber.Router) {
	r.Get("/characters", handleGetAllCharacters)
	fmt.Println("CharacterController: /characters - GET")
}

func handleGetAllCharacters(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
