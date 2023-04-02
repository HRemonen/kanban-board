package services

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateBoardList(c *fiber.Ctx) (model.List, error) {
	db := database.DB.Db
	var board model.Board
	var list model.List

	payload := new(model.ListUserInput)

	boardID := c.Params("id")

	err := db.Model(&board).Preload("Lists").Find(&board, "id = ?", boardID).Error

	if err != nil {
		return list, errors.New("Board not found")
	}

	c.BodyParser(payload)

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return list, err
	}

	var currentPosition uint

	db.Model(&model.List{}).Select("COALESCE(MAX(position), 0)").Where("board_id = ?", board.ID).Row().Scan(&currentPosition)

	if len(board.Lists) == 0 {
		currentPosition = 0
	} else {
		currentPosition++
	}

	list = model.List{
		Name:     payload.Name,
		Position: currentPosition,
		BoardID:  board.ID,
	}

	err = db.Create(&list).Error

	return list, err
}

func UpdateBoardListPosition(c *fiber.Ctx) (model.List, error) {
	db := database.DB.Db
	var board model.Board
	var list model.List

	payload := new(model.ListPositionInput)

	boardID := c.Params("id")

	err := db.Model(&board).Preload("Lists").Find(&board, "id = ?", boardID).Error

	if err != nil {
		return list, errors.New("Board not found")
	}

	c.BodyParser(payload)

	var validate = utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return list, err
	}

	listID := c.Params("list")

	err = db.Find(&list, "id = ?", listID).Error

	if err != nil {
		return list, errors.New("List not found")
	}

	if list.BoardID != board.ID {
		return list, errors.New("Unauthorized action")
	}

	currentPosition := list.Position

	if payload.Position == currentPosition {
		return list, errors.New("List position not modified") // nothing to update
	}

	if payload.Position < currentPosition {
		// shift items between new and old position up by 1
		err = db.Model(&model.List{}).Where("board_id = ? AND position between ? and ?", board.ID, payload.Position, currentPosition).Update("position", gorm.Expr("position + 1")).Error
	} else {
		// shift items between new and old position down by 1
		err = db.Model(&model.List{}).Where("board_id = ? AND position between ? and ?", board.ID, currentPosition, payload.Position).Update("position", gorm.Expr("position - 1")).Error
	}

	if err != nil {
		return list, errors.New("Could not update list positions")
	}

	list.Position = payload.Position

	err = db.Save(&list).Error

	return list, err
}

func DeleteBoardList(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.Board
	var list model.List

	boardID := c.Params("id")

	err := db.Find(&board, "id = ?", boardID).Error

	if err != nil {
		return errors.New("Board not found")
	}

	listID := c.Params("list")

	err = db.Find(&list, "id = ?", listID).Error

	if err != nil {
		return errors.New("List not found")
	}

	if list.BoardID != board.ID {
		return errors.New("Unauthorized action")
	}

	err = db.Select(clause.Associations).Delete(&list).Error

	if err != nil {
		return errors.New("Failed to delete list")
	}

	err = db.Model(&model.List{}).Where("board_id = ? AND position > ?", board.ID, list.Position).Update("position", gorm.Expr("position - 1")).Error

	return err
}
