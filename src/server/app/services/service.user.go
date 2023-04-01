package services

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"gorm.io/gorm/clause"
)

func GetAllUsers() ([]model.APIUser, error) {
	db := database.DB.Db
	var users []model.APIUser

	db.Model(&model.User{}).Preload(clause.Associations).Find(&users)

	return users, nil
}
