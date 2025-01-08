package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	Amount int
	UserID uuid.UUID `gorm:"not null"`
}
