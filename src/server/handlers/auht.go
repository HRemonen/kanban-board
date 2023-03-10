package handlers

import (
	"fmt"
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
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	var user model.User
	result := db.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"status": "fail", "message": "User not found, check username"})
	}

	fmt.Println(user.Password)

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"status": "fail", "message": "Invalid password"})
	}

	if user.Provider == "Google" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Use OAuth for google login"})
	}

	token, err := utils.GenerateNewAccessToken(user.ID)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged in successfully",
		"token":   token,
		"user":    user,
	})
}
