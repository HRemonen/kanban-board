package helpers

import (
	"fmt"

	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedTestUsers(db *gorm.DB) error {
	hash, _ := utils.HashPassword("salainensalasana")

	users := []model.User{
		{ID: uuid.MustParse("817c0ae8-f2ee-45c8-9e8f-d3bb2b42335d"), Username: "Alice", Email: "alice@example.com", Password: hash},
		{ID: uuid.MustParse("d8f22bf2-3550-43cf-af53-85b71d69742c"), Username: "Bob", Email: "bob@example.com", Password: hash},
		{ID: uuid.MustParse("b4ce4645-6da6-46d2-abb5-26aa55885d76"), Username: "Charlie", Email: "charlie@example.com", Password: hash},
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
		ID:       uuid.MustParse("b9f56152-25e3-49a0-812e-a40c1ad431eb"),
		Name:     "",
		Username: "Kaalebbi",
		Email:    "kale@kaalebbi.fi",
		Password: hash,
	}

	err := db.Create(&user).Error

	if err != nil {
		return fmt.Errorf("failed to seed test user data: %w", err)
	}

	boards := []model.Board{
		{ID: uuid.MustParse("a70d4aa4-d913-4b85-abbe-86271aaba147"), Name: "Project Alpha", Description: "Planning and execution of Alpha project", UserID: user.ID},
		{ID: uuid.MustParse("7a02f5d7-75aa-46b7-a698-c073ce49b12f"), Name: "Marketing Campaign", Description: "Ideas and tasks for our upcoming marketing campaign", UserID: user.ID},
		{ID: uuid.MustParse("4075b27e-ba24-400a-8f61-c2bc51fbf8cc"), Name: "Team Building Activities", Description: "Planning and organizing team building activities for the company", UserID: user.ID},
		{ID: uuid.MustParse("bceb0bc4-aa05-430f-b543-cfd808d61427"), Name: "Product Development", Description: "Tracking progress and tasks for our product development team", UserID: user.ID},
		{ID: uuid.MustParse("2ae97c29-46a4-4342-91f7-d552920b6578"), Name: "Bug Tracking", Description: "Tracking and fixing bugs for our software development projects", UserID: user.ID},
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
