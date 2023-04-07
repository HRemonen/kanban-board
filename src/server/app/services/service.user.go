package services

import (
	"errors"
	"strings"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetAllUsers() ([]model.APIUser, error) {
	db := database.DB.Db
	var users []model.APIUser

	err := db.Model(&model.User{}).Preload(clause.Associations).Find(&users).Error

	return users, err
}

func GetSingleUser(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db
	id := c.Params("id")

	var user model.User

	err := db.Model(&model.User{}).Preload("Boards").Find(&user, "id = ?", id).Error

	if _, err := utils.IsAuthorized(c, user); err != nil {
		return user, err
	}

	return user, err
}

func CreateUser(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db
	payload := new(model.RegisterUserInput)

	c.BodyParser(payload)

	hash, _ := utils.HashPassword(payload.Password)

	user := model.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hash,
	}

	if payload.Password != payload.PasswordConfirm {
		return user, errors.New("Passwords do not match")
	}

	var validate = utils.NewValidator()

	err := validate.Struct(payload)

	if err != nil {
		return user, err
	}

	err = db.Create(&user).Error

	return user, err
}

func UpdateUser(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db
	id := c.Params("id")
	payload := new(model.UpdateUser)

	var user model.User
	db.Find(&user, "id = ?", id)

	if _, err := utils.IsAuthorized(c, user); err != nil {
		return user, err
	}

	c.BodyParser(payload)

	var validate = utils.NewValidator()

	err := validate.Struct(payload)

	if err != nil {
		return user, err
	}

	user.Name = payload.Name

	err = db.Save(&user).Error

	return user, err
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var user model.User
	db.Find(&user, "id = ?", id)

	if _, err := utils.IsAuthorized(c, user); err != nil {
		return err
	}

	err := db.Delete(&user, "id = ?", user.ID).Error

	return err
}
