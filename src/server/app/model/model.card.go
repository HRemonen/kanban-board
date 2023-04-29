package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(255);default:'Initial card';"`
	Description string    `gorm:"type:text;"`
	Position    uint      `gorm:"type:integer;default:0;"`
	Status      string    `gorm:"type:varchar(10);default:'open';"`
	Label       string    `gorm:"type:varchar(10);"`
	ListID      uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (card *Card) BeforeCreate(*gorm.DB) error {
	card.ID = uuid.New()

	return nil
}

type CardUserInput struct {
	Title       string `json:"title" validate:"required,ascii,gte=3,lte=255"`
	Description string `json:"description" validate:"omitempty,ascii"`
}

type CardPositionInput struct {
	Position uint `json:"position" validate:"numeric,gte=0"`
}

type UpdateCard struct {
	Title       string `json:"title" validate:"omitempty,ascii,gte=3,lte=255"`
	Description string `json:"description" validate:"omitempty,ascii"`
	Status      string `json:"status" validate:"omitempty,alpha,lte=10"`
	Label       string `json:"label" validate:"omitempty,alpha,lte=10"`
}
