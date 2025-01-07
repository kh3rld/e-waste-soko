package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;not null"`
	Password  string    `gorm:"not null"`
	UserType  string    `gorm:"size:20;not null"`
	Points    []Tokens  `gorm:"foreignKey:UserID"`
	WasteType []Waste   `gorm:"foreignKey:BroughtByID"`
}
