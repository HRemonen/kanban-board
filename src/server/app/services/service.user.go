package services

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetAllUsers() ([]model.APIUser, error) {
	db := database.DB.Db
	var users []model.APIUser

	db.Model(&model.User{}).Preload(clause.Associations).Find(&users)

	return users, nil
}

func GetSingleUser(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db
	id := c.Params("id")

	var user model.User
	db.Find(&user, "id = ?", id)

	if _, err := utils.IsAuthorized(c, user); err != nil {
		return user, err
	}

	return user, nil
}
