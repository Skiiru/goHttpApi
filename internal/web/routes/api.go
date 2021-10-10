package routes

import (
	"context"
	"github.com/Skiiru/goGrpcApi/pkg/api"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func Hello(context *fiber.Ctx) error {
	return context.SendString("Hello world")
}

func Method(ctx *fiber.Ctx) error {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	person := new(Person)
	if err := ctx.BodyParser(person); err != nil {
		return err
	}

	client := api.NewApiServiceClient(conn)
	response, err := client.Method(context.Background(), &api.MethodRequest{Message: person.Name})
	if err != nil {
		panic(err)
	}

	return ctx.SendString(response.Result)
}
