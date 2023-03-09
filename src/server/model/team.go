package model

type Team struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Members []*User `gorm:"many2many:user_teams;"`
	Boards  []*Board
}
