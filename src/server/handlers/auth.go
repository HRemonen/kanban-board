package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	db := database.DB.Db
	payload := new(model.LoginUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	var user model.User
	result := db.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"status": "fail", "message": "User not found, check username", "data": nil})
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"status": "fail", "message": "Invalid password", "data": nil})
	}

	if user.Provider == "Google" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Use OAuth for google login", "data": nil})
	}

	token, err := utils.GenerateNewAccessToken(user.ID)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	var data model.LoginData
	data.Token = token
	data.User = user

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged in successfully",
		"data":    data,
	})
}
