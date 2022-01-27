package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ExtID   string  `json:"extID"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
	Picture string  `json:"picture"`
	PIN     string  `json:"-"`
}
