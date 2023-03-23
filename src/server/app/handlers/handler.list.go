package handlers

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
// @Failure 500 {object} object
// @Router /board/{id}/list [post]
func CreateBoardList(c *fiber.Ctx) error {
	db := database.DB.Db
	var board model.Board

	boardID := c.Params("id")

	db.Model(&board).Preload("Lists").Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	payload := new(model.ListUserInput)

	err := c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	validate := utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
	}

	var currentPosition uint

	db.Model(&model.List{}).Select("COALESCE(MAX(position), 0)").Where("board_id = ?", board.ID).Row().Scan(&currentPosition)

	if len(board.Lists) == 0 {
		currentPosition = 0
	} else {
		currentPosition++
	}

	newList := model.List{
		Name:     payload.Name,
		Position: currentPosition,
		BoardID:  board.ID,
	}

	err = db.Create(&newList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a initial list for the board", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "A new list has been created", "data": newList})
}

// UpdateBoardListPosition ... Update list position on the board
// @Summary Update list position on the board
// @Description update list position on the board
// @Tags Boards
// @Accept json
// @Param id path string true "Board ID"
// @Param list path string true "List ID"
// @Param position body model.ListPositionInput true "List position"
// @Success 200 {object} object
// @Failure 404 {object} object
// @Failure 422 {object} object
// @Failure 500 {object} object
// @Router /board/{id}/list/{list} [put]
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

	validate := utils.NewValidator()

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": utils.ValidatorErrors(err)})
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
		err = db.Model(&model.List{}).Where("board_id = ? AND position between ? and ?", board.ID, payload.Position, currentPosition).Update("position", gorm.Expr("position + 1")).Error
	} else {
		// shift items between new and old position down by 1
		err = db.Model(&model.List{}).Where("board_id = ? AND position between ? and ?", board.ID, currentPosition, payload.Position).Update("position", gorm.Expr("position - 1")).Error
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

// DeleteBoardList ... Delete a list from board
// @Summary Delete a list from board
// @Description delete a list from board
// @Tags Boards
// @Param id path string true "Board ID"
// @Param list path string true "List ID"
// @Success 200 {object} object
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Router /board/{id}/list/{list} [delete]
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
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete list", "data": nil})
	}

	db.Model(&model.List{}).Where("board_id = ? AND position > ?", board.ID, list.Position).Update("position", gorm.Expr("position - 1"))

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "List deleted", "data": nil})
}