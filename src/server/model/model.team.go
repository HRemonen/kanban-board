package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"not null"`
	Users     []*User   `gorm:"many2many:team_users"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (team *Team) BeforeCreate(*gorm.DB) error {
	team.ID = uuid.New()

	return nil
}
