package model

import "time"

type List struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Position  uint   `gorm:"not null"`
	Cards     []Card
	BoardID   uint  `gorm:"not null"`
	Board     Board `gorm:"foreignKey:BoardID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
