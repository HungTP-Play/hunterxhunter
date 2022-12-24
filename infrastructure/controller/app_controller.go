package controller

import (
	"github.com/HungTP-Play/hunterxhunter/app/characters"
	"github.com/gofiber/fiber/v2"
)

type IController interface {
	ApplyController(r fiber.Router)
}

func ApplyController(r fiber.Router) {
	// Characters
	characterController := characters.CharacterController{}
	characterController.ApplyController(r)
}
