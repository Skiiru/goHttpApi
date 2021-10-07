package routes

import (
	"github.com/gofiber/fiber/v2"
)

func First(context *fiber.Ctx) error {
	return context.SendString("Hello world")
}
