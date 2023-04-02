package services

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllBoards(c *fiber.Ctx) ([]model.APIBoard, error) {
	db := database.DB.Db
	var boards []model.APIBoard

	err := db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&boards).Error

	return boards, err
}

func GetSingleBoard(c *fiber.Ctx) (model.APIBoard, error) {
	db := database.DB.Db
	boardID := c.Params("id")

	var board model.APIBoard

	err := db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&board, "id = ?", boardID).Error

	return board, err
}

func CreateBoard(c *fiber.Ctx) (model.Board, error) {
	db := database.DB.Db

	var board model.Board

	user, err := utils.ExtractUser(c)

	if err != nil {
		return board, errors.New("Could not get user")
	}

	payload := new(model.BoardUserInput)

	c.BodyParser(payload)

	validate := utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return board, err
	}

	board = model.Board{
		Name:        payload.Name,
		Description: payload.Description,
		UserID:      user.ID,
	}

	err = db.Create(&board).Error

	return board, err
}

func DeleteBoard(c *fiber.Ctx) error {
	db := database.DB.Db
	boardID := c.Params("id")

	var board model.Board

	user, err := utils.ExtractUser(c)

	if err != nil {
		return errors.New("Could not get user")
	}

	err = db.Model(&model.Board{}).Preload("Lists.Cards").Find(&board, "id = ?", boardID).Error

	if board.ID == uuid.Nil {
		return errors.New("Board not found")
	}

	if board.UserID != user.ID {
		return errors.New("Unauthorized action")
	}

	// Loop through the board list's cards and delete each card
	for _, column := range board.Lists {
		for _, card := range column.Cards {
			if err := db.Delete(&card).Error; err != nil {
				return errors.New("Could not delete board cards")
			}
		}
	}

	err = db.Delete(&board.Lists).Error

	if err != nil {
		return errors.New("Could not delete board lists")
	}

	err = db.Delete(&board).Error

	return err
}
