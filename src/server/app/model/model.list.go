package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:varchar(100);default:'List title';"`
	Position  uint      `gorm:"type:integer;not null;"`
	Cards     []Card    `gorm:"ForeignKey:ListID;references:ID;"`
	BoardID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (list *List) BeforeCreate(*gorm.DB) error {
	list.ID = uuid.New()

	return nil
}

type ListUserInput struct {
	Name string `json:"name" binding:"required"`
}

type ListPositionInput struct {
	Position uint `json:"position" binding:"required"`
}
