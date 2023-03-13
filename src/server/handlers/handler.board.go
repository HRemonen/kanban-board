package handlers

import (
	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Board found", "data": board})
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

func CreateBoardList(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.Board

	boardID := c.Params("id")

	db.Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	payload := new(model.ListUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	var currentPosition uint

	db.Model(&model.List{}).Select("COALESCE(MAX(position), 1)").Row().Scan(&currentPosition)

	newList := model.List{
		Name:     payload.Name,
		Position: currentPosition + 1,
		BoardID:  board.ID,
	}

	err = db.Create(&newList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a initial list for the board", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "A new list has been created", "data": newList})
}

func UpdateBoardListPosition(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.Board

	boardID := c.Params("id")

	db.Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	payload := new(model.ListPositionInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	var list model.List

	listID := c.Params("list")

	db.Find(&list, "id = ?", listID)

	currentPosition := list.Position

	if payload.Position == currentPosition {
		return c.Status(304).JSON(fiber.Map{"status": "success", "message": "Position not modified", "data": nil}) // nothing to update
	}

	if payload.Position < currentPosition {
		// shift items between new and old position up by 1
		err = db.Model(&model.List{}).Where("position between ? and ?", payload.Position, currentPosition).Update("position", gorm.Expr("position + 1")).Error
	} else {
		// shift items between new and old position up by 1
		err = db.Model(&model.List{}).Where("position between ? and ?", currentPosition, payload.Position).Update("position", gorm.Expr("position - 1")).Error
	}

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not update list positions", "data": nil})
	}

	list.Position = payload.Position

	err = db.Save(&list).Error

	if err != nil {
		// rollback position update on error
		db.Model(&list).Where("position = ?", currentPosition).Update("position", payload.Position)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "List positions updated", "data": nil})
}

func DeleteBoardList(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.Board

	boardID := c.Params("id")

	db.Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	var list model.List

	listID := c.Params("list")

	db.Find(&list, "id = ?", listID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "List not found", "data": nil})
	} else if board.ID != list.BoardID {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	}

	err := db.Select(clause.Associations).Delete(&list).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete list", "data": nil})
	}

	db.Model(&model.List{}).Where("position > ?", list.Position).Update("position", gorm.Expr("position - 1"))

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "List deleted", "data": nil})
}
