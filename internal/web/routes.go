package web

import (
	"first/internal/web/routes"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Path     string
	Method   string
	Callback func(context *fiber.Ctx) error
}

func GetRoutes() []Route {
	return []Route{
		{Path: "/api/v1/first", Method: "Get", Callback: routes.First},
	}
}
