package handlers

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetSingleCard ... Get a single card by ID
// @Summary Get a single card by ID
// @Description get a single card by ID
// @Tags Cards
// @Param id path string true "Card ID"
// @Success 200 {object} model.Card
// @Failure 404 {object} object
// @Router /card/{id} [get]
func GetSingleCard(c *fiber.Ctx) error {
	db := database.DB.Db
	var card model.Card

	cardID := c.Params("id")

	db.Model(&model.Card{}).Find(&card, "id = ?", cardID)

	if card.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Card not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card found", "data": card})
}

// CreateListCard ... Create a new card for a list
// @Summary Create a new card for a list
// @Description create a new card for a list
// @Tags cards
// @Accept json
// @Param id path string true "List ID"
// @Param card_attrs body model.CardUserInput true "Card attributes"
// @Success 201 {object} model.Card
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Router /list/{id}/card [post]
func CreateListCard(c *fiber.Ctx) error {
	db := database.DB.Db
	var list model.List

	listID := c.Params("id")

	db.Model(&list).Preload("Cards").Find(&list, "id = ?", listID)

	if list.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "List not found", "data": nil})
	}

	payload := new(model.CardUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": nil})
	}

	var currentPosition uint

	db.Model(&model.Card{}).Select("COALESCE(MAX(position), 0)").Where("list_id = ?", list.ID).Row().Scan(&currentPosition)

	if len(list.Cards) == 0 {
		currentPosition = 0
	} else {
		currentPosition++
	}

	newCard := model.Card{
		Title:    payload.Title,
		Position: currentPosition,
		ListID:   list.ID,
	}

	err = db.Create(&newCard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a card for the list", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "A new card has been created", "data": newCard})
}

// UpdateListCardPosition ... Update card position on the list
// @Summary Update card position on the list
// @Description update card position on the list
// @Tags Cards
// @Accept json
// @Param id path string true "List ID"
// @Param card path string true "Card ID"
// @Param position body model.CardPositionInput true "Card position"
// @Success 200 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Failure 500 {object} object
// @Router /list/{id}/card/{card} [put]
func UpdateListCardPosition(c *fiber.Ctx) error {
	db := database.DB.Db
	var list model.List

	listID := c.Params("id")

	db.Find(&list, "id = ?", listID)

	if list.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "List not found", "data": nil})
	}

	payload := new(model.CardPositionInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": nil})
	}

	var card model.Card

	cardID := c.Params("card")

	db.Find(&card, "id = ?", cardID)

	currentPosition := card.Position

	if payload.Position == currentPosition {
		return c.Status(304).JSON(fiber.Map{"status": "success", "message": "Position not modified", "data": nil}) // nothing to update
	}

	if payload.Position < currentPosition {
		// shift items between new and old position up by 1
		err = db.Model(&model.Card{}).Where("list_id = ? AND position between ? and ?", list.ID, payload.Position, currentPosition).Update("position", gorm.Expr("position + 1")).Error
	} else {
		// shift items between new and old position down by 1
		err = db.Model(&model.Card{}).Where("list_id = ? AND position between ? and ?", list.ID, currentPosition, payload.Position).Update("position", gorm.Expr("position - 1")).Error
	}

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not update list positions", "data": nil})
	}

	card.Position = payload.Position

	err = db.Save(&card).Error

	if err != nil {
		// rollback position update on error
		db.Model(&card).Where("position = ?", currentPosition).Update("position", payload.Position)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card positions updated", "data": nil})
}
