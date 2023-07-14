package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mateus-rodriguess/contacts-api/pkg/database"
	"github.com/mateus-rodriguess/contacts-api/pkg/model"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error", "data": err})
	}

	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	} // Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User
	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}
