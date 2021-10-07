package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(context *fiber.Ctx) error {
	return context.SendString("Hello world")
}
