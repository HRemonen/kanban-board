package model

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string
	Lists       []List `gorm:"ForeignKey:BoardID"`
	TeamID      uuid.UUID
	Team        Team `gorm:"ForeignKey:TeamID"`
	OwnerID     uuid.UUID
	Owner       User `gorm:"ForeignKey:OwnerID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
