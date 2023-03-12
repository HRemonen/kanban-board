package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(20);default:'user';"`
	Photo     string    `gorm:"default:'default.png';"`
	Verified  bool      `gorm:"default:false;"`
	Provider  string    `gorm:"default:'local';"`
	Boards    []*Board  `gorm:"ForeignKey:UserID;references:ID;"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	user.ID = uuid.New()

	return nil
}

type RegisterUserInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginData struct {
	Token string       `json:"token" binding:"required"`
	User  UserResponse `json:"user" binding:"required"`
}

type APIUser struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name   string    `gorm:"type:varchar(100);not null"`
	Email  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Boards []*Board  `gorm:"ForeignKey:UserID;references:ID;"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Provider  string    `json:"provider,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Verified  bool      `json:"verified,omitempty"`
	Boards    []Board   `gorm:"many2many:boards;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		Verified:  user.Verified,
		Photo:     user.Photo,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
