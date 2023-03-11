package handlers

import (
	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateBoard(c *fiber.Ctx) error {
	db := database.DB.Db
	user, err := utils.ExtractUser(c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	userResponse := model.FilteredResponse(&user)

	payload := new(model.BoardUserInput)

	err = c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	newBoard := model.Board{
		Name:        payload.Name,
		Description: payload.Description,
		UserID:      userResponse.ID,
		User:        userResponse,
	}

	err = db.Create(&newBoard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a new board", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Board has been created", "data": newBoard})
}

func GetAllBoards(c *fiber.Ctx) error {
	db := database.DB.Db
	var boards []model.Board

	db.Find(&boards)

	if len(boards) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Boards not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Boards found", "data": boards})
}
