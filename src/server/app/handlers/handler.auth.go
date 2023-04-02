package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// Login ... Login user
// @Summary Login user
// @Description login user and generate JWT token
// @Tags Login
// @Param login_attrs body model.LoginUserInput true "Login attributes"
// @Success 201 {object} model.LoginData
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	loginData, err := services.Login(c)

	if err != nil && strings.Contains(err.Error(), "Invalid password") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Use Oauth login for Google login") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Use Oauth login for Google login", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged in successfully",
		"data":    loginData,
	})
}
