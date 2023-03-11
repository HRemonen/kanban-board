package model

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"not null"`
	Description string
	Position    uint `gorm:"not null"`
	ListID      uuid.UUID
	List        List `gorm:"ForeignKey:ListID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
