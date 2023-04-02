package handlers

import (
	"strings"

	"github.com/HRemonen/kanban-board/app/services"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllBoards ... Get all boards
// @Summary Get all boards
// @Description get all boards
// @Tags Boards
// @Success 200 {array} model.APIBoard
// @Failure 404 {object} object
// @Router /board [get]
func GetAllBoards(c *fiber.Ctx) error {
	boards, err := services.GetAllBoards(c)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	if len(boards) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Boards not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Boards found", "data": boards})
}

// GetSingleBoard ... Get a single board
// @Summary Get a single board
// @Description get a single board
// @Tags Boards
// @Param id path string true "Board ID"
// @Success 200 {object} model.APIBoard
// @Failure 404 {object} object
// @Router /board/{id} [get]
func GetSingleBoard(c *fiber.Ctx) error {
	board, err := services.GetSingleBoard(c)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Board found", "data": board})
}

// CreateBoard ... Create a new board
// @Summary Create a new board
// @Description create a new board
// @Tags Boards
// @Accept json
// @Param board_attrs body model.BoardUserInput true "Board attributes"
// @Success 201 {object} model.APIBoard
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Router /board [post]
func CreateBoard(c *fiber.Ctx) error {
	board, err := services.CreateBoard(c)

	if err != nil && strings.Contains(err.Error(), "Key:") {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the inputs failed", "data": utils.ValidatorErrors(err)})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Board has been created", "data": board})
}

// DeleteBoardByID ... Delete a board by ID
// @Summary Delete a board by ID
// @Description delete a board by ID
// @Tags Boards
// @Param id path string true "Board ID"
// @Success 200 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /board/{id} [delete]
func DeleteBoard(c *fiber.Ctx) error {
	err := services.DeleteBoard(c)

	if err != nil && strings.Contains(err.Error(), "Unauthorized action") {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	} else if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Board deleted", "data": nil})
}
