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
	ListID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (card *Card) BeforeCreate(*gorm.DB) error {
	card.ID = uuid.New()

	return nil
}
