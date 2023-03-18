package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `gorm:"type:varchar(20);default:'Todo';"`
	Position uint      `gorm:"type:integer;default:0;"`
	Cards    []Card    `gorm:"ForeignKey:ListID;references:ID;"`
	BoardID  uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (list *List) BeforeCreate(*gorm.DB) error {
	list.ID = uuid.New()

	return nil
}

type ListUserInput struct {
	Name string `json:"name" validate:"required,ascii,gte=1,lte=20"`
}

type ListPositionInput struct {
	Position uint `json:"position" validate:"numeric,gte=0"`
}
