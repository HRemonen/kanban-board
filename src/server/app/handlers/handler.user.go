package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} model.User
// @Failure 404 {object} object
// @Router /user [get]
func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong accessing data", "data": nil})
	}

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

// GetSingleUser ... Get a single user by ID
// @Summary Get a single user by ID
// @Description get a single user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} model.UserResponse
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /user/{id} [get]
func GetSingleUser(c *fiber.Ctx) error {
	user, err := services.GetSingleUser(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// CreateUser ... Create User
// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user_attrs body model.RegisterUserInput true "User attributes"
// @Success 201 {object} model.UserResponse
// @Failure 409 {object} object
// @Failure 422 {object} object
// @Failure 500 {object} object
// @Router /user [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	payload := new(model.RegisterUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match", "data": nil})
	}

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
	}

	hash, _ := utils.HashPassword(payload.Password)

	newUser := model.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hash,
	}

	err = db.Create(&newUser).Error

	if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		return c.Status(409).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists", "data": nil})

	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": nil})
	}

	userResponse := model.FilteredResponse(&newUser)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": userResponse})
}

// UpdateUser ... Update user name by ID
// @Summary Update user name by ID
// @Description update user name by ID
// @Tags Users
// @Accept json
// @Param id path string true "User ID"
// @Param name body model.UpdateUser true "User name"
// @Success 200 {object} model.UserResponse
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Failure 500 {object} object
// @Router /user/{id} [put]
func UpdateUser(c *fiber.Ctx) error {

	db := database.DB.Db

	user, err := utils.CheckAuthorization(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	var updateUserData model.UpdateUser

	err = c.BodyParser(&updateUserData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	var validate = utils.NewValidator()

	err = validate.Struct(updateUserData)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
	}

	user.Name = updateUserData.Name

	db.Save(&user)

	userResponse := model.FilteredResponse(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User updated", "data": userResponse})
}

// UpdateUser ... Delete user by ID
// @Summary Delete user by ID
// @Description delete user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Router /user/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db

	user, err := utils.CheckAuthorization(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	err = db.Delete(&user, "id = ?", user.ID).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted", "data": nil})
}
