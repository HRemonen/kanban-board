package model

import "github.com/google/uuid"

type Team struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string    `gorm:"not null"`
}
