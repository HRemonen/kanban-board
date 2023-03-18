package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(20);default:'Initial card';"`
	Description string    `gorm:"type:varchar(100);"`
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
	Title       string `json:"title" validate:"required,ascii,gte=3, lte=20"`
	Description string `json:"description" validate:"ascii,lte=100"`
	Status      string `json:"status" validate:"alpha,gte=1,lte=10"`
	Label       string `json:"label" validate:"alpha,gte=1,lte=10"`
}

type CardPositionInput struct {
	Position uint `json:"position" validate:"numeric,gte=0"`
}

type CardStatusInput struct {
	Status uint `json:"status" validate:"required,alpha,lte=10"`
}

type CardLabelInput struct {
	Label uint `json:"label" validate:"required,alpha,lte=10"`
}
