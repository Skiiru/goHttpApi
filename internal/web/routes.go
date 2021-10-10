package web

import (
	"github.com/gofiber/fiber/v2"
	"goHttpApi/internal/web/routes"
)

type Route struct {
	Path     string
	Method   string
	Callback func(context *fiber.Ctx) error
}

func GetRoutes() []Route {
	return []Route{
		{Path: "/api/v1/hello", Method: "Get", Callback: routes.Hello},
		{Path: "/api/v1/method", Method: "Post", Callback: routes.Method},
	}
}
