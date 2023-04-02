package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateBoardList ... Create a new list for a board
// @Summary Create a new list for a board
// @Description create a new list for a board
// @Tags Boards
// @Accept json
// @Param id path string true "Board ID"
// @Param list_attrs body model.ListUserInput true "List attributes"
// @Success 201 {object} model.List
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /board/{id}/list [post]
func CreateBoardList(c *fiber.Ctx) error {
	list, err := services.CreateBoardList(c)

	if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "A new list has been created", "data": list})
}

// UpdateBoardListPosition ... Update list position on the board
// @Summary Update list position on the board
// @Description update list position on the board
// @Tags Boards
// @Accept json
// @Param id path string true "Board ID"
// @Param list path string true "List ID"
// @Param position body model.ListPositionInput true "List position"
// @Success 200 {object} model.List
// @Success 304 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /board/{id}/list/{list} [put]
func UpdateBoardListPosition(c *fiber.Ctx) error {
	list, err := services.UpdateBoardListPosition(c)

	if err != nil && strings.Contains(err.Error(), "List position not modified") {
		return c.Status(304).JSON(fiber.Map{"status": "success", "message": "List position not modified", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "List positions updated", "data": list})
}

// DeleteBoardList ... Delete a list from board
// @Summary Delete a list from board
// @Description delete a list from board
// @Tags Boards
// @Param id path string true "Board ID"
// @Param list path string true "List ID"
// @Success 200 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /board/{id}/list/{list} [delete]
func DeleteBoardList(c *fiber.Ctx) error {
	err := services.DeleteBoardList(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "List deleted", "data": nil})
}
