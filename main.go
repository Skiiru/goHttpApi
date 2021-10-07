package main

import (
	"first/internal/web"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

func main() {
	app := fiber.New()
	routes := web.GetRoutes()

	for _, route := range routes {
		reflect.ValueOf(app).MethodByName(route.Method).Call([]reflect.Value{
			reflect.ValueOf(route.Path),
			reflect.ValueOf(route.Callback),
		})
	}

	app.Listen(":8080")
}
