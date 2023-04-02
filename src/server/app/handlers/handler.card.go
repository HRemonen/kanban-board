package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
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
	card, err := services.GetSingleCard(c)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Card not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card found", "data": card})
}

// CreateListCard ... Create a new card for a list
// @Summary Create a new card for a list
// @Description create a new card for a list
// @Tags Cards
// @Accept json
// @Param id path string true "List ID"
// @Param card_attrs body model.CardUserInput true "Card attributes"
// @Success 201 {object} model.Card
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /list/{id}/card [post]
func CreateListCard(c *fiber.Ctx) error {
	card, err := services.CreateListCard(c)

	if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "A new card has been created", "data": card})
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

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
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

// DeleteListCard ... Delete a card from list
// @Summary Delete a card from list
// @Description delete a card from list
// @Tags Cards
// @Param id path string true "List ID"
// @Param card path string true "card ID"
// @Success 200 {object} object
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Router /list/{id}/card/{list} [delete]
func DeleteListCard(c *fiber.Ctx) error {
	db := database.DB.Db
	var list model.List

	listID := c.Params("id")

	db.Find(&list, "id = ?", listID)

	if list.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "List not found", "data": nil})
	}

	var card model.Card

	cardID := c.Params("card")

	db.Find(&card, "id = ?", cardID)

	if list.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "List not found", "data": nil})
	} else if list.ID != card.ListID {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	}

	err := db.Select(clause.Associations).Delete(&card).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete card", "data": nil})
	}

	db.Model(&model.Card{}).Where("list_id = ? AND position > ?", list.ID, card.Position).Update("position", gorm.Expr("position - 1"))

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card deleted", "data": nil})
}
