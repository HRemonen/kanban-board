package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
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
// @Failure 500 {object} object
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	db := database.DB.Db
	payload := new(model.LoginUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
	}

	var user model.User

	result := db.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "fail", "message": "User not found, check username", "data": nil})
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

	userResponse := model.FilteredResponse(&user)

	var data model.LoginData
	data.Token = token
	data.User = userResponse

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged in successfully",
		"data":    data,
	})
}
