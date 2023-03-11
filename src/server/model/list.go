package model

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"not null"`
	Position  uint      `gorm:"not null"`
	Cards     []Card    `gorm:"ForeignKey:ListID"`
	BoardID   uuid.UUID
	Board     Board `gorm:"ForeignKey:BoardID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
