package model

import "time"

type Board struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Lists       []List
	TeamID      uint
	Team        Team
	OwnerID     uint
	Owner       User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
