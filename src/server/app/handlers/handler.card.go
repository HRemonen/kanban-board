package handlers

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
