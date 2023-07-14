package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mateus-rodriguess/contacts-api/pkg/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Get("/", handler.GetAllUsers)
	v1.Post("/", handler.CreateUser)
}
