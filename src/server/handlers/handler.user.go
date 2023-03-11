package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	db.Omit("password").Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

func GetSingleUser(c *fiber.Ctx) error {
	user, err := utils.CheckAuthorization(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	userResponse := model.FilteredResponse(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": userResponse})
}

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

func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Name string `json:"name"`
	}
	db := database.DB.Db

	user, err := utils.CheckAuthorization(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	var updateUserData updateUser

	err = c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	user.Name = updateUserData.Name

	db.Save(&user)

	userResponse := model.FilteredResponse(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": userResponse})
}

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
