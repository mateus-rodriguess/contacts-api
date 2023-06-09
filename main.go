package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mateus-rodriguess/contacts-api/pkg/database"
	router "github.com/mateus-rodriguess/contacts-api/pkg/routes"
)

func main() {
	database.Connet()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":8080")

}
