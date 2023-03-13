package handlers

import (
	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllBoards(c *fiber.Ctx) error {
	db := database.DB.Db
	var boards []model.APIBoard

	db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&boards)

	if len(boards) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Boards not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Boards found", "data": boards})
}

func GetSingleBoard(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.APIBoard

	boardID := c.Params("id")

	db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Boards found", "data": board})
}

func CreateBoard(c *fiber.Ctx) error {
	db := database.DB.Db
	user, err := utils.ExtractUser(c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	payload := new(model.BoardUserInput)

	err = c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	newBoard := model.Board{
		Name:        payload.Name,
		Description: payload.Description,
		UserID:      user.ID,
	}

	err = db.Create(&newBoard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a new board", "data": err.Error()})
	}

	newList := model.List{
		Name:     "In progress",
		Position: 1,
		BoardID:  newBoard.ID,
	}

	err = db.Create(&newList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a initial list for the board", "data": err.Error()})
	}

	newCard := model.Card{
		Title:       "Initial card",
		Description: "You can create cards here",
		Position:    1,
		ListID:      newList.ID,
	}

	err = db.Create(&newCard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create an initial card for the list", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Board has been created", "data": newBoard})
}
