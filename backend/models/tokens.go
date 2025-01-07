package models

import (
	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	Amount int
}
