package helpers

import (
	"fmt"

	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"gorm.io/gorm"
)

func SeedTestUsers(db *gorm.DB) error {
	hash, _ := utils.HashPassword("salainensalasana")

	users := []model.User {
		{Username: "Alice", Email: "alice@example.com", Password: hash},
		{Username: "Bob", Email: "bob@example.com", Password: hash},
		{Username: "Charlie", Email: "charlie@example.com", Password: hash},
	}

	err := db.Create(&users).Error
	if err != nil {
		return fmt.Errorf("failed to seed test data: %w", err)
	}

	return nil
}

func ClearTestUsers(db *gorm.DB) error {
	db.Exec("SET CONSTRAINTS ALL DEFERRED")
	err := db.Exec("TRUNCATE TABLE users CASCADE").Error
	if err != nil {
		return fmt.Errorf("failed to clear test user data: %w", err)
	}

	return nil
}

func SeedTestBoards(db *gorm.DB) error {
	hash, _ := utils.HashPassword("salainen")

	user := model.User{
		Name: "",
        Username: "Kaalebbi",
        Email: "kale@kaalebbi.fi",
        Password: hash,
	}

	err := db.Create(&user).Error

	if err != nil {
		return fmt.Errorf("failed to seed test user data: %w", err)
	}

	boards := []model.Board {
		{Name: "Project Alpha", Description: "Planning and execution of Alpha project", UserID: user.ID},
		{Name: "Marketing Campaign",Description: "Ideas and tasks for our upcoming marketing campaign", UserID: user.ID},
		{Name: "Team Building Activities", Description: "Planning and organizing team building activities for the company",  UserID: user.ID},
		{Name: "Product Development", Description: "Tracking progress and tasks for our product development team",  UserID: user.ID},
		{Name: "Bug Tracking",Description: "Tracking and fixing bugs for our software development projects", UserID: user.ID},
	}

	err = db.Create(&boards).Error
	if err != nil {
		return fmt.Errorf("failed to seed test data: %w", err)
	}

	return nil
}

func ClearTestBoards(db *gorm.DB) error {
	db.Exec("SET CONSTRAINTS ALL DEFERRED")
	err := db.Exec("TRUNCATE TABLE boards CASCADE").Error
	if err != nil {
		return fmt.Errorf("failed to clear test board data: %w", err)
	}

	return nil
}