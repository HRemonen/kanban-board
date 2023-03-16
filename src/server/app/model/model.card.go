package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(100);default:'Card title';"`
	Description string    `gorm:"type:varchar(100);"`
	Position    uint      `gorm:"type:integer;not null;"`
	Status      string    `gorm:"type:varchar(10);default:'open';"`
	Label       string    `gorm:"type:varchar(10);;"`
	ListID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (card *Card) BeforeCreate(*gorm.DB) error {
	card.ID = uuid.New()

	return nil
}

type CardUserInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Label       string `json:"label"`
}

type CardPositionInput struct {
	Position uint `json:"position" binding:"required"`
}

type CardStatusInput struct {
	Status uint `json:"status" binding:"required"`
}

type CardLabelInput struct {
	Label uint `json:"label" binding:"required"`
}
