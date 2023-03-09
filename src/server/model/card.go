package model

import "time"

type Card struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Position    uint `gorm:"not null"`
	ListID      uint `gorm:"not null"`
	List        List `gorm:"foreignKey:ListID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
