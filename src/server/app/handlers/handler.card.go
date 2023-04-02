package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
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
// @Success 200 {object} model.Card
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /list/{id}/card/{card} [put]
func UpdateListCardPosition(c *fiber.Ctx) error {
	card, err := services.UpdateListCardPosition(c)

	if err != nil && strings.Contains(err.Error(), "Card position not modified") {
		return c.Status(304).JSON(fiber.Map{"status": "success", "message": "Card position not modified", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card positions updated", "data": card})
}

// DeleteListCard ... Delete a card from list
// @Summary Delete a card from list
// @Description delete a card from list
// @Tags Cards
// @Param id path string true "List ID"
// @Param card path string true "card ID"
// @Success 200 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /list/{id}/card/{list} [delete]
func DeleteListCard(c *fiber.Ctx) error {
	err := services.DeleteListCard(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Card deleted", "data": nil})
}
