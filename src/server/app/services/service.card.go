package services

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func GetSingleCard(c *fiber.Ctx) (model.Card, error) {
	db := database.DB.Db
	var card model.Card

	cardID := c.Params("id")

	err := db.Model(&model.Card{}).Find(&card, "id = ?", cardID).Error

	return card, err
}

func CreateListCard(c *fiber.Ctx) (model.Card, error) {
	db := database.DB.Db
	var list model.List
	var card model.Card

	listID := c.Params("id")
	payload := new(model.CardUserInput)

	err := db.Model(&list).Preload("Cards").Find(&list, "id = ?", listID).Error

	if err != nil {
		return card, errors.New("List not found")
	}

	c.BodyParser(payload)

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return card, err
	}

	var currentPosition uint

	db.Model(&model.Card{}).Select("COALESCE(MAX(position), 0)").Where("list_id = ?", list.ID).Row().Scan(&currentPosition)

	if len(list.Cards) == 0 {
		currentPosition = 0
	} else {
		currentPosition++
	}

	card = model.Card{
		Title:    payload.Title,
		Position: currentPosition,
		ListID:   list.ID,
	}

	err = db.Create(&card).Error

	return card, err
}
