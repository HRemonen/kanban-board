package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Board struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Description string    `gorm:"type:text;"`
	UserID      uuid.UUID
	Lists       []List `gorm:"ForeignKey:BoardID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (board *Board) BeforeCreate(*gorm.DB) error {
	board.ID = uuid.New()

	return nil
}

type APIBoard struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;"`
	UserID      uuid.UUID
	Lists       []List `gorm:"ForeignKey:BoardID;references:ID;"`
}

type BoardUserInput struct {
	Name        string `json:"name" validate:"required,ascii,gte=3,lte=255"`
	Description string `json:"description" validate:"omitempty,ascii"`
}
