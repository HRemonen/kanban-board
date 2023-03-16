package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(20);default:'Card title';"`
	Description string    `gorm:"type:varchar(100);"`
	Position    uint      `gorm:"type:integer;not null;"`
	Status      string    `gorm:"type:varchar(10);default:'open';"`
	Label       string    `gorm:"type:varchar(10);"`
	ListID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (card *Card) BeforeCreate(*gorm.DB) error {
	card.ID = uuid.New()

	return nil
}

type CardUserInput struct {
	Title       string `json:"title" binding:"required, alpha, lte=20"`
	Description string `json:"description" binding:"lte=100, alphanum"`
	Status      string `json:"status" binding:"alpha, lte=10"`
	Label       string `json:"label" binding:"alpha, lte=10"`
}

type CardPositionInput struct {
	Position uint `json:"position" binding:"required, numeric, gte=1"`
}

type CardStatusInput struct {
	Status uint `json:"status" binding:"required, alpha, lte=10"`
}

type CardLabelInput struct {
	Label uint `json:"label" binding:"required, alpha, lte=10"`
}
