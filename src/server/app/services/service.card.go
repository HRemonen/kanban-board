package services

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func UpdateListCardPosition(c *fiber.Ctx) (model.Card, error) {
	db := database.DB.Db
	var list model.List
	var card model.Card

	listID := c.Params("id")
	payload := new(model.CardPositionInput)

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

	cardID := c.Params("card")

	err = db.Find(&card, "id = ?", cardID).Error

	if err != nil {
		return card, errors.New("Card not found")
	}

	if list.ID != card.ListID {
		return card, errors.New("Unauthorized action")
	}

	currentPosition := card.Position

	if payload.Position == currentPosition {
		return card, errors.New("Card position not modified") // nothing to update
	}

	if payload.Position < currentPosition {
		// shift items between new and old position up by 1
		err = db.Model(&model.Card{}).Where("list_id = ? AND position between ? and ?", list.ID, payload.Position, currentPosition).Update("position", gorm.Expr("position + 1")).Error
	} else {
		// shift items between new and old position down by 1
		err = db.Model(&model.Card{}).Where("list_id = ? AND position between ? and ?", list.ID, currentPosition, payload.Position).Update("position", gorm.Expr("position - 1")).Error
	}

	if err != nil {
		return card, errors.New("Could not update list positions")
	}

	card.Position = payload.Position

	err = db.Save(&card).Error

	if err != nil {
		// rollback position update on error
		db.Model(&card).Where("position = ?", currentPosition).Update("position", payload.Position)
	}

	return card, err
}

func DeleteListCard(c *fiber.Ctx) error {
	db := database.DB.Db
	var list model.List
	var card model.Card

	listID := c.Params("id")

	err := db.Model(&list).Preload("Cards").Find(&list, "id = ?", listID).Error

	if err != nil {
		return errors.New("List not found")
	}

	cardID := c.Params("card")

	err = db.Find(&card, "id = ?", cardID).Error

	if err != nil {
		return errors.New("Card not found")
	}

	if list.ID != card.ListID {
		return errors.New("Unauthorized action")
	}

	err = db.Select(clause.Associations).Delete(&card).Error

	if err != nil {
		return errors.New("Failed to delete card")
	}

	err = db.Model(&model.Card{}).Where("list_id = ? AND position > ?", list.ID, card.Position).Update("position", gorm.Expr("position - 1")).Error

	return err
}
