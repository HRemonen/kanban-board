package handlers

import (
	"fmt"

	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/HRemonen/kanban-board/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetAllBoards(c *fiber.Ctx) error {
	db := database.DB.Db
	var boards []model.Board

	db.Preload(clause.Associations).Find(&boards)

	if len(boards) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Boards not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Boards found", "data": boards})
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

	newTeam := model.Team{
		Name: "Team, " + payload.Name,
	}

	fmt.Println(&newTeam)

	err = db.Create(&newTeam).Error

	db.Model(&newTeam).Association("Users").Append(&user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a new team for board", "data": err.Error()})
	}

	newBoard := model.Board{
		Name:        payload.Name,
		Description: payload.Description,
		UserID:      user.ID,
		TeamID:      newTeam.ID,
	}

	err = db.Create(&newBoard).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create a new board", "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Board has been created", "data": newBoard})
}
