package helpers

import (
	"fmt"

	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"gorm.io/gorm"
)

func SeedTestUsers(db *gorm.DB) error {
	hash, _ := utils.HashPassword("salainensalasana")

	users := []model.User{
		{Name: "Alice", Email: "alice@example.com", Password: hash},
		{Name: "Bob", Email: "bob@example.com", Password: hash},
		{Name: "Charlie", Email: "charlie@example.com", Password: hash},
	}

	err := db.Create(&users).Error
	if err != nil {
		return fmt.Errorf("failed to seed test data: %w", err)
	}

	return nil
}

func ClearTestUsers(db *gorm.DB) error {
	err := db.Exec("TRUNCATE TABLE users").Error
	if err != nil {
		return fmt.Errorf("failed to clear test data: %w", err)
	}

	return nil
}
