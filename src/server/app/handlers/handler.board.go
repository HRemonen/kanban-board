package handlers

import (
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

// GetAllBoards ... Get all boards
// @Summary Get all boards
// @Description get all boards
// @Tags Boards
// @Success 200 {array} model.APIBoard
// @Failure 404 {object} object
// @Router /board [get]
func GetAllBoards(c *fiber.Ctx) error {
	db := database.DB.Db
	var boards []model.APIBoard

	db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&boards)

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
	db := database.DB.Db
	var board model.APIBoard

	boardID := c.Params("id")

	db.Model(&model.Board{}).Preload("User").Preload("Lists.Cards").Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
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
// @Failure 500 {object} object
// @Router /board [post]
func CreateBoard(c *fiber.Ctx) error {
	db := database.DB.Db
	user, err := utils.ExtractUser(c)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	payload := new(model.BoardUserInput)

	err = c.BodyParser(payload)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": nil})
	}

	err = validate.Struct(payload)

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Validation of the input failed", "data": nil})
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
		BoardID: newBoard.ID,
	}

	err = db.Create(&newList).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a initial list for the board", "data": err.Error()})
	}

	newCard := model.Card{
		ListID: newList.ID,
	}

	err = db.Create(&newCard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create an initial card for the list", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Board has been created", "data": newBoard})
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
func DeleteBoardByID(c *fiber.Ctx) error {
	db := database.DB.Db
	user, err := utils.ExtractUser(c)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not fetch user", "data": nil})
	}

	var board model.Board

	boardID := c.Params("id")

	db.Model(&model.Board{}).Preload("Lists.Cards").Find(&board, "id = ?", boardID)

	if board.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Board not found", "data": nil})
	}

	if board.UserID != user.ID {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized action", "data": nil})
	}

	// Loop through the board list's cards and delete each card
	for _, column := range board.Lists {
		for _, card := range column.Cards {
			if err := db.Delete(&card).Error; err != nil {
				return err
			}
		}
	}

	err = db.Delete(&board.Lists).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete board lists", "data": nil})
	}

	err = db.Delete(&board).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete board", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Board deleted", "data": nil})
}
