package services

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
)

func GetSingleCard(c *fiber.Ctx) (model.Card, error) {
	db := database.DB.Db
	var card model.Card

	cardID := c.Params("id")

	err := db.Model(&model.Card{}).Find(&card, "id = ?", cardID).Error

	return card, err
}
