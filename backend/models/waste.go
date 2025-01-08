package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Waste struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text;not null"`
	Condition   string    `gorm:"size:100;not null"`
	BroughtByID uuid.UUID `gorm:"not null"`
	BroughtBy   User      `gorm:"foreignKey:BroughtByID"`
}
